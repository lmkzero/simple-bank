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

## API Definition
### pb validation
protoc插件：https://github.com/bufbuild/protoc-gen-validate

使用规则：https://github.com/bufbuild/protoc-gen-validate?tab=readme-ov-file#constraint-rules

### method annotations
使用option进行标注，定义为HTTP方法。

protoc插件：
- https://pkg.go.dev/github.com/zhufuyi/sponge/cmd/protoc-gen-go-gin （最好用这个，因为后续会使用Gin框架，所以插件生成的handler需要和Gin handler函数签名保持一致。）
- https://pkg.go.dev/github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2

### third_party
bank.proto的依赖pb文件。

### generate
```shell
protoc --proto_path=./api/bank/v1 \
        --proto_path=./third_party \
        --go_out=paths=source_relative:./api/bank/v1 \
        --validate_out=paths=source_relative,lang=go:./api/bank/v1 \
        bank.proto
```