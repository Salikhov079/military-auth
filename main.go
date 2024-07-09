package main

import (
	"fmt"
	"log"

	"root/api"
	"root/api/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "root/genprotos"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8087"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()
	us:=pb.NewUserServiceClient(UserConn)
	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8081")
	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
