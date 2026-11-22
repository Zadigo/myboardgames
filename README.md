# My Board Games



## Commands

### Generate protocole buffer files

```Shell
# install: https://buf.build/docs/installation

buf lint

buf generate

cd gen

go mod init myboardgames/gen

go mod tidy   # pulls google.golang.org/grpc and protobuf

cd ..

go work init ./gen ./services/minimalgo ./services/minimalgame
```

```Go

// Write these in services/minimalgame/go.mod and services/minimalgo/go.mod

require myboardgames/gen v0.0.0

replace myboardgames/gen => ../../gen
```
