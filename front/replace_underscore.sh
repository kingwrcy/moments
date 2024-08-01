#!/bin/sh

TARGET_DIR="/app/.output/public"
# 遍历当前目录及子目录中的所有以 _ 开头的文件
find "$TARGET_DIR" -type f -name '_*' | while read -r file; do
    # 获取文件的目录和新文件名
    dir=$(dirname "$file")
    new_file="${file/_/a}"

    # 重命名文件
    mv "$file" "$new_file"

    # 更新所有引用到这个文件的地方
    # 使用 sed 替换文件内容中的引用
    find "$dir" -type f -exec sed -i "s|$file|$new_file|g" {} +
done