package main

import (
	"context"
	"log"
	"time"

	pb "myservice/grpcservices/proto" // Adjust this path to where your proto package is generated

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051"
	defaultName = "test_id"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDataServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(name) == 0 {
		name = defaultName
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.StoreData(ctx, &pb.Data{Id: name, Content: "Test content"})
	if err != nil {
		log.Fatalf("could not store data: %v", err)
	}
	log.Printf("Response: %v", r.GetSuccess())
}
