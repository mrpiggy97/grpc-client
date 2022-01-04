package consumeNum

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/num"
)

func ConsumeRnd(waiter *sync.WaitGroup, client num.NumServiceClient) {
	waiter.Add(1)
	//make request
	var request *num.NumRequest = &num.NumRequest{
		From:   0,
		To:     67,
		Number: 100,
	}

	streamContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)

	stream, streamErr := client.Rnd(streamContext, request)

	if streamErr != nil {
		panic(streamErr)
	}
	stream.CloseSend()

	//consume stream
	for {
		_, resError := stream.Recv()
		if resError != nil && resError != io.EOF {
			panic(resError)
		}

		if resError == io.EOF {
			fmt.Println("finished consuming service Rnd")
			break
		}
	}
	defer waiter.Done()
	defer cancelFunc()
}
