#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

ENV_FILE="/app/.env"

grep -v -e '^#' -e '^$' "$ENV_FILE" | while IFS='=' read -r key value; do
    if [ -z "$key" ] || [ -z "$value" ]; then
        continue
    fi

    if ! declare -p "$key" &>/dev/null; then
        echo "Setting $key=$value"
        export "$key=$value"
    fi
done

node /app/.output/server/index.mjs > output.log 2>&1