package main

import (
	"log"
	"time"

	pb "github.com/aynp/grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.HelloService_SayHelloServerStreamingServer) error {
	log.Printf("Got names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name + " 👋",
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
