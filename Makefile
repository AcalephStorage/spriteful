APP_NAME = spriteful

all: clean deps build

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

build: format test
	@echo "--> Building application"
	@gb build ...
	@tar cfz bin/${APP_NAME}-linux-amd64.tar.gz -C bin ${APP_NAME}
	@rm bin/${APP_NAME}
