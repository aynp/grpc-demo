package main

import (
	"context"
	pb "github.com/aynp/grpc-demo/proto"
	"log"
	"time"
)

func callAddToSum(client pb.SumServiceClient, nums []int64) {
	log.Printf("Starting streaming")

	stream, err := client.AddToSum(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, num := range nums {
		req := &pb.SumRequest{
			Num: num,
		}

		// req is of type *pb.SumRequest
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error: %v", err)
		}

		log.Printf("Sent request with number %d", num)
		time.Sleep(time.Second)
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// reply has type *pb.SumResponse
	log.Println(reply.Sum)
}
