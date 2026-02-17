---
name: docker-compose
description: "Advanced Docker Compose: multi-service, scaling, profiles"
---
# Docker Compose Advanced

Use the **bash** tool for advanced Compose workflows. Modern Docker includes Compose as a plugin: use `docker compose` (with space), not `docker-compose`.

## Setup

1. **Check if installed:**
   ```bash
   command -v docker && docker --version && docker compose version
   ```

2. **Install:**
   ```bash
   # macOS
   brew install docker

   # Ubuntu / Debian (see https://docs.docker.com/engine/install/ubuntu/)
   sudo apt update && sudo apt install -y ca-certificates curl
   sudo install -m 0755 -d /etc/apt/keyrings
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
   sudo chmod a+r /etc/apt/keyrings/docker.gpg
   echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
   sudo apt update && sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
   ```

## Profiles & Environment
```bash
docker compose --profile debug up -d
docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
docker compose --env-file .env.staging up -d
```

## Scaling & Health
```bash
docker compose up -d --scale worker=3
docker compose ps --format json | jq '.[] | {Name, State, Health}'
docker compose top
```

## Networking
```bash
docker network ls
docker network inspect <network>
docker compose exec <service> ping <other-service>
```

## Tips
- Use profiles for dev/staging/prod separation
- Use multiple compose files for environment overrides
- Always check health status after scaling
