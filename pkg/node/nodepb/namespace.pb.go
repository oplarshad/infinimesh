// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/node/nodepb/namespace.proto

package nodepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Namespace struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{0}
}

func (m *Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespace.Unmarshal(m, b)
}
func (m *Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespace.Marshal(b, m, deterministic)
}
func (m *Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespace.Merge(m, src)
}
func (m *Namespace) XXX_Size() int {
	return xxx_messageInfo_Namespace.Size(m)
}
func (m *Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_Namespace proto.InternalMessageInfo

func (m *Namespace) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Namespace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateNamespaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceRequest) Reset()         { *m = CreateNamespaceRequest{} }
func (m *CreateNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceRequest) ProtoMessage()    {}
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{1}
}

func (m *CreateNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceRequest.Unmarshal(m, b)
}
func (m *CreateNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceRequest.Merge(m, src)
}
func (m *CreateNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceRequest.Size(m)
}
func (m *CreateNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceRequest proto.InternalMessageInfo

func (m *CreateNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetNamespaceRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNamespaceRequest) Reset()         { *m = GetNamespaceRequest{} }
func (m *GetNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*GetNamespaceRequest) ProtoMessage()    {}
func (*GetNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{2}
}

func (m *GetNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNamespaceRequest.Unmarshal(m, b)
}
func (m *GetNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *GetNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNamespaceRequest.Merge(m, src)
}
func (m *GetNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_GetNamespaceRequest.Size(m)
}
func (m *GetNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNamespaceRequest proto.InternalMessageInfo

func (m *GetNamespaceRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListNamespacesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNamespacesRequest) Reset()         { *m = ListNamespacesRequest{} }
func (m *ListNamespacesRequest) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesRequest) ProtoMessage()    {}
func (*ListNamespacesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{3}
}

func (m *ListNamespacesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesRequest.Unmarshal(m, b)
}
func (m *ListNamespacesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesRequest.Marshal(b, m, deterministic)
}
func (m *ListNamespacesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesRequest.Merge(m, src)
}
func (m *ListNamespacesRequest) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesRequest.Size(m)
}
func (m *ListNamespacesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesRequest proto.InternalMessageInfo

type ListNamespacesForAccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNamespacesForAccountRequest) Reset()         { *m = ListNamespacesForAccountRequest{} }
func (m *ListNamespacesForAccountRequest) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesForAccountRequest) ProtoMessage()    {}
func (*ListNamespacesForAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{4}
}

func (m *ListNamespacesForAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Unmarshal(m, b)
}
func (m *ListNamespacesForAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Marshal(b, m, deterministic)
}
func (m *ListNamespacesForAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesForAccountRequest.Merge(m, src)
}
func (m *ListNamespacesForAccountRequest) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesForAccountRequest.Size(m)
}
func (m *ListNamespacesForAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesForAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesForAccountRequest proto.InternalMessageInfo

func (m *ListNamespacesForAccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type ListNamespacesResponse struct {
	Namespaces           []*Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListNamespacesResponse) Reset()         { *m = ListNamespacesResponse{} }
func (m *ListNamespacesResponse) String() string { return proto.CompactTextString(m) }
func (*ListNamespacesResponse) ProtoMessage()    {}
func (*ListNamespacesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ba94f8fc77af896, []int{5}
}

func (m *ListNamespacesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNamespacesResponse.Unmarshal(m, b)
}
func (m *ListNamespacesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNamespacesResponse.Marshal(b, m, deterministic)
}
func (m *ListNamespacesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNamespacesResponse.Merge(m, src)
}
func (m *ListNamespacesResponse) XXX_Size() int {
	return xxx_messageInfo_ListNamespacesResponse.Size(m)
}
func (m *ListNamespacesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNamespacesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNamespacesResponse proto.InternalMessageInfo

func (m *ListNamespacesResponse) GetNamespaces() []*Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Namespace)(nil), "infinimesh.node.Namespace")
	proto.RegisterType((*CreateNamespaceRequest)(nil), "infinimesh.node.CreateNamespaceRequest")
	proto.RegisterType((*GetNamespaceRequest)(nil), "infinimesh.node.GetNamespaceRequest")
	proto.RegisterType((*ListNamespacesRequest)(nil), "infinimesh.node.ListNamespacesRequest")
	proto.RegisterType((*ListNamespacesForAccountRequest)(nil), "infinimesh.node.ListNamespacesForAccountRequest")
	proto.RegisterType((*ListNamespacesResponse)(nil), "infinimesh.node.ListNamespacesResponse")
}

func init() { proto.RegisterFile("pkg/node/nodepb/namespace.proto", fileDescriptor_2ba94f8fc77af896) }

var fileDescriptor_2ba94f8fc77af896 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x2a, 0x95, 0x8c, 0x92, 0xc2, 0x14, 0x6b, 0x08, 0x42, 0x4b, 0x10, 0xdb, 0x83,
	0x24, 0xd2, 0xde, 0xf4, 0x64, 0x05, 0xbd, 0x88, 0x84, 0x22, 0x1e, 0xbc, 0xa5, 0xc9, 0xaa, 0x8b,
	0x64, 0x37, 0x66, 0x93, 0xf7, 0xf0, 0x91, 0x25, 0x69, 0xfe, 0x35, 0x1b, 0xdb, 0x5e, 0x96, 0xdd,
	0xd9, 0xdf, 0x7c, 0xb3, 0xfb, 0x0d, 0x03, 0xe3, 0xe8, 0xfb, 0xd3, 0x61, 0x3c, 0x20, 0xf9, 0x12,
	0xad, 0x1d, 0xe6, 0x85, 0x44, 0x44, 0x9e, 0x4f, 0xec, 0x28, 0xe6, 0x09, 0xc7, 0x01, 0x65, 0x1f,
	0x94, 0xd1, 0x90, 0x88, 0x2f, 0x3b, 0x43, 0x2c, 0x07, 0xb4, 0x97, 0x92, 0x41, 0x1d, 0x54, 0x1a,
	0x18, 0xca, 0x44, 0x99, 0x69, 0x2b, 0x95, 0x06, 0x88, 0x70, 0x94, 0x09, 0x18, 0x6a, 0x1e, 0xc9,
	0xf7, 0xd6, 0x35, 0x8c, 0x1e, 0x62, 0xe2, 0x25, 0xa4, 0x4a, 0x5b, 0x91, 0x9f, 0x94, 0x88, 0xa4,
	0xa2, 0x95, 0x06, 0xbd, 0x80, 0xe1, 0x13, 0x49, 0x24, 0xf4, 0x02, 0xb4, 0xea, 0x65, 0x05, 0x5f,
	0x07, 0xac, 0x73, 0x38, 0x7b, 0xa6, 0xa2, 0xce, 0x12, 0x45, 0x9a, 0x75, 0x07, 0xe3, 0xed, 0x8b,
	0x47, 0x1e, 0xdf, 0xfb, 0x3e, 0x4f, 0x59, 0x52, 0x2a, 0x1b, 0x70, 0xec, 0x6d, 0x22, 0x85, 0x6e,
	0x79, 0xb4, 0x5e, 0x61, 0xd4, 0x56, 0x15, 0x11, 0x67, 0x82, 0xe0, 0x2d, 0x40, 0x55, 0x5c, 0x18,
	0xca, 0xa4, 0x37, 0x3b, 0x99, 0x9b, 0x76, 0xcb, 0x29, 0xbb, 0xfe, 0x44, 0x83, 0x9e, 0xff, 0xf6,
	0x00, 0x6a, 0x49, 0x7c, 0x83, 0x41, 0xcb, 0x1d, 0x9c, 0x4a, 0x4a, 0xdd, 0xfe, 0x99, 0x3b, 0x4a,
	0xa2, 0x0b, 0xa7, 0x4d, 0x1f, 0xf1, 0x52, 0x62, 0x3b, 0x6c, 0xde, 0xa9, 0xe8, 0x81, 0xbe, 0x6d,
	0x07, 0x5e, 0x49, 0x74, 0x67, 0x17, 0xcc, 0xe9, 0x5e, 0xae, 0xf0, 0x35, 0x05, 0xe3, 0xbf, 0x76,
	0xe1, 0xcd, 0x1e, 0x11, 0xa9, 0xb3, 0x07, 0x97, 0x5d, 0x3a, 0x30, 0xf4, 0x79, 0xd8, 0xa6, 0x97,
	0x7a, 0x85, 0xba, 0xd9, 0x28, 0xb8, 0xca, 0x7b, 0x7f, 0x33, 0x24, 0xeb, 0x7e, 0x3e, 0x1b, 0x8b,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xc5, 0x02, 0xf1, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespacesClient is the client API for Namespaces service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespacesClient interface {
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error)
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	ListNamespacesForAccount(ctx context.Context, in *ListNamespacesForAccountRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type namespacesClient struct {
	cc *grpc.ClientConn
}

func NewNamespacesClient(cc *grpc.ClientConn) NamespacesClient {
	return &namespacesClient{cc}
}

func (c *namespacesClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Namespaces/CreateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*Namespace, error) {
	out := new(Namespace)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Namespaces/GetNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Namespaces/ListNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespacesClient) ListNamespacesForAccount(ctx context.Context, in *ListNamespacesForAccountRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Namespaces/ListNamespacesForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespacesServer is the server API for Namespaces service.
type NamespacesServer interface {
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*Namespace, error)
	GetNamespace(context.Context, *GetNamespaceRequest) (*Namespace, error)
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	ListNamespacesForAccount(context.Context, *ListNamespacesForAccountRequest) (*ListNamespacesResponse, error)
}

func RegisterNamespacesServer(s *grpc.Server, srv NamespacesServer) {
	s.RegisterService(&_Namespaces_serviceDesc, srv)
}

func _Namespaces_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Namespaces/CreateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Namespaces/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Namespaces/ListNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Namespaces_ListNamespacesForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesForAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServer).ListNamespacesForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Namespaces/ListNamespacesForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServer).ListNamespacesForAccount(ctx, req.(*ListNamespacesForAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Namespaces_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.node.Namespaces",
	HandlerType: (*NamespacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _Namespaces_CreateNamespace_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _Namespaces_GetNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _Namespaces_ListNamespaces_Handler,
		},
		{
			MethodName: "ListNamespacesForAccount",
			Handler:    _Namespaces_ListNamespacesForAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/node/nodepb/namespace.proto",
}