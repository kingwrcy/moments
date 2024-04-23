#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

ENV_FILE="/app/.env"
export $(grep -v '^#' "$ENV_FILE" | grep -v '^$' | xargs)

node /app/.output/server/index.mjs > output.log 2>&1