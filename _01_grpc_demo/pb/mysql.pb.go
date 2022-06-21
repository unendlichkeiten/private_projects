// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: mysql.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MysqlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sql string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
}

func (x *MysqlReq) Reset() {
	*x = MysqlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MysqlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlReq) ProtoMessage() {}

func (x *MysqlReq) ProtoReflect() protoreflect.Message {
	mi := &file_mysql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlReq.ProtoReflect.Descriptor instead.
func (*MysqlReq) Descriptor() ([]byte, []int) {
	return file_mysql_proto_rawDescGZIP(), []int{0}
}

func (x *MysqlReq) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

type MysqlRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *MysqlRes) Reset() {
	*x = MysqlRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mysql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MysqlRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MysqlRes) ProtoMessage() {}

func (x *MysqlRes) ProtoReflect() protoreflect.Message {
	mi := &file_mysql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MysqlRes.ProtoReflect.Descriptor instead.
func (*MysqlRes) Descriptor() ([]byte, []int) {
	return file_mysql_proto_rawDescGZIP(), []int{1}
}

func (x *MysqlRes) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MysqlRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *MysqlRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_mysql_proto protoreflect.FileDescriptor

var file_mysql_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x5f,
	0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x1c, 0x0a, 0x08,
	0x4d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x22, 0x44, 0x0a, 0x08, 0x4d, 0x79,
	0x73, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xa6, 0x02, 0x0a, 0x0c, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x42, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x17, 0x2e, 0x5f, 0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x5f, 0x30, 0x31,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x5f, 0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17,
	0x2e, 0x5f, 0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d,
	0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x44, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x5f, 0x30,
	0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71,
	0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x5f, 0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x46, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x12, 0x17, 0x2e, 0x5f, 0x30, 0x31, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x5f, 0x30, 0x31,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x79, 0x73, 0x71, 0x6c,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mysql_proto_rawDescOnce sync.Once
	file_mysql_proto_rawDescData = file_mysql_proto_rawDesc
)

func file_mysql_proto_rawDescGZIP() []byte {
	file_mysql_proto_rawDescOnce.Do(func() {
		file_mysql_proto_rawDescData = protoimpl.X.CompressGZIP(file_mysql_proto_rawDescData)
	})
	return file_mysql_proto_rawDescData
}

var file_mysql_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mysql_proto_goTypes = []interface{}{
	(*MysqlReq)(nil), // 0: _01_grpc_demo.MysqlReq
	(*MysqlRes)(nil), // 1: _01_grpc_demo.MysqlRes
}
var file_mysql_proto_depIdxs = []int32{
	0, // 0: _01_grpc_demo.MysqlService.SelectRecord:input_type -> _01_grpc_demo.MysqlReq
	0, // 1: _01_grpc_demo.MysqlService.InsertRecord:input_type -> _01_grpc_demo.MysqlReq
	0, // 2: _01_grpc_demo.MysqlService.DeleteRecord:input_type -> _01_grpc_demo.MysqlReq
	0, // 3: _01_grpc_demo.MysqlService.UpdateRecord:input_type -> _01_grpc_demo.MysqlReq
	1, // 4: _01_grpc_demo.MysqlService.SelectRecord:output_type -> _01_grpc_demo.MysqlRes
	1, // 5: _01_grpc_demo.MysqlService.InsertRecord:output_type -> _01_grpc_demo.MysqlRes
	1, // 6: _01_grpc_demo.MysqlService.DeleteRecord:output_type -> _01_grpc_demo.MysqlRes
	1, // 7: _01_grpc_demo.MysqlService.UpdateRecord:output_type -> _01_grpc_demo.MysqlRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mysql_proto_init() }
func file_mysql_proto_init() {
	if File_mysql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mysql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MysqlReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mysql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MysqlRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mysql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mysql_proto_goTypes,
		DependencyIndexes: file_mysql_proto_depIdxs,
		MessageInfos:      file_mysql_proto_msgTypes,
	}.Build()
	File_mysql_proto = out.File
	file_mysql_proto_rawDesc = nil
	file_mysql_proto_goTypes = nil
	file_mysql_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MysqlServiceClient is the client API for MysqlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MysqlServiceClient interface {
	SelectRecord(ctx context.Context, in *MysqlReq, opts ...grpc.CallOption) (*MysqlRes, error)
	InsertRecord(ctx context.Context, opts ...grpc.CallOption) (MysqlService_InsertRecordClient, error)
	DeleteRecord(ctx context.Context, in *MysqlReq, opts ...grpc.CallOption) (MysqlService_DeleteRecordClient, error)
	UpdateRecord(ctx context.Context, opts ...grpc.CallOption) (MysqlService_UpdateRecordClient, error)
}

type mysqlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMysqlServiceClient(cc grpc.ClientConnInterface) MysqlServiceClient {
	return &mysqlServiceClient{cc}
}

func (c *mysqlServiceClient) SelectRecord(ctx context.Context, in *MysqlReq, opts ...grpc.CallOption) (*MysqlRes, error) {
	out := new(MysqlRes)
	err := c.cc.Invoke(ctx, "/_01_grpc_demo.MysqlService/SelectRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mysqlServiceClient) InsertRecord(ctx context.Context, opts ...grpc.CallOption) (MysqlService_InsertRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MysqlService_serviceDesc.Streams[0], "/_01_grpc_demo.MysqlService/InsertRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &mysqlServiceInsertRecordClient{stream}
	return x, nil
}

type MysqlService_InsertRecordClient interface {
	Send(*MysqlReq) error
	CloseAndRecv() (*MysqlRes, error)
	grpc.ClientStream
}

type mysqlServiceInsertRecordClient struct {
	grpc.ClientStream
}

func (x *mysqlServiceInsertRecordClient) Send(m *MysqlReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mysqlServiceInsertRecordClient) CloseAndRecv() (*MysqlRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MysqlRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mysqlServiceClient) DeleteRecord(ctx context.Context, in *MysqlReq, opts ...grpc.CallOption) (MysqlService_DeleteRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MysqlService_serviceDesc.Streams[1], "/_01_grpc_demo.MysqlService/DeleteRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &mysqlServiceDeleteRecordClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MysqlService_DeleteRecordClient interface {
	Recv() (*MysqlRes, error)
	grpc.ClientStream
}

type mysqlServiceDeleteRecordClient struct {
	grpc.ClientStream
}

func (x *mysqlServiceDeleteRecordClient) Recv() (*MysqlRes, error) {
	m := new(MysqlRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mysqlServiceClient) UpdateRecord(ctx context.Context, opts ...grpc.CallOption) (MysqlService_UpdateRecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MysqlService_serviceDesc.Streams[2], "/_01_grpc_demo.MysqlService/UpdateRecord", opts...)
	if err != nil {
		return nil, err
	}
	x := &mysqlServiceUpdateRecordClient{stream}
	return x, nil
}

type MysqlService_UpdateRecordClient interface {
	Send(*MysqlReq) error
	Recv() (*MysqlRes, error)
	grpc.ClientStream
}

type mysqlServiceUpdateRecordClient struct {
	grpc.ClientStream
}

func (x *mysqlServiceUpdateRecordClient) Send(m *MysqlReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mysqlServiceUpdateRecordClient) Recv() (*MysqlRes, error) {
	m := new(MysqlRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MysqlServiceServer is the server API for MysqlService service.
type MysqlServiceServer interface {
	SelectRecord(context.Context, *MysqlReq) (*MysqlRes, error)
	InsertRecord(MysqlService_InsertRecordServer) error
	DeleteRecord(*MysqlReq, MysqlService_DeleteRecordServer) error
	UpdateRecord(MysqlService_UpdateRecordServer) error
}

// UnimplementedMysqlServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMysqlServiceServer struct {
}

func (*UnimplementedMysqlServiceServer) SelectRecord(context.Context, *MysqlReq) (*MysqlRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectRecord not implemented")
}
func (*UnimplementedMysqlServiceServer) InsertRecord(MysqlService_InsertRecordServer) error {
	return status.Errorf(codes.Unimplemented, "method InsertRecord not implemented")
}
func (*UnimplementedMysqlServiceServer) DeleteRecord(*MysqlReq, MysqlService_DeleteRecordServer) error {
	return status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (*UnimplementedMysqlServiceServer) UpdateRecord(MysqlService_UpdateRecordServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}

func RegisterMysqlServiceServer(s *grpc.Server, srv MysqlServiceServer) {
	s.RegisterService(&_MysqlService_serviceDesc, srv)
}

func _MysqlService_SelectRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MysqlReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MysqlServiceServer).SelectRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/_01_grpc_demo.MysqlService/SelectRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MysqlServiceServer).SelectRecord(ctx, req.(*MysqlReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MysqlService_InsertRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MysqlServiceServer).InsertRecord(&mysqlServiceInsertRecordServer{stream})
}

type MysqlService_InsertRecordServer interface {
	SendAndClose(*MysqlRes) error
	Recv() (*MysqlReq, error)
	grpc.ServerStream
}

type mysqlServiceInsertRecordServer struct {
	grpc.ServerStream
}

func (x *mysqlServiceInsertRecordServer) SendAndClose(m *MysqlRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mysqlServiceInsertRecordServer) Recv() (*MysqlReq, error) {
	m := new(MysqlReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MysqlService_DeleteRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MysqlReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MysqlServiceServer).DeleteRecord(m, &mysqlServiceDeleteRecordServer{stream})
}

type MysqlService_DeleteRecordServer interface {
	Send(*MysqlRes) error
	grpc.ServerStream
}

type mysqlServiceDeleteRecordServer struct {
	grpc.ServerStream
}

func (x *mysqlServiceDeleteRecordServer) Send(m *MysqlRes) error {
	return x.ServerStream.SendMsg(m)
}

func _MysqlService_UpdateRecord_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MysqlServiceServer).UpdateRecord(&mysqlServiceUpdateRecordServer{stream})
}

type MysqlService_UpdateRecordServer interface {
	Send(*MysqlRes) error
	Recv() (*MysqlReq, error)
	grpc.ServerStream
}

type mysqlServiceUpdateRecordServer struct {
	grpc.ServerStream
}

func (x *mysqlServiceUpdateRecordServer) Send(m *MysqlRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mysqlServiceUpdateRecordServer) Recv() (*MysqlReq, error) {
	m := new(MysqlReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MysqlService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "_01_grpc_demo.MysqlService",
	HandlerType: (*MysqlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectRecord",
			Handler:    _MysqlService_SelectRecord_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InsertRecord",
			Handler:       _MysqlService_InsertRecord_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DeleteRecord",
			Handler:       _MysqlService_DeleteRecord_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateRecord",
			Handler:       _MysqlService_UpdateRecord_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "mysql.proto",
}
