#!/bin/bash

# 设置目标目录
target_directory="."

# 遍历目标目录及其子目录下的文件
find "$target_directory" -type f -name "gitea" | while read source_file; do
  # 提取目录名
  directory_name=$(dirname "$source_file")

  # 提取文件名（不包括路径）
  # file_name=$(basename "$source_file")

  # 新文件名为目录名
  new_file_name="$directory_name/gitbundle"

  # 重命名文件为新文件名
  mv "$source_file" "$new_file_name"

  echo "已将文件 '$source_file' 重命名为 '$new_file_name'"
done
