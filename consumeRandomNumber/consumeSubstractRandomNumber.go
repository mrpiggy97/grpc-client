package consumeRandomNumber

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/randomNumber"
)

func ConsumeSubstractRandomNumber(waiter *sync.WaitGroup, client randomNumber.RandomServiceClient) {
	waiter.Add(1)
	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*10,
	)

	var request *randomNumber.RandomNumberRequest = &randomNumber.RandomNumberRequest{
		Number: 100000,
	}

	//make request
	stream, streamErr := client.SubstractRandomNumber(reqContext, request)
	if streamErr != nil {
		panic(streamErr)
	}
	stream.CloseSend()

	//consume stream
	for {
		res, resErr := stream.Recv()
		if resErr == io.EOF {
			fmt.Println("finished consuming service SubstractRandomNumber")
			break
		} else if resErr != nil && resErr != io.EOF {
			panic(resErr)
		}

		fmt.Println("SubstractRandomNumber service ", res.String())
	}

	defer waiter.Done()
	defer cancelFunc()
}
