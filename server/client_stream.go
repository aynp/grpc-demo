package main

import (
	"io"
	"log"

	pb "github.com/aynp/grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.HelloService_SayHelloClientStreamingServer) error {
	log.Printf("Client streaming started")
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}

		if err != nil {
			return err
		}

		log.Printf("Got requests with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
