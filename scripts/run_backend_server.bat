chcp 65001

@REM 进入到根目录
@echo off
cd ../../example

@REM 运行后端服务
@echo off
air -- -mode=dev -port=4000
pause
