package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-test/proto"
	"log"
	"time"
)

const PORT = "localhost:4000"

func main() {
	conn, err := grpc.Dial(PORT, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Failed to dial : %v", err)
	}

	defer conn.Close()

	// connect using proto buffer greeter client
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	response, err := client.Enter(ctx, &pb.EnterMessage{
		Name: "msh",
	})

	if err != nil {
		log.Fatalf("Failed to enter : %v", err)
	}

	log.Printf(`
		Name : %s entered!
		at Time : %s
	`, response.Name, response.Time)

	response, err = client.Exit(ctx, &pb.ExitMessage{
		Name: "milkcoke",
	})

	if err != nil {
		log.Fatalf("Failed to exit : %v", err)
	}

	log.Printf(`
		Name: %s exited!
		at Time : %s
	`, response.Name, response.Time)
}
