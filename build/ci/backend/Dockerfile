# Stage 1: Build Golang binary
FROM alpine:latest
MAINTAINER Lookeke
ENV VERSION 1.0

WORKDIR /app

# 拷贝当前目录下可以执行文件
COPY . .

RUN go mod download

# 编译成目标操作系统的可执行文件
RUN CC=x86_64-unknown-linux-gnu-gcc CXX=x86_64-unknown-linux-gnu-g++ GOOS=linux GOARCH=amd64 go build -o app .

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 4000

# 运行golang程序的命令
ENTRYPOINT ["/app"]
