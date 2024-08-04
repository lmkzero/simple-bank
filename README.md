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

protoc插件：https://github.com/varluffy/protoc-gen-go-gin
因为后续会使用Gin框架，所以插件生成的handler需要和Gin handler函数签名保持一致。
同时生成的桩代码包含了RegisterService()方法，所以无需手动注册路由。

### pb tag
Gin在绑定uri（ShouldBindUri）时，需要对应结构体包含"uri"标记字段。

cli工具：https://github.com/favadi/protoc-go-inject-tag

### third_party
bank.proto的依赖pb文件。

### generate
```shell
protoc --proto_path=./api/bank/v1 \
        --proto_path=./third_party \
        --go_out ./api/bank/v1 --go_opt=paths=source_relative \
        --go-gin_out ./api/bank/v1 --go-gin_opt=paths=source_relative \
        --validate_out=paths=source_relative,lang=go:./api/bank/v1 \
        bank.proto

protoc-go-inject-tag -input="./api/bank/v1/*.pb.go"
```