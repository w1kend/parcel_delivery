generate:
	protoc -I=./ --go_out=../../../ --go-grpc_out=../../../ ./api/parcel_delivery.proto

compose-up:
	docker compose build && docker compose up

clean-img:
	docker image prune

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

jet:
	jet -source=postgres -dsn=${DB_DSN} -path=./internal/generated/


