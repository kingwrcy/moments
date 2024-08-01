#!/bin/sh

# 目标目录
TARGET_DIR="/app/.output/public/_nuxt"

# 列出目标目录中的文件
echo "Listing files in $TARGET_DIR:"
ls -l "$TARGET_DIR"

# 遍历目标目录中的所有以 _ 开头的文件
find "$TARGET_DIR" -type f -name '_*' | while read -r file; do
    # 获取文件的目录和新文件名
    dir=$(dirname "$file")
    new_file="${file/_/a}"

    # 检查文件是否存在
    if [ -f "$file" ]; then
        # 重命名文件
        echo "Renamed $file to $new_file"
        mv "$file" "$new_file"
        echo "Renamed $file to $new_file"
    else
        echo "File $file does not exist."
    fi

    # 更新所有引用到这个文件的地方
    # 使用 sed 替换文件内容中的引用
    find "$dir" -type f -exec sed -i "s|$file|$new_file|g" {} +
done

echo "Replacement complete."