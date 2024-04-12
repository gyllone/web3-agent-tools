# 使用go:1.21作为基础镜像
FROM golang:1.20

# 更新软件包列表并安装gcc-x86-64-linux-gnu
RUN apt-get update && apt-get install -y gcc-x86-64-linux-gnu

# 设置工作目录
WORKDIR /app