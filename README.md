# isucon-tools

scripts for ISUCON !

- local/Makefile
  - ローカル用の便利スクリプト
    - goのbuild,deployなど

- server/Makefile
  - サーバ用の便利スクリプト
    - inspired by https://github.com/tohutohu/isucon9/blob/master/Makefile
    - log rotate
    - alp
    - slow queryのon/off

- templates/alp
  - alp用のltsvテンプレート
    - nginx
    - apache
    - h2o

- templates/configutations
  - パフォーマンスチューニング系 (メモ程度)
    - nginx
    - mysql

- slack.sh
  - pipeでslack投稿するシンプルなスクリプト (TODO)
