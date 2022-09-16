# MaxBlog 后端微服务模板

protobuf
```shell
# protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
# grpc
go get -u google.golang.org/grpc
# protobuf
go get -u google.golang.org/protobuf
```

```shell
protoc -I . --go_out=plugins=grpc:. *.proto
```