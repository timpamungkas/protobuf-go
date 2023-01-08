// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/dummy/application.proto

package dummy

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

// Job application
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationId     uint32 `protobuf:"varint,1,opt,name=application_id,proto3" json:"application_id,omitempty"`
	ApplicantFullName string `protobuf:"bytes,2,opt,name=applicant_full_name,proto3" json:"applicant_full_name,omitempty"`
	Phone             string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email             string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dummy_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dummy_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_proto_dummy_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetApplicationId() uint32 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *Application) GetApplicantFullName() string {
	if x != nil {
		return x.ApplicantFullName
	}
	return ""
}

func (x *Application) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Application) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_proto_dummy_application_proto protoreflect.FileDescriptor

var file_proto_dummy_application_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x74, 0x68, 0x65, 0x2e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x5f,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dummy_application_proto_rawDescOnce sync.Once
	file_proto_dummy_application_proto_rawDescData = file_proto_dummy_application_proto_rawDesc
)

func file_proto_dummy_application_proto_rawDescGZIP() []byte {
	file_proto_dummy_application_proto_rawDescOnce.Do(func() {
		file_proto_dummy_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dummy_application_proto_rawDescData)
	})
	return file_proto_dummy_application_proto_rawDescData
}

var file_proto_dummy_application_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_dummy_application_proto_goTypes = []interface{}{
	(*Application)(nil), // 0: the.dummy.Application
}
var file_proto_dummy_application_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_dummy_application_proto_init() }
func file_proto_dummy_application_proto_init() {
	if File_proto_dummy_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dummy_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
			RawDescriptor: file_proto_dummy_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_dummy_application_proto_goTypes,
		DependencyIndexes: file_proto_dummy_application_proto_depIdxs,
		MessageInfos:      file_proto_dummy_application_proto_msgTypes,
	}.Build()
	File_proto_dummy_application_proto = out.File
	file_proto_dummy_application_proto_rawDesc = nil
	file_proto_dummy_application_proto_goTypes = nil
	file_proto_dummy_application_proto_depIdxs = nil
}
