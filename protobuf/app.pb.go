// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type StoreReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte   `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	AccountToken         string   `protobuf:"bytes,3,opt,name=account_token,json=accountToken,proto3" json:"account_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreReq) Reset()         { *m = StoreReq{} }
func (m *StoreReq) String() string { return proto.CompactTextString(m) }
func (*StoreReq) ProtoMessage()    {}
func (*StoreReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *StoreReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreReq.Unmarshal(m, b)
}
func (m *StoreReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreReq.Marshal(b, m, deterministic)
}
func (m *StoreReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreReq.Merge(m, src)
}
func (m *StoreReq) XXX_Size() int {
	return xxx_messageInfo_StoreReq.Size(m)
}
func (m *StoreReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreReq.DiscardUnknown(m)
}

var xxx_messageInfo_StoreReq proto.InternalMessageInfo

func (m *StoreReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreReq) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *StoreReq) GetAccountToken() string {
	if m != nil {
		return m.AccountToken
	}
	return ""
}

type StoreResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StoreResp) Reset()         { *m = StoreResp{} }
func (m *StoreResp) String() string { return proto.CompactTextString(m) }
func (*StoreResp) ProtoMessage()    {}
func (*StoreResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *StoreResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreResp.Unmarshal(m, b)
}
func (m *StoreResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreResp.Marshal(b, m, deterministic)
}
func (m *StoreResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreResp.Merge(m, src)
}
func (m *StoreResp) XXX_Size() int {
	return xxx_messageInfo_StoreResp.Size(m)
}
func (m *StoreResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreResp.DiscardUnknown(m)
}

var xxx_messageInfo_StoreResp proto.InternalMessageInfo

type GetReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReq) Reset()         { *m = GetReq{} }
func (m *GetReq) String() string { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()    {}
func (*GetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}

func (m *GetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReq.Unmarshal(m, b)
}
func (m *GetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReq.Marshal(b, m, deterministic)
}
func (m *GetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReq.Merge(m, src)
}
func (m *GetReq) XXX_Size() int {
	return xxx_messageInfo_GetReq.Size(m)
}
func (m *GetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetReq proto.InternalMessageInfo

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResp struct {
	Val                  []byte   `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResp) Reset()         { *m = GetResp{} }
func (m *GetResp) String() string { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()    {}
func (*GetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{3}
}

func (m *GetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResp.Unmarshal(m, b)
}
func (m *GetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResp.Marshal(b, m, deterministic)
}
func (m *GetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResp.Merge(m, src)
}
func (m *GetResp) XXX_Size() int {
	return xxx_messageInfo_GetResp.Size(m)
}
func (m *GetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetResp proto.InternalMessageInfo

func (m *GetResp) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

type GetByTokenReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByTokenReq) Reset()         { *m = GetByTokenReq{} }
func (m *GetByTokenReq) String() string { return proto.CompactTextString(m) }
func (*GetByTokenReq) ProtoMessage()    {}
func (*GetByTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{4}
}

func (m *GetByTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByTokenReq.Unmarshal(m, b)
}
func (m *GetByTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByTokenReq.Marshal(b, m, deterministic)
}
func (m *GetByTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByTokenReq.Merge(m, src)
}
func (m *GetByTokenReq) XXX_Size() int {
	return xxx_messageInfo_GetByTokenReq.Size(m)
}
func (m *GetByTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetByTokenReq proto.InternalMessageInfo

func (m *GetByTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetByTokenResp struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByTokenResp) Reset()         { *m = GetByTokenResp{} }
func (m *GetByTokenResp) String() string { return proto.CompactTextString(m) }
func (*GetByTokenResp) ProtoMessage()    {}
func (*GetByTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{5}
}

func (m *GetByTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByTokenResp.Unmarshal(m, b)
}
func (m *GetByTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByTokenResp.Marshal(b, m, deterministic)
}
func (m *GetByTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByTokenResp.Merge(m, src)
}
func (m *GetByTokenResp) XXX_Size() int {
	return xxx_messageInfo_GetByTokenResp.Size(m)
}
func (m *GetByTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetByTokenResp proto.InternalMessageInfo

func (m *GetByTokenResp) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type Account struct {
	MaxCacheKeys         int64    `protobuf:"varint,1,opt,name=max_cache_keys,json=maxCacheKeys,proto3" json:"max_cache_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{6}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetMaxCacheKeys() int64 {
	if m != nil {
		return m.MaxCacheKeys
	}
	return 0
}

func init() {
	proto.RegisterType((*StoreReq)(nil), "rpc.StoreReq")
	proto.RegisterType((*StoreResp)(nil), "rpc.StoreResp")
	proto.RegisterType((*GetReq)(nil), "rpc.GetReq")
	proto.RegisterType((*GetResp)(nil), "rpc.GetResp")
	proto.RegisterType((*GetByTokenReq)(nil), "rpc.GetByTokenReq")
	proto.RegisterType((*GetByTokenResp)(nil), "rpc.GetByTokenResp")
	proto.RegisterType((*Account)(nil), "rpc.Account")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4f, 0x84, 0x30,
	0x10, 0xc5, 0x45, 0xc2, 0xb2, 0x0c, 0x1f, 0x31, 0xa3, 0x07, 0x82, 0x97, 0x4d, 0xfd, 0x08, 0x27,
	0x4c, 0xf0, 0xa0, 0x57, 0xdd, 0x03, 0x07, 0x6f, 0xac, 0x9e, 0x49, 0x6d, 0x9a, 0x98, 0xe0, 0x2e,
	0x95, 0x56, 0xb3, 0xfc, 0xf7, 0xa6, 0xa5, 0xc8, 0x26, 0xee, 0xad, 0xef, 0x37, 0xaf, 0x6f, 0x3a,
	0x53, 0x08, 0xa8, 0x10, 0x85, 0xe8, 0x3b, 0xd5, 0xa1, 0xdb, 0x0b, 0x46, 0x36, 0xb0, 0xdc, 0xa8,
	0xae, 0xe7, 0x35, 0xff, 0xc2, 0x33, 0x70, 0x5b, 0x3e, 0xa4, 0xce, 0xca, 0xc9, 0x83, 0x5a, 0x1f,
	0x35, 0xf9, 0xa1, 0x9f, 0xe9, 0xe9, 0xca, 0xc9, 0xa3, 0x5a, 0x1f, 0xf1, 0x0a, 0x62, 0xca, 0x58,
	0xf7, 0xbd, 0x53, 0x8d, 0xea, 0x5a, 0xbe, 0x4b, 0x5d, 0xe3, 0x8e, 0x2c, 0x7c, 0xd5, 0x8c, 0x84,
	0x10, 0xd8, 0x50, 0x29, 0x48, 0x06, 0x8b, 0x8a, 0xab, 0xa3, 0xf9, 0xe4, 0x12, 0x7c, 0x53, 0x93,
	0x62, 0x6a, 0xe5, 0xfc, 0xb5, 0x22, 0x37, 0x10, 0x57, 0x5c, 0x3d, 0x0f, 0x26, 0x53, 0xdf, 0xbf,
	0x00, 0x6f, 0xec, 0x39, 0x26, 0x8c, 0x82, 0x3c, 0x42, 0x72, 0x68, 0x93, 0x02, 0x6f, 0xc1, 0xb7,
	0xcf, 0x31, 0xce, 0xb0, 0x8c, 0x8a, 0x5e, 0xb0, 0xe2, 0x69, 0x64, 0xf5, 0x54, 0x24, 0x77, 0xe0,
	0x5b, 0x86, 0xd7, 0x90, 0x6c, 0xe9, 0xbe, 0x61, 0x94, 0x7d, 0xf0, 0xa6, 0xe5, 0x83, 0x34, 0x37,
	0xdd, 0x3a, 0xda, 0xd2, 0xfd, 0x5a, 0xc3, 0x17, 0x3e, 0xc8, 0xf2, 0x0d, 0x3c, 0x23, 0x30, 0x07,
	0xcf, 0x0c, 0x88, 0xb1, 0x49, 0x9e, 0x36, 0x98, 0x25, 0x87, 0x52, 0x0a, 0x72, 0x82, 0x04, 0xdc,
	0x8a, 0x2b, 0x0c, 0x4d, 0x61, 0xdc, 0x43, 0x16, 0xcd, 0x42, 0x7b, 0xca, 0x35, 0x2c, 0xed, 0x3b,
	0x24, 0x3e, 0x00, 0xcc, 0xd3, 0x20, 0x4e, 0xce, 0x79, 0x0b, 0xd9, 0xf9, 0x3f, 0xa6, 0x43, 0xde,
	0x17, 0xe6, 0x53, 0xef, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x56, 0xa9, 0x4b, 0xd4, 0xe1, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error) {
	out := new(StoreResp)
	err := c.cc.Invoke(ctx, "/rpc.Cache/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/rpc.Cache/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	Store(context.Context, *StoreReq) (*StoreResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountsClient interface {
	GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error) {
	out := new(GetByTokenResp)
	err := c.cc.Invoke(ctx, "/rpc.Accounts/GetByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
type AccountsServer interface {
	GetByToken(context.Context, *GetByTokenReq) (*GetByTokenResp, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Accounts/GetByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetByToken(ctx, req.(*GetByTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByToken",
			Handler:    _Accounts_GetByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
