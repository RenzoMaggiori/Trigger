#! /usr/bin/env bash

docker compose down
docker compose build
docker compose up --no-start
docker compose start db
docker compose start auth
docker compose start user
docker compose start session
docker compose start action
docker compose start gmail
docker compose start sync
docker compose start settings
docker compose start github

