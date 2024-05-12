## Installation
### Download go module
Install go runtime protoc-gen-go with each folder(eg. user and transaction)
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
Install depedencies grpc library
```
go get -u google.golang.org/grpc@v1.26.0
```
Run this command below for each folder
```
go mod download
```
### Running application
Run application within inside 2 folder
```
go run cmd/main.go -config=config/.env
```
