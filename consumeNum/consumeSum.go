package consumeNum

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/num"
)

func ConsumeSum(waiter *sync.WaitGroup, client num.NumServiceClient) {
	waiter.Add(1)
	streamContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	stream, streamError := client.Sum(streamContext)
	if streamError != nil {
		panic(streamError)
	}
	for i := 0; i <= 5; i++ {
		var request *num.SumRequest = &num.SumRequest{
			Number: int64(i),
		}
		var sendingErr error = stream.Send(request)
		if sendingErr != nil {
			panic(sendingErr)
		}
	}
	res, resError := stream.CloseAndRecv()
	if resError != nil && resError != io.EOF {
		panic(resError)
	}
	fmt.Println(resError)
	fmt.Println(res.String())
	fmt.Println("finished consuming SumService")
	defer waiter.Done()
	defer cancelFunc()
}
