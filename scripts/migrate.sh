#!/usr/bin/env bash
while : 
do
  nc -z $DB_HOST $DB_PORT >/dev/null 2>&1
  success=$?

  if [[ $success -eq 0 ]]; then
    break
  else 
    >&2 echo "Postgres is unavailable on $DB_HOST:$DB_PORT"
    sleep 1
  fi
done

goose -dir=./migrations/ postgres "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" up