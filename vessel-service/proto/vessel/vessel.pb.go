// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Spefication struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spefication) Reset()         { *m = Spefication{} }
func (m *Spefication) String() string { return proto.CompactTextString(m) }
func (*Spefication) ProtoMessage()    {}
func (*Spefication) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{1}
}

func (m *Spefication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spefication.Unmarshal(m, b)
}
func (m *Spefication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spefication.Marshal(b, m, deterministic)
}
func (m *Spefication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spefication.Merge(m, src)
}
func (m *Spefication) XXX_Size() int {
	return xxx_messageInfo_Spefication.Size(m)
}
func (m *Spefication) XXX_DiscardUnknown() {
	xxx_messageInfo_Spefication.DiscardUnknown(m)
}

var xxx_messageInfo_Spefication proto.InternalMessageInfo

func (m *Spefication) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Spefication) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ef66862bb50716, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Spefication)(nil), "vessel.Spefication")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_04ef66862bb50716) }

var fileDescriptor_04ef66862bb50716 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x4a, 0xc3, 0x50,
	0x10, 0xf5, 0xa6, 0x6d, 0x9a, 0x4c, 0x69, 0x91, 0x71, 0x73, 0x5b, 0x14, 0x42, 0x16, 0x92, 0x55,
	0x85, 0xba, 0x71, 0xeb, 0x46, 0xec, 0x36, 0x05, 0xdd, 0x08, 0xe5, 0x36, 0x19, 0x75, 0x20, 0x2f,
	0x92, 0x90, 0xd6, 0xbf, 0xf1, 0x53, 0x85, 0x49, 0x52, 0x1f, 0x48, 0x57, 0x39, 0x8f, 0xcc, 0x70,
	0xe6, 0x5c, 0x98, 0x17, 0x65, 0x5e, 0xe7, 0x37, 0x0d, 0x55, 0x15, 0x25, 0xdd, 0x67, 0x29, 0x1a,
	0xda, 0x2d, 0xf3, 0x3f, 0x15, 0xd8, 0x4f, 0x02, 0x71, 0x06, 0x16, 0xc7, 0x5a, 0x79, 0x2a, 0x70,
	0x43, 0x8b, 0x63, 0x5c, 0x80, 0x13, 0x99, 0xc2, 0x44, 0x5c, 0x7f, 0x68, 0xcb, 0x53, 0xc1, 0x28,
	0x3c, 0x72, 0xbc, 0x02, 0x48, 0xcd, 0x61, 0xbb, 0x27, 0x7e, 0x7b, 0xaf, 0xf5, 0x40, 0x5c, 0x37,
	0x35, 0x87, 0x67, 0x11, 0x10, 0x61, 0x98, 0x99, 0x94, 0xf4, 0x50, 0x96, 0x09, 0xc6, 0x4b, 0x70,
	0x4d, 0x63, 0x38, 0x31, 0xbb, 0x84, 0xf4, 0xc8, 0x53, 0x81, 0x13, 0x7e, 0x0b, 0x38, 0x07, 0x27,
	0xdf, 0x67, 0x54, 0x6e, 0x39, 0xd6, 0xb6, 0x4c, 0x8d, 0x85, 0xaf, 0x63, 0xff, 0x11, 0x26, 0x9b,
	0x82, 0x5e, 0x39, 0x32, 0x35, 0xe7, 0xd9, 0xaf, 0x58, 0xea, 0x64, 0x2c, 0xeb, 0x4f, 0x2c, 0xff,
	0x05, 0x9c, 0x90, 0xaa, 0x22, 0xcf, 0x2a, 0xc2, 0x6b, 0xe8, 0x2a, 0x90, 0x25, 0x93, 0xd5, 0x6c,
	0xd9, 0xf5, 0xd3, 0xb6, 0x11, 0x76, 0x2e, 0x06, 0x30, 0x6e, 0x51, 0xa5, 0x2d, 0x6f, 0xf0, 0xcf,
	0x8f, 0xbd, 0xbd, 0x5a, 0xc3, 0xb4, 0x95, 0x36, 0x54, 0x36, 0x1c, 0x11, 0xde, 0xc1, 0xf4, 0x81,
	0xb3, 0xf8, 0xfe, 0x78, 0xe4, 0x45, 0x3f, 0xfa, 0xe3, 0x9e, 0xc5, 0x79, 0x2f, 0xf6, 0xd1, 0xfc,
	0xb3, 0x9d, 0x2d, 0x8f, 0x74, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x92, 0x93, 0x04, 0x30, 0xc1,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Spefication, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Spefication, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Spefication, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Spefication, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}