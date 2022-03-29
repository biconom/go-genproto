// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/client/rank/v1/rank.proto

package service_client_rank_pb

import (
	condition "github.com/biconom/go-genproto/biconom/type/condition"
	rank "github.com/biconom/go-genproto/biconom/type/rank"
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

type RankListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *rank.RankSystem_Rank_ID `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	Sort *sort.Sort               `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankListRequest) Reset() {
	*x = RankListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_client_rank_v1_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankListRequest) ProtoMessage() {}

func (x *RankListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_client_rank_v1_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankListRequest.ProtoReflect.Descriptor instead.
func (*RankListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_client_rank_v1_rank_proto_rawDescGZIP(), []int{0}
}

func (x *RankListRequest) GetStep() *rank.RankSystem_Rank_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type RankListByRankSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankSystem *rank.RankSystem_ID            `protobuf:"bytes,1,opt,name=rank_system,json=rankSystem,proto3" json:"rank_system,omitempty"`
	Step       *rank.RankSystem_Rank_ID_Inner `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort       *sort.Sort                     `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankListByRankSystemRequest) Reset() {
	*x = RankListByRankSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_client_rank_v1_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankListByRankSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankListByRankSystemRequest) ProtoMessage() {}

func (x *RankListByRankSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_client_rank_v1_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankListByRankSystemRequest.ProtoReflect.Descriptor instead.
func (*RankListByRankSystemRequest) Descriptor() ([]byte, []int) {
	return file_biconom_client_rank_v1_rank_proto_rawDescGZIP(), []int{1}
}

func (x *RankListByRankSystemRequest) GetRankSystem() *rank.RankSystem_ID {
	if x != nil {
		return x.RankSystem
	}
	return nil
}

func (x *RankListByRankSystemRequest) GetStep() *rank.RankSystem_Rank_ID_Inner {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankListByRankSystemRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

var File_biconom_client_rank_v1_rank_proto protoreflect.FileDescriptor

var file_biconom_client_rank_v1_rank_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0f, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x1b, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x49, 0x44, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49,
	0x44, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x32, 0xbc, 0x07, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x4a,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x6c, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x1a, 0x1d, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6f, 0x0a, 0x15, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x1a, 0x24, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x7c, 0x0a, 0x1b, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x61, 0x6e,
	0x6b, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_client_rank_v1_rank_proto_rawDescOnce sync.Once
	file_biconom_client_rank_v1_rank_proto_rawDescData = file_biconom_client_rank_v1_rank_proto_rawDesc
)

func file_biconom_client_rank_v1_rank_proto_rawDescGZIP() []byte {
	file_biconom_client_rank_v1_rank_proto_rawDescOnce.Do(func() {
		file_biconom_client_rank_v1_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_client_rank_v1_rank_proto_rawDescData)
	})
	return file_biconom_client_rank_v1_rank_proto_rawDescData
}

var file_biconom_client_rank_v1_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_biconom_client_rank_v1_rank_proto_goTypes = []interface{}{
	(*RankListRequest)(nil),               // 0: biconom.client.rank.v1.RankListRequest
	(*RankListByRankSystemRequest)(nil),   // 1: biconom.client.rank.v1.RankListByRankSystemRequest
	(*rank.RankSystem_Rank_ID)(nil),       // 2: biconom.type.RankSystem.Rank.ID
	(*sort.Sort)(nil),                     // 3: biconom.type.Sort
	(*rank.RankSystem_ID)(nil),            // 4: biconom.type.RankSystem.ID
	(*rank.RankSystem_Rank_ID_Inner)(nil), // 5: biconom.type.RankSystem.Rank.ID.Inner
	(*rank.RankSystem_Option)(nil),        // 6: biconom.type.RankSystem.Option
	(*rank.RankSystem_Rank)(nil),          // 7: biconom.type.RankSystem.Rank
	(*rank.RankSystem_Rank_Header)(nil),   // 8: biconom.type.RankSystem.Rank.Header
	(*condition.Condition)(nil),           // 9: biconom.type.Condition
}
var file_biconom_client_rank_v1_rank_proto_depIdxs = []int32{
	2,  // 0: biconom.client.rank.v1.RankListRequest.step:type_name -> biconom.type.RankSystem.Rank.ID
	3,  // 1: biconom.client.rank.v1.RankListRequest.sort:type_name -> biconom.type.Sort
	4,  // 2: biconom.client.rank.v1.RankListByRankSystemRequest.rank_system:type_name -> biconom.type.RankSystem.ID
	5,  // 3: biconom.client.rank.v1.RankListByRankSystemRequest.step:type_name -> biconom.type.RankSystem.Rank.ID.Inner
	3,  // 4: biconom.client.rank.v1.RankListByRankSystemRequest.sort:type_name -> biconom.type.Sort
	2,  // 5: biconom.client.rank.v1.Rank.Get:input_type -> biconom.type.RankSystem.Rank.ID
	0,  // 6: biconom.client.rank.v1.Rank.List:input_type -> biconom.client.rank.v1.RankListRequest
	1,  // 7: biconom.client.rank.v1.Rank.ListByRankSystem:input_type -> biconom.client.rank.v1.RankListByRankSystemRequest
	2,  // 8: biconom.client.rank.v1.Rank.ValueGet:input_type -> biconom.type.RankSystem.Rank.ID
	0,  // 9: biconom.client.rank.v1.Rank.ValueList:input_type -> biconom.client.rank.v1.RankListRequest
	1,  // 10: biconom.client.rank.v1.Rank.ValueListByRankSystem:input_type -> biconom.client.rank.v1.RankListByRankSystemRequest
	2,  // 11: biconom.client.rank.v1.Rank.ValueHeaderGet:input_type -> biconom.type.RankSystem.Rank.ID
	0,  // 12: biconom.client.rank.v1.Rank.ValueHeaderList:input_type -> biconom.client.rank.v1.RankListRequest
	1,  // 13: biconom.client.rank.v1.Rank.ValueHeaderListByRankSystem:input_type -> biconom.client.rank.v1.RankListByRankSystemRequest
	2,  // 14: biconom.client.rank.v1.Rank.ConditionGet:input_type -> biconom.type.RankSystem.Rank.ID
	6,  // 15: biconom.client.rank.v1.Rank.Get:output_type -> biconom.type.RankSystem.Option
	6,  // 16: biconom.client.rank.v1.Rank.List:output_type -> biconom.type.RankSystem.Option
	6,  // 17: biconom.client.rank.v1.Rank.ListByRankSystem:output_type -> biconom.type.RankSystem.Option
	7,  // 18: biconom.client.rank.v1.Rank.ValueGet:output_type -> biconom.type.RankSystem.Rank
	7,  // 19: biconom.client.rank.v1.Rank.ValueList:output_type -> biconom.type.RankSystem.Rank
	7,  // 20: biconom.client.rank.v1.Rank.ValueListByRankSystem:output_type -> biconom.type.RankSystem.Rank
	8,  // 21: biconom.client.rank.v1.Rank.ValueHeaderGet:output_type -> biconom.type.RankSystem.Rank.Header
	8,  // 22: biconom.client.rank.v1.Rank.ValueHeaderList:output_type -> biconom.type.RankSystem.Rank.Header
	8,  // 23: biconom.client.rank.v1.Rank.ValueHeaderListByRankSystem:output_type -> biconom.type.RankSystem.Rank.Header
	9,  // 24: biconom.client.rank.v1.Rank.ConditionGet:output_type -> biconom.type.Condition
	15, // [15:25] is the sub-list for method output_type
	5,  // [5:15] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_biconom_client_rank_v1_rank_proto_init() }
func file_biconom_client_rank_v1_rank_proto_init() {
	if File_biconom_client_rank_v1_rank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_client_rank_v1_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankListRequest); i {
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
		file_biconom_client_rank_v1_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankListByRankSystemRequest); i {
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
			RawDescriptor: file_biconom_client_rank_v1_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_client_rank_v1_rank_proto_goTypes,
		DependencyIndexes: file_biconom_client_rank_v1_rank_proto_depIdxs,
		MessageInfos:      file_biconom_client_rank_v1_rank_proto_msgTypes,
	}.Build()
	File_biconom_client_rank_v1_rank_proto = out.File
	file_biconom_client_rank_v1_rank_proto_rawDesc = nil
	file_biconom_client_rank_v1_rank_proto_goTypes = nil
	file_biconom_client_rank_v1_rank_proto_depIdxs = nil
}
