#!/bin/sh
envsubst < /etc/nginx/nginx.template.conf > /etc/nginx/nginx.conf
exec "$@"

