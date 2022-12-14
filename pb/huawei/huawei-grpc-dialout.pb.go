// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: huawei-grpc-dialout.proto

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

type ServiceArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqId int64 `protobuf:"varint,1,opt,name=ReqId,proto3" json:"ReqId,omitempty"` //Request ID.
	// Types that are assignable to MessageData:
	//	*ServiceArgs_Data
	//	*ServiceArgs_DataJson
	MessageData isServiceArgs_MessageData `protobuf_oneof:"MessageData"`
	Errors      string                    `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"` //Error description.
}

func (x *ServiceArgs) Reset() {
	*x = ServiceArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huawei_grpc_dialout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceArgs) ProtoMessage() {}

func (x *ServiceArgs) ProtoReflect() protoreflect.Message {
	mi := &file_huawei_grpc_dialout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceArgs.ProtoReflect.Descriptor instead.
func (*ServiceArgs) Descriptor() ([]byte, []int) {
	return file_huawei_grpc_dialout_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceArgs) GetReqId() int64 {
	if x != nil {
		return x.ReqId
	}
	return 0
}

func (m *ServiceArgs) GetMessageData() isServiceArgs_MessageData {
	if m != nil {
		return m.MessageData
	}
	return nil
}

func (x *ServiceArgs) GetData() []byte {
	if x, ok := x.GetMessageData().(*ServiceArgs_Data); ok {
		return x.Data
	}
	return nil
}

func (x *ServiceArgs) GetDataJson() string {
	if x, ok := x.GetMessageData().(*ServiceArgs_DataJson); ok {
		return x.DataJson
	}
	return ""
}

func (x *ServiceArgs) GetErrors() string {
	if x != nil {
		return x.Errors
	}
	return ""
}

type isServiceArgs_MessageData interface {
	isServiceArgs_MessageData()
}

type ServiceArgs_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"` //Sampled data in GPB encoding format is carried.
}

type ServiceArgs_DataJson struct {
	DataJson string `protobuf:"bytes,4,opt,name=data_json,json=dataJson,proto3,oneof"` //Sampled data in JSON encoding format is carried.
}

func (*ServiceArgs_Data) isServiceArgs_MessageData() {}

func (*ServiceArgs_DataJson) isServiceArgs_MessageData() {}

var File_huawei_grpc_dialout_proto protoreflect.FileDescriptor

var file_huawei_grpc_dialout_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x75, 0x74, 0x22, 0x7f, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65,
	0x71, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x52, 0x65, 0x71, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x0d, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x32, 0x60, 0x0a, 0x0f,
	0x67, 0x52, 0x50, 0x43, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b,
	0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x75, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x1b, 0x2e, 0x68, 0x75,
	0x61, 0x77, 0x65, 0x69, 0x5f, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x75, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x67, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_huawei_grpc_dialout_proto_rawDescOnce sync.Once
	file_huawei_grpc_dialout_proto_rawDescData = file_huawei_grpc_dialout_proto_rawDesc
)

func file_huawei_grpc_dialout_proto_rawDescGZIP() []byte {
	file_huawei_grpc_dialout_proto_rawDescOnce.Do(func() {
		file_huawei_grpc_dialout_proto_rawDescData = protoimpl.X.CompressGZIP(file_huawei_grpc_dialout_proto_rawDescData)
	})
	return file_huawei_grpc_dialout_proto_rawDescData
}

var file_huawei_grpc_dialout_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_huawei_grpc_dialout_proto_goTypes = []interface{}{
	(*ServiceArgs)(nil), // 0: huawei_dialout.serviceArgs
}
var file_huawei_grpc_dialout_proto_depIdxs = []int32{
	0, // 0: huawei_dialout.gRPCDataservice.dataPublish:input_type -> huawei_dialout.serviceArgs
	0, // 1: huawei_dialout.gRPCDataservice.dataPublish:output_type -> huawei_dialout.serviceArgs
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_huawei_grpc_dialout_proto_init() }
func file_huawei_grpc_dialout_proto_init() {
	if File_huawei_grpc_dialout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_huawei_grpc_dialout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceArgs); i {
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
	file_huawei_grpc_dialout_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServiceArgs_Data)(nil),
		(*ServiceArgs_DataJson)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_huawei_grpc_dialout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_huawei_grpc_dialout_proto_goTypes,
		DependencyIndexes: file_huawei_grpc_dialout_proto_depIdxs,
		MessageInfos:      file_huawei_grpc_dialout_proto_msgTypes,
	}.Build()
	File_huawei_grpc_dialout_proto = out.File
	file_huawei_grpc_dialout_proto_rawDesc = nil
	file_huawei_grpc_dialout_proto_goTypes = nil
	file_huawei_grpc_dialout_proto_depIdxs = nil
}
