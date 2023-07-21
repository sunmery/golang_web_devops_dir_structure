#!/bin/bash
# 该脚本用于在主机上创建临时目录

# 需要创建的目录
dir=/home/nginx/html
# 列出该目录所有文件
ls -A $dir
# 获取上个指令的返回值, 为0则该目录存在
equal=$?
# 检查该目录是否有文件, 存在删除目录, 没有则创建
if [ "$equal" -eq 0 ]; then
  echo "已存在临时目录,正在删除"
  rm -rf /home/nginx/html/backend
fi
echo "正在创建临时目录"
mkdir -p /home/nginx/html/backend
