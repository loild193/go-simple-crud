#!/usr/bin/env sh
set -e

echo "Migrating databases.."
go run /app/migrations/migrate.go
echo "Starting server..."

exec "$@"