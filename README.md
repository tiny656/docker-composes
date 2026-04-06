# Docker Compose Collection

> A curated collection of Docker Compose configurations for self-hosted applications and services.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/docker-composes.git
cd docker-composes

# Navigate to the desired service
cd <service-name>

# Start the service
docker-compose up -d
```

## Services

### Security & Certificates

| Service | Description |
|---------|-------------|
| [acme.sh](acme.sh/) | Automated HTTPS certificate management with Azure DNS & Telegram notifications ([acme.sh](https://github.com/acmesh-official/acme.sh)) |

### Network & Proxy

| Service | Description |
|---------|-------------|
| [anyproxy](anyproxy/) | HTTP/HTTPS proxy server for intercepting and inspecting requests ([AnyProxy](https://github.com/alibaba/anyproxy)) |
| [nginx-proxy-manager](nginx-proxy-manager/) | Web-based reverse proxy management with GoAccess log analytics ([NPM](https://nginxproxymanager.com/)) |
| [singbox](singbox/) | Universal proxy platform with multi-protocol support ([sing-box](https://github.com/SagerNet/sing-box)) |
| [snell](snell/) | Lightweight encrypted proxy server with obfuscation support ([Snell](https://github.com/surge-networks/snell)) |
| [trojan-go](trojan-go/) | Proxy protocol disguising traffic as HTTPS, with Caddy for camouflage ([Trojan-Go](https://github.com/p4gefau1t/trojan-go)) |
| [xray](xray/) | Proxy server supporting Vmess, Vless, Trojan and other protocols ([Xray-core](https://github.com/XTLS/Xray-core)) |
| [zerotier](zerotier/) | Software-defined networking with ZeroNSD DNS resolution ([ZeroTier](https://www.zerotier.com/)) |

### AI & LLM

| Service | Description |
|---------|-------------|
| [azure-openai-proxy](azure-openai-proxy/) | Proxy server for Azure OpenAI API with Nginx HTTPS termination |
| [chatgpt-next-web](chatgp-next-web/) | ChatGPT web interface with customizable API backend ([ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web)) |
| [xiaogpt](xiaogpt/) | ChatGPT integration for Xiaomi AI Speaker with OpenAI API proxy ([XiaoGPT](https://github.com/yihong0618/xiaogpt)) |

### Cloud & Serverless

| Service | Description |
|---------|-------------|
| [azure-functions-dotnet-isolated](azure-functions-dotnet-isolated/) | Azure Functions with .NET isolated worker and Azurite local storage emulator |

### Productivity & Notes

| Service | Description |
|---------|-------------|
| [memos](memos/) | Privacy-focused self-hosted memo/notes application with Nginx reverse proxy ([Memos](https://github.com/usememos/memos)) |
| [siyuan](siyuan/) | Local-first personal knowledge management with Nginx reverse proxy ([SiYuan](https://github.com/siyuan-note/siyuan)) |

### RSS & Content

| Service | Description |
|---------|-------------|
| [miniflux](miniflux/) | Minimalist RSS feed reader with PostgreSQL backend ([Miniflux](https://miniflux.app/)) |

### File & Sync

| Service | Description |
|---------|-------------|
| [syncthing](syncthing/) | Continuous file synchronization with relay server for peer discovery ([Syncthing](https://syncthing.net/)) |

### Documents & Tools

| Service | Description |
|---------|-------------|
| [stirling-pdf](stirling-pdf/) | Self-hosted PDF toolkit for merge, split, OCR and more ([Stirling PDF](https://stirlingpdf.com/)) |

### Monitoring & Logging

| Service | Description |
|---------|-------------|
| [dozzle](dozzle/) | Real-time Docker container log viewer with file-based user auth ([Dozzle](https://dozzle.dev/)) |
| [watchtower](watchtower/) | Automated Docker container updates with Telegram notifications ([Watchtower](https://containrrr.dev/watchtower/)) |

### Automation & Utilities

| Service | Description |
|---------|-------------|
| [fizzy](fizzy/) | Self-hosted team chat platform with SMTP-to-Telegram bridge ([Fizzy/Campfire](https://github.com/basecamp/campfire)) |
| [free-games-claimer](free-games-claimer/) | Automatically claim free games from Epic Games with VNC interface ([free-games-claimer](https://github.com/vogler/free-games-claimer)) |
| [qiandao](qiandao/) | Automatic task scheduler for website check-ins with Redis cache ([Qiandao](https://github.com/qd-today/qd)) |

### Security & Password Management

| Service | Description |
|---------|-------------|
| [vaultwarden](vaultwarden/) | Self-hosted Bitwarden-compatible password manager with Nginx reverse proxy ([Vaultwarden](https://github.com/dani-garcia/vaultwarden)) |

### Lifestyle

| Service | Description |
|---------|-------------|
| [wanderer](wanderer/) | Self-hosted trail database with Meilisearch and PocketBase ([Wanderer](https://github.com/Flomp/wanderer)) |

## Directory Structure

```
docker-composes/
├── acme.sh/
├── anyproxy/
├── azure-functions-dotnet-isolated/
├── azure-openai-proxy/
├── chatgp-next-web/
├── dozzle/
├── fizzy/
├── free-games-claimer/
├── memos/
├── miniflux/
├── nginx-proxy-manager/
├── qiandao/
├── singbox/
├── siyuan/
├── snell/
├── stirling-pdf/
├── syncthing/
├── trojan-go/
├── vaultwarden/
├── wanderer/
├── watchtower/
├── xiaogpt/
├── xray/
└── zerotier/
```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or additions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
