// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mr_gun.proto

package mr_gun

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

type Request struct {
	Sort                 string   `protobuf:"bytes,1,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdc78a6994181c46, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

type GunInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Damage               string   `protobuf:"bytes,2,opt,name=damage,proto3" json:"damage,omitempty"`
	BulletsShot          string   `protobuf:"bytes,3,opt,name=bulletsShot,proto3" json:"bulletsShot,omitempty"`
	TotalDamage          string   `protobuf:"bytes,4,opt,name=totalDamage,proto3" json:"totalDamage,omitempty"`
	Scatter              string   `protobuf:"bytes,5,opt,name=scatter,proto3" json:"scatter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GunInfo) Reset()         { *m = GunInfo{} }
func (m *GunInfo) String() string { return proto.CompactTextString(m) }
func (*GunInfo) ProtoMessage()    {}
func (*GunInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdc78a6994181c46, []int{1}
}

func (m *GunInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GunInfo.Unmarshal(m, b)
}
func (m *GunInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GunInfo.Marshal(b, m, deterministic)
}
func (m *GunInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GunInfo.Merge(m, src)
}
func (m *GunInfo) XXX_Size() int {
	return xxx_messageInfo_GunInfo.Size(m)
}
func (m *GunInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GunInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GunInfo proto.InternalMessageInfo

func (m *GunInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GunInfo) GetDamage() string {
	if m != nil {
		return m.Damage
	}
	return ""
}

func (m *GunInfo) GetBulletsShot() string {
	if m != nil {
		return m.BulletsShot
	}
	return ""
}

func (m *GunInfo) GetTotalDamage() string {
	if m != nil {
		return m.TotalDamage
	}
	return ""
}

func (m *GunInfo) GetScatter() string {
	if m != nil {
		return m.Scatter
	}
	return ""
}

type Reply struct {
	GunInfo              []*GunInfo `protobuf:"bytes,1,rep,name=gunInfo,proto3" json:"gunInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdc78a6994181c46, []int{2}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetGunInfo() []*GunInfo {
	if m != nil {
		return m.GunInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "mr_gun.Request")
	proto.RegisterType((*GunInfo)(nil), "mr_gun.GunInfo")
	proto.RegisterType((*Reply)(nil), "mr_gun.Reply")
}

func init() { proto.RegisterFile("mr_gun.proto", fileDescriptor_fdc78a6994181c46) }

var fileDescriptor_fdc78a6994181c46 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4a, 0xc4, 0x40,
	0x10, 0x85, 0x89, 0x33, 0x49, 0xb0, 0x46, 0x11, 0x6a, 0x21, 0x8d, 0x20, 0x84, 0xac, 0xc6, 0xcd,
	0x2c, 0xe2, 0x15, 0x84, 0xe0, 0xc2, 0x4d, 0x7b, 0x00, 0xe9, 0xd1, 0x32, 0x82, 0x9d, 0xee, 0xd8,
	0xa9, 0x5e, 0xcc, 0x39, 0xbc, 0xb0, 0xf4, 0x4f, 0x9c, 0xec, 0xea, 0xbd, 0xef, 0x15, 0x54, 0x3d,
	0xb8, 0x1a, 0xdd, 0xdb, 0xe0, 0xcd, 0x61, 0x72, 0x96, 0x2d, 0x56, 0x49, 0xb5, 0xf7, 0x50, 0x4b,
	0xfa, 0xf1, 0x34, 0x33, 0x22, 0x6c, 0x67, 0xeb, 0x58, 0x14, 0x4d, 0xb1, 0xbf, 0x94, 0x71, 0x6e,
	0x7f, 0x0b, 0xa8, 0x7b, 0x6f, 0x9e, 0xcd, 0xa7, 0x0d, 0xdc, 0xa8, 0x91, 0x16, 0x1e, 0x66, 0xbc,
	0x85, 0xea, 0x43, 0x8d, 0x6a, 0x20, 0x71, 0x11, 0xdd, 0xac, 0xb0, 0x81, 0xdd, 0xd1, 0x6b, 0x4d,
	0x3c, 0xbf, 0x7e, 0x59, 0x16, 0x9b, 0x08, 0xd7, 0x56, 0x48, 0xb0, 0x65, 0xa5, 0x9f, 0xd2, 0xfa,
	0x36, 0x25, 0x56, 0x16, 0x0a, 0xa8, 0xe7, 0x77, 0xc5, 0x4c, 0x4e, 0x94, 0x91, 0x2e, 0xb2, 0xed,
	0xa0, 0x94, 0x34, 0xe9, 0x13, 0x3e, 0x40, 0x3d, 0xa4, 0xeb, 0x44, 0xd1, 0x6c, 0xf6, 0xbb, 0xee,
	0xe6, 0x90, 0xbf, 0xcc, 0x47, 0xcb, 0x85, 0x77, 0x1d, 0x94, 0x2f, 0xae, 0xf7, 0x26, 0xec, 0xf4,
	0xc4, 0x52, 0x99, 0x6f, 0xfc, 0x4f, 0xe7, 0x0a, 0xee, 0xae, 0xcf, 0xc6, 0xa4, 0x4f, 0xc7, 0x2a,
	0x76, 0xf5, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x64, 0x6c, 0xd3, 0xb5, 0x3b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MrGunClient is the client API for MrGun service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MrGunClient interface {
	GetRank(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type mrGunClient struct {
	cc *grpc.ClientConn
}

func NewMrGunClient(cc *grpc.ClientConn) MrGunClient {
	return &mrGunClient{cc}
}

func (c *mrGunClient) GetRank(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/mr_gun.MrGun/GetRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MrGunServer is the server API for MrGun service.
type MrGunServer interface {
	GetRank(context.Context, *Request) (*Reply, error)
}

func RegisterMrGunServer(s *grpc.Server, srv MrGunServer) {
	s.RegisterService(&_MrGun_serviceDesc, srv)
}

func _MrGun_GetRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MrGunServer).GetRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mr_gun.MrGun/GetRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MrGunServer).GetRank(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _MrGun_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mr_gun.MrGun",
	HandlerType: (*MrGunServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRank",
			Handler:    _MrGun_GetRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mr_gun.proto",
}
