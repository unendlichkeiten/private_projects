package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"_01_grpc_demo"
	"_01_grpc_demo/pb"
)

func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()

	// step 1: news a mysql server
	mySqlServer := _01_grpc_demo.NewMysqlServer()

	// step 2: sets a listened address
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("start server err: ", err)
	}

	// step 3: news a grpc server
	gRPCServer := grpc.NewServer()

	// step 4: register grpc server
	pb.RegisterMysqlServiceServer(gRPCServer, mySqlServer)
	reflection.Register(gRPCServer)

	// step 5: starts serve
	log.Printf("start GRPC server at %s\n", listener.Addr().String())
	gRPCServer.Serve(listener)
}
