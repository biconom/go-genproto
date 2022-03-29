// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/type/bot_key.proto

package bot_key_pb

import (
	until "github.com/biconom/go-genproto/biconom/type/until"
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

type BotKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotKeyId    uint32       `protobuf:"varint,1,opt,name=bot_key_id,json=botKeyId,proto3" json:"bot_key_id,omitempty"`
	BotKeyTitle string       `protobuf:"bytes,2,opt,name=bot_key_title,json=botKeyTitle,proto3" json:"bot_key_title,omitempty"`
	Token       string       `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Ttl         *until.Until `protobuf:"bytes,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	IpMasks     []string     `protobuf:"bytes,5,rep,name=ip_masks,json=ipMasks,proto3" json:"ip_masks,omitempty"`
}

func (x *BotKey) Reset() {
	*x = BotKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_bot_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotKey) ProtoMessage() {}

func (x *BotKey) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_bot_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotKey.ProtoReflect.Descriptor instead.
func (*BotKey) Descriptor() ([]byte, []int) {
	return file_biconom_type_bot_key_proto_rawDescGZIP(), []int{0}
}

func (x *BotKey) GetBotKeyId() uint32 {
	if x != nil {
		return x.BotKeyId
	}
	return 0
}

func (x *BotKey) GetBotKeyTitle() string {
	if x != nil {
		return x.BotKeyTitle
	}
	return ""
}

func (x *BotKey) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *BotKey) GetTtl() *until.Until {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *BotKey) GetIpMasks() []string {
	if x != nil {
		return x.IpMasks
	}
	return nil
}

type BotKey_ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotKeyId uint32 `protobuf:"varint,1,opt,name=bot_key_id,json=botKeyId,proto3" json:"bot_key_id,omitempty"`
}

func (x *BotKey_ID) Reset() {
	*x = BotKey_ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_bot_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BotKey_ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BotKey_ID) ProtoMessage() {}

func (x *BotKey_ID) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_bot_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BotKey_ID.ProtoReflect.Descriptor instead.
func (*BotKey_ID) Descriptor() ([]byte, []int) {
	return file_biconom_type_bot_key_proto_rawDescGZIP(), []int{0, 0}
}

func (x *BotKey_ID) GetBotKeyId() uint32 {
	if x != nil {
		return x.BotKeyId
	}
	return 0
}

var File_biconom_type_bot_key_proto protoreflect.FileDescriptor

var file_biconom_type_bot_key_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x62,
	0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x18, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x06, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x0a, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x70, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x69, 0x70, 0x4d, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x22, 0x0a, 0x02, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x0a, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x62, 0x6f, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x3b, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_type_bot_key_proto_rawDescOnce sync.Once
	file_biconom_type_bot_key_proto_rawDescData = file_biconom_type_bot_key_proto_rawDesc
)

func file_biconom_type_bot_key_proto_rawDescGZIP() []byte {
	file_biconom_type_bot_key_proto_rawDescOnce.Do(func() {
		file_biconom_type_bot_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_type_bot_key_proto_rawDescData)
	})
	return file_biconom_type_bot_key_proto_rawDescData
}

var file_biconom_type_bot_key_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_biconom_type_bot_key_proto_goTypes = []interface{}{
	(*BotKey)(nil),      // 0: biconom.type.BotKey
	(*BotKey_ID)(nil),   // 1: biconom.type.BotKey.ID
	(*until.Until)(nil), // 2: biconom.type.Until
}
var file_biconom_type_bot_key_proto_depIdxs = []int32{
	2, // 0: biconom.type.BotKey.ttl:type_name -> biconom.type.Until
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_biconom_type_bot_key_proto_init() }
func file_biconom_type_bot_key_proto_init() {
	if File_biconom_type_bot_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_type_bot_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotKey); i {
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
		file_biconom_type_bot_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BotKey_ID); i {
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
			RawDescriptor: file_biconom_type_bot_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_biconom_type_bot_key_proto_goTypes,
		DependencyIndexes: file_biconom_type_bot_key_proto_depIdxs,
		MessageInfos:      file_biconom_type_bot_key_proto_msgTypes,
	}.Build()
	File_biconom_type_bot_key_proto = out.File
	file_biconom_type_bot_key_proto_rawDesc = nil
	file_biconom_type_bot_key_proto_goTypes = nil
	file_biconom_type_bot_key_proto_depIdxs = nil
}