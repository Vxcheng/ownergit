# 打包execel_generator
version="1.0.0"
# 定义包目录变量，并加入时间戳
time=$(date +%Y%m%d%H%M%S)
PACKAGE_DIR="execel_generator"
PACKAGE_TAR_NAME="execel_generator_v${version}_${time}.tar.gz"

echo "build start"
GOOS=linux GOARCH=amd64 go build -o excel_generator_linux
GOOS=windows GOARCH=amd64 go build -o excel_generator_windows.exe
GOOS=darwin GOARCH=amd64 go build -o excel_generator_mac
echo "build end"

echo "package start"
# 创建包目录
mkdir -p $PACKAGE_DIR/output
mkdir -p $PACKAGE_DIR/logs

# 复制编译文件到包目录
cp excel_generator_linux $PACKAGE_DIR/
cp excel_generator_windows.exe $PACKAGE_DIR/
cp excel_generator_mac $PACKAGE_DIR/

# 复制output、templates、config.yaml到包目录
cp -r templates $PACKAGE_DIR/
cp config.yaml $PACKAGE_DIR/
cp README.md $PACKAGE_DIR/


# 压缩包目录
tar -czvf $PACKAGE_TAR_NAME $PACKAGE_DIR
echo $PACKAGE_TAR_NAME
# 清理临时文件
rm -rf $PACKAGE_DIR

echo "package end"