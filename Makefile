POSTGRESQL_URL ?= "postgres://user:pass@localhost:8003/artilight?sslmode=disable"
MIGRATIONS_PATH ?= "internal/common/database/migrations"

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ${MIGRATIONS_PATH} up

migrate-down:
	migrate -database -verbose ${POSTGRESQL_URL} -path ${MIGRATIONS_PATH} down

create-migration:
	migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(name)