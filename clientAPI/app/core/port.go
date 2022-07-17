package core

import (
	"context"
	"ekszuki/uploader/clientAPI/parsers"
	protoport "ekszuki/uploader/portDomain/protos/port"
	"io"

	"github.com/sirupsen/logrus"
)

type Port struct {
	client protoport.PortServiceClient
}

func NewPort(client protoport.PortServiceClient) *Port {
	return &Port{
		client: client,
	}
}

func (p *Port) UploadPortFile(ctx context.Context, dataStream io.ReadCloser) error {
	logCtx := logrus.WithFields(
		logrus.Fields{"package": "core", "function": "UploadPortFile"},
	)

	stream, err := p.client.UploadPorts(ctx)
	if err != nil {
		logCtx.Errorf("could not create stream to gRPC server: %v", err)
		return err
	}
	defer func() {
		_, err = stream.CloseAndRecv()
		if err != nil {
			logCtx.Warnf("error on close and recv stream: %v", err)
		} else {
			logCtx.Info("stream closed successfully")
		}
	}()

	// closure function to interact with the stream
	jsonFilePortParserFunc, err := parsers.ParseJsonPortFile(dataStream)
	if err != nil {
		return err
	}

	for {
		data, err := jsonFilePortParserFunc()
		if err != nil {
			logCtx.Errorf("could not parse json object: %v", err)
			return err
		}

		if data == nil {
			logCtx.Info("file data parsed completly")
			return nil
		}
		loadPortRequest := parsers.ToUpLoadPortRequest(data)
		err = stream.Send(loadPortRequest)
		if err != nil {
			logCtx.Errorf("stream error: %v", err)
			return err
		}
	}
}
