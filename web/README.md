# web — Blipw frontend

Vite + React + TypeScript. API живёт в корне репо (`cmd/server`, порт `:8081`).

## Запуск

```bash
# из корня репо: подними Postgres + API
docker compose up -d
go run ./cmd/server

# в другом терминале:
cd web
npm install
npm run dev
```

Открой `http://localhost:5173`. Запросы `/api/*` проксируются на `http://localhost:8081` (см. `vite.config.ts`).

## Важно

- После login пользователь пока в `localStorage` — временно, пока нет JWT/session.
- `POST /api/tweets` пока шлёт `user_id` с клиента.
- Когда появится auth на бэкенде — см. комментарии в `src/api/client.ts`.
