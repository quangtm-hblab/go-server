package main

import (
	"context"
	"log"
	"time"

	pb "github.com/quangtm-hblab/core-server/calculatorpb"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50053"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Sum(ctx, &pb.SumRequest{Num1: 6, Num2: 7})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.GetResult())
}
