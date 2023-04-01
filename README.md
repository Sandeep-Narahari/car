# Protobuf Setup for grpc-gateway and grpc using Golang 
#### Practices and steps
1. In the proto folder google/api/annotations.proto has to be there then only we can able to avoid source relative error 
2. Try to maintain the separate folder for proto generated files Ex: types folder
3. Install Go
4. Install 
`go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest` <br>
`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` <br>
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

#### Run to generate the proto files
`protoc --proto_path=./proto --go_out=paths=source_relative:./types --go-grpc_out=paths=source_relative:./types --grpc-gateway_out=paths=source_relative:./types ./proto/query.proto ./proto/car.proto `
