include backend/.env
export

# Determine which DB URL to use based on ENV
ifeq ($(ENV),prod)
    DB_URL=$(DB_URL_PROD)
else
    DB_URL=$(DB_URL_DEV)
endif

# Database migration commands
.PHONY: migrate-up migrate-down migrate-status migrate-create db-reset

migrate-up:
	@echo "Running migrations on $(ENV) database..."
	cd backend && goose -dir migrations postgres "$(DB_URL)" up

migrate-down:
	@echo "Rolling back migration on $(ENV) database..."
	cd backend && goose -dir migrations postgres "$(DB_URL)" down

migrate-status:
	@echo "Migration status for $(ENV) database:"
	cd backend && goose -dir migrations postgres "$(DB_URL)" status

migrate-create:
	@read -p "Enter migration name: " name; \
	cd backend && goose -dir migrations create $$name sql

db-reset:
	@echo "⚠️  WARNING: Resetting $(ENV) database..."
	@read -p "Are you sure? (y/N): " confirm; \
	if [ "$$confirm" = "y" ]; then \
		cd backend && goose -dir migrations postgres "$(DB_URL)" reset && \
		goose -dir migrations postgres "$(DB_URL)" up; \
	else \
		echo "Cancelled."; \
	fi

# Backend - Server
.PHONY: backend-run backend-dev

backend-run:
	cd backend && go run cmd/server/main.go

backend-dev:
	cd backend && air

# Environment switching helpers
.PHONY: use-dev use-prod

use-dev:
	@echo "Switching to DEV database"
	@sed -i.bak 's/ENV=.*/ENV=dev/' backend/.env && rm backend/.env.bak

use-prod:
	@echo "Switching to PROD database"
	@sed -i.bak 's/ENV=.*/ENV=prod/' backend/.env && rm backend/.env.bak