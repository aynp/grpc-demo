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

type helloServer struct {
	pb.HelloServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})

	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
