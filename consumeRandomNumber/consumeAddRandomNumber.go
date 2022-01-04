package consumeRandomNumber

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/randomNumber"
)

func ConsumeAddRandomNumber(waiter *sync.WaitGroup, client randomNumber.RandomServiceClient) {
	waiter.Add(1)
	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*10,
	)

	var request *randomNumber.RandomNumberRequest = &randomNumber.RandomNumberRequest{
		Number: 100000,
	}

	//make request
	stream, streamErr := client.AddRandomNumber(reqContext, request)
	if streamErr != nil {
		panic(streamErr)
	}
	stream.CloseSend()

	//consume stream
	for {
		res, resErr := stream.Recv()
		//when a stream has no more items to recieve it should return
		//a io.EOF
		if resErr == io.EOF {
			fmt.Println("finished consuming service AddRandomNumber")
			break
		} else if resErr != nil && resErr != io.EOF {
			panic(resErr)
		}

		fmt.Println("AddRandomNumber service ", res.String())
	}

	defer waiter.Done()
	defer cancelFunc()
}
