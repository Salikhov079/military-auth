package handler

import pb "root/genprotos"

type Handler struct {
	User   pb.UserServiceClient
}

func NewHandler(us pb.UserServiceClient) *Handler {
	return &Handler{us}
}
