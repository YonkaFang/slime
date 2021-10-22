// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Limiter_RateLimitBackend int32

const (
	Limiter_netEaseLocalFlowControl Limiter_RateLimitBackend = 0
	Limiter_envoyLocalRateLimit     Limiter_RateLimitBackend = 1
)

var Limiter_RateLimitBackend_name = map[int32]string{
	0: "netEaseLocalFlowControl",
	1: "envoyLocalRateLimit",
}

var Limiter_RateLimitBackend_value = map[string]int32{
	"netEaseLocalFlowControl": 0,
	"envoyLocalRateLimit":     1,
}

func (x Limiter_RateLimitBackend) String() string {
	return proto.EnumName(Limiter_RateLimitBackend_name, int32(x))
}

func (Limiter_RateLimitBackend) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3, 0}
}

type Prometheus_Source_Type int32

const (
	Prometheus_Source_Value Prometheus_Source_Type = 0
	Prometheus_Source_Group Prometheus_Source_Type = 1
)

var Prometheus_Source_Type_name = map[int32]string{
	0: "Value",
	1: "Group",
}

var Prometheus_Source_Type_value = map[string]int32{
	"Value": 0,
	"Group": 1,
}

func (x Prometheus_Source_Type) String() string {
	return proto.EnumName(Prometheus_Source_Type_name, int32(x))
}

func (Prometheus_Source_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{7, 0}
}

type Config_Mode int32

const (
	Config_Module     Config_Mode = 0
	Config_BundleItem Config_Mode = 1
)

var Config_Mode_name = map[int32]string{
	0: "Module",
	1: "BundleItem",
}

var Config_Mode_value = map[string]int32{
	"Module":     0,
	"BundleItem": 1,
}

func (x Config_Mode) String() string {
	return proto.EnumName(Config_Mode_name, int32(x))
}

func (Config_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{12, 0}
}

type LocalSource struct {
	Mount                string   `protobuf:"bytes,1,opt,name=mount,proto3" json:"mount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalSource) Reset()         { *m = LocalSource{} }
func (m *LocalSource) String() string { return proto.CompactTextString(m) }
func (*LocalSource) ProtoMessage()    {}
func (*LocalSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}
func (m *LocalSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalSource.Unmarshal(m, b)
}
func (m *LocalSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalSource.Marshal(b, m, deterministic)
}
func (m *LocalSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalSource.Merge(m, src)
}
func (m *LocalSource) XXX_Size() int {
	return xxx_messageInfo_LocalSource.Size(m)
}
func (m *LocalSource) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalSource.DiscardUnknown(m)
}

var xxx_messageInfo_LocalSource proto.InternalMessageInfo

func (m *LocalSource) GetMount() string {
	if m != nil {
		return m.Mount
	}
	return ""
}

type RemoteSource struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteSource) Reset()         { *m = RemoteSource{} }
func (m *RemoteSource) String() string { return proto.CompactTextString(m) }
func (*RemoteSource) ProtoMessage()    {}
func (*RemoteSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}
func (m *RemoteSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteSource.Unmarshal(m, b)
}
func (m *RemoteSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteSource.Marshal(b, m, deterministic)
}
func (m *RemoteSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteSource.Merge(m, src)
}
func (m *RemoteSource) XXX_Size() int {
	return xxx_messageInfo_RemoteSource.Size(m)
}
func (m *RemoteSource) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteSource.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteSource proto.InternalMessageInfo

func (m *RemoteSource) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Plugin struct {
	// Types that are valid to be assigned to WasmSource:
	//	*Plugin_Local
	//	*Plugin_Remote
	WasmSource           isPlugin_WasmSource `protobuf_oneof:"wasm_source"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}
func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Plugin.Unmarshal(m, b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return xxx_messageInfo_Plugin.Size(m)
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

type isPlugin_WasmSource interface {
	isPlugin_WasmSource()
}

type Plugin_Local struct {
	Local *LocalSource `protobuf:"bytes,2,opt,name=local,proto3,oneof"`
}
type Plugin_Remote struct {
	Remote *RemoteSource `protobuf:"bytes,3,opt,name=remote,proto3,oneof"`
}

func (*Plugin_Local) isPlugin_WasmSource()  {}
func (*Plugin_Remote) isPlugin_WasmSource() {}

func (m *Plugin) GetWasmSource() isPlugin_WasmSource {
	if m != nil {
		return m.WasmSource
	}
	return nil
}

func (m *Plugin) GetLocal() *LocalSource {
	if x, ok := m.GetWasmSource().(*Plugin_Local); ok {
		return x.Local
	}
	return nil
}

func (m *Plugin) GetRemote() *RemoteSource {
	if x, ok := m.GetWasmSource().(*Plugin_Remote); ok {
		return x.Remote
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Plugin) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Plugin_OneofMarshaler, _Plugin_OneofUnmarshaler, _Plugin_OneofSizer, []interface{}{
		(*Plugin_Local)(nil),
		(*Plugin_Remote)(nil),
	}
}

func _Plugin_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Plugin)
	// wasm_source
	switch x := m.WasmSource.(type) {
	case *Plugin_Local:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Local); err != nil {
			return err
		}
	case *Plugin_Remote:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Remote); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Plugin.WasmSource has unexpected type %T", x)
	}
	return nil
}

func _Plugin_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Plugin)
	switch tag {
	case 2: // wasm_source.local
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LocalSource)
		err := b.DecodeMessage(msg)
		m.WasmSource = &Plugin_Local{msg}
		return true, err
	case 3: // wasm_source.remote
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RemoteSource)
		err := b.DecodeMessage(msg)
		m.WasmSource = &Plugin_Remote{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Plugin_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Plugin)
	// wasm_source
	switch x := m.WasmSource.(type) {
	case *Plugin_Local:
		s := proto.Size(x.Local)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Plugin_Remote:
		s := proto.Size(x.Remote)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Limiter struct {
	Multicluster         bool                     `protobuf:"varint,2,opt,name=Multicluster,proto3" json:"Multicluster,omitempty"`
	Backend              Limiter_RateLimitBackend `protobuf:"varint,3,opt,name=backend,proto3,enum=slime.config.v1alpha1.Limiter_RateLimitBackend" json:"backend,omitempty"`
	Refresh              *time.Duration           `protobuf:"bytes,4,opt,name=refresh,proto3,stdduration" json:"refresh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Limiter) Reset()         { *m = Limiter{} }
func (m *Limiter) String() string { return proto.CompactTextString(m) }
func (*Limiter) ProtoMessage()    {}
func (*Limiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{3}
}
func (m *Limiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Limiter.Unmarshal(m, b)
}
func (m *Limiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Limiter.Marshal(b, m, deterministic)
}
func (m *Limiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Limiter.Merge(m, src)
}
func (m *Limiter) XXX_Size() int {
	return xxx_messageInfo_Limiter.Size(m)
}
func (m *Limiter) XXX_DiscardUnknown() {
	xxx_messageInfo_Limiter.DiscardUnknown(m)
}

var xxx_messageInfo_Limiter proto.InternalMessageInfo

func (m *Limiter) GetMulticluster() bool {
	if m != nil {
		return m.Multicluster
	}
	return false
}

func (m *Limiter) GetBackend() Limiter_RateLimitBackend {
	if m != nil {
		return m.Backend
	}
	return Limiter_netEaseLocalFlowControl
}

func (m *Limiter) GetRefresh() *time.Duration {
	if m != nil {
		return m.Refresh
	}
	return nil
}

type Global struct {
	Service              string            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Multicluster         string            `protobuf:"bytes,2,opt,name=multicluster,proto3" json:"multicluster,omitempty"`
	IstioNamespace       string            `protobuf:"bytes,3,opt,name=istioNamespace,proto3" json:"istioNamespace,omitempty"`
	SlimeNamespace       string            `protobuf:"bytes,4,opt,name=slimeNamespace,proto3" json:"slimeNamespace,omitempty"`
	Log                  *Log              `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	Misc                 map[string]string `protobuf:"bytes,6,rep,name=misc,proto3" json:"misc,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IstioRev             string            `protobuf:"bytes,7,opt,name=istioRev,proto3" json:"istioRev,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Global) Reset()         { *m = Global{} }
func (m *Global) String() string { return proto.CompactTextString(m) }
func (*Global) ProtoMessage()    {}
func (*Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{4}
}
func (m *Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Global.Unmarshal(m, b)
}
func (m *Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Global.Marshal(b, m, deterministic)
}
func (m *Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Global.Merge(m, src)
}
func (m *Global) XXX_Size() int {
	return xxx_messageInfo_Global.Size(m)
}
func (m *Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Global.DiscardUnknown(m)
}

var xxx_messageInfo_Global proto.InternalMessageInfo

func (m *Global) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Global) GetMulticluster() string {
	if m != nil {
		return m.Multicluster
	}
	return ""
}

func (m *Global) GetIstioNamespace() string {
	if m != nil {
		return m.IstioNamespace
	}
	return ""
}

func (m *Global) GetSlimeNamespace() string {
	if m != nil {
		return m.SlimeNamespace
	}
	return ""
}

func (m *Global) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Global) GetMisc() map[string]string {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Global) GetIstioRev() string {
	if m != nil {
		return m.IstioRev
	}
	return ""
}

type Log struct {
	LogLevel             string   `protobuf:"bytes,1,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	KlogLevel            int32    `protobuf:"varint,2,opt,name=klogLevel,proto3" json:"klogLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{5}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *Log) GetKlogLevel() int32 {
	if m != nil {
		return m.KlogLevel
	}
	return 0
}

type Fence struct {
	WormholePort         []string `protobuf:"bytes,2,rep,name=wormholePort,proto3" json:"wormholePort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fence) Reset()         { *m = Fence{} }
func (m *Fence) String() string { return proto.CompactTextString(m) }
func (*Fence) ProtoMessage()    {}
func (*Fence) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{6}
}
func (m *Fence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fence.Unmarshal(m, b)
}
func (m *Fence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fence.Marshal(b, m, deterministic)
}
func (m *Fence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fence.Merge(m, src)
}
func (m *Fence) XXX_Size() int {
	return xxx_messageInfo_Fence.Size(m)
}
func (m *Fence) XXX_DiscardUnknown() {
	xxx_messageInfo_Fence.DiscardUnknown(m)
}

var xxx_messageInfo_Fence proto.InternalMessageInfo

func (m *Fence) GetWormholePort() []string {
	if m != nil {
		return m.WormholePort
	}
	return nil
}

type Prometheus_Source struct {
	Address              string                                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Handlers             map[string]*Prometheus_Source_Handler `protobuf:"bytes,2,rep,name=handlers,proto3" json:"handlers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *Prometheus_Source) Reset()         { *m = Prometheus_Source{} }
func (m *Prometheus_Source) String() string { return proto.CompactTextString(m) }
func (*Prometheus_Source) ProtoMessage()    {}
func (*Prometheus_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{7}
}
func (m *Prometheus_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prometheus_Source.Unmarshal(m, b)
}
func (m *Prometheus_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prometheus_Source.Marshal(b, m, deterministic)
}
func (m *Prometheus_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prometheus_Source.Merge(m, src)
}
func (m *Prometheus_Source) XXX_Size() int {
	return xxx_messageInfo_Prometheus_Source.Size(m)
}
func (m *Prometheus_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Prometheus_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Prometheus_Source proto.InternalMessageInfo

func (m *Prometheus_Source) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Prometheus_Source) GetHandlers() map[string]*Prometheus_Source_Handler {
	if m != nil {
		return m.Handlers
	}
	return nil
}

type Prometheus_Source_Handler struct {
	Query                string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Type                 Prometheus_Source_Type `protobuf:"varint,2,opt,name=type,proto3,enum=slime.config.v1alpha1.Prometheus_Source_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Prometheus_Source_Handler) Reset()         { *m = Prometheus_Source_Handler{} }
func (m *Prometheus_Source_Handler) String() string { return proto.CompactTextString(m) }
func (*Prometheus_Source_Handler) ProtoMessage()    {}
func (*Prometheus_Source_Handler) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{7, 0}
}
func (m *Prometheus_Source_Handler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prometheus_Source_Handler.Unmarshal(m, b)
}
func (m *Prometheus_Source_Handler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prometheus_Source_Handler.Marshal(b, m, deterministic)
}
func (m *Prometheus_Source_Handler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prometheus_Source_Handler.Merge(m, src)
}
func (m *Prometheus_Source_Handler) XXX_Size() int {
	return xxx_messageInfo_Prometheus_Source_Handler.Size(m)
}
func (m *Prometheus_Source_Handler) XXX_DiscardUnknown() {
	xxx_messageInfo_Prometheus_Source_Handler.DiscardUnknown(m)
}

var xxx_messageInfo_Prometheus_Source_Handler proto.InternalMessageInfo

func (m *Prometheus_Source_Handler) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Prometheus_Source_Handler) GetType() Prometheus_Source_Type {
	if m != nil {
		return m.Type
	}
	return Prometheus_Source_Value
}

type K8S_Source struct {
	Handlers             []string `protobuf:"bytes,1,rep,name=handlers,proto3" json:"handlers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K8S_Source) Reset()         { *m = K8S_Source{} }
func (m *K8S_Source) String() string { return proto.CompactTextString(m) }
func (*K8S_Source) ProtoMessage()    {}
func (*K8S_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{8}
}
func (m *K8S_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8S_Source.Unmarshal(m, b)
}
func (m *K8S_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8S_Source.Marshal(b, m, deterministic)
}
func (m *K8S_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8S_Source.Merge(m, src)
}
func (m *K8S_Source) XXX_Size() int {
	return xxx_messageInfo_K8S_Source.Size(m)
}
func (m *K8S_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_K8S_Source.DiscardUnknown(m)
}

var xxx_messageInfo_K8S_Source proto.InternalMessageInfo

func (m *K8S_Source) GetHandlers() []string {
	if m != nil {
		return m.Handlers
	}
	return nil
}

type Metric struct {
	Prometheus           *Prometheus_Source `protobuf:"bytes,1,opt,name=prometheus,proto3" json:"prometheus,omitempty"`
	K8S                  *K8S_Source        `protobuf:"bytes,2,opt,name=k8s,proto3" json:"k8s,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{9}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetPrometheus() *Prometheus_Source {
	if m != nil {
		return m.Prometheus
	}
	return nil
}

func (m *Metric) GetK8S() *K8S_Source {
	if m != nil {
		return m.K8S
	}
	return nil
}

type General struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *General) Reset()         { *m = General{} }
func (m *General) String() string { return proto.CompactTextString(m) }
func (*General) ProtoMessage()    {}
func (*General) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{10}
}
func (m *General) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_General.Unmarshal(m, b)
}
func (m *General) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_General.Marshal(b, m, deterministic)
}
func (m *General) XXX_Merge(src proto.Message) {
	xxx_messageInfo_General.Merge(m, src)
}
func (m *General) XXX_Size() int {
	return xxx_messageInfo_General.Size(m)
}
func (m *General) XXX_DiscardUnknown() {
	xxx_messageInfo_General.DiscardUnknown(m)
}

var xxx_messageInfo_General proto.InternalMessageInfo

type Bundle struct {
	Modules              []*Bundle_Item `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{11}
}
func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetModules() []*Bundle_Item {
	if m != nil {
		return m.Modules
	}
	return nil
}

type Bundle_Item struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle_Item) Reset()         { *m = Bundle_Item{} }
func (m *Bundle_Item) String() string { return proto.CompactTextString(m) }
func (*Bundle_Item) ProtoMessage()    {}
func (*Bundle_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{11, 0}
}
func (m *Bundle_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle_Item.Unmarshal(m, b)
}
func (m *Bundle_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle_Item.Marshal(b, m, deterministic)
}
func (m *Bundle_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle_Item.Merge(m, src)
}
func (m *Bundle_Item) XXX_Size() int {
	return xxx_messageInfo_Bundle_Item.Size(m)
}
func (m *Bundle_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle_Item proto.InternalMessageInfo

func (m *Bundle_Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Config struct {
	Plugin               *Plugin     `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Limiter              *Limiter    `protobuf:"bytes,2,opt,name=limiter,proto3" json:"limiter,omitempty"`
	Global               *Global     `protobuf:"bytes,3,opt,name=global,proto3" json:"global,omitempty"`
	Fence                *Fence      `protobuf:"bytes,4,opt,name=fence,proto3" json:"fence,omitempty"`
	Metric               *Metric     `protobuf:"bytes,6,opt,name=metric,proto3" json:"metric,omitempty"`
	Name                 string      `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Enable               bool        `protobuf:"varint,7,opt,name=enable,proto3" json:"enable,omitempty"`
	General              *General    `protobuf:"bytes,8,opt,name=general,proto3" json:"general,omitempty"`
	Bundle               *Bundle     `protobuf:"bytes,9,opt,name=bundle,proto3" json:"bundle,omitempty"`
	Mode                 Config_Mode `protobuf:"varint,10,opt,name=mode,proto3,enum=slime.config.v1alpha1.Config_Mode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{12}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetPlugin() *Plugin {
	if m != nil {
		return m.Plugin
	}
	return nil
}

func (m *Config) GetLimiter() *Limiter {
	if m != nil {
		return m.Limiter
	}
	return nil
}

func (m *Config) GetGlobal() *Global {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *Config) GetFence() *Fence {
	if m != nil {
		return m.Fence
	}
	return nil
}

func (m *Config) GetMetric() *Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *Config) GetGeneral() *General {
	if m != nil {
		return m.General
	}
	return nil
}

func (m *Config) GetBundle() *Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *Config) GetMode() Config_Mode {
	if m != nil {
		return m.Mode
	}
	return Config_Module
}

func init() {
	proto.RegisterEnum("slime.config.v1alpha1.Limiter_RateLimitBackend", Limiter_RateLimitBackend_name, Limiter_RateLimitBackend_value)
	proto.RegisterEnum("slime.config.v1alpha1.Prometheus_Source_Type", Prometheus_Source_Type_name, Prometheus_Source_Type_value)
	proto.RegisterEnum("slime.config.v1alpha1.Config_Mode", Config_Mode_name, Config_Mode_value)
	proto.RegisterType((*LocalSource)(nil), "slime.config.v1alpha1.LocalSource")
	proto.RegisterType((*RemoteSource)(nil), "slime.config.v1alpha1.RemoteSource")
	proto.RegisterType((*Plugin)(nil), "slime.config.v1alpha1.Plugin")
	proto.RegisterType((*Limiter)(nil), "slime.config.v1alpha1.Limiter")
	proto.RegisterType((*Global)(nil), "slime.config.v1alpha1.Global")
	proto.RegisterMapType((map[string]string)(nil), "slime.config.v1alpha1.Global.MiscEntry")
	proto.RegisterType((*Log)(nil), "slime.config.v1alpha1.Log")
	proto.RegisterType((*Fence)(nil), "slime.config.v1alpha1.Fence")
	proto.RegisterType((*Prometheus_Source)(nil), "slime.config.v1alpha1.Prometheus_Source")
	proto.RegisterMapType((map[string]*Prometheus_Source_Handler)(nil), "slime.config.v1alpha1.Prometheus_Source.HandlersEntry")
	proto.RegisterType((*Prometheus_Source_Handler)(nil), "slime.config.v1alpha1.Prometheus_Source.Handler")
	proto.RegisterType((*K8S_Source)(nil), "slime.config.v1alpha1.K8S_Source")
	proto.RegisterType((*Metric)(nil), "slime.config.v1alpha1.Metric")
	proto.RegisterType((*General)(nil), "slime.config.v1alpha1.General")
	proto.RegisterType((*Bundle)(nil), "slime.config.v1alpha1.Bundle")
	proto.RegisterType((*Bundle_Item)(nil), "slime.config.v1alpha1.Bundle.Item")
	proto.RegisterType((*Config)(nil), "slime.config.v1alpha1.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x5d, 0x73, 0xdb, 0x44,
	0x17, 0x8e, 0x62, 0x59, 0x8e, 0x8f, 0xd3, 0x8c, 0xdf, 0x7d, 0x81, 0x0a, 0x91, 0x96, 0xa0, 0xce,
	0x50, 0x33, 0x50, 0x99, 0xba, 0x43, 0x09, 0x05, 0x86, 0x21, 0xa5, 0x49, 0x3a, 0x24, 0x4c, 0x66,
	0xcb, 0x70, 0xc1, 0x4d, 0x47, 0x96, 0x4f, 0x64, 0x8d, 0x57, 0x5a, 0xb1, 0x92, 0x9c, 0xf1, 0x2f,
	0xe0, 0x9e, 0xe1, 0x82, 0x5f, 0xc6, 0x2d, 0xff, 0x81, 0x7b, 0x66, 0x98, 0xfd, 0x90, 0xed, 0x84,
	0xca, 0x94, 0xbb, 0x3d, 0x47, 0xcf, 0xf9, 0x7a, 0x9e, 0xdd, 0x23, 0xd8, 0x8d, 0x78, 0x76, 0x99,
	0xc4, 0x41, 0x2e, 0x78, 0xc9, 0xc9, 0x9b, 0x05, 0x4b, 0x52, 0x0c, 0x8c, 0x6f, 0xfe, 0x30, 0x64,
	0xf9, 0x34, 0x7c, 0xe8, 0x3d, 0x88, 0x93, 0x72, 0x5a, 0x8d, 0x83, 0x88, 0xa7, 0xc3, 0x98, 0xc7,
	0x7c, 0xa8, 0xd0, 0xe3, 0xea, 0x52, 0x59, 0xca, 0x50, 0x27, 0x9d, 0xc5, 0xbb, 0x1b, 0x73, 0x1e,
	0x33, 0x5c, 0xa1, 0x26, 0x95, 0x08, 0xcb, 0x84, 0x67, 0xfa, 0xbb, 0x7f, 0x0f, 0x7a, 0x67, 0x3c,
	0x0a, 0xd9, 0x0b, 0x5e, 0x89, 0x08, 0xc9, 0x1b, 0xd0, 0x4e, 0x79, 0x95, 0x95, 0xae, 0x75, 0x60,
	0x0d, 0xba, 0x54, 0x1b, 0xfe, 0x00, 0x76, 0x29, 0xa6, 0xbc, 0x44, 0x83, 0x72, 0xa1, 0x13, 0x4e,
	0x26, 0x02, 0x8b, 0xc2, 0xe0, 0x6a, 0xd3, 0xff, 0xc5, 0x02, 0xe7, 0x82, 0x55, 0x71, 0x92, 0x91,
	0x27, 0xd0, 0x66, 0x32, 0xb3, 0xbb, 0x7d, 0x60, 0x0d, 0x7a, 0x23, 0x3f, 0x78, 0xe5, 0x3c, 0xc1,
	0x5a, 0xf5, 0xd3, 0x2d, 0xaa, 0x43, 0xc8, 0x97, 0xe0, 0x08, 0x55, 0xd0, 0x6d, 0xa9, 0xe0, 0x7b,
	0x0d, 0xc1, 0xeb, 0x5d, 0x9d, 0x6e, 0x51, 0x13, 0x74, 0x74, 0x0b, 0x7a, 0x57, 0x61, 0x91, 0xbe,
	0x2c, 0xd4, 0x07, 0xff, 0x2f, 0x0b, 0x3a, 0x67, 0x49, 0x9a, 0x94, 0x28, 0x88, 0x0f, 0xbb, 0xe7,
	0x15, 0x2b, 0x93, 0x88, 0x55, 0x45, 0x89, 0x42, 0x35, 0xb7, 0x43, 0xaf, 0xf9, 0xc8, 0x73, 0xe8,
	0x8c, 0xc3, 0x68, 0x86, 0xd9, 0x44, 0x95, 0xdf, 0x1b, 0x0d, 0x9b, 0x7a, 0xd7, 0x49, 0x03, 0x1a,
	0x96, 0xa8, 0xce, 0x47, 0x3a, 0x8c, 0xd6, 0xf1, 0xe4, 0x33, 0xe8, 0x08, 0xbc, 0x14, 0x58, 0x4c,
	0x5d, 0x5b, 0x4d, 0xf2, 0x76, 0xa0, 0x05, 0x09, 0x6a, 0x41, 0x82, 0x6f, 0x8c, 0x20, 0x47, 0xf6,
	0x6f, 0x7f, 0xbc, 0x6b, 0xd1, 0x1a, 0xef, 0x9f, 0x42, 0xff, 0x66, 0x5e, 0xf2, 0x0e, 0xdc, 0xce,
	0xb0, 0x7c, 0x16, 0x16, 0xa8, 0x68, 0x3b, 0x66, 0xfc, 0xea, 0x29, 0xcf, 0x4a, 0xc1, 0x59, 0x7f,
	0x8b, 0xdc, 0x86, 0xff, 0x63, 0x36, 0xe7, 0x0b, 0xf5, 0x69, 0x19, 0xda, 0xb7, 0xfc, 0xdf, 0xb7,
	0xc1, 0x39, 0x61, 0x7c, 0x1c, 0x32, 0xa9, 0x5c, 0x81, 0x62, 0x9e, 0x44, 0x58, 0x2b, 0x67, 0x4c,
	0x49, 0x4c, 0x7a, 0x93, 0x98, 0x2e, 0xbd, 0xe6, 0x23, 0xef, 0xc3, 0x5e, 0x52, 0x94, 0x09, 0xff,
	0x2e, 0x4c, 0xb1, 0xc8, 0xc3, 0x48, 0xcb, 0xd3, 0xa5, 0x37, 0xbc, 0x12, 0xa7, 0x08, 0x5b, 0xe1,
	0x6c, 0x8d, 0xbb, 0xee, 0x25, 0x1f, 0x41, 0x8b, 0xf1, 0xd8, 0x6d, 0x2b, 0x66, 0xbc, 0xc6, 0x0b,
	0x12, 0x53, 0x09, 0x23, 0x9f, 0x83, 0x9d, 0x26, 0x45, 0xe4, 0x3a, 0x07, 0xad, 0x41, 0x6f, 0x74,
	0xbf, 0x01, 0xae, 0x07, 0x0d, 0xce, 0x93, 0x22, 0x7a, 0x96, 0x95, 0x62, 0x41, 0x55, 0x10, 0xf1,
	0x60, 0x47, 0x35, 0x49, 0x71, 0xee, 0x76, 0x54, 0x33, 0x4b, 0xdb, 0xfb, 0x14, 0xba, 0x4b, 0x38,
	0xe9, 0x43, 0x6b, 0x86, 0x0b, 0xc3, 0x8e, 0x3c, 0xca, 0x37, 0x31, 0x0f, 0x59, 0x85, 0x86, 0x12,
	0x6d, 0x3c, 0xd9, 0x3e, 0xb4, 0xfc, 0xaf, 0xa0, 0x75, 0xc6, 0x63, 0x99, 0x9b, 0xf1, 0xf8, 0x0c,
	0xe7, 0xc8, 0x4c, 0xdc, 0xd2, 0x26, 0xfb, 0xd0, 0x9d, 0x2d, 0x3f, 0xca, 0x04, 0x6d, 0xba, 0x72,
	0xf8, 0x1f, 0x42, 0xfb, 0x18, 0x33, 0xcd, 0xfe, 0x15, 0x17, 0xe9, 0x94, 0x33, 0xbc, 0xe0, 0xa2,
	0x74, 0xb7, 0x0f, 0x5a, 0x92, 0xfd, 0x75, 0x9f, 0xff, 0xe7, 0x36, 0xfc, 0xef, 0x42, 0xf0, 0x14,
	0xcb, 0x29, 0x56, 0xc5, 0xcb, 0x7f, 0x7b, 0x8b, 0x84, 0xc2, 0xce, 0x34, 0xcc, 0x26, 0x0c, 0x45,
	0xa1, 0xf2, 0xf5, 0x46, 0x8f, 0x1b, 0x38, 0xfb, 0x47, 0xd6, 0xe0, 0xd4, 0x04, 0x6a, 0x0a, 0x97,
	0x79, 0xbc, 0x31, 0x74, 0xcc, 0x27, 0x49, 0xcb, 0x4f, 0x15, 0x8a, 0x9a, 0x2a, 0x6d, 0x90, 0xaf,
	0xc1, 0x2e, 0x17, 0xb9, 0xe6, 0x6a, 0x6f, 0xf4, 0xe0, 0xb5, 0x0b, 0x7e, 0xbf, 0xc8, 0x91, 0xaa,
	0x50, 0x2f, 0x85, 0x5b, 0xd7, 0xca, 0xbf, 0x42, 0x92, 0xe3, 0x75, 0x49, 0x7a, 0xa3, 0x8f, 0xff,
	0xeb, 0x5c, 0xeb, 0x22, 0xee, 0x83, 0x2d, 0x8b, 0x93, 0x2e, 0xb4, 0x7f, 0x90, 0xce, 0xfe, 0x96,
	0x3c, 0x9e, 0x08, 0x5e, 0xe5, 0x7d, 0xcb, 0x1f, 0x00, 0x7c, 0x7b, 0xf8, 0xa2, 0x26, 0xdb, 0x5b,
	0xa3, 0xd4, 0x52, 0x12, 0x2d, 0x6d, 0xff, 0x67, 0x0b, 0x9c, 0x73, 0x2c, 0x45, 0x12, 0x91, 0x53,
	0x80, 0x7c, 0x59, 0x5a, 0xf5, 0xdd, 0x1b, 0x0d, 0x5e, 0xb7, 0x47, 0xba, 0x16, 0x4b, 0x1e, 0x41,
	0x6b, 0x76, 0x58, 0x98, 0x31, 0xdf, 0x6b, 0x48, 0xb1, 0x6a, 0x90, 0x4a, 0xb4, 0xdf, 0x85, 0xce,
	0x09, 0x66, 0x28, 0x42, 0xe6, 0x8f, 0xc1, 0x39, 0xaa, 0x64, 0x83, 0xe4, 0x0b, 0xe8, 0xa4, 0x7c,
	0x52, 0x31, 0xd4, 0x9d, 0x37, 0x2f, 0x64, 0x8d, 0x0f, 0x9e, 0x97, 0x98, 0xd2, 0x3a, 0xc4, 0xf3,
	0xc0, 0x96, 0x0e, 0x42, 0xc0, 0xce, 0xc2, 0xb4, 0x5e, 0x1e, 0xea, 0xec, 0xff, 0x6a, 0x83, 0xf3,
	0x54, 0x25, 0x21, 0x9f, 0x80, 0x93, 0xab, 0xed, 0x6f, 0x86, 0xbe, 0xd3, 0x34, 0xb4, 0x02, 0x51,
	0x03, 0x26, 0x87, 0xd0, 0x61, 0x7a, 0x95, 0x9a, 0x49, 0xef, 0x6e, 0x5e, 0xb8, 0xb4, 0x86, 0xcb,
	0x82, 0xb1, 0x7a, 0xf0, 0xe6, 0x47, 0x71, 0x67, 0xe3, 0x56, 0xa0, 0x06, 0x4c, 0x46, 0xd0, 0xbe,
	0x94, 0xef, 0xce, 0x2c, 0xe5, 0xfd, 0x86, 0x28, 0xf5, 0x36, 0xa9, 0x86, 0xca, 0x52, 0xa9, 0x92,
	0xd7, 0x75, 0x36, 0x96, 0xd2, 0x77, 0x80, 0x1a, 0xf0, 0x92, 0xb1, 0xf6, 0x8a, 0x31, 0xf2, 0x16,
	0x38, 0x98, 0x85, 0x63, 0x86, 0x6a, 0x15, 0xed, 0x50, 0x63, 0x49, 0x1e, 0x62, 0x2d, 0x9c, 0xbb,
	0xb3, 0x91, 0x07, 0x23, 0x2f, 0xad, 0xe1, 0xb2, 0xb9, 0xb1, 0xd2, 0xcd, 0xed, 0x6e, 0x6c, 0x4e,
	0x8b, 0x4b, 0x0d, 0x98, 0x3c, 0x06, 0x3b, 0xe5, 0x13, 0x74, 0x41, 0xbd, 0xd6, 0xa6, 0x1b, 0xa1,
	0xc5, 0x0d, 0xce, 0xf9, 0x04, 0xa9, 0xc2, 0xfb, 0x3e, 0xd8, 0xd2, 0x22, 0x00, 0xce, 0xb9, 0xba,
	0x21, 0xfd, 0x2d, 0xb2, 0x07, 0xa0, 0xb3, 0xcb, 0x8b, 0xd2, 0xb7, 0x8e, 0x3e, 0xf8, 0xf1, 0xbe,
	0x4e, 0x97, 0xf0, 0xa1, 0x3a, 0x0c, 0xf3, 0x59, 0x3c, 0x0c, 0xf3, 0xa4, 0x18, 0xea, 0x02, 0xc3,
	0xba, 0xc0, 0xd8, 0x51, 0x3f, 0xc3, 0x47, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x17, 0x85, 0x03,
	0xea, 0x01, 0x09, 0x00, 0x00,
}