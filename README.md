# simple-bank
my project for backend master class

## Setting up
### db migration
cli工具: https://github.com/golang-migrate/migrate

初始化：
`migrate create -ext sql -dir db/migration -seq $(name)`

执行变更到db：
`migrate -path db/migration -database "$(DB_URL)" -verbose up`