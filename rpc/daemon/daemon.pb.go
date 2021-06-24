// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: rpc/daemon/daemon.proto

package daemon

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	common "github.com/telepresenceio/telepresence/rpc/v2/common"
	manager "github.com/telepresenceio/telepresence/rpc/v2/manager"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DaemonStatus_ErrType int32

const (
	DaemonStatus_UNSPECIFIED DaemonStatus_ErrType = 0
	DaemonStatus_NO_NETWORK  DaemonStatus_ErrType = 3
)

// Enum value maps for DaemonStatus_ErrType.
var (
	DaemonStatus_ErrType_name = map[int32]string{
		0: "UNSPECIFIED",
		3: "NO_NETWORK",
	}
	DaemonStatus_ErrType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"NO_NETWORK":  3,
	}
)

func (x DaemonStatus_ErrType) Enum() *DaemonStatus_ErrType {
	p := new(DaemonStatus_ErrType)
	*p = x
	return p
}

func (x DaemonStatus_ErrType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DaemonStatus_ErrType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_daemon_daemon_proto_enumTypes[0].Descriptor()
}

func (DaemonStatus_ErrType) Type() protoreflect.EnumType {
	return &file_rpc_daemon_daemon_proto_enumTypes[0]
}

func (x DaemonStatus_ErrType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DaemonStatus_ErrType.Descriptor instead.
func (DaemonStatus_ErrType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_daemon_daemon_proto_rawDescGZIP(), []int{0, 0}
}

type DaemonStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error DaemonStatus_ErrType `protobuf:"varint,1,opt,name=error,proto3,enum=telepresence.daemon.DaemonStatus_ErrType" json:"error,omitempty"`
	Dns   string               `protobuf:"bytes,2,opt,name=dns,proto3" json:"dns,omitempty"`
}

func (x *DaemonStatus) Reset() {
	*x = DaemonStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_daemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaemonStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaemonStatus) ProtoMessage() {}

func (x *DaemonStatus) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_daemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaemonStatus.ProtoReflect.Descriptor instead.
func (*DaemonStatus) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_daemon_proto_rawDescGZIP(), []int{0}
}

func (x *DaemonStatus) GetError() DaemonStatus_ErrType {
	if x != nil {
		return x.Error
	}
	return DaemonStatus_UNSPECIFIED
}

func (x *DaemonStatus) GetDns() string {
	if x != nil {
		return x.Dns
	}
	return ""
}

type Paths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *Paths) Reset() {
	*x = Paths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_daemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paths) ProtoMessage() {}

func (x *Paths) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_daemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paths.ProtoReflect.Descriptor instead.
func (*Paths) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_daemon_proto_rawDescGZIP(), []int{1}
}

func (x *Paths) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

// DNS configuration for the local DNS resolver
type DNSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// local_ip is the address of the local DNS server. Only used by Linux systems that have no
	// systemd.resolved configured. Defaults to the first line of /etc/resolv.conf
	LocalIp []byte `protobuf:"bytes,1,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	// remote_ip is the address of the kube-dns.kube-system, dns-default.openshift-dns, or similar service,
	RemoteIp []byte `protobuf:"bytes,2,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// Suffixes to exclude
	ExcludeSuffixes []string `protobuf:"bytes,3,rep,name=exclude_suffixes,json=excludeSuffixes,proto3" json:"exclude_suffixes,omitempty"`
	// Suffixes to include. Has higher prio than the excludes
	IncludeSuffixes []string `protobuf:"bytes,4,rep,name=include_suffixes,json=includeSuffixes,proto3" json:"include_suffixes,omitempty"`
	// The maximum time wait for a cluster side host lookup.
	LookupTimeout int64 `protobuf:"varint,5,opt,name=lookup_timeout,json=lookupTimeout,proto3" json:"lookup_timeout,omitempty"`
}

func (x *DNSConfig) Reset() {
	*x = DNSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_daemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSConfig) ProtoMessage() {}

func (x *DNSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_daemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSConfig.ProtoReflect.Descriptor instead.
func (*DNSConfig) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_daemon_proto_rawDescGZIP(), []int{2}
}

func (x *DNSConfig) GetLocalIp() []byte {
	if x != nil {
		return x.LocalIp
	}
	return nil
}

func (x *DNSConfig) GetRemoteIp() []byte {
	if x != nil {
		return x.RemoteIp
	}
	return nil
}

func (x *DNSConfig) GetExcludeSuffixes() []string {
	if x != nil {
		return x.ExcludeSuffixes
	}
	return nil
}

func (x *DNSConfig) GetIncludeSuffixes() []string {
	if x != nil {
		return x.IncludeSuffixes
	}
	return nil
}

func (x *DNSConfig) GetLookupTimeout() int64 {
	if x != nil {
		return x.LookupTimeout
	}
	return 0
}

// OutboundInfo contains all information that the root daemon needs in order to
// establish outbound traffic to the cluster.
type OutboundInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// manager_port is the port for the traffic-manager gRPC API.
	ManagerPort int32 `protobuf:"varint,1,opt,name=manager_port,json=managerPort,proto3" json:"manager_port,omitempty"`
	// session makes it possible for the root daemon to identify itself as the
	// same client as the user daemon.
	Session *manager.SessionInfo `protobuf:"bytes,2,opt,name=session,proto3" json:"session,omitempty"`
	// DNS configuration
	Dns *DNSConfig `protobuf:"bytes,3,opt,name=dns,proto3" json:"dns,omitempty"`
	// also_proxy are user-added subnets.
	AlsoProxySubnets []*manager.IPNet `protobuf:"bytes,5,rep,name=also_proxy_subnets,json=alsoProxySubnets,proto3" json:"also_proxy_subnets,omitempty"`
}

func (x *OutboundInfo) Reset() {
	*x = OutboundInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_daemon_daemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutboundInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutboundInfo) ProtoMessage() {}

func (x *OutboundInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_daemon_daemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutboundInfo.ProtoReflect.Descriptor instead.
func (*OutboundInfo) Descriptor() ([]byte, []int) {
	return file_rpc_daemon_daemon_proto_rawDescGZIP(), []int{3}
}

func (x *OutboundInfo) GetManagerPort() int32 {
	if x != nil {
		return x.ManagerPort
	}
	return 0
}

func (x *OutboundInfo) GetSession() *manager.SessionInfo {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *OutboundInfo) GetDns() *DNSConfig {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *OutboundInfo) GetAlsoProxySubnets() []*manager.IPNet {
	if x != nil {
		return x.AlsoProxySubnets
	}
	return nil
}

var File_rpc_daemon_daemon_proto protoreflect.FileDescriptor

var file_rpc_daemon_daemon_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x99, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3f, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x29, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x64, 0x6e, 0x73, 0x22, 0x30, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x03,
	0x22, 0x04, 0x08, 0x01, 0x10, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x1d, 0x0a, 0x05,
	0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x09,
	0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x49, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49,
	0x70, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0xf1,
	0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d,
	0x6f, 0x6e, 0x2e, 0x44, 0x4e, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x64, 0x6e,
	0x73, 0x12, 0x49, 0x0a, 0x12, 0x61, 0x6c, 0x73, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x50, 0x4e, 0x65, 0x74, 0x52, 0x10, 0x61, 0x6c, 0x73, 0x6f,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x4a, 0x04, 0x08, 0x04,
	0x10, 0x05, 0x32, 0xe0, 0x02, 0x0a, 0x06, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x43, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x51, 0x75, 0x69, 0x74, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x4c, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x46, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x69, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_daemon_daemon_proto_rawDescOnce sync.Once
	file_rpc_daemon_daemon_proto_rawDescData = file_rpc_daemon_daemon_proto_rawDesc
)

func file_rpc_daemon_daemon_proto_rawDescGZIP() []byte {
	file_rpc_daemon_daemon_proto_rawDescOnce.Do(func() {
		file_rpc_daemon_daemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_daemon_daemon_proto_rawDescData)
	})
	return file_rpc_daemon_daemon_proto_rawDescData
}

var file_rpc_daemon_daemon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_daemon_daemon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_daemon_daemon_proto_goTypes = []interface{}{
	(DaemonStatus_ErrType)(0),   // 0: telepresence.daemon.DaemonStatus.ErrType
	(*DaemonStatus)(nil),        // 1: telepresence.daemon.DaemonStatus
	(*Paths)(nil),               // 2: telepresence.daemon.Paths
	(*DNSConfig)(nil),           // 3: telepresence.daemon.DNSConfig
	(*OutboundInfo)(nil),        // 4: telepresence.daemon.OutboundInfo
	(*manager.SessionInfo)(nil), // 5: telepresence.manager.SessionInfo
	(*manager.IPNet)(nil),       // 6: telepresence.manager.IPNet
	(*empty.Empty)(nil),         // 7: google.protobuf.Empty
	(*common.VersionInfo)(nil),  // 8: telepresence.common.VersionInfo
}
var file_rpc_daemon_daemon_proto_depIdxs = []int32{
	0, // 0: telepresence.daemon.DaemonStatus.error:type_name -> telepresence.daemon.DaemonStatus.ErrType
	5, // 1: telepresence.daemon.OutboundInfo.session:type_name -> telepresence.manager.SessionInfo
	3, // 2: telepresence.daemon.OutboundInfo.dns:type_name -> telepresence.daemon.DNSConfig
	6, // 3: telepresence.daemon.OutboundInfo.also_proxy_subnets:type_name -> telepresence.manager.IPNet
	7, // 4: telepresence.daemon.Daemon.Version:input_type -> google.protobuf.Empty
	7, // 5: telepresence.daemon.Daemon.Status:input_type -> google.protobuf.Empty
	7, // 6: telepresence.daemon.Daemon.Quit:input_type -> google.protobuf.Empty
	4, // 7: telepresence.daemon.Daemon.SetOutboundInfo:input_type -> telepresence.daemon.OutboundInfo
	2, // 8: telepresence.daemon.Daemon.SetDnsSearchPath:input_type -> telepresence.daemon.Paths
	8, // 9: telepresence.daemon.Daemon.Version:output_type -> telepresence.common.VersionInfo
	1, // 10: telepresence.daemon.Daemon.Status:output_type -> telepresence.daemon.DaemonStatus
	7, // 11: telepresence.daemon.Daemon.Quit:output_type -> google.protobuf.Empty
	7, // 12: telepresence.daemon.Daemon.SetOutboundInfo:output_type -> google.protobuf.Empty
	7, // 13: telepresence.daemon.Daemon.SetDnsSearchPath:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_daemon_daemon_proto_init() }
func file_rpc_daemon_daemon_proto_init() {
	if File_rpc_daemon_daemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_daemon_daemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DaemonStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_daemon_daemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paths); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_daemon_daemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_daemon_daemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutboundInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_daemon_daemon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_daemon_daemon_proto_goTypes,
		DependencyIndexes: file_rpc_daemon_daemon_proto_depIdxs,
		EnumInfos:         file_rpc_daemon_daemon_proto_enumTypes,
		MessageInfos:      file_rpc_daemon_daemon_proto_msgTypes,
	}.Build()
	File_rpc_daemon_daemon_proto = out.File
	file_rpc_daemon_daemon_proto_rawDesc = nil
	file_rpc_daemon_daemon_proto_goTypes = nil
	file_rpc_daemon_daemon_proto_depIdxs = nil
}
