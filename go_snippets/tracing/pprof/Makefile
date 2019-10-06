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

pprof: ## pprof intaractive mode. enter command. ( e.g. list fib )
	go tool pprof $(BINARY_DARWIN_PATH) cpu.pprof

pprof-gui: ## pprof-gui
	go tool pprof -http=":9999" $(BINARY_DARWIN_PATH) cpu.pprof

pprof-png: ## pprof-png
	go tool pprof -png $(BINARY_DARWIN_PATH) cpu.pprof > out.png

bench: ## send sample requests
	curl localhost:8000/hoge
	curl localhost:8000/hoge2
	curl localhost:8000/hoge3

help: ## command lists
	@echo "=========================="
	@echo "local/Makefile"
	@echo "ISUCON用の便利スクリプト雛形"
	@echo "(ローカルマシンで動かす用)"
	@echo "=========================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
