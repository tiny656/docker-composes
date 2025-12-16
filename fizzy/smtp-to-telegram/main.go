package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/emersion/go-message"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

// Config holds the application configuration
type Config struct {
	SMTPPort         string
	SMTPUsername     string
	SMTPPassword     string
	TelegramBotToken string
	TelegramChatID   string
}

func loadConfig() (*Config, error) {
	cfg := &Config{
		SMTPPort:         getEnv("SMTP_PORT", "2525"),
		SMTPUsername:     os.Getenv("SMTP_USERNAME"),
		SMTPPassword:     os.Getenv("SMTP_PASSWORD"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatID:   os.Getenv("TELEGRAM_CHAT_ID"),
	}

	if cfg.TelegramBotToken == "" {
		return nil, errors.New("TELEGRAM_BOT_TOKEN is required")
	}
	if cfg.TelegramChatID == "" {
		return nil, errors.New("TELEGRAM_CHAT_ID is required")
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// TelegramClient handles sending messages to Telegram
type TelegramClient struct {
	botToken string
	chatID   string
	client   *http.Client
}

func NewTelegramClient(botToken, chatID string) *TelegramClient {
	return &TelegramClient{
		botToken: botToken,
		chatID:   chatID,
		client:   &http.Client{},
	}
}

func (t *TelegramClient) SendMessage(text string) error {
	// Telegram message limit is 4096 characters
	if len(text) > 4000 {
		text = text[:4000] + "\n\n... (message truncated)"
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.botToken)

	payload := map[string]interface{}{
		"chat_id":    t.chatID,
		"text":       text,
		"parse_mode": "HTML",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := t.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", string(body))
	}

	return nil
}

// Backend implements smtp.Backend
type Backend struct {
	config   *Config
	telegram *TelegramClient
}

func (b *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		config:   b.config,
		telegram: b.telegram,
	}, nil
}

// Session implements smtp.Session and smtp.AuthSession
type Session struct {
	config   *Config
	telegram *TelegramClient
	from     string
	to       []string
	authed   bool
}

// AuthMechanisms returns available auth mechanisms - implements smtp.AuthSession
func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

// Auth handles authentication - implements smtp.AuthSession
func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		// If no credentials configured, allow all
		if s.config.SMTPUsername == "" && s.config.SMTPPassword == "" {
			s.authed = true
			return nil
		}
		// Validate credentials
		if username != s.config.SMTPUsername || password != s.config.SMTPPassword {
			return errors.New("invalid credentials")
		}
		s.authed = true
		return nil
	}), nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.from = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.to = append(s.to, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	log.Printf("Receiving email from %s to %v", s.from, s.to)

	// Read and parse the email
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read email data: %w", err)
	}
	log.Printf("Email data received, size: %d bytes", len(data))

	// Parse the email
	msg, err := message.Read(bytes.NewReader(data))
	if err != nil {
		log.Printf("Warning: failed to parse email: %v", err)
		// Send raw content if parsing fails
		text := fmt.Sprintf("📧 <b>New Email</b>\n\nFrom: %s\nTo: %s\n\n---\n%s",
			escapeHTML(s.from),
			escapeHTML(strings.Join(s.to, ", ")),
			escapeHTML(string(data)))
		return s.telegram.SendMessage(text)
	}

	// Extract headers
	header := msg.Header
	subject := header.Get("Subject")
	from := header.Get("From")
	to := header.Get("To")

	if from == "" {
		from = s.from
	}
	if to == "" {
		to = strings.Join(s.to, ", ")
	}

	// Extract body
	body := extractBody(msg)

	// Format message for Telegram
	text := fmt.Sprintf("📧 <b>New Email</b>\n\n<b>From:</b> %s\n<b>To:</b> %s\n<b>Subject:</b> %s\n\n---\n%s",
		escapeHTML(from),
		escapeHTML(to),
		escapeHTML(subject),
		escapeHTML(body))

	if err := s.telegram.SendMessage(text); err != nil {
		log.Printf("Failed to send to Telegram: %v", err)
		return err
	}
	log.Println("Email successfully forwarded to Telegram")
	return nil
}

func (s *Session) Reset() {
	s.from = ""
	s.to = nil
}

func (s *Session) Logout() error {
	return nil
}

// extractBody extracts the text body from an email message
func extractBody(msg *message.Entity) string {
	mediaType, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	if err != nil {
		// Try to read as plain text
		body, _ := io.ReadAll(msg.Body)
		return string(body)
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		mr := multipart.NewReader(msg.Body, params["boundary"])
		var textBody, htmlBody string

		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			partType := part.Header.Get("Content-Type")
			content, _ := io.ReadAll(part)

			if strings.HasPrefix(partType, "text/plain") {
				textBody = string(content)
			} else if strings.HasPrefix(partType, "text/html") {
				htmlBody = string(content)
			}
		}

		// Prefer text/plain over text/html
		if textBody != "" {
			return textBody
		}
		if htmlBody != "" {
			return stripHTML(htmlBody)
		}
		return ""
	}

	body, _ := io.ReadAll(msg.Body)

	if strings.HasPrefix(mediaType, "text/html") {
		return stripHTML(string(body))
	}

	return string(body)
}

// stripHTML removes HTML tags and decodes common HTML entities
func stripHTML(html string) string {
	// Remove script and style elements
	re := regexp.MustCompile(`(?is)<(script|style)[^>]*>.*?</\1>`)
	html = re.ReplaceAllString(html, "")

	// Remove HTML tags
	re = regexp.MustCompile(`<[^>]+>`)
	text := re.ReplaceAllString(html, "")

	// Decode common HTML entities
	text = strings.ReplaceAll(text, "&nbsp;", " ")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&#39;", "'")

	// Clean up whitespace
	re = regexp.MustCompile(`\n\s*\n`)
	text = re.ReplaceAllString(text, "\n\n")
	text = strings.TrimSpace(text)

	return text
}

// escapeHTML escapes HTML special characters for Telegram HTML mode
func escapeHTML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

func main() {
	log.Println("Starting SMTP to Telegram service...")

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	telegram := NewTelegramClient(config.TelegramBotToken, config.TelegramChatID)

	backend := &Backend{
		config:   config,
		telegram: telegram,
	}

	server := smtp.NewServer(backend)
	server.Addr = ":" + config.SMTPPort
	server.Domain = "localhost"
	server.AllowInsecureAuth = true

	authDisabled := config.SMTPUsername == "" && config.SMTPPassword == ""
	log.Printf("SMTP server listening on port %s", config.SMTPPort)
	if authDisabled {
		log.Println("Authentication is disabled (no credentials configured)")
	} else {
		log.Println("Authentication is enabled")
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
