package main

import (
	"context"
	pb "github.com/aynp/grpc-demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStreaming(client pb.HelloServiceClient, names *pb.NameList) {
	log.Printf("Started streaming")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error: %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error: %v", err)
		}

		time.Sleep(time.Second)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Streaming closed")
}
