#! /usr/bin/env bash

./scripts/api.sh
if [ $? -ne 0 ]; then
    exit 1
fi

docker compose build web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose build web' failed."
    docker compose down
    exit 1
fi

docker compose up --no-start web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose up --no-start web' failed."
    docker compose down
    exit 1
fi

docker compose start web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose start web' failed."
    docker compose down
    exit 1
fi

