# 使用已有的 redis alpine 镜像作为基础（本地已有，无需下载）
FROM redis:7-alpine

# 安装 ca-certificates（用于 HTTPS 请求）
RUN apk add --no-cache ca-certificates 2>/dev/null || true

WORKDIR /root/

# 复制本地编译好的二进制文件
COPY bin/shortlink ./main

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]
