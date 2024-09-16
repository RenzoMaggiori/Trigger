#!/bin/sh

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <HOST> <PORT> <CMD>"
    exit 1
fi

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

echo "Waiting for $host:$port..."
until nc -z "$host" $port; do
  >&2 echo "Service $host:$port is unavailable - sleeping"
  sleep 1
done

>&2 echo "Service $host:$port is ready - executing command"
exec $cmd
