#!/bin/bash

# 后端服务文件
BACKEND_FILE=/app/ipnas6

# 启动后端服务
nohup ${BACKEND_FILE} > /dev/null 2>&1 &

# nginx 守护进程
nginx -g "daemon off;"
