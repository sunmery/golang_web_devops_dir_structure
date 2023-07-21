# Golang Web DevOps Dir

本项目是根据个人的喜好做的一个 Golang Web DevOps的项目目录结构, 包含一些示例

- 面向接口编程
- 后端使用 gin
- 鉴权使用 JWT
- 前端使用 React + Electron + Vite
- 数据库使用 GORM + Postgres + Redis
- 运维部署使用 Kubectl + Kubernetes 与 Dockerfile + docker-compose 或 gitlab + gitlab-Runner方式
- 命名:
    - 目录名全部小写
    - 文件名遵循蛇形命名法
    - 导出的函数遵循大驼峰命名法, 包内部私有字段以及不对外暴露的函数,方法,变量使用小驼峰命名法

## 目录结构

- assets 共享的静态资产(静态文件, 图像, 音视频...)
- build 打包目录
    - ci 存放`gitlab-ci.yml`等`CI`脚本
    - package 存放云,操作系统, Dockerfile, docker-compose 等
- client 存放客户端文件
- cmd 存放可执行文件, 非 sh,bat 脚本的指令文件
- docs 设计文档, 说明文件
- examples 存储应用程序示例目录
- internal 私有文件夹, 不对外共享
    - api 存放 `JSON`等文件数据
    - handler 处理 HTTP 请求
    - middleware 中间件, 例如 CORS, 拦截器
    - repository 数据库
    - router 路由服务
    - schema/model/dao 存放与数据库映射的类型与方法
    - Kubernetes 存放 k8s 部署的文件
- pkg 存储公共方法
    - config 对外开放的 config 配置, 例如数据库配置
        - ssl 存储 HTTPS证书
    - helper 工具函数
- scripts 存放脚本
    - migrations sql迁移文件
- test 测试文件夹

## Backend list

1. Web 服务器: [gin](https://github.com/gin-gonic/gin)
2. Web 跨域处理: [cors](https://github.com/gin-contrib/cors)
3. 日志: [slog](https://github.com/gookit/slog)
4. 环境变量: [godotenv](https://github.com/joho/godotenv)
5. 开发热重启: [air](https://github.com/cosmtrek/air)

## Frontend

### Cors

1. [React](https://react.dev/)
2. [Electron](https://electron-vite.github.io)
3. [Typescript](https://www.typescriptlang.org/)
4. [Vite](https://vitejs.dev/)
5. vite-plugin-electron

### Web 项目工程化

1. husky
2. commitlint
3. ESLint
    1. lint-staged
    2. eslint-config-alloy
    3. eslint-import-resolver-typescript
    4. eslint-plugin-import
    5. eslint-plugin-react
    6. eslint-plugin-react-hooks
    7. eslint-plugin-react-hooks-addons
    8. eslint-plugin-react-refresh
    9. eslint-plugin-simple-import-sort
    10. eslint-plugin-typescript-enum
4. Stylelint+
    1. stylelint-config-recess-order
    2. stylelint-config-standard
    3. stylelint-config-standard-scss

#### 安装

1. 安装依赖

   ```shell
   p i -D lint-staged husky @commitlint/cli @commitlint/config-conventional eslint eslint-config-alloy @typescript-eslint/parser eslint-plugin-import eslint-import-resolver-typescript eslint-import-resolver-typescript eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-react-hooks-addons eslint-plugin-simple-import-sort eslint-plugin-typescript-enum lint-staged typescript
   ```

2. 先创建`.git`仓库
   ```git
   git init
   ```

3. `package.json`添加`prepare`指令
   ```json
   { 
       "scripts": { 
           "prepare": "husky install"
        }
   }
   ```

4. prepare脚本会在执行npm install之后自动执行。执行npm install安装完项目依赖后会执行 husky install命令。
   ```shell
   pnpm i
   ```

   (可选)命令行
   ```shell
   npm set-script prepare "husky install" && npm run prepare
   ```

5. 添加`git hooks`, 创建一条 `pre-commit hook`

   ```shell
   npx husky add .husky/pre-commit "pnpm lint"
   ```

   执行该命令后，会看到.husky/目录下新增了一个名为`pre-commit`的`shell`脚本。之后执行`git commit`
   命令时会先触发`pre-commit`这个脚本。
   脚本内容:
   ```
   #!/bin/sh 
   . "$(dirname "$0")/_/husky.sh" 
   pnpm lint
   ```

6. 修改`pre-commit`脚本, 改成项目所需的指令, 例如`ESLint`,`Stylelint`,`commitlint`检查
   ```
   #!/usr/bin/env sh
   . "$(dirname -- "$0")/_/husky.sh"
   
   tsc --noEmit && pnpm stylelint:fix && pnpm dlx lint-staged
   ```

7. 添加`commitlint`规范
   ```shell
   npx husky add .husky/commit-msg 'npx --no-install commitlint --edit "$1"'
   ```

## Database

- ORM管理器: [GORM](https://gorm.io/gorm)
- 数据库驱动: [Postgres](https://gorm.io/driver/postgres)
- Redis: [Go-Redis](https://github.com/redis/go-redis)

- 迁移数据库 [github](https://github.com/golang-migrate/migrate)
    - [迁移库](https://ggithub.com/golang-migrate/migrate/v4)
    - [Postgres数据库](https://github.com/golang-migrate/migrate/v4/database/postgres)
    - [sql文件读取](https://ggithub.com/golang-migrate/migrate/v4/source/file)

## (可选) Deployment

1. kubernetes
2. Dockerfile
3. docker-compose
4. gitlab-runner

### kubernetes

1. 进入根目录的`kubernetes`的 `deployment`目录下
2. 部署 Golang 应用 在`kubernetes`集群运行:
   ```shell
   kubectl apply -f golang-app-deployment.yaml
   ```

3. 部署 Web 应用 在`kubernetes`集群运行:
   ```shell
   kubectl apply -f golang-app-deployment.yaml
   ```

4. 部署 Golang 应用 在`kubernetes`集群运行:
   ```shell
   kubectl apply -f web-app-deployment.yaml
   ```

5. 部署 Postgres 数据库 在`kubernetes`集群运行:
   ```shell
   kubectl apply -f postgres-deployment.yaml
   ```

### gitlab-runner

1. 可根据此篇文章安装[Gitlab](https://juejin.cn/post/7205954522722828344)
2. 可根据此篇文章安装[gitlab-runner](https://juejin.cn/editor/drafts/7257866619204124732)
3. 在根目录创建`.gitlab-ci.yml`文件,参考`examples/ci/backend/.gitlab-ci.yml`
   后端部署文件或`examples/ci/frontend/.gitlab-ci.yml`前端部署文件针对你的需求个性化

## Used

### 前置条件

1. `Node.js` 14+ 版本以上运行环境
2. `Golang` 1.16+ 版本以上运行环境
3. `Postgres` PG 数据库环境. 项目自带`Postgres`部署环境, `Kubernetes`
   部署方式在根目录的`Kubernetes`/`deplotment`/`postgres-deployment.yaml`, Docker 部署方式在根目录的`cmd`/`Dockerfile`
4. (可选) `Kubernetes`(K8S) 版本推荐 1.20 以上, 使用`Containerd` 容器运行时(CRI)
5. (可选) 至少 3台Linux服务器, 1个 `master` 主控节点, 2 个`Node`/`Worker` 节点. 推荐 2个 `master` 主控节点,
   4个`Node`/`Worker` 节点


### 编写配置文件

- 数据库配置: 填写在根目录的`pkg` -> `config` -> `db.development.yaml`开发配置与`db.production.yaml`生产模式配置文件
- (可选) 将 `HTTPS`服务器的 `crt` 与 `key`证书的文件复制至 `pkg` -> `config` -> `ssl`目录即可运行`TLS`服务
- (可选) 修改`pkg` -> `config` -> `jwt.yaml` JWT 秘钥

### 安装后端服务

1. `gin`
    ```shell
    go get -u github.com/gin-gonic/gin
    ```

2. `gin-CORS`

    ```shell
    go get github.com/gin-contrib/cors
    ```

3. `slog` 日志库
    ```shell
    go get github.com/gookit/slog
    ```

4. `godotenv` 环境变量
    ```shell
    go get github.com/joho/godotenv
    ```

5. `air` 开发热重启
    ```shell
    go install github.com/cosmtrek/air@latest
    ```

6. sql迁移.

    ```shell
    go get github.com/golang-migrate/migrate/v4
    go get github.com/golang-migrate/migrate/v4/database/postgres
    go get github.com/golang-migrate/migrate/v4/source/file
    ```

7. `go-redis` Redis 客户端
    ```shell
    go get github.com/redis/go-redis/v9
    ```
   
### 安装前端服务

1. 进入到项目的前端目录`Client`.
    ```shell
    cd client
    ```

2. 任选你喜欢的一个`包管理器`安装依赖, 例如`npm`,`pnpm`,`yarn`, 只要你安装了它.
    - npm:
    ```shell
    npm install
    ```

    - pnpm:
    ```shell
    pnpm i
    ```

    - yarn:
   ```shell
   yarn
   ```

### 运行后端服务
#### 开发
- 普通运行
    ```shell
    go run .
    ```

  - `air` 方式运行
      - `mode`: 运行模式, 项目内置`dev`开发与`prod`生产模式, 不同模式使用不同的配置文件运行程序的方式, 默认为`dev`开发模式
      - `port`: 运行端口, 项目运行的端口,默认端口为`4000`.
  
      1. 示例1: 使用`8080` 端口的`开发`模式.
          ```shell
              air -- -mode=dev -port=8080
          ```
    
      2. 示例2: 使用`443` 端口的`生产`模式.
          ```shell
              air -- -mode=prod -port=443
          ```

- (推荐)`Markfile`方式运行 `Golang` 程序与`Swagger`
```shell
make all
```

#### 打包

- (推荐)`Markfile`方式运行
```shell
make build
```

- go 自带
```shell
go build -o .
```

### 运行 `Web` 服务.
    - npm:
    ```shell
    npm run dev
    ```

    - pnpm:
    ```shell
    pnpm dev
    ```

    - yarn
   ```shell
   yarn dev
   ```

## Test

- 运行所有测试用例
   ```shell
   go test
   ```

- 单个测试
   ```shell
   go test ./test/file_test.go
   ```

- HTTP 测试 (仅适用于 jetbrains系列的 IDEA, WebStorm, Goland IDE)

1. 参考`<project>/examples/http/test/api.http`文件
2. 修改`<project>/test/http-client.env.json`为你的项目URI地址
3. 打开`<project>/test/api.http`, 修改这些示例为你项目真实的接口地址, 点击左侧运行图标即可运行 `HTTP` 请求

- sql迁移测试

1. 打开`test`目录
2. 运行`init_pg_data_test.go`
   ```shell
   go test init_pg_data_test.go
   ```
3. 到`Postgres`的 `tests`数据库查看sql 文件数据迁移是否成功
