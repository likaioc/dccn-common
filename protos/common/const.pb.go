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

// Hub application status
type AppStatus int32

const (
	AppStatus_APP_FAILED        AppStatus = 0
	AppStatus_APP_DISPATCHING   AppStatus = 1
	AppStatus_APP_LAUNCHING     AppStatus = 2
	AppStatus_APP_RUNNING       AppStatus = 3
	AppStatus_APP_UPDATING      AppStatus = 4
	AppStatus_APP_UPDATE_FAILED AppStatus = 5
	AppStatus_APP_CANCELING     AppStatus = 6
	AppStatus_APP_CANCELED      AppStatus = 7
)

var AppStatus_name = map[int32]string{
	0: "APP_FAILED",
	1: "APP_DISPATCHING",
	2: "APP_LAUNCHING",
	3: "APP_RUNNING",
	4: "APP_UPDATING",
	5: "APP_UPDATE_FAILED",
	6: "APP_CANCELING",
	7: "APP_CANCELED",
}

var AppStatus_value = map[string]int32{
	"APP_FAILED":        0,
	"APP_DISPATCHING":   1,
	"APP_LAUNCHING":     2,
	"APP_RUNNING":       3,
	"APP_UPDATING":      4,
	"APP_UPDATE_FAILED": 5,
	"APP_CANCELING":     6,
	"APP_CANCELED":      7,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

// Hub application event
type AppEvent int32

const (
	AppEvent_DISPATCH_APP       AppEvent = 0
	AppEvent_LAUNCH_APP         AppEvent = 1
	AppEvent_LAUNCH_APP_FAILED  AppEvent = 2
	AppEvent_LAUNCH_APP_SUCCEED AppEvent = 3
	AppEvent_UPDATE_APP         AppEvent = 4
	AppEvent_UPDATE_APP_FAILED  AppEvent = 5
	AppEvent_UPDATE_APP_SUCCEED AppEvent = 6
	AppEvent_CANCEL_APP         AppEvent = 7
	AppEvent_CANCEL_APP_FAILED  AppEvent = 8
	AppEvent_CANCEL_APP_SUCCEED AppEvent = 9
)

var AppEvent_name = map[int32]string{
	0: "DISPATCH_APP",
	1: "LAUNCH_APP",
	2: "LAUNCH_APP_FAILED",
	3: "LAUNCH_APP_SUCCEED",
	4: "UPDATE_APP",
	5: "UPDATE_APP_FAILED",
	6: "UPDATE_APP_SUCCEED",
	7: "CANCEL_APP",
	8: "CANCEL_APP_FAILED",
	9: "CANCEL_APP_SUCCEED",
}

var AppEvent_value = map[string]int32{
	"DISPATCH_APP":       0,
	"LAUNCH_APP":         1,
	"LAUNCH_APP_FAILED":  2,
	"LAUNCH_APP_SUCCEED": 3,
	"UPDATE_APP":         4,
	"UPDATE_APP_FAILED":  5,
	"UPDATE_APP_SUCCEED": 6,
	"CANCEL_APP":         7,
	"CANCEL_APP_FAILED":  8,
	"CANCEL_APP_SUCCEED": 9,
}

func (x AppEvent) String() string {
	return proto.EnumName(AppEvent_name, int32(x))
}

func (AppEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_FAILED        NamespaceStatus = 0
	NamespaceStatus_NS_DISPATCHING   NamespaceStatus = 1
	NamespaceStatus_NS_LAUNCHING     NamespaceStatus = 2
	NamespaceStatus_NS_RUNNING       NamespaceStatus = 3
	NamespaceStatus_NS_UPDATING      NamespaceStatus = 4
	NamespaceStatus_NS_UPDATE_FAILED NamespaceStatus = 5
	NamespaceStatus_NS_CANCELING     NamespaceStatus = 6
	NamespaceStatus_NS_CANCELED      NamespaceStatus = 7
)

var NamespaceStatus_name = map[int32]string{
	0: "NS_FAILED",
	1: "NS_DISPATCHING",
	2: "NS_LAUNCHING",
	3: "NS_RUNNING",
	4: "NS_UPDATING",
	5: "NS_UPDATE_FAILED",
	6: "NS_CANCELING",
	7: "NS_CANCELED",
}

var NamespaceStatus_value = map[string]int32{
	"NS_FAILED":        0,
	"NS_DISPATCHING":   1,
	"NS_LAUNCHING":     2,
	"NS_RUNNING":       3,
	"NS_UPDATING":      4,
	"NS_UPDATE_FAILED": 5,
	"NS_CANCELING":     6,
	"NS_CANCELED":      7,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

// Hub namespace event
type NamespaceEvent int32

const (
	NamespaceEvent_DISPATCH_NS       NamespaceEvent = 0
	NamespaceEvent_LAUNCH_NS         NamespaceEvent = 1
	NamespaceEvent_LAUNCH_NS_FAILED  NamespaceEvent = 2
	NamespaceEvent_LAUNCH_NS_SUCCEED NamespaceEvent = 3
	NamespaceEvent_UPDATE_NS         NamespaceEvent = 4
	NamespaceEvent_UPDATE_NS_FAILED  NamespaceEvent = 5
	NamespaceEvent_UPDATE_NS_SUCCEED NamespaceEvent = 6
	NamespaceEvent_CANCEL_NS         NamespaceEvent = 7
	NamespaceEvent_CANCEL_NS_FAILED  NamespaceEvent = 8
	NamespaceEvent_CANCEL_NS_SUCCEED NamespaceEvent = 9
)

var NamespaceEvent_name = map[int32]string{
	0: "DISPATCH_NS",
	1: "LAUNCH_NS",
	2: "LAUNCH_NS_FAILED",
	3: "LAUNCH_NS_SUCCEED",
	4: "UPDATE_NS",
	5: "UPDATE_NS_FAILED",
	6: "UPDATE_NS_SUCCEED",
	7: "CANCEL_NS",
	8: "CANCEL_NS_FAILED",
	9: "CANCEL_NS_SUCCEED",
}

var NamespaceEvent_value = map[string]int32{
	"DISPATCH_NS":       0,
	"LAUNCH_NS":         1,
	"LAUNCH_NS_FAILED":  2,
	"LAUNCH_NS_SUCCEED": 3,
	"UPDATE_NS":         4,
	"UPDATE_NS_FAILED":  5,
	"UPDATE_NS_SUCCEED": 6,
	"CANCEL_NS":         7,
	"CANCEL_NS_FAILED":  8,
	"CANCEL_NS_SUCCEED": 9,
}

func (x NamespaceEvent) String() string {
	return proto.EnumName(NamespaceEvent_name, int32(x))
}

func (NamespaceEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func init() {
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.AppEvent", AppEvent_name, AppEvent_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
	proto.RegisterEnum("common.proto.NamespaceEvent", NamespaceEvent_name, NamespaceEvent_value)
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x19, 0xb8, 0xb7, 0xd0, 0x53, 0x68, 0x87, 0xb9, 0x57, 0x1e, 0x82, 0x85, 0x1b, 0x9f,
	0x60, 0xd2, 0x8e, 0x4a, 0x42, 0x4e, 0x1a, 0x0f, 0x5d, 0x13, 0x24, 0x2c, 0xa1, 0x8d, 0x54, 0x9f,
	0xc7, 0xf8, 0x4e, 0x2e, 0x7c, 0x1b, 0x73, 0x3a, 0x9d, 0xce, 0xe8, 0xf2, 0xff, 0xf5, 0xeb, 0xfc,
	0x39, 0x5f, 0x80, 0xe4, 0x58, 0x5f, 0xae, 0xed, 0x6d, 0xf3, 0x52, 0xb7, 0xb5, 0x9a, 0x1f, 0xeb,
	0xf3, 0xb9, 0xbe, 0xd8, 0xb4, 0xfe, 0x10, 0x10, 0xeb, 0xa6, 0xa1, 0xf6, 0xd0, 0xbe, 0x5e, 0x55,
	0x0a, 0xa0, 0xcb, 0x72, 0x7f, 0xaf, 0x37, 0x5b, 0x53, 0xc8, 0x91, 0xfa, 0x07, 0x19, 0xe7, 0x62,
	0x43, 0xa5, 0xde, 0xe5, 0x8f, 0x1b, 0x7c, 0x90, 0x42, 0x2d, 0x61, 0xc1, 0xe5, 0x56, 0x57, 0x68,
	0xab, 0xb1, 0xca, 0x20, 0xe1, 0xea, 0xa9, 0x42, 0xe4, 0x62, 0xa2, 0x24, 0xcc, 0xb9, 0xa8, 0xca,
	0x42, 0xef, 0xb8, 0xf9, 0xa3, 0x6e, 0x60, 0x39, 0x34, 0xc6, 0xbd, 0xf0, 0xd7, 0x7d, 0x2c, 0xd7,
	0x98, 0x9b, 0x2d, 0xff, 0x67, 0xe4, 0x58, 0x5b, 0x99, 0x42, 0x4e, 0xd7, 0x5f, 0x02, 0x66, 0xba,
	0x69, 0xcc, 0xdb, 0xe9, 0xd2, 0xf2, 0x9f, 0xdd, 0x9e, 0xbd, 0x2e, 0x4b, 0x39, 0xe2, 0xd5, 0x76,
	0x4c, 0x97, 0x05, 0x3f, 0xe5, 0xb3, 0x7b, 0x6a, 0xac, 0x56, 0xa0, 0x82, 0x9a, 0xaa, 0x3c, 0x37,
	0xa6, 0x90, 0x13, 0xc6, 0xfb, 0x55, 0x8c, 0x77, 0x4b, 0x7d, 0xf6, 0x4b, 0x57, 0xa0, 0x82, 0xda,
	0xe1, 0x11, 0xe3, 0x76, 0x6a, 0x87, 0x4f, 0x19, 0xf7, 0xd9, 0xe1, 0x33, 0xc6, 0x83, 0xda, 0xe1,
	0xf1, 0xfa, 0x5d, 0x40, 0x86, 0x87, 0xf3, 0xe9, 0xda, 0x1c, 0x8e, 0xa7, 0x5e, 0xc3, 0x02, 0x62,
	0x24, 0x6f, 0x41, 0x41, 0x8a, 0xf4, 0x4b, 0x82, 0x84, 0x39, 0xd2, 0x0f, 0x07, 0x29, 0x00, 0x52,
	0xa0, 0x20, 0x83, 0x04, 0x29, 0x34, 0xf0, 0x1f, 0xa4, 0x2b, 0x02, 0x01, 0xf6, 0x43, 0xe1, 0xfd,
	0x2d, 0x18, 0x9c, 0xff, 0x53, 0x40, 0x3a, 0x4c, 0xb4, 0x12, 0x32, 0x48, 0x06, 0x09, 0x48, 0x72,
	0xc4, 0x93, 0xfb, 0xe3, 0x22, 0x49, 0xc1, 0x6f, 0x0d, 0xd1, 0x1b, 0xf0, 0x62, 0x90, 0x02, 0x01,
	0x0b, 0x88, 0xfb, 0x55, 0x48, 0x76, 0xe7, 0x10, 0xfd, 0x4e, 0x6f, 0x25, 0x60, 0x23, 0x66, 0xfb,
	0xb3, 0x22, 0xc9, 0x29, 0xb3, 0x43, 0xf4, 0xb7, 0xf7, 0x4a, 0x02, 0x36, 0x7e, 0x8e, 0xba, 0x9f,
	0xc0, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x8e, 0x64, 0x08, 0x1f, 0x03, 0x00, 0x00,
}
