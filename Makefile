DB_URL = "postgres://belovetech:@localhost:5432/gator"
GOOSE_CMD = goose
DB_TYPE = postgres
MIGRATION_DIR = ./sql/schema

SQLC_CMD = sqlc


.PHONY: migrate-up migrate-down create-migration generate-db


migrate-up:
	@echo "Applying migrations..."
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_TYPE) $(DB_URL) up


migrate-down:
	@echo "Rolling back last migration..."
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) $(DB_TYPE) $(DB_URL) down


create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Please provide a migration name using 'make create-migration name=<migration_name>'"; \
		exit 1; \
	fi
	@echo "Creating new migration..."
	$(GOOSE_CMD) -dir $(MIGRATION_DIR) create $(name) sql

generate-db:
	@echo "Generating database..."
	$(SQLC_CMD) generate
