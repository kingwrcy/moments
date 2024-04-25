#!/bin/sh

mkdir -p /app/data/upload
npx prisma migrate deploy

export VERSION=$(cat /app/version)

CONFIG_FILE="/app/config"

if [ -f "$CONFIG_FILE" ]; then
    # 使用进程替换来过滤掉注释行
    while IFS='=' read -r key value; do
        if [[ -z "$key" ]] || [[ $key == \#* ]]; then
            continue
        fi

        if [[ -z "$value" ]]; then
            continue
        fi

        # 检查环境变量是否已经存在，存在则忽略
        if [ -z "$(printenv "$key")" ]; then
            export "$key"="$value"
            echo "Set ENV: $key => Value: $value"
        else
            echo "环境变量 $key 已经存在，忽略"
        fi
    done < <(cat "$CONFIG_FILE" | grep -v '^#') 
else
    echo "配置文件 $CONFIG_FILE 不存在，忽略读取并导出环境变量的过程"
fi

node /app/.output/server/index.mjs > output.log 2>&1