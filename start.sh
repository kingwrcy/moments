#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)


CONFIG_FILE="/app/config"

if [ -f "$CONFIG_FILE" ]; then
    while IFS='=' read -r key value; do
        if [[ -z "$key" ]]; then
            continue
        fi

        if [[ $key == \#* ]]; then
            continue
        fi

        if [[ -z "$value" ]]; then
            continue
        fi

        key=$(echo "$key" | xargs)
        value=$(echo "$value" | xargs)

        # 检查环境变量是否已经存在，存在则忽略
        if [ -z "$(printenv "$key")" ]; then
            export "$key"="$value"
            echo "Set ENV: $key => Value: $value"
        else
            echo "环境变量 $key 已经存在，忽略"
        fi
    done < "$CONFIG_FILE"
else
    echo "配置文件 $CONFIG_FILE 不存在，忽略读取并导出环境变量的过程"
fi

echo $NUXT_PUBLIC_MOMENTS_TOOLBAR_ENABLE_DOUBAN
node /app/.output/server/index.mjs > output.log 2>&1