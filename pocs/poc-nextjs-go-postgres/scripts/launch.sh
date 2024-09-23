#!/bin/bash

docker compose down
docker compose rm
docker compose build
docker compose up --no-start
docker compose start db
docker compose start api
docker compose start web
docker compose logs -f

