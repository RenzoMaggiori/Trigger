#! /usr/bin/env bash

docker compose down
if [ $? -ne 0 ]; then
    echo "Error: 'docker compose down' failed."
    exit 1
fi

services="db auth user session action gmail sync settings github spotify"

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
