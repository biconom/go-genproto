// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/rank/v1/rank_system_account.proto

package service_admin_rank_pb

import (
	account "github.com/biconom/go-genproto/biconom/type/account"
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

type RankSystemAccountListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account                 *account.Account_ID        `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Step                    *rank_system.RankSystem_ID `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort                    *sort.Sort                 `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedRankSystem bool                       `protobuf:"varint,4,opt,name=only_published_rank_system,json=onlyPublishedRankSystem,proto3" json:"only_published_rank_system,omitempty"`
}

func (x *RankSystemAccountListRequest) Reset() {
	*x = RankSystemAccountListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSystemAccountListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSystemAccountListRequest) ProtoMessage() {}

func (x *RankSystemAccountListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSystemAccountListRequest.ProtoReflect.Descriptor instead.
func (*RankSystemAccountListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_system_account_proto_rawDescGZIP(), []int{0}
}

func (x *RankSystemAccountListRequest) GetAccount() *account.Account_ID {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *RankSystemAccountListRequest) GetStep() *rank_system.RankSystem_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankSystemAccountListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *RankSystemAccountListRequest) GetOnlyPublishedRankSystem() bool {
	if x != nil {
		return x.OnlyPublishedRankSystem
	}
	return false
}

type RankSystemAccountWithRankActivatedAndNextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header        *rank_system.RankSystem_Header   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	RankActivated *account.Account_RankSystem_Rank `protobuf:"bytes,2,opt,name=rank_activated,json=rankActivated,proto3" json:"rank_activated,omitempty"`
	RankNext      *account.Account_RankSystem_Rank `protobuf:"bytes,3,opt,name=rank_next,json=rankNext,proto3" json:"rank_next,omitempty"`
}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) Reset() {
	*x = RankSystemAccountWithRankActivatedAndNextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankSystemAccountWithRankActivatedAndNextResponse) ProtoMessage() {}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankSystemAccountWithRankActivatedAndNextResponse.ProtoReflect.Descriptor instead.
func (*RankSystemAccountWithRankActivatedAndNextResponse) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_system_account_proto_rawDescGZIP(), []int{1}
}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) GetHeader() *rank_system.RankSystem_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) GetRankActivated() *account.Account_RankSystem_Rank {
	if x != nil {
		return x.RankActivated
	}
	return nil
}

func (x *RankSystemAccountWithRankActivatedAndNextResponse) GetRankNext() *account.Account_RankSystem_Rank {
	if x != nil {
		return x.RankNext
	}
	return nil
}

var File_biconom_admin_rank_v1_rank_system_account_proto protoreflect.FileDescriptor

var file_biconom_admin_rank_v1_rank_system_account_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01,
	0x0a, 0x1c, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x6f,
	0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x72, 0x61,
	0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x17, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0xfe, 0x01, 0x0a, 0x31, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x6b, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x0d, 0x72, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x42, 0x0a, 0x09, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52,
	0x08, 0x72, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x78, 0x74, 0x32, 0xde, 0x05, 0x0a, 0x11, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x4e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a, 0x20, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x66, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x6e,
	0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x23, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a,
	0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x8e, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e,
	0x64, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x23, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x1a, 0x48, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa1, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x33, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x48, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x6b, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_admin_rank_v1_rank_system_account_proto_rawDescOnce sync.Once
	file_biconom_admin_rank_v1_rank_system_account_proto_rawDescData = file_biconom_admin_rank_v1_rank_system_account_proto_rawDesc
)

func file_biconom_admin_rank_v1_rank_system_account_proto_rawDescGZIP() []byte {
	file_biconom_admin_rank_v1_rank_system_account_proto_rawDescOnce.Do(func() {
		file_biconom_admin_rank_v1_rank_system_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_rank_v1_rank_system_account_proto_rawDescData)
	})
	return file_biconom_admin_rank_v1_rank_system_account_proto_rawDescData
}

var file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_biconom_admin_rank_v1_rank_system_account_proto_goTypes = []interface{}{
	(*RankSystemAccountListRequest)(nil),                      // 0: biconom.admin.rank.v1.RankSystemAccountListRequest
	(*RankSystemAccountWithRankActivatedAndNextResponse)(nil), // 1: biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse
	(*account.Account_ID)(nil),                                // 2: biconom.type.Account.ID
	(*rank_system.RankSystem_ID)(nil),                         // 3: biconom.type.RankSystem.ID
	(*sort.Sort)(nil),                                         // 4: biconom.type.Sort
	(*rank_system.RankSystem_Header)(nil),                     // 5: biconom.type.RankSystem.Header
	(*account.Account_RankSystem_Rank)(nil),                   // 6: biconom.type.Account.RankSystem.Rank
	(*account.Account_RankSystem_ID)(nil),                     // 7: biconom.type.Account.RankSystem.ID
	(*account.Account_RankSystem)(nil),                        // 8: biconom.type.Account.RankSystem
	(*account.Account_RankSystem_Option)(nil),                 // 9: biconom.type.Account.RankSystem.Option
}
var file_biconom_admin_rank_v1_rank_system_account_proto_depIdxs = []int32{
	2,  // 0: biconom.admin.rank.v1.RankSystemAccountListRequest.account:type_name -> biconom.type.Account.ID
	3,  // 1: biconom.admin.rank.v1.RankSystemAccountListRequest.step:type_name -> biconom.type.RankSystem.ID
	4,  // 2: biconom.admin.rank.v1.RankSystemAccountListRequest.sort:type_name -> biconom.type.Sort
	5,  // 3: biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse.header:type_name -> biconom.type.RankSystem.Header
	6,  // 4: biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse.rank_activated:type_name -> biconom.type.Account.RankSystem.Rank
	6,  // 5: biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse.rank_next:type_name -> biconom.type.Account.RankSystem.Rank
	7,  // 6: biconom.admin.rank.v1.RankSystemAccount.Get:input_type -> biconom.type.Account.RankSystem.ID
	0,  // 7: biconom.admin.rank.v1.RankSystemAccount.List:input_type -> biconom.admin.rank.v1.RankSystemAccountListRequest
	7,  // 8: biconom.admin.rank.v1.RankSystemAccount.GetWithRankActivated:input_type -> biconom.type.Account.RankSystem.ID
	0,  // 9: biconom.admin.rank.v1.RankSystemAccount.ListWithRankActivated:input_type -> biconom.admin.rank.v1.RankSystemAccountListRequest
	7,  // 10: biconom.admin.rank.v1.RankSystemAccount.GetWithRankActivatedAndNext:input_type -> biconom.type.Account.RankSystem.ID
	0,  // 11: biconom.admin.rank.v1.RankSystemAccount.ListWithRankActivatedAndNext:input_type -> biconom.admin.rank.v1.RankSystemAccountListRequest
	8,  // 12: biconom.admin.rank.v1.RankSystemAccount.Get:output_type -> biconom.type.Account.RankSystem
	8,  // 13: biconom.admin.rank.v1.RankSystemAccount.List:output_type -> biconom.type.Account.RankSystem
	9,  // 14: biconom.admin.rank.v1.RankSystemAccount.GetWithRankActivated:output_type -> biconom.type.Account.RankSystem.Option
	9,  // 15: biconom.admin.rank.v1.RankSystemAccount.ListWithRankActivated:output_type -> biconom.type.Account.RankSystem.Option
	1,  // 16: biconom.admin.rank.v1.RankSystemAccount.GetWithRankActivatedAndNext:output_type -> biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse
	1,  // 17: biconom.admin.rank.v1.RankSystemAccount.ListWithRankActivatedAndNext:output_type -> biconom.admin.rank.v1.RankSystemAccountWithRankActivatedAndNextResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_biconom_admin_rank_v1_rank_system_account_proto_init() }
func file_biconom_admin_rank_v1_rank_system_account_proto_init() {
	if File_biconom_admin_rank_v1_rank_system_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSystemAccountListRequest); i {
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
		file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankSystemAccountWithRankActivatedAndNextResponse); i {
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
			RawDescriptor: file_biconom_admin_rank_v1_rank_system_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_rank_v1_rank_system_account_proto_goTypes,
		DependencyIndexes: file_biconom_admin_rank_v1_rank_system_account_proto_depIdxs,
		MessageInfos:      file_biconom_admin_rank_v1_rank_system_account_proto_msgTypes,
	}.Build()
	File_biconom_admin_rank_v1_rank_system_account_proto = out.File
	file_biconom_admin_rank_v1_rank_system_account_proto_rawDesc = nil
	file_biconom_admin_rank_v1_rank_system_account_proto_goTypes = nil
	file_biconom_admin_rank_v1_rank_system_account_proto_depIdxs = nil
}
