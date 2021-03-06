// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/type/position.proto

package currency_pb

import (
	currency "github.com/biconom/go-genproto/biconom/type/currency"
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

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Position_Value       `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Tree  *currency.Tree_Header `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_position_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_position_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_biconom_type_position_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetValue() *Position_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Position) GetTree() *currency.Tree_Header {
	if x != nil {
		return x.Tree
	}
	return nil
}

type Position_ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UniqueField:
	//	*Position_ID_PositionId
	//	*Position_ID_PositionRef
	UniqueField isPosition_ID_UniqueField `protobuf_oneof:"unique_field"`
}

func (x *Position_ID) Reset() {
	*x = Position_ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_position_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position_ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position_ID) ProtoMessage() {}

func (x *Position_ID) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_position_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position_ID.ProtoReflect.Descriptor instead.
func (*Position_ID) Descriptor() ([]byte, []int) {
	return file_biconom_type_position_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Position_ID) GetUniqueField() isPosition_ID_UniqueField {
	if m != nil {
		return m.UniqueField
	}
	return nil
}

func (x *Position_ID) GetPositionId() uint32 {
	if x, ok := x.GetUniqueField().(*Position_ID_PositionId); ok {
		return x.PositionId
	}
	return 0
}

func (x *Position_ID) GetPositionRef() *Position_ID_Ref {
	if x, ok := x.GetUniqueField().(*Position_ID_PositionRef); ok {
		return x.PositionRef
	}
	return nil
}

type isPosition_ID_UniqueField interface {
	isPosition_ID_UniqueField()
}

type Position_ID_PositionId struct {
	PositionId uint32 `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3,oneof"`
}

type Position_ID_PositionRef struct {
	PositionRef *Position_ID_Ref `protobuf:"bytes,2,opt,name=position_ref,json=positionRef,proto3,oneof"`
}

func (*Position_ID_PositionId) isPosition_ID_UniqueField() {}

func (*Position_ID_PositionRef) isPosition_ID_UniqueField() {}

type Position_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionId      uint32 `protobuf:"varint,1,opt,name=position_id,json=positionId,proto3" json:"position_id,omitempty"`
	PositionRefId   uint32 `protobuf:"varint,2,opt,name=position_ref_id,json=positionRefId,proto3" json:"position_ref_id,omitempty"`
	PositionRefLine uint32 `protobuf:"varint,3,opt,name=position_ref_line,json=positionRefLine,proto3" json:"position_ref_line,omitempty"`
	PositionWidth   uint32 `protobuf:"varint,4,opt,name=position_width,json=positionWidth,proto3" json:"position_width,omitempty"`
	PositionLevel   uint32 `protobuf:"varint,5,opt,name=position_level,json=positionLevel,proto3" json:"position_level,omitempty"`
	AccountId       uint32 `protobuf:"varint,6,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	CreatedAt       uint32 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Position_Value) Reset() {
	*x = Position_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_position_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position_Value) ProtoMessage() {}

func (x *Position_Value) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_position_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position_Value.ProtoReflect.Descriptor instead.
func (*Position_Value) Descriptor() ([]byte, []int) {
	return file_biconom_type_position_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Position_Value) GetPositionId() uint32 {
	if x != nil {
		return x.PositionId
	}
	return 0
}

func (x *Position_Value) GetPositionRefId() uint32 {
	if x != nil {
		return x.PositionRefId
	}
	return 0
}

func (x *Position_Value) GetPositionRefLine() uint32 {
	if x != nil {
		return x.PositionRefLine
	}
	return 0
}

func (x *Position_Value) GetPositionWidth() uint32 {
	if x != nil {
		return x.PositionWidth
	}
	return 0
}

func (x *Position_Value) GetPositionLevel() uint32 {
	if x != nil {
		return x.PositionLevel
	}
	return 0
}

func (x *Position_Value) GetAccountId() uint32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Position_Value) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Position_ID_Ref struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PositionRefId   uint32 `protobuf:"varint,1,opt,name=position_ref_id,json=positionRefId,proto3" json:"position_ref_id,omitempty"`
	PositionRefLine uint32 `protobuf:"varint,2,opt,name=position_ref_line,json=positionRefLine,proto3" json:"position_ref_line,omitempty"`
}

func (x *Position_ID_Ref) Reset() {
	*x = Position_ID_Ref{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_position_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position_ID_Ref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position_ID_Ref) ProtoMessage() {}

func (x *Position_ID_Ref) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_position_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position_ID_Ref.ProtoReflect.Descriptor instead.
func (*Position_ID_Ref) Descriptor() ([]byte, []int) {
	return file_biconom_type_position_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Position_ID_Ref) GetPositionRefId() uint32 {
	if x != nil {
		return x.PositionRefId
	}
	return 0
}

func (x *Position_ID_Ref) GetPositionRefLine() uint32 {
	if x != nil {
		return x.PositionRefLine
	}
	return 0
}

var File_biconom_type_position_proto protoreflect.FileDescriptor

var file_biconom_type_position_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x17, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x04, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x04,
	0x74, 0x72, 0x65, 0x65, 0x1a, 0xd6, 0x01, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x42,
	0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x2e,
	0x52, 0x65, 0x66, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x1a, 0x59, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x66, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x6e, 0x65, 0x42, 0x0e, 0x0a,
	0x0c, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x88, 0x02,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69,
	0x64, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x3b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_type_position_proto_rawDescOnce sync.Once
	file_biconom_type_position_proto_rawDescData = file_biconom_type_position_proto_rawDesc
)

func file_biconom_type_position_proto_rawDescGZIP() []byte {
	file_biconom_type_position_proto_rawDescOnce.Do(func() {
		file_biconom_type_position_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_type_position_proto_rawDescData)
	})
	return file_biconom_type_position_proto_rawDescData
}

var file_biconom_type_position_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_biconom_type_position_proto_goTypes = []interface{}{
	(*Position)(nil),             // 0: biconom.type.Position
	(*Position_ID)(nil),          // 1: biconom.type.Position.ID
	(*Position_Value)(nil),       // 2: biconom.type.Position.Value
	(*Position_ID_Ref)(nil),      // 3: biconom.type.Position.ID.Ref
	(*currency.Tree_Header)(nil), // 4: biconom.type.Tree.Header
}
var file_biconom_type_position_proto_depIdxs = []int32{
	2, // 0: biconom.type.Position.value:type_name -> biconom.type.Position.Value
	4, // 1: biconom.type.Position.tree:type_name -> biconom.type.Tree.Header
	3, // 2: biconom.type.Position.ID.position_ref:type_name -> biconom.type.Position.ID.Ref
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_biconom_type_position_proto_init() }
func file_biconom_type_position_proto_init() {
	if File_biconom_type_position_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_type_position_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_biconom_type_position_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position_ID); i {
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
		file_biconom_type_position_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position_Value); i {
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
		file_biconom_type_position_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position_ID_Ref); i {
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
	file_biconom_type_position_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Position_ID_PositionId)(nil),
		(*Position_ID_PositionRef)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_biconom_type_position_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_biconom_type_position_proto_goTypes,
		DependencyIndexes: file_biconom_type_position_proto_depIdxs,
		MessageInfos:      file_biconom_type_position_proto_msgTypes,
	}.Build()
	File_biconom_type_position_proto = out.File
	file_biconom_type_position_proto_rawDesc = nil
	file_biconom_type_position_proto_goTypes = nil
	file_biconom_type_position_proto_depIdxs = nil
}
