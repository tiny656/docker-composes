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
| [acme.sh](acme.sh/) | Automated Let's Encrypt certificate management using [acme.sh](https://github.com/acmesh-official/acme.sh) |

### Network & Proxy

| Service | Description |
|---------|-------------|
| [anyproxy](anyproxy/) | HTTP/HTTPS proxy server for capturing and modifying requests ([AnyProxy](https://github.com/alibaba/anyproxy)) |
| [nginx-proxy-manager](nginx-proxy-manager/) | Web-based reverse proxy management ([NPM](https://nginxproxymanager.com/)) |
| [singbox](singbox/) | Universal proxy platform with multi-protocol support ([sing-box](https://github.com/SagerNet/sing-box)) |
| [snell](snell/) | Lightweight encrypted proxy server ([Snell](https://github.com/surge-networks/snell)) |
| [trojan-go](trojan-go/) | Unidentifiable proxy mechanism ([Trojan-Go](https://github.com/p4gefau1t/trojan-go)) |
| [xray](xray/) | Proxy platform for bypassing network restrictions ([Xray-core](https://github.com/XTLS/Xray-core)) |
| [zerotier](zerotier/) | Software-defined networking ([ZeroTier](https://www.zerotier.com/)) |

### AI & LLM

| Service | Description |
|---------|-------------|
| [azure-openai-proxy](azure-openai-proxy/) | Proxy server for Azure OpenAI API |
| [chatgpt-next-web](chatgp-next-web/) | ChatGPT web interface |
| [xiaogpt](xiaogpt/) | ChatGPT integration for Xiaomi AI Speaker ([XiaoGPT](https://github.com/yihong0618/xiaogpt)) |

### Cloud & Serverless

| Service | Description |
|---------|-------------|
| [azure-functions-dotnet-isolated](azure-functions-dotnet-isolated/) | Azure Functions with .NET isolated worker |

### Productivity & Notes

| Service | Description |
|---------|-------------|
| [memos](memos/) | Self-hosted note-taking application ([Memos](https://github.com/usememos/memos)) |
| [siyuan](siyuan/) | Local-first personal knowledge management ([SiYuan](https://github.com/siyuan-note/siyuan)) |

### RSS & Content

| Service | Description |
|---------|-------------|
| [miniflux](miniflux/) | Minimalist feed reader ([Miniflux](https://miniflux.app/)) |

### File & Sync

| Service | Description |
|---------|-------------|
| [syncthing](syncthing/) | Continuous file synchronization ([Syncthing](https://syncthing.net/)) |

### Documents & Tools

| Service | Description |
|---------|-------------|
| [stirling-pdf](stirling-pdf/) | Self-hosted PDF toolkit ([Stirling PDF](https://stirlingpdf.com/)) |

### Automation & Utilities

| Service | Description |
|---------|-------------|
| [free-games-claimer](free-games-claimer/) | Automatically claim free games from various platforms |
| [qiandao](qiandao/) | Open-source sign-in system ([Qiandao](https://github.com/qd-today/qd)) |
| [watchtower](watchtower/) | Automated Docker container updates ([Watchtower](https://containrrr.dev/watchtower/)) |
| [fizzy](fizzy/) | Web application service |

### Security & Password Management

| Service | Description |
|---------|-------------|
| [vaultwarden](vaultwarden/) | Unofficial Bitwarden server implementation ([Vaultwarden](https://github.com/dani-garcia/vaultwarden)) |

### Lifestyle

| Service | Description |
|---------|-------------|
| [wanderer](wanderer/) | Self-hosted trail database for adventures ([Wanderer](https://github.com/Flomp/wanderer)) |

## Directory Structure

```
docker-composes/
├── acme.sh/
├── anyproxy/
├── azure-functions-dotnet-isolated/
├── azure-openai-proxy/
├── chatgp-next-web/
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
