// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protos/cacheping.proto

package client_sdk_go

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

type XPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *XPingRequest) Reset() {
	*x = XPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XPingRequest) ProtoMessage() {}

func (x *XPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XPingRequest.ProtoReflect.Descriptor instead.
func (*XPingRequest) Descriptor() ([]byte, []int) {
	return file_protos_cacheping_proto_rawDescGZIP(), []int{0}
}

type XPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *XPingResponse) Reset() {
	*x = XPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_cacheping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XPingResponse) ProtoMessage() {}

func (x *XPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_cacheping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XPingResponse.ProtoReflect.Descriptor instead.
func (*XPingResponse) Descriptor() ([]byte, []int) {
	return file_protos_cacheping_proto_rawDescGZIP(), []int{1}
}

var File_protos_cacheping_proto protoreflect.FileDescriptor

var file_protos_cacheping_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x49, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x41, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x62, 0x0a, 0x11, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x68, 0x71, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x3b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x64, 0x6b, 0x5f, 0x67, 0x6f, 0xaa, 0x02, 0x18, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_cacheping_proto_rawDescOnce sync.Once
	file_protos_cacheping_proto_rawDescData = file_protos_cacheping_proto_rawDesc
)

func file_protos_cacheping_proto_rawDescGZIP() []byte {
	file_protos_cacheping_proto_rawDescOnce.Do(func() {
		file_protos_cacheping_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_cacheping_proto_rawDescData)
	})
	return file_protos_cacheping_proto_rawDescData
}

var file_protos_cacheping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_cacheping_proto_goTypes = []interface{}{
	(*XPingRequest)(nil),  // 0: cache_client._PingRequest
	(*XPingResponse)(nil), // 1: cache_client._PingResponse
}
var file_protos_cacheping_proto_depIdxs = []int32{
	0, // 0: cache_client.Ping.Ping:input_type -> cache_client._PingRequest
	1, // 1: cache_client.Ping.Ping:output_type -> cache_client._PingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_cacheping_proto_init() }
func file_protos_cacheping_proto_init() {
	if File_protos_cacheping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_cacheping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XPingRequest); i {
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
		file_protos_cacheping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XPingResponse); i {
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
			RawDescriptor: file_protos_cacheping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_cacheping_proto_goTypes,
		DependencyIndexes: file_protos_cacheping_proto_depIdxs,
		MessageInfos:      file_protos_cacheping_proto_msgTypes,
	}.Build()
	File_protos_cacheping_proto = out.File
	file_protos_cacheping_proto_rawDesc = nil
	file_protos_cacheping_proto_goTypes = nil
	file_protos_cacheping_proto_depIdxs = nil
}
