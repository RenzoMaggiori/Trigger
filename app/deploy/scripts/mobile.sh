#! /usr/bin/env bash

./scripts/api.sh
if [ $? -ne 0 ]; then
    exit 1
fi

docker compose build mobile
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose build mobile' failed."
    docker compose down
    exit 1
fi

docker compose up --no-start mobile
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose up --no-start mobile' failed."
    docker compose down
    exit 1
fi

docker compose start mobile
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose start mobile' failed."
    docker compose down
    exit 1
fi

