// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package echopb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TestReq struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReq) Reset()         { *m = TestReq{} }
func (m *TestReq) String() string { return proto.CompactTextString(m) }
func (*TestReq) ProtoMessage()    {}
func (*TestReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReq.Unmarshal(m, b)
}
func (m *TestReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReq.Marshal(b, m, deterministic)
}
func (m *TestReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReq.Merge(m, src)
}
func (m *TestReq) XXX_Size() int {
	return xxx_messageInfo_TestReq.Size(m)
}
func (m *TestReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReq.DiscardUnknown(m)
}

var xxx_messageInfo_TestReq proto.InternalMessageInfo

func (m *TestReq) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TestResp struct {
	Error                string      `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Message              string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code                 int32       `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Details              *TestDetail `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TestResp) Reset()         { *m = TestResp{} }
func (m *TestResp) String() string { return proto.CompactTextString(m) }
func (*TestResp) ProtoMessage()    {}
func (*TestResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResp.Unmarshal(m, b)
}
func (m *TestResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResp.Marshal(b, m, deterministic)
}
func (m *TestResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResp.Merge(m, src)
}
func (m *TestResp) XXX_Size() int {
	return xxx_messageInfo_TestResp.Size(m)
}
func (m *TestResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResp.DiscardUnknown(m)
}

var xxx_messageInfo_TestResp proto.InternalMessageInfo

func (m *TestResp) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TestResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TestResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TestResp) GetDetails() *TestDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type TestDetail struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestDetail) Reset()         { *m = TestDetail{} }
func (m *TestDetail) String() string { return proto.CompactTextString(m) }
func (*TestDetail) ProtoMessage()    {}
func (*TestDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *TestDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestDetail.Unmarshal(m, b)
}
func (m *TestDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestDetail.Marshal(b, m, deterministic)
}
func (m *TestDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDetail.Merge(m, src)
}
func (m *TestDetail) XXX_Size() int {
	return xxx_messageInfo_TestDetail.Size(m)
}
func (m *TestDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TestDetail proto.InternalMessageInfo

func (m *TestDetail) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type MakeCodingRequest struct {
	ServerName           string   `protobuf:"bytes,13,opt,name=serverName,proto3" json:"serverName,omitempty"`
	ModuleName           string   `protobuf:"bytes,14,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	TableName            string   `protobuf:"bytes,10,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Name                 string   `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	DatabaseName         string   `protobuf:"bytes,15,opt,name=databaseName,proto3" json:"databaseName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeCodingRequest) Reset()         { *m = MakeCodingRequest{} }
func (m *MakeCodingRequest) String() string { return proto.CompactTextString(m) }
func (*MakeCodingRequest) ProtoMessage()    {}
func (*MakeCodingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *MakeCodingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeCodingRequest.Unmarshal(m, b)
}
func (m *MakeCodingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeCodingRequest.Marshal(b, m, deterministic)
}
func (m *MakeCodingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeCodingRequest.Merge(m, src)
}
func (m *MakeCodingRequest) XXX_Size() int {
	return xxx_messageInfo_MakeCodingRequest.Size(m)
}
func (m *MakeCodingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeCodingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeCodingRequest proto.InternalMessageInfo

func (m *MakeCodingRequest) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *MakeCodingRequest) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *MakeCodingRequest) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *MakeCodingRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MakeCodingRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func init() {
	proto.RegisterType((*TestReq)(nil), "echopb.TestReq")
	proto.RegisterType((*TestResp)(nil), "echopb.TestResp")
	proto.RegisterType((*TestDetail)(nil), "echopb.TestDetail")
	proto.RegisterType((*MakeCodingRequest)(nil), "echopb.MakeCodingRequest")
}

func init() {
	proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e)
}

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4e, 0xeb, 0x40,
	0x0c, 0x87, 0x95, 0xbe, 0xfe, 0x75, 0x1f, 0xb4, 0x58, 0x2c, 0xa2, 0xaa, 0x82, 0x2a, 0xab, 0x2e,
	0x50, 0x23, 0xca, 0x09, 0x10, 0x48, 0x88, 0x45, 0x59, 0x44, 0x5c, 0xc0, 0x69, 0xac, 0x10, 0x91,
	0x66, 0xd2, 0xf1, 0xb4, 0x2b, 0x56, 0x5c, 0x81, 0x53, 0x70, 0x1e, 0xae, 0xc0, 0x41, 0xd0, 0xcc,
	0x50, 0xb5, 0x5d, 0xb0, 0xb3, 0xbf, 0xef, 0x37, 0x8a, 0x1d, 0x03, 0x18, 0x16, 0x33, 0xab, 0xb5,
	0x32, 0x0a, 0xdb, 0xbc, 0x7c, 0x51, 0x75, 0x3a, 0x1a, 0xe7, 0x4a, 0xe5, 0x25, 0xc7, 0x54, 0x17,
	0x31, 0x55, 0x95, 0x32, 0x64, 0x0a, 0x55, 0x89, 0x4f, 0x45, 0x97, 0xd0, 0x79, 0x66, 0x31, 0x09,
	0xaf, 0xf1, 0x1c, 0x5a, 0x5b, 0x2a, 0x37, 0x1c, 0x06, 0x93, 0x60, 0xda, 0x4b, 0x7c, 0x13, 0xbd,
	0x41, 0xd7, 0x07, 0xa4, 0xb6, 0x09, 0xd6, 0x5a, 0xe9, 0x5d, 0xc2, 0x35, 0x18, 0x42, 0x67, 0xc5,
	0x22, 0x94, 0x73, 0xd8, 0x70, 0x7c, 0xd7, 0x22, 0x42, 0x73, 0xa9, 0x32, 0x0e, 0xff, 0x4d, 0x82,
	0x69, 0x2b, 0x71, 0x35, 0x5e, 0x41, 0x27, 0x63, 0x43, 0x45, 0x29, 0x61, 0x73, 0x12, 0x4c, 0xfb,
	0x73, 0x9c, 0xf9, 0x41, 0x67, 0xf6, 0x33, 0xf7, 0x4e, 0x25, 0xbb, 0x48, 0x14, 0x01, 0xec, 0xf1,
	0x1f, 0x13, 0x7e, 0x06, 0x70, 0xb6, 0xa0, 0x57, 0xbe, 0x53, 0x59, 0x51, 0xe5, 0x09, 0xaf, 0x37,
	0x2c, 0x06, 0x2f, 0x00, 0x84, 0xf5, 0x96, 0xf5, 0x13, 0xad, 0x38, 0x3c, 0x71, 0x0f, 0x0e, 0x88,
	0xf5, 0x2b, 0x95, 0x6d, 0x4a, 0x76, 0xfe, 0xd4, 0xfb, 0x3d, 0xc1, 0x31, 0xf4, 0x0c, 0xa5, 0xbf,
	0x1a, 0x9c, 0xde, 0x03, 0xbb, 0x59, 0x65, 0x45, 0xdf, 0x09, 0x57, 0x63, 0x04, 0xff, 0x33, 0x32,
	0x94, 0x92, 0xf8, 0x47, 0x03, 0xe7, 0x8e, 0xd8, 0xfc, 0x11, 0x9a, 0x76, 0x1f, 0xbc, 0x05, 0x78,
	0x60, 0x63, 0xcb, 0x85, 0xe4, 0x38, 0x38, 0xfc, 0x05, 0x09, 0xaf, 0x47, 0xc3, 0x63, 0x20, 0x75,
	0x34, 0x7c, 0xff, 0xfa, 0xfe, 0x68, 0x00, 0x76, 0xe3, 0xed, 0x75, 0x6c, 0xaf, 0x9c, 0xb6, 0xdd,
	0x01, 0x6f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x1b, 0x8a, 0x4b, 0xf4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	//测试方法
	GetTestMsg(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) GetTestMsg(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestResp, error) {
	out := new(TestResp)
	err := c.cc.Invoke(ctx, "/echopb.Test/GetTestMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	//测试方法
	GetTestMsg(context.Context, *TestReq) (*TestResp, error)
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) GetTestMsg(ctx context.Context, req *TestReq) (*TestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestMsg not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_GetTestMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).GetTestMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echopb.Test/GetTestMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).GetTestMsg(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echopb.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTestMsg",
			Handler:    _Test_GetTestMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
