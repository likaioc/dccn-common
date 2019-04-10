// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package common_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Hub app status
type AppStatus int32

const (
	AppStatus_APP_STARTING      AppStatus = 0
	AppStatus_APP_START_FAILED  AppStatus = 1
	AppStatus_APP_RUNNING       AppStatus = 2
	AppStatus_APP_UPDATING      AppStatus = 3
	AppStatus_APP_UPDATE_FAILED AppStatus = 4
	AppStatus_APP_CANCELLING    AppStatus = 5
	AppStatus_APP_CANCELLED     AppStatus = 6
	AppStatus_APP_CANCEL_FAILED AppStatus = 7
)

var AppStatus_name = map[int32]string{
	0: "APP_STARTING",
	1: "APP_START_FAILED",
	2: "APP_RUNNING",
	3: "APP_UPDATING",
	4: "APP_UPDATE_FAILED",
	5: "APP_CANCELLING",
	6: "APP_CANCELLED",
	7: "APP_CANCEL_FAILED",
}

var AppStatus_value = map[string]int32{
	"APP_STARTING":      0,
	"APP_START_FAILED":  1,
	"APP_RUNNING":       2,
	"APP_UPDATING":      3,
	"APP_UPDATE_FAILED": 4,
	"APP_CANCELLING":    5,
	"APP_CANCELLED":     6,
	"APP_CANCEL_FAILED": 7,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_STARTING      NamespaceStatus = 0
	NamespaceStatus_NS_START_FAILED  NamespaceStatus = 1
	NamespaceStatus_NS_RUNNING       NamespaceStatus = 2
	NamespaceStatus_NS_UPDATING      NamespaceStatus = 3
	NamespaceStatus_NS_UPDATE_FAILED NamespaceStatus = 4
	NamespaceStatus_NS_CANCELLING    NamespaceStatus = 5
	NamespaceStatus_NS_CANCELLED     NamespaceStatus = 6
	NamespaceStatus_NS_CANCEL_FAILED NamespaceStatus = 7
)

var NamespaceStatus_name = map[int32]string{
	0: "NS_STARTING",
	1: "NS_START_FAILED",
	2: "NS_RUNNING",
	3: "NS_UPDATING",
	4: "NS_UPDATE_FAILED",
	5: "NS_CANCELLING",
	6: "NS_CANCELLED",
	7: "NS_CANCEL_FAILED",
}

var NamespaceStatus_value = map[string]int32{
	"NS_STARTING":      0,
	"NS_START_FAILED":  1,
	"NS_RUNNING":       2,
	"NS_UPDATING":      3,
	"NS_UPDATE_FAILED": 4,
	"NS_CANCELLING":    5,
	"NS_CANCELLED":     6,
	"NS_CANCEL_FAILED": 7,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func init() {
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x46, 0x09, 0x3f, 0x45, 0x7c, 0x49, 0x1b, 0x67, 0x28, 0x97, 0xe8, 0x82, 0x0d, 0x27, 0xb0,
	0x1a, 0x83, 0x2a, 0x45, 0xa3, 0xa8, 0x6e, 0xd7, 0x55, 0x88, 0xba, 0x4c, 0x6d, 0x11, 0x73, 0x28,
	0xc4, 0x25, 0x91, 0x63, 0xdc, 0x1a, 0x96, 0xf3, 0xf4, 0x3d, 0xe9, 0x69, 0x90, 0xf7, 0xe6, 0x34,
	0xba, 0x67, 0xfb, 0x61, 0x9c, 0xa1, 0xa2, 0x37, 0xc3, 0x60, 0x4e, 0xe1, 0x5a, 0x7d, 0x67, 0x78,
	0x90, 0xd6, 0x6a, 0xd7, 0xb9, 0xcf, 0x91, 0x04, 0x0a, 0xd9, 0xb6, 0x07, 0xbd, 0x93, 0xdb, 0xdd,
	0x86, 0xdf, 0xc4, 0x15, 0x2d, 0x21, 0xce, 0xe4, 0xf0, 0x2a, 0x37, 0x8d, 0xaa, 0x45, 0x46, 0x25,
	0x72, 0x4f, 0xb7, 0x7b, 0x66, 0x3f, 0xbb, 0x8e, 0xe2, 0xbe, 0xad, 0xe5, 0x24, 0xde, 0xd0, 0x13,
	0xaa, 0x33, 0x51, 0xd1, 0xbc, 0x25, 0xc2, 0xc2, 0xe3, 0xb5, 0xe4, 0xb5, 0x6a, 0x1a, 0x3f, 0xbd,
	0xa3, 0x0a, 0xf3, 0x84, 0xa9, 0x5a, 0xcc, 0xa2, 0x1d, 0x50, 0xb4, 0xef, 0x57, 0x5f, 0x19, 0x4a,
	0xee, 0x86, 0xe3, 0x68, 0xbb, 0xfe, 0xf8, 0xdb, 0x5c, 0x22, 0x67, 0x9d, 0x26, 0x3f, 0xa2, 0x8c,
	0xe0, 0x52, 0xbc, 0x00, 0x58, 0x27, 0xc1, 0xc1, 0x4a, 0x7a, 0x97, 0x10, 0x11, 0x24, 0xb9, 0x15,
	0xe6, 0xac, 0xff, 0xd6, 0x0a, 0x14, 0x17, 0x34, 0xc5, 0x06, 0xf5, 0x5f, 0xeb, 0xfb, 0x6c, 0x7a,
	0xf0, 0xcb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x8a, 0xf3, 0x25, 0x7d, 0x01, 0x00, 0x00,
}
