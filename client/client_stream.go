package main

import (
	"context"
	pb "github.com/aynp/grpc-demo/proto"
	"log"
	"time"
)

func callSayHelloClientStreaming(client pb.HelloServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Sent request with name: %s", name)
		time.Sleep(time.Second)
	}

	log.Printf("Client streaming stopped")
}
