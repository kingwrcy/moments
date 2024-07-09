#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

if [ ! -f "/app/data/config.json" ]; then
    cp "/app/config.json" "/app/data/config.json"
    echo "配置文件不存在,启用默认配置"
fi

if [ -z "${NUXT_JWT_KEY}" ]; then
  echo "建议设置NUXT_JWT_KEY变量,以免重启后需重新登录"
  export NUXT_JWT_KEY=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
fi

node /app/.output/server/index.mjs > output.log 2>&1
