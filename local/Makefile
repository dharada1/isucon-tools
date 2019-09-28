all: help

SOURCE_PATH := "./"
BINARY_LINUX_PATH := "./bin/linux"
BINARY_DARWIN_PATH := "./bin/darwin"

build: build-linux build-darwin ## build for all

build-linux: ## build for linux
	GO111MODULE=on \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o $(BINARY_LINUX_PATH) $(SOURCE_PATH)

build-darwin: ## build for mac
	GO111MODULE=on \
	GOOS=darwin \
	GOARCH=amd64 \
	go build -o $(BINARY_DARWIN_PATH) $(SOURCE_PATH)

deploy: ## deploy (scp) to the server
	echo "do something"

help: ## command lists
	@echo "=========================="
	@echo "local/Makefile"
	@echo "ISUCON用の便利スクリプト雛形"
	@echo "(ローカルマシンで動かす用)"
	@echo "=========================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
