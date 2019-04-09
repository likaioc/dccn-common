// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

// App Events operation code
type DCOperation int32

const (
	DCOperation_TASK_CREATE DCOperation = 0
	DCOperation_TASK_CANCEL DCOperation = 1
	DCOperation_TASK_UPDATE DCOperation = 2
	DCOperation_HEARTBEAT   DCOperation = 3
)

var DCOperation_name = map[int32]string{
	0: "TASK_CREATE",
	1: "TASK_CANCEL",
	2: "TASK_UPDATE",
	3: "HEARTBEAT",
}

var DCOperation_value = map[string]int32{
	"TASK_CREATE": 0,
	"TASK_CANCEL": 1,
	"TASK_UPDATE": 2,
	"HEARTBEAT":   3,
}

func (x DCOperation) String() string {
	return proto.EnumName(DCOperation_name, int32(x))
}

func (DCOperation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// Hub app status
type AppStatus int32

const (
	AppStatus_APP_STARTING       AppStatus = 0
	AppStatus_APP_START_SUCCESS  AppStatus = 1
	AppStatus_APP_START_FAILED   AppStatus = 2
	AppStatus_APP_RUNNING        AppStatus = 3
	AppStatus_APP_UPDATING       AppStatus = 4
	AppStatus_APP_UPDATE_SUCCESS AppStatus = 5
	AppStatus_APP_UPDATE_FAILED  AppStatus = 6
	AppStatus_APP_CANCELLING     AppStatus = 7
	AppStatus_APP_CANCELLED      AppStatus = 8
	AppStatus_APP_CANCEL_FAILED  AppStatus = 9
	AppStatus_APP_DONE           AppStatus = 10
)

var AppStatus_name = map[int32]string{
	0:  "APP_STARTING",
	1:  "APP_START_SUCCESS",
	2:  "APP_START_FAILED",
	3:  "APP_RUNNING",
	4:  "APP_UPDATING",
	5:  "APP_UPDATE_SUCCESS",
	6:  "APP_UPDATE_FAILED",
	7:  "APP_CANCELLING",
	8:  "APP_CANCELLED",
	9:  "APP_CANCEL_FAILED",
	10: "APP_DONE",
}

var AppStatus_value = map[string]int32{
	"APP_STARTING":       0,
	"APP_START_SUCCESS":  1,
	"APP_START_FAILED":   2,
	"APP_RUNNING":        3,
	"APP_UPDATING":       4,
	"APP_UPDATE_SUCCESS": 5,
	"APP_UPDATE_FAILED":  6,
	"APP_CANCELLING":     7,
	"APP_CANCELLED":      8,
	"APP_CANCEL_FAILED":  9,
	"APP_DONE":           10,
}

func (x AppStatus) String() string {
	return proto.EnumName(AppStatus_name, int32(x))
}

func (AppStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

// Hub namespace status
type NamespaceStatus int32

const (
	NamespaceStatus_NS_STARTING       NamespaceStatus = 0
	NamespaceStatus_NS_START_SUCCESS  NamespaceStatus = 1
	NamespaceStatus_NS_START_FAILED   NamespaceStatus = 2
	NamespaceStatus_NS_RUNNING        NamespaceStatus = 3
	NamespaceStatus_NS_UPDATING       NamespaceStatus = 4
	NamespaceStatus_NS_UPDATE_SUCCESS NamespaceStatus = 5
	NamespaceStatus_NS_UPDATE_FAILED  NamespaceStatus = 6
	NamespaceStatus_NS_CANCELLING     NamespaceStatus = 7
	NamespaceStatus_NS_CANCELLED      NamespaceStatus = 8
	NamespaceStatus_NS_CANCEL_FAILED  NamespaceStatus = 9
	NamespaceStatus_NS_DONE           NamespaceStatus = 10
)

var NamespaceStatus_name = map[int32]string{
	0:  "NS_STARTING",
	1:  "NS_START_SUCCESS",
	2:  "NS_START_FAILED",
	3:  "NS_RUNNING",
	4:  "NS_UPDATING",
	5:  "NS_UPDATE_SUCCESS",
	6:  "NS_UPDATE_FAILED",
	7:  "NS_CANCELLING",
	8:  "NS_CANCELLED",
	9:  "NS_CANCEL_FAILED",
	10: "NS_DONE",
}

var NamespaceStatus_value = map[string]int32{
	"NS_STARTING":       0,
	"NS_START_SUCCESS":  1,
	"NS_START_FAILED":   2,
	"NS_RUNNING":        3,
	"NS_UPDATING":       4,
	"NS_UPDATE_SUCCESS": 5,
	"NS_UPDATE_FAILED":  6,
	"NS_CANCELLING":     7,
	"NS_CANCELLED":      8,
	"NS_CANCEL_FAILED":  9,
	"NS_DONE":           10,
}

func (x NamespaceStatus) String() string {
	return proto.EnumName(NamespaceStatus_name, int32(x))
}

func (NamespaceStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

// Data center status
type DCStatus int32

const (
	DCStatus_AVAILABLE   DCStatus = 0
	DCStatus_UNAVAILABLE DCStatus = 1
)

var DCStatus_name = map[int32]string{
	0: "AVAILABLE",
	1: "UNAVAILABLE",
}

var DCStatus_value = map[string]int32{
	"AVAILABLE":   0,
	"UNAVAILABLE": 1,
}

func (x DCStatus) String() string {
	return proto.EnumName(DCStatus_name, int32(x))
}

func (DCStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

// Emtpy Message
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// App Data structure
type App struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TypeData:
	//	*App_NamespaceId
	//	*App_Namespace
	TypeData             isApp_TypeData `protobuf_oneof:"type_data"`
	Status               AppStatus      `protobuf:"varint,5,opt,name=status,proto3,enum=common.proto.AppStatus" json:"status,omitempty"`
	Uid                  string         `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Chart                *Chart         `protobuf:"bytes,7,opt,name=chart,proto3" json:"chart,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isApp_TypeData interface {
	isApp_TypeData()
}

type App_NamespaceId struct {
	NamespaceId string `protobuf:"bytes,3,opt,name=namespace_id,json=namespaceId,proto3,oneof"`
}

type App_Namespace struct {
	Namespace *Namespace `protobuf:"bytes,4,opt,name=namespace,proto3,oneof"`
}

func (*App_NamespaceId) isApp_TypeData() {}

func (*App_Namespace) isApp_TypeData() {}

func (m *App) GetTypeData() isApp_TypeData {
	if m != nil {
		return m.TypeData
	}
	return nil
}

func (m *App) GetNamespaceId() string {
	if x, ok := m.GetTypeData().(*App_NamespaceId); ok {
		return x.NamespaceId
	}
	return ""
}

func (m *App) GetNamespace() *Namespace {
	if x, ok := m.GetTypeData().(*App_Namespace); ok {
		return x.Namespace
	}
	return nil
}

func (m *App) GetStatus() AppStatus {
	if m != nil {
		return m.Status
	}
	return AppStatus_APP_STARTING
}

func (m *App) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *App) GetChart() *Chart {
	if m != nil {
		return m.Chart
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*App) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*App_NamespaceId)(nil),
		(*App_Namespace)(nil),
	}
}

type GeoLocation struct {
	Lat                  string   `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lng                  string   `protobuf:"bytes,2,opt,name=lng,proto3" json:"lng,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoLocation) Reset()         { *m = GeoLocation{} }
func (m *GeoLocation) String() string { return proto.CompactTextString(m) }
func (*GeoLocation) ProtoMessage()    {}
func (*GeoLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *GeoLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoLocation.Unmarshal(m, b)
}
func (m *GeoLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoLocation.Marshal(b, m, deterministic)
}
func (m *GeoLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoLocation.Merge(m, src)
}
func (m *GeoLocation) XXX_Size() int {
	return xxx_messageInfo_GeoLocation.Size(m)
}
func (m *GeoLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoLocation.DiscardUnknown(m)
}

var xxx_messageInfo_GeoLocation proto.InternalMessageInfo

func (m *GeoLocation) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *GeoLocation) GetLng() string {
	if m != nil {
		return m.Lng
	}
	return ""
}

func (m *GeoLocation) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

// Data Center structure
type DataCenter struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	GeoLocation          *GeoLocation          `protobuf:"bytes,3,opt,name=geo_location,json=geoLocation,proto3" json:"geo_location,omitempty"`
	Status               DCStatus              `protobuf:"varint,4,opt,name=status,proto3,enum=common.proto.DCStatus" json:"status,omitempty"`
	DcAttributes         *DataCenterAttributes `protobuf:"bytes,5,opt,name=dc_attributes,json=dcAttributes,proto3" json:"dc_attributes,omitempty"`
	DcHeartbeatReport    *DCHeartbeatReport    `protobuf:"bytes,6,opt,name=dc_heartbeat_report,json=dcHeartbeatReport,proto3" json:"dc_heartbeat_report,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DataCenter) Reset()         { *m = DataCenter{} }
func (m *DataCenter) String() string { return proto.CompactTextString(m) }
func (*DataCenter) ProtoMessage()    {}
func (*DataCenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *DataCenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenter.Unmarshal(m, b)
}
func (m *DataCenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenter.Marshal(b, m, deterministic)
}
func (m *DataCenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenter.Merge(m, src)
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

func (m *DataCenter) GetGeoLocation() *GeoLocation {
	if m != nil {
		return m.GeoLocation
	}
	return nil
}

func (m *DataCenter) GetStatus() DCStatus {
	if m != nil {
		return m.Status
	}
	return DCStatus_AVAILABLE
}

func (m *DataCenter) GetDcAttributes() *DataCenterAttributes {
	if m != nil {
		return m.DcAttributes
	}
	return nil
}

func (m *DataCenter) GetDcHeartbeatReport() *DCHeartbeatReport {
	if m != nil {
		return m.DcHeartbeatReport
	}
	return nil
}

type DataCenterAttributes struct {
	WalletAddress        string   `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
	CreationDate         uint64   `protobuf:"varint,2,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	LastModifiedDate     uint64   `protobuf:"varint,3,opt,name=last_modified_date,json=lastModifiedDate,proto3" json:"last_modified_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataCenterAttributes) Reset()         { *m = DataCenterAttributes{} }
func (m *DataCenterAttributes) String() string { return proto.CompactTextString(m) }
func (*DataCenterAttributes) ProtoMessage()    {}
func (*DataCenterAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *DataCenterAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataCenterAttributes.Unmarshal(m, b)
}
func (m *DataCenterAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataCenterAttributes.Marshal(b, m, deterministic)
}
func (m *DataCenterAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataCenterAttributes.Merge(m, src)
}
func (m *DataCenterAttributes) XXX_Size() int {
	return xxx_messageInfo_DataCenterAttributes.Size(m)
}
func (m *DataCenterAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DataCenterAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DataCenterAttributes proto.InternalMessageInfo

func (m *DataCenterAttributes) GetWalletAddress() string {
	if m != nil {
		return m.WalletAddress
	}
	return ""
}

func (m *DataCenterAttributes) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *DataCenterAttributes) GetLastModifiedDate() uint64 {
	if m != nil {
		return m.LastModifiedDate
	}
	return 0
}

type DCHeartbeatReport struct {
	Metrics              string   `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	Report               string   `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	ReportTime           uint64   `protobuf:"varint,3,opt,name=report_time,json=reportTime,proto3" json:"report_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCHeartbeatReport) Reset()         { *m = DCHeartbeatReport{} }
func (m *DCHeartbeatReport) String() string { return proto.CompactTextString(m) }
func (*DCHeartbeatReport) ProtoMessage()    {}
func (*DCHeartbeatReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *DCHeartbeatReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCHeartbeatReport.Unmarshal(m, b)
}
func (m *DCHeartbeatReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCHeartbeatReport.Marshal(b, m, deterministic)
}
func (m *DCHeartbeatReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCHeartbeatReport.Merge(m, src)
}
func (m *DCHeartbeatReport) XXX_Size() int {
	return xxx_messageInfo_DCHeartbeatReport.Size(m)
}
func (m *DCHeartbeatReport) XXX_DiscardUnknown() {
	xxx_messageInfo_DCHeartbeatReport.DiscardUnknown(m)
}

var xxx_messageInfo_DCHeartbeatReport proto.InternalMessageInfo

func (m *DCHeartbeatReport) GetMetrics() string {
	if m != nil {
		return m.Metrics
	}
	return ""
}

func (m *DCHeartbeatReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

func (m *DCHeartbeatReport) GetReportTime() uint64 {
	if m != nil {
		return m.ReportTime
	}
	return 0
}

type AppReport struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Report               string   `protobuf:"bytes,2,opt,name=report,proto3" json:"report,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppReport) Reset()         { *m = AppReport{} }
func (m *AppReport) String() string { return proto.CompactTextString(m) }
func (*AppReport) ProtoMessage()    {}
func (*AppReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *AppReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppReport.Unmarshal(m, b)
}
func (m *AppReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppReport.Marshal(b, m, deterministic)
}
func (m *AppReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppReport.Merge(m, src)
}
func (m *AppReport) XXX_Size() int {
	return xxx_messageInfo_AppReport.Size(m)
}
func (m *AppReport) XXX_DiscardUnknown() {
	xxx_messageInfo_AppReport.DiscardUnknown(m)
}

var xxx_messageInfo_AppReport proto.InternalMessageInfo

func (m *AppReport) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *AppReport) GetReport() string {
	if m != nil {
		return m.Report
	}
	return ""
}

// data center communicate with dc manager
type DCStream struct {
	OpType DCOperation `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=common.proto.DCOperation" json:"op_type,omitempty"`
	// Types that are valid to be assigned to OpPayload:
	//	*DCStream_App
	//	*DCStream_AppReport
	//	*DCStream_DataCenter
	OpPayload            isDCStream_OpPayload `protobuf_oneof:"op_payload"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DCStream) Reset()         { *m = DCStream{} }
func (m *DCStream) String() string { return proto.CompactTextString(m) }
func (*DCStream) ProtoMessage()    {}
func (*DCStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *DCStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCStream.Unmarshal(m, b)
}
func (m *DCStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCStream.Marshal(b, m, deterministic)
}
func (m *DCStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCStream.Merge(m, src)
}
func (m *DCStream) XXX_Size() int {
	return xxx_messageInfo_DCStream.Size(m)
}
func (m *DCStream) XXX_DiscardUnknown() {
	xxx_messageInfo_DCStream.DiscardUnknown(m)
}

var xxx_messageInfo_DCStream proto.InternalMessageInfo

func (m *DCStream) GetOpType() DCOperation {
	if m != nil {
		return m.OpType
	}
	return DCOperation_TASK_CREATE
}

type isDCStream_OpPayload interface {
	isDCStream_OpPayload()
}

type DCStream_App struct {
	App *App `protobuf:"bytes,2,opt,name=app,proto3,oneof"`
}

type DCStream_AppReport struct {
	AppReport *AppReport `protobuf:"bytes,3,opt,name=app_report,json=appReport,proto3,oneof"`
}

type DCStream_DataCenter struct {
	DataCenter *DataCenter `protobuf:"bytes,4,opt,name=data_center,json=dataCenter,proto3,oneof"`
}

func (*DCStream_App) isDCStream_OpPayload() {}

func (*DCStream_AppReport) isDCStream_OpPayload() {}

func (*DCStream_DataCenter) isDCStream_OpPayload() {}

func (m *DCStream) GetOpPayload() isDCStream_OpPayload {
	if m != nil {
		return m.OpPayload
	}
	return nil
}

func (m *DCStream) GetApp() *App {
	if x, ok := m.GetOpPayload().(*DCStream_App); ok {
		return x.App
	}
	return nil
}

func (m *DCStream) GetAppReport() *AppReport {
	if x, ok := m.GetOpPayload().(*DCStream_AppReport); ok {
		return x.AppReport
	}
	return nil
}

func (m *DCStream) GetDataCenter() *DataCenter {
	if x, ok := m.GetOpPayload().(*DCStream_DataCenter); ok {
		return x.DataCenter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DCStream) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DCStream_App)(nil),
		(*DCStream_AppReport)(nil),
		(*DCStream_DataCenter)(nil),
	}
}

type Chart struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IconUrl              string   `protobuf:"bytes,4,opt,name=icon_url,json=iconUrl,proto3" json:"icon_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chart) Reset()         { *m = Chart{} }
func (m *Chart) String() string { return proto.CompactTextString(m) }
func (*Chart) ProtoMessage()    {}
func (*Chart) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{8}
}

func (m *Chart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chart.Unmarshal(m, b)
}
func (m *Chart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chart.Marshal(b, m, deterministic)
}
func (m *Chart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chart.Merge(m, src)
}
func (m *Chart) XXX_Size() int {
	return xxx_messageInfo_Chart.Size(m)
}
func (m *Chart) XXX_DiscardUnknown() {
	xxx_messageInfo_Chart.DiscardUnknown(m)
}

var xxx_messageInfo_Chart proto.InternalMessageInfo

func (m *Chart) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chart) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *Chart) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Chart) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

type ChartDetail struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Repo                 string   `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	AppVersion           string   `protobuf:"bytes,4,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Readme               string   `protobuf:"bytes,5,opt,name=readme,proto3" json:"readme,omitempty"`
	Values               string   `protobuf:"bytes,6,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChartDetail) Reset()         { *m = ChartDetail{} }
func (m *ChartDetail) String() string { return proto.CompactTextString(m) }
func (*ChartDetail) ProtoMessage()    {}
func (*ChartDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{9}
}

func (m *ChartDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChartDetail.Unmarshal(m, b)
}
func (m *ChartDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChartDetail.Marshal(b, m, deterministic)
}
func (m *ChartDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChartDetail.Merge(m, src)
}
func (m *ChartDetail) XXX_Size() int {
	return xxx_messageInfo_ChartDetail.Size(m)
}
func (m *ChartDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ChartDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ChartDetail proto.InternalMessageInfo

func (m *ChartDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChartDetail) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *ChartDetail) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ChartDetail) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *ChartDetail) GetReadme() string {
	if m != nil {
		return m.Readme
	}
	return ""
}

func (m *ChartDetail) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

type Namespace struct {
	Id                   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ClusterId            string          `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ClusterName          string          `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	CreationDate         uint64          `protobuf:"varint,5,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	CpuLimit             float32         `protobuf:"fixed32,6,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	MemLimit             string          `protobuf:"bytes,7,opt,name=mem_limit,json=memLimit,proto3" json:"mem_limit,omitempty"`
	StorageLimit         string          `protobuf:"bytes,8,opt,name=storage_limit,json=storageLimit,proto3" json:"storage_limit,omitempty"`
	NamespaceStatus      NamespaceStatus `protobuf:"varint,9,opt,name=namespace_status,json=namespaceStatus,proto3,enum=common.proto.NamespaceStatus" json:"namespace_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Namespace) Reset()         { *m = Namespace{} }
func (m *Namespace) String() string { return proto.CompactTextString(m) }
func (*Namespace) ProtoMessage()    {}
func (*Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{10}
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

func (m *Namespace) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *Namespace) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *Namespace) GetCreationDate() uint64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Namespace) GetCpuLimit() float32 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *Namespace) GetMemLimit() string {
	if m != nil {
		return m.MemLimit
	}
	return ""
}

func (m *Namespace) GetStorageLimit() string {
	if m != nil {
		return m.StorageLimit
	}
	return ""
}

func (m *Namespace) GetNamespaceStatus() NamespaceStatus {
	if m != nil {
		return m.NamespaceStatus
	}
	return NamespaceStatus_NS_STARTING
}

func init() {
	proto.RegisterEnum("common.proto.DCOperation", DCOperation_name, DCOperation_value)
	proto.RegisterEnum("common.proto.AppStatus", AppStatus_name, AppStatus_value)
	proto.RegisterEnum("common.proto.NamespaceStatus", NamespaceStatus_name, NamespaceStatus_value)
	proto.RegisterEnum("common.proto.DCStatus", DCStatus_name, DCStatus_value)
	proto.RegisterType((*Empty)(nil), "common.proto.Empty")
	proto.RegisterType((*App)(nil), "common.proto.App")
	proto.RegisterType((*GeoLocation)(nil), "common.proto.GeoLocation")
	proto.RegisterType((*DataCenter)(nil), "common.proto.DataCenter")
	proto.RegisterType((*DataCenterAttributes)(nil), "common.proto.DataCenterAttributes")
	proto.RegisterType((*DCHeartbeatReport)(nil), "common.proto.DCHeartbeatReport")
	proto.RegisterType((*AppReport)(nil), "common.proto.AppReport")
	proto.RegisterType((*DCStream)(nil), "common.proto.DCStream")
	proto.RegisterType((*Chart)(nil), "common.proto.Chart")
	proto.RegisterType((*ChartDetail)(nil), "common.proto.ChartDetail")
	proto.RegisterType((*Namespace)(nil), "common.proto.Namespace")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1073 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4b, 0x6e, 0xe3, 0x46,
	0x13, 0x16, 0xf5, 0x66, 0x51, 0x92, 0xa9, 0xb6, 0xc7, 0xbf, 0x06, 0x3f, 0x06, 0xe3, 0xd0, 0x18,
	0xc0, 0x31, 0x02, 0x07, 0x50, 0x16, 0x09, 0x90, 0x6c, 0x68, 0x89, 0xb1, 0x8c, 0x51, 0x68, 0x83,
	0x94, 0x67, 0x4b, 0xb4, 0xc9, 0x1e, 0x0f, 0x01, 0x52, 0x6c, 0x90, 0xad, 0x09, 0xbc, 0xcd, 0x09,
	0x72, 0x84, 0x1c, 0x23, 0xbb, 0x1c, 0x26, 0x9b, 0x20, 0xa7, 0x08, 0xfa, 0x41, 0x8a, 0xd2, 0xd8,
	0x80, 0x57, 0xec, 0xfa, 0xaa, 0xba, 0xbe, 0x7a, 0x74, 0x15, 0x61, 0x10, 0x66, 0x69, 0x9a, 0xad,
	0x2f, 0x68, 0x9e, 0xb1, 0x0c, 0xed, 0x48, 0x56, 0x0f, 0x3a, 0x4e, 0x4a, 0xd9, 0xa3, 0xf5, 0x5b,
	0x13, 0x5a, 0x36, 0xa5, 0x68, 0x04, 0xcd, 0x38, 0x9a, 0x68, 0x27, 0xda, 0x99, 0xee, 0x35, 0xe3,
	0x08, 0x21, 0x68, 0xaf, 0x71, 0x4a, 0x26, 0x4d, 0x81, 0x88, 0x33, 0x3a, 0x85, 0x01, 0xff, 0x16,
	0x14, 0x87, 0x24, 0x88, 0xa3, 0x49, 0x8b, 0xeb, 0x16, 0x0d, 0xcf, 0xa8, 0xd0, 0xeb, 0x08, 0x7d,
	0x0f, 0x7a, 0x25, 0x4e, 0xda, 0x27, 0xda, 0x99, 0x31, 0xfd, 0xdf, 0x45, 0x9d, 0xfb, 0xc2, 0x2d,
	0xd5, 0x8b, 0x86, 0xb7, 0xb5, 0x45, 0xdf, 0x42, 0xb7, 0x60, 0x98, 0x6d, 0x8a, 0x49, 0xe7, 0x44,
	0x3b, 0x1b, 0xed, 0xdf, 0xb2, 0x29, 0xf5, 0x85, 0xda, 0x53, 0x66, 0xc8, 0x84, 0xd6, 0x26, 0x8e,
	0x26, 0x5d, 0x11, 0x21, 0x3f, 0xa2, 0xaf, 0xa1, 0x13, 0x7e, 0xc2, 0x39, 0x9b, 0xf4, 0x04, 0xef,
	0xe1, 0xae, 0x87, 0x19, 0x57, 0x79, 0xd2, 0xe2, 0xd2, 0x00, 0x9d, 0x3d, 0x52, 0x12, 0x44, 0x98,
	0x61, 0xeb, 0x3d, 0x18, 0x57, 0x24, 0x5b, 0x66, 0x21, 0x66, 0x71, 0xb6, 0xe6, 0x8e, 0x13, 0xcc,
	0x54, 0x31, 0xf8, 0x51, 0x20, 0xeb, 0x07, 0x55, 0x0c, 0x7e, 0x44, 0x13, 0xe8, 0x85, 0xd9, 0x66,
	0xcd, 0xf2, 0x47, 0x59, 0x06, 0xaf, 0x14, 0xad, 0x3f, 0x9b, 0x00, 0x73, 0xcc, 0xf0, 0x8c, 0xac,
	0x19, 0xc9, 0x5f, 0x54, 0xd8, 0x9f, 0x60, 0xf0, 0x40, 0xb2, 0x20, 0x51, 0x01, 0x08, 0x8f, 0xc6,
	0xf4, 0xf5, 0x6e, 0xf8, 0xb5, 0x08, 0x3d, 0xe3, 0xa1, 0x16, 0xee, 0x45, 0x55, 0xb8, 0xb6, 0x28,
	0xdc, 0xf1, 0xee, 0xbd, 0xf9, 0x6c, 0xaf, 0x6e, 0x57, 0x30, 0x8c, 0xc2, 0x00, 0x33, 0x96, 0xc7,
	0xf7, 0x1b, 0x46, 0x64, 0xbd, 0x8d, 0xa9, 0xb5, 0x77, 0xad, 0x4a, 0xc1, 0xae, 0x2c, 0xbd, 0x41,
	0x14, 0x6e, 0x25, 0x74, 0x03, 0x87, 0x51, 0x18, 0x7c, 0x22, 0x38, 0x67, 0xf7, 0x04, 0xb3, 0x20,
	0x27, 0x34, 0xcb, 0x99, 0x68, 0x88, 0x31, 0x7d, 0xbb, 0x1f, 0xc5, 0xa2, 0xb4, 0xf3, 0x84, 0x99,
	0x37, 0x8e, 0xc2, 0x3d, 0xc8, 0xfa, 0x5d, 0x83, 0xa3, 0xa7, 0x78, 0xd1, 0x3b, 0x18, 0xfd, 0x8a,
	0x93, 0x84, 0xb0, 0x00, 0x47, 0x51, 0x4e, 0x8a, 0x42, 0x15, 0x74, 0x28, 0x51, 0x5b, 0x82, 0xe8,
	0x14, 0x86, 0x61, 0x4e, 0x44, 0x55, 0x78, 0x63, 0x65, 0x91, 0xdb, 0xde, 0xa0, 0x04, 0xe7, 0x98,
	0x11, 0xf4, 0x0d, 0xa0, 0x04, 0x17, 0x2c, 0x48, 0xb3, 0x28, 0xfe, 0x18, 0x93, 0x48, 0x5a, 0xb6,
	0x84, 0xa5, 0xc9, 0x35, 0xbf, 0x28, 0x05, 0xb7, 0xb6, 0x3e, 0xc2, 0xf8, 0x8b, 0xd0, 0x79, 0xf3,
	0x53, 0xc2, 0xf2, 0x38, 0x2c, 0xe3, 0x28, 0x45, 0x74, 0x0c, 0x5d, 0x55, 0x05, 0xd9, 0x5f, 0x25,
	0xa1, 0xb7, 0x60, 0xc8, 0x53, 0xc0, 0xe2, 0xb4, 0x64, 0x03, 0x09, 0xad, 0xe2, 0x94, 0x58, 0x0b,
	0xd0, 0x6d, 0x4a, 0x95, 0xff, 0x53, 0x68, 0x61, 0x4a, 0x85, 0x6f, 0x63, 0x3a, 0xfe, 0x62, 0x0e,
	0x3c, 0xae, 0x7d, 0x8e, 0xca, 0xfa, 0x57, 0x83, 0x3e, 0xef, 0x79, 0x4e, 0x70, 0x8a, 0xa6, 0xd0,
	0xcb, 0x68, 0xc0, 0x5f, 0xba, 0xf0, 0x36, 0xda, 0x7f, 0x54, 0xf3, 0xd9, 0x0d, 0x25, 0xb9, 0x7c,
	0x54, 0xdd, 0x8c, 0xae, 0x1e, 0x29, 0x41, 0xef, 0x24, 0x7b, 0xf3, 0x19, 0xf6, 0x45, 0x43, 0xf2,
	0xff, 0x00, 0x80, 0x29, 0x2d, 0x9b, 0xde, 0x7a, 0x6a, 0xd2, 0xab, 0x8c, 0xf8, 0xa4, 0xe3, 0x2a,
	0xbd, 0x1f, 0xc1, 0xe0, 0x63, 0x17, 0x84, 0xa2, 0xcd, 0x6a, 0x49, 0x4c, 0x9e, 0x7b, 0x7e, 0x8b,
	0x86, 0x07, 0x51, 0x25, 0x5d, 0x0e, 0x00, 0x32, 0x1a, 0x50, 0xfc, 0x98, 0x64, 0x38, 0xb2, 0x12,
	0xe8, 0x88, 0xb1, 0xae, 0xc6, 0x4a, 0xab, 0x8d, 0x15, 0x82, 0x36, 0x8f, 0xae, 0x1c, 0x35, 0x7e,
	0x46, 0x27, 0x60, 0x44, 0xa4, 0x08, 0xf3, 0x98, 0x56, 0x93, 0xa6, 0x7b, 0x75, 0x08, 0xbd, 0x86,
	0x7e, 0x1c, 0x66, 0xeb, 0x60, 0x93, 0x27, 0x22, 0x34, 0xdd, 0xeb, 0x71, 0xf9, 0x2e, 0x4f, 0xac,
	0x3f, 0x34, 0x30, 0x04, 0xdd, 0x9c, 0x30, 0x1c, 0x27, 0x2f, 0x26, 0x9d, 0x40, 0xef, 0x33, 0xc9,
	0x8b, 0x2d, 0x61, 0x29, 0xf2, 0x77, 0xc1, 0x8b, 0x58, 0x6a, 0x25, 0x1f, 0xaf, 0xeb, 0x07, 0x65,
	0x20, 0xba, 0x8c, 0xa3, 0x94, 0x88, 0x29, 0x15, 0x5d, 0xe6, 0x12, 0xc7, 0x3f, 0xe3, 0x64, 0x43,
	0x0a, 0xb5, 0xff, 0x94, 0x64, 0xfd, 0xd5, 0x04, 0xbd, 0x5a, 0xb0, 0x2f, 0x5a, 0x3e, 0x6f, 0x00,
	0xc2, 0x64, 0x53, 0x30, 0x92, 0x57, 0x3b, 0xdd, 0xd3, 0x15, 0x72, 0x1d, 0xa1, 0xaf, 0x60, 0x50,
	0xaa, 0xc5, 0x55, 0x19, 0xa2, 0xa1, 0x30, 0x57, 0xfe, 0x17, 0xf6, 0xc6, 0xae, 0xf3, 0xc4, 0xd8,
	0xfd, 0x1f, 0xf4, 0x90, 0x6e, 0x82, 0x24, 0x4e, 0x63, 0xb9, 0x22, 0x9a, 0x5e, 0x3f, 0xa4, 0x9b,
	0x25, 0x97, 0xb9, 0x32, 0x25, 0xa9, 0x52, 0xf6, 0x04, 0x43, 0x3f, 0x25, 0xa9, 0x54, 0x9e, 0xc2,
	0xb0, 0x60, 0x59, 0x8e, 0x1f, 0x88, 0x32, 0xe8, 0x0b, 0x83, 0x81, 0x02, 0xa5, 0xd1, 0x02, 0xcc,
	0xed, 0xbf, 0x49, 0xad, 0x43, 0x5d, 0xbc, 0xf8, 0x37, 0xcf, 0xfc, 0x7d, 0xd4, 0x56, 0x3c, 0x58,
	0xef, 0x02, 0xe7, 0x37, 0x60, 0xd4, 0xa6, 0x02, 0x1d, 0x80, 0xb1, 0xb2, 0xfd, 0xf7, 0xc1, 0xcc,
	0x73, 0xec, 0x95, 0x63, 0x36, 0xb6, 0x80, 0xed, 0xce, 0x9c, 0xa5, 0xa9, 0x55, 0xc0, 0xdd, 0xed,
	0x9c, 0x5b, 0x34, 0xd1, 0x10, 0xf4, 0x85, 0x63, 0x7b, 0xab, 0x4b, 0xc7, 0x5e, 0x99, 0xad, 0xf3,
	0x7f, 0x34, 0x31, 0xdb, 0x7e, 0xf9, 0xd7, 0x1a, 0xd8, 0xb7, 0xb7, 0x81, 0xbf, 0xb2, 0xbd, 0xd5,
	0xb5, 0x7b, 0x65, 0x36, 0xd0, 0x2b, 0x18, 0x57, 0x48, 0xe0, 0xdf, 0xcd, 0x66, 0x8e, 0xef, 0x9b,
	0x1a, 0x3a, 0x02, 0x73, 0x0b, 0xff, 0x6c, 0x5f, 0x2f, 0x9d, 0xb9, 0xd9, 0xe4, 0x64, 0x1c, 0xf5,
	0xee, 0x5c, 0x97, 0xdf, 0x6e, 0x95, 0xfe, 0x04, 0x39, 0x47, 0xda, 0xe8, 0x18, 0x50, 0x85, 0x38,
	0x95, 0xc3, 0x4e, 0xc9, 0xa3, 0x70, 0xe5, 0xb1, 0x8b, 0x10, 0x8c, 0x38, 0x2c, 0xd3, 0x59, 0x72,
	0x17, 0x3d, 0x34, 0x86, 0x61, 0x0d, 0x73, 0xe6, 0x66, 0xbf, 0xbc, 0x2d, 0xa1, 0xf2, 0xb6, 0x8e,
	0x06, 0xd0, 0xe7, 0xf0, 0xfc, 0xc6, 0x75, 0x4c, 0x38, 0xff, 0x5b, 0x83, 0x83, 0xbd, 0x02, 0xf3,
	0x88, 0x5d, 0xbf, 0x9e, 0xef, 0x11, 0x98, 0x25, 0x50, 0x4b, 0xf7, 0x10, 0x0e, 0x2a, 0xb4, 0xca,
	0x76, 0x04, 0xe0, 0xfa, 0xb5, 0x64, 0xa5, 0xaf, 0x5a, 0xae, 0xaf, 0x60, 0x5c, 0x02, 0xf5, 0x54,
	0x25, 0xc5, 0x7e, 0xa6, 0x63, 0x18, 0xba, 0xfe, 0x6e, 0xa2, 0x26, 0x0c, 0xb6, 0x90, 0xc8, 0x53,
	0x5e, 0xdd, 0x4f, 0xd3, 0x80, 0x9e, 0xeb, 0x97, 0x59, 0x9e, 0xcb, 0x05, 0x2b, 0xb2, 0x1b, 0x82,
	0x6e, 0x7f, 0xb0, 0xaf, 0x97, 0xf6, 0xe5, 0x52, 0x3d, 0x8e, 0x3b, 0x77, 0x0b, 0x68, 0xf7, 0x5d,
	0xf1, 0xea, 0xbe, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x67, 0x0c, 0xd0, 0x9f, 0x8d, 0x09, 0x00,
	0x00,
}
