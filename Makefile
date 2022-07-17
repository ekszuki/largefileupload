generate-ports-proto:
	@protoc --go-grpc_out=portDomain/protos --go_out=portDomain/protos portDomain/protos/port.proto