# SMTP to Telegram

A lightweight SMTP server written in Go that receives emails and forwards them to Telegram.

## Features

- Receives emails via SMTP protocol
- Forwards email content to Telegram chat
- Supports PLAIN authentication
- Handles multipart emails (text/plain and text/html)
- Converts HTML to plain text
- Docker-ready with multi-stage build

## Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `SMTP_PORT` | SMTP server listening port | No | 2525 |
| `SMTP_USERNAME` | SMTP authentication username | No | - |
| `SMTP_PASSWORD` | SMTP authentication password | No | - |
| `TELEGRAM_BOT_TOKEN` | Telegram Bot API token | Yes | - |
| `TELEGRAM_CHAT_ID` | Target Telegram chat ID | Yes | - |

## Usage

### With Docker Compose (Recommended)

Create a `.env` file in the fizzy directory:

```env
TELEGRAM_BOT_TOKEN=your_bot_token_here
TELEGRAM_CHAT_ID=your_chat_id_here
```

Then run:

```bash
docker-compose up -d
```

### Standalone Docker

```bash
docker build -t smtp-to-telegram .
docker run -d \
  -p 2525:2525 \
  -e TELEGRAM_BOT_TOKEN=your_bot_token \
  -e TELEGRAM_CHAT_ID=your_chat_id \
  -e SMTP_USERNAME=user \
  -e SMTP_PASSWORD=pass \
  smtp-to-telegram
```

### Local Development

```bash
export TELEGRAM_BOT_TOKEN=your_bot_token
export TELEGRAM_CHAT_ID=your_chat_id
go run .
```

## Getting Telegram Credentials

### Bot Token

1. Open Telegram and search for `@BotFather`
2. Send `/newbot` and follow the instructions
3. Copy the API token provided

### Chat ID

1. Start a chat with your bot or add it to a group
2. Send a message to the bot/group
3. Visit `https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates`
4. Find the `chat.id` in the response

## Message Format

Emails are formatted as:

```
📧 New Email

From: sender@example.com
To: recipient@example.com
Subject: Email Subject

---
Email body content...
```

## License

MIT
