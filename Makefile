generate:
	protoc -I=./ --go_out=../../../ --go-grpc_out=../../../ ./api/parcel_delivery.proto

compose-up:
	docker compose build && docker compose up

clean-img:
	docker image prune

docker_app_id := $(shell docker ps | grep parcel_delivery_test_app_1 | awk '{print $$1}')
restart_app:
	docker restart ${docker_app_id}
restart_app-%:


run:
	go run cmd/main.go

DB_HOST=localhost
DB_USER=postgres_user
DB_PASSWORD=password
DB_NAME=parcel_delivery
DB_PORT=5432

DB_DSN="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

db-dump:#??
	pg_dump ${DB_DSN} --schema-only --no-owner > ./db_schema.sql

db-reset:
	psql -c 'drop database ${DB_NAME} with (FORCE);' "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/postgres?sslmode=disable"	
	psql -c 'create database ${DB_NAME};' "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/postgres?sslmode=disable"
	DB_HOST=${DB_HOST} DB_USER=${DB_USER} DB_PASSWORD=${DB_PASSWORD} DB_NAME=${DB_NAME} DB_PORT=${DB_PORT} scripts/migrate.sh

jet:
	jet -source=postgres -dsn=${DB_DSN} -path=./internal/generated/ -ignore-tables goose_db_version


