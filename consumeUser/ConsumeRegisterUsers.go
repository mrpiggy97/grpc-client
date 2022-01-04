package consumeUser

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/user"
)

func ConsumeRegisterUsers(waiter *sync.WaitGroup, client user.UserServiceClient) {
	waiter.Add(1)
	//get stream
	streamContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	stream, streamError := client.RegisterUsers(streamContext)
	if streamError != nil {
		panic(streamError)
	}
	//make requests
	for i := 0; i < 100; i++ {
		var request *user.RegisterUserRequest = &user.RegisterUserRequest{
			Username: fmt.Sprintf("%v username", i),
			Password: fmt.Sprintf("%v", rand.Int63()),
		}
		//make request
		var requestError error = stream.Send(request)
		if requestError != nil {
			panic(requestError)
		}
	}
	stream.CloseSend()

	//consume stream
	for {
		response, resError := stream.Recv()
		if resError != nil && resError != io.EOF {
			panic(resError)
		}
		if resError == io.EOF {
			fmt.Println("finished consuming service RegisterUsers")
			break
		}
		fmt.Println("register users service ", response.String())
	}
	defer waiter.Done()
	defer cancelFunc()
}
