package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/basic-gRPC-Go/proto"
)

const (
	port = ":60001"
)

type server struct {}

func (s *server) RegisterEmployee(ctx context.Context, employeeReq *pb.Employee) (*pb.EmployeeRes, error) {
	log.Printf("Employee Req: %v", employeeReq)
	return &pb.EmployeeRes{
		Created: true,
	}, nil
}

func (s *server) GetEmployeeById(ctx context.Context, req *pb.EmployeeIdReq) (*pb.Employee, error) {
	log.Printf("By Id Req: %v", req.EmployeeId)
	return &pb.Employee{}, nil
}

func main () {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listening to port %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmployeeDetailsServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}


