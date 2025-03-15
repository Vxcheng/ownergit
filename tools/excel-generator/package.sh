echo "build start"
GOOS=linux GOARCH=amd64 go build -o excel_generator_linux
GOOS=windows GOARCH=amd64 go build -o excel_generator_windows.exe
GOOS=darwin GOARCH=amd64 go build -o excel_generator_mac
echo "build end"

