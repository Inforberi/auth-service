MIGRATIONS_DIR := migrations
DATABASE_URL ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_NAME_DB)?sslmode=disable
VERBOSE ?= -verbose

.PHONY: help-migrate check-migrate migrate-create migrate-up migrate-down-1 migrate-down-all migrate-force migrate-version

help-migrate:
	@echo ""
	@echo "Команды миграций:"
	@echo "  make migrate-create     — создать новую миграцию"
	@echo "  make migrate-up         — применить все миграции"
	@echo "  make migrate-down-1     — откатить одну миграцию назад"
	@echo "  make migrate-down-all   — откатить ВСЕ миграции (ОПАСНО)"
	@echo "  make migrate-force      — принудительно установить версию (если база dirty)"
	@echo "  make migrate-version    — показать текущую версию миграций"
	@echo ""

check-migrate:
	@command -v migrate >/dev/null 2>&1 || { \
		echo "❌ Утилита migrate не найдена."; \
		echo "Установите командой:"; \
		echo "go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
		exit 1; \
	}

migrate-create: check-migrate
	@read -p "Введите имя миграции (например: init_auth): " name; \
	if [ -z "$$name" ]; then echo "❌ Имя миграции не должно быть пустым"; exit 1; fi; \
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" -seq "$$name"

migrate-up: check-migrate
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" $(VERBOSE) up

migrate-down-1: check-migrate
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" $(VERBOSE) down 1

migrate-down-all: check-migrate
	@read -p "⚠ Это откатит ВСЕ миграции. Введите YES для подтверждения: " ans; \
	if [ "$$ans" != "YES" ]; then echo "Отменено"; exit 1; fi; \
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" $(VERBOSE) down

migrate-force: check-migrate
	@read -p "Введите версию для force (например 1): " ver; \
	if [ -z "$$ver" ]; then echo "❌ Версия не должна быть пустой"; exit 1; fi; \
	read -p "Введите YES для подтверждения force версии $$ver: " ans; \
	if [ "$$ans" != "YES" ]; then echo "Отменено"; exit 1; fi; \
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" $(VERBOSE) force "$$ver"

migrate-version: check-migrate
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" version