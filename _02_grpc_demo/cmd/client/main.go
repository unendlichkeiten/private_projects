package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"

	"_02_grpc_demo"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s\n", *serverAddress)

	transportOption := grpc.WithInsecure()

	// step 1: dials a gRPCServer address
	cc, err := grpc.Dial(*serverAddress, transportOption,
		grpc.WithUnaryInterceptor(mySqlUnaryClientInterceptor),
		grpc.WithStreamInterceptor(mysqlClientStreamInterceptor))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	// step 2: news a service client
	mysqlClient := _02_grpc_demo.NewMysqlClient(cc)

	// step 3: start request gRPCServices
	//// InsertRecord client-streaming RPC
	//log.Printf(mysqlClient.InsertRecord(
	//	`INSERT INTO
	//		users
	//		(name, age, birthday, uuid, role, created_at, updated_at)
	//		VALUE
	//		( "hamming_test", 18, "2004-04-29", "c4e448b8-560a-4e8f-be89-b296ab5accd5",
	//		"stu", "2022-05-01 00:00:00", "2022-05-01 00:00:00");`))
	//
	// SelectRecord unary RPC
	log.Printf(mysqlClient.SelectRecord("SELECT * FROM users WHERE id=1;"))

	// UpdateRecord bidirectional-streaming RPC
	//log.Printf(mysqlClient.UpdateRecord(
	//	`UPDATE users SET role = "stu+tech" WHERE name = "hamming_test"`))
	//
	//// DeleteRecord server-streaming RPC
	//log.Printf(mysqlClient.DeleteRecord(`DELETE FROM users where name = "hamming_test";`))
}
