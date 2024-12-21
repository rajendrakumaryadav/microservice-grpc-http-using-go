package main

import (
	"log"
	"net"

	"myservice/db"
	pb "myservice/grpcservices/proto"
	"myservice/server"

	"google.golang.org/grpc"
)

func init() {
	log.Println("Initializing database...")
	db.InitDB()
	log.Println("Database initialization completed.")
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("gRPC: failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDataServiceServer(s, &server.Server{})

	go func() {
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("gRPC: failed to serve: %v", err)
		}
	}()

	go func() {
		server.StartHTTPServer()
		log.Println("HTTP server started successfully")
	}()

	select {}
}
