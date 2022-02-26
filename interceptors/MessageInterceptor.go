package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func MessageInterceptor(
	cxt context.Context,
	method string,
	req interface{},
	reply interface{},
	conn *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	options ...grpc.CallOption,
) error {
	cxt = metadata.AppendToOutgoingContext(cxt, "message", "this is my message")
	return invoker(cxt, method, req, reply, conn, options...)
}

func WithMessageInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(MessageInterceptor)
}
