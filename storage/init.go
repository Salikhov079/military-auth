package postgres

import (
	pb "root/genprotos"
)

type InitRoot interface {
	User() User
}
type User interface {
	LoginUser(user *pb.LoginUserRequest ) (*pb.LoginUserResponse, error)
}
