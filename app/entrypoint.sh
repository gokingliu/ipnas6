#!/bin/sh

# 后端服务文件
BACKEND_FILE=/app/ipnas6

# 启动后端服务
nohup ${BACKEND_FILE} > /dev/null 2>&1 &
