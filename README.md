# Docker-Compose Collection

This repository contains a collection of Docker-Compose configurations for various applications and services. Each directory contains the necessary files to deploy and manage the respective application using Docker-Compose.

## Directory Structure

- **acme.sh**: Docker-Compose setup for [acme.sh](https://github.com/acmesh-official/acme.sh), a shell script for managing Let's Encrypt certificates.
- **anyproxy**: Docker-Compose configuration for [AnyProxy](https://github.com/alibaba/anyproxy), a proxy server for capturing and modifying HTTP/HTTPS requests.
- **azure-functions-dotnet-isolated**: Docker-Compose setup for Azure Functions with .NET isolated worker.
- **azure-openai-proxy**: Docker-Compose configuration for an Azure OpenAI proxy server.
- **chatgpt-next-web**: Docker-Compose setup for a ChatGPT web interface.
- **free-games-claimer**: Docker-Compose configuration for a service that automatically claims free games.
- **memos**: Docker-Compose setup for [Memos](https://github.com/justmemos/memos), a self-hosted note-taking application.
- **miniflux**: Docker-Compose configuration for [Miniflux](https://miniflux.app/), a minimalist and opinionated feed reader.
- **nginx-proxy-manager**: Docker-Compose setup for [Nginx Proxy Manager](https://nginxproxymanager.com/), a simple and powerful reverse proxy.
- **qiandao**: Docker-Compose configuration for [Qiandao](https://github.com/qd-today/qd), an open-source sign-in system.
- **siyuan**: Docker-Compose configuration for [SiYuan](https://github.com/siyuan-note/siyuan), a local-first personal knowledge management system.
- **singbox**: Docker-Compose configuration for [sing-box](https://github.com/SagerNet/sing-box), a universal proxy platform supporting multiple protocols.
- **snell**: Docker-Compose setup for [Snell](https://github.com/surge-networks/snell), a proxy server.
- **stirling-pdf**: Docker-Compose configuration for [Stirling PDF](https://stirlingpdf.com/), an open-source PDF toolkit.
- **syncthing**: Docker-Compose configuration for [Syncthing](https://github.com/syncthing/syncthing), a continuous file synchronization program.
- **trojan-go**: Docker-Compose setup for [Trojan-Go](https://github.com/p4gefau1t/trojan-go), an unidentifiable mechanism that helps you bypass GFW.
- **vaultwarden**: Docker-Compose configuration for [Vaultwarden](https://github.com/dani-garcia/vaultwarden), an unofficial Bitwarden server implementation.
- **wanderer**: Docker-Compose setup for [Wanderer](https://github.com/Flomp/wanderer), a self-hosted trail database. Save your adventures.
- **watchtower**: Docker-Compose configuration for [Watchtower](https://containrrr.dev/watchtower/), a process for automating Docker container base image updates.
- **xiaogpt**: Docker-Compose setup for [XiaoGPT](https://github.com/yihong0618/xiaogpt), play ChatGPT and other LLM with Xiaomi AI Speaker.
- **xray**: Docker-Compose configuration for [Xray](https://github.com/XTLS/Xray-core), a platform for building proxies to bypass network restrictions.
- **zerotier**: Docker-Compose setup for [ZeroTier](https://www.zerotier.com/), a smart Ethernet switch for planet Earth.

## Usage

To use any of the Docker-Compose configurations, navigate to the respective directory and follow the instructions provided in the `README.md` file within that directory. Typically, the steps involve:

1. Cloning the repository:
   ```bash
   git clone https://github.com/yourusername/docker-composes.git
   cd docker-composes
   ```

2. Navigating to the desired directory:
   ```bash
   cd <directory-name>
   ```

3. Running Docker-Compose:
   ```bash
   docker-compose up -d
   ```

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes or additions.

## License

This repository is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

This repository aims to provide a comprehensive collection of Docker-Compose configurations to streamline the deployment and management of various applications and services. Feel free to explore and use the configurations as needed.