package consumeFormatting

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/formatting"
)

func ConsumeToCamelCase(waiter *sync.WaitGroup, client formatting.FormattingServiceClient) {
	waiter.Add(1)
	//make request
	var request *formatting.FormattingRequest = &formatting.FormattingRequest{
		StringToConvert: "Ccamel-case-example",
	}

	reqContext, cancelFunc := context.WithTimeout(
		context.Background(),
		time.Second*30,
	)

	res, resError := client.ToCamelCase(reqContext, request)
	if resError != nil {
		panic(resError)
	}

	fmt.Println("ToCamelCase service ", res.String())

	defer waiter.Done()
	defer cancelFunc()
}
