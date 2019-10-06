all: help

#####################################################
# Configurations
#####################################################

SOURCE_PATH := "./"
BINARY_LINUX_PATH := "./bin/linux"
BINARY_DARWIN_PATH := "./bin/darwin"

# log path
NGINX_ACCESS_LOG_PATH:=/tmp/nginx/access.log
MYSQL_SLOW_QUERY_LOG_PATH:=/tmp/mysql/slow-query.log
PPROF_FILE_PATH:=/path/to/cpu.pprof

# ログのsuffixにつける用
$(eval EXEC_DATE := $(shell date +%Y%m%d%I%M%S))

#####################################################
# Build Go Application
#####################################################

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

deploy-%: ## deploy. deploy-isucon-1, deploy-isucon-2, deploy-isucon-3. 前提として ~/.ssh/config にisucon-1とかでno pass sshできるように記載が必要
	# deploy
	scp ./bin/linux $*:/home/isucon/go/bin/

	# stop
	ssh $* "systemctl stop app.golang"

	# rotate logs
	ssh $* "systemctl restart app.golang"
	ssh $* "make nginx-access-log-rotate"
	ssh $* "make mysql-slow-query-log-rotate"
	# ssh $* "make go-pprof-rotate"

	# start
	ssh $* systemctl start app.golang

#####################################################
# pprof
#####################################################

pprof: ## pprof intaractive mode. enter command. ( e.g. list fib )
	go tool pprof $(BINARY_DARWIN_PATH) cpu.pprof

pprof-gui: ## pprof-gui
	go tool pprof -http=":9999" $(BINARY_DARWIN_PATH) cpu.pprof

pprof-png: ## pprof-png
	go tool pprof -png $(BINARY_DARWIN_PATH) cpu.pprof > out.png

#####################################################
# help
#####################################################

help: ## command lists
	@echo "========================================="
	@echo "github.com/dharada1/isucon-tools/Makefile"
	@echo "========================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

help-detailed: ## cat Makefile
	cat ./Makefile

#####################################################
# Installation ( Ansible のレシピ書いたほうが良い
#####################################################

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
	cat $(NGINX_ACCESS_LOG_PATH) | alp

#####################################################
# Rotate Logs!
#####################################################

nginx-access-log-rotate: ## lotate nginx log.
	mv $(NGINX_ACCESS_LOG_PATH) $(NGINX_ACCESS_LOG_PATH)_$(EXEC_DATE)
	touch $(NGINX_ACCESS_LOG_PATH)
	sudo systemctl restart nginx

go-pprof-rotate: ## lotate cpu.pprof
	mv $(PPROF_FILE_PATH) $(PPROF_FILE_PATH)_$(EXEC_DATE)

mysql-slow-query-log-rotate: ## lotate mysql slow query log.
	mv $(MYSQL_SLOW_QUERY_LOG_PATH) $(MYSQL_SLOW_QUERY_LOG_PATH)_$(EXEC_DATE)
	touch $(MYSQL_SLOW_QUERY_LOG_PATH)
	sudo systemctl restart mysql

#####################################################
# mysql
#####################################################

# todo
# 	docker内から吐く場合MYSQL_SLOW_QUERY_LOG_PATHの設定に注意
# 	ポートも
# 	username, passwordも!!!!!
mysql-slow-query-log-enable: ## enable mysql slow query log with long_query_time 0
	sudo mysql -e "set global slow_query_log_file = '$(MYSQL_SLOW_QUERY_LOG_PATH)'; set global long_query_time = 0; set global slow_query_log = ON;"

mysql-slow-query-log-disable: ## disable mysql slow query log
	sudo mysql -e "set global slow_query_log = OFF;"


#####################################################
# tail logs!
#####################################################

tail-log: ## check logs.
	sudo journalctl -f
