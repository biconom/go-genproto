// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/rank/v1/rank_network_statistics.proto

package service_admin_rank_pb

import (
	account "github.com/biconom/go-genproto/biconom/type/account"
	rank_statistics "github.com/biconom/go-genproto/biconom/type/rank_statistics"
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

type RankNetworkStatisticsAccountCountListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Referral *account.Account_ID             `protobuf:"bytes,1,opt,name=referral,proto3" json:"referral,omitempty"`
	Step     *rank_system.RankSystem_Rank_ID `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort     *sort.Sort                      `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankNetworkStatisticsAccountCountListRequest) Reset() {
	*x = RankNetworkStatisticsAccountCountListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankNetworkStatisticsAccountCountListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankNetworkStatisticsAccountCountListRequest) ProtoMessage() {}

func (x *RankNetworkStatisticsAccountCountListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankNetworkStatisticsAccountCountListRequest.ProtoReflect.Descriptor instead.
func (*RankNetworkStatisticsAccountCountListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *RankNetworkStatisticsAccountCountListRequest) GetReferral() *account.Account_ID {
	if x != nil {
		return x.Referral
	}
	return nil
}

func (x *RankNetworkStatisticsAccountCountListRequest) GetStep() *rank_system.RankSystem_Rank_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankNetworkStatisticsAccountCountListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type RankNetworkStatisticsAccountCountListByRankSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Referral   *account.Account_ID                   `protobuf:"bytes,1,opt,name=referral,proto3" json:"referral,omitempty"`
	RankSystem *rank_system.RankSystem_ID            `protobuf:"bytes,2,opt,name=rank_system,json=rankSystem,proto3" json:"rank_system,omitempty"`
	Step       *rank_system.RankSystem_Rank_ID_Inner `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	Sort       *sort.Sort                            `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) Reset() {
	*x = RankNetworkStatisticsAccountCountListByRankSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankNetworkStatisticsAccountCountListByRankSystemRequest) ProtoMessage() {}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankNetworkStatisticsAccountCountListByRankSystemRequest.ProtoReflect.Descriptor instead.
func (*RankNetworkStatisticsAccountCountListByRankSystemRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescGZIP(), []int{1}
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) GetReferral() *account.Account_ID {
	if x != nil {
		return x.Referral
	}
	return nil
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) GetRankSystem() *rank_system.RankSystem_ID {
	if x != nil {
		return x.RankSystem
	}
	return nil
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) GetStep() *rank_system.RankSystem_Rank_ID_Inner {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankNetworkStatisticsAccountCountListByRankSystemRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type RankNetworkStatisticsAccountTopRankListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Referral *account.Account_ID             `protobuf:"bytes,1,opt,name=referral,proto3" json:"referral,omitempty"`
	TopLimit uint32                          `protobuf:"varint,2,opt,name=top_limit,json=topLimit,proto3" json:"top_limit,omitempty"`
	Step     *rank_system.RankSystem_Rank_ID `protobuf:"bytes,3,opt,name=step,proto3" json:"step,omitempty"`
	Sort     *sort.Sort                      `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) Reset() {
	*x = RankNetworkStatisticsAccountTopRankListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankNetworkStatisticsAccountTopRankListRequest) ProtoMessage() {}

func (x *RankNetworkStatisticsAccountTopRankListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankNetworkStatisticsAccountTopRankListRequest.ProtoReflect.Descriptor instead.
func (*RankNetworkStatisticsAccountTopRankListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescGZIP(), []int{2}
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) GetReferral() *account.Account_ID {
	if x != nil {
		return x.Referral
	}
	return nil
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) GetTopLimit() uint32 {
	if x != nil {
		return x.TopLimit
	}
	return 0
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) GetStep() *rank_system.RankSystem_Rank_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankNetworkStatisticsAccountTopRankListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type RankNetworkStatisticsAccountTopRankListByRankSystemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Referral   *account.Account_ID             `protobuf:"bytes,1,opt,name=referral,proto3" json:"referral,omitempty"`
	RankSystem *rank_system.RankSystem_ID      `protobuf:"bytes,2,opt,name=rank_system,json=rankSystem,proto3" json:"rank_system,omitempty"`
	TopLimit   uint32                          `protobuf:"varint,3,opt,name=top_limit,json=topLimit,proto3" json:"top_limit,omitempty"`
	Step       *rank_system.RankSystem_Rank_ID `protobuf:"bytes,4,opt,name=step,proto3" json:"step,omitempty"`
	Sort       *sort.Sort                      `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) Reset() {
	*x = RankNetworkStatisticsAccountTopRankListByRankSystemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RankNetworkStatisticsAccountTopRankListByRankSystemRequest) ProtoMessage() {}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RankNetworkStatisticsAccountTopRankListByRankSystemRequest.ProtoReflect.Descriptor instead.
func (*RankNetworkStatisticsAccountTopRankListByRankSystemRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescGZIP(), []int{3}
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) GetReferral() *account.Account_ID {
	if x != nil {
		return x.Referral
	}
	return nil
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) GetRankSystem() *rank_system.RankSystem_ID {
	if x != nil {
		return x.RankSystem
	}
	return nil
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) GetTopLimit() uint32 {
	if x != nil {
		return x.TopLimit
	}
	return 0
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) GetStep() *rank_system.RankSystem_Rank_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *RankNetworkStatisticsAccountTopRankListByRankSystemRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

var File_biconom_admin_rank_v1_rank_network_statistics_proto protoreflect.FileDescriptor

var file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDesc = []byte{
	0x0a, 0x33, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x2c, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x49, 0x44, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x92, 0x02, 0x0a, 0x38, 0x52,
	0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x49, 0x44, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x3c, 0x0a,
	0x0b, 0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x52,
	0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x3a, 0x0a, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22,
	0xe1, 0x01, 0x0a, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x08,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x6f, 0x70,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x22, 0xab, 0x02, 0x0a, 0x3a, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x52, 0x08,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x3c, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x6b,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e,
	0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x49, 0x44, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x6b,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x70, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x32, 0xc5, 0x06, 0x0a, 0x15, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x68, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x12, 0x28,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x2e, 0x49, 0x44, 0x1a, 0x29, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x43, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x9e,
	0x01, 0x0a, 0x1c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12,
	0x4f, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x6a, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e,
	0x6b, 0x47, 0x65, 0x74, 0x12, 0x2a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x2e, 0x49, 0x44,
	0x1a, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x12,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x45, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x6f, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12, 0xa0, 0x01, 0x0a, 0x1e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x51, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f,
	0x70, 0x52, 0x61, 0x6e, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x6f, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescOnce sync.Once
	file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescData = file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDesc
)

func file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescGZIP() []byte {
	file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescOnce.Do(func() {
		file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescData)
	})
	return file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDescData
}

var file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_biconom_admin_rank_v1_rank_network_statistics_proto_goTypes = []interface{}{
	(*RankNetworkStatisticsAccountCountListRequest)(nil),               // 0: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListRequest
	(*RankNetworkStatisticsAccountCountListByRankSystemRequest)(nil),   // 1: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest
	(*RankNetworkStatisticsAccountTopRankListRequest)(nil),             // 2: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListRequest
	(*RankNetworkStatisticsAccountTopRankListByRankSystemRequest)(nil), // 3: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest
	(*account.Account_ID)(nil),                                         // 4: biconom.type.Account.ID
	(*rank_system.RankSystem_Rank_ID)(nil),                             // 5: biconom.type.RankSystem.Rank.ID
	(*sort.Sort)(nil),                                                  // 6: biconom.type.Sort
	(*rank_system.RankSystem_ID)(nil),                                  // 7: biconom.type.RankSystem.ID
	(*rank_system.RankSystem_Rank_ID_Inner)(nil),                       // 8: biconom.type.RankSystem.Rank.ID.Inner
	(*account.Account_RankSystem_Rank_ID)(nil),                         // 9: biconom.type.Account.RankSystem.Rank.ID
	(*rank_statistics.RankStatistics_AccountTop_ID)(nil),               // 10: biconom.type.RankStatistics.AccountTop.ID
	(*rank_statistics.RankStatistics_AccountCount)(nil),                // 11: biconom.type.RankStatistics.AccountCount
	(*rank_statistics.RankStatistics_AccountTop)(nil),                  // 12: biconom.type.RankStatistics.AccountTop
}
var file_biconom_admin_rank_v1_rank_network_statistics_proto_depIdxs = []int32{
	4,  // 0: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListRequest.referral:type_name -> biconom.type.Account.ID
	5,  // 1: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListRequest.step:type_name -> biconom.type.RankSystem.Rank.ID
	6,  // 2: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListRequest.sort:type_name -> biconom.type.Sort
	4,  // 3: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest.referral:type_name -> biconom.type.Account.ID
	7,  // 4: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest.rank_system:type_name -> biconom.type.RankSystem.ID
	8,  // 5: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest.step:type_name -> biconom.type.RankSystem.Rank.ID.Inner
	6,  // 6: biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest.sort:type_name -> biconom.type.Sort
	4,  // 7: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListRequest.referral:type_name -> biconom.type.Account.ID
	5,  // 8: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListRequest.step:type_name -> biconom.type.RankSystem.Rank.ID
	6,  // 9: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListRequest.sort:type_name -> biconom.type.Sort
	4,  // 10: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest.referral:type_name -> biconom.type.Account.ID
	7,  // 11: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest.rank_system:type_name -> biconom.type.RankSystem.ID
	5,  // 12: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest.step:type_name -> biconom.type.RankSystem.Rank.ID
	6,  // 13: biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest.sort:type_name -> biconom.type.Sort
	9,  // 14: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountGet:input_type -> biconom.type.Account.RankSystem.Rank.ID
	0,  // 15: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountList:input_type -> biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListRequest
	1,  // 16: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountListByRankSystem:input_type -> biconom.admin.rank.v1.RankNetworkStatisticsAccountCountListByRankSystemRequest
	10, // 17: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankGet:input_type -> biconom.type.RankStatistics.AccountTop.ID
	2,  // 18: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankList:input_type -> biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListRequest
	3,  // 19: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankListByRankSystem:input_type -> biconom.admin.rank.v1.RankNetworkStatisticsAccountTopRankListByRankSystemRequest
	11, // 20: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountGet:output_type -> biconom.type.RankStatistics.AccountCount
	11, // 21: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountList:output_type -> biconom.type.RankStatistics.AccountCount
	11, // 22: biconom.admin.rank.v1.RankNetworkStatistics.AccountCountListByRankSystem:output_type -> biconom.type.RankStatistics.AccountCount
	12, // 23: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankGet:output_type -> biconom.type.RankStatistics.AccountTop
	12, // 24: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankList:output_type -> biconom.type.RankStatistics.AccountTop
	12, // 25: biconom.admin.rank.v1.RankNetworkStatistics.AccountTopRankListByRankSystem:output_type -> biconom.type.RankStatistics.AccountTop
	20, // [20:26] is the sub-list for method output_type
	14, // [14:20] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_biconom_admin_rank_v1_rank_network_statistics_proto_init() }
func file_biconom_admin_rank_v1_rank_network_statistics_proto_init() {
	if File_biconom_admin_rank_v1_rank_network_statistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankNetworkStatisticsAccountCountListRequest); i {
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
		file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankNetworkStatisticsAccountCountListByRankSystemRequest); i {
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
		file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankNetworkStatisticsAccountTopRankListRequest); i {
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
		file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RankNetworkStatisticsAccountTopRankListByRankSystemRequest); i {
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
			RawDescriptor: file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_rank_v1_rank_network_statistics_proto_goTypes,
		DependencyIndexes: file_biconom_admin_rank_v1_rank_network_statistics_proto_depIdxs,
		MessageInfos:      file_biconom_admin_rank_v1_rank_network_statistics_proto_msgTypes,
	}.Build()
	File_biconom_admin_rank_v1_rank_network_statistics_proto = out.File
	file_biconom_admin_rank_v1_rank_network_statistics_proto_rawDesc = nil
	file_biconom_admin_rank_v1_rank_network_statistics_proto_goTypes = nil
	file_biconom_admin_rank_v1_rank_network_statistics_proto_depIdxs = nil
}
