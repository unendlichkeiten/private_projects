package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	"_03_grpc_deadline_and_cancel/pb"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s\n", *serverAddress)

	transportOption := grpc.WithInsecure()

	// step 1: dials a gRPCServer address
	cc, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatalf("cannot dial server: %s", err)
	}

	// step 2: news a service client
	helloClient := NewHelloClient(cc)

	// step 3: start request gRPCServices
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// test for custom cancel
	//log.Printf("delay 2 seconds and cancel context ... \n")
	//time.Sleep(2*time.Second)
	//cancel()

	helloClient.UnaryHello(ctx, "Hello, I am unary client hamming")
	helloClient.BiStreamHello(ctx, "Hello, I am bi-stream client hamming")

	// test for without deadline
	//helloClient.UnaryHello(context.Background(), "Hello, I am unary client hamming")
	//helloClient.BiStreamHello(context.Background(), "Hello, I am bi-stream client hamming")
}

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
