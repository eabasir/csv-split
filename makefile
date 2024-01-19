build-windows:
	@GOOS=windows GOARCH=amd64 go build -o split-win.exe main.go

build-macos:
	@GOOS=darwin GOARCH=amd64 go build -o split-mac main.go

clean:
	@rm -f split-win.exe split-mac