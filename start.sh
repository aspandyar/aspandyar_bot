#!/bin/sh

set -e

echo "add env data"

echo "SERVER_ADDRESS=0.0.0.0:8080" > "/app/.env"

echo "start the app"
exec "$@"
