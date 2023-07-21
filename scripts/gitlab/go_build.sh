#!/bin/bash
# 该脚本用于Go项目部署

echo "开始部署139.198.165.102节点"
echo "传输Golang文件除了web前端目录,.cache缓存,.ideaIDE缓存,tmp缓存目录以外的所有文件发送至主机root@139.198.165.102的/home/nginx/html/backend目录下"
sshpass -e rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={web,.cache,.idea,tmp,vendor,test,.air.toml,.dockerignore,.gitlab.ci.yml} root@139.198.165.102:/home/nginx/html/backend
echo "进入到主机root@139.198.165.102的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass -e ssh -o stricthostkeychecking=no root@139.198.165.102 'cd /home/nginx/html/backend && bash ./cmd/go_deploy.sh'
echo "完成部署139.198.165.102节点"

echo "开始部署192.168.0.158节点"
echo "传输Golang文件除了web前端目录,.cache缓存,.ideaIDE缓存,tmp缓存目录以外的所有文件发送至主机root@192.168.0.158的/home/nginx/html/backend目录下"
sshpass -e rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={web,.cache,.idea,tmp,vendor,test,.air.toml,.dockerignore,.gitlab.ci.yml} root@192.168.0.158:/home/nginx/html/backend
echo "进入到主机root@192.168.0.158的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass -e ssh -o stricthostkeychecking=no root@192.168.0.158 'cd /home/nginx/html/backend && bash ./cmd/go_deploy.sh'
echo "完成部署192.168.0.158节点"

echo "开始部署119.91.208.160节点"
echo "传输Golang文件除了web前端目录,.cache缓存,.ideaIDE缓存,tmp缓存目录以外的所有文件发送至主机root@119.91.208.160的/home/nginx/html/backend目录下"
sshpass -e rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={web,.cache,.idea,tmp,vendor,test,.air.toml,.dockerignore,.gitlab.ci.yml} root@119.91.208.160:/home/nginx/html/backend
echo "进入到主机root@1119.91.208.160的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass -e ssh -o stricthostkeychecking=no root@119.91.208.160 'cd /home/nginx/html/backend && bash ./cmd/go_deploy.sh'
echo "完成部署139.198.172.110节点"

echo "开始部署139.198.172.110节点"
echo "传输Golang文件除了web前端目录,.cache缓存,.ideaIDE缓存,tmp缓存目录以外的所有文件发送至主机root@139.198.172.110的/home/nginx/html/backend目录下"
sshpass -e rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={web,.cache,.idea,tmp,vendor,test,.air.toml,.dockerignore,.gitlab.ci.yml} root@139.198.172.110:/home/nginx/html/backend
echo "进入到主机root@139.198.172.110的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass -e ssh -o stricthostkeychecking=no root@139.198.172.110 'cd /home/nginx/html/backend && bash ./cmd/go_deploy.sh'
echo "完成部署139.198.172.110节点"
