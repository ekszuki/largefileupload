generate-ports-proto:
	@protoc --go-grpc_out=portDomain/protos --go_out=portDomain/protos portDomain/protos/port.proto

generate-mock-portDomain-repository:
	@mockgen  -destination=portDomain/mocks/port/repository.go -package=port -source portDomain/app/contracts/repository.go PortRepository
