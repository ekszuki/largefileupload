package muxengine

import (
	protoport "ekszuki/uploader/portDomain/protos/port"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxEnginer struct {
	client protoport.PortServiceClient
}

func NewMuxEnginer(portGRPCClient protoport.PortServiceClient) *MuxEnginer {
	return &MuxEnginer{
		client: portGRPCClient,
	}
}

func (m *MuxEnginer) CreateMuxEnginer() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/port/fileupload", m.UpLoadFilePortsHandler).Methods(http.MethodPost)
	r.HandleFunc("/port/{key}", m.FindPortByKeyHandler).Methods(http.MethodGet)

	return r
}
