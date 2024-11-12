#!/bin/sh

set -e

echo "add env data"

. /app/app.env

echo "start the app"
exec "$@"
