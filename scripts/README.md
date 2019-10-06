# scripts

便利なMakefile

## 鍵配置

Makefileは `ssh isucon-a` `ssh isucon-b` `ssh isucon-c` でno passでsshできる前提で書く.

`~/.ssh/config` が書いてあって、かつチーム全員の鍵が全サーバの `~/.ssh/authorized_keys` に配置されてることが前提

```
Host isucon-a
    Hostname xx.xxx.xxx.xxx
    IdentityFile    ~/.ssh/isucon-20191005.pem
    User            isucon
    Port            22

Host isucon-b
    Hostname xx.xxx.xxx.xxx
    IdentityFile    ~/.ssh/isucon-20191005.pem
    User            isucon
    Port            22

Host isucon-c
    Hostname xx.xxx.xxx.xxx
    IdentityFile    ~/.ssh/isucon-20191005.pem
    User            isucon
    Port            22
```

(todo テスト↑)

事前に `ssh-keygen` してもらう

## xxxx

Makefileはサーバー上に配置せず全てローカルから叩く。
サーバー上で何か実行したい場合は `ssh isucon-a "pwd"` のように書く.

```
install-htop: ## install htop for all server.
	ssh isucon-a "sudo apt-get install htop"
	ssh isucon-b "sudo apt-get install htop"
	ssh isucon-c "sudo apt-get install htop"

init-db-c: ## db reset for isucon-c server.
	ssh isucon-c "sudo systemctl stop isutrain-go"
	ssh isucon-c "sudo docker rm webapp_mysql_1"
	ssh isucon-c "sudo docker volume rm webapp_mysql"
	ssh isucon-c "sudo systemctl start isutrain-go"
```


todo server/Makefile の移植

## ディレクトリ構成 (本戦)

`/home/isucon/isutrain/webapp/` 下 で `git init` する前提。 ( `/home/isucon/isutrain/webapp/go` でinitしてはいけない)

```
.
├── Makefile
├── docker-compose.go.yml
├── docker-compose.yml
├── go
│   ├── Dockerfile
│   ├── main.go
│   ├── tmp.go
│   ├── utils.go
│   └── utils_test.go
├── go.mod
├── go.sum
├── mysql
│   └── conf.d
├── nginx
│   └── conf.d
├── public
│   └── static
└── sql
    └── hoge.sql
```

Makefileもこれに応じて応じて書く.
