package main

import (
	"context"
	"io"
	"log"

	pb "github.com/aynp/grpc-demo/proto"
)

func callSayHelloServerStreaming(client pb.HelloServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Recieved - %v", message)
	}

	log.Printf("Streaming finished")
}
