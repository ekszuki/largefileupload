package muxengine

import (
	"ekszuki/uploader/clientAPI/app/core"
	"net/http"

	"github.com/sirupsen/logrus"
)

// UpLoadPortsHandler implements contracts.PortApi
func (s *MuxEnginer) UpLoadFilePortsHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logCtx := logrus.WithFields(
		logrus.Fields{"package": "muxenginer", "function": "UpLoadFilePortsHandler"},
	)
	logCtx.Info("starting request")

	portCore := core.NewPort(s.client)
	err := portCore.UploadPortFile(ctx, req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
