推荐版本组合（用于 Go 1.22 和 grpc@v1.68.0）：

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.33.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```
然后升级 protoc 到 3.21.12 或 21.x（推荐）

buf generate 命令生成hello.pb.go代码
