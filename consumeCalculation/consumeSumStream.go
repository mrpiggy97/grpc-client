package consumeCalculation

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/calculation"
)

func ConsumeSumStream(waiter *sync.WaitGroup, client calculation.CalculationServiceClient) {
	waiter.Add(1)
	streamContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*120,
	)

	stream, streamErr := client.SumStream(streamContext)
	if streamErr != nil {
		panic(streamErr)
	}

	//send requests to server
	for i := 0; i < 100; i++ {
		var request *calculation.SumStreamRequest = &calculation.SumStreamRequest{
			A: int32(i),
			B: int32(i * 5),
		}

		var sendingError error = stream.Send(request)
		if sendingError != nil {
			panic(sendingError)
		}
	}

	stream.CloseSend()

	//consume stream
	for {
		res, resError := stream.Recv()
		if resError != nil && resError != io.EOF {
			panic(resError)
		}
		if resError == io.EOF {
			fmt.Println("finished consuming service SumStream")
			break
		}
		fmt.Println("SumStreamService ", res.String())
	}

	defer waiter.Done()
	defer cancelFunc()
}
