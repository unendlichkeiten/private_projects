package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// mySqlUnaryClientInterceptor 客户端一元拦截器
func mySqlUnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	// Pre-processing logic
	// set requestID
	uuID, _ := uuid.NewRandom()
	ctx = context.WithValue(ctx, "requestID", uuID)

	start := time.Now()
	log.Printf("client unary interceptor pre-processing: requstID [%s]\n",
		ctx.Value("requestID"))

	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Post processing logic
	end := time.Now()
	log.Printf("client unary interceptor post processing: time [%s]\n", end.Sub(start))

	return err
}

// wrappedStream  用于包装 grpc.ClientStream 结构体并拦截其对应的方法。
type wrappedStream struct {
	grpc.ClientStream
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Client Stream Interceptor Post Processing RecvMsg (Type: %T) at %v",
		m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Client Stream Interceptor Post Processing SendMsg (Type: %T) at %v",
		m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func mysqlClientStreamInterceptor(
	ctx context.Context, desc *grpc.StreamDesc,
	cc *grpc.ClientConn, method string,
	streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	// Pre-processing
	log.Println("Client Stream Interceptor Pre-processing : ", method)

	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}

	return newWrappedStream(s), nil
}
