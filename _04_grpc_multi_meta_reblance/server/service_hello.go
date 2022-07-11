package main

import (
	"context"
	"io"
	"log"

	"_04_grpc_multi_meta_reblance/pb"
)

type HelloServer struct{}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (*HelloServer) UnaryHello(ctx context.Context, req *pb.Req) (*pb.Res, error) {
	say := req.GetSay()
	log.Printf("receives a client message [%s]\n", say)

	// test deadline
	//ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	//defer cancel()

	if err := ctx.Err(); err != nil {
		log.Printf("ctx terminated [%s]\n", err)
		return nil, err
	}

	res := &pb.Res{Reply: "hello, I am unary hamming server"}

	return res, nil
}

func (*HelloServer) BiStreamHello(stream pb.Chat_BiStreamHelloServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("read client ends [%s]\n", err)
			return err
		}

		if err != nil {
			log.Printf("stream receive err [%s]\n", err)
			return err
		}

		// response
		log.Printf("received a client message [%s]\n", req.GetSay())

		// test no delay
		if err := stream.Send(&pb.Res{Reply: "Hello, I am bi-stream server hamming"}); err != nil {
			return err
		}

		// test EOF
		//for count := 0; count < 5; count++ {
		//	if err := stream.Send(&pb.Res{Reply:"Hello, I am bi-stream server hamming"}); err != nil {
		//		return err
		//	}
		//
		//	sleep := rand.Intn(5)
		//	time.Sleep(time.Duration(sleep))
		//}
		//
		//break
	}
	// test EOF
	//return nil
}
