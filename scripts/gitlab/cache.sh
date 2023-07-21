#!/bin/bash
# 该脚本用于超过一定次数删除缓存

# 判断当前目录下是否存在 tag.txt 文件
if test -e "tag.txt"; then
  echo "tag.txt 文件存在"
  echo ./tag.txt
else
  # 如果不存在，则创建 tag.txt 文件
  echo "创建 tag.txt 文件"
  echo >"tag.txt" 0
fi

# 查看缓存列表
echo "查看缓存列表"
ls "$CI_PROJECT_DIR"/app
ls "$CI_PROJECT_DIR"/vendor

# 文件的值赋值给一个tag环境变量
echo "文件的值赋值给一个tag环境变量"
export tag=$(cat ./tag.txt)

# 运行一次Job tag自增一次
echo "查看tag值"
echo $tag

# 如果缓存存在超过定义的次数则删除并重置tag为0次
#if [ "$TAG" -gt 0 ]; then
if (("$tag" == 0)); then
  echo "删除第7次产生的缓存"
  rm -rf "$CI_PROJECT_DIR"/app
  rm -rf "$CI_PROJECT_DIR"/vendor
  rm -rf "$CI_PROJECT_DIR"/.cahce
  echo "重置tag值"
  tag=0
fi

echo "tag自增"
tag=$(cat ./tag.txt)
tag=$((tag + 1))
echo "缓存tag值"
echo $tag >./tag.txt