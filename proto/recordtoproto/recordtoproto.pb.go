// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: github.com/google/xtoproto/proto/recordtoproto/recordtoproto.proto

package recordtoproto

import (
	proto "github.com/golang/protobuf/proto"
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

type RecordProtoMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName           string                  `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	MessageName           string                  `protobuf:"bytes,2,opt,name=message_name,json=messageName,proto3" json:"message_name,omitempty"`
	ColumnToFieldMappings []*ColumnToFieldMapping `protobuf:"bytes,3,rep,name=column_to_field_mappings,json=columnToFieldMappings,proto3" json:"column_to_field_mappings,omitempty"`
	GoOptions             *GoOptions              `protobuf:"bytes,4,opt,name=go_options,json=goOptions,proto3" json:"go_options,omitempty"`
	ExtraFieldDefinitions []*FieldDefinition      `protobuf:"bytes,5,rep,name=extra_field_definitions,json=extraFieldDefinitions,proto3" json:"extra_field_definitions,omitempty"`
}

func (x *RecordProtoMapping) Reset() {
	*x = RecordProtoMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordProtoMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordProtoMapping) ProtoMessage() {}

func (x *RecordProtoMapping) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordProtoMapping.ProtoReflect.Descriptor instead.
func (*RecordProtoMapping) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{0}
}

func (x *RecordProtoMapping) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *RecordProtoMapping) GetMessageName() string {
	if x != nil {
		return x.MessageName
	}
	return ""
}

func (x *RecordProtoMapping) GetColumnToFieldMappings() []*ColumnToFieldMapping {
	if x != nil {
		return x.ColumnToFieldMappings
	}
	return nil
}

func (x *RecordProtoMapping) GetGoOptions() *GoOptions {
	if x != nil {
		return x.GoOptions
	}
	return nil
}

func (x *RecordProtoMapping) GetExtraFieldDefinitions() []*FieldDefinition {
	if x != nil {
		return x.ExtraFieldDefinitions
	}
	return nil
}

type ColumnToFieldMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnIndex  int32    `protobuf:"varint,1,opt,name=column_index,json=columnIndex,proto3" json:"column_index,omitempty"`
	ColName      string   `protobuf:"bytes,2,opt,name=col_name,json=colName,proto3" json:"col_name,omitempty"`
	ProtoName    string   `protobuf:"bytes,3,opt,name=proto_name,json=protoName,proto3" json:"proto_name,omitempty"`
	ProtoType    string   `protobuf:"bytes,4,opt,name=proto_type,json=protoType,proto3" json:"proto_type,omitempty"`
	ProtoTag     int32    `protobuf:"varint,5,opt,name=proto_tag,json=protoTag,proto3" json:"proto_tag,omitempty"`
	Ignored      bool     `protobuf:"varint,6,opt,name=ignored,proto3" json:"ignored,omitempty"`
	ProtoImports []string `protobuf:"bytes,7,rep,name=proto_imports,json=protoImports,proto3" json:"proto_imports,omitempty"`
	Comment      string   `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"`
	// Types that are assignable to ParsingInfo:
	//	*ColumnToFieldMapping_TimeFormat
	//	*ColumnToFieldMapping_DurationFormat
	ParsingInfo isColumnToFieldMapping_ParsingInfo `protobuf_oneof:"parsing_info"`
}

func (x *ColumnToFieldMapping) Reset() {
	*x = ColumnToFieldMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnToFieldMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnToFieldMapping) ProtoMessage() {}

func (x *ColumnToFieldMapping) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnToFieldMapping.ProtoReflect.Descriptor instead.
func (*ColumnToFieldMapping) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{1}
}

func (x *ColumnToFieldMapping) GetColumnIndex() int32 {
	if x != nil {
		return x.ColumnIndex
	}
	return 0
}

func (x *ColumnToFieldMapping) GetColName() string {
	if x != nil {
		return x.ColName
	}
	return ""
}

func (x *ColumnToFieldMapping) GetProtoName() string {
	if x != nil {
		return x.ProtoName
	}
	return ""
}

func (x *ColumnToFieldMapping) GetProtoType() string {
	if x != nil {
		return x.ProtoType
	}
	return ""
}

func (x *ColumnToFieldMapping) GetProtoTag() int32 {
	if x != nil {
		return x.ProtoTag
	}
	return 0
}

func (x *ColumnToFieldMapping) GetIgnored() bool {
	if x != nil {
		return x.Ignored
	}
	return false
}

func (x *ColumnToFieldMapping) GetProtoImports() []string {
	if x != nil {
		return x.ProtoImports
	}
	return nil
}

func (x *ColumnToFieldMapping) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (m *ColumnToFieldMapping) GetParsingInfo() isColumnToFieldMapping_ParsingInfo {
	if m != nil {
		return m.ParsingInfo
	}
	return nil
}

func (x *ColumnToFieldMapping) GetTimeFormat() *TimeFormat {
	if x, ok := x.GetParsingInfo().(*ColumnToFieldMapping_TimeFormat); ok {
		return x.TimeFormat
	}
	return nil
}

func (x *ColumnToFieldMapping) GetDurationFormat() *DurationFormat {
	if x, ok := x.GetParsingInfo().(*ColumnToFieldMapping_DurationFormat); ok {
		return x.DurationFormat
	}
	return nil
}

type isColumnToFieldMapping_ParsingInfo interface {
	isColumnToFieldMapping_ParsingInfo()
}

type ColumnToFieldMapping_TimeFormat struct {
	TimeFormat *TimeFormat `protobuf:"bytes,8,opt,name=time_format,json=timeFormat,proto3,oneof"`
}

type ColumnToFieldMapping_DurationFormat struct {
	DurationFormat *DurationFormat `protobuf:"bytes,10,opt,name=duration_format,json=durationFormat,proto3,oneof"`
}

func (*ColumnToFieldMapping_TimeFormat) isColumnToFieldMapping_ParsingInfo() {}

func (*ColumnToFieldMapping_DurationFormat) isColumnToFieldMapping_ParsingInfo() {}

type FieldDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoName    string   `protobuf:"bytes,1,opt,name=proto_name,json=protoName,proto3" json:"proto_name,omitempty"`
	ProtoType    string   `protobuf:"bytes,2,opt,name=proto_type,json=protoType,proto3" json:"proto_type,omitempty"`
	ProtoTag     int32    `protobuf:"varint,3,opt,name=proto_tag,json=protoTag,proto3" json:"proto_tag,omitempty"`
	ProtoImports []string `protobuf:"bytes,4,rep,name=proto_imports,json=protoImports,proto3" json:"proto_imports,omitempty"`
	Comment      string   `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *FieldDefinition) Reset() {
	*x = FieldDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldDefinition) ProtoMessage() {}

func (x *FieldDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldDefinition.ProtoReflect.Descriptor instead.
func (*FieldDefinition) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{2}
}

func (x *FieldDefinition) GetProtoName() string {
	if x != nil {
		return x.ProtoName
	}
	return ""
}

func (x *FieldDefinition) GetProtoType() string {
	if x != nil {
		return x.ProtoType
	}
	return ""
}

func (x *FieldDefinition) GetProtoTag() int32 {
	if x != nil {
		return x.ProtoTag
	}
	return 0
}

func (x *FieldDefinition) GetProtoImports() []string {
	if x != nil {
		return x.ProtoImports
	}
	return nil
}

func (x *FieldDefinition) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type TimeFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoLayout     string `protobuf:"bytes,1,opt,name=go_layout,json=goLayout,proto3" json:"go_layout,omitempty"`
	TimeZoneName string `protobuf:"bytes,2,opt,name=time_zone_name,json=timeZoneName,proto3" json:"time_zone_name,omitempty"`
}

func (x *TimeFormat) Reset() {
	*x = TimeFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFormat) ProtoMessage() {}

func (x *TimeFormat) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFormat.ProtoReflect.Descriptor instead.
func (*TimeFormat) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{3}
}

func (x *TimeFormat) GetGoLayout() string {
	if x != nil {
		return x.GoLayout
	}
	return ""
}

func (x *TimeFormat) GetTimeZoneName() string {
	if x != nil {
		return x.TimeZoneName
	}
	return ""
}

type DurationFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoUnitSuffix string `protobuf:"bytes,1,opt,name=go_unit_suffix,json=goUnitSuffix,proto3" json:"go_unit_suffix,omitempty"`
}

func (x *DurationFormat) Reset() {
	*x = DurationFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DurationFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DurationFormat) ProtoMessage() {}

func (x *DurationFormat) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DurationFormat.ProtoReflect.Descriptor instead.
func (*DurationFormat) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{4}
}

func (x *DurationFormat) GetGoUnitSuffix() string {
	if x != nil {
		return x.GoUnitSuffix
	}
	return ""
}

type GoOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoPackageName string `protobuf:"bytes,1,opt,name=go_package_name,json=goPackageName,proto3" json:"go_package_name,omitempty"`
	ProtoImport   string `protobuf:"bytes,2,opt,name=proto_import,json=protoImport,proto3" json:"proto_import,omitempty"`
}

func (x *GoOptions) Reset() {
	*x = GoOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoOptions) ProtoMessage() {}

func (x *GoOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoOptions.ProtoReflect.Descriptor instead.
func (*GoOptions) Descriptor() ([]byte, []int) {
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP(), []int{5}
}

func (x *GoOptions) GetGoPackageName() string {
	if x != nil {
		return x.GoPackageName
	}
	return ""
}

func (x *GoOptions) GetProtoImport() string {
	if x != nil {
		return x.ProtoImport
	}
	return ""
}

var File_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto protoreflect.FileDescriptor

var file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x78, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x78, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba,
	0x02, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x18, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x78, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54,
	0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x15, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x78, 0x74, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x09, 0x67,
	0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x78, 0x74, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x65, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x96, 0x03, 0x0a, 0x14,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x6f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x61, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x78, 0x74,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x48, 0x00, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x43, 0x0a, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x78, 0x74, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x54, 0x61, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x67, 0x6f, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x6f, 0x5f, 0x75, 0x6e, 0x69, 0x74,
	0x5f, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67,
	0x6f, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x56, 0x0a, 0x09, 0x47,
	0x6f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x67, 0x6f, 0x5f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x67, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x78, 0x74, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescOnce sync.Once
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescData = file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDesc
)

func file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescGZIP() []byte {
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescOnce.Do(func() {
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescData)
	})
	return file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDescData
}

var file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_goTypes = []interface{}{
	(*RecordProtoMapping)(nil),   // 0: xtoproto.RecordProtoMapping
	(*ColumnToFieldMapping)(nil), // 1: xtoproto.ColumnToFieldMapping
	(*FieldDefinition)(nil),      // 2: xtoproto.FieldDefinition
	(*TimeFormat)(nil),           // 3: xtoproto.TimeFormat
	(*DurationFormat)(nil),       // 4: xtoproto.DurationFormat
	(*GoOptions)(nil),            // 5: xtoproto.GoOptions
}
var file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_depIdxs = []int32{
	1, // 0: xtoproto.RecordProtoMapping.column_to_field_mappings:type_name -> xtoproto.ColumnToFieldMapping
	5, // 1: xtoproto.RecordProtoMapping.go_options:type_name -> xtoproto.GoOptions
	2, // 2: xtoproto.RecordProtoMapping.extra_field_definitions:type_name -> xtoproto.FieldDefinition
	3, // 3: xtoproto.ColumnToFieldMapping.time_format:type_name -> xtoproto.TimeFormat
	4, // 4: xtoproto.ColumnToFieldMapping.duration_format:type_name -> xtoproto.DurationFormat
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_init() }
func file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_init() {
	if File_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordProtoMapping); i {
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
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnToFieldMapping); i {
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
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldDefinition); i {
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
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFormat); i {
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
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DurationFormat); i {
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
		file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoOptions); i {
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
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ColumnToFieldMapping_TimeFormat)(nil),
		(*ColumnToFieldMapping_DurationFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_goTypes,
		DependencyIndexes: file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_depIdxs,
		MessageInfos:      file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_msgTypes,
	}.Build()
	File_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto = out.File
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_rawDesc = nil
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_goTypes = nil
	file_github_com_google_xtoproto_proto_recordtoproto_recordtoproto_proto_depIdxs = nil
}
