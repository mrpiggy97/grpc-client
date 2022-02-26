package main

import (
	"sync"
	"time"

	"github.com/mrpiggy97/sharedProtofiles/calculation"
	"github.com/mrpiggy97/sharedProtofiles/formatting"
	"github.com/mrpiggy97/sharedProtofiles/num"
	"github.com/mrpiggy97/sharedProtofiles/randomNumber"
	"github.com/mrpiggy97/sharedProtofiles/user"

	"google.golang.org/grpc"

	"github.com/mrpiggy97/grpcClient/consumeCalculation"
	"github.com/mrpiggy97/grpcClient/consumeFormatting"
	"github.com/mrpiggy97/grpcClient/consumeNum"
	"github.com/mrpiggy97/grpcClient/consumeRandomNumber"
	"github.com/mrpiggy97/grpcClient/consumeUser"
	"github.com/mrpiggy97/grpcClient/interceptors"
)

func main() {
	//run server
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	// ip below refers to a docker container ip connected to the same
	// docker network
	connection, err := grpc.Dial(
		"172.27.0.2:50051",
		grpc.WithInsecure(),
		interceptors.WithAuthInterceptor(),
		interceptors.WithClientLoggerInterceptor(),
		interceptors.WithMessageInterceptor(),
	)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	//get clients
	var userClient user.UserServiceClient = user.NewUserServiceClient(connection)
	var randomNumberClient randomNumber.RandomServiceClient = randomNumber.NewRandomServiceClient(
		connection,
	)
	var numClient num.NumServiceClient = num.NewNumServiceClient(connection)
	var formattingClient formatting.FormattingServiceClient = formatting.NewFormattingServiceClient(
		connection,
	)
	var calculationClient calculation.CalculationServiceClient = calculation.NewCalculationServiceClient(
		connection,
	)

	//make clients consume api
	go consumeUser.ConsumeRegisterUsers(waiter, userClient)
	go consumeUser.ConsumeGetUser(waiter, userClient)
	go consumeRandomNumber.ConsumeAddRandomNumber(waiter, randomNumberClient)
	go consumeRandomNumber.ConsumeSubstractRandomNumber(waiter, randomNumberClient)
	go consumeNum.ConsumeRnd(waiter, numClient)
	go consumeFormatting.ConsumeToCamelCase(waiter, formattingClient)
	go consumeFormatting.ConsumeToLowerCase(waiter, formattingClient)
	go consumeFormatting.ConsumeToUpperCase(waiter, formattingClient)
	go consumeCalculation.ConsumeSumStream(waiter, calculationClient)
	go consumeNum.ConsumeSum(waiter, numClient)
	time.Sleep(time.Second * 1)
	waiter.Wait()
}
