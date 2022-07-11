package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"_04_grpc_multi_meta_reblance/pb"
)

type SmileClient struct {
	service pb.SmileClient
}

func NewSmileClient(cc *grpc.ClientConn) *SmileClient {
	service := pb.NewSmileClient(cc)

	return &SmileClient{service}
}

func (SmileClient *SmileClient) UnarySmile(ctx context.Context, say string) {
	req := &pb.Req{
		Say: say,
	}

	res, err := SmileClient.service.UnarySmile(ctx, req)
	if err != nil {
		log.Fatalf("client smile err [%s]\n", err)
		return
	}

	log.Printf("received a server smile reply [%s]\n", res.GetReply())
}
