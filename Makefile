DB_URL=postgres://togo_user:togo_password@localhost:5432/togo_db?sslmode=disable
MIGRATE=migrate -database "$(DB_URL)" -path migrations

.PHONY: migrate-up migrate-down migrate-force migrate-create build ToGo down

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-force:
	$(MIGRATE) force VERSION

migrate-create:
	@read -p "Enter migration name: " name; \
	touch migrations/`date +"%Y%m%d%H%M%S"`_$$name.up.sql; \
	touch migrations/`date +"%Y%m%d%H%M%S"`_$$name.down.sql; \
	echo "Migration files created"

build:
	docker compose build

down:
	docker compose down

ToGo: build
	docker compose up --build -d
