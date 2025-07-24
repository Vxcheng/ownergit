protoc --go_out=plugins=grpc:. customer.proto

protoc --go_out=./pb --go-grpc_out=./pb ./pb/greeter.proto
