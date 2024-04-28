TARGET_TOOL=$1
if [ -z "$1" ]; then
    echo "没有指定编译项目";
    exit 1;
else
    echo "编译项目: $1"
fi

BUILD_CMD="cd ${TARGET_TOOL} && CGO_ENABLED=1 go build -ldflags \"-s -w\" -buildvcs=false -o outputs/${TARGET_TOOL}.so -buildmode=c-shared"
go clean -cache
bash -c "${BUILD_CMD}"