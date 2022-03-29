// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/currency/v1/currency_pair.proto

package service_admin_currency_pb

import (
	currency "github.com/biconom/go-genproto/biconom/type/currency"
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

type CurrencyPairListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step              *currency.Currency_Pair_ID `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	Sort              *sort.Sort                 `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedPair bool                       `protobuf:"varint,3,opt,name=only_published_pair,json=onlyPublishedPair,proto3" json:"only_published_pair,omitempty"`
}

func (x *CurrencyPairListRequest) Reset() {
	*x = CurrencyPairListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairListRequest) ProtoMessage() {}

func (x *CurrencyPairListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairListRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyPairListRequest) GetStep() *currency.Currency_Pair_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *CurrencyPairListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairListRequest) GetOnlyPublishedPair() bool {
	if x != nil {
		return x.OnlyPublishedPair
	}
	return false
}

type CurrencyPairListBySourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceCurrency    *currency.Currency_ID `protobuf:"bytes,1,opt,name=source_currency,json=sourceCurrency,proto3" json:"source_currency,omitempty"`
	Step              *currency.Currency_ID `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort              *sort.Sort            `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedPair bool                  `protobuf:"varint,4,opt,name=only_published_pair,json=onlyPublishedPair,proto3" json:"only_published_pair,omitempty"`
}

func (x *CurrencyPairListBySourceRequest) Reset() {
	*x = CurrencyPairListBySourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairListBySourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairListBySourceRequest) ProtoMessage() {}

func (x *CurrencyPairListBySourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairListBySourceRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairListBySourceRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyPairListBySourceRequest) GetSourceCurrency() *currency.Currency_ID {
	if x != nil {
		return x.SourceCurrency
	}
	return nil
}

func (x *CurrencyPairListBySourceRequest) GetStep() *currency.Currency_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *CurrencyPairListBySourceRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairListBySourceRequest) GetOnlyPublishedPair() bool {
	if x != nil {
		return x.OnlyPublishedPair
	}
	return false
}

type CurrencyPairListByTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetCurrency    *currency.Currency_ID `protobuf:"bytes,1,opt,name=target_currency,json=targetCurrency,proto3" json:"target_currency,omitempty"`
	Step              *currency.Currency_ID `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort              *sort.Sort            `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedPair bool                  `protobuf:"varint,4,opt,name=only_published_pair,json=onlyPublishedPair,proto3" json:"only_published_pair,omitempty"`
}

func (x *CurrencyPairListByTargetRequest) Reset() {
	*x = CurrencyPairListByTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairListByTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairListByTargetRequest) ProtoMessage() {}

func (x *CurrencyPairListByTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairListByTargetRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairListByTargetRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP(), []int{2}
}

func (x *CurrencyPairListByTargetRequest) GetTargetCurrency() *currency.Currency_ID {
	if x != nil {
		return x.TargetCurrency
	}
	return nil
}

func (x *CurrencyPairListByTargetRequest) GetStep() *currency.Currency_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *CurrencyPairListByTargetRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairListByTargetRequest) GetOnlyPublishedPair() bool {
	if x != nil {
		return x.OnlyPublishedPair
	}
	return false
}

type CurrencyPairListByCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Currency          *currency.Currency_ID `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	Sort              *sort.Sort            `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	OnlyPublishedPair bool                  `protobuf:"varint,3,opt,name=only_published_pair,json=onlyPublishedPair,proto3" json:"only_published_pair,omitempty"`
}

func (x *CurrencyPairListByCurrencyRequest) Reset() {
	*x = CurrencyPairListByCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairListByCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairListByCurrencyRequest) ProtoMessage() {}

func (x *CurrencyPairListByCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairListByCurrencyRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairListByCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP(), []int{3}
}

func (x *CurrencyPairListByCurrencyRequest) GetCurrency() *currency.Currency_ID {
	if x != nil {
		return x.Currency
	}
	return nil
}

func (x *CurrencyPairListByCurrencyRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *CurrencyPairListByCurrencyRequest) GetOnlyPublishedPair() bool {
	if x != nil {
		return x.OnlyPublishedPair
	}
	return false
}

type CurrencyPairCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceCurrency *currency.Currency_ID `protobuf:"bytes,1,opt,name=source_currency,json=sourceCurrency,proto3" json:"source_currency,omitempty"`
	TargetCurrency *currency.Currency_ID `protobuf:"bytes,2,opt,name=target_currency,json=targetCurrency,proto3" json:"target_currency,omitempty"`
	Published      bool                  `protobuf:"varint,3,opt,name=published,proto3" json:"published,omitempty"`
}

func (x *CurrencyPairCreateRequest) Reset() {
	*x = CurrencyPairCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyPairCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyPairCreateRequest) ProtoMessage() {}

func (x *CurrencyPairCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyPairCreateRequest.ProtoReflect.Descriptor instead.
func (*CurrencyPairCreateRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP(), []int{4}
}

func (x *CurrencyPairCreateRequest) GetSourceCurrency() *currency.Currency_ID {
	if x != nil {
		return x.SourceCurrency
	}
	return nil
}

func (x *CurrencyPairCreateRequest) GetTargetCurrency() *currency.Currency_ID {
	if x != nil {
		return x.TargetCurrency
	}
	return nil
}

func (x *CurrencyPairCreateRequest) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

var File_biconom_admin_currency_v1_currency_pair_proto protoreflect.FileDescriptor

var file_biconom_admin_currency_v1_currency_pair_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa5, 0x01, 0x0a, 0x17, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f,
	0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x6e, 0x6c, 0x79,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0xec, 0x01, 0x0a, 0x1f, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44,
	0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x2d, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12,
	0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72,
	0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x6e, 0x6c, 0x79, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0xec, 0x01, 0x0a, 0x1f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44, 0x52,
	0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x2d, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0xb2, 0x01, 0x0a, 0x21, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6f,
	0x6e, 0x6c, 0x79, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61,
	0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6f, 0x6e, 0x6c, 0x79, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x50, 0x61, 0x69, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x19,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49, 0x44, 0x52, 0x0e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x42, 0x0a,
	0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x49,
	0x44, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x32,
	0xf1, 0x05, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x50, 0x61, 0x69, 0x72, 0x2e, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x6b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x3a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x6b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x3a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6f, 0x0a,
	0x0e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x3c, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5d,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x69,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x48, 0x0a,
	0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x50, 0x61, 0x69, 0x72, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x1e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69,
	0x72, 0x2e, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x50, 0x61, 0x69,
	0x72, 0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_biconom_admin_currency_v1_currency_pair_proto_rawDescOnce sync.Once
	file_biconom_admin_currency_v1_currency_pair_proto_rawDescData = file_biconom_admin_currency_v1_currency_pair_proto_rawDesc
)

func file_biconom_admin_currency_v1_currency_pair_proto_rawDescGZIP() []byte {
	file_biconom_admin_currency_v1_currency_pair_proto_rawDescOnce.Do(func() {
		file_biconom_admin_currency_v1_currency_pair_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_currency_v1_currency_pair_proto_rawDescData)
	})
	return file_biconom_admin_currency_v1_currency_pair_proto_rawDescData
}

var file_biconom_admin_currency_v1_currency_pair_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_biconom_admin_currency_v1_currency_pair_proto_goTypes = []interface{}{
	(*CurrencyPairListRequest)(nil),           // 0: biconom.admin.currency.v1.CurrencyPairListRequest
	(*CurrencyPairListBySourceRequest)(nil),   // 1: biconom.admin.currency.v1.CurrencyPairListBySourceRequest
	(*CurrencyPairListByTargetRequest)(nil),   // 2: biconom.admin.currency.v1.CurrencyPairListByTargetRequest
	(*CurrencyPairListByCurrencyRequest)(nil), // 3: biconom.admin.currency.v1.CurrencyPairListByCurrencyRequest
	(*CurrencyPairCreateRequest)(nil),         // 4: biconom.admin.currency.v1.CurrencyPairCreateRequest
	(*currency.Currency_Pair_ID)(nil),         // 5: biconom.type.Currency.Pair.ID
	(*sort.Sort)(nil),                         // 6: biconom.type.Sort
	(*currency.Currency_ID)(nil),              // 7: biconom.type.Currency.ID
	(*currency.Currency_Pair)(nil),            // 8: biconom.type.Currency.Pair
}
var file_biconom_admin_currency_v1_currency_pair_proto_depIdxs = []int32{
	5,  // 0: biconom.admin.currency.v1.CurrencyPairListRequest.step:type_name -> biconom.type.Currency.Pair.ID
	6,  // 1: biconom.admin.currency.v1.CurrencyPairListRequest.sort:type_name -> biconom.type.Sort
	7,  // 2: biconom.admin.currency.v1.CurrencyPairListBySourceRequest.source_currency:type_name -> biconom.type.Currency.ID
	7,  // 3: biconom.admin.currency.v1.CurrencyPairListBySourceRequest.step:type_name -> biconom.type.Currency.ID
	6,  // 4: biconom.admin.currency.v1.CurrencyPairListBySourceRequest.sort:type_name -> biconom.type.Sort
	7,  // 5: biconom.admin.currency.v1.CurrencyPairListByTargetRequest.target_currency:type_name -> biconom.type.Currency.ID
	7,  // 6: biconom.admin.currency.v1.CurrencyPairListByTargetRequest.step:type_name -> biconom.type.Currency.ID
	6,  // 7: biconom.admin.currency.v1.CurrencyPairListByTargetRequest.sort:type_name -> biconom.type.Sort
	7,  // 8: biconom.admin.currency.v1.CurrencyPairListByCurrencyRequest.currency:type_name -> biconom.type.Currency.ID
	6,  // 9: biconom.admin.currency.v1.CurrencyPairListByCurrencyRequest.sort:type_name -> biconom.type.Sort
	7,  // 10: biconom.admin.currency.v1.CurrencyPairCreateRequest.source_currency:type_name -> biconom.type.Currency.ID
	7,  // 11: biconom.admin.currency.v1.CurrencyPairCreateRequest.target_currency:type_name -> biconom.type.Currency.ID
	5,  // 12: biconom.admin.currency.v1.CurrencyPair.Get:input_type -> biconom.type.Currency.Pair.ID
	0,  // 13: biconom.admin.currency.v1.CurrencyPair.List:input_type -> biconom.admin.currency.v1.CurrencyPairListRequest
	1,  // 14: biconom.admin.currency.v1.CurrencyPair.ListBySource:input_type -> biconom.admin.currency.v1.CurrencyPairListBySourceRequest
	2,  // 15: biconom.admin.currency.v1.CurrencyPair.ListByTarget:input_type -> biconom.admin.currency.v1.CurrencyPairListByTargetRequest
	3,  // 16: biconom.admin.currency.v1.CurrencyPair.ListByCurrency:input_type -> biconom.admin.currency.v1.CurrencyPairListByCurrencyRequest
	4,  // 17: biconom.admin.currency.v1.CurrencyPair.Create:input_type -> biconom.admin.currency.v1.CurrencyPairCreateRequest
	5,  // 18: biconom.admin.currency.v1.CurrencyPair.Publish:input_type -> biconom.type.Currency.Pair.ID
	5,  // 19: biconom.admin.currency.v1.CurrencyPair.Unpublish:input_type -> biconom.type.Currency.Pair.ID
	8,  // 20: biconom.admin.currency.v1.CurrencyPair.Get:output_type -> biconom.type.Currency.Pair
	8,  // 21: biconom.admin.currency.v1.CurrencyPair.List:output_type -> biconom.type.Currency.Pair
	8,  // 22: biconom.admin.currency.v1.CurrencyPair.ListBySource:output_type -> biconom.type.Currency.Pair
	8,  // 23: biconom.admin.currency.v1.CurrencyPair.ListByTarget:output_type -> biconom.type.Currency.Pair
	8,  // 24: biconom.admin.currency.v1.CurrencyPair.ListByCurrency:output_type -> biconom.type.Currency.Pair
	8,  // 25: biconom.admin.currency.v1.CurrencyPair.Create:output_type -> biconom.type.Currency.Pair
	8,  // 26: biconom.admin.currency.v1.CurrencyPair.Publish:output_type -> biconom.type.Currency.Pair
	8,  // 27: biconom.admin.currency.v1.CurrencyPair.Unpublish:output_type -> biconom.type.Currency.Pair
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_biconom_admin_currency_v1_currency_pair_proto_init() }
func file_biconom_admin_currency_v1_currency_pair_proto_init() {
	if File_biconom_admin_currency_v1_currency_pair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairListRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairListBySourceRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairListByTargetRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairListByCurrencyRequest); i {
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
		file_biconom_admin_currency_v1_currency_pair_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyPairCreateRequest); i {
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
			RawDescriptor: file_biconom_admin_currency_v1_currency_pair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_currency_v1_currency_pair_proto_goTypes,
		DependencyIndexes: file_biconom_admin_currency_v1_currency_pair_proto_depIdxs,
		MessageInfos:      file_biconom_admin_currency_v1_currency_pair_proto_msgTypes,
	}.Build()
	File_biconom_admin_currency_v1_currency_pair_proto = out.File
	file_biconom_admin_currency_v1_currency_pair_proto_rawDesc = nil
	file_biconom_admin_currency_v1_currency_pair_proto_goTypes = nil
	file_biconom_admin_currency_v1_currency_pair_proto_depIdxs = nil
}