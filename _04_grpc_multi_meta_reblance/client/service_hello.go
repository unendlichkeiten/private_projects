package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	"_04_grpc_multi_meta_reblance/pb"
)

type HelloClient struct {
	service pb.ChatClient
}

func NewHelloClient(cc *grpc.ClientConn) *HelloClient {
	service := pb.NewChatClient(cc)

	return &HelloClient{service}
}

func (helloClient *HelloClient) UnaryHello(ctx context.Context, say string) {
	req := &pb.Req{
		Say: say,
	}

	res, err := helloClient.service.UnaryHello(ctx, req)
	if err != nil {
		log.Fatalf("client say hello err [%s]\n", err)
		return
	}

	log.Printf("received a server reply [%s]\n", res.GetReply())
}

func (helloClient *HelloClient) BiStreamHello(ctx context.Context, say string) {

	stream, err := helloClient.service.BiStreamHello(ctx)
	if err != nil {
		log.Fatalf("client say hello err [%s]\n", err)
		return
	}

	stream.Send(&pb.Req{Say: say})

	// go routine to receive responses
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more responses [%s]\n", err)
			break
		}
		if err != nil {
			log.Printf("stream receive err [%s]\n", err)
			return
		}

		log.Printf("received a server reply [%s]\n", res.GetReply())
	}
}
