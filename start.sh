#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

if [ ! -f "/app/data/config.json" ]; then
    cp "/app/config.json" "/app/data/config.json"
    echo "配置文件不存在,启用默认配置"
fi
node /app/.output/server/index.mjs > output.log 2>&1