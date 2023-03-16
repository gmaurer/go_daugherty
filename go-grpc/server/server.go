package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "source.cloud.google.com/dl-gcp-cngo-sbox-devenv-b1/go_daugherty/go-grpc/gen/proto"
)

type employeeApiServer struct {
	pb.UnimplementedEmployeeApiServer
}

type Employee struct {
	FirstName  string
	LastName   string
	Age        int64
	Email      string
	Department string
}

func (s *employeeApiServer) SayHello(ctx context.Context, name *pb.HelloRequest) (*pb.HelloResponse, error) {
	greeting := "Hello, " + name.Msg
	return &pb.HelloResponse{Msg: greeting}, nil
}

func (s *employeeApiServer) GetAllEmployees(ctx context.Context, _ *emptypb.Empty) (*pb.EmployeesResponse, error) {

	return &pb.EmployeesResponse{
		Data: []*pb.Employee{
			{
				FirstName:  "Alex",
				LastName:   "Sheperd",
				Age:        35,
				Email:      "alex.sheperd@daugherty.com",
				Department: "SAE",
			},
			{
				FirstName:  "Jim",
				LastName:   "Pratt",
				Age:        45,
				Email:      "jim.pratt@daugherty.com",
				Department: "DA",
			},
			{
				FirstName:  "Lindsay",
				LastName:   "Elliot",
				Age:        27,
				Email:      "lindsay.elliot@daugherty.com",
				Department: "SAE",
			},
		},
	}, nil
}

func (s *employeeApiServer) GetEmployee(ctx context.Context, _ *emptypb.Empty) (*pb.Employee, error) {
	employee := &pb.Employee{FirstName: "Alex", LastName: "Sheperd", Age: 35, Email: "alex.sheperd@daugherty.com", Department: "SAE"}
	return employee, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8090")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterEmployeeApiServer(grpcServer, &employeeApiServer{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
}
