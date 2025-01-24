// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: monitor.proto

package monitor

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

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_monitor_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{0}
}

type OrganizationId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgID         string                 `protobuf:"bytes,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationId) Reset() {
	*x = OrganizationId{}
	mi := &file_monitor_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationId) ProtoMessage() {}

func (x *OrganizationId) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationId.ProtoReflect.Descriptor instead.
func (*OrganizationId) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationId) GetOrgID() string {
	if x != nil {
		return x.OrgID
	}
	return ""
}

type MetricsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgID         *OrganizationId        `protobuf:"bytes,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricsResponse) Reset() {
	*x = MetricsResponse{}
	mi := &file_monitor_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsResponse) ProtoMessage() {}

func (x *MetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsResponse.ProtoReflect.Descriptor instead.
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *MetricsResponse) GetOrgID() *OrganizationId {
	if x != nil {
		return x.OrgID
	}
	return nil
}

type MetricsQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrgID         *OrganizationId        `protobuf:"bytes,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricsQuery) Reset() {
	*x = MetricsQuery{}
	mi := &file_monitor_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsQuery) ProtoMessage() {}

func (x *MetricsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_monitor_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsQuery.ProtoReflect.Descriptor instead.
func (*MetricsQuery) Descriptor() ([]byte, []int) {
	return file_monitor_proto_rawDescGZIP(), []int{3}
}

func (x *MetricsQuery) GetOrgID() *OrganizationId {
	if x != nil {
		return x.OrgID
	}
	return nil
}

var File_monitor_proto protoreflect.FileDescriptor

var file_monitor_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x26, 0x0a, 0x0e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x0f, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05,
	0x4f, 0x72, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x52, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x22, 0x3d, 0x0a, 0x0c, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x4f,
	0x72, 0x67, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x52, 0x05, 0x4f, 0x72, 0x67, 0x49, 0x44, 0x32, 0x8a, 0x01, 0x0a, 0x0e, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x18, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0e, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x70, 0x72, 0x6f, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitor_proto_rawDescOnce sync.Once
	file_monitor_proto_rawDescData = file_monitor_proto_rawDesc
)

func file_monitor_proto_rawDescGZIP() []byte {
	file_monitor_proto_rawDescOnce.Do(func() {
		file_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitor_proto_rawDescData)
	})
	return file_monitor_proto_rawDescData
}

var file_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_monitor_proto_goTypes = []any{
	(*Empty)(nil),           // 0: monitor.Empty
	(*OrganizationId)(nil),  // 1: monitor.OrganizationId
	(*MetricsResponse)(nil), // 2: monitor.MetricsResponse
	(*MetricsQuery)(nil),    // 3: monitor.MetricsQuery
}
var file_monitor_proto_depIdxs = []int32{
	1, // 0: monitor.MetricsResponse.OrgID:type_name -> monitor.OrganizationId
	1, // 1: monitor.MetricsQuery.OrgID:type_name -> monitor.OrganizationId
	3, // 2: monitor.MonitorService.GetMetrics:input_type -> monitor.MetricsQuery
	0, // 3: monitor.MonitorService.GetAllMetrics:input_type -> monitor.Empty
	2, // 4: monitor.MonitorService.GetMetrics:output_type -> monitor.MetricsResponse
	2, // 5: monitor.MonitorService.GetAllMetrics:output_type -> monitor.MetricsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_monitor_proto_init() }
func file_monitor_proto_init() {
	if File_monitor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitor_proto_goTypes,
		DependencyIndexes: file_monitor_proto_depIdxs,
		MessageInfos:      file_monitor_proto_msgTypes,
	}.Build()
	File_monitor_proto = out.File
	file_monitor_proto_rawDesc = nil
	file_monitor_proto_goTypes = nil
	file_monitor_proto_depIdxs = nil
}