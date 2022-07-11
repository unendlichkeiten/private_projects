package main

import (
	"context"
	"flag"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	"google.golang.org/grpc"
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
	smileClient := NewSmileClient(cc)

	// step 3: start request gRPCServices
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ctx.Err()
	// set metadata to server
	md := metadata.New(map[string]string{"token": uuid.New().String()})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// test for custom cancel
	//log.Printf("delay 2 seconds and cancel context ... \n")
	//time.Sleep(2*time.Second)
	//cancel()

	smileClient.UnarySmile(ctx, "client smile ━○o｡.d(*´∀`*)b.｡o○━")
	helloClient.UnaryHello(ctx, "Hello, I am unary client hamming")
	helloClient.BiStreamHello(ctx, "Hello, I am bi-stream client hamming")

	// test for without deadline
	//helloClient.UnaryHello(context.Background(), "Hello, I am unary client hamming")
	//helloClient.BiStreamHello(context.Background(), "Hello, I am bi-stream client hamming")
}
