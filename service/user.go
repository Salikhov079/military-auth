package service

import (
	"context"
	"log"
	pb "root/genprotos"
	s "root/storage"

)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) LoginUser(ctx context.Context, user *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	prod, err := c.stg.User().LoginUser(user)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}
