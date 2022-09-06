// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/twap/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/osmosis-labs/osmosis/v11/x/twap/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetArithmeticTwapRequest struct {
	PoolId     uint64     `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	BaseAsset  string     `protobuf:"bytes,2,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string     `protobuf:"bytes,3,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
	StartTime  time.Time  `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime    *time.Time `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty" yaml:"end_time"`
}

func (m *GetArithmeticTwapRequest) Reset()         { *m = GetArithmeticTwapRequest{} }
func (m *GetArithmeticTwapRequest) String() string { return proto.CompactTextString(m) }
func (*GetArithmeticTwapRequest) ProtoMessage()    {}
func (*GetArithmeticTwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141a22dba58615af, []int{0}
}
func (m *GetArithmeticTwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArithmeticTwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArithmeticTwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArithmeticTwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArithmeticTwapRequest.Merge(m, src)
}
func (m *GetArithmeticTwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetArithmeticTwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArithmeticTwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArithmeticTwapRequest proto.InternalMessageInfo

func (m *GetArithmeticTwapRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *GetArithmeticTwapRequest) GetBaseAsset() string {
	if m != nil {
		return m.BaseAsset
	}
	return ""
}

func (m *GetArithmeticTwapRequest) GetQuoteAsset() string {
	if m != nil {
		return m.QuoteAsset
	}
	return ""
}

func (m *GetArithmeticTwapRequest) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *GetArithmeticTwapRequest) GetEndTime() *time.Time {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type GetArithmeticTwapResponse struct {
	ArithmeticTwap github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=arithmetic_twap,json=arithmeticTwap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"arithmetic_twap" yaml:"arithmetic_twap"`
}

func (m *GetArithmeticTwapResponse) Reset()         { *m = GetArithmeticTwapResponse{} }
func (m *GetArithmeticTwapResponse) String() string { return proto.CompactTextString(m) }
func (*GetArithmeticTwapResponse) ProtoMessage()    {}
func (*GetArithmeticTwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141a22dba58615af, []int{1}
}
func (m *GetArithmeticTwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArithmeticTwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArithmeticTwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArithmeticTwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArithmeticTwapResponse.Merge(m, src)
}
func (m *GetArithmeticTwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetArithmeticTwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArithmeticTwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArithmeticTwapResponse proto.InternalMessageInfo

type GetArithmeticTwapToNowRequest struct {
	PoolId     uint64    `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	BaseAsset  string    `protobuf:"bytes,2,opt,name=base_asset,json=baseAsset,proto3" json:"base_asset,omitempty"`
	QuoteAsset string    `protobuf:"bytes,3,opt,name=quote_asset,json=quoteAsset,proto3" json:"quote_asset,omitempty"`
	StartTime  time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
}

func (m *GetArithmeticTwapToNowRequest) Reset()         { *m = GetArithmeticTwapToNowRequest{} }
func (m *GetArithmeticTwapToNowRequest) String() string { return proto.CompactTextString(m) }
func (*GetArithmeticTwapToNowRequest) ProtoMessage()    {}
func (*GetArithmeticTwapToNowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_141a22dba58615af, []int{2}
}
func (m *GetArithmeticTwapToNowRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArithmeticTwapToNowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArithmeticTwapToNowRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArithmeticTwapToNowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArithmeticTwapToNowRequest.Merge(m, src)
}
func (m *GetArithmeticTwapToNowRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetArithmeticTwapToNowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArithmeticTwapToNowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArithmeticTwapToNowRequest proto.InternalMessageInfo

func (m *GetArithmeticTwapToNowRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *GetArithmeticTwapToNowRequest) GetBaseAsset() string {
	if m != nil {
		return m.BaseAsset
	}
	return ""
}

func (m *GetArithmeticTwapToNowRequest) GetQuoteAsset() string {
	if m != nil {
		return m.QuoteAsset
	}
	return ""
}

func (m *GetArithmeticTwapToNowRequest) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

type GetArithmeticTwapToNowResponse struct {
	ArithmeticTwap github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=arithmetic_twap,json=arithmeticTwap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"arithmetic_twap" yaml:"arithmetic_twap"`
}

func (m *GetArithmeticTwapToNowResponse) Reset()         { *m = GetArithmeticTwapToNowResponse{} }
func (m *GetArithmeticTwapToNowResponse) String() string { return proto.CompactTextString(m) }
func (*GetArithmeticTwapToNowResponse) ProtoMessage()    {}
func (*GetArithmeticTwapToNowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_141a22dba58615af, []int{3}
}
func (m *GetArithmeticTwapToNowResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetArithmeticTwapToNowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetArithmeticTwapToNowResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetArithmeticTwapToNowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArithmeticTwapToNowResponse.Merge(m, src)
}
func (m *GetArithmeticTwapToNowResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetArithmeticTwapToNowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArithmeticTwapToNowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArithmeticTwapToNowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetArithmeticTwapRequest)(nil), "osmosis.twap.v1beta1.GetArithmeticTwapRequest")
	proto.RegisterType((*GetArithmeticTwapResponse)(nil), "osmosis.twap.v1beta1.GetArithmeticTwapResponse")
	proto.RegisterType((*GetArithmeticTwapToNowRequest)(nil), "osmosis.twap.v1beta1.GetArithmeticTwapToNowRequest")
	proto.RegisterType((*GetArithmeticTwapToNowResponse)(nil), "osmosis.twap.v1beta1.GetArithmeticTwapToNowResponse")
}

func init() { proto.RegisterFile("osmosis/twap/v1beta1/query.proto", fileDescriptor_141a22dba58615af) }

var fileDescriptor_141a22dba58615af = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0x84, 0x3e, 0xc8, 0x54, 0xa2, 0xaa, 0x55, 0x95, 0x34, 0x50, 0x3b, 0xf2, 0xa2, 0x14,
	0x44, 0x67, 0x94, 0xb6, 0x2b, 0xc4, 0xa6, 0x15, 0x12, 0xb0, 0x41, 0xc2, 0xaa, 0x10, 0x62, 0x63,
	0x4d, 0xec, 0xc1, 0xb5, 0x88, 0x3d, 0x8e, 0x67, 0xd2, 0x92, 0x2d, 0x1f, 0x80, 0x2a, 0xf5, 0x0b,
	0x58, 0xf0, 0x11, 0xfc, 0x41, 0x96, 0x15, 0x6c, 0x10, 0x0b, 0x83, 0x12, 0x76, 0xec, 0xf2, 0x05,
	0x68, 0x66, 0xec, 0x10, 0x9a, 0x54, 0x84, 0x15, 0x62, 0x65, 0xcf, 0x3d, 0xe7, 0x9e, 0x7b, 0xee,
	0x9d, 0x07, 0xac, 0x33, 0x1e, 0x31, 0x1e, 0x72, 0x2c, 0x4e, 0x48, 0x82, 0x8f, 0x1b, 0x4d, 0x2a,
	0x48, 0x03, 0xb7, 0x3b, 0x34, 0xed, 0xa2, 0x24, 0x65, 0x82, 0x19, 0xab, 0x39, 0x03, 0x49, 0x06,
	0xca, 0x19, 0xb5, 0xd5, 0x80, 0x05, 0x4c, 0x11, 0xb0, 0xfc, 0xd3, 0xdc, 0xda, 0xe6, 0x54, 0x35,
	0xb9, 0x70, 0x53, 0xea, 0xb1, 0xd4, 0xcf, 0x79, 0xa6, 0xa7, 0x88, 0xb8, 0x49, 0x38, 0x1d, 0xd1,
	0x3c, 0x16, 0xc6, 0x39, 0x7e, 0x67, 0x1c, 0x57, 0x66, 0x46, 0xac, 0x84, 0x04, 0x61, 0x4c, 0x44,
	0xc8, 0x0a, 0xee, 0xcd, 0x80, 0xb1, 0xa0, 0x45, 0x31, 0x49, 0x42, 0x4c, 0xe2, 0x98, 0x09, 0x05,
	0xf2, 0x1c, 0x5d, 0xcf, 0x51, 0xb5, 0x6a, 0x76, 0x5e, 0x62, 0x12, 0x77, 0x0b, 0x48, 0x17, 0x71,
	0x75, 0x17, 0x7a, 0x91, 0x43, 0xd6, 0xc5, 0x2c, 0x11, 0x46, 0x94, 0x0b, 0x12, 0x25, 0x9a, 0x60,
	0xbf, 0x2b, 0xc3, 0xea, 0x43, 0x2a, 0xf6, 0xd3, 0x50, 0x1c, 0x45, 0x54, 0x84, 0xde, 0xe1, 0x09,
	0x49, 0x1c, 0xda, 0xee, 0x50, 0x2e, 0x8c, 0xeb, 0x70, 0x31, 0x61, 0xac, 0xe5, 0x86, 0x7e, 0x15,
	0xd4, 0xc1, 0xd6, 0x9c, 0xb3, 0x20, 0x97, 0x8f, 0x7d, 0x63, 0x03, 0x42, 0xd9, 0x91, 0x4b, 0x38,
	0xa7, 0xa2, 0x5a, 0xae, 0x83, 0xad, 0x8a, 0x53, 0x91, 0x91, 0x7d, 0x19, 0x30, 0x2c, 0xb8, 0xd4,
	0xee, 0x30, 0x51, 0xe0, 0x57, 0x14, 0x0e, 0x55, 0x48, 0x13, 0x9e, 0x43, 0xc8, 0x05, 0x49, 0x85,
	0x2b, 0xed, 0x54, 0xe7, 0xea, 0x60, 0x6b, 0x69, 0xa7, 0x86, 0xb4, 0x57, 0x54, 0x78, 0x45, 0x87,
	0x85, 0xd7, 0x83, 0x8d, 0x5e, 0x66, 0x95, 0x86, 0x99, 0xb5, 0xd2, 0x25, 0x51, 0xeb, 0x9e, 0xfd,
	0x2b, 0xd7, 0x3e, 0xfd, 0x6a, 0x01, 0xa7, 0xa2, 0x02, 0x92, 0x6e, 0x38, 0xf0, 0x2a, 0x8d, 0x7d,
	0xad, 0x3b, 0xff, 0x47, 0xdd, 0x1b, 0xbd, 0xcc, 0x02, 0xc3, 0xcc, 0x5a, 0xd6, 0xba, 0x45, 0xa6,
	0x56, 0x5d, 0xa4, 0xb1, 0x2f, 0xa9, 0xf6, 0x5b, 0x00, 0xd7, 0xa7, 0xcc, 0x88, 0x27, 0x2c, 0xe6,
	0xd4, 0x68, 0xc3, 0x65, 0x32, 0x42, 0x5c, 0x79, 0x44, 0xd4, 0xb0, 0x2a, 0x07, 0x8f, 0xa4, 0xe9,
	0x2f, 0x99, 0xb5, 0x19, 0x84, 0xe2, 0xa8, 0xd3, 0x44, 0x1e, 0x8b, 0xf2, 0xcd, 0xc9, 0x3f, 0xdb,
	0xdc, 0x7f, 0x85, 0x45, 0x37, 0xa1, 0x1c, 0x3d, 0xa0, 0xde, 0x30, 0xb3, 0xd6, 0xb4, 0x8d, 0x0b,
	0x72, 0xb6, 0x73, 0x8d, 0xfc, 0x56, 0xda, 0xfe, 0x08, 0xe0, 0xc6, 0x84, 0xa1, 0x43, 0xf6, 0x84,
	0x9d, 0xfc, 0xbf, 0x3b, 0x67, 0x9f, 0x01, 0x68, 0x5e, 0xd6, 0xd4, 0x3f, 0x1b, 0xf5, 0xce, 0x8f,
	0x32, 0x9c, 0x7f, 0x2a, 0xef, 0xad, 0xf1, 0x1e, 0xc0, 0x95, 0x09, 0x7f, 0x06, 0x42, 0xd3, 0x5e,
	0x15, 0x74, 0xd9, 0x95, 0xaa, 0xe1, 0x99, 0xf9, 0xba, 0x67, 0x1b, 0xbf, 0xf9, 0xf4, 0xfd, 0xac,
	0x7c, 0xdb, 0xb8, 0x85, 0xa7, 0x3e, 0x49, 0x93, 0x8e, 0x3e, 0x00, 0xb8, 0x36, 0x7d, 0x8e, 0xc6,
	0xee, 0x8c, 0xc5, 0xc7, 0x8f, 0x52, 0x6d, 0xef, 0xef, 0x92, 0x72, 0xdb, 0x7b, 0xca, 0x36, 0x32,
	0xee, 0xce, 0x68, 0x5b, 0x65, 0x1f, 0x3c, 0xeb, 0xf5, 0x4d, 0x70, 0xde, 0x37, 0xc1, 0xb7, 0xbe,
	0x09, 0x4e, 0x07, 0x66, 0xe9, 0x7c, 0x60, 0x96, 0x3e, 0x0f, 0xcc, 0xd2, 0x8b, 0xfb, 0x63, 0x3b,
	0x9b, 0x2b, 0x6e, 0xb7, 0x48, 0x93, 0x8f, 0xe4, 0x8f, 0x1b, 0x0d, 0xfc, 0x5a, 0x17, 0xf1, 0x5a,
	0x21, 0x8d, 0x85, 0x7e, 0x6e, 0xf5, 0x91, 0x5c, 0x50, 0x9f, 0xdd, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x6a, 0x07, 0x65, 0x9a, 0x25, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	GetArithmeticTwap(ctx context.Context, in *GetArithmeticTwapRequest, opts ...grpc.CallOption) (*GetArithmeticTwapResponse, error)
	GetArithmeticTwapToNow(ctx context.Context, in *GetArithmeticTwapToNowRequest, opts ...grpc.CallOption) (*GetArithmeticTwapToNowResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetArithmeticTwap(ctx context.Context, in *GetArithmeticTwapRequest, opts ...grpc.CallOption) (*GetArithmeticTwapResponse, error) {
	out := new(GetArithmeticTwapResponse)
	err := c.cc.Invoke(ctx, "/osmosis.twap.v1beta1.Query/GetArithmeticTwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetArithmeticTwapToNow(ctx context.Context, in *GetArithmeticTwapToNowRequest, opts ...grpc.CallOption) (*GetArithmeticTwapToNowResponse, error) {
	out := new(GetArithmeticTwapToNowResponse)
	err := c.cc.Invoke(ctx, "/osmosis.twap.v1beta1.Query/GetArithmeticTwapToNow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetArithmeticTwap(context.Context, *GetArithmeticTwapRequest) (*GetArithmeticTwapResponse, error)
	GetArithmeticTwapToNow(context.Context, *GetArithmeticTwapToNowRequest) (*GetArithmeticTwapToNowResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetArithmeticTwap(ctx context.Context, req *GetArithmeticTwapRequest) (*GetArithmeticTwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArithmeticTwap not implemented")
}
func (*UnimplementedQueryServer) GetArithmeticTwapToNow(ctx context.Context, req *GetArithmeticTwapToNowRequest) (*GetArithmeticTwapToNowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArithmeticTwapToNow not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetArithmeticTwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArithmeticTwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetArithmeticTwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.twap.v1beta1.Query/GetArithmeticTwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetArithmeticTwap(ctx, req.(*GetArithmeticTwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetArithmeticTwapToNow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArithmeticTwapToNowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetArithmeticTwapToNow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.twap.v1beta1.Query/GetArithmeticTwapToNow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetArithmeticTwapToNow(ctx, req.(*GetArithmeticTwapToNowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.twap.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArithmeticTwap",
			Handler:    _Query_GetArithmeticTwap_Handler,
		},
		{
			MethodName: "GetArithmeticTwapToNow",
			Handler:    _Query_GetArithmeticTwapToNow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/twap/v1beta1/query.proto",
}

func (m *GetArithmeticTwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArithmeticTwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArithmeticTwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintQuery(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x2a
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintQuery(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if len(m.QuoteAsset) > 0 {
		i -= len(m.QuoteAsset)
		copy(dAtA[i:], m.QuoteAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QuoteAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAsset) > 0 {
		i -= len(m.BaseAsset)
		copy(dAtA[i:], m.BaseAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BaseAsset)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetArithmeticTwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArithmeticTwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArithmeticTwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ArithmeticTwap.Size()
		i -= size
		if _, err := m.ArithmeticTwap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GetArithmeticTwapToNowRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArithmeticTwapToNowRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArithmeticTwapToNowRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintQuery(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	if len(m.QuoteAsset) > 0 {
		i -= len(m.QuoteAsset)
		copy(dAtA[i:], m.QuoteAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.QuoteAsset)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BaseAsset) > 0 {
		i -= len(m.BaseAsset)
		copy(dAtA[i:], m.BaseAsset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BaseAsset)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetArithmeticTwapToNowResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetArithmeticTwapToNowResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetArithmeticTwapToNowResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ArithmeticTwap.Size()
		i -= size
		if _, err := m.ArithmeticTwap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetArithmeticTwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovQuery(uint64(m.PoolId))
	}
	l = len(m.BaseAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.QuoteAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovQuery(uint64(l))
	if m.EndTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetArithmeticTwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ArithmeticTwap.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *GetArithmeticTwapToNowRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovQuery(uint64(m.PoolId))
	}
	l = len(m.BaseAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.QuoteAsset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *GetArithmeticTwapToNowResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ArithmeticTwap.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetArithmeticTwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetArithmeticTwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArithmeticTwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetArithmeticTwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetArithmeticTwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArithmeticTwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArithmeticTwap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ArithmeticTwap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetArithmeticTwapToNowRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetArithmeticTwapToNowRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArithmeticTwapToNowRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuoteAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuoteAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetArithmeticTwapToNowResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetArithmeticTwapToNowResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetArithmeticTwapToNowResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArithmeticTwap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ArithmeticTwap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
