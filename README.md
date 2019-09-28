# isucon-tools

scripts for ISUCON !

- go
  - goのサンプルとか
    - daemon化
    - goroutineでの並列化
    - インメモリキャッシュ
    - net/http, sql.DBのチューニング

- scripts
  - local/Makefile
    - ローカル用の便利スクリプト
      - goのbuild,deployなど
  - server/Makefile
    - サーバ用の便利スクリプト
      - inspired by https://github.com/tohutohu/isucon9/blob/master/Makefile
      - log rotate
      - journalctl
      - alp
      - slow queryのon/off

- templates/alp
  - alp用のltsvテンプレート
    - nginx
    - apache
    - h2o

- templates/configutations
  - パフォーマンスチューニング系の秘伝のタレメモ
    - nginx
    - mysql
    - kernel系

- slack.sh
  - pipeでslack投稿するシンプルなスクリプト (TODO)
