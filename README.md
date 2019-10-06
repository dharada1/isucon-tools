# isucon-tools

This is my ISUCON toolbox!

```
.
├── Makefile
├── README.md
├── bin
│   ├── darwin
│   └── linux
├── go
│   └── main.go
├── go.mod
├── go.sum
├── go_snippets
│   ├── README.md
│   ├── daemon
│   │   └── main.go
│   ├── http
│   │   └── main.go
│   ├── mysql
│   │   ├── mysql-connections.go
│   │   └── sql-proxy.go
│   ├── parallel
│   │   ├── Makefile
│   │   ├── main.go
│   │   └── main_test.go
│   ├── slice
│   │   └── main.go
│   └── tracing
│       ├── datadog
│       │   ├── Makefile
│       │   ├── README.md
│       │   ├── apptracer.go
│       │   ├── bin
│       │   │   └── darwin
│       │   ├── hoge_handler.go
│       │   └── main.go
│       └── pprof
│           ├── Makefile
│           ├── bin
│           │   └── darwin
│           ├── cpu.pprof
│           ├── main.go
│           └── out.png
├── mysql
│   ├── README.md
│   ├── alter_table_add_index_sample.sql
│   └── conf.d
│       └── my.conf
├── nginx
│   └── conf.d
│       └── defaults.conf
├── playbook
│   ├── 000_競技開始.md
│   ├── 010_鍵配置.md
│   └── 020_gitリポジトリの作成.md
└── templates
    └── alp
        ├── README.md
        ├── apache_ltsv.conf
        ├── h2o_ltsv.conf
        └── nginx_ltsv.conf
```