package main

import (
	"log"
	"net"

	pb "github.com/aynp/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type sumServer struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSumServiceServer(grpcServer, &sumServer{})

	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
