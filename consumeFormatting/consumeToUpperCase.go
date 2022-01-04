package consumeFormatting

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/formatting"
)

func ConsumeToUpperCase(waiter *sync.WaitGroup, client formatting.FormattingServiceClient) {
	waiter.Add(1)
	//make request
	var request *formatting.FormattingRequest = &formatting.FormattingRequest{
		StringToConvert: "Michael-j-fox",
	}

	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)

	res, resError := client.ToUpperCase(reqContext, request)
	if resError != nil {
		panic(resError)
	}

	fmt.Println("ToUpperCase service ", res.String())

	defer waiter.Done()
	defer cancelFunc()
}
