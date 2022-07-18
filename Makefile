generate-port-proto:
	@protoc --go-grpc_out=portDomain/protos --go_out=portDomain/protos portDomain/protos/port.proto

generate-mock-portDomain-repository:
	@mockgen  -destination=portDomain/mocks/port/repository.go -package=port -source portDomain/app/contracts/repository.go PortRepository

build-containers:
	@docker-compose build

daemon-start-containers:
	@docker-compose up -d

start-containers:
	@docker-compose up

stop-containers:
	@docker-compose down

run-tests:
	@go test -cover ./...