package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "github.com/basic-gRPC-Go/proto"
	"time"
)

const (
	address = "localhost:60001"
)

func main () {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewEmployeeDetailsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.RegisterEmployee(ctx, &pb.Employee{
		EmployeeId: "123",
		Name:       "Cool Name",
		Department: "Awesome Dept",
	})
	if err != nil {
		log.Fatalf("error communicating %v", err)
	}

	log.Printf("response from server %v",response)
}