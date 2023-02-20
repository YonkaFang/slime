// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service_fence.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Destinations_Status int32

const (
	Destinations_ACTIVE Destinations_Status = 0
	Destinations_EXPIRE Destinations_Status = 1
	// In order to avoid deleting frequently, add status EXPIREWAIT between ACTIVE and EXPIRE.
	// When new metric does not contain host of ACTIVE status, its status will change to EXPIREWAIT. If new metric does not contain
	// host of EXPIREWAIT status, which means this host is not contained in the last two metrics, the status will change to EXPIRE.
	// Otherwise, EXPIREWAIT status will change back to ACTIVE.
	// Hosts of ACTIVE or EXPIREWAIT status are all valid for sidecar.
	// For prometheus metric source, as metric can continuously be watched, we can set status update interval in the future version,
	// refer to RecentlyCalled of RecyclingStrategy. But for accesslog metric source, metric only stores in lazyload controller memory.
	// Metric can not continuously produce after host added to sidecar. So after lazyload controller rebooting, we can not tell whether
	// old host is valid or not until it is removed from sidecar and goes to global-sidecar again.
	// We do not have a proper solution to do same thing for accesslog metric source so far. Need further thinking.
	Destinations_EXPIREWAIT Destinations_Status = 2
)

var Destinations_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "EXPIRE",
	2: "EXPIREWAIT",
}

var Destinations_Status_value = map[string]int32{
	"ACTIVE":     0,
	"EXPIRE":     1,
	"EXPIREWAIT": 2,
}

func (x Destinations_Status) String() string {
	return proto.EnumName(Destinations_Status_name, int32(x))
}

func (Destinations_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{5, 0}
}

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

// Spec Example
//   spec:
//    enable: true
//    host:
//      reviews.default.svc.cluster.local: # static dependency of reviews.default service
//        stable:
//      test/*: {} # static dependency of all services in namespace 'test'
//    namespaceSelector: # Match namespace names, multiple namespaces are 'or' relations, static dependency
//      - foo
//      - bar
//    labelSelector: # Match service label, multiple selectors are 'or' relationship, static dependency
//      - selector:
//          project: back
//      - selector: # labels in same selector are 'and' relationship
//          project: front
//          group: web
//    workloadSelector:
//      labels:
//        group: foo
//        zone: hz
//      fromService: false
type ServiceFenceSpec struct {
	Host map[string]*RecyclingStrategy `protobuf:"bytes,1,rep,name=host,proto3" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Switch to render servicefence as sidecar
	Enable bool `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
	// services in these namespaces are all static dependency, will not expire
	NamespaceSelector []string `protobuf:"bytes,3,rep,name=namespaceSelector,proto3" json:"namespaceSelector,omitempty"`
	// services match one selector of the label selector are all static dependency, will not expire
	LabelSelector        []*Selector       `protobuf:"bytes,4,rep,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	WorkloadSelector     *WorkloadSelector `protobuf:"bytes,5,opt,name=workloadSelector,proto3" json:"workloadSelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceFenceSpec) Reset()         { *m = ServiceFenceSpec{} }
func (m *ServiceFenceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceSpec) ProtoMessage()    {}
func (*ServiceFenceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{1}
}
func (m *ServiceFenceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceSpec.Unmarshal(m, b)
}
func (m *ServiceFenceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceFenceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceSpec.Merge(m, src)
}
func (m *ServiceFenceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceSpec.Size(m)
}
func (m *ServiceFenceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceSpec proto.InternalMessageInfo

func (m *ServiceFenceSpec) GetHost() map[string]*RecyclingStrategy {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ServiceFenceSpec) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *ServiceFenceSpec) GetNamespaceSelector() []string {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *ServiceFenceSpec) GetLabelSelector() []*Selector {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *ServiceFenceSpec) GetWorkloadSelector() *WorkloadSelector {
	if m != nil {
		return m.WorkloadSelector
	}
	return nil
}

type Selector struct {
	Selector             map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{2}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

type WorkloadSelector struct {
	// take effect when labels is empty
	// true: sidecar.workloadSelector.labels = svc.spec.selector
	// false: sidecar.workloadSelector.labels = map[string]string{env.config.global.service: svc.name}
	FromService bool `protobuf:"varint,1,opt,name=fromService,proto3" json:"fromService,omitempty"`
	// top priority, if labels is not empty, sidecar.workloadSelector.labels = sf.spec.workloadSelector.labels
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkloadSelector) Reset()         { *m = WorkloadSelector{} }
func (m *WorkloadSelector) String() string { return proto.CompactTextString(m) }
func (*WorkloadSelector) ProtoMessage()    {}
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{3}
}
func (m *WorkloadSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadSelector.Unmarshal(m, b)
}
func (m *WorkloadSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadSelector.Marshal(b, m, deterministic)
}
func (m *WorkloadSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSelector.Merge(m, src)
}
func (m *WorkloadSelector) XXX_Size() int {
	return xxx_messageInfo_WorkloadSelector.Size(m)
}
func (m *WorkloadSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSelector.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSelector proto.InternalMessageInfo

func (m *WorkloadSelector) GetFromService() bool {
	if m != nil {
		return m.FromService
	}
	return false
}

func (m *WorkloadSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type RecyclingStrategy struct {
	// Configuration that will not be cleaned up
	Stable *RecyclingStrategy_Stable `protobuf:"bytes,1,opt,name=stable,proto3" json:"stable,omitempty"`
	// Configurations that expire after expiration
	Deadline *RecyclingStrategy_Deadline `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Deprecated
	Auto                 *RecyclingStrategy_Auto `protobuf:"bytes,3,opt,name=auto,proto3" json:"auto,omitempty"`
	RecentlyCalled       *Timestamp              `protobuf:"bytes,4,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RecyclingStrategy) Reset()         { *m = RecyclingStrategy{} }
func (m *RecyclingStrategy) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy) ProtoMessage()    {}
func (*RecyclingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4}
}
func (m *RecyclingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy.Unmarshal(m, b)
}
func (m *RecyclingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy.Merge(m, src)
}
func (m *RecyclingStrategy) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy.Size(m)
}
func (m *RecyclingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy proto.InternalMessageInfo

func (m *RecyclingStrategy) GetStable() *RecyclingStrategy_Stable {
	if m != nil {
		return m.Stable
	}
	return nil
}

func (m *RecyclingStrategy) GetDeadline() *RecyclingStrategy_Deadline {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *RecyclingStrategy) GetAuto() *RecyclingStrategy_Auto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *RecyclingStrategy) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

type RecyclingStrategy_Stable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecyclingStrategy_Stable) Reset()         { *m = RecyclingStrategy_Stable{} }
func (m *RecyclingStrategy_Stable) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Stable) ProtoMessage()    {}
func (*RecyclingStrategy_Stable) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4, 0}
}
func (m *RecyclingStrategy_Stable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Stable.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Stable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Stable.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Stable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Stable.Merge(m, src)
}
func (m *RecyclingStrategy_Stable) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Stable.Size(m)
}
func (m *RecyclingStrategy_Stable) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Stable.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Stable proto.InternalMessageInfo

type RecyclingStrategy_Deadline struct {
	Expire               *Timestamp `protobuf:"bytes,1,opt,name=expire,proto3" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Deadline) Reset()         { *m = RecyclingStrategy_Deadline{} }
func (m *RecyclingStrategy_Deadline) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Deadline) ProtoMessage()    {}
func (*RecyclingStrategy_Deadline) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4, 1}
}
func (m *RecyclingStrategy_Deadline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Deadline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Deadline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Deadline.Merge(m, src)
}
func (m *RecyclingStrategy_Deadline) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Size(m)
}
func (m *RecyclingStrategy_Deadline) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Deadline.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Deadline proto.InternalMessageInfo

func (m *RecyclingStrategy_Deadline) GetExpire() *Timestamp {
	if m != nil {
		return m.Expire
	}
	return nil
}

type RecyclingStrategy_Auto struct {
	Duration             *Timestamp `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Auto) Reset()         { *m = RecyclingStrategy_Auto{} }
func (m *RecyclingStrategy_Auto) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Auto) ProtoMessage()    {}
func (*RecyclingStrategy_Auto) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4, 2}
}
func (m *RecyclingStrategy_Auto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Auto.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Auto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Auto.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Auto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Auto.Merge(m, src)
}
func (m *RecyclingStrategy_Auto) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Auto.Size(m)
}
func (m *RecyclingStrategy_Auto) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Auto.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Auto proto.InternalMessageInfo

func (m *RecyclingStrategy_Auto) GetDuration() *Timestamp {
	if m != nil {
		return m.Duration
	}
	return nil
}

type Destinations struct {
	// Deprecated
	RecentlyCalled       *Timestamp          `protobuf:"bytes,1,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	Hosts                []string            `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Status               Destinations_Status `protobuf:"varint,3,opt,name=status,proto3,enum=slime.microservice.lazyload.v1alpha1.Destinations_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Destinations) Reset()         { *m = Destinations{} }
func (m *Destinations) String() string { return proto.CompactTextString(m) }
func (*Destinations) ProtoMessage()    {}
func (*Destinations) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{5}
}
func (m *Destinations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destinations.Unmarshal(m, b)
}
func (m *Destinations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destinations.Marshal(b, m, deterministic)
}
func (m *Destinations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destinations.Merge(m, src)
}
func (m *Destinations) XXX_Size() int {
	return xxx_messageInfo_Destinations.Size(m)
}
func (m *Destinations) XXX_DiscardUnknown() {
	xxx_messageInfo_Destinations.DiscardUnknown(m)
}

var xxx_messageInfo_Destinations proto.InternalMessageInfo

func (m *Destinations) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

func (m *Destinations) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Destinations) GetStatus() Destinations_Status {
	if m != nil {
		return m.Status
	}
	return Destinations_ACTIVE
}

type ServiceFenceStatus struct {
	Domains      map[string]*Destinations `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetricStatus map[string]string        `protobuf:"bytes,3,rep,name=metricStatus,proto3" json:"metricStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Deprecated
	Visitor              map[string]bool `protobuf:"bytes,2,rep,name=visitor,proto3" json:"visitor,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ServiceFenceStatus) Reset()         { *m = ServiceFenceStatus{} }
func (m *ServiceFenceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceStatus) ProtoMessage()    {}
func (*ServiceFenceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{6}
}
func (m *ServiceFenceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceStatus.Unmarshal(m, b)
}
func (m *ServiceFenceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceStatus.Marshal(b, m, deterministic)
}
func (m *ServiceFenceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceStatus.Merge(m, src)
}
func (m *ServiceFenceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceStatus.Size(m)
}
func (m *ServiceFenceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceStatus proto.InternalMessageInfo

func (m *ServiceFenceStatus) GetDomains() map[string]*Destinations {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *ServiceFenceStatus) GetMetricStatus() map[string]string {
	if m != nil {
		return m.MetricStatus
	}
	return nil
}

func (m *ServiceFenceStatus) GetVisitor() map[string]bool {
	if m != nil {
		return m.Visitor
	}
	return nil
}

func init() {
	proto.RegisterEnum("slime.microservice.lazyload.v1alpha1.Destinations_Status", Destinations_Status_name, Destinations_Status_value)
	proto.RegisterType((*Timestamp)(nil), "slime.microservice.lazyload.v1alpha1.Timestamp")
	proto.RegisterType((*ServiceFenceSpec)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceSpec")
	proto.RegisterMapType((map[string]*RecyclingStrategy)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceSpec.HostEntry")
	proto.RegisterType((*Selector)(nil), "slime.microservice.lazyload.v1alpha1.Selector")
	proto.RegisterMapType((map[string]string)(nil), "slime.microservice.lazyload.v1alpha1.Selector.SelectorEntry")
	proto.RegisterType((*WorkloadSelector)(nil), "slime.microservice.lazyload.v1alpha1.WorkloadSelector")
	proto.RegisterMapType((map[string]string)(nil), "slime.microservice.lazyload.v1alpha1.WorkloadSelector.LabelsEntry")
	proto.RegisterType((*RecyclingStrategy)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy")
	proto.RegisterType((*RecyclingStrategy_Stable)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Stable")
	proto.RegisterType((*RecyclingStrategy_Deadline)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Deadline")
	proto.RegisterType((*RecyclingStrategy_Auto)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Auto")
	proto.RegisterType((*Destinations)(nil), "slime.microservice.lazyload.v1alpha1.Destinations")
	proto.RegisterType((*ServiceFenceStatus)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus")
	proto.RegisterMapType((map[string]*Destinations)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.DomainsEntry")
	proto.RegisterMapType((map[string]string)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.MetricStatusEntry")
	proto.RegisterMapType((map[string]bool)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.VisitorEntry")
}

func init() { proto.RegisterFile("service_fence.proto", fileDescriptor_b4b8d9f0db3c7310) }

var fileDescriptor_b4b8d9f0db3c7310 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xcd, 0x6e, 0xe3, 0x36,
	0x10, 0xc7, 0x2b, 0x7f, 0x45, 0x1e, 0x3b, 0x81, 0xc3, 0x06, 0x85, 0xa1, 0x93, 0x61, 0xf4, 0xe0,
	0x43, 0x20, 0x37, 0x2e, 0xd0, 0x36, 0x1f, 0x68, 0xf3, 0xe5, 0x36, 0x69, 0x1b, 0x20, 0xa5, 0x8c,
	0x24, 0x08, 0x0a, 0x04, 0xb4, 0xc4, 0x24, 0x42, 0x24, 0x51, 0x10, 0x29, 0xb7, 0xee, 0xb1, 0x6f,
	0xb2, 0xb7, 0xbd, 0xef, 0x63, 0xec, 0x1b, 0xec, 0xd3, 0x2c, 0x44, 0x4a, 0x82, 0x6c, 0x07, 0x58,
	0xdb, 0xbb, 0x37, 0x52, 0xd4, 0xfc, 0x66, 0xfe, 0x33, 0xa3, 0xa1, 0xe0, 0x6b, 0x4e, 0xa3, 0x89,
	0x6b, 0xd3, 0x87, 0x47, 0x1a, 0xd8, 0xd4, 0x0c, 0x23, 0x26, 0x18, 0xfa, 0x96, 0x7b, 0xae, 0x4f,
	0x4d, 0xdf, 0xb5, 0x23, 0x96, 0x9e, 0x9b, 0x1e, 0xf9, 0x6f, 0xea, 0x31, 0xe2, 0x98, 0x93, 0x3d,
	0xe2, 0x85, 0xcf, 0x64, 0xaf, 0x7b, 0x08, 0xf5, 0x91, 0xeb, 0x53, 0x2e, 0x88, 0x1f, 0xa2, 0x36,
	0x6c, 0x70, 0x6a, 0xb3, 0xc0, 0xe1, 0x6d, 0xad, 0xa3, 0xf5, 0xca, 0x38, 0xdb, 0xa2, 0x1d, 0xa8,
	0x06, 0x24, 0x60, 0xbc, 0x5d, 0xea, 0x68, 0xbd, 0x2a, 0x56, 0x9b, 0xee, 0x87, 0x32, 0xb4, 0x2c,
	0x85, 0xfe, 0x35, 0xf1, 0x6c, 0x85, 0xd4, 0x46, 0x23, 0xa8, 0x3c, 0x33, 0x2e, 0xda, 0x5a, 0xa7,
	0xdc, 0x6b, 0x0c, 0x8e, 0xcd, 0x65, 0xc2, 0x30, 0xe7, 0x29, 0xe6, 0x05, 0xe3, 0x62, 0x18, 0x88,
	0x68, 0x8a, 0x25, 0x0d, 0x7d, 0x03, 0x35, 0x1a, 0x90, 0xb1, 0x47, 0x65, 0x04, 0x3a, 0x4e, 0x77,
	0x68, 0x17, 0xb6, 0x03, 0xe2, 0x53, 0x1e, 0x12, 0x9b, 0x5a, 0xd4, 0xa3, 0xb6, 0x60, 0x51, 0xbb,
	0xdc, 0x29, 0xf7, 0xea, 0x78, 0xf1, 0x00, 0x8d, 0x60, 0xd3, 0x23, 0x63, 0xea, 0xe5, 0x6f, 0x56,
	0x64, 0x90, 0xe6, 0xb2, 0x41, 0x2a, 0x2b, 0x3c, 0x0b, 0x41, 0x63, 0x68, 0xfd, 0xc3, 0xa2, 0x97,
	0xe4, 0xe5, 0x1c, 0x5c, 0xed, 0x68, 0xbd, 0xc6, 0xe0, 0x87, 0xe5, 0xc0, 0xb7, 0x73, 0xd6, 0x78,
	0x81, 0x67, 0x84, 0x50, 0xcf, 0x53, 0x82, 0x5a, 0x50, 0x7e, 0xa1, 0x53, 0x59, 0xa3, 0x3a, 0x4e,
	0x96, 0xe8, 0x0a, 0xaa, 0x13, 0xe2, 0xc5, 0x2a, 0x3b, 0x8d, 0xc1, 0x8f, 0xcb, 0xf9, 0xc5, 0xd4,
	0x9e, 0xda, 0x9e, 0x1b, 0x3c, 0x59, 0x22, 0x22, 0x82, 0x3e, 0x4d, 0xb1, 0xa2, 0x1c, 0x94, 0x7e,
	0xd2, 0xba, 0x6f, 0x34, 0xd0, 0x73, 0x89, 0x77, 0xa0, 0xf3, 0x4c, 0x9a, 0x2a, 0xec, 0xd1, 0x6a,
	0x39, 0xcb, 0x17, 0xaa, 0xa8, 0x39, 0xcd, 0x38, 0x84, 0xcd, 0x99, 0xa3, 0x57, 0xc4, 0xed, 0x14,
	0xc5, 0xd5, 0x8b, 0x31, 0xbe, 0xd7, 0xa0, 0x35, 0x9f, 0x3c, 0xd4, 0x81, 0xc6, 0x63, 0xc4, 0xfc,
	0xb4, 0xa5, 0x24, 0x48, 0xc7, 0xc5, 0x47, 0xe8, 0x1e, 0x6a, 0xb2, 0x82, 0x49, 0x3b, 0x27, 0x5a,
	0x4e, 0xd7, 0x2b, 0x93, 0xf9, 0xa7, 0x84, 0x28, 0x45, 0x29, 0xd1, 0xd8, 0x87, 0x46, 0xe1, 0xf1,
	0x4a, 0x6a, 0xde, 0x56, 0x60, 0x7b, 0xa1, 0x24, 0xe8, 0x06, 0x6a, 0x5c, 0xc8, 0xce, 0xd7, 0x64,
	0x6d, 0x7f, 0x5e, 0xb3, 0xb6, 0xa6, 0x25, 0x29, 0x38, 0xa5, 0xa1, 0xbf, 0x41, 0x77, 0x28, 0x71,
	0x3c, 0x37, 0xc8, 0xba, 0xe6, 0x78, 0x5d, 0xf2, 0x79, 0xca, 0xc1, 0x39, 0x11, 0x5d, 0x43, 0x85,
	0xc4, 0x82, 0xb5, 0xcb, 0x92, 0x7c, 0xb4, 0x2e, 0xf9, 0x24, 0x16, 0x0c, 0x4b, 0x12, 0xba, 0x85,
	0x2d, 0x4c, 0x6d, 0x1a, 0x08, 0x6f, 0x7a, 0x46, 0x3c, 0x8f, 0x3a, 0xed, 0x8a, 0x64, 0xf7, 0x97,
	0x63, 0xe7, 0x53, 0x0e, 0xcf, 0x61, 0x0c, 0x1d, 0x6a, 0x2a, 0x35, 0x86, 0x05, 0x7a, 0x26, 0x05,
	0xfd, 0x06, 0x35, 0xfa, 0x6f, 0xe8, 0x46, 0x59, 0xda, 0x57, 0x76, 0x93, 0x9a, 0x1b, 0x16, 0x54,
	0x12, 0x15, 0xe8, 0x0f, 0xd0, 0x9d, 0x38, 0x22, 0xc2, 0x65, 0xc1, 0xba, 0xc8, 0x1c, 0xd0, 0xfd,
	0xbf, 0x04, 0xcd, 0x73, 0xca, 0x85, 0x1b, 0xc8, 0x3d, 0x7f, 0x25, 0x3b, 0xda, 0x17, 0xc9, 0x4e,
	0xd2, 0xae, 0xc9, 0x00, 0x56, 0x9f, 0x4a, 0x1d, 0xab, 0x0d, 0xfa, 0x4b, 0x36, 0xa5, 0x88, 0xb9,
	0x2c, 0xf0, 0xd6, 0x60, 0x7f, 0x39, 0x37, 0xc5, 0x90, 0x93, 0x7e, 0x14, 0x31, 0xc7, 0x29, 0xa8,
	0xfb, 0x9d, 0x2c, 0x83, 0x88, 0x39, 0x02, 0xa8, 0x9d, 0x9c, 0x8d, 0x2e, 0x6f, 0x86, 0xad, 0xaf,
	0x92, 0xf5, 0xf0, 0xee, 0xfa, 0x12, 0x0f, 0x5b, 0x1a, 0xda, 0x02, 0x50, 0xeb, 0xdb, 0x93, 0xcb,
	0x51, 0xab, 0xd4, 0x7d, 0x57, 0x01, 0x34, 0x73, 0x71, 0x28, 0xf3, 0x07, 0xd8, 0x70, 0x98, 0x4f,
	0xdc, 0x80, 0xa7, 0xa3, 0x6a, 0xb8, 0xc6, 0x1d, 0x24, 0x51, 0xe6, 0xb9, 0xe2, 0xa8, 0x2f, 0x3c,
	0xa3, 0xa2, 0x00, 0x9a, 0x3e, 0x15, 0x91, 0x6b, 0x5b, 0x59, 0x0a, 0x12, 0x2f, 0xbf, 0xaf, 0xed,
	0xe5, 0xaa, 0x00, 0x53, 0xae, 0x66, 0xf8, 0x89, 0xa0, 0x89, 0xcb, 0xdd, 0x64, 0xf6, 0x96, 0x3e,
	0x53, 0xd0, 0x8d, 0xe2, 0xa4, 0x82, 0x52, 0xaa, 0x11, 0x40, 0xb3, 0xa8, 0xf4, 0x95, 0xa1, 0x75,
	0x31, 0x7b, 0xbf, 0x0c, 0x56, 0x2f, 0x77, 0x61, 0xd0, 0x19, 0xbf, 0xc0, 0xf6, 0x82, 0xe6, 0x55,
	0x26, 0xa5, 0x71, 0x00, 0xcd, 0xa2, 0x92, 0x4f, 0xd9, 0xea, 0x05, 0xdb, 0x53, 0xf3, 0x7e, 0x57,
	0x05, 0xef, 0xb2, 0xbe, 0x5c, 0xf4, 0x7d, 0xe6, 0xc4, 0x1e, 0xe5, 0xfd, 0x4c, 0x40, 0x9f, 0x84,
	0x6e, 0x3f, 0x13, 0x31, 0xae, 0xc9, 0xdf, 0xa9, 0xef, 0x3f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0xa9, 0x94, 0x30, 0x65, 0x09, 0x00, 0x00,
}
