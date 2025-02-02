// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: integration_insights.proto

package rpc

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

type IntegrationInsightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A log ID to get the integration insight for
	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *IntegrationInsightRequest) Reset() {
	*x = IntegrationInsightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_insights_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationInsightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationInsightRequest) ProtoMessage() {}

func (x *IntegrationInsightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_integration_insights_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationInsightRequest.ProtoReflect.Descriptor instead.
func (*IntegrationInsightRequest) Descriptor() ([]byte, []int) {
	return file_integration_insights_proto_rawDescGZIP(), []int{0}
}

func (x *IntegrationInsightRequest) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

type IntegrationInsightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Integration insight message
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IntegrationInsightResponse) Reset() {
	*x = IntegrationInsightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_integration_insights_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegrationInsightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrationInsightResponse) ProtoMessage() {}

func (x *IntegrationInsightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_integration_insights_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrationInsightResponse.ProtoReflect.Descriptor instead.
func (*IntegrationInsightResponse) Descriptor() ([]byte, []int) {
	return file_integration_insights_proto_rawDescGZIP(), []int{1}
}

func (x *IntegrationInsightResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_integration_insights_proto protoreflect.FileDescriptor

var file_integration_insights_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x73, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70,
	0x63, 0x22, 0x2d, 0x0a, 0x19, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67,
	0x22, 0x36, 0x0a, 0x1a, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x2f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x2d, 0x63, 0x6c, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_integration_insights_proto_rawDescOnce sync.Once
	file_integration_insights_proto_rawDescData = file_integration_insights_proto_rawDesc
)

func file_integration_insights_proto_rawDescGZIP() []byte {
	file_integration_insights_proto_rawDescOnce.Do(func() {
		file_integration_insights_proto_rawDescData = protoimpl.X.CompressGZIP(file_integration_insights_proto_rawDescData)
	})
	return file_integration_insights_proto_rawDescData
}

var file_integration_insights_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_integration_insights_proto_goTypes = []interface{}{
	(*IntegrationInsightRequest)(nil),  // 0: rpc.IntegrationInsightRequest
	(*IntegrationInsightResponse)(nil), // 1: rpc.IntegrationInsightResponse
}
var file_integration_insights_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_integration_insights_proto_init() }
func file_integration_insights_proto_init() {
	if File_integration_insights_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_integration_insights_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationInsightRequest); i {
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
		file_integration_insights_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegrationInsightResponse); i {
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
			RawDescriptor: file_integration_insights_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integration_insights_proto_goTypes,
		DependencyIndexes: file_integration_insights_proto_depIdxs,
		MessageInfos:      file_integration_insights_proto_msgTypes,
	}.Build()
	File_integration_insights_proto = out.File
	file_integration_insights_proto_rawDesc = nil
	file_integration_insights_proto_goTypes = nil
	file_integration_insights_proto_depIdxs = nil
}
