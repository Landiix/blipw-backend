# Blipw

Учебный клон Twitter: Go API + React UI в одном репо.

```
blipw-backend/   (этот репозиторий)
├── cmd/server/     # API, порт :8081
├── internal/       # handlers, service, repository, …
├── web/            # фронтенд (Vite + React)
├── docker-compose.yml
└── init.sql
```

## Backend

```bash
docker compose up -d
cp .env.example .env   # подставь значения из compose
go run ./cmd/server
```

## Frontend

```bash
cd web && npm install && npm run dev
```

Подробнее: [`web/README.md`](./web/README.md).
