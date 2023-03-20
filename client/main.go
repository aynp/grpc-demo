package main

import (
	"log"

	pb "github.com/aynp/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	nums := []int64{2, 5, 10, 6, 7, 20}

	callAddToSum(client, nums)
}
