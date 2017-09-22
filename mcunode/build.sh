#Ubuntu test ok,install upx，set ：
PATH="$HOME/go1.9/go/bin:$PATH"  #PATH for Go bin file
export GOPATH="$HOME/go"         #GOPATH


env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./mcunode_darwin_amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o ./mcunode_linux_386
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./mcunode_linux_amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ./mcunode_linux_arm
env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o ./mcunode_windows_386.exe
env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./mcunode_windows_amd64.exe
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64 go build -ldflags "-s -w" -o ./mcunode_linux_mips64
env CGO_ENABLED=0 GOOS=linux GOARCH=mips64le go build -ldflags "-s -w" -o ./mcunode_linux_mips64le
env CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -ldflags "-s -w" -o ./mcunode_linux_mips
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags "-s -w" -o ./mcunode_linux_mipsle

upx ./mcunode_darwin_amd64
upx ./mcunode_linux_386
upx ./mcunode_linux_amd64
upx ./mcunode_linux_arm
upx ./mcunode_windows_386.exe
upx ./mcunode_windows_amd64.exe
upx ./mcunode_linux_mips64
upx ./mcunode_linux_mips64le
upx ./mcunode_linux_mips
upx ./mcunode_linux_mipsle

tar czvf mcunode_darwin_amd64.tar.gz views static mcunode_darwin_amd64
tar czvf mcunode_linux_386.tar.gz views static mcunode_linux_386
tar czvf mcunode_linux_amd64.tar.gz views static mcunode_linux_amd64
tar czvf mcunode_linux_arm.tar.gz views static mcunode_linux_arm
tar czvf mcunode_windows_386.tar.gz views static mcunode_windows_386.exe
tar czvf mcunode_windows_amd64.tar.gz views static mcunode_windows_amd64.exe
tar czvf mcunode_linux_mips64.tar.gz views static mcunode_linux_mips64
tar czvf mcunode_linux_mips64le.tar.gz views static mcunode_linux_mips64le
tar czvf mcunode_linux_mips.tar.gz views static mcunode_linux_mips
tar czvf mcunode_linux_mipsle.tar.gz views static mcunode_linux_mipsle
