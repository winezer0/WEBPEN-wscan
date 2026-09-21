CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o wscan_linux_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64  go build -o wscan_linux_arm64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o wscan_darwin_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o wscan_darwin_arm64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o wscan_windows_amd64.exe


prefix="wscan"

# 遍历目录并压缩文件
for file in "./$prefix"*; do
    if [ -f "$file" ]; then  # 确保是文件而不是目录
        # 压缩文件为zip
        zip -r "${file}.zip" "$file" >/dev/null 2>&1

        # 删除源文件
        # rm "$file"
    fi
done