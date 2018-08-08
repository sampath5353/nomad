// Code generated by protoc-gen-go. DO NOT EDIT.
// source: driver.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type RecordRequest struct {
	Driver               string     `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	TaskId               string     `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Event                *TaskEvent `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecordRequest) Reset()         { *m = RecordRequest{} }
func (m *RecordRequest) String() string { return proto.CompactTextString(m) }
func (*RecordRequest) ProtoMessage()    {}
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{0}
}
func (m *RecordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordRequest.Unmarshal(m, b)
}
func (m *RecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordRequest.Marshal(b, m, deterministic)
}
func (dst *RecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordRequest.Merge(dst, src)
}
func (m *RecordRequest) XXX_Size() int {
	return xxx_messageInfo_RecordRequest.Size(m)
}
func (m *RecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecordRequest proto.InternalMessageInfo

func (m *RecordRequest) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *RecordRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordRequest) GetEvent() *TaskEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type TaskEvent struct {
	EventType            string               `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Driver               string               `protobuf:"bytes,3,opt,name=driver,proto3" json:"driver,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Attrs                map[string]string    `protobuf:"bytes,5,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskEvent) Reset()         { *m = TaskEvent{} }
func (m *TaskEvent) String() string { return proto.CompactTextString(m) }
func (*TaskEvent) ProtoMessage()    {}
func (*TaskEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{1}
}
func (m *TaskEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskEvent.Unmarshal(m, b)
}
func (m *TaskEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskEvent.Marshal(b, m, deterministic)
}
func (dst *TaskEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskEvent.Merge(dst, src)
}
func (m *TaskEvent) XXX_Size() int {
	return xxx_messageInfo_TaskEvent.Size(m)
}
func (m *TaskEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskEvent proto.InternalMessageInfo

func (m *TaskEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *TaskEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TaskEvent) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *TaskEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskEvent) GetAttrs() map[string]string {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type RecoverTaskRequest struct {
	TaskId               string       `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Events               []*TaskEvent `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RecoverTaskRequest) Reset()         { *m = RecoverTaskRequest{} }
func (m *RecoverTaskRequest) String() string { return proto.CompactTextString(m) }
func (*RecoverTaskRequest) ProtoMessage()    {}
func (*RecoverTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{2}
}
func (m *RecoverTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverTaskRequest.Unmarshal(m, b)
}
func (m *RecoverTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverTaskRequest.Marshal(b, m, deterministic)
}
func (dst *RecoverTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverTaskRequest.Merge(dst, src)
}
func (m *RecoverTaskRequest) XXX_Size() int {
	return xxx_messageInfo_RecoverTaskRequest.Size(m)
}
func (m *RecoverTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverTaskRequest proto.InternalMessageInfo

func (m *RecoverTaskRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecoverTaskRequest) GetEvents() []*TaskEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

type CreateTaskRequest struct {
	Task                 *TaskInfo `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{3}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetTask() *TaskInfo {
	if m != nil {
		return m.Task
	}
	return nil
}

type TaskInfo struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	User                 string            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	DriverConfig         *_struct.Struct   `protobuf:"bytes,3,opt,name=driver_config,json=driverConfig,proto3" json:"driver_config,omitempty"`
	Env                  map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources            *Resources        `protobuf:"bytes,5,opt,name=resources,proto3" json:"resources,omitempty"`
	Mounts               []*Mount          `protobuf:"bytes,6,rep,name=mounts,proto3" json:"mounts,omitempty"`
	Devices              []*Device         `protobuf:"bytes,7,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{4}
}
func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInfo.Unmarshal(m, b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
}
func (dst *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(dst, src)
}
func (m *TaskInfo) XXX_Size() int {
	return xxx_messageInfo_TaskInfo.Size(m)
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *TaskInfo) GetDriverConfig() *_struct.Struct {
	if m != nil {
		return m.DriverConfig
	}
	return nil
}

func (m *TaskInfo) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *TaskInfo) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskInfo) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *TaskInfo) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type Resources struct {
	// CPU CFS (Completely Fair Scheduler) period. Default: 0 (not specified).
	CpuPeriod int64 `protobuf:"varint,1,opt,name=cpu_period,json=cpuPeriod,proto3" json:"cpu_period,omitempty"`
	// CPU CFS (Completely Fair Scheduler) quota. Default: 0 (not specified).
	CpuQuota int64 `protobuf:"varint,2,opt,name=cpu_quota,json=cpuQuota,proto3" json:"cpu_quota,omitempty"`
	// CPU shares (relative weight vs. other containers). Default: 0 (not specified).
	CpuShares int64 `protobuf:"varint,3,opt,name=cpu_shares,json=cpuShares,proto3" json:"cpu_shares,omitempty"`
	// Memory limit in bytes. Default: 0 (not specified).
	MemoryLimitInBytes int64 `protobuf:"varint,4,opt,name=memory_limit_in_bytes,json=memoryLimitInBytes,proto3" json:"memory_limit_in_bytes,omitempty"`
	// OOMScoreAdj adjusts the oom-killer score. Default: 0 (not specified).
	OomScoreAdj int64 `protobuf:"varint,5,opt,name=oom_score_adj,json=oomScoreAdj,proto3" json:"oom_score_adj,omitempty"`
	// CpusetCpus constrains the allowed set of logical CPUs. Default: "" (not specified).
	CpusetCpus string `protobuf:"bytes,6,opt,name=cpuset_cpus,json=cpusetCpus,proto3" json:"cpuset_cpus,omitempty"`
	// CpusetMems constrains the allowed set of memory nodes. Default: "" (not specified).
	CpusetMems           string   `protobuf:"bytes,7,opt,name=cpuset_mems,json=cpusetMems,proto3" json:"cpuset_mems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{5}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetCpuPeriod() int64 {
	if m != nil {
		return m.CpuPeriod
	}
	return 0
}

func (m *Resources) GetCpuQuota() int64 {
	if m != nil {
		return m.CpuQuota
	}
	return 0
}

func (m *Resources) GetCpuShares() int64 {
	if m != nil {
		return m.CpuShares
	}
	return 0
}

func (m *Resources) GetMemoryLimitInBytes() int64 {
	if m != nil {
		return m.MemoryLimitInBytes
	}
	return 0
}

func (m *Resources) GetOomScoreAdj() int64 {
	if m != nil {
		return m.OomScoreAdj
	}
	return 0
}

func (m *Resources) GetCpusetCpus() string {
	if m != nil {
		return m.CpusetCpus
	}
	return ""
}

func (m *Resources) GetCpusetMems() string {
	if m != nil {
		return m.CpusetMems
	}
	return ""
}

type Mount struct {
	TaskPath             string   `protobuf:"bytes,1,opt,name=task_path,json=taskPath,proto3" json:"task_path,omitempty"`
	HostPath             string   `protobuf:"bytes,2,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	Readonly             bool     `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mount) Reset()         { *m = Mount{} }
func (m *Mount) String() string { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()    {}
func (*Mount) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{6}
}
func (m *Mount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mount.Unmarshal(m, b)
}
func (m *Mount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mount.Marshal(b, m, deterministic)
}
func (dst *Mount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mount.Merge(dst, src)
}
func (m *Mount) XXX_Size() int {
	return xxx_messageInfo_Mount.Size(m)
}
func (m *Mount) XXX_DiscardUnknown() {
	xxx_messageInfo_Mount.DiscardUnknown(m)
}

var xxx_messageInfo_Mount proto.InternalMessageInfo

func (m *Mount) GetTaskPath() string {
	if m != nil {
		return m.TaskPath
	}
	return ""
}

func (m *Mount) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *Mount) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

type Device struct {
	TaskPath             string   `protobuf:"bytes,1,opt,name=task_path,json=taskPath,proto3" json:"task_path,omitempty"`
	HostPath             string   `protobuf:"bytes,2,opt,name=host_path,json=hostPath,proto3" json:"host_path,omitempty"`
	Permissions          string   `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_driver_f9132606fcb84447, []int{7}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetTaskPath() string {
	if m != nil {
		return m.TaskPath
	}
	return ""
}

func (m *Device) GetHostPath() string {
	if m != nil {
		return m.HostPath
	}
	return ""
}

func (m *Device) GetPermissions() string {
	if m != nil {
		return m.Permissions
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
	return fileDescriptor_driver_f9132606fcb84447, []int{8}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
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

func init() {
	proto.RegisterType((*RecordRequest)(nil), "hashicorp.nomad.plugins.drivers.base.proto.RecordRequest")
	proto.RegisterType((*TaskEvent)(nil), "hashicorp.nomad.plugins.drivers.base.proto.TaskEvent")
	proto.RegisterMapType((map[string]string)(nil), "hashicorp.nomad.plugins.drivers.base.proto.TaskEvent.AttrsEntry")
	proto.RegisterType((*RecoverTaskRequest)(nil), "hashicorp.nomad.plugins.drivers.base.proto.RecoverTaskRequest")
	proto.RegisterType((*CreateTaskRequest)(nil), "hashicorp.nomad.plugins.drivers.base.proto.CreateTaskRequest")
	proto.RegisterType((*TaskInfo)(nil), "hashicorp.nomad.plugins.drivers.base.proto.TaskInfo")
	proto.RegisterMapType((map[string]string)(nil), "hashicorp.nomad.plugins.drivers.base.proto.TaskInfo.EnvEntry")
	proto.RegisterType((*Resources)(nil), "hashicorp.nomad.plugins.drivers.base.proto.Resources")
	proto.RegisterType((*Mount)(nil), "hashicorp.nomad.plugins.drivers.base.proto.Mount")
	proto.RegisterType((*Device)(nil), "hashicorp.nomad.plugins.drivers.base.proto.Device")
	proto.RegisterType((*CreateTaskResponse)(nil), "hashicorp.nomad.plugins.drivers.base.proto.CreateTaskResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskRecorderClient is the client API for TaskRecorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskRecorderClient interface {
	Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type taskRecorderClient struct {
	cc *grpc.ClientConn
}

func NewTaskRecorderClient(cc *grpc.ClientConn) TaskRecorderClient {
	return &taskRecorderClient{cc}
}

func (c *taskRecorderClient) Record(ctx context.Context, in *RecordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.drivers.base.proto.TaskRecorder/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskRecorderServer is the server API for TaskRecorder service.
type TaskRecorderServer interface {
	Record(context.Context, *RecordRequest) (*empty.Empty, error)
}

func RegisterTaskRecorderServer(s *grpc.Server, srv TaskRecorderServer) {
	s.RegisterService(&_TaskRecorder_serviceDesc, srv)
}

func _TaskRecorder_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskRecorderServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.drivers.base.proto.TaskRecorder/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskRecorderServer).Record(ctx, req.(*RecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskRecorder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.drivers.base.proto.TaskRecorder",
	HandlerType: (*TaskRecorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _TaskRecorder_Record_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}

// DriverClient is the client API for Driver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DriverClient interface {
	RecoverTask(ctx context.Context, in *RecoverTaskRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
}

type driverClient struct {
	cc *grpc.ClientConn
}

func NewDriverClient(cc *grpc.ClientConn) DriverClient {
	return &driverClient{cc}
}

func (c *driverClient) RecoverTask(ctx context.Context, in *RecoverTaskRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.drivers.base.proto.Driver/RecoverTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, "/hashicorp.nomad.plugins.drivers.base.proto.Driver/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServer is the server API for Driver service.
type DriverServer interface {
	RecoverTask(context.Context, *RecoverTaskRequest) (*empty.Empty, error)
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
}

func RegisterDriverServer(s *grpc.Server, srv DriverServer) {
	s.RegisterService(&_Driver_serviceDesc, srv)
}

func _Driver_RecoverTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoverTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).RecoverTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.drivers.base.proto.Driver/RecoverTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).RecoverTask(ctx, req.(*RecoverTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Driver_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hashicorp.nomad.plugins.drivers.base.proto.Driver/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Driver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hashicorp.nomad.plugins.drivers.base.proto.Driver",
	HandlerType: (*DriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecoverTask",
			Handler:    _Driver_RecoverTask_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Driver_CreateTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "driver.proto",
}

func init() { proto.RegisterFile("driver.proto", fileDescriptor_driver_f9132606fcb84447) }

var fileDescriptor_driver_f9132606fcb84447 = []byte{
	// 836 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xef, 0x6e, 0xdb, 0x36,
	0x10, 0xaf, 0xad, 0x58, 0xb1, 0x4f, 0x0d, 0xb0, 0x11, 0x5b, 0x6b, 0xb8, 0x1b, 0x6a, 0xe8, 0x53,
	0x30, 0x60, 0x2a, 0xea, 0xfd, 0x41, 0x36, 0x6c, 0xc5, 0xda, 0x34, 0xc0, 0x82, 0x35, 0x58, 0xc7,
	0x04, 0xfb, 0x30, 0xa0, 0x10, 0x18, 0xe9, 0x62, 0xab, 0x31, 0x45, 0x96, 0xa4, 0x0c, 0x08, 0xd8,
	0x97, 0x3d, 0x40, 0xdf, 0x65, 0x6f, 0xb5, 0xc7, 0xd8, 0x40, 0x52, 0x4a, 0x94, 0x06, 0x05, 0x26,
	0xf7, 0x93, 0xc8, 0xdf, 0x1d, 0x7f, 0xbc, 0xdf, 0x1d, 0xef, 0x04, 0x77, 0x73, 0x55, 0x6c, 0x50,
	0x25, 0x52, 0x09, 0x23, 0xc8, 0x17, 0x2b, 0xa6, 0x57, 0x45, 0x26, 0x94, 0x4c, 0x4a, 0xc1, 0x59,
	0x9e, 0xc8, 0x75, 0xb5, 0x2c, 0x4a, 0x9d, 0x78, 0x2f, 0x9d, 0x9c, 0x33, 0x8d, 0xde, 0x77, 0xf6,
	0x60, 0x29, 0xc4, 0x72, 0x8d, 0x8f, 0xdc, 0xee, 0xbc, 0xba, 0x78, 0x84, 0x5c, 0x9a, 0xba, 0x31,
	0x7e, 0xf6, 0xae, 0x51, 0x1b, 0x55, 0x65, 0xa6, 0xb1, 0x3e, 0x7c, 0xd7, 0x6a, 0x0a, 0x8e, 0xda,
	0x30, 0x2e, 0xbd, 0x43, 0xfc, 0x76, 0x00, 0x7b, 0x14, 0x33, 0xa1, 0x72, 0x8a, 0x6f, 0x2a, 0xd4,
	0x86, 0xdc, 0x83, 0xd0, 0xc7, 0x30, 0x1d, 0xcc, 0x07, 0xfb, 0x13, 0xda, 0xec, 0xc8, 0x7d, 0xd8,
	0x35, 0x4c, 0x5f, 0xa6, 0x45, 0x3e, 0x1d, 0x7a, 0x83, 0xdd, 0x1e, 0xe7, 0xe4, 0x17, 0x18, 0xe1,
	0x06, 0x4b, 0x33, 0x0d, 0xe6, 0x83, 0xfd, 0x68, 0xf1, 0x4d, 0xf2, 0xff, 0xa5, 0x25, 0x67, 0x4c,
	0x5f, 0x1e, 0xd9, 0xc3, 0xd4, 0x73, 0xc4, 0x7f, 0x0f, 0x61, 0x72, 0x05, 0x92, 0xcf, 0x01, 0x1c,
	0x9c, 0x9a, 0x5a, 0x62, 0x13, 0xcf, 0xc4, 0x21, 0x67, 0xb5, 0x44, 0x72, 0x00, 0x93, 0x2b, 0x3d,
	0x2e, 0xa8, 0x68, 0x31, 0x4b, 0xbc, 0xe2, 0xa4, 0x55, 0x9c, 0x9c, 0xb5, 0x1e, 0xf4, 0xda, 0xb9,
	0x23, 0x32, 0xb8, 0x21, 0x72, 0x0e, 0x51, 0x8e, 0x3a, 0x53, 0x85, 0x34, 0x85, 0x28, 0xa7, 0x3b,
	0xce, 0xd8, 0x85, 0xc8, 0xef, 0x30, 0x62, 0xc6, 0x28, 0x3d, 0x1d, 0xcd, 0x83, 0xfd, 0x68, 0xf1,
	0xd3, 0x56, 0x6a, 0x93, 0xa7, 0x96, 0xe2, 0xa8, 0x34, 0xaa, 0xa6, 0x9e, 0x6e, 0x76, 0x00, 0x70,
	0x0d, 0x92, 0x8f, 0x20, 0xb8, 0xc4, 0xba, 0x51, 0x6c, 0x97, 0xe4, 0x13, 0x18, 0x6d, 0xd8, 0xba,
	0xc2, 0x26, 0xf9, 0x7e, 0xf3, 0xfd, 0xf0, 0x60, 0x10, 0xff, 0x09, 0xc4, 0x56, 0x70, 0x83, 0xca,
	0xf2, 0xb7, 0x65, 0xec, 0x94, 0x6b, 0x70, 0xa3, 0x5c, 0x27, 0x10, 0xba, 0x0c, 0xea, 0xe9, 0xd0,
	0x29, 0xd8, 0xb2, 0x5e, 0x0d, 0x49, 0xfc, 0x0a, 0x3e, 0x3e, 0x54, 0xc8, 0x0c, 0x76, 0x2f, 0xff,
	0x19, 0x76, 0xec, 0x6d, 0xee, 0xe6, 0x68, 0xf1, 0x75, 0xdf, 0x1b, 0x8e, 0xcb, 0x0b, 0x41, 0x1d,
	0x43, 0xfc, 0x4f, 0x00, 0xe3, 0x16, 0x22, 0x04, 0x76, 0x4a, 0xc6, 0xdb, 0x87, 0xe0, 0xd6, 0x16,
	0xab, 0x34, 0xaa, 0x26, 0x2d, 0x6e, 0x4d, 0x7e, 0x80, 0x3d, 0xcf, 0x9c, 0x66, 0xa2, 0xbc, 0x28,
	0x96, 0xcd, 0xcb, 0xbc, 0x7f, 0xeb, 0x6d, 0x9c, 0xba, 0x5e, 0xa1, 0x4d, 0x6b, 0x1e, 0x3a, 0x67,
	0xf2, 0x2b, 0x04, 0x58, 0x6e, 0xa6, 0x3b, 0x2e, 0x3b, 0x3f, 0x6e, 0x13, 0x7b, 0x72, 0x54, 0x6e,
	0x7c, 0x71, 0x2d, 0x13, 0x39, 0x85, 0x89, 0x42, 0x2d, 0x2a, 0x95, 0xa1, 0x7d, 0x36, 0xbd, 0x9b,
	0x84, 0xb6, 0x87, 0xe9, 0x35, 0x0f, 0x39, 0x86, 0x90, 0x8b, 0xca, 0x96, 0x31, 0x74, 0x81, 0x3e,
	0xee, 0xc3, 0x78, 0x62, 0x4f, 0xd2, 0x86, 0x80, 0xbc, 0x80, 0xdd, 0x1c, 0x37, 0x85, 0x8d, 0x6e,
	0xd7, 0x71, 0x2d, 0xfa, 0x70, 0x3d, 0x77, 0x47, 0x69, 0x4b, 0x31, 0xfb, 0x16, 0xc6, 0xad, 0xfc,
	0x5e, 0xcf, 0xf8, 0xdf, 0x01, 0x4c, 0xae, 0x94, 0xda, 0xce, 0xcf, 0x64, 0x95, 0x4a, 0x54, 0x85,
	0xf0, 0x2f, 0x38, 0xa0, 0x93, 0x4c, 0x56, 0x2f, 0x1d, 0x40, 0x1e, 0x80, 0xdd, 0xa4, 0x6f, 0x2a,
	0x61, 0x98, 0xa3, 0x0a, 0xe8, 0x38, 0x93, 0xd5, 0x6f, 0x76, 0xdf, 0x9e, 0xd5, 0x2b, 0xa6, 0x50,
	0xbb, 0xda, 0xfb, 0xb3, 0xa7, 0x0e, 0x20, 0x8f, 0xe1, 0x53, 0x8e, 0x5c, 0xa8, 0x3a, 0x5d, 0x17,
	0xbc, 0x30, 0x69, 0x51, 0xa6, 0xe7, 0xb5, 0x41, 0xed, 0xba, 0x3d, 0xa0, 0xc4, 0x1b, 0x5f, 0x58,
	0xdb, 0x71, 0xf9, 0xcc, 0x5a, 0x48, 0x0c, 0x7b, 0x42, 0xf0, 0x54, 0x67, 0x42, 0x61, 0xca, 0xf2,
	0xd7, 0xae, 0x8a, 0x01, 0x8d, 0x84, 0xe0, 0xa7, 0x16, 0x7b, 0x9a, 0xbf, 0x26, 0x0f, 0x21, 0xca,
	0x64, 0xa5, 0xd1, 0xa4, 0xf6, 0x33, 0x0d, 0x9d, 0x3e, 0xf0, 0xd0, 0xa1, 0xac, 0x74, 0xc7, 0x81,
	0x23, 0xb7, 0xa9, 0xee, 0x38, 0x9c, 0x20, 0xb7, 0xad, 0x34, 0x72, 0x85, 0xb1, 0xea, 0x5c, 0xef,
	0x4a, 0x66, 0x56, 0x4d, 0xf2, 0xc6, 0x16, 0x78, 0xc9, 0xcc, 0xca, 0x1a, 0x57, 0x42, 0x1b, 0x6f,
	0xf4, 0x59, 0x1c, 0x5b, 0xc0, 0x19, 0x67, 0x30, 0x56, 0xc8, 0x72, 0x51, 0xae, 0x6b, 0x27, 0x7c,
	0x4c, 0xaf, 0xf6, 0x71, 0x0e, 0xa1, 0xaf, 0xd5, 0x07, 0xf0, 0xcf, 0x21, 0x92, 0xa8, 0x78, 0xa1,
	0x75, 0x21, 0x4a, 0xdd, 0x0c, 0xcf, 0x2e, 0x14, 0x7f, 0x09, 0xa4, 0x3b, 0x0f, 0xb4, 0x14, 0xa5,
	0xc6, 0xf7, 0x4e, 0xa3, 0x05, 0x87, 0xbb, 0xde, 0xd1, 0xfe, 0x82, 0x50, 0x91, 0x57, 0x10, 0xfa,
	0x35, 0xf9, 0xae, 0x5f, 0x8b, 0x74, 0x7e, 0x61, 0xb3, 0x7b, 0xb7, 0x1a, 0xfd, 0xc8, 0xfe, 0x31,
	0xe3, 0x3b, 0x8b, 0xbf, 0x86, 0x10, 0x3e, 0xf7, 0xa3, 0x7e, 0x09, 0x51, 0x67, 0x6c, 0x92, 0x27,
	0x7d, 0xaf, 0xbb, 0x39, 0x6f, 0xdf, 0x7f, 0x27, 0x79, 0x3b, 0x00, 0xb8, 0x4e, 0x09, 0xe9, 0x35,
	0x51, 0x6e, 0x8d, 0xd6, 0xd9, 0x93, 0x6d, 0x8f, 0xfb, 0x4a, 0xc4, 0x77, 0x9e, 0xed, 0xfe, 0x31,
	0xf2, 0x31, 0x86, 0xee, 0xf3, 0xd5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x75, 0xd8, 0x9e,
	0x9a, 0x08, 0x00, 0x00,
}