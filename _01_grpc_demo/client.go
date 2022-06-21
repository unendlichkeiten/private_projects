package _01_grpc_demo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	"_01_grpc_demo/pb"
)

const (
	_empty = ""
)

type MysqlClient struct {
	service pb.MysqlServiceClient
}

func NewMysqlClient(cc *grpc.ClientConn) *MysqlClient {
	service := pb.NewMysqlServiceClient(cc)

	return &MysqlClient{service}
}

func (mysqlClient *MysqlClient) SelectRecord(sql string) string {
	req := &pb.MysqlReq{
		Sql: sql,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := mysqlClient.service.SelectRecord(ctx, req)
	if err != nil {
		log.Fatal("select record err ", err)
		return _empty
	}

	return fmt.Sprintf("select record success %s", res.GetData())
}

func (mysqlClient *MysqlClient) InsertRecord(sql string) string {

	req := &pb.MysqlReq{
		Sql: sql,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := mysqlClient.service.InsertRecord(ctx)
	if err != nil {
		log.Fatal("insert record err ", err)
		return _empty
	}

	if err = stream.Send(req); err != nil {
		log.Fatal("send req err ", err)
		return _empty
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("receive res err ", err)
		return _empty
	}

	resBytes, _ := json.Marshal(res)
	return string(resBytes)
}

func (mysqlClient *MysqlClient) DeleteRecord(sql string) string {

	req := &pb.MysqlReq{
		Sql: sql,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := mysqlClient.service.DeleteRecord(ctx, req)
	if err != nil {
		log.Fatal("delete record err ", err)
		return _empty
	}

	resBytes := make([]byte, 0)
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
			return _empty
		}

		resBytes, _ = json.Marshal(res)
	}

	return string(resBytes)
}

func (mysqlClient *MysqlClient) UpdateRecord(sql string) string {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := mysqlClient.service.UpdateRecord(ctx)
	if err != nil {
		log.Fatal("update record err ", err)
		return _empty
	}

	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	req := &pb.MysqlReq{
		Sql: sql,
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Sprintf("send stream err: %s", err)
	}
	err = stream.CloseSend()
	if err != nil {
		return fmt.Sprintf("cannot close send: %s", err)
	}

	err = <-waitResponse
	return fmt.Sprintf("%s", err)
}
