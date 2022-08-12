package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/xcasluw/fullcycle-grpc/pb"
)

// type UserSeviceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserSeviceServer()
// AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserSevice_AddUserVerboseClient, error)
// AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserSevice_AddUsersClient, error)
// AddUserStreamBoth(ctx context.Context, opts ...grpc.CallOption) (UserSevice_AddUserStreamBothClient, error)
// }

type UserService struct {
	pb.UnimplementedUserSeviceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	// Insert database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil

}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserSevice_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserSevice_AddUsersServer) error {
	users := []*pb.User{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatal("error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("adding", req.GetName())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserSevice_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("error receiving stream from the client: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User:   req,
		})

		if err != nil {
			log.Fatal("error sending stream to the client: %v", err)
		}
	}
}
