package main

import (
	"context"
	"fmt"
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
	AddUser(client)

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
