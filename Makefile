all: help

#####################################################
# Configurations
#####################################################

# local path
LOCAL_PATH_GO_SOURCE     := ./go/
LOCAL_PATH_BINARY_LINUX  := ./bin/linux
LOCAL_PATH_BINARY_DARWIN := ./bin/darwin

# server path
SERVER_PATH_GO_BINARY            := /home/isucon/go/bin/
SERVER_PATH_NGINX_ACCESS_LOG     := /tmp/nginx/access.log
SERVER_PATH_MYSQL_SLOW_QUERY_LOG := /tmp/mysql/slow-query.log
SERVER_PATH_PPROF_FILE           := /path/to/cpu.pprof

# 実行時刻のタイムスタンプ
# $(MAKEFILE_EXECUTE_TIMESTAMP) で呼んでログの入替などにつかう
$(eval MAKEFILE_EXECUTE_TIMESTAMP := $(shell date +%Y%m%d%I%M%S))

#####################################################
# Build & Run Go Application (Local)
#####################################################

build: build-darwin ## alias of build-darwin
run: run-darwin ## alias of run-darwin

build-linux: # build for linux
	GO111MODULE=on GOOS=linux GOARCH=amd64 \
	go build -o $(LOCAL_PATH_BINARY_LINUX) $(LOCAL_PATH_GO_SOURCE)

build-darwin: # build for darwin (mac)
	GO111MODULE=on GOOS=darwin GOARCH=amd64 \
	go build -o $(LOCAL_PATH_BINARY_DARWIN) $(LOCAL_PATH_GO_SOURCE)

run-linux: build-linux # run in linux
	$(LOCAL_PATH_BINARY_LINUX)

run-darwin: build-darwin # run in darwin (mac)
	$(LOCAL_PATH_BINARY_DARWIN)

#####################################################
# Deploy Go Application
#####################################################

deploy-scp-a: ## バイナリ直送型 (server a)
	scp $(LOCAL_PATH_BINARY_DARWIN) isucon-a:$(SERVER_PATH_GO_BINARY)
	ssh isucon-a "systemctl stop app.golang"
	# do something
	# restartの方が適切だけど、
	# ログの入替とかdbのinit走らせたい場合ここでやる
	ssh isucon-a "systemctl start app.golang"

deploy-scp-all: ## バイナリ直送型 (server abc)
	# TODO

deploy-git-a: ## go run in dockerの場合こっち (server a)
	# TODO

deploy-git-all: ## go run in dockerの場合こっち (server abc)
	# TODO

#####################################################
# pprof
#####################################################

pprof: ## pprof intaractive mode. enter command. ( e.g. list fib )
	go tool pprof $(LOCAL_PATH_BINARY_DARWIN) cpu.pprof

pprof-gui: ## pprof-gui
	go tool pprof -http=":9999" $(LOCAL_PATH_BINARY_DARWIN) cpu.pprof

pprof-png: ## pprof-png
	go tool pprof -png $(LOCAL_PATH_BINARY_DARWIN) cpu.pprof > out.png

#####################################################
# help
#####################################################

help: ## command lists
	@echo "========================================="
	@echo "github.com/dharada1/isucon-tools/Makefile"
	@echo "========================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

help-detailed: ## cat Makefile
	cat ./Makefile

#####################################################
# Installation ( Ansible のレシピ書いたほうが良い
#####################################################

install-htop: ## install htop for all server.
	ssh isucon-a "sudo apt-get install htop"
	ssh isucon-b "sudo apt-get install htop"
	ssh isucon-c "sudo apt-get install htop"

install-alp: ## install alp
	# TODO

go-pprof: ## go-pprof
	# TODO

mysql-myprofiler: ## myprofiler
	# TODO

#####################################################
# alp
#####################################################

alp-exec: ## analyze nginx access log with alp
	# その場でaggregatesは書いてみてください
	# alp --aggregates="/api/users/.*","/api/events/.*","/admin/api/reports/events/.*"
	cat $(SERVER_PATH_NGINX_ACCESS_LOG) | alp

#####################################################
# Rotate Logs!
#####################################################

nginx-access-log-rotate: ## lotate nginx log.
	mv $(SERVER_PATH_NGINX_ACCESS_LOG) $(SERVER_PATH_NGINX_ACCESS_LOG)_$(MAKEFILE_EXECUTE_TIMESTAMP)
	touch $(SERVER_PATH_NGINX_ACCESS_LOG)
	sudo systemctl restart nginx

go-pprof-rotate: ## lotate cpu.pprof
	mv $(SERVER_PATH_PPROF_FILE) $(SERVER_PATH_PPROF_FILE)_$(MAKEFILE_EXECUTE_TIMESTAMP)

mysql-slow-query-log-rotate: ## lotate mysql slow query log.
	mv $(SERVER_PATH_MYSQL_SLOW_QUERY_LOG) $(SERVER_PATH_MYSQL_SLOW_QUERY_LOG)_$(MAKEFILE_EXECUTE_TIMESTAMP)
	touch $(SERVER_PATH_MYSQL_SLOW_QUERY_LOG)
	sudo systemctl restart mysql

#####################################################
# mysql
#####################################################

# todo
# 	docker内から吐く場合SERVER_PATH_MYSQL_SLOW_QUERY_LOGの設定に注意
# 	ポートも
# 	username, passwordも!!!!!
mysql-slow-query-log-enable: ## enable mysql slow query log with long_query_time 0
	sudo mysql -e "set global slow_query_log_file = '$(SERVER_PATH_MYSQL_SLOW_QUERY_LOG)'; set global long_query_time = 0; set global slow_query_log = ON;"

mysql-slow-query-log-disable: ## disable mysql slow query log
	sudo mysql -e "set global slow_query_log = OFF;"


#####################################################
# tail logs!
#####################################################

tail-log: ## check logs.
	sudo journalctl -f


# init-db-c: ## db reset for isucon-c server.
#   ssh isucon-c "sudo systemctl stop isutrain-go"
#   ssh isucon-c "sudo docker rm webapp_mysql_1"
#   ssh isucon-c "sudo docker volume rm webapp_mysql"
#   ssh isucon-c "sudo systemctl start isutrain-go"