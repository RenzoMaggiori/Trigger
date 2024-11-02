#!/bin/sh

if [ -f /shared/client.apk ]; then
  cp /shared/client.apk /app/public/client.apk
  echo "client.apk copied to /app/public"
else
  echo "client.apk not found in /shared"
fi
exec "$@"
