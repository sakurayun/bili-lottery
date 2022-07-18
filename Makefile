gen:
	go build -o bili-lottery-tool *.go

amd64linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bili-lottery-tool *.go
	tar zcvf bili-lottery-tool-linux-amd64.tar.gz bili-lottery-tool example.toml README.md

arm64linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bili-lottery-tool *.go
	tar zcvf bili-lottery-tool-linux-arm64.tar.gz bili-lottery-tool example.toml README.md

amd64windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o bili-lottery-tool.exe *.go
	tar zcvf bili-lottery-tool-windows-amd64.tar.gz bili-lottery-tool.exe example.toml README.md

amd64mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bili-lottery-tool *.go
	tar zcvf bili-lottery-tool-macOS-amd64.tar.gz bili-lottery-tool example.toml README.md

clean:
	rm bili-lottery-tool*

