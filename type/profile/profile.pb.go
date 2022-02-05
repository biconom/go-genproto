// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: profile.proto

package profile_pb

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

type Profile_Data_Gender int32

const (
	Profile_Data_UNKNOWN Profile_Data_Gender = 0
	Profile_Data_MALE    Profile_Data_Gender = 1
	Profile_Data_FEMALE  Profile_Data_Gender = 2
)

// Enum value maps for Profile_Data_Gender.
var (
	Profile_Data_Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	Profile_Data_Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
	}
)

func (x Profile_Data_Gender) Enum() *Profile_Data_Gender {
	p := new(Profile_Data_Gender)
	*p = x
	return p
}

func (x Profile_Data_Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Profile_Data_Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_profile_proto_enumTypes[0].Descriptor()
}

func (Profile_Data_Gender) Type() protoreflect.EnumType {
	return &file_profile_proto_enumTypes[0]
}

func (x Profile_Data_Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Profile_Data_Gender.Descriptor instead.
func (Profile_Data_Gender) EnumDescriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data               *Profile_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`                                                        // профиль
	ModificationLocked bool          `protobuf:"varint,2,opt,name=modification_locked,json=modificationLocked,proto3" json:"modification_locked,omitempty"` // заблокировано для редактирования
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *Profile) GetData() *Profile_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Profile) GetModificationLocked() bool {
	if x != nil {
		return x.ModificationLocked
	}
	return false
}

type Profile_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                            // имя
	Surname    string              `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`                                      // фамилия
	Patronymic string              `protobuf:"bytes,3,opt,name=patronymic,proto3" json:"patronymic,omitempty"`                                // отчество
	Gender     Profile_Data_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=biconom.type.Profile_Data_Gender" json:"gender,omitempty"` // пол
	Birthday   uint32              `protobuf:"varint,5,opt,name=birthday,proto3" json:"birthday,omitempty"`                                   // дата рождения в формате YYYYMMDD
	National   string              `protobuf:"bytes,6,opt,name=national,proto3" json:"national,omitempty"`                                    // гражданство
}

func (x *Profile_Data) Reset() {
	*x = Profile_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile_Data) ProtoMessage() {}

func (x *Profile_Data) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile_Data.ProtoReflect.Descriptor instead.
func (*Profile_Data) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Profile_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile_Data) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Profile_Data) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *Profile_Data) GetGender() Profile_Data_Gender {
	if x != nil {
		return x.Gender
	}
	return Profile_Data_UNKNOWN
}

func (x *Profile_Data) GetBirthday() uint32 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *Profile_Data) GetNational() string {
	if x != nil {
		return x.National
	}
	return ""
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe1, 0x02,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x13, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x1a, 0xf4, 0x01, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79, 0x6d, 0x69,
	0x63, 0x12, 0x39, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_profile_proto_goTypes = []interface{}{
	(Profile_Data_Gender)(0), // 0: biconom.type.Profile.Data.Gender
	(*Profile)(nil),          // 1: biconom.type.Profile
	(*Profile_Data)(nil),     // 2: biconom.type.Profile.Data
}
var file_profile_proto_depIdxs = []int32{
	2, // 0: biconom.type.Profile.data:type_name -> biconom.type.Profile.Data
	0, // 1: biconom.type.Profile.Data.gender:type_name -> biconom.type.Profile.Data.Gender
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile_Data); i {
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
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		EnumInfos:         file_profile_proto_enumTypes,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
