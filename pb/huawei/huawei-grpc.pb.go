// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: huawei-grpc.proto

package huawei

import (
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

type Grpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service     *Grpc_Service     `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceIpv6 *Grpc_ServiceIpv6 `protobuf:"bytes,2,opt,name=service_ipv6,json=service-ipv6,proto3" json:"service_ipv6,omitempty"`
	Client      *Grpc_Client      `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	SessionCar  *Grpc_SessionCar  `protobuf:"bytes,4,opt,name=session_car,json=session-car,proto3" json:"session_car,omitempty"`
}

func (x *Grpc) Reset() {
	*x = Grpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc) ProtoMessage() {}

func (x *Grpc) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc.ProtoReflect.Descriptor instead.
func (*Grpc) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Grpc) GetService() *Grpc_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Grpc) GetServiceIpv6() *Grpc_ServiceIpv6 {
	if x != nil {
		return x.ServiceIpv6
	}
	return nil
}

func (x *Grpc) GetClient() *Grpc_Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Grpc) GetSessionCar() *Grpc_SessionCar {
	if x != nil {
		return x.SessionCar
	}
	return nil
}

type Grpc_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerEnable    bool   `protobuf:"varint,1,opt,name=server_enable,json=server-enable,proto3" json:"server_enable,omitempty"`
	SourceAddress   string `protobuf:"bytes,2,opt,name=source_address,json=source-address,proto3" json:"source_address,omitempty"`
	VpnName         string `protobuf:"bytes,3,opt,name=vpn_name,json=vpn-name,proto3" json:"vpn_name,omitempty"`
	ServerPort      uint32 `protobuf:"varint,4,opt,name=server_port,json=server-port,proto3" json:"server_port,omitempty"`
	IdleTimeout     uint32 `protobuf:"varint,5,opt,name=idle_timeout,json=idle-timeout,proto3" json:"idle_timeout,omitempty"`
	AclNumberOrName string `protobuf:"bytes,6,opt,name=acl_number_or_name,json=acl-number-or-name,proto3" json:"acl_number_or_name,omitempty"`
	SslPolicy       string `protobuf:"bytes,7,opt,name=ssl_policy,json=ssl-policy,proto3" json:"ssl_policy,omitempty"`
	SslPolicyPeer   bool   `protobuf:"varint,8,opt,name=ssl_policy_peer,json=ssl-policy-peer,proto3" json:"ssl_policy_peer,omitempty"`
	Dscp            uint32 `protobuf:"varint,9,opt,name=dscp,proto3" json:"dscp,omitempty"`
	Compression     string `protobuf:"bytes,10,opt,name=compression,proto3" json:"compression,omitempty"`
}

func (x *Grpc_Service) Reset() {
	*x = Grpc_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc_Service) ProtoMessage() {}

func (x *Grpc_Service) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc_Service.ProtoReflect.Descriptor instead.
func (*Grpc_Service) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Grpc_Service) GetServerEnable() bool {
	if x != nil {
		return x.ServerEnable
	}
	return false
}

func (x *Grpc_Service) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *Grpc_Service) GetVpnName() string {
	if x != nil {
		return x.VpnName
	}
	return ""
}

func (x *Grpc_Service) GetServerPort() uint32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *Grpc_Service) GetIdleTimeout() uint32 {
	if x != nil {
		return x.IdleTimeout
	}
	return 0
}

func (x *Grpc_Service) GetAclNumberOrName() string {
	if x != nil {
		return x.AclNumberOrName
	}
	return ""
}

func (x *Grpc_Service) GetSslPolicy() string {
	if x != nil {
		return x.SslPolicy
	}
	return ""
}

func (x *Grpc_Service) GetSslPolicyPeer() bool {
	if x != nil {
		return x.SslPolicyPeer
	}
	return false
}

func (x *Grpc_Service) GetDscp() uint32 {
	if x != nil {
		return x.Dscp
	}
	return 0
}

func (x *Grpc_Service) GetCompression() string {
	if x != nil {
		return x.Compression
	}
	return ""
}

type Grpc_ServiceIpv6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerEnable    bool   `protobuf:"varint,1,opt,name=server_enable,json=server-enable,proto3" json:"server_enable,omitempty"`
	SourceAddress   string `protobuf:"bytes,2,opt,name=source_address,json=source-address,proto3" json:"source_address,omitempty"`
	VpnName         string `protobuf:"bytes,3,opt,name=vpn_name,json=vpn-name,proto3" json:"vpn_name,omitempty"`
	ServerPort      uint32 `protobuf:"varint,4,opt,name=server_port,json=server-port,proto3" json:"server_port,omitempty"`
	IdleTimeout     uint32 `protobuf:"varint,5,opt,name=idle_timeout,json=idle-timeout,proto3" json:"idle_timeout,omitempty"`
	AclNumberOrName string `protobuf:"bytes,6,opt,name=acl_number_or_name,json=acl-number-or-name,proto3" json:"acl_number_or_name,omitempty"`
	SslPolicy       string `protobuf:"bytes,7,opt,name=ssl_policy,json=ssl-policy,proto3" json:"ssl_policy,omitempty"`
	SslPolicyPeer   bool   `protobuf:"varint,8,opt,name=ssl_policy_peer,json=ssl-policy-peer,proto3" json:"ssl_policy_peer,omitempty"`
	Dscp            uint32 `protobuf:"varint,9,opt,name=dscp,proto3" json:"dscp,omitempty"`
	Compression     string `protobuf:"bytes,10,opt,name=compression,proto3" json:"compression,omitempty"`
}

func (x *Grpc_ServiceIpv6) Reset() {
	*x = Grpc_ServiceIpv6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc_ServiceIpv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc_ServiceIpv6) ProtoMessage() {}

func (x *Grpc_ServiceIpv6) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc_ServiceIpv6.ProtoReflect.Descriptor instead.
func (*Grpc_ServiceIpv6) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Grpc_ServiceIpv6) GetServerEnable() bool {
	if x != nil {
		return x.ServerEnable
	}
	return false
}

func (x *Grpc_ServiceIpv6) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *Grpc_ServiceIpv6) GetVpnName() string {
	if x != nil {
		return x.VpnName
	}
	return ""
}

func (x *Grpc_ServiceIpv6) GetServerPort() uint32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *Grpc_ServiceIpv6) GetIdleTimeout() uint32 {
	if x != nil {
		return x.IdleTimeout
	}
	return 0
}

func (x *Grpc_ServiceIpv6) GetAclNumberOrName() string {
	if x != nil {
		return x.AclNumberOrName
	}
	return ""
}

func (x *Grpc_ServiceIpv6) GetSslPolicy() string {
	if x != nil {
		return x.SslPolicy
	}
	return ""
}

func (x *Grpc_ServiceIpv6) GetSslPolicyPeer() bool {
	if x != nil {
		return x.SslPolicyPeer
	}
	return false
}

func (x *Grpc_ServiceIpv6) GetDscp() uint32 {
	if x != nil {
		return x.Dscp
	}
	return 0
}

func (x *Grpc_ServiceIpv6) GetCompression() string {
	if x != nil {
		return x.Compression
	}
	return ""
}

type Grpc_Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslPolicy       string `protobuf:"bytes,1,opt,name=ssl_policy,json=ssl-policy,proto3" json:"ssl_policy,omitempty"`
	SslVerifyCnName string `protobuf:"bytes,2,opt,name=ssl_verify_cn_name,json=ssl-verify-cn-name,proto3" json:"ssl_verify_cn_name,omitempty"`
	SslPolicyPeer   bool   `protobuf:"varint,3,opt,name=ssl_policy_peer,json=ssl-policy-peer,proto3" json:"ssl_policy_peer,omitempty"`
	Dscp            uint32 `protobuf:"varint,4,opt,name=dscp,proto3" json:"dscp,omitempty"`
	SslVerifySan    string `protobuf:"bytes,5,opt,name=ssl_verify_san,json=ssl-verify-san,proto3" json:"ssl_verify_san,omitempty"`
}

func (x *Grpc_Client) Reset() {
	*x = Grpc_Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc_Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc_Client) ProtoMessage() {}

func (x *Grpc_Client) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc_Client.ProtoReflect.Descriptor instead.
func (*Grpc_Client) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Grpc_Client) GetSslPolicy() string {
	if x != nil {
		return x.SslPolicy
	}
	return ""
}

func (x *Grpc_Client) GetSslVerifyCnName() string {
	if x != nil {
		return x.SslVerifyCnName
	}
	return ""
}

func (x *Grpc_Client) GetSslPolicyPeer() bool {
	if x != nil {
		return x.SslPolicyPeer
	}
	return false
}

func (x *Grpc_Client) GetDscp() uint32 {
	if x != nil {
		return x.Dscp
	}
	return 0
}

func (x *Grpc_Client) GetSslVerifySan() string {
	if x != nil {
		return x.SslVerifySan
	}
	return ""
}

type Grpc_SessionCar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable      bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	CirInterval uint32 `protobuf:"varint,2,opt,name=cir_interval,json=cir-interval,proto3" json:"cir_interval,omitempty"`
	CbsInterval uint32 `protobuf:"varint,3,opt,name=cbs_interval,json=cbs-interval,proto3" json:"cbs_interval,omitempty"`
	PirInterval uint32 `protobuf:"varint,4,opt,name=pir_interval,json=pir-interval,proto3" json:"pir_interval,omitempty"`
	PbsInterval uint32 `protobuf:"varint,5,opt,name=pbs_interval,json=pbs-interval,proto3" json:"pbs_interval,omitempty"`
}

func (x *Grpc_SessionCar) Reset() {
	*x = Grpc_SessionCar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grpc_SessionCar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grpc_SessionCar) ProtoMessage() {}

func (x *Grpc_SessionCar) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grpc_SessionCar.ProtoReflect.Descriptor instead.
func (*Grpc_SessionCar) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Grpc_SessionCar) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Grpc_SessionCar) GetCirInterval() uint32 {
	if x != nil {
		return x.CirInterval
	}
	return 0
}

func (x *Grpc_SessionCar) GetCbsInterval() uint32 {
	if x != nil {
		return x.CbsInterval
	}
	return 0
}

func (x *Grpc_SessionCar) GetPirInterval() uint32 {
	if x != nil {
		return x.PirInterval
	}
	return 0
}

func (x *Grpc_SessionCar) GetPbsInterval() uint32 {
	if x != nil {
		return x.PbsInterval
	}
	return 0
}

var File_huawei_grpc_proto protoreflect.FileDescriptor

var file_huawei_grpc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x22, 0xc4, 0x0a, 0x0a, 0x04, 0x47, 0x72, 0x70, 0x63, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x70, 0x76, 0x36, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x70, 0x76, 0x36, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x70, 0x76,
	0x36, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65,
	0x69, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x72, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d,
	0x63, 0x61, 0x72, 0x1a, 0xe9, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x70, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x70, 0x6e, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x2e, 0x0a, 0x12, 0x61, 0x63, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x63, 0x6c,
	0x2d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x6f, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x73, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x70, 0x65,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2d, 0x70, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x73, 0x63,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x73, 0x63, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0xed, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x70, 0x76, 0x36, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x70, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x76, 0x70, 0x6e, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69,
	0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x2e, 0x0a, 0x12, 0x61, 0x63, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x63, 0x6c,
	0x2d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x6f, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x73, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x73, 0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x70, 0x65,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2d, 0x70, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x73, 0x63,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x64, 0x73, 0x63, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0xbe, 0x01, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x73,
	0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x73,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x73, 0x6c, 0x2d, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x2d, 0x63, 0x6e, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x73,
	0x6c, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x73, 0x6c, 0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2d,
	0x70, 0x65, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x73, 0x63, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x64, 0x73, 0x63, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x73, 0x6c, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x73, 0x6c, 0x2d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x73, 0x61, 0x6e,
	0x1a, 0xb4, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x69, 0x72, 0x5f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63,
	0x69, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x62, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x63, 0x62, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x69, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x69, 0x72, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x62, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x62, 0x73, 0x2d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_huawei_grpc_proto_rawDescOnce sync.Once
	file_huawei_grpc_proto_rawDescData = file_huawei_grpc_proto_rawDesc
)

func file_huawei_grpc_proto_rawDescGZIP() []byte {
	file_huawei_grpc_proto_rawDescOnce.Do(func() {
		file_huawei_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_huawei_grpc_proto_rawDescData)
	})
	return file_huawei_grpc_proto_rawDescData
}

var file_huawei_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_huawei_grpc_proto_goTypes = []interface{}{
	(*Grpc)(nil),             // 0: huawei_grpc.Grpc
	(*Grpc_Service)(nil),     // 1: huawei_grpc.Grpc.Service
	(*Grpc_ServiceIpv6)(nil), // 2: huawei_grpc.Grpc.ServiceIpv6
	(*Grpc_Client)(nil),      // 3: huawei_grpc.Grpc.Client
	(*Grpc_SessionCar)(nil),  // 4: huawei_grpc.Grpc.SessionCar
}
var file_huawei_grpc_proto_depIdxs = []int32{
	1, // 0: huawei_grpc.Grpc.service:type_name -> huawei_grpc.Grpc.Service
	2, // 1: huawei_grpc.Grpc.service_ipv6:type_name -> huawei_grpc.Grpc.ServiceIpv6
	3, // 2: huawei_grpc.Grpc.client:type_name -> huawei_grpc.Grpc.Client
	4, // 3: huawei_grpc.Grpc.session_car:type_name -> huawei_grpc.Grpc.SessionCar
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_huawei_grpc_proto_init() }
func file_huawei_grpc_proto_init() {
	if File_huawei_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_huawei_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grpc); i {
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
		file_huawei_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grpc_Service); i {
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
		file_huawei_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grpc_ServiceIpv6); i {
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
		file_huawei_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grpc_Client); i {
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
		file_huawei_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grpc_SessionCar); i {
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
			RawDescriptor: file_huawei_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_huawei_grpc_proto_goTypes,
		DependencyIndexes: file_huawei_grpc_proto_depIdxs,
		MessageInfos:      file_huawei_grpc_proto_msgTypes,
	}.Build()
	File_huawei_grpc_proto = out.File
	file_huawei_grpc_proto_rawDesc = nil
	file_huawei_grpc_proto_goTypes = nil
	file_huawei_grpc_proto_depIdxs = nil
}
