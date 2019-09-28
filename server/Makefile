# log path
NGINX_ACCESS_LOG_PATH:=/tmp/nginx/access.log
MYSQL_SLOW_QUERY_LOG_PATH:=/tmp/mysql/slow-query.log

SLACKCAT:=slackcat --tee --channel general
SLACKRAW:=slackcat --channel general

# ログのsuffixにつける用
$(eval EXEC_DATE := $(shell date +%Y%m%d%I%M%S))

all: help

hsg:
	history | grep

xxx: ## do something

install-alp: ## install alp
	# TODO

alp-exec: ## analyze nginx access log with alp
	# その場でaggregatesは書いてみてください
	# alp --aggregates="/api/users/.*","/api/events/.*","/admin/api/reports/events/.*"
	cat $(NGINX_ACCESS_LOG_PATH) | alp

nginx-access-log-rotate: ## lotate nginx log.
	mv $(NGINX_ACCESS_LOG_PATH) $(NGINX_ACCESS_LOG_PATH)_$(EXEC_DATE)
	touch $(NGINX_ACCESS_LOG_PATH)
	sudo systemctl restart nginx

mysql-slow-query-log-rotate: ## lotate mysql slow query log.
	echo $(EXEC_DATE)
	mv $(MYSQL_SLOW_QUERY_LOG_PATH) $(MYSQL_SLOW_QUERY_LOG_PATH)_$(EXEC_DATE)
	touch $(MYSQL_SLOW_QUERY_LOG_PATH)
	sudo systemctl restart mysql

mysql-slow-query-log-enable: ## enable mysql slow query log with long_query_time 0
	sudo mysql -e "set global slow_query_log_file = '$(MYSQL_SLOW_QUERY_LOG_PATH)'; set global long_query_time = 0; set global slow_query_log = ON;"

mysql-slow-query-log-disable: ## disable mysql slow query log
	sudo mysql -e "set global slow_query_log = OFF;"

tail-log: ## check logs.
	sudo journalctl

help: ## command lists
	@echo "=========================="
	@echo "server/Makefile"
	@echo "ISUCON用の便利スクリプト雛形"
	@echo "(サーバーで動かす用)"
	@echo "=========================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

help-detailed: ## for environment variables check
	@cat ./Makefile
