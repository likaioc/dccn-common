// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string     `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Image                string     `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Replica              int32      `protobuf:"varint,6,opt,name=replica,proto3" json:"replica,omitempty"`
	DataCenter           string     `protobuf:"bytes,7,opt,name=data_center,json=dataCenter,proto3" json:"data_center,omitempty"`
	DataCenterId         string     `protobuf:"bytes,8,opt,name=data_center_id,json=dataCenterId,proto3" json:"data_center_id,omitempty"`
	Status               TaskStatus `protobuf:"varint,9,opt,name=status,proto3,enum=common.proto.TaskStatus" json:"status,omitempty"`
	UniqueName           string     `protobuf:"bytes,10,opt,name=unique_name,json=uniqueName,proto3" json:"unique_name,omitempty"`
	Url                  string     `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Hidden               bool       `protobuf:"varint,12,opt,name=hidden,proto3" json:"hidden,omitempty"`
	Uptime               uint32     `protobuf:"varint,13,opt,name=uptime,proto3" json:"uptime,omitempty"`
	CreationDate         uint64     `protobuf:"varint,14,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Report               string     `protobuf:"bytes,15,opt,name=report,proto3" json:"report,omitempty"`
	Extra                []byte     `protobuf:"bytes,16,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_f5e74526e14bc6ce, []int{0}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Task) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Task) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *Task) GetDataCenter() string {
	if m != nil {
		return m.DataCenter
	}
	return ""
}

func (m *Task) GetDataCenterId() string {
	if m != nil {
		return m.DataCenterId
	}
	return ""
}

func (m *Task) GetStatus() TaskStatus {
	if m != nil {
		return m.Status
	}
	return TaskStatus_START
}

func (m *Task) GetUniqueName() string {
	if m != nil {
		return m.UniqueName
	}
	return ""
}

func (m *Task) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Task) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Task) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *Task) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Task) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *Task) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type DataCenter struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lat                  string   `protobuf:"bytes,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,4,opt,name=lng,proto3" json:"lng,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Metrics              string   `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Extra                string   `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	Report               string   `protobuf:"bytes,8,opt,name=report,proto3" json:"report,omitempty"`
	LastReportTime       string   `protobuf:"bytes,9,opt,name=last_report_time,json=lastReportTime,proto3" json:"last_report_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_f5e74526e14bc6ce, []int{1}
}
func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (dst *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(dst, src)
}
func (m *DataCenter) XXX_Size() int {
	return xxx_messageInfo_DataCenter.Size(m)
}
func (m *DataCenter) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenter.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenter proto.InternalMessageInfo

func (m *DataCenter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataCenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataCenter) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *DataCenter) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *DataCenter) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DataCenter) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DataCenter) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *DataCenter) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DataCenter) GetLastReportTime() string {
	if m != nil {
		return m.LastReportTime
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "common.proto.Task")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_f5e74526e14bc6ce) }

var fileDescriptor_common_f5e74526e14bc6ce = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xe5, 0xfe, 0xa4, 0xed, 0x9d, 0x34, 0x54, 0x16, 0x82, 0x2b, 0x36, 0x44, 0x03, 0x0b,
	0xaf, 0x2a, 0x04, 0x8f, 0xc0, 0x6c, 0x66, 0xc3, 0xc2, 0xcc, 0x3e, 0x32, 0xf1, 0xd5, 0x60, 0xd1,
	0xfc, 0xe0, 0x38, 0x12, 0xbc, 0x18, 0x8f, 0xc3, 0xb3, 0x20, 0x5f, 0x27, 0x6a, 0xca, 0xee, 0x9e,
	0xcf, 0xd7, 0xb2, 0xcf, 0x39, 0x90, 0xd7, 0x5d, 0xd3, 0x74, 0xed, 0xb9, 0xf7, 0x5d, 0xe8, 0xe4,
	0x8d, 0x7a, 0x03, 0x75, 0x67, 0x29, 0xcd, 0xf7, 0x7f, 0xd6, 0xb0, 0x79, 0x32, 0xc3, 0x0f, 0x59,
	0xc0, 0xca, 0x59, 0x14, 0xa5, 0x50, 0x07, 0xbd, 0x72, 0x56, 0xbe, 0x86, 0xdd, 0x38, 0x90, 0xaf,
	0x9c, 0xc5, 0x15, 0xc3, 0x2c, 0xca, 0x47, 0x2b, 0x25, 0x6c, 0x5a, 0xd3, 0x10, 0xae, 0x99, 0xf2,
	0x1c, 0x59, 0xf8, 0xdd, 0x13, 0x6e, 0x12, 0x8b, 0xb3, 0x7c, 0x09, 0x5b, 0xd7, 0x98, 0x67, 0xc2,
	0x2d, 0xc3, 0x24, 0x24, 0xc2, 0xce, 0x53, 0x7f, 0x71, 0xb5, 0xc1, 0xac, 0x14, 0x6a, 0xab, 0x67,
	0x29, 0xdf, 0xc2, 0x9d, 0x35, 0xc1, 0x54, 0x35, 0xb5, 0x81, 0x3c, 0xee, 0xf8, 0x16, 0x44, 0xf4,
	0x99, 0x89, 0x7c, 0x0f, 0xc5, 0x62, 0x21, 0x7e, 0x6c, 0xcf, 0x3b, 0xf9, 0x75, 0xe7, 0xd1, 0xca,
	0x0f, 0x90, 0x0d, 0xc1, 0x84, 0x71, 0xc0, 0x43, 0x29, 0x54, 0xf1, 0x11, 0xcf, 0x4b, 0xef, 0xe7,
	0xe8, 0xf5, 0x2b, 0x9f, 0xeb, 0x69, 0x2f, 0x3e, 0x3c, 0xb6, 0xee, 0xe7, 0x48, 0x15, 0xfb, 0x82,
	0xf4, 0x70, 0x42, 0x5f, 0xa2, 0xbb, 0x13, 0xac, 0x47, 0x7f, 0xc1, 0x3b, 0x3e, 0x88, 0xa3, 0x7c,
	0x05, 0xd9, 0x77, 0x67, 0x2d, 0xb5, 0x98, 0x97, 0x42, 0xed, 0xf5, 0xa4, 0x22, 0x1f, 0xfb, 0xe0,
	0x1a, 0xc2, 0x63, 0x29, 0xd4, 0x51, 0x4f, 0x4a, 0xbe, 0x83, 0x63, 0xed, 0xc9, 0x04, 0xd7, 0xb5,
	0x95, 0x35, 0x81, 0xb0, 0x28, 0x85, 0xda, 0xe8, 0x7c, 0x86, 0x0f, 0x26, 0x50, 0xbc, 0xec, 0xa9,
	0xef, 0x7c, 0xc0, 0x17, 0x29, 0xf0, 0xa4, 0x62, 0x90, 0xf4, 0x2b, 0x78, 0x83, 0xa7, 0x52, 0xa8,
	0x5c, 0x27, 0x71, 0xff, 0x57, 0x00, 0x3c, 0x5c, 0xc3, 0xf9, 0xbf, 0xbe, 0xb9, 0xa5, 0xd5, 0xa2,
	0xa5, 0x13, 0xac, 0x2f, 0x26, 0x4c, 0xc5, 0xc5, 0x91, 0x49, 0xfb, 0x3c, 0xd5, 0x16, 0xc7, 0xf8,
	0x89, 0x29, 0xbe, 0x54, 0xdb, 0x1c, 0x12, 0xc2, 0xae, 0xa1, 0xe0, 0x5d, 0x3d, 0x70, 0x6f, 0x07,
	0x3d, 0xcb, 0xeb, 0xf7, 0x52, 0x63, 0x49, 0x2c, 0xcc, 0xec, 0x6f, 0xcc, 0x28, 0x38, 0x5d, 0xcc,
	0x10, 0xaa, 0x24, 0x2b, 0xce, 0xea, 0xc0, 0x1b, 0x45, 0xe4, 0x9a, 0xf1, 0x93, 0x6b, 0xe8, 0x5b,
	0xc6, 0x85, 0x7d, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x9e, 0xa7, 0xd2, 0xca, 0x02, 0x00,
	0x00,
}
