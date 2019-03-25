// Code generated by protoc-gen-go. DO NOT EDIT.
// source: taskmgr/v1/taskmgr.proto

package gwtaskmgr

import (
	context "context"
	fmt "fmt"
	common "github.com/Ankr-network/dccn-common/protos/common"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The dccn client request message containing the user's token
type CreateTaskRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Schedule             string   `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Replica              int32    `protobuf:"varint,5,opt,name=replica,proto3" json:"replica,omitempty"`
	DataCenterName       string   `protobuf:"bytes,6,opt,name=data_center_name,json=dataCenterName,proto3" json:"data_center_name,omitempty"`
	Cpu                  int32    `protobuf:"varint,7,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               int32    `protobuf:"varint,8,opt,name=memory,proto3" json:"memory,omitempty"`
	Disk                 int32    `protobuf:"varint,9,opt,name=disk,proto3" json:"disk,omitempty"`
	Uid                  string   `protobuf:"bytes,11,opt,name=uid,proto3" json:"uid,omitempty"`
	ChartName            string   `protobuf:"bytes,12,opt,name=chart_name,json=chartName,proto3" json:"chart_name,omitempty"`
	ChartVer             string   `protobuf:"bytes,13,opt,name=chart_ver,json=chartVer,proto3" json:"chart_ver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{0}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTaskRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateTaskRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *CreateTaskRequest) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *CreateTaskRequest) GetReplica() int32 {
	if m != nil {
		return m.Replica
	}
	return 0
}

func (m *CreateTaskRequest) GetDataCenterName() string {
	if m != nil {
		return m.DataCenterName
	}
	return ""
}

func (m *CreateTaskRequest) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *CreateTaskRequest) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *CreateTaskRequest) GetDisk() int32 {
	if m != nil {
		return m.Disk
	}
	return 0
}

func (m *CreateTaskRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CreateTaskRequest) GetChartName() string {
	if m != nil {
		return m.ChartName
	}
	return ""
}

func (m *CreateTaskRequest) GetChartVer() string {
	if m != nil {
		return m.ChartVer
	}
	return ""
}

type CreateTaskResponse struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{1}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskListRequest struct {
	TaskFilter           *TaskFilter `protobuf:"bytes,1,opt,name=task_filter,json=taskFilter,proto3" json:"task_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskListRequest) Reset()         { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()    {}
func (*TaskListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{2}
}

func (m *TaskListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListRequest.Unmarshal(m, b)
}
func (m *TaskListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListRequest.Marshal(b, m, deterministic)
}
func (m *TaskListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListRequest.Merge(m, src)
}
func (m *TaskListRequest) XXX_Size() int {
	return xxx_messageInfo_TaskListRequest.Size(m)
}
func (m *TaskListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListRequest proto.InternalMessageInfo

func (m *TaskListRequest) GetTaskFilter() *TaskFilter {
	if m != nil {
		return m.TaskFilter
	}
	return nil
}

type TaskListResponse struct {
	Tasks                []*common.Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskListResponse) Reset()         { *m = TaskListResponse{} }
func (m *TaskListResponse) String() string { return proto.CompactTextString(m) }
func (*TaskListResponse) ProtoMessage()    {}
func (*TaskListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{3}
}

func (m *TaskListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskListResponse.Unmarshal(m, b)
}
func (m *TaskListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskListResponse.Marshal(b, m, deterministic)
}
func (m *TaskListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskListResponse.Merge(m, src)
}
func (m *TaskListResponse) XXX_Size() int {
	return xxx_messageInfo_TaskListResponse.Size(m)
}
func (m *TaskListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskListResponse proto.InternalMessageInfo

func (m *TaskListResponse) GetTasks() []*common.Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type TaskFilter struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskFilter) Reset()         { *m = TaskFilter{} }
func (m *TaskFilter) String() string { return proto.CompactTextString(m) }
func (*TaskFilter) ProtoMessage()    {}
func (*TaskFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{4}
}

func (m *TaskFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskFilter.Unmarshal(m, b)
}
func (m *TaskFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskFilter.Marshal(b, m, deterministic)
}
func (m *TaskFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskFilter.Merge(m, src)
}
func (m *TaskFilter) XXX_Size() int {
	return xxx_messageInfo_TaskFilter.Size(m)
}
func (m *TaskFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TaskFilter proto.InternalMessageInfo

func (m *TaskFilter) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type TaskID struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskID) Reset()         { *m = TaskID{} }
func (m *TaskID) String() string { return proto.CompactTextString(m) }
func (*TaskID) ProtoMessage()    {}
func (*TaskID) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{5}
}

func (m *TaskID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskID.Unmarshal(m, b)
}
func (m *TaskID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskID.Marshal(b, m, deterministic)
}
func (m *TaskID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskID.Merge(m, src)
}
func (m *TaskID) XXX_Size() int {
	return xxx_messageInfo_TaskID.Size(m)
}
func (m *TaskID) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskID.DiscardUnknown(m)
}

var xxx_messageInfo_TaskID proto.InternalMessageInfo

func (m *TaskID) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type UpdateTaskRequest struct {
	Task                 *common.Task `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{6}
}

func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(m, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetTask() *common.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskOverviewResponse struct {
	ClusterCount         int32    `protobuf:"varint,1,opt,name=cluster_count,json=clusterCount,proto3" json:"cluster_count,omitempty"`
	EnvironmentCount     int32    `protobuf:"varint,2,opt,name=environment_count,json=environmentCount,proto3" json:"environment_count,omitempty"`
	RegionCount          int32    `protobuf:"varint,3,opt,name=region_count,json=regionCount,proto3" json:"region_count,omitempty"`
	TotalTaskCount       int32    `protobuf:"varint,4,opt,name=total_task_count,json=totalTaskCount,proto3" json:"total_task_count,omitempty"`
	HealthTaskCount      int32    `protobuf:"varint,5,opt,name=health_task_count,json=healthTaskCount,proto3" json:"health_task_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskOverviewResponse) Reset()         { *m = TaskOverviewResponse{} }
func (m *TaskOverviewResponse) String() string { return proto.CompactTextString(m) }
func (*TaskOverviewResponse) ProtoMessage()    {}
func (*TaskOverviewResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{7}
}

func (m *TaskOverviewResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskOverviewResponse.Unmarshal(m, b)
}
func (m *TaskOverviewResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskOverviewResponse.Marshal(b, m, deterministic)
}
func (m *TaskOverviewResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskOverviewResponse.Merge(m, src)
}
func (m *TaskOverviewResponse) XXX_Size() int {
	return xxx_messageInfo_TaskOverviewResponse.Size(m)
}
func (m *TaskOverviewResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskOverviewResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskOverviewResponse proto.InternalMessageInfo

func (m *TaskOverviewResponse) GetClusterCount() int32 {
	if m != nil {
		return m.ClusterCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetEnvironmentCount() int32 {
	if m != nil {
		return m.EnvironmentCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetRegionCount() int32 {
	if m != nil {
		return m.RegionCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetTotalTaskCount() int32 {
	if m != nil {
		return m.TotalTaskCount
	}
	return 0
}

func (m *TaskOverviewResponse) GetHealthTaskCount() int32 {
	if m != nil {
		return m.HealthTaskCount
	}
	return 0
}

type TaskLeaderBoardResponse struct {
	List                 []*TaskLeaderBoardDetail `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TaskLeaderBoardResponse) Reset()         { *m = TaskLeaderBoardResponse{} }
func (m *TaskLeaderBoardResponse) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardResponse) ProtoMessage()    {}
func (*TaskLeaderBoardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{8}
}

func (m *TaskLeaderBoardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardResponse.Unmarshal(m, b)
}
func (m *TaskLeaderBoardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardResponse.Marshal(b, m, deterministic)
}
func (m *TaskLeaderBoardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardResponse.Merge(m, src)
}
func (m *TaskLeaderBoardResponse) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardResponse.Size(m)
}
func (m *TaskLeaderBoardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardResponse proto.InternalMessageInfo

func (m *TaskLeaderBoardResponse) GetList() []*TaskLeaderBoardDetail {
	if m != nil {
		return m.List
	}
	return nil
}

type TaskLeaderBoardDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               float64  `protobuf:"fixed64,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskLeaderBoardDetail) Reset()         { *m = TaskLeaderBoardDetail{} }
func (m *TaskLeaderBoardDetail) String() string { return proto.CompactTextString(m) }
func (*TaskLeaderBoardDetail) ProtoMessage()    {}
func (*TaskLeaderBoardDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbc0cfcc64f892be, []int{9}
}

func (m *TaskLeaderBoardDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskLeaderBoardDetail.Unmarshal(m, b)
}
func (m *TaskLeaderBoardDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskLeaderBoardDetail.Marshal(b, m, deterministic)
}
func (m *TaskLeaderBoardDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskLeaderBoardDetail.Merge(m, src)
}
func (m *TaskLeaderBoardDetail) XXX_Size() int {
	return xxx_messageInfo_TaskLeaderBoardDetail.Size(m)
}
func (m *TaskLeaderBoardDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskLeaderBoardDetail.DiscardUnknown(m)
}

var xxx_messageInfo_TaskLeaderBoardDetail proto.InternalMessageInfo

func (m *TaskLeaderBoardDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskLeaderBoardDetail) GetNumber() float64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "gwtaskmgr.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "gwtaskmgr.CreateTaskResponse")
	proto.RegisterType((*TaskListRequest)(nil), "gwtaskmgr.TaskListRequest")
	proto.RegisterType((*TaskListResponse)(nil), "gwtaskmgr.TaskListResponse")
	proto.RegisterType((*TaskFilter)(nil), "gwtaskmgr.TaskFilter")
	proto.RegisterType((*TaskID)(nil), "gwtaskmgr.TaskID")
	proto.RegisterType((*UpdateTaskRequest)(nil), "gwtaskmgr.UpdateTaskRequest")
	proto.RegisterType((*TaskOverviewResponse)(nil), "gwtaskmgr.TaskOverviewResponse")
	proto.RegisterType((*TaskLeaderBoardResponse)(nil), "gwtaskmgr.TaskLeaderBoardResponse")
	proto.RegisterType((*TaskLeaderBoardDetail)(nil), "gwtaskmgr.TaskLeaderBoardDetail")
}

func init() { proto.RegisterFile("taskmgr/v1/taskmgr.proto", fileDescriptor_dbc0cfcc64f892be) }

var fileDescriptor_dbc0cfcc64f892be = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0xf5, 0x6b, 0x8d, 0x64, 0x5b, 0x5a, 0xdb, 0x32, 0x41, 0xdb, 0xa8, 0xcc, 0xa2, 0x85,
	0xe0, 0xa2, 0x12, 0xea, 0x16, 0x3d, 0xb8, 0x39, 0x45, 0x8e, 0x01, 0x21, 0x71, 0x1c, 0x10, 0xf9,
	0xbb, 0x09, 0x2b, 0x72, 0x43, 0x11, 0x26, 0xb9, 0xcc, 0x72, 0x29, 0xc3, 0xd7, 0xbc, 0x42, 0x9e,
	0x27, 0x4f, 0x91, 0x27, 0x08, 0x90, 0x73, 0x9e, 0x21, 0xd8, 0x59, 0x5a, 0x92, 0x65, 0x09, 0x48,
	0x4e, 0x9c, 0xf9, 0xe6, 0x9b, 0x6f, 0x67, 0x76, 0x66, 0x09, 0xa6, 0xa4, 0xe9, 0x75, 0xe4, 0x8b,
	0xfe, 0xf4, 0xef, 0x7e, 0x6e, 0xf6, 0x12, 0xc1, 0x25, 0x27, 0x35, 0xff, 0x26, 0x07, 0xac, 0x43,
	0x9f, 0x73, 0x3f, 0x64, 0x7d, 0x9a, 0x04, 0x7d, 0x1a, 0xc7, 0x5c, 0x52, 0x19, 0xf0, 0x38, 0xd5,
	0x44, 0x6b, 0xc7, 0xe5, 0x51, 0xc4, 0xe3, 0xbe, 0xfe, 0x68, 0xd0, 0xfe, 0x54, 0x80, 0xd6, 0x40,
	0x30, 0x2a, 0xd9, 0x4b, 0x9a, 0x5e, 0x3b, 0xec, 0x7d, 0xc6, 0x52, 0x49, 0x08, 0x94, 0x62, 0x1a,
	0x31, 0xd3, 0xe8, 0x18, 0xdd, 0x9a, 0x83, 0xb6, 0xc2, 0xe4, 0x6d, 0xc2, 0xcc, 0x82, 0xc6, 0x94,
	0x4d, 0x76, 0xa1, 0x1c, 0x44, 0xd4, 0x67, 0x66, 0x11, 0x41, 0xed, 0x10, 0x0b, 0x36, 0x52, 0x77,
	0xc2, 0xbc, 0x2c, 0x64, 0x66, 0x09, 0x03, 0x33, 0x9f, 0x98, 0x50, 0x15, 0x2c, 0x09, 0x03, 0x97,
	0x9a, 0xe5, 0x8e, 0xd1, 0x2d, 0x3b, 0x77, 0x2e, 0xe9, 0x42, 0xd3, 0xa3, 0x92, 0x8e, 0x5c, 0x16,
	0x4b, 0x26, 0x46, 0x78, 0x7e, 0x05, 0xb3, 0xb7, 0x14, 0x3e, 0x40, 0xf8, 0xb9, 0xaa, 0xa4, 0x09,
	0x45, 0x37, 0xc9, 0xcc, 0x2a, 0xe6, 0x2b, 0x93, 0xb4, 0xa1, 0x12, 0xb1, 0x88, 0x8b, 0x5b, 0x73,
	0x03, 0xc1, 0xdc, 0x53, 0x35, 0x7b, 0x41, 0x7a, 0x6d, 0xd6, 0x10, 0x45, 0x5b, 0x65, 0x67, 0x81,
	0x67, 0xd6, 0x51, 0x5a, 0x99, 0xe4, 0x08, 0xc0, 0x9d, 0x50, 0x21, 0xf5, 0x99, 0x0d, 0x0c, 0xd4,
	0x10, 0xc1, 0xe3, 0x0e, 0x40, 0x3b, 0xa3, 0x29, 0x13, 0xe6, 0xa6, 0xee, 0x07, 0x81, 0xd7, 0x4c,
	0xd8, 0x7f, 0x01, 0x59, 0xbc, 0xbe, 0x34, 0xe1, 0x71, 0xca, 0xc8, 0x3e, 0x54, 0xd5, 0x4c, 0x46,
	0x81, 0x97, 0x5f, 0x61, 0x45, 0xb9, 0x43, 0xcf, 0x1e, 0xc2, 0xb6, 0x22, 0x3e, 0x0b, 0x52, 0x79,
	0x77, 0xd7, 0xff, 0x41, 0x1d, 0xb9, 0xef, 0x82, 0x50, 0x32, 0x81, 0xfc, 0xfa, 0xe9, 0x5e, 0x6f,
	0x36, 0xd5, 0x9e, 0x4a, 0xb8, 0xc0, 0xa0, 0x03, 0x72, 0x66, 0xdb, 0x8f, 0xa0, 0x39, 0x97, 0xca,
	0xcf, 0xed, 0x42, 0x59, 0x31, 0x52, 0xd3, 0xe8, 0x14, 0xbb, 0xf5, 0x53, 0xd2, 0x5b, 0x9c, 0x35,
	0x0a, 0x39, 0x9a, 0x60, 0xff, 0x0e, 0x30, 0xd7, 0x5d, 0x5f, 0xef, 0x31, 0x54, 0x14, 0x6d, 0x78,
	0xbe, 0x9e, 0xf2, 0x3f, 0xb4, 0x5e, 0x25, 0xde, 0xd2, 0x02, 0xfd, 0x01, 0x25, 0x15, 0xce, 0xbb,
	0x59, 0x55, 0x07, 0xc6, 0xed, 0x2f, 0x06, 0xec, 0x2a, 0xf7, 0x6a, 0xca, 0xc4, 0x34, 0x60, 0x37,
	0xb3, 0x4e, 0x7e, 0x83, 0x4d, 0x37, 0xcc, 0x52, 0xb5, 0x09, 0x2e, 0xcf, 0x62, 0x89, 0x4a, 0x65,
	0xa7, 0x91, 0x83, 0x03, 0x85, 0x91, 0x3f, 0xa1, 0xc5, 0xe2, 0x69, 0x20, 0x78, 0x1c, 0xb1, 0x58,
	0xe6, 0xc4, 0x02, 0x12, 0x9b, 0x0b, 0x01, 0x4d, 0x3e, 0x86, 0x86, 0x60, 0x7e, 0xc0, 0xe3, 0x9c,
	0x57, 0x44, 0x5e, 0x5d, 0x63, 0x9a, 0xd2, 0x85, 0xa6, 0xe4, 0x92, 0x86, 0x23, 0xec, 0x54, 0xd3,
	0x4a, 0x48, 0xdb, 0x42, 0x5c, 0x55, 0xaa, 0x99, 0x27, 0xd0, 0x9a, 0x30, 0x1a, 0xca, 0xc9, 0x22,
	0x55, 0x2f, 0xf4, 0xb6, 0x0e, 0xcc, 0xb8, 0xf6, 0x15, 0xec, 0xe3, 0xa0, 0x18, 0xf5, 0x98, 0x78,
	0xcc, 0xa9, 0xf0, 0x66, 0x5d, 0xfe, 0x0b, 0xa5, 0x30, 0x48, 0x65, 0x3e, 0xae, 0xce, 0xd2, 0xd0,
	0x17, 0x32, 0xce, 0x99, 0xa4, 0x41, 0xe8, 0x20, 0xdb, 0x1e, 0xc0, 0xde, 0xca, 0xf0, 0xca, 0x67,
	0xdb, 0x86, 0x4a, 0x9c, 0x45, 0x63, 0x26, 0xf0, 0x62, 0x0c, 0x27, 0xf7, 0x4e, 0xbf, 0x95, 0xa0,
	0xaa, 0x54, 0x2e, 0x7d, 0x41, 0xc6, 0x00, 0xf3, 0x25, 0x26, 0x87, 0x0b, 0x65, 0x3c, 0xf8, 0x35,
	0x58, 0x47, 0x6b, 0xa2, 0xba, 0x23, 0x7b, 0xff, 0xc3, 0xe7, 0xaf, 0x1f, 0x0b, 0x2d, 0xbb, 0x81,
	0x7f, 0xa9, 0xbe, 0x8b, 0x8c, 0x33, 0xe3, 0x84, 0xbc, 0x81, 0x8d, 0xbb, 0x75, 0x25, 0xd6, 0x72,
	0xa3, 0xf3, 0xe7, 0x60, 0x1d, 0xac, 0x8c, 0xe5, 0xea, 0x04, 0xd5, 0x1b, 0x04, 0xb4, 0xba, 0xba,
	0x0d, 0x72, 0x09, 0x30, 0xa0, 0xb1, 0xcb, 0x70, 0x3a, 0xa4, 0xb5, 0x94, 0x3e, 0x3c, 0xb7, 0x76,
	0xee, 0x6f, 0xdf, 0x93, 0x28, 0x91, 0xb7, 0x0f, 0xea, 0x44, 0x05, 0x55, 0xe7, 0x53, 0xa8, 0xbd,
	0xc8, 0x84, 0xcf, 0x7e, 0x4a, 0xad, 0x8d, 0x6a, 0x4d, 0xbb, 0xae, 0xd5, 0x12, 0x25, 0xa0, 0xc4,
	0xde, 0x02, 0xcc, 0xdf, 0xc6, 0xbd, 0x8b, 0x7d, 0xf0, 0x64, 0x7e, 0xa8, 0xcc, 0x0c, 0xb3, 0x94,
	0xf2, 0x05, 0x34, 0x16, 0xdf, 0x0d, 0x59, 0x95, 0x6d, 0xfd, 0xba, 0x54, 0xfe, 0xf2, 0x2b, 0xb3,
	0x7f, 0x21, 0x93, 0xfc, 0x87, 0x34, 0xdf, 0xa5, 0xd5, 0x52, 0xf6, 0xfa, 0xdd, 0x9c, 0xa9, 0x1d,
	0x62, 0xb1, 0x6d, 0xbb, 0x95, 0x4f, 0x07, 0x29, 0x63, 0x45, 0x39, 0x33, 0x4e, 0xc6, 0x15, 0x94,
	0xfb, 0xe7, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x5a, 0x62, 0x09, 0xca, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskMgrClient is the client API for TaskMgr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskMgrClient interface {
	// Sends request to start a task and list task
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error)
	CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error)
	TaskOverview(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskOverviewResponse, error)
	TaskLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskLeaderBoardResponse, error)
}

type taskMgrClient struct {
	cc *grpc.ClientConn
}

func NewTaskMgrClient(cc *grpc.ClientConn) TaskMgrClient {
	return &taskMgrClient{cc}
}

func (c *taskMgrClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListResponse, error) {
	out := new(TaskListResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) CancelTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/CancelTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) PurgeTask(ctx context.Context, in *TaskID, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/PurgeTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskOverview(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskOverviewResponse, error) {
	out := new(TaskOverviewResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskOverview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskMgrClient) TaskLeaderBoard(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (*TaskLeaderBoardResponse, error) {
	out := new(TaskLeaderBoardResponse)
	err := c.cc.Invoke(ctx, "/gwtaskmgr.TaskMgr/TaskLeaderBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskMgrServer is the server API for TaskMgr service.
type TaskMgrServer interface {
	// Sends request to start a task and list task
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListResponse, error)
	CancelTask(context.Context, *TaskID) (*common.Empty, error)
	PurgeTask(context.Context, *TaskID) (*common.Empty, error)
	UpdateTask(context.Context, *UpdateTaskRequest) (*common.Empty, error)
	TaskOverview(context.Context, *common.Empty) (*TaskOverviewResponse, error)
	TaskLeaderBoard(context.Context, *common.Empty) (*TaskLeaderBoardResponse, error)
}

// UnimplementedTaskMgrServer can be embedded to have forward compatible implementations.
type UnimplementedTaskMgrServer struct {
}

func (*UnimplementedTaskMgrServer) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (*UnimplementedTaskMgrServer) TaskList(ctx context.Context, req *TaskListRequest) (*TaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (*UnimplementedTaskMgrServer) CancelTask(ctx context.Context, req *TaskID) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (*UnimplementedTaskMgrServer) PurgeTask(ctx context.Context, req *TaskID) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeTask not implemented")
}
func (*UnimplementedTaskMgrServer) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*common.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (*UnimplementedTaskMgrServer) TaskOverview(ctx context.Context, req *common.Empty) (*TaskOverviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskOverview not implemented")
}
func (*UnimplementedTaskMgrServer) TaskLeaderBoard(ctx context.Context, req *common.Empty) (*TaskLeaderBoardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskLeaderBoard not implemented")
}

func RegisterTaskMgrServer(s *grpc.Server, srv TaskMgrServer) {
	s.RegisterService(&_TaskMgr_serviceDesc, srv)
}

func _TaskMgr_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/CancelTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).CancelTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_PurgeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).PurgeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/PurgeTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).PurgeTask(ctx, req.(*TaskID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).UpdateTask(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskOverview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskOverview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskOverview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskOverview(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskMgr_TaskLeaderBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMgrServer).TaskLeaderBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gwtaskmgr.TaskMgr/TaskLeaderBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMgrServer).TaskLeaderBoard(ctx, req.(*common.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskMgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gwtaskmgr.TaskMgr",
	HandlerType: (*TaskMgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskMgr_CreateTask_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskMgr_TaskList_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _TaskMgr_CancelTask_Handler,
		},
		{
			MethodName: "PurgeTask",
			Handler:    _TaskMgr_PurgeTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TaskMgr_UpdateTask_Handler,
		},
		{
			MethodName: "TaskOverview",
			Handler:    _TaskMgr_TaskOverview_Handler,
		},
		{
			MethodName: "TaskLeaderBoard",
			Handler:    _TaskMgr_TaskLeaderBoard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taskmgr/v1/taskmgr.proto",
}
