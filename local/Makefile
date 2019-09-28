all: help

build: ## build go application
	echo "do something"

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
