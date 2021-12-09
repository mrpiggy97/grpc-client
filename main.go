package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mrpiggy97/grpc-client/user"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("172.27.0.2:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	var client user.UserServiceClient = user.NewUserServiceClient(connection)
	cxt, cancelF := context.WithTimeout(context.Background(), time.Second)
	defer cancelF()
	var request *user.UserRequest = &user.UserRequest{
		UserId: "John",
	}
	response, err2 := client.GetUser(cxt, request)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("client recieved:", response.String())
}
