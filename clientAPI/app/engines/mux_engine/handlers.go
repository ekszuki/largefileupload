package muxengine

import (
	"ekszuki/uploader/clientAPI/app/core"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// UpLoadFilePortsHandler Upload large json file
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

// FindPortByKeyHandler - Find Port by key
func (s *MuxEnginer) FindPortByKeyHandler(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	logCtx := logrus.WithFields(
		logrus.Fields{"package": "muxenginer", "function": "FindPortByKeyHandler"},
	)
	logCtx.Info("starting request")

	vars := mux.Vars(req)
	key := vars["key"]

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("key is invalid"))
		return
	}

	portCore := core.NewPort(s.client)
	resp, err := portCore.FindByKey(ctx, key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}

	bs, err := json.Marshal(resp)
	if err != nil {
		logCtx.Errorf("error on parse json response, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-server-time", time.Now().UTC().Format(time.RFC3339))
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
