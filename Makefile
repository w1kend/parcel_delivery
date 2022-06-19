generate:
	protoc -I=./ --go_out=../../../ --go-grpc_out=../../../ ./api/parcel_delivery.proto

up:
	docker compose build && docker compose up

envs := $(shell cat env.local)
run:
	${envs} go run cmd/main.go

DB_HOST=localhost
DB_USER=postgres_user
DB_PASSWORD=password
DB_NAME=parcel_delivery
DB_PORT=5432

DB_DSN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

db-reset:
	psql -c 'drop database if exists ${DB_NAME} with (FORCE);' "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/postgres?sslmode=disable"	
	psql -c 'create database ${DB_NAME};' "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/postgres?sslmode=disable"
	DB_HOST=${DB_HOST} DB_USER=${DB_USER} DB_PASSWORD=${DB_PASSWORD} DB_NAME=${DB_NAME} DB_PORT=${DB_PORT} scripts/migrate.sh

jet:
	jet -source=postgres -dsn=${DB_DSN} -path=./internal/generated/ -ignore-tables goose_db_version


