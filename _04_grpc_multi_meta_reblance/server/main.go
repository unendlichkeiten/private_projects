package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"_04_grpc_multi_meta_reblance/pb"
)

func main() {
	port := flag.Int("port", 808, "the server port")
	flag.Parse()

	// step 1: news a mysql server
	helloServer := NewHelloServer()
	smileServer := NewSmileServer()

	// step 2: sets a listened address
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("start server err: [%s]", err)
	}

	// step 3: news a grpc server
	gRPCServer := grpc.NewServer()

	// step 4: register grpc server
	pb.RegisterChatServer(gRPCServer, helloServer)
	pb.RegisterSmileServer(gRPCServer, smileServer)
	reflection.Register(gRPCServer)

	// step 5: starts serve
	gRPCServer.Serve(listener)
}
