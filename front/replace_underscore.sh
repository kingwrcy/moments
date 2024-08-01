#!/bin/bash

# 定义要处理的目录
TARGET_DIR="/app/.output/public"

# 遍历所有以_开头的文件
find "$TARGET_DIR" -type f -name "_*" | while read -r file; do
    # 获取文件的目录和扩展名
    dir=$(dirname "$file")
    base=$(basename "$file")
    ext="${base##*.}"
    new_base=$(echo "$base" | sed 's/^_//') # 去掉开头的下划线

    # 生成一个随机字符串作为新文件名
    random_string=$(tr -dc 'a-zA-Z0-9' < /dev/urandom | head -c 8)
    new_file="${dir}/${random_string}.${ext}"

    # 重命名文件
    mv "$file" "$new_file"

    # 在整个public目录中替换旧文件名为新文件名
    find "$TARGET_DIR" -type f -exec sed -i "s|$base|$random_string.$ext|g" {} +

    echo "Renamed '$file' to '$new_file' and updated references."
done
