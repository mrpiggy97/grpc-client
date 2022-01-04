package consumeUser

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/user"
)

func ConsumeGetUser(waiter *sync.WaitGroup, client user.UserServiceClient) {
	waiter.Add(1)
	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)
	var request *user.UserRequest = &user.UserRequest{
		UserId: "this  is my user",
	}

	response, resError := client.GetUser(
		reqContext,
		request,
	)

	if resError != nil {
		panic(resError)
	}
	fmt.Println("get user service ", response.String())
	defer waiter.Done()
	defer cancelFunc()
}
