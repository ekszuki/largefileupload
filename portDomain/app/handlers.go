package app

import (
	"context"
	"ekszuki/uploader/portDomain/app/core"
	"ekszuki/uploader/portDomain/app/parsers"
	protoport "ekszuki/uploader/portDomain/protos/port"
	"io"

	"github.com/sirupsen/logrus"
)

func (s *Server) UploadPorts(stream protoport.PortService_UploadPortsServer) error {
	logCtx := logrus.WithFields(logrus.Fields{"package": "app", "function": "UploadPorts"})
	ctx := stream.Context()
	portCore := core.NewPortApplication(s.Repositories.Port)

	for {
		stdPort, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&protoport.Empty{})
		}
		if err != nil {
			logCtx.Errorf("client stream died, %v", err)
			return err
		}

		dmPort := parsers.FromUpLoadPortRequestToDomain(stdPort)
		err = portCore.SaveOrUpdate(ctx, dmPort)
		if err != nil {
			return err
		}
	}
}

func (s *Server) FindByKey(ctx context.Context, req *protoport.FindByKeyRequest) (*protoport.FindByKeyResponse, error) {
	portCore := core.NewPortApplication(s.Repositories.Port)
	dmPort, err := portCore.FindByKey(ctx, req.GetKey())

	return parsers.FromDomainToFindByKeyResponse(dmPort), err
}
