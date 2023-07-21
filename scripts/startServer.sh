#!/bin/bash

# 打开第一个窗口，并执行任务1
osascript -e 'tell application "Terminal" to do script "sudo sh ./cmd/run_frontend_server.sh"'

# 打开第二个窗口，并执行任务2
osascript -e 'tell application "Terminal" to do script "sudo sh ./cmd/run_backend_server.sh"'
