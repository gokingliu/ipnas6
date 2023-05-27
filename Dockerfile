# 指定基础镜像
FROM alpine:3.17

# 复制执行文件
COPY app /app/

ENV PUID=0 PGID=0 UMASK=022

# 执行脚本
CMD ["/app/entrypoint.sh"]
