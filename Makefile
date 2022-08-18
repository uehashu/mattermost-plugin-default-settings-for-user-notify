default: clean build archive

clean:
	rm -rf server
	rm -rf plugin.tar.gz

build:
	GOOS=linux GOARCH=amd64 go build -o server/dist/plugin-linux-amd64 plugin.go
	GOOS=linux GOARCH=arm64 go build -o server/dist/plugin-linux-arm64 plugin.go
	GOOS=darwin GOARCH=amd64 go build -o server/dist/plugin-darwin-amd64 plugin.go
	GOOS=darwin GOARCH=arm64 go build -o server/dist/plugin-darwin-arm64 plugin.go
	GOOS=windows GOARCH=amd64 go build -o server/dist/plugin-windows-amd64.exe plugin.go

archive:
	tar -czvf plugin.tar.gz server plugin.json