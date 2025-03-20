all: fmt build_play

fmt:
	@mkdir -p bin
	go fmt ./...
	go mod tidy

build_play: 
	GOOS=linux   GOARCH=amd64 go build -o bin/googleplay_dl     internal/play/play.go
	GOOS=windows GOARCH=amd64 go build -o bin/googleplay_dl.exe internal/play/play.go

package: build_play
	@mkdir -p bin
	@zip -r bin/googleplay_dl_linux.zip   bin/googleplay_dl bin/apks bin/token
	@zip -r bin/googleplay_dl_windows.zip bin/googleplay_dl.exe bin/apks bin/token
