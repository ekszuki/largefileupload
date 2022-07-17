package main

import (
	"context"
	"ekszuki/uploader/clientAPI/app"
	"ekszuki/uploader/clientAPI/clients/port"
	"ekszuki/uploader/utils"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logCtx := logrus.WithFields(
		logrus.Fields{"component": "cmd", "function": "main"},
	)

	logCtx.Info("Starting Client API")
	ctx := context.Background()

	cancelCtx, cancelFunc := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	apiPort, err := strconv.Atoi(utils.GetEnv("API_PORT", "3000"))
	if err != nil {
		apiPort = 3000
	}

	grpcClientURL := utils.GetEnv("PORT_GRPC_URL", "localhost:9090")
	logCtx.Infof("Connecting with GRPC Port Service on url: %s", grpcClientURL)
	portGrpcClient := port.NewGRPCClient(grpcClientURL)

	logCtx.Infof("Starting HTTP Server on port %d", apiPort)
	server := app.NewServer(cancelCtx, portGrpcClient)
	server.Run(fmt.Sprintf(":%d", apiPort))

	<-c
	if err := server.Shutdown(); err != nil {
		logCtx.Infof("Error on graceful shutdown: %v", err)
	}
	cancelFunc()

	logCtx.Info("Waiting 5 seconds for turning server off")
	time.Sleep(5 * time.Second)
	logCtx.Info("HTTP Server Graceful Shutdown")
}
