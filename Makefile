APP_NAME = spriteful

all: build

clean:
	@echo "--> Cleaning build"
	@rm -rf ./bin

format:
	@echo "--> Formatting source code"
	@go fmt ./...

deps:
	@echo "--> Getting dependencies"
	@gb vendor restore

test: format
	@echo "--> Testing application"
	@gb test ...

build: format deps
	@echo "--> Building application"
	@gb build ...
	@tar cf bin/${APP_NAME}-linux-amd64.tar -C bin ${APP_NAME}
	@rm bin/${APP_NAME}
