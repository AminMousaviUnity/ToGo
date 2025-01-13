DB_URL=postgres://togo_user:togo_password@localhost:5432/togo_db?sslmode=disable
MIGRATE=migrate -database "$(DB_URL)" -path migrations

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
