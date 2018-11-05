// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/nodes.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type BareMetalNode struct {
	// Unique node identifier.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Hostname. Is not unique.
	Hostname             string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalNode) Reset()         { *m = BareMetalNode{} }
func (m *BareMetalNode) String() string { return proto.CompactTextString(m) }
func (*BareMetalNode) ProtoMessage()    {}
func (*BareMetalNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_dbb3174ed4d9d994, []int{0}
}
func (m *BareMetalNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalNode.Unmarshal(m, b)
}
func (m *BareMetalNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalNode.Marshal(b, m, deterministic)
}
func (dst *BareMetalNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalNode.Merge(dst, src)
}
func (m *BareMetalNode) XXX_Size() int {
	return xxx_messageInfo_BareMetalNode.Size(m)
}
func (m *BareMetalNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalNode.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalNode proto.InternalMessageInfo

func (m *BareMetalNode) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BareMetalNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BareMetalNode) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type AddBareMetalNodeRequest struct {
	// Unique node name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Hostname. Is not unique.
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddBareMetalNodeRequest) Reset()         { *m = AddBareMetalNodeRequest{} }
func (m *AddBareMetalNodeRequest) String() string { return proto.CompactTextString(m) }
func (*AddBareMetalNodeRequest) ProtoMessage()    {}
func (*AddBareMetalNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_dbb3174ed4d9d994, []int{1}
}
func (m *AddBareMetalNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBareMetalNodeRequest.Unmarshal(m, b)
}
func (m *AddBareMetalNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBareMetalNodeRequest.Marshal(b, m, deterministic)
}
func (dst *AddBareMetalNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBareMetalNodeRequest.Merge(dst, src)
}
func (m *AddBareMetalNodeRequest) XXX_Size() int {
	return xxx_messageInfo_AddBareMetalNodeRequest.Size(m)
}
func (m *AddBareMetalNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBareMetalNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddBareMetalNodeRequest proto.InternalMessageInfo

func (m *AddBareMetalNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddBareMetalNodeRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type AddBareMetalNodeResponse struct {
	Node                 *BareMetalNode `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddBareMetalNodeResponse) Reset()         { *m = AddBareMetalNodeResponse{} }
func (m *AddBareMetalNodeResponse) String() string { return proto.CompactTextString(m) }
func (*AddBareMetalNodeResponse) ProtoMessage()    {}
func (*AddBareMetalNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodes_dbb3174ed4d9d994, []int{2}
}
func (m *AddBareMetalNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddBareMetalNodeResponse.Unmarshal(m, b)
}
func (m *AddBareMetalNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddBareMetalNodeResponse.Marshal(b, m, deterministic)
}
func (dst *AddBareMetalNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddBareMetalNodeResponse.Merge(dst, src)
}
func (m *AddBareMetalNodeResponse) XXX_Size() int {
	return xxx_messageInfo_AddBareMetalNodeResponse.Size(m)
}
func (m *AddBareMetalNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddBareMetalNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddBareMetalNodeResponse proto.InternalMessageInfo

func (m *AddBareMetalNodeResponse) GetNode() *BareMetalNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*BareMetalNode)(nil), "inventory.BareMetalNode")
	proto.RegisterType((*AddBareMetalNodeRequest)(nil), "inventory.AddBareMetalNodeRequest")
	proto.RegisterType((*AddBareMetalNodeResponse)(nil), "inventory.AddBareMetalNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodesClient is the client API for Nodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodesClient interface {
	// Add Bare Metal node.
	AddBareMetal(ctx context.Context, in *AddBareMetalNodeRequest, opts ...grpc.CallOption) (*AddBareMetalNodeResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) AddBareMetal(ctx context.Context, in *AddBareMetalNodeRequest, opts ...grpc.CallOption) (*AddBareMetalNodeResponse, error) {
	out := new(AddBareMetalNodeResponse)
	err := c.cc.Invoke(ctx, "/inventory.Nodes/AddBareMetal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesServer is the server API for Nodes service.
type NodesServer interface {
	// Add Bare Metal node.
	AddBareMetal(context.Context, *AddBareMetalNodeRequest) (*AddBareMetalNodeResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_AddBareMetal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBareMetalNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).AddBareMetal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Nodes/AddBareMetal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).AddBareMetal(ctx, req.(*AddBareMetalNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBareMetal",
			Handler:    _Nodes_AddBareMetal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/nodes.proto",
}

func init() { proto.RegisterFile("inventory/nodes.proto", fileDescriptor_nodes_dbb3174ed4d9d994) }

var fileDescriptor_nodes_dbb3174ed4d9d994 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcc, 0x2b, 0x4b,
	0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0xcf, 0xcb, 0x4f, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x4b, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16,
	0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x41, 0x15, 0x2a, 0xf9,
	0x73, 0xf1, 0x3a, 0x25, 0x16, 0xa5, 0xfa, 0xa6, 0x96, 0x24, 0xe6, 0xf8, 0xe5, 0xa7, 0xa4, 0x0a,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x31, 0x65, 0xa6, 0x08,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x42, 0x52, 0x5c, 0x1c, 0x19, 0xf9, 0xc5, 0x25, 0x60, 0x71, 0x66, 0xb0, 0x38, 0x9c, 0xaf, 0xe4,
	0xc9, 0x25, 0xee, 0x98, 0x92, 0x82, 0x62, 0x66, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09, 0xdc,
	0x28, 0x46, 0x1c, 0x46, 0x31, 0xa1, 0x19, 0xe5, 0xc1, 0x25, 0x81, 0x69, 0x54, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x0e, 0x17, 0x0b, 0xc8, 0xbf, 0x60, 0xb3, 0xb8, 0x8d, 0x24, 0xf4, 0xe0,
	0xfe, 0xd5, 0x43, 0x55, 0x0f, 0x56, 0x65, 0xd4, 0xc7, 0xc8, 0xc5, 0x0a, 0xe2, 0x16, 0x0b, 0xb5,
	0x30, 0x72, 0xf1, 0x20, 0x1b, 0x2a, 0xa4, 0x84, 0xa4, 0x15, 0x87, 0xc3, 0xa5, 0x94, 0xf1, 0xaa,
	0x81, 0xb8, 0x48, 0x49, 0xbb, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xaa, 0x4a, 0x0a, 0xfa, 0x65, 0x06,
	0xfa, 0x88, 0x58, 0x01, 0x5b, 0xab, 0x8f, 0xac, 0xcb, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0x1c, 0xfa,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x62, 0xc4, 0xc7, 0x45, 0xbf, 0x01, 0x00, 0x00,
}