package app

import (
	"context"
	muxengine "ekszuki/uploader/clientAPI/app/engines/mux_engine"
	protoport "ekszuki/uploader/portDomain/protos/port"
	"fmt"
	"net/http"
)

type Server struct {
	ctx    context.Context
	server *http.Server
}

func NewServer(
	ctx context.Context, portGRPCClient protoport.PortServiceClient,
) *Server {
	enginer := muxengine.NewMuxEnginer(portGRPCClient).CreateMuxEnginer()

	return &Server{
		ctx: ctx,
		server: &http.Server{
			Handler: enginer,
		},
	}
}

func (s *Server) Run(addr string) {
	s.server.Addr = addr

	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			fmt.Printf("error to start server: %v", err)
		}
	}()
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(s.ctx)
}
