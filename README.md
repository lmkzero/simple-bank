# simple-bank
my project for [backend master class](https://www.youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)

## Setting up
### db migration
cli工具: https://github.com/golang-migrate/migrate

初始化：
`migrate create -ext sql -dir db/migration -seq $(name)`

执行变更到db：
`migrate -path db/migration -database "$(DB_URL)" -verbose up`

### sql compiler (golang)
cli工具：https://github.com/sqlc-dev/sqlc

sqlc基本配置：https://docs.sqlc.dev/en/latest/reference/config.html

sqlc基本语法：https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html#schema-and-queries