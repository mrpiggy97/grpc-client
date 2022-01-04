package consumeFormatting

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/formatting"
)

func ConsumeToLowerCase(waiter *sync.WaitGroup, client formatting.FormattingServiceClient) {
	waiter.Add(1)
	//make request
	var request *formatting.FormattingRequest = &formatting.FormattingRequest{
		StringToConvert: "MICHELEE-aT-thIS",
	}
	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)
	res, resError := client.ToLowerCase(reqContext, request)
	if resError != nil {
		panic(resError)
	}

	fmt.Println("toLowerCase service ", res.String())
	defer waiter.Done()
	defer cancelFunc()
}
