# Collab API Editor

A real-time, collaborative OpenAPI design platform ("Figma for OpenAPI").

## Features
- Real-time collaborative editing of OpenAPI specs (YAML/JSON)
- Schema validation in real time (OpenAPI 3.0+)
- Inline commenting on specific fields
- Version history and diffs
- Role-based permissions (Owner, Editor, Viewer)
- OAuth login (Google)
- Integration with a Rust-based Raft KV store (session metadata)

## Tech Stack
- **Frontend:** React, TypeScript, Tailwind, CodeMirror/Monaco
- **Backend:** Go, Echo, JWT, OAuth2, WebSocket
- **Database:** Postgres, Redis (pub/sub + cache)
- **Infra:** Docker, Kubernetes, AWS
- **Raft KV Store:** Rust-based microservice

## Monorepo Structure
```
/ (root)
  backend/    # Go API server
  frontend/   # React app (coming soon)
  shared/     # Shared types (coming soon)
  infra/      # Docker, k8s, scripts
```

## Getting Started
1. Clone the repo
2. Copy `backend/.env.example` to `backend/.env` and fill in secrets
3. Build and run the backend:
   ```sh
   cd backend
   go mod tidy
   go run ./cmd/main.go
   ```
4. (Frontend coming soon)

## License
MIT
