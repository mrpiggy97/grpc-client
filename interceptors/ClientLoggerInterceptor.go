package interceptors

import (
	"context"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientLoggerInterceptor(
	cxt context.Context,
	method string,
	req interface{},
	reply interface{},
	connection *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	var os string = runtime.GOOS
	zone, _ := time.Now().Zone()
	cxt = metadata.AppendToOutgoingContext(cxt, "os", os)
	cxt = metadata.AppendToOutgoingContext(cxt, "zone", zone)
	var err error = invoker(cxt, method, req, reply, connection, opts...)
	return err
}

func WithClientLoggerInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(ClientLoggerInterceptor)
}
