# 指定基础镜像
FROM alpine:latest

# 复制执行文件
COPY app /app/

# 执行脚本
ENTRYPOINT ["/app/entrypoint.sh"]
