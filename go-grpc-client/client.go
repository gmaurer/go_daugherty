package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "source.cloud.google.com/dl-gcp-cngo-sbox-devenv-b1/go_daugherty/go-grpc/gen/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewEmployeeApiClient(conn)

	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Msg: "World"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Msg)

	otherResponse, err := client.GetAllEmployees(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(otherResponse.Data)

	empResponse, err := client.GetEmployee(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(empResponse)

}
