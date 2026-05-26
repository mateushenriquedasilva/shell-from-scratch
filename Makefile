run:
	go run cmd/shell/main.go

build: build-windows build-mac build-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o build/shell-windows-amd64.exe cmd/shell/main.go

build-mac:
	GOOS=darwin GOARCH=arm64 go build -o build/shell-darwin-arm64 cmd/shell/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/shell-darwin-amd64 cmd/shell/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/shell-linux-amd64 cmd/shell/main.go
	GOOS=linux GOARCH=arm64 go build -o build/shell-linux-arm64 cmd/shell/main.go

clean:
	rm -rf build/

