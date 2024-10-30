#! /usr/bin/env bash

./scripts/api.sh
if [ $? -ne 0 ]; then
    exit 1
fi

services="client_web client_mobile"

docker compose build $services
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose build $services' failed."
    exit 1
fi

docker compose up --no-start $services
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose up --no-start $services' failed."
    exit 1
fi

docker compose start $services
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose start $services' failed."
    exit 1
fi
