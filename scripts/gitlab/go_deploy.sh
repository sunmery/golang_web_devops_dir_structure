#!/bin/bash
# 该脚本用于打包后端镜像

# 检测goback(.gitlab-ci.yml提供的golang打包后的二进制文件的名称)是否已在运行, 如果已在运行则停止该容器
# docker ps 列出Docker容器列表
# grep goback 搜索docker ps 列出Docker容器列表有没有goback这个字串
# awk '{print $12}' 从docker ps 列出Docker容器列表的第12列(容器名)
# docker stop goback 停止运行goback容器
# shellcheck disable=SC2046
echo "检测goback(.gitlab-ci.yml提供的golang打包后的二进制文件的名称)是否已在运行, 如果已在运行则停止该容器"
if [ $(docker ps | grep goback | awk '{print $12}') ]; then
  echo "goback容器已在运行, 正在停止该容器"
  docker stop goback
fi

# 如果goback容器存在 则删除该容器
# docker ps -a 列出全部容器
# docker ps -q 静默模式,只列出容器编号
# docker ps -aq 列出全部容器编号
# --filter name=goback 过滤出容器名称为goback的容器
# docker rm -f goback 删除goback容器
echo "如果goback容器存在 则删除该容器"
if [ $(docker ps -aq --filter name=goback) ]; then
  echo "goback容器存在, 正在删除该容器"
  docker rm -f goback
fi

# docker images 列出所有镜像
# docker rmi -f 强制删除
# docker rmi -f goback 强制删除goback镜像
echo "docker images 列出所有镜像"
if [ $(docker images | grep goback | awk '{print $1}') ]; then
  echo "goback镜像存在, 正在删除该镜像"
  docker rmi -f goback
fi

# 在Dockerfile当前目录的所有文件的镜像打包至标签为goback的二进制文件可运行的Docker镜像
# docker build --tag 镜像的名字及标签，通常 name:tag 或者 name 格式；可以在一次构建中为一个镜像设置多个标签。
echo "在Dockerfile当前目录的所有文件的镜像打包至标签为goback的二进制文件可运行的Docker镜像"
docker build --tag goback -f /home/nginx/html/backend/Dockerfile .

# docker run 运行Golang项目
# -d 后台运行容器，并返回容器ID
# -p 映射主机与容器的端口
# --name 容器名称
echo "docker run 运行Golang项目"
docker run -d -p 443:443 --name goback goback

echo "查看goback容器运行日志"
docker logs goback

echo "构建镜像完成, 删除目录"
rm -rf /home/nginx/html/backend
