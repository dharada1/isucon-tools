# MySQL

## パラメータチューニング

todo 秘伝のタレ
innodb_buffer_pool_size

## 心得メモ

- まずは順当に slow query log の確認、その結果からindexを張る.
- カラムのスキーマ変更は難易度高めなのでやらない (isucon9はこれで死んだ...)
  - Go側の構造体の型をいちいち変更していくのは、時間的に無理.
