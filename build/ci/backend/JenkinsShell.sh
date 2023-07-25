#!/bin/bash

go version

#go env -w GOPROXY=https://proxy.golang.com.cn,direct
#go get
#go test -v ./test
#go build -o app
#chmod +x app

echo "开始部署139.198.165.102节点"
echo "传输Golang文件除了不必要的文件的所有文件发送至主机root@139.198.165.102的/home/nginx/html/backend目录下"
sshpass rsync -av -e "ssh -o stricthostkeychecking=no" ./ --exclude={build,client,examples,tmp,vendor,.air.toml} root@139.198.165.102:/home/nginx/html/backend
echo "进入到主机root@139.198.165.102的/home/nginx/html/backend Golang部署script,构建镜像完成之后删除/home/nginx/html/web/backend目录"
sshpass ssh -o stricthostkeychecking=no root@139.198.165.102 'cd /home/nginx/html/backend && bash ./kubernetes/deploy/golang-app-deploy.yml'
echo "完成部署139.198.165.102节点"
