#!/bin/bash

docker compose down
docker compose rm
docker compose build
docker compose up --no-start
docker compose start frontend
docker compose logs -f