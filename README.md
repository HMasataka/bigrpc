# grpc streaming example

## Generate proto files

```bash
protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    server.proto
```

## Requirements

```bash
$ protoc --version
libprotoc 29.3
```

```bash
$ protoc-gen-go --version
protoc-gen-go v1.36.6
```

```bash
$ protoc-gen-go-grpc --version
protoc-gen-go-grpc 1.3.0
```
