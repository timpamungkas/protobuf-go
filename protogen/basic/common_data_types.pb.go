// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/basic/common_data_types.proto

package basic

import (
	date "google.golang.org/genproto/googleapis/type/date"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AdditionalDataTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyDuration   *durationpb.Duration   `protobuf:"bytes,1,opt,name=my_duration,json=myDuration,proto3" json:"my_duration,omitempty"`
	MyTimestamp  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=my_timestamp,json=myTimestamp,proto3" json:"my_timestamp,omitempty"`
	MyDate       *date.Date             `protobuf:"bytes,16,opt,name=my_date,json=myDate,proto3" json:"my_date,omitempty"`
	MyCoordinate *latlng.LatLng         `protobuf:"bytes,17,opt,name=my_coordinate,json=myCoordinate,proto3" json:"my_coordinate,omitempty"`
}

func (x *AdditionalDataTypes) Reset() {
	*x = AdditionalDataTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_common_data_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionalDataTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionalDataTypes) ProtoMessage() {}

func (x *AdditionalDataTypes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_common_data_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionalDataTypes.ProtoReflect.Descriptor instead.
func (*AdditionalDataTypes) Descriptor() ([]byte, []int) {
	return file_proto_basic_common_data_types_proto_rawDescGZIP(), []int{0}
}

func (x *AdditionalDataTypes) GetMyDuration() *durationpb.Duration {
	if x != nil {
		return x.MyDuration
	}
	return nil
}

func (x *AdditionalDataTypes) GetMyTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.MyTimestamp
	}
	return nil
}

func (x *AdditionalDataTypes) GetMyDate() *date.Date {
	if x != nil {
		return x.MyDate
	}
	return nil
}

func (x *AdditionalDataTypes) GetMyCoordinate() *latlng.LatLng {
	if x != nil {
		return x.MyCoordinate
	}
	return nil
}

var File_proto_basic_common_data_types_proto protoreflect.FileDescriptor

var file_proto_basic_common_data_types_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0b,
	0x6d, 0x79, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6d, 0x79,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x79, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6d, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x06, 0x6d, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6d, 0x79, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52,
	0x0c, 0x6d, 0x79, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x42, 0x1c, 0x5a,
	0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_common_data_types_proto_rawDescOnce sync.Once
	file_proto_basic_common_data_types_proto_rawDescData = file_proto_basic_common_data_types_proto_rawDesc
)

func file_proto_basic_common_data_types_proto_rawDescGZIP() []byte {
	file_proto_basic_common_data_types_proto_rawDescOnce.Do(func() {
		file_proto_basic_common_data_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_common_data_types_proto_rawDescData)
	})
	return file_proto_basic_common_data_types_proto_rawDescData
}

var file_proto_basic_common_data_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_basic_common_data_types_proto_goTypes = []interface{}{
	(*AdditionalDataTypes)(nil),   // 0: AdditionalDataTypes
	(*durationpb.Duration)(nil),   // 1: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*date.Date)(nil),             // 3: google.type.Date
	(*latlng.LatLng)(nil),         // 4: google.type.LatLng
}
var file_proto_basic_common_data_types_proto_depIdxs = []int32{
	1, // 0: AdditionalDataTypes.my_duration:type_name -> google.protobuf.Duration
	2, // 1: AdditionalDataTypes.my_timestamp:type_name -> google.protobuf.Timestamp
	3, // 2: AdditionalDataTypes.my_date:type_name -> google.type.Date
	4, // 3: AdditionalDataTypes.my_coordinate:type_name -> google.type.LatLng
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_basic_common_data_types_proto_init() }
func file_proto_basic_common_data_types_proto_init() {
	if File_proto_basic_common_data_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_basic_common_data_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionalDataTypes); i {
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
			RawDescriptor: file_proto_basic_common_data_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_common_data_types_proto_goTypes,
		DependencyIndexes: file_proto_basic_common_data_types_proto_depIdxs,
		MessageInfos:      file_proto_basic_common_data_types_proto_msgTypes,
	}.Build()
	File_proto_basic_common_data_types_proto = out.File
	file_proto_basic_common_data_types_proto_rawDesc = nil
	file_proto_basic_common_data_types_proto_goTypes = nil
	file_proto_basic_common_data_types_proto_depIdxs = nil
}
