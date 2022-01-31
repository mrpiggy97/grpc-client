package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func AuthInterceptor(
	cxt context.Context,
	method string,
	req interface{},
	reply interface{},
	connection *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	cxt = metadata.AppendToOutgoingContext(cxt, "password", "go")
	return invoker(cxt, method, req, reply, connection, opts...)
}

func WithAuthInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(AuthInterceptor)
}
