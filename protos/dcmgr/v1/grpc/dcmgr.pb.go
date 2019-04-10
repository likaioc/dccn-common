// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dcmgr/v1/grpc/dcmgr.proto

package dcmgr

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/Ankr-network/dccn-common/protos/common"

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

type DataCenterListResponse struct {
	DcList               []*common.DataCenter `protobuf:"bytes,1,rep,name=dcList,proto3" json:"dcList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DataCenterListResponse) Reset()         { *m = DataCenterListResponse{} }
func (m *DataCenterListResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterListResponse) ProtoMessage()    {}
func (*DataCenterListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_2be9a507cdd3de40, []int{0}
}
func (m *DataCenterListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterListResponse.Unmarshal(m, b)
}
func (m *DataCenterListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterListResponse.Marshal(b, m, deterministic)
}
func (dst *DataCenterListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterListResponse.Merge(dst, src)
}
func (m *DataCenterListResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterListResponse.Size(m)
}
func (m *DataCenterListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterListResponse proto.InternalMessageInfo

func (m *DataCenterListResponse) GetDcList() []*common.DataCenter {
	if m != nil {
		return m.DcList
	}
	return nil
}

type NetworkInfoResponse struct {
	UserCount            int32    `protobuf:"varint,1,opt,name=user_count,json=userCount,proto3" json:"user_count,omitempty"`
	HostCount            int32    `protobuf:"varint,2,opt,name=host_count,json=hostCount,proto3" json:"host_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,3,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	ContainerCount       int32    `protobuf:"varint,4,opt,name=container_count,json=containerCount,proto3" json:"container_count,omitempty"`
	Traffic              int32    `protobuf:"varint,5,opt,name=traffic,proto3" json:"traffic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInfoResponse) Reset()         { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()    {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_2be9a507cdd3de40, []int{1}
}
func (m *NetworkInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInfoResponse.Unmarshal(m, b)
}
func (m *NetworkInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInfoResponse.Marshal(b, m, deterministic)
}
func (dst *NetworkInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInfoResponse.Merge(dst, src)
}
func (m *NetworkInfoResponse) XXX_Size() int {
	return xxx_messageInfo_NetworkInfoResponse.Size(m)
}
func (m *NetworkInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInfoResponse proto.InternalMessageInfo

func (m *NetworkInfoResponse) GetUserCount() int32 {
	if m != nil {
		return m.UserCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetHostCount() int32 {
	if m != nil {
		return m.HostCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetContainerCount() int32 {
	if m != nil {
		return m.ContainerCount
	}
	return 0
}

func (m *NetworkInfoResponse) GetTraffic() int32 {
	if m != nil {
		return m.Traffic
	}
	return 0
}

type DataCenterLeaderBoardResponse struct {
	List                 []*DataCenterLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *DataCenterLeaderBoardResponse) Reset()         { *m = DataCenterLeaderBoardResponse{} }
func (m *DataCenterLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardResponse) ProtoMessage()    {}
func (*DataCenterLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_2be9a507cdd3de40, []int{2}
}
func (m *DataCenterLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (dst *DataCenterLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardResponse.Merge(dst, src)
}
func (m *DataCenterLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardResponse.Size(m)
}
func (m *DataCenterLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardResponse proto.InternalMessageInfo

func (m *DataCenterLeaderBoardResponse) GetList() []*DataCenterLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type DataCenterLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterLeaderBoardDetail) Reset()         { *m = DataCenterLeaderBoardDetail{} }
func (m *DataCenterLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*DataCenterLeaderBoardDetail) ProtoMessage()    {}
func (*DataCenterLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcmgr_2be9a507cdd3de40, []int{3}
}
func (m *DataCenterLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Unmarshal(m, b)
}
func (m *DataCenterLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (dst *DataCenterLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterLeaderBoardDetail.Merge(dst, src)
}
func (m *DataCenterLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_DataCenterLeaderBoardDetail.Size(m)
}
func (m *DataCenterLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterLeaderBoardDetail proto.InternalMessageInfo

func (m *DataCenterLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenterLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*DataCenterListResponse)(nil), "dcmgr.DataCenterListResponse")
	proto.RegisterType((*NetworkInfoResponse)(nil), "dcmgr.NetworkInfoResponse")
	proto.RegisterType((*DataCenterLeaderBoardResponse)(nil), "dcmgr.DataCenterLeaderBoardResponse")
	proto.RegisterType((*DataCenterLeaderBoardDetail)(nil), "dcmgr.DataCenterLeaderBoardDetail")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DCClient is the client API for DC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCClient interface {
	CreateApp(ctx context.Context, in *common.AppDeployment, opts ...grpc.CallOption) (*common.Empty, error)
	UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...grpc.CallOption) (*common.Empty, error)
	PurgeApp(ctx context.Context, in *common.AppID, opts ...grpc.CallOption) (*common.Empty, error)
	Status(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.DataCenterStatus, error)
}

type dCClient struct {
	cc *grpc.ClientConn
}

func NewDCClient(cc *grpc.ClientConn) DCClient {
	return &dCClient{cc}
}

func (c *dCClient) CreateApp(ctx context.Context, in *common.AppDeployment, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dcmgr.DC/CreateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCClient) UpdateApp(ctx context.Context, in *common.AppDeployment, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dcmgr.DC/UpdateApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCClient) PurgeApp(ctx context.Context, in *common.AppID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/dcmgr.DC/PurgeApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCClient) Status(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*common.DataCenterStatus, error) {
	out := new(common.DataCenterStatus)
	err := c.cc.Invoke(ctx, "/dcmgr.DC/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DCServer is the server API for DC service.
type DCServer interface {
	CreateApp(context.Context, *common.AppDeployment) (*common.Empty, error)
	UpdateApp(context.Context, *common.AppDeployment) (*common.Empty, error)
	PurgeApp(context.Context, *common.AppID) (*common.Empty, error)
	Status(context.Context, *common.Empty) (*common.DataCenterStatus, error)
}

func RegisterDCServer(s *grpc.Server, srv DCServer) {
	s.RegisterService(&_DC_serviceDesc, srv)
}

func _DC_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AppDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DC/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCServer).CreateApp(ctx, req.(*common.AppDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _DC_UpdateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AppDeployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCServer).UpdateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DC/UpdateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCServer).UpdateApp(ctx, req.(*common.AppDeployment))
	}
	return interceptor(ctx, in, info, handler)
}

func _DC_PurgeApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.AppID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCServer).PurgeApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DC/PurgeApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCServer).PurgeApp(ctx, req.(*common.AppID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DC_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DC/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCServer).Status(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcmgr.DC",
	HandlerType: (*DCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _DC_CreateApp_Handler,
		},
		{
			MethodName: "UpdateApp",
			Handler:    _DC_UpdateApp_Handler,
		},
		{
			MethodName: "PurgeApp",
			Handler:    _DC_PurgeApp_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DC_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcmgr/v1/grpc/dcmgr.proto",
}

// DCAPIClient is the client API for DCAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DCAPIClient interface {
	DataCenterList(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterListResponse, error)
	DataCenterLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterLeaderBoardResponse, error)
	NetworkInfo(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
}

type dCAPIClient struct {
	cc *grpc.ClientConn
}

func NewDCAPIClient(cc *grpc.ClientConn) DCAPIClient {
	return &dCAPIClient{cc}
}

func (c *dCAPIClient) DataCenterList(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterListResponse, error) {
	out := new(DataCenterListResponse)
	err := c.cc.Invoke(ctx, "/dcmgr.DCAPI/DataCenterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIClient) DataCenterLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*DataCenterLeaderBoardResponse, error) {
	out := new(DataCenterLeaderBoardResponse)
	err := c.cc.Invoke(ctx, "/dcmgr.DCAPI/DataCenterLeaderBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dCAPIClient) NetworkInfo(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := c.cc.Invoke(ctx, "/dcmgr.DCAPI/NetworkInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DCAPIServer is the server API for DCAPI service.
type DCAPIServer interface {
	DataCenterList(context.Context, *common.Empty) (*DataCenterListResponse, error)
	DataCenterLeaderBoard(context.Context, *common.Empty) (*DataCenterLeaderBoardResponse, error)
	NetworkInfo(context.Context, *common.Empty) (*NetworkInfoResponse, error)
}

func RegisterDCAPIServer(s *grpc.Server, srv DCAPIServer) {
	s.RegisterService(&_DCAPI_serviceDesc, srv)
}

func _DCAPI_DataCenterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCAPIServer).DataCenterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DCAPI/DataCenterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCAPIServer).DataCenterList(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DCAPI_DataCenterLeaderBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCAPIServer).DataCenterLeaderBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DCAPI/DataCenterLeaderBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCAPIServer).DataCenterLeaderBoard(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DCAPI_NetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DCAPIServer).NetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dcmgr.DCAPI/NetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DCAPIServer).NetworkInfo(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DCAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dcmgr.DCAPI",
	HandlerType: (*DCAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DataCenterList",
			Handler:    _DCAPI_DataCenterList_Handler,
		},
		{
			MethodName: "DataCenterLeaderBoard",
			Handler:    _DCAPI_DataCenterLeaderBoard_Handler,
		},
		{
			MethodName: "NetworkInfo",
			Handler:    _DCAPI_NetworkInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dcmgr/v1/grpc/dcmgr.proto",
}

func init() { proto.RegisterFile("dcmgr/v1/grpc/dcmgr.proto", fileDescriptor_dcmgr_2be9a507cdd3de40) }

var fileDescriptor_dcmgr_2be9a507cdd3de40 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5f, 0x6b, 0xd4, 0x40,
	0x10, 0xbf, 0xb4, 0x77, 0xa7, 0x37, 0x85, 0xaa, 0x7b, 0x58, 0xce, 0x94, 0x13, 0x59, 0x04, 0x0b,
	0x42, 0xa2, 0x27, 0xf4, 0x55, 0xcf, 0x44, 0x21, 0x22, 0x52, 0xa2, 0xe2, 0xa3, 0x6c, 0x93, 0xb9,
	0x33, 0x78, 0xd9, 0x5d, 0x36, 0x9b, 0x4a, 0x3f, 0x9f, 0x1f, 0xc6, 0x8f, 0xa1, 0xec, 0x6e, 0x9a,
	0xe6, 0x68, 0x7a, 0x2f, 0x3e, 0x25, 0x33, 0xbf, 0x3f, 0xcb, 0xfe, 0x66, 0x16, 0x1e, 0xe5, 0x59,
	0xb9, 0x56, 0xe1, 0xc5, 0xcb, 0x70, 0xad, 0x64, 0x16, 0xda, 0x2a, 0x90, 0x4a, 0x68, 0x41, 0x46,
	0xb6, 0xf0, 0xa7, 0x99, 0x28, 0x4b, 0xc1, 0x43, 0xf7, 0x71, 0x18, 0xfd, 0x00, 0x47, 0x31, 0xd3,
	0x2c, 0x42, 0xae, 0x51, 0x7d, 0x2c, 0x2a, 0x9d, 0x62, 0x25, 0x05, 0xaf, 0x90, 0xbc, 0x80, 0x71,
	0x9e, 0x99, 0xce, 0xcc, 0x7b, 0xb2, 0x7f, 0x72, 0xb0, 0x98, 0x05, 0x5d, 0x61, 0x70, 0xad, 0x4a,
	0x1b, 0x1e, 0xfd, 0xed, 0xc1, 0xf4, 0x13, 0xea, 0x5f, 0x42, 0xfd, 0x4c, 0xf8, 0x4a, 0xb4, 0x4e,
	0x73, 0x80, 0xba, 0x42, 0xf5, 0x3d, 0x13, 0x35, 0x37, 0x6e, 0xde, 0xc9, 0x28, 0x9d, 0x98, 0x4e,
	0x64, 0x1a, 0x06, 0xfe, 0x21, 0x2a, 0xdd, 0xc0, 0x7b, 0x0e, 0x36, 0x1d, 0x07, 0x3f, 0x87, 0x07,
	0xc8, 0x2f, 0x0a, 0x25, 0x78, 0x89, 0xfc, 0x8a, 0xb5, 0x6f, 0x59, 0xf7, 0x3b, 0x80, 0x23, 0x3f,
	0x83, 0x7b, 0x99, 0xe0, 0x9a, 0x15, 0xbc, 0x3d, 0x6f, 0x68, 0xa9, 0x87, 0x6d, 0xdb, 0x11, 0x67,
	0x70, 0x47, 0x2b, 0xb6, 0x5a, 0x15, 0xd9, 0x6c, 0x64, 0x09, 0x57, 0x25, 0xfd, 0x06, 0xf3, 0x4e,
	0x22, 0xc8, 0x72, 0x54, 0x6f, 0x05, 0x53, 0x79, 0x7b, 0x9d, 0x53, 0x18, 0x6e, 0xae, 0x63, 0xa1,
	0x81, 0x8b, 0xba, 0x57, 0x13, 0xa3, 0x66, 0xc5, 0x26, 0xb5, 0x7c, 0x9a, 0xc0, 0xf1, 0x0e, 0x12,
	0x21, 0x30, 0xe4, 0xac, 0x44, 0x9b, 0xcf, 0x24, 0xb5, 0xff, 0xe4, 0x08, 0xc6, 0xbc, 0x2e, 0xcf,
	0x51, 0xd9, 0x58, 0xbc, 0xb4, 0xa9, 0x16, 0x7f, 0x3d, 0xd8, 0x8b, 0x23, 0xf2, 0x1a, 0x26, 0x91,
	0x42, 0xa6, 0x71, 0x29, 0x25, 0x39, 0xde, 0x9e, 0xcf, 0x52, 0xca, 0x18, 0xe5, 0x46, 0x5c, 0x9a,
	0x70, 0xfc, 0xe9, 0x36, 0xf8, 0xae, 0x94, 0xfa, 0x92, 0x0e, 0x8c, 0xc1, 0x57, 0x99, 0xff, 0x87,
	0xc1, 0x29, 0xdc, 0x3d, 0xab, 0xd5, 0xda, 0xea, 0xa7, 0x37, 0xf4, 0x49, 0x7c, 0xfb, 0xc1, 0xe3,
	0xcf, 0x9a, 0xe9, 0xba, 0x22, 0x7d, 0x04, 0xff, 0xf1, 0x6d, 0xbb, 0xe6, 0x44, 0x74, 0xb0, 0xf8,
	0xe3, 0xc1, 0x28, 0x8e, 0x96, 0x67, 0x09, 0x79, 0x0f, 0x87, 0xdb, 0x1b, 0xdc, 0x6f, 0x39, 0xbf,
	0x39, 0xa7, 0xce, 0xb6, 0xd3, 0x01, 0xf9, 0x02, 0x0f, 0x7b, 0xc7, 0xd3, 0x6f, 0xf7, 0x74, 0xd7,
	0xd8, 0x3b, 0xae, 0x6f, 0xe0, 0xa0, 0xf3, 0x24, 0xfa, 0xbd, 0xfc, 0xc6, 0xab, 0xe7, 0xed, 0xd0,
	0xc1, 0xf9, 0xd8, 0x52, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x9f, 0xde, 0x2f, 0xe1,
	0x03, 0x00, 0x00,
}
