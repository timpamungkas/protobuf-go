// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/basic/user.proto

package basic

import (
	date "google.golang.org/genproto/googleapis/type/date"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Gender int32

const (
	Gender_GENDER_UNSPECIFIED Gender = 0
	Gender_GENDER_FEMALE      Gender = 1
	Gender_GENDER_MALE        Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNSPECIFIED",
		1: "GENDER_FEMALE",
		2: "GENDER_MALE",
	}
	Gender_value = map[string]int32{
		"GENDER_UNSPECIFIED": 0,
		"GENDER_FEMALE":      1,
		"GENDER_MALE":        2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_basic_user_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_proto_basic_user_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{0}
}

// Represents user
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IsActive             bool       `protobuf:"varint,3,opt,name=is_active,proto3" json:"is_active,omitempty"`
	Password             []byte     `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Emails               []string   `protobuf:"bytes,16,rep,name=emails,proto3" json:"emails,omitempty"`
	Gender               Gender     `protobuf:"varint,17,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`
	Addresss             *Addresss  `protobuf:"bytes,18,opt,name=addresss,proto3" json:"addresss,omitempty"`
	CommunicationChannel *anypb.Any `protobuf:"bytes,19,opt,name=communication_channel,proto3" json:"communication_channel,omitempty"`
	// Types that are assignable to ElectronicCommChannel:
	//
	//	*User_SocialMedia
	//	*User_InstantMessaging
	ElectronicCommChannel isUser_ElectronicCommChannel `protobuf_oneof:"electronic_comm_channel"`
	SkillRating           map[string]uint32            `protobuf:"bytes,22,rep,name=skill_rating,proto3" json:"skill_rating,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LastLoginTimestamp    *timestamppb.Timestamp       `protobuf:"bytes,23,opt,name=last_login_timestamp,proto3" json:"last_login_timestamp,omitempty"`
	BirthDate             *date.Date                   `protobuf:"bytes,24,opt,name=birth_date,proto3" json:"birth_date,omitempty"`
	LastKnownLocation     *latlng.LatLng               `protobuf:"bytes,25,opt,name=last_known_location,proto3" json:"last_known_location,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *User) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *User) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_UNSPECIFIED
}

func (x *User) GetAddresss() *Addresss {
	if x != nil {
		return x.Addresss
	}
	return nil
}

func (x *User) GetCommunicationChannel() *anypb.Any {
	if x != nil {
		return x.CommunicationChannel
	}
	return nil
}

func (m *User) GetElectronicCommChannel() isUser_ElectronicCommChannel {
	if m != nil {
		return m.ElectronicCommChannel
	}
	return nil
}

func (x *User) GetSocialMedia() *SocialMedia {
	if x, ok := x.GetElectronicCommChannel().(*User_SocialMedia); ok {
		return x.SocialMedia
	}
	return nil
}

func (x *User) GetInstantMessaging() *InstantMessaging {
	if x, ok := x.GetElectronicCommChannel().(*User_InstantMessaging); ok {
		return x.InstantMessaging
	}
	return nil
}

func (x *User) GetSkillRating() map[string]uint32 {
	if x != nil {
		return x.SkillRating
	}
	return nil
}

func (x *User) GetLastLoginTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LastLoginTimestamp
	}
	return nil
}

func (x *User) GetBirthDate() *date.Date {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *User) GetLastKnownLocation() *latlng.LatLng {
	if x != nil {
		return x.LastKnownLocation
	}
	return nil
}

type isUser_ElectronicCommChannel interface {
	isUser_ElectronicCommChannel()
}

type User_SocialMedia struct {
	SocialMedia *SocialMedia `protobuf:"bytes,20,opt,name=social_media,proto3,oneof"`
}

type User_InstantMessaging struct {
	InstantMessaging *InstantMessaging `protobuf:"bytes,21,opt,name=instant_messaging,proto3,oneof"`
}

func (*User_SocialMedia) isUser_ElectronicCommChannel() {}

func (*User_InstantMessaging) isUser_ElectronicCommChannel() {}

type Addresss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street     string               `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City       string               `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Country    string               `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Coordinate *Addresss_Coordinate `protobuf:"bytes,16,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *Addresss) Reset() {
	*x = Addresss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addresss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addresss) ProtoMessage() {}

func (x *Addresss) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addresss.ProtoReflect.Descriptor instead.
func (*Addresss) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{1}
}

func (x *Addresss) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Addresss) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Addresss) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Addresss) GetCoordinate() *Addresss_Coordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type PaperMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaperMailAddress string `protobuf:"bytes,1,opt,name=paper_mail_address,proto3" json:"paper_mail_address,omitempty"`
}

func (x *PaperMail) Reset() {
	*x = PaperMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaperMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaperMail) ProtoMessage() {}

func (x *PaperMail) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaperMail.ProtoReflect.Descriptor instead.
func (*PaperMail) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{2}
}

func (x *PaperMail) GetPaperMailAddress() string {
	if x != nil {
		return x.PaperMailAddress
	}
	return ""
}

type SocialMedia struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SocialMediaPlatform string `protobuf:"bytes,1,opt,name=social_media_platform,proto3" json:"social_media_platform,omitempty"`
	SocialMediaUsername string `protobuf:"bytes,2,opt,name=social_media_username,proto3" json:"social_media_username,omitempty"`
}

func (x *SocialMedia) Reset() {
	*x = SocialMedia{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocialMedia) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocialMedia) ProtoMessage() {}

func (x *SocialMedia) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocialMedia.ProtoReflect.Descriptor instead.
func (*SocialMedia) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{3}
}

func (x *SocialMedia) GetSocialMediaPlatform() string {
	if x != nil {
		return x.SocialMediaPlatform
	}
	return ""
}

func (x *SocialMedia) GetSocialMediaUsername() string {
	if x != nil {
		return x.SocialMediaUsername
	}
	return ""
}

type InstantMessaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstantMessagingProduct  string `protobuf:"bytes,1,opt,name=instant_messaging_product,proto3" json:"instant_messaging_product,omitempty"`
	InstantMessagingUsername string `protobuf:"bytes,2,opt,name=instant_messaging_username,proto3" json:"instant_messaging_username,omitempty"`
}

func (x *InstantMessaging) Reset() {
	*x = InstantMessaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantMessaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantMessaging) ProtoMessage() {}

func (x *InstantMessaging) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantMessaging.ProtoReflect.Descriptor instead.
func (*InstantMessaging) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{4}
}

func (x *InstantMessaging) GetInstantMessagingProduct() string {
	if x != nil {
		return x.InstantMessagingProduct
	}
	return ""
}

func (x *InstantMessaging) GetInstantMessagingUsername() string {
	if x != nil {
		return x.InstantMessagingUsername
	}
	return ""
}

type Addresss_Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Addresss_Coordinate) Reset() {
	*x = Addresss_Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_basic_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Addresss_Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Addresss_Coordinate) ProtoMessage() {}

func (x *Addresss_Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_basic_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Addresss_Coordinate.ProtoReflect.Descriptor instead.
func (*Addresss_Coordinate) Descriptor() ([]byte, []int) {
	return file_proto_basic_user_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Addresss_Coordinate) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Addresss_Coordinate) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_proto_basic_user_proto protoreflect.FileDescriptor

var file_proto_basic_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf0, 0x05, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x08, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x73, 0x12, 0x4a, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x32, 0x0a,
	0x0c, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x12, 0x41, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x16, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x4e, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x31, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x45, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c,
	0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3e, 0x0a, 0x10, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x19, 0x0a, 0x17, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xce, 0x01, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x1a, 0x46,
	0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x3b, 0x0a, 0x09, 0x50, 0x61, 0x70, 0x65, 0x72, 0x4d,
	0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x79, 0x0a, 0x0b, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x90,
	0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x3e, 0x0a, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x2a, 0x44, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_basic_user_proto_rawDescOnce sync.Once
	file_proto_basic_user_proto_rawDescData = file_proto_basic_user_proto_rawDesc
)

func file_proto_basic_user_proto_rawDescGZIP() []byte {
	file_proto_basic_user_proto_rawDescOnce.Do(func() {
		file_proto_basic_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_basic_user_proto_rawDescData)
	})
	return file_proto_basic_user_proto_rawDescData
}

var file_proto_basic_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_basic_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_basic_user_proto_goTypes = []interface{}{
	(Gender)(0),                   // 0: Gender
	(*User)(nil),                  // 1: User
	(*Addresss)(nil),              // 2: Addresss
	(*PaperMail)(nil),             // 3: PaperMail
	(*SocialMedia)(nil),           // 4: SocialMedia
	(*InstantMessaging)(nil),      // 5: InstantMessaging
	nil,                           // 6: User.SkillRatingEntry
	(*Addresss_Coordinate)(nil),   // 7: Addresss.Coordinate
	(*anypb.Any)(nil),             // 8: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*date.Date)(nil),             // 10: google.type.Date
	(*latlng.LatLng)(nil),         // 11: google.type.LatLng
}
var file_proto_basic_user_proto_depIdxs = []int32{
	0,  // 0: User.gender:type_name -> Gender
	2,  // 1: User.addresss:type_name -> Addresss
	8,  // 2: User.communication_channel:type_name -> google.protobuf.Any
	4,  // 3: User.social_media:type_name -> SocialMedia
	5,  // 4: User.instant_messaging:type_name -> InstantMessaging
	6,  // 5: User.skill_rating:type_name -> User.SkillRatingEntry
	9,  // 6: User.last_login_timestamp:type_name -> google.protobuf.Timestamp
	10, // 7: User.birth_date:type_name -> google.type.Date
	11, // 8: User.last_known_location:type_name -> google.type.LatLng
	7,  // 9: Addresss.coordinate:type_name -> Addresss.Coordinate
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_basic_user_proto_init() }
func file_proto_basic_user_proto_init() {
	if File_proto_basic_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_basic_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_basic_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addresss); i {
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
		file_proto_basic_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaperMail); i {
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
		file_proto_basic_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SocialMedia); i {
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
		file_proto_basic_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantMessaging); i {
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
		file_proto_basic_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Addresss_Coordinate); i {
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
	file_proto_basic_user_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*User_SocialMedia)(nil),
		(*User_InstantMessaging)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_basic_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_basic_user_proto_goTypes,
		DependencyIndexes: file_proto_basic_user_proto_depIdxs,
		EnumInfos:         file_proto_basic_user_proto_enumTypes,
		MessageInfos:      file_proto_basic_user_proto_msgTypes,
	}.Build()
	File_proto_basic_user_proto = out.File
	file_proto_basic_user_proto_rawDesc = nil
	file_proto_basic_user_proto_goTypes = nil
	file_proto_basic_user_proto_depIdxs = nil
}
