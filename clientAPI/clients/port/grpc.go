package port

import (
	protoport "ekszuki/uploader/portDomain/protos/port"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(url string) protoport.PortServiceClient {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("could not get grpc connection, %v", err)
	}
	cc := protoport.NewPortServiceClient(conn)

	return cc
}
