package main

import (
	"context"
	"ekszuki/uploader/portDomain/app"
	"ekszuki/uploader/portDomain/repository/mongodb"
	"ekszuki/uploader/utils"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	time.Local = time.UTC
	logCtx := logrus.WithFields(logrus.Fields{"package": "main", "function": "main"})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	grpcAdd := fmt.Sprintf(":%s", utils.GetEnv("GRPC-PORT", "9090"))
	logCtx.Infof("Creating TCP Listener on address %s", grpcAdd)
	listen, err := net.Listen("tcp", grpcAdd)
	if err != nil {
		logCtx.Fatalf("could not create listener on address %s", grpcAdd)
	}

	logCtx.Info("Creating database connection")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	portRepo, err := mongodb.NewPortRepository(ctx)
	if err != nil {
		logCtx.Fatalf("could not create port repository, %v", err)
	}
	repos := app.Repositories{
		Port: portRepo,
	}

	logCtx.Infof("Starting gRPC Server on address %s", grpcAdd)
	server := app.NewServer(repos)
	go func() {
		err := server.GRPCServer.Serve(listen)
		if err != nil {
			logCtx.Fatalf("could not initialize grpc server %v", err)
		}
	}()

	<-c
	server.GRPCServer.GracefulStop()
}
