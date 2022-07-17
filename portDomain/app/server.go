package app

import (
	"ekszuki/uploader/portDomain/app/contracts"
	protoport "ekszuki/uploader/portDomain/protos/port"
	"ekszuki/uploader/utils"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Repositories struct {
	Port contracts.PortRepository
}

type Server struct {
	protoport.UnimplementedPortServiceServer

	GRPCServer   *grpc.Server
	Repositories Repositories
}

func NewServer(repositories Repositories) *Server {
	grpcServer := grpc.NewServer()
	s := &Server{
		GRPCServer:   grpcServer,
		Repositories: repositories,
	}

	protoport.RegisterPortServiceServer(grpcServer, s)
	if strings.EqualFold(utils.GetEnv("GRPC_REFLECTION", "S"), "S") {
		reflection.Register(grpcServer)
	}

	return s
}
