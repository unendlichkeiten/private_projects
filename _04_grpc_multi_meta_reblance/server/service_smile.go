package main

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"log"

	"_04_grpc_multi_meta_reblance/pb"
)

type SmileServer struct{}

func NewSmileServer() *SmileServer {
	return &SmileServer{}
}

func (*SmileServer) UnarySmile(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	// reading metadata from client
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("read metadata error")
	}
	if t, ok := md["token"]; ok {
		log.Printf("received client token [%s]\n", t)
	}

	say := req.GetSay()
	log.Printf("receives a client message [%s]\n", say)

	if err := ctx.Err(); err != nil {
		log.Printf("ctx terminated [%s]\n", err)
		return nil, err
	}

	res := &pb.Res{Reply: "server smile ━d(ゝω･*)b━"}

	return res, nil
}
