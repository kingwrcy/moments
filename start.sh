#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

ENV_FILE="/app/.env"

# 读取 .env 文件并导出环境变量，跳过已存在的环境变量
grep -v -e '^#' -e '^$' "$ENV_FILE" | while read -r line; do
    key=$(echo "$line" | cut -d '=' -f 1)
    if [ -z "${!key+x}" ]; then
        # 导出整个键值对
        export "$line"
    fi
done

node /app/.output/server/index.mjs > output.log 2>&1