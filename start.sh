#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)
node /app/.output/server/index.mjs > output.log 2>&1