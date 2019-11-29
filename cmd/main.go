package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	grpc_reflection "google.golang.org/grpc/reflection"

	"github.com/akhilesharora/todo/env"
	"github.com/akhilesharora/todo/internal/repo"
	"github.com/akhilesharora/todo/internal/server"
	"github.com/akhilesharora/todo/pb"
)

func main() {
	cfg, err := env.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	serverAddress := net.JoinHostPort(cfg.Host, cfg.Port)
	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	log.Info("server started on ", serverAddress)

	grpcServer := grpc.NewServer()
	todoServer := server.NewTodoServer(repo.NewTodoService())
	pb.RegisterTodoServer(grpcServer, todoServer)

	// for GRPC ui
	grpc_reflection.Register(grpcServer)

	// Register exit function to gracefully stop grpc server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-c:
			log.Info("gracefully stopping server...")
			grpcServer.GracefulStop()
			log.Info("done")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
