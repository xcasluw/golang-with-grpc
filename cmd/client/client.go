package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/xcasluw/fullcycle-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserSeviceClient(connection)
	// AddUser(client)
	AddUserVerbose(client)

}

func AddUser(client pb.UserSeviceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "joao@email.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserSeviceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "joao@email.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("could not receive the msg: %v", err)
		}

		fmt.Println("status:", stream.Status)
	}
}
