// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/type/device.proto

package device_pb

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

type Device_Type_Id int32

const (
	Device_Type_UNKNOWN Device_Type_Id = 0
	Device_Type_BOT     Device_Type_Id = 1
	Device_Type_WEB     Device_Type_Id = 2
	Device_Type_PHONE   Device_Type_Id = 3
	Device_Type_TABLET  Device_Type_Id = 4
	Device_Type_DESKTOP Device_Type_Id = 5
)

// Enum value maps for Device_Type_Id.
var (
	Device_Type_Id_name = map[int32]string{
		0: "UNKNOWN",
		1: "BOT",
		2: "WEB",
		3: "PHONE",
		4: "TABLET",
		5: "DESKTOP",
	}
	Device_Type_Id_value = map[string]int32{
		"UNKNOWN": 0,
		"BOT":     1,
		"WEB":     2,
		"PHONE":   3,
		"TABLET":  4,
		"DESKTOP": 5,
	}
)

func (x Device_Type_Id) Enum() *Device_Type_Id {
	p := new(Device_Type_Id)
	*p = x
	return p
}

func (x Device_Type_Id) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_Type_Id) Descriptor() protoreflect.EnumDescriptor {
	return file_biconom_type_device_proto_enumTypes[0].Descriptor()
}

func (Device_Type_Id) Type() protoreflect.EnumType {
	return &file_biconom_type_device_proto_enumTypes[0]
}

func (x Device_Type_Id) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_Type_Id.Descriptor instead.
func (Device_Type_Id) EnumDescriptor() ([]byte, []int) {
	return file_biconom_type_device_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Device_OperatingSystem_Id int32

const (
	Device_OperatingSystem_UNKNOWN Device_OperatingSystem_Id = 0
	Device_OperatingSystem_IOS     Device_OperatingSystem_Id = 1
	Device_OperatingSystem_MACOS   Device_OperatingSystem_Id = 2
	Device_OperatingSystem_ANDROID Device_OperatingSystem_Id = 3
	Device_OperatingSystem_WINDOWS Device_OperatingSystem_Id = 4
	Device_OperatingSystem_LINUX   Device_OperatingSystem_Id = 5
)

// Enum value maps for Device_OperatingSystem_Id.
var (
	Device_OperatingSystem_Id_name = map[int32]string{
		0: "UNKNOWN",
		1: "IOS",
		2: "MACOS",
		3: "ANDROID",
		4: "WINDOWS",
		5: "LINUX",
	}
	Device_OperatingSystem_Id_value = map[string]int32{
		"UNKNOWN": 0,
		"IOS":     1,
		"MACOS":   2,
		"ANDROID": 3,
		"WINDOWS": 4,
		"LINUX":   5,
	}
)

func (x Device_OperatingSystem_Id) Enum() *Device_OperatingSystem_Id {
	p := new(Device_OperatingSystem_Id)
	*p = x
	return p
}

func (x Device_OperatingSystem_Id) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Device_OperatingSystem_Id) Descriptor() protoreflect.EnumDescriptor {
	return file_biconom_type_device_proto_enumTypes[1].Descriptor()
}

func (Device_OperatingSystem_Id) Type() protoreflect.EnumType {
	return &file_biconom_type_device_proto_enumTypes[1]
}

func (x Device_OperatingSystem_Id) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Device_OperatingSystem_Id.Descriptor instead.
func (Device_OperatingSystem_Id) EnumDescriptor() ([]byte, []int) {
	return file_biconom_type_device_proto_rawDescGZIP(), []int{0, 1, 0}
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceType      *Device_Type            `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	OperatingSystem *Device_OperatingSystem `protobuf:"bytes,4,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_biconom_type_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetDeviceType() *Device_Type {
	if x != nil {
		return x.DeviceType
	}
	return nil
}

func (x *Device) GetOperatingSystem() *Device_OperatingSystem {
	if x != nil {
		return x.OperatingSystem
	}
	return nil
}

type Device_Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      Device_Type_Id `protobuf:"varint,1,opt,name=id,proto3,enum=biconom.type.Device_Type_Id" json:"id,omitempty"`
	Version string         `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Device_Type) Reset() {
	*x = Device_Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device_Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device_Type) ProtoMessage() {}

func (x *Device_Type) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device_Type.ProtoReflect.Descriptor instead.
func (*Device_Type) Descriptor() ([]byte, []int) {
	return file_biconom_type_device_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Device_Type) GetId() Device_Type_Id {
	if x != nil {
		return x.Id
	}
	return Device_Type_UNKNOWN
}

func (x *Device_Type) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Device_OperatingSystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      Device_OperatingSystem_Id `protobuf:"varint,1,opt,name=id,proto3,enum=biconom.type.Device_OperatingSystem_Id" json:"id,omitempty"`
	Version string                    `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Device_OperatingSystem) Reset() {
	*x = Device_OperatingSystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device_OperatingSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device_OperatingSystem) ProtoMessage() {}

func (x *Device_OperatingSystem) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device_OperatingSystem.ProtoReflect.Descriptor instead.
func (*Device_OperatingSystem) Descriptor() ([]byte, []int) {
	return file_biconom_type_device_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Device_OperatingSystem) GetId() Device_OperatingSystem_Id {
	if x != nil {
		return x.Id
	}
	return Device_OperatingSystem_UNKNOWN
}

func (x *Device_OperatingSystem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_biconom_type_device_proto protoreflect.FileDescriptor

var file_biconom_type_device_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe2, 0x03, 0x0a, 0x06, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x4f, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x1a, 0x97, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x4f, 0x54, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x54, 0x10, 0x04, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x53, 0x4b, 0x54, 0x4f, 0x50, 0x10, 0x05, 0x1a, 0xb0, 0x01, 0x0a, 0x0f,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x37, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x4d, 0x41, 0x43, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x4e, 0x44,
	0x52, 0x4f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57,
	0x53, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x05, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_type_device_proto_rawDescOnce sync.Once
	file_biconom_type_device_proto_rawDescData = file_biconom_type_device_proto_rawDesc
)

func file_biconom_type_device_proto_rawDescGZIP() []byte {
	file_biconom_type_device_proto_rawDescOnce.Do(func() {
		file_biconom_type_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_type_device_proto_rawDescData)
	})
	return file_biconom_type_device_proto_rawDescData
}

var file_biconom_type_device_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_biconom_type_device_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_biconom_type_device_proto_goTypes = []interface{}{
	(Device_Type_Id)(0),            // 0: biconom.type.Device.Type.Id
	(Device_OperatingSystem_Id)(0), // 1: biconom.type.Device.OperatingSystem.Id
	(*Device)(nil),                 // 2: biconom.type.Device
	(*Device_Type)(nil),            // 3: biconom.type.Device.Type
	(*Device_OperatingSystem)(nil), // 4: biconom.type.Device.OperatingSystem
}
var file_biconom_type_device_proto_depIdxs = []int32{
	3, // 0: biconom.type.Device.device_type:type_name -> biconom.type.Device.Type
	4, // 1: biconom.type.Device.operating_system:type_name -> biconom.type.Device.OperatingSystem
	0, // 2: biconom.type.Device.Type.id:type_name -> biconom.type.Device.Type.Id
	1, // 3: biconom.type.Device.OperatingSystem.id:type_name -> biconom.type.Device.OperatingSystem.Id
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_biconom_type_device_proto_init() }
func file_biconom_type_device_proto_init() {
	if File_biconom_type_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_type_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_biconom_type_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device_Type); i {
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
		file_biconom_type_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device_OperatingSystem); i {
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
			RawDescriptor: file_biconom_type_device_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_biconom_type_device_proto_goTypes,
		DependencyIndexes: file_biconom_type_device_proto_depIdxs,
		EnumInfos:         file_biconom_type_device_proto_enumTypes,
		MessageInfos:      file_biconom_type_device_proto_msgTypes,
	}.Build()
	File_biconom_type_device_proto = out.File
	file_biconom_type_device_proto_rawDesc = nil
	file_biconom_type_device_proto_goTypes = nil
	file_biconom_type_device_proto_depIdxs = nil
}
