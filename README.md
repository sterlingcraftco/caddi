# Caddi 🛠️

**Caddi** is a CLI tool for managing Caddy server configuration files and blocks via JSON.

### Features

- List Caddy config blocks
- Show details of specific blocks
- Modify ports and upstreams
- (Optional) Push updated config to Caddy server via API

### Usage

```bash
# List all config blocks
caddi list

# Show config details
caddi show example.com

# Modify port or reverse proxy
caddi modify example.com --port :8081 --upstream localhost:3000
