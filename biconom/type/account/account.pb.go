// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/type/account.proto

package account_pb

import (
	condition "github.com/biconom/go-genproto/biconom/type/condition"
	contact "github.com/biconom/go-genproto/biconom/type/contact"
	profile "github.com/biconom/go-genproto/biconom/type/profile"
	rank "github.com/biconom/go-genproto/biconom/type/rank"
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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId    uint32                       `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountLogin string                       `protobuf:"bytes,2,opt,name=account_login,json=accountLogin,proto3" json:"account_login,omitempty"`
	Profile      *profile.Profile             `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	Contacts     []*Account_WrappedContact    `protobuf:"bytes,4,rep,name=contacts,proto3" json:"contacts,omitempty"`
	RankBuckets  []*Account_WrappedRankBucket `protobuf:"bytes,5,rep,name=rank_buckets,json=rankBuckets,proto3" json:"rank_buckets,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetAccountId() uint32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Account) GetAccountLogin() string {
	if x != nil {
		return x.AccountLogin
	}
	return ""
}

func (x *Account) GetProfile() *profile.Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *Account) GetContacts() []*Account_WrappedContact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

func (x *Account) GetRankBuckets() []*Account_WrappedRankBucket {
	if x != nil {
		return x.RankBuckets
	}
	return nil
}

type Account_Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UniqueField:
	//	*Account_Id_AccountId
	//	*Account_Id_AccountLogin
	UniqueField isAccount_Id_UniqueField `protobuf_oneof:"unique_field"`
}

func (x *Account_Id) Reset() {
	*x = Account_Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_Id) ProtoMessage() {}

func (x *Account_Id) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_Id.ProtoReflect.Descriptor instead.
func (*Account_Id) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Account_Id) GetUniqueField() isAccount_Id_UniqueField {
	if m != nil {
		return m.UniqueField
	}
	return nil
}

func (x *Account_Id) GetAccountId() uint32 {
	if x, ok := x.GetUniqueField().(*Account_Id_AccountId); ok {
		return x.AccountId
	}
	return 0
}

func (x *Account_Id) GetAccountLogin() string {
	if x, ok := x.GetUniqueField().(*Account_Id_AccountLogin); ok {
		return x.AccountLogin
	}
	return ""
}

type isAccount_Id_UniqueField interface {
	isAccount_Id_UniqueField()
}

type Account_Id_AccountId struct {
	AccountId uint32 `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3,oneof"`
}

type Account_Id_AccountLogin struct {
	AccountLogin string `protobuf:"bytes,2,opt,name=account_login,json=accountLogin,proto3,oneof"`
}

func (*Account_Id_AccountId) isAccount_Id_UniqueField() {}

func (*Account_Id_AccountLogin) isAccount_Id_UniqueField() {}

type Account_WrappedContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contact     *contact.Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Confirmed   bool             `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	CreatedAt   int64            `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ConfirmedAt int64            `protobuf:"varint,4,opt,name=confirmed_at,json=confirmedAt,proto3" json:"confirmed_at,omitempty"`
}

func (x *Account_WrappedContact) Reset() {
	*x = Account_WrappedContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_WrappedContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_WrappedContact) ProtoMessage() {}

func (x *Account_WrappedContact) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_WrappedContact.ProtoReflect.Descriptor instead.
func (*Account_WrappedContact) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Account_WrappedContact) GetContact() *contact.Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Account_WrappedContact) GetConfirmed() bool {
	if x != nil {
		return x.Confirmed
	}
	return false
}

func (x *Account_WrappedContact) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Account_WrappedContact) GetConfirmedAt() int64 {
	if x != nil {
		return x.ConfirmedAt
	}
	return 0
}

type Account_WrappedRank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank        *rank.Rank `protobuf:"bytes,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Activated   bool       `protobuf:"varint,2,opt,name=activated,proto3" json:"activated,omitempty"`
	ActivatedAt int64      `protobuf:"varint,3,opt,name=activated_at,json=activatedAt,proto3" json:"activated_at,omitempty"`
}

func (x *Account_WrappedRank) Reset() {
	*x = Account_WrappedRank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_WrappedRank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_WrappedRank) ProtoMessage() {}

func (x *Account_WrappedRank) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_WrappedRank.ProtoReflect.Descriptor instead.
func (*Account_WrappedRank) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Account_WrappedRank) GetRank() *rank.Rank {
	if x != nil {
		return x.Rank
	}
	return nil
}

func (x *Account_WrappedRank) GetActivated() bool {
	if x != nil {
		return x.Activated
	}
	return false
}

func (x *Account_WrappedRank) GetActivatedAt() int64 {
	if x != nil {
		return x.ActivatedAt
	}
	return 0
}

type Account_WrappedRankBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *rank.RankBucket_Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Ranks  []*Account_WrappedRank  `protobuf:"bytes,2,rep,name=ranks,proto3" json:"ranks,omitempty"`
}

func (x *Account_WrappedRankBucket) Reset() {
	*x = Account_WrappedRankBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_WrappedRankBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_WrappedRankBucket) ProtoMessage() {}

func (x *Account_WrappedRankBucket) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_WrappedRankBucket.ProtoReflect.Descriptor instead.
func (*Account_WrappedRankBucket) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Account_WrappedRankBucket) GetHeader() *rank.RankBucket_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Account_WrappedRankBucket) GetRanks() []*Account_WrappedRank {
	if x != nil {
		return x.Ranks
	}
	return nil
}

type Account_WrappedConditionBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *condition.ConditionBucket_Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Target *condition.ConditionBucket_Target `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Result bool                              `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Account_WrappedConditionBucket) Reset() {
	*x = Account_WrappedConditionBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_type_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account_WrappedConditionBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account_WrappedConditionBucket) ProtoMessage() {}

func (x *Account_WrappedConditionBucket) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_type_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account_WrappedConditionBucket.ProtoReflect.Descriptor instead.
func (*Account_WrappedConditionBucket) Descriptor() ([]byte, []int) {
	return file_biconom_type_account_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Account_WrappedConditionBucket) GetHeader() *condition.ConditionBucket_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Account_WrappedConditionBucket) GetTarget() *condition.ConditionBucket_Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *Account_WrappedConditionBucket) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_biconom_type_account_proto protoreflect.FileDescriptor

var file_biconom_type_account_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x1a, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x72,
	0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x07, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x4a, 0x0a, 0x0c,
	0x72, 0x61, 0x6e, 0x6b, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0b, 0x72, 0x61, 0x6e,
	0x6b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x1a, 0x5c, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0xa1, 0x01, 0x0a, 0x0e, 0x57, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x69, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x76, 0x0a, 0x0b, 0x57, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x1a, 0x85, 0x01, 0x0a, 0x11, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x52, 0x61,
	0x6e, 0x6b, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x37, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x52,
	0x61, 0x6e, 0x6b, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x6b, 0x73, 0x1a, 0xac, 0x01, 0x0a, 0x16, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_biconom_type_account_proto_rawDescOnce sync.Once
	file_biconom_type_account_proto_rawDescData = file_biconom_type_account_proto_rawDesc
)

func file_biconom_type_account_proto_rawDescGZIP() []byte {
	file_biconom_type_account_proto_rawDescOnce.Do(func() {
		file_biconom_type_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_type_account_proto_rawDescData)
	})
	return file_biconom_type_account_proto_rawDescData
}

var file_biconom_type_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_biconom_type_account_proto_goTypes = []interface{}{
	(*Account)(nil),                          // 0: biconom.type.Account
	(*Account_Id)(nil),                       // 1: biconom.type.Account.Id
	(*Account_WrappedContact)(nil),           // 2: biconom.type.Account.WrappedContact
	(*Account_WrappedRank)(nil),              // 3: biconom.type.Account.WrappedRank
	(*Account_WrappedRankBucket)(nil),        // 4: biconom.type.Account.WrappedRankBucket
	(*Account_WrappedConditionBucket)(nil),   // 5: biconom.type.Account.WrappedConditionBucket
	(*profile.Profile)(nil),                  // 6: biconom.type.Profile
	(*contact.Contact)(nil),                  // 7: biconom.type.Contact
	(*rank.Rank)(nil),                        // 8: biconom.type.Rank
	(*rank.RankBucket_Header)(nil),           // 9: biconom.type.RankBucket.Header
	(*condition.ConditionBucket_Header)(nil), // 10: biconom.type.ConditionBucket.Header
	(*condition.ConditionBucket_Target)(nil), // 11: biconom.type.ConditionBucket.Target
}
var file_biconom_type_account_proto_depIdxs = []int32{
	6,  // 0: biconom.type.Account.profile:type_name -> biconom.type.Profile
	2,  // 1: biconom.type.Account.contacts:type_name -> biconom.type.Account.WrappedContact
	4,  // 2: biconom.type.Account.rank_buckets:type_name -> biconom.type.Account.WrappedRankBucket
	7,  // 3: biconom.type.Account.WrappedContact.contact:type_name -> biconom.type.Contact
	8,  // 4: biconom.type.Account.WrappedRank.rank:type_name -> biconom.type.Rank
	9,  // 5: biconom.type.Account.WrappedRankBucket.header:type_name -> biconom.type.RankBucket.Header
	3,  // 6: biconom.type.Account.WrappedRankBucket.ranks:type_name -> biconom.type.Account.WrappedRank
	10, // 7: biconom.type.Account.WrappedConditionBucket.header:type_name -> biconom.type.ConditionBucket.Header
	11, // 8: biconom.type.Account.WrappedConditionBucket.target:type_name -> biconom.type.ConditionBucket.Target
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_biconom_type_account_proto_init() }
func file_biconom_type_account_proto_init() {
	if File_biconom_type_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_type_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_biconom_type_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_Id); i {
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
		file_biconom_type_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_WrappedContact); i {
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
		file_biconom_type_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_WrappedRank); i {
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
		file_biconom_type_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_WrappedRankBucket); i {
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
		file_biconom_type_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account_WrappedConditionBucket); i {
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
	file_biconom_type_account_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Account_Id_AccountId)(nil),
		(*Account_Id_AccountLogin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_biconom_type_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_biconom_type_account_proto_goTypes,
		DependencyIndexes: file_biconom_type_account_proto_depIdxs,
		MessageInfos:      file_biconom_type_account_proto_msgTypes,
	}.Build()
	File_biconom_type_account_proto = out.File
	file_biconom_type_account_proto_rawDesc = nil
	file_biconom_type_account_proto_goTypes = nil
	file_biconom_type_account_proto_depIdxs = nil
}
