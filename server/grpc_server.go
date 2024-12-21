package server

import (
	"context"
	"log"

	"myservice/db"
	pb "myservice/grpcservices/proto"
)

type Server struct {
	pb.UnimplementedDataServiceServer
}

func (s *Server) StoreData(ctx context.Context, in *pb.Data) (*pb.Response, error) {
	log.Printf("gRPC: Received StoreData request with ID: %s, Content: %s", in.Id, in.Content)
	err := db.StoreData(in.Id, in.Content)
	if err != nil {
		log.Printf("gRPC: StoreData failed: %v", err)
		return &pb.Response{Success: false}, err
	}
	log.Printf("gRPC: StoreData successful for ID: %s", in.Id)
	return &pb.Response{Success: true}, nil
}
