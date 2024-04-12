#!/bin/bash

# 定义镜像和容器名称
IMAGE_NAME="go-tool-compiler"
CONTAINER_NAME="go-tool-compiler-build"
TARGET_TOOL=$1
if [ -z "$1" ]; then
    echo "没有指定编译项目";
    exit 1;
else
    echo "编译项目: $1"
fi


echo "开始在docker中编译 ${TARGET_TOOL}"

# 构建镜像
if docker image ls | grep -q $IMAGE_NAME; then
  echo "镜像已存在"
else
  echo "开始构建镜像"
  docker build -t $IMAGE_NAME .
fi

container_exists=$(docker ps -a --filter "name=^/${CONTAINER_NAME}$" --format '{{.Names}}')

if [ -z "$container_exists" ]; then
    echo "容器 $CONTAINER_NAME 不存在，创建并启动..."
    # 创建并启动容器，这里假设使用的镜像名为go-tool-compiler
    # 根据需要调整docker run命令的参数
    docker run --name $CONTAINER_NAME -d $IMAGE_NAME tail -f /dev/null
elif [ "$(docker inspect -f '{{.State.Running}}' $CONTAINER_NAME)" = "false" ]; then
    echo "容器 $CONTAINER_NAME 已存在但停止了，启动它..."
    # 启动容器
    docker start $CONTAINER_NAME
else
    echo "容器 $CONTAINER_NAME 已在运行。"
fi
# 启动容器
#docker run --name $CONTAINER_NAME -d $IMAGE_NAME tail -f /dev/null

# 将当前项目拷贝到容器中
docker cp ../go-tools $CONTAINER_NAME:/app

BUILD_CMD="cd /app/go-tools/${TARGET_TOOL} && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build -ldflags \"-s -w\" -buildvcs=false -o outputs/${TARGET_TOOL}.so -buildmode=c-shared"

SRC_ARCH=$(docker exec $CONTAINER_NAME go env GOARCH)

# 在容器内运行编译命令，这里假设是go build
docker exec $CONTAINER_NAME bash -c "${BUILD_CMD}"

# 从容器中拷贝编译内容出来，假设输出文件名为app
HOST_DIR_PATH=${TARGET_TOOL}/docker_output_${SRC_ARCH}_2_amd64
if [ -d "$HOST_DIR_PATH" ]; then
    echo "删除现有内容"
    rm -r "$HOST_DIR_PATH"
fi
docker cp "$CONTAINER_NAME:/app/go-tools/${TARGET_TOOL}/outputs" "${TARGET_TOOL}/docker_output_${SRC_ARCH}_2_amd64"

# 停止并删除容器
docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME
