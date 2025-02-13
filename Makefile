TESTABLE_PACKAGES = `go list ./... | grep -v examples | grep -v constants | grep -v mocks | grep -v helpers | grep -v interfaces | grep -v protos | grep -v e2e | grep -v benchmark`
GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = Cocos3DPlayableAdPlugin
ifeq ($(OS),Windows_NT)
    detected_OS := Windows
else
    detected_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif
setup: init-submodules
	@go get ./...

init-submodules:
	@git submodule init

setup-ci:
	@go get github.com/mattn/goveralls
	@go get -u github.com/wadey/gocovmerge

setup-protobuf-macos:
	@brew install protobuf
	@go get github.com/golang/protobuf/protoc-gen-go

run:
	@go run main.go

build-windows:
	@SET CGO_ENABLED=0&&SET GOOS=windows&&SET GOARCH=amd64&&go build -o ./out/windows/$(BINARY_NAME).exe main.go
  	# @CGO_ENABLED=0&&GOOS=darwin&&GOARCH=amd64&&go build -o ./out/windows/$(BINARY_NAME).exe main.go

build-darwin:
	@SET CGO_ENABLED=0&&SET GOOS=darwin&&SET GOARCH=amd64&&go build -o ./out/darwin/$(BINARY_NAME) main.go

build-linux:
	@SET CGO_ENABLED=0&&SET GOOS=linux&&SET GOARCH=amd64&&go build -o ./out/linux/$(BINARY_NAME) main.go

build:
	@make build-windows
	@make build-darwin
	@make build-linux