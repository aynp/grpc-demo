package main

import (
	pb "github.com/aynp/grpc-demo/proto"
	"io"
	"log"
)

func (s *sumServer) AddToSum(stream pb.SumService_AddToSumServer) error {
	log.Printf("Streaming started from client")
	var sum int64 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.SumResponse{Sum: sum})
		}
		if err != nil {
			return err
		}

		sum += req.Num
		log.Printf("Recieved = %d, Current Sum = %d", req.Num, sum)
	}
}
