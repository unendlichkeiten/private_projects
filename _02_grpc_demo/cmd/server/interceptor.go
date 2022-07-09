package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// mySqlUnaryServerInterceptor 服务端一元拦截器
func mySqlUnaryServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Pre-processing logic
	// set requestID
	uuID, _ := uuid.NewRandom()
	ctx = context.WithValue(ctx, "requestID", uuID)

	start := time.Now()
	log.Printf("server unary interceptor pre-processing: requstID [%s]\n", ctx.Value("requestID"))

	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)

	// Post processing logic
	end := time.Now()
	log.Printf("server unary interceptor post processing: time [%s]\n", end.Sub(start))
	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Server Stream Interceptor Post Processing RecvMsg (Type: %T) at %s",
		m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Server Stream Interceptor Post Processing SendMsg (Type: %T) at %v",
		m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func mysqlServerStreamInterceptor(
	srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Pre-processing
	log.Println("Server Stream Interceptor Pre-processing : ", info.FullMethod)

	// Invoking the StreamHandler to complete the execution of RPC invocation
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}
