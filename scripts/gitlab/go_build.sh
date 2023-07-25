#!/bin/bash
# 该脚本用于Go项目部署

echo "开始部署139.198.165.102节点"
echo "传输Golang文件除了不必要的文件的所有文件发送至主机root@139.198.165.102的/home/nginx/html/backend目录下"
sshpass -e rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={build,client,examples,tmp,vendor,.air.toml} root@139.198.165.102:/home/nginx/html/backend
echo "进入到主机root@139.198.165.102的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass -e ssh -o stricthostkeychecking=no root@139.198.165.102 'cd /home/nginx/html/backend && bash ./cmd/go_deploy.sh'
echo "完成部署139.198.165.102节点"
