package services

import (
	"context"
	"fmt"

	"github.com/xcasluw/fullcycle-grpc/pb"
)

// type UserSeviceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserSeviceServer()
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
