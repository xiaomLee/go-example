// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: user.proto

package api

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// gender enum.
type User_Gender int32

const (
	User_GENDER_UNKNOWN    User_Gender = 0
	User_GENDER_MALE       User_Gender = 1
	User_GENDER_FEMALE     User_Gender = 2
	User_GENDER_GENDERLESS User_Gender = 3
)

// Enum value maps for User_Gender.
var (
	User_Gender_name = map[int32]string{
		0: "GENDER_UNKNOWN",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
		3: "GENDER_GENDERLESS",
	}
	User_Gender_value = map[string]int32{
		"GENDER_UNKNOWN":    0,
		"GENDER_MALE":       1,
		"GENDER_FEMALE":     2,
		"GENDER_GENDERLESS": 3,
	}
)

func (x User_Gender) Enum() *User_Gender {
	p := new(User_Gender)
	*p = x
	return p
}

func (x User_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (User_Gender) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x User_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Gender.Descriptor instead.
func (User_Gender) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0, 0}
}

// user status enum.
type User_Status int32

const (
	User_STATUS_UNKNOWN User_Status = 0
	User_STATUS_NORMAL  User_Status = 1
	User_STATUS_FROZEN  User_Status = 2
	User_STATUS_DELETED User_Status = 3
)

// Enum value maps for User_Status.
var (
	User_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "STATUS_NORMAL",
		2: "STATUS_FROZEN",
		3: "STATUS_DELETED",
	}
	User_Status_value = map[string]int32{
		"STATUS_UNKNOWN": 0,
		"STATUS_NORMAL":  1,
		"STATUS_FROZEN":  2,
		"STATUS_DELETED": 3,
	}
)

func (x User_Status) Enum() *User_Status {
	p := new(User_Status)
	*p = x
	return p
}

func (x User_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (User_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[1].Descriptor()
}

func (User_Status) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[1]
}

func (x User_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use User_Status.Descriptor instead.
func (User_Status) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0, 1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// gender.
	Gender User_Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=service.user.User_Gender" json:"gender,omitempty"`
	// birthday yy-mm-dd.
	Birthday string `protobuf:"bytes,4,opt,name=birthday,proto3" json:"birthday,omitempty"`
	// user status.
	Status User_Status `protobuf:"varint,5,opt,name=status,proto3,enum=service.user.User_Status" json:"status,omitempty"`
	// password.
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// account.
	Account string `protobuf:"bytes,7,opt,name=account,proto3" json:"account,omitempty"`
	// email.
	Email string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	// tel.
	Tel string `protobuf:"bytes,9,opt,name=tel,proto3" json:"tel,omitempty"`
	// create time.
	CreateTime int64 `protobuf:"varint,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// update time.
	UpdateTime int64 `protobuf:"varint,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
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
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetGender() User_Gender {
	if x != nil {
		return x.Gender
	}
	return User_GENDER_UNKNOWN
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetStatus() User_Status {
	if x != nil {
		return x.Status
	}
	return User_STATUS_UNKNOWN
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x0e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46,
	0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x45, 0x4e, 0x44, 0x45,
	0x52, 0x5f, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x03, 0x22, 0x56,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e,
	0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_proto_goTypes = []interface{}{
	(User_Gender)(0), // 0: service.user.User.Gender
	(User_Status)(0), // 1: service.user.User.Status
	(*User)(nil),     // 2: service.user.User
}
var file_user_proto_depIdxs = []int32{
	0, // 0: service.user.User.gender:type_name -> service.user.User.Gender
	1, // 1: service.user.User.status:type_name -> service.user.User.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
