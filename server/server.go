package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-test/proto"
	"log"
	"net"
	"time"
)

const PORT = "localhost:4000"

type server struct {
	pb.UnimplementedGreeterServer
}

func (server *server) Enter(ctx context.Context, req *pb.EnterMessage) (*pb.GreetMessage, error) {
	return &pb.GreetMessage{Name: req.GetName(), Time: time.Now().String()}, nil
}

func (server *server) Exit(ctx context.Context, req *pb.ExitMessage) (*pb.GreetMessage, error) {
	return &pb.GreetMessage{Name: req.GetName(), Time: time.Now().String()}, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}
	// Defined at grpc module for connecting TCP listener
	gRPCServer := grpc.NewServer()

	// connect with Interface *.pb.go with gRPC server instance
	pb.RegisterGreeterServer(gRPCServer, &server{})

	// Conncet
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve listener : %v", err)
	}

}
