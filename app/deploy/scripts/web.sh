#! /usr/bin/env bash

./scripts/api.sh
if [ $? -ne 0 ]; then
    exit 1
fi

docker compose build client_web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose build client_web' failed."
    docker compose down
    exit 1
fi

docker compose up --no-start client_web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose up --no-start client_web' failed."
    docker compose down
    exit 1
fi

docker compose start client_web
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose start client_web' failed."
    docker compose down
    exit 1
fi

