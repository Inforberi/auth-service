# auth-service

Сервис аутентификации на Go (`go 1.25.1`) с PostgreSQL.

## Текущий этап

На текущем этапе реализован каркас сервиса:
- загрузка конфигурации из `.env`;
- инициализация структурированного логгера (`slog`);
- модуль подключения к PostgreSQL через `pgxpool` с `Ping` и настройками пула;
- первая миграция схемы БД для auth-домена (`users`, `user_identities`, `user_passwords`, `user_profiles`, `sessions`).

Важно: в `cmd/api/main.go` пока выполняются только загрузка конфига и инициализация логгера. Создание `App` и запуск HTTP/API еще не подключены.

## Стек

- Go 1.25.1+
- PostgreSQL 18 (через `docker-compose`)
- `pgx/v5`
- `cleanenv`
- `slog`
- `golang-migrate` (CLI для миграций)

## Структура

- `cmd/api/main.go` — точка входа.
- `internal/config/config.go` — конфигурация и DSN.
- `internal/logger/logger.go` — логгер (json/text, dev/prod поведение).
- `internal/infra/postgres/pool.go` — создание пула подключений PostgreSQL.
- `migrations/` — SQL-миграции.
- `make/migrate.mk` — команды миграций.

## Переменные окружения

Обязательные:
- `POSTGRES_HOST`
- `POSTGRES_PORT`
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_NAME_DB`

Опциональные:
- `APP_ENV` (default: `prod`)
- `LOG_LEVEL` (default: `info`)
- `LOG_FORMAT` (default: `json`)
- `PG_SSL_MODE` (default: `disable`)
- `PG_MAX_CONNS` (default: `20`)
- `PG_MIN_CONNS` (default: `2`)
- `PG_HEALTH_TIMEOUT` (default: `3s`)
- `PG_MAX_CONN_LIFETIME` (default: `30m`)
- `PG_MAX_CONN_IDLE_TIME` (default: `5m`)

## Быстрый старт

1. Заполнить `.env` обязательными переменными.
2. Поднять PostgreSQL:

```bash
docker compose up -d postgres
```

3. Установить migrate CLI (если не установлен):

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

4. Применить миграции:

```bash
make migrate-up
```

5. Запустить приложение:

```bash
go run ./cmd/api
```

## Доступные команды make

```bash
make help
make help-migrate
```
