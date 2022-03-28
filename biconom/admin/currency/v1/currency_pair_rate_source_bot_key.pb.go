// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/currency/v1/currency_pair_rate_source_bot_key.proto

package service_admin_currency_pb

import (
	bot_key "github.com/biconom/go-genproto/biconom/type/bot_key"
	currency "github.com/biconom/go-genproto/biconom/type/currency"
	sort "github.com/biconom/go-genproto/biconom/type/sort"
	until "github.com/biconom/go-genproto/biconom/type/until"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrencyPairRateSourceBotKeyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step                *currency.Currency_Pair_Rate_Source_BotKey_ID `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	Sort                *sort.Sort                                    `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedSource bool                                          `protobuf:"varint,3,opt,name=only_published_source,json=onlyPublishedSource,proto3" json:"only_published_source,omitempty"`
}

func (x *CurrencyPairRateSourceBotKeyListRequest) Reset() {
	*x = CurrencyPairRateSourceBotKeyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairRateSourceBotKeyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairRateSourceBotKeyListRequest) ProtoMessage() {}

func (x *CurrencyPairRateSourceBotKeyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairRateSourceBotKeyListRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairRateSourceBotKeyListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyPairRateSourceBotKeyListRequest) GetStep() *currency.Currency_Pair_Rate_Source_BotKey_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyListRequest) GetOnlyPublishedSource() bool {
	if x != nil {
		return x.OnlyPublishedSource
	}
	return false
}

type CurrencyPairRateSourceBotKeyListBySourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateSource          *currency.Currency_Pair_Rate_Source_ID `protobuf:"bytes,1,opt,name=rate_source,json=rateSource,proto3" json:"rate_source,omitempty"`
	Step                *bot_key.BotKey_ID                     `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort                *sort.Sort                             `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedSource bool                                   `protobuf:"varint,4,opt,name=only_published_source,json=onlyPublishedSource,proto3" json:"only_published_source,omitempty"`
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) Reset() {
	*x = CurrencyPairRateSourceBotKeyListBySourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairRateSourceBotKeyListBySourceRequest) ProtoMessage() {}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairRateSourceBotKeyListBySourceRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairRateSourceBotKeyListBySourceRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) GetRateSource() *currency.Currency_Pair_Rate_Source_ID {
	if x != nil {
		return x.RateSource
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) GetStep() *bot_key.BotKey_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyListBySourceRequest) GetOnlyPublishedSource() bool {
	if x != nil {
		return x.OnlyPublishedSource
	}
	return false
}

type CurrencyPairRateSourceBotKeyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateSource  *currency.Currency_Pair_Rate_Source_ID `protobuf:"bytes,1,opt,name=rate_source,json=rateSource,proto3" json:"rate_source,omitempty"`
	BotKeyTitle string                                 `protobuf:"bytes,2,opt,name=bot_key_title,json=botKeyTitle,proto3" json:"bot_key_title,omitempty"`
	Ttl         *until.Until                           `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	IpMasks     []string                               `protobuf:"bytes,4,rep,name=ip_masks,json=ipMasks,proto3" json:"ip_masks,omitempty"`
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) Reset() {
	*x = CurrencyPairRateSourceBotKeyCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairRateSourceBotKeyCreateRequest) ProtoMessage() {}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairRateSourceBotKeyCreateRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairRateSourceBotKeyCreateRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) GetRateSource() *currency.Currency_Pair_Rate_Source_ID {
	if x != nil {
		return x.RateSource
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) GetBotKeyTitle() string {
	if x != nil {
		return x.BotKeyTitle
	}
	return ""
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) GetTtl() *until.Until {
	if x != nil {
		return x.Ttl
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyCreateRequest) GetIpMasks() []string {
	if x != nil {
		return x.IpMasks
	}
	return nil
}

type CurrencyPairRateSourceBotKeyIpMasksSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotKey  *currency.Currency_Pair_Rate_Source_BotKey_ID `protobuf:"bytes,1,opt,name=bot_key,json=botKey,proto3" json:"bot_key,omitempty"`
	IpMasks []string                                      `protobuf:"bytes,2,rep,name=ip_masks,json=ipMasks,proto3" json:"ip_masks,omitempty"`
}

func (x *CurrencyPairRateSourceBotKeyIpMasksSetRequest) Reset() {
	*x = CurrencyPairRateSourceBotKeyIpMasksSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairRateSourceBotKeyIpMasksSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairRateSourceBotKeyIpMasksSetRequest) ProtoMessage() {}

func (x *CurrencyPairRateSourceBotKeyIpMasksSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairRateSourceBotKeyIpMasksSetRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairRateSourceBotKeyIpMasksSetRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescGZIP(), []int{3}
}

func (x *CurrencyPairRateSourceBotKeyIpMasksSetRequest) GetBotKey() *currency.Currency_Pair_Rate_Source_BotKey_ID {
	if x != nil {
		return x.BotKey
	}
	return nil
}

func (x *CurrencyPairRateSourceBotKeyIpMasksSetRequest) GetIpMasks() []string {
	if x != nil {
		return x.IpMasks
	}
	return nil
}

var File_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto protoreflect.FileDescriptor

var file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDesc = []byte{
	0x0a, 0x41, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1a,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x62, 0x6f, 0x74,
	0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x75,
	0x6e, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x27, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65,
	0x79, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x87, 0x02, 0x0a, 0x2f, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x52, 0x61, 0x74,
	0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x52, 0x0a, 0x72, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x32, 0x0a, 0x15,
	0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6f, 0x6e, 0x6c,
	0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xde, 0x01, 0x0a, 0x29, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65,
	0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b,
	0x0a, 0x0b, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x49, 0x44, 0x52,
	0x0a, 0x72, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x62,
	0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x25, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x74, 0x69,
	0x6c, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x70, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x70, 0x4d, 0x61, 0x73, 0x6b,
	0x73, 0x22, 0x96, 0x01, 0x0a, 0x2d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b,
	0x65, 0x79, 0x49, 0x70, 0x4d, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x07, 0x62, 0x6f, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69,
	0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f,
	0x74, 0x4b, 0x65, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x06, 0x62, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x70, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x70, 0x4d, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xe5, 0x05, 0x0a, 0x1c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x6a, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x31, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e,
	0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b,
	0x65, 0x79, 0x2e, 0x49, 0x44, 0x1a, 0x2e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61,
	0x69, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42,
	0x6f, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x42, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74,
	0x4b, 0x65, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x74, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x4a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x80, 0x01,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65,
	0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x31, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x0a, 0x49, 0x70, 0x4d, 0x61,
	0x73, 0x6b, 0x73, 0x53, 0x65, 0x74, 0x12, 0x48, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79, 0x49,
	0x70, 0x4d, 0x61, 0x73, 0x6b, 0x73, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x74, 0x4b, 0x65, 0x79,
	0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescOnce sync.Once
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescData = file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDesc
)

func file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescGZIP() []byte {
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescOnce.Do(func() {
		file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescData)
	})
	return file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDescData
}

var file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_goTypes = []interface{}{
	(*CurrencyPairRateSourceBotKeyListRequest)(nil),         // 0: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListRequest
	(*CurrencyPairRateSourceBotKeyListBySourceRequest)(nil), // 1: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListBySourceRequest
	(*CurrencyPairRateSourceBotKeyCreateRequest)(nil),       // 2: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyCreateRequest
	(*CurrencyPairRateSourceBotKeyIpMasksSetRequest)(nil),   // 3: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyIpMasksSetRequest
	(*currency.Currency_Pair_Rate_Source_BotKey_ID)(nil),    // 4: biconom.type.Currency.Pair.Rate.Source.BotKey.ID
	(*sort.Sort)(nil), // 5: biconom.type.Sort
	(*currency.Currency_Pair_Rate_Source_ID)(nil),     // 6: biconom.type.Currency.Pair.Rate.Source.ID
	(*bot_key.BotKey_ID)(nil),                         // 7: biconom.type.BotKey.ID
	(*until.Until)(nil),                               // 8: biconom.type.Until
	(*currency.Currency_Pair_Rate_Source_BotKey)(nil), // 9: biconom.type.Currency.Pair.Rate.Source.BotKey
	(*bot_key.BotKey)(nil),                            // 10: biconom.type.BotKey
	(*emptypb.Empty)(nil),                             // 11: google.protobuf.Empty
}
var file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_depIdxs = []int32{
	4,  // 0: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListRequest.step:type_name -> biconom.type.Currency.Pair.Rate.Source.BotKey.ID
	5,  // 1: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListRequest.sort:type_name -> biconom.type.Sort
	6,  // 2: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListBySourceRequest.rate_source:type_name -> biconom.type.Currency.Pair.Rate.Source.ID
	7,  // 3: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListBySourceRequest.step:type_name -> biconom.type.BotKey.ID
	5,  // 4: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListBySourceRequest.sort:type_name -> biconom.type.Sort
	6,  // 5: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyCreateRequest.rate_source:type_name -> biconom.type.Currency.Pair.Rate.Source.ID
	8,  // 6: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyCreateRequest.ttl:type_name -> biconom.type.Until
	4,  // 7: biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyIpMasksSetRequest.bot_key:type_name -> biconom.type.Currency.Pair.Rate.Source.BotKey.ID
	4,  // 8: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Get:input_type -> biconom.type.Currency.Pair.Rate.Source.BotKey.ID
	0,  // 9: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.List:input_type -> biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListRequest
	1,  // 10: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.ListBySource:input_type -> biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyListBySourceRequest
	2,  // 11: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Create:input_type -> biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyCreateRequest
	4,  // 12: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Revoke:input_type -> biconom.type.Currency.Pair.Rate.Source.BotKey.ID
	3,  // 13: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.IpMasksSet:input_type -> biconom.admin.currency.v1.CurrencyPairRateSourceBotKeyIpMasksSetRequest
	9,  // 14: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Get:output_type -> biconom.type.Currency.Pair.Rate.Source.BotKey
	9,  // 15: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.List:output_type -> biconom.type.Currency.Pair.Rate.Source.BotKey
	10, // 16: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.ListBySource:output_type -> biconom.type.BotKey
	9,  // 17: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Create:output_type -> biconom.type.Currency.Pair.Rate.Source.BotKey
	11, // 18: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.Revoke:output_type -> google.protobuf.Empty
	9,  // 19: biconom.admin.currency.v1.CurrencyPairRateSourceBotKey.IpMasksSet:output_type -> biconom.type.Currency.Pair.Rate.Source.BotKey
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_init() }
func file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_init() {
	if File_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairRateSourceBotKeyListRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairRateSourceBotKeyListBySourceRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairRateSourceBotKeyCreateRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairRateSourceBotKeyIpMasksSetRequest); i {
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
			RawDescriptor: file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_goTypes,
		DependencyIndexes: file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_depIdxs,
		MessageInfos:      file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_msgTypes,
	}.Build()
	File_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto = out.File
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_rawDesc = nil
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_goTypes = nil
	file_biconom_admin_currency_v1_currency_pair_rate_source_bot_key_proto_depIdxs = nil
}
