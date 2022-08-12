package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// AddUserVerbose(client)
	AddUsers(client)
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

func AddUsers(client pb.UserSeviceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "j1",
			Name:  "João",
			Email: "joao1@email.com",
		},
		&pb.User{
			Id:    "j2",
			Name:  "João",
			Email: "joao2@email.com",
		},
		&pb.User{
			Id:    "j3",
			Name:  "João",
			Email: "joao3@email.com",
		},
		&pb.User{
			Id:    "j4",
			Name:  "João",
			Email: "joao4@email.com",
		},
		&pb.User{
			Id:    "j5",
			Name:  "João",
			Email: "joao5@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}

	fmt.Println(res)
}
