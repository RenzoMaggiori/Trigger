#!/bin/sh

# This SCRIPT is ONLY to be used in the DOCKERFILE

./scripts/wait-for-it.sh db:5432 5432 '-- echo "Postgres is ready"'
./api/api -port 8000 -env-path ./api/.env
