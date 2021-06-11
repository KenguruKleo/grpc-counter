### Genarate

```
protoc --go_out=models --go_opt=paths=source_relative --go-grpc_out=models --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false  protos/services/counter/*.proto
```

### Run

```
go mod init main.go
go mod tidy
go run .
```