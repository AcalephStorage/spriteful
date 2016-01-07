APP_NAME = spriteful

all: clean format deps test build

clean:
	@echo "--> Cleaning build"
	@rm -rf ./bin

format:
	@echo "--> Formatting source code"
	@go fmt ./...

deps:
	@echo "--> Getting dependencies"
	@gb vendor restore

test:
	@echo "--> Testing application"
	@gb test ...

build: format test
	@echo "--> Building application"
	@gb build ...
	@tar cf bin/${APP_NAME}-linux-amd64.tar -C bin ${APP_NAME}
	@rm bin/${APP_NAME}
