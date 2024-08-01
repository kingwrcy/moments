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

    # 使用 basename 提取文件名
    base_file=$(basename "$file")

    # 检查文件名的第一个字符是否为下划线
    if [ "${base_file:0:1}" = "_" ]; then
        new_file="a${base_file:1}"  # 替换开头的下划线为 'a'

        # 调试信息
        echo "Current file: $file"
        echo "New file name: $new_file"

        # 检查文件是否存在
        if [ -f "$file" ]; then
            # 重命名文件
            echo "Renaming $file to $new_file"
            mv "$file" "$dir/$new_file"  # 确保使用目录路径
        else
            echo "File $file does not exist."
        fi

        # 更新所有引用到这个文件的地方
        find "$dir" -type f -exec sed -i "s|$base_file|$new_file|g" {} +
    else
        echo "File $file does not start with an underscore, skipping."
    fi
done

echo "Replacement complete."
