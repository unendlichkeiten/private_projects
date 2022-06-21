package _01_grpc_demo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"_01_grpc_demo/conf"
	"_01_grpc_demo/pb"
)

type MysqlServer struct{}

func NewMysqlServer() *MysqlServer {
	return &MysqlServer{}
}

func (*MysqlServer) SelectRecord(ctx context.Context, req *pb.MysqlReq) (*pb.MysqlRes, error) {

	sql := req.GetSql()
	log.Printf("receives a select record request with sql : %s", sql)

	// test sleep
	// time.Sleep(time.Second * 5)

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	users, res := make([]*conf.User, 0), &pb.MysqlRes{}
	if err := conf.DB.Raw(sql).Scan(&users).Error; err != nil {
		return nil, err
	}

	rows, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	res.Code = 200
	res.Msg = "select success"
	res.Data = string(rows)

	return res, nil
}

func (*MysqlServer) InsertRecord(stream pb.MysqlService_InsertRecordServer) error {

	var sql string
	for {
		req, err := stream.Recv()
		if err == io.EOF || req == nil {
			log.Printf("read ends")
			break
		}

		if len(sql) == 0 {
			sql = req.Sql
		} else {
			sql = fmt.Sprintf("%s;%s", sql, req.Sql)
		}
	}

	log.Printf("receives a insert record request with sql : %s", sql)

	// test sleep
	// time.Sleep(time.Second * 5)

	res := &pb.MysqlRes{
		Code: 200,
		Msg:  fmt.Sprintf("insert records sql [%s] succ", sql),
		Data: _empty,
	}

	if err := conf.DB.Exec(sql).Error; err != nil {
		res.Code = 500
		res.Msg = fmt.Sprintf("exec %s err [%s]", sql, err)
		return err
	}

	// send res
	if err := stream.SendAndClose(res); err != nil {
		errs := fmt.Sprintf("send res [%s] err [%s]", res.Msg, err)
		return errors.New(errs)
	}

	return nil
}

func (*MysqlServer) DeleteRecord(
	req *pb.MysqlReq,
	stream pb.MysqlService_DeleteRecordServer,
) error {

	sql := req.GetSql()
	log.Printf("receives a delete record request with sql : %s", sql)

	// test sleep
	// time.Sleep(time.Second * 5)

	res := &pb.MysqlRes{
		Code: 200,
		Msg:  fmt.Sprintf("delete records sql [%s] succ", sql),
		Data: _empty,
	}

	if err := conf.DB.Exec(sql).Error; err != nil {
		res.Code = 500
		res.Msg = fmt.Sprintf("exec %s err [%s]", sql, err)
	}

	// send res
	if err := stream.Send(res); err != nil {
		resBytes, _ := json.Marshal(res)
		return fmt.Errorf("send res [%s] to stream err [%s]", string(resBytes), err)
	}

	return nil
}

func (*MysqlServer) UpdateRecord(stream pb.MysqlService_UpdateRecordServer) error {
	var sql string
	for {
		req, err := stream.Recv()
		if err == io.EOF || req == nil {
			log.Printf("read ends\n")
			break
		}

		if len(sql) == 0 {
			sql = req.Sql
		} else {
			sql = fmt.Sprintf("%s;%s", sql, req.Sql)
		}
	}

	log.Printf("receives a update record request with sql : %s\n", sql)

	res := &pb.MysqlRes{
		Code: 200,
		Msg:  fmt.Sprintf("update records sql [%s] succ", sql),
		Data: _empty,
	}

	if err := conf.DB.Exec(sql).Error; err != nil {
		res.Code = 500
		res.Msg = fmt.Sprintf("exec %s err [%s]", sql, err)
	}

	if err := stream.Send(res); err != nil {
		resBytes, _ := json.Marshal(res)
		return fmt.Errorf("send res [%s] to stream err [%s]", string(resBytes), err)
	}

	return nil
}
