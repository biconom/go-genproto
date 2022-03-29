// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/rank/v1/rank_system.proto

package service_admin_rank_pb

import (
	rank_system "github.com/biconom/go-genproto/biconom/type/rank_system"
	sort "github.com/biconom/go-genproto/biconom/type/sort"
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

type RankSystemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step                    *rank_system.RankSystem_ID `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	Sort                    *sort.Sort                 `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedRankSystem bool                       `protobuf:"varint,3,opt,name=only_published_rank_system,json=onlyPublishedRankSystem,proto3" json:"only_published_rank_system,omitempty"`
}

func (x *RankSystemListRequest) Reset() {
	*x = RankSystemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSystemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSystemListRequest) ProtoMessage() {}

func (x *RankSystemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSystemListRequest.ProtoReflect.Descriptor instead.
func (*RankSystemListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_system_proto_rawDescGZIP(), []int{0}
}

func (x *RankSystemListRequest) GetStep() *rank_system.RankSystem_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankSystemListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *RankSystemListRequest) GetOnlyPublishedRankSystem() bool {
	if x != nil {
		return x.OnlyPublishedRankSystem
	}
	return false
}

type RankSystemCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankSystemName string `protobuf:"bytes,1,opt,name=rank_system_name,json=rankSystemName,proto3" json:"rank_system_name,omitempty"`
	Published      bool   `protobuf:"varint,2,opt,name=published,proto3" json:"published,omitempty"`
}

func (x *RankSystemCreateRequest) Reset() {
	*x = RankSystemCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSystemCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSystemCreateRequest) ProtoMessage() {}

func (x *RankSystemCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSystemCreateRequest.ProtoReflect.Descriptor instead.
func (*RankSystemCreateRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_system_proto_rawDescGZIP(), []int{1}
}

func (x *RankSystemCreateRequest) GetRankSystemName() string {
	if x != nil {
		return x.RankSystemName
	}
	return ""
}

func (x *RankSystemCreateRequest) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

var File_biconom_admin_rank_v1_rank_system_proto protoreflect.FileDescriptor

var file_biconom_admin_rank_v1_rank_system_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x15, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x1a,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x72,
	0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x61, 0x0a, 0x17, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x32, 0xbc, 0x04, 0x0a,
	0x0a, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x3e, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a,
	0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x4b, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0a,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x54, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x09, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x61, 0x6e, 0x6b,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_biconom_admin_rank_v1_rank_system_proto_rawDescOnce sync.Once
	file_biconom_admin_rank_v1_rank_system_proto_rawDescData = file_biconom_admin_rank_v1_rank_system_proto_rawDesc
)

func file_biconom_admin_rank_v1_rank_system_proto_rawDescGZIP() []byte {
	file_biconom_admin_rank_v1_rank_system_proto_rawDescOnce.Do(func() {
		file_biconom_admin_rank_v1_rank_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_rank_v1_rank_system_proto_rawDescData)
	})
	return file_biconom_admin_rank_v1_rank_system_proto_rawDescData
}

var file_biconom_admin_rank_v1_rank_system_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_biconom_admin_rank_v1_rank_system_proto_goTypes = []interface{}{
	(*RankSystemListRequest)(nil),         // 0: biconom.admin.rank.v1.RankSystemListRequest
	(*RankSystemCreateRequest)(nil),       // 1: biconom.admin.rank.v1.RankSystemCreateRequest
	(*rank_system.RankSystem_ID)(nil),     // 2: biconom.type.RankSystem.ID
	(*sort.Sort)(nil),                     // 3: biconom.type.Sort
	(*rank_system.RankSystem)(nil),        // 4: biconom.type.RankSystem
	(*rank_system.RankSystem_Header)(nil), // 5: biconom.type.RankSystem.Header
}
var file_biconom_admin_rank_v1_rank_system_proto_depIdxs = []int32{
	2, // 0: biconom.admin.rank.v1.RankSystemListRequest.step:type_name -> biconom.type.RankSystem.ID
	3, // 1: biconom.admin.rank.v1.RankSystemListRequest.sort:type_name -> biconom.type.Sort
	2, // 2: biconom.admin.rank.v1.RankSystem.Get:input_type -> biconom.type.RankSystem.ID
	0, // 3: biconom.admin.rank.v1.RankSystem.List:input_type -> biconom.admin.rank.v1.RankSystemListRequest
	2, // 4: biconom.admin.rank.v1.RankSystem.HeaderGet:input_type -> biconom.type.RankSystem.ID
	0, // 5: biconom.admin.rank.v1.RankSystem.HeaderList:input_type -> biconom.admin.rank.v1.RankSystemListRequest
	1, // 6: biconom.admin.rank.v1.RankSystem.Create:input_type -> biconom.admin.rank.v1.RankSystemCreateRequest
	2, // 7: biconom.admin.rank.v1.RankSystem.Publish:input_type -> biconom.type.RankSystem.ID
	2, // 8: biconom.admin.rank.v1.RankSystem.Unpublish:input_type -> biconom.type.RankSystem.ID
	4, // 9: biconom.admin.rank.v1.RankSystem.Get:output_type -> biconom.type.RankSystem
	4, // 10: biconom.admin.rank.v1.RankSystem.List:output_type -> biconom.type.RankSystem
	5, // 11: biconom.admin.rank.v1.RankSystem.HeaderGet:output_type -> biconom.type.RankSystem.Header
	5, // 12: biconom.admin.rank.v1.RankSystem.HeaderList:output_type -> biconom.type.RankSystem.Header
	4, // 13: biconom.admin.rank.v1.RankSystem.Create:output_type -> biconom.type.RankSystem
	5, // 14: biconom.admin.rank.v1.RankSystem.Publish:output_type -> biconom.type.RankSystem.Header
	5, // 15: biconom.admin.rank.v1.RankSystem.Unpublish:output_type -> biconom.type.RankSystem.Header
	9, // [9:16] is the sub-list for method output_type
	2, // [2:9] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_biconom_admin_rank_v1_rank_system_proto_init() }
func file_biconom_admin_rank_v1_rank_system_proto_init() {
	if File_biconom_admin_rank_v1_rank_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_rank_v1_rank_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSystemListRequest); i {
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
		file_biconom_admin_rank_v1_rank_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSystemCreateRequest); i {
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
			RawDescriptor: file_biconom_admin_rank_v1_rank_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_rank_v1_rank_system_proto_goTypes,
		DependencyIndexes: file_biconom_admin_rank_v1_rank_system_proto_depIdxs,
		MessageInfos:      file_biconom_admin_rank_v1_rank_system_proto_msgTypes,
	}.Build()
	File_biconom_admin_rank_v1_rank_system_proto = out.File
	file_biconom_admin_rank_v1_rank_system_proto_rawDesc = nil
	file_biconom_admin_rank_v1_rank_system_proto_goTypes = nil
	file_biconom_admin_rank_v1_rank_system_proto_depIdxs = nil
}
