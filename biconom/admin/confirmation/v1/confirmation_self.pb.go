// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: biconom/admin/confirmation/v1/confirmation_self.proto

package service_admin_confirmation_pb

import (
	agent "github.com/biconom/go-genproto/biconom/type/agent"
	confirmation "github.com/biconom/go-genproto/biconom/type/confirmation"
	sort "github.com/biconom/go-genproto/biconom/type/sort"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type ConfirmationSelfListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step *confirmation.Confirmation_ID `protobuf:"bytes,1,opt,name=step,proto3" json:"step,omitempty"`
	Sort *sort.Sort                    `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ConfirmationSelfListRequest) Reset() {
	*x = ConfirmationSelfListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationSelfListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationSelfListRequest) ProtoMessage() {}

func (x *ConfirmationSelfListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationSelfListRequest.ProtoReflect.Descriptor instead.
func (*ConfirmationSelfListRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP(), []int{0}
}

func (x *ConfirmationSelfListRequest) GetStep() *confirmation.Confirmation_ID {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *ConfirmationSelfListRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type ConfirmationSelfLogListByConfirmationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmation *confirmation.Confirmation_ID           `protobuf:"bytes,1,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
	Step         *confirmation.Confirmation_Log_ID_Inner `protobuf:"bytes,2,opt,name=step,proto3" json:"step,omitempty"`
	Sort         *sort.Sort                              `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *ConfirmationSelfLogListByConfirmationRequest) Reset() {
	*x = ConfirmationSelfLogListByConfirmationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationSelfLogListByConfirmationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationSelfLogListByConfirmationRequest) ProtoMessage() {}

func (x *ConfirmationSelfLogListByConfirmationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationSelfLogListByConfirmationRequest.ProtoReflect.Descriptor instead.
func (*ConfirmationSelfLogListByConfirmationRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmationSelfLogListByConfirmationRequest) GetConfirmation() *confirmation.Confirmation_ID {
	if x != nil {
		return x.Confirmation
	}
	return nil
}

func (x *ConfirmationSelfLogListByConfirmationRequest) GetStep() *confirmation.Confirmation_Log_ID_Inner {
	if x != nil {
		return x.Step
	}
	return nil
}

func (x *ConfirmationSelfLogListByConfirmationRequest) GetSort() *sort.Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

type ConfirmationSelfGenerateOneTimePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent        *agent.Agent                  `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`                     // ???????????? ???? ??????????-???????????? ?? ???????????????? ?????????????????????? ????????????
	Confirmation *confirmation.Confirmation_ID `protobuf:"bytes,2,opt,name=confirmation,proto3" json:"confirmation,omitempty"`       // ?????????????????????????? ?????????? ??????????????????????????
	FieldId      uint32                        `protobuf:"varint,3,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"` // ?????????????????????????? ????????
}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) Reset() {
	*x = ConfirmationSelfGenerateOneTimePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationSelfGenerateOneTimePasswordRequest) ProtoMessage() {}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationSelfGenerateOneTimePasswordRequest.ProtoReflect.Descriptor instead.
func (*ConfirmationSelfGenerateOneTimePasswordRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) GetAgent() *agent.Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) GetConfirmation() *confirmation.Confirmation_ID {
	if x != nil {
		return x.Confirmation
	}
	return nil
}

func (x *ConfirmationSelfGenerateOneTimePasswordRequest) GetFieldId() uint32 {
	if x != nil {
		return x.FieldId
	}
	return 0
}

type ConfirmationSelfConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent        *agent.Agent                  `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`                                                                                                                         // ???????????? ???? ??????????-???????????? ?? ???????????????? ?????????????????????? ????????????
	Confirmation *confirmation.Confirmation_ID `protobuf:"bytes,2,opt,name=confirmation,proto3" json:"confirmation,omitempty"`                                                                                                           // ?????????????????????????? ?????????? ??????????????????????????
	FieldValues  map[uint32]string             `protobuf:"bytes,3,rep,name=field_values,json=fieldValues,proto3" json:"field_values,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map key is "field_id"
}

func (x *ConfirmationSelfConfirmRequest) Reset() {
	*x = ConfirmationSelfConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationSelfConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationSelfConfirmRequest) ProtoMessage() {}

func (x *ConfirmationSelfConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationSelfConfirmRequest.ProtoReflect.Descriptor instead.
func (*ConfirmationSelfConfirmRequest) Descriptor() ([]byte, []int) {
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP(), []int{3}
}

func (x *ConfirmationSelfConfirmRequest) GetAgent() *agent.Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

func (x *ConfirmationSelfConfirmRequest) GetConfirmation() *confirmation.Confirmation_ID {
	if x != nil {
		return x.Confirmation
	}
	return nil
}

func (x *ConfirmationSelfConfirmRequest) GetFieldValues() map[uint32]string {
	if x != nil {
		return x.FieldValues
	}
	return nil
}

type ConfirmationSelfConfirmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool                                               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	FieldStatuses map[uint32]*confirmation.Confirmation_Field_Status `protobuf:"bytes,2,rep,name=field_statuses,json=fieldStatuses,proto3" json:"field_statuses,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map key is "field_id"
	Metadata      *anypb.Any                                         `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *ConfirmationSelfConfirmResponse) Reset() {
	*x = ConfirmationSelfConfirmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmationSelfConfirmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmationSelfConfirmResponse) ProtoMessage() {}

func (x *ConfirmationSelfConfirmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmationSelfConfirmResponse.ProtoReflect.Descriptor instead.
func (*ConfirmationSelfConfirmResponse) Descriptor() ([]byte, []int) {
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP(), []int{4}
}

func (x *ConfirmationSelfConfirmResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ConfirmationSelfConfirmResponse) GetFieldStatuses() map[uint32]*confirmation.Confirmation_Field_Status {
	if x != nil {
		return x.FieldStatuses
	}
	return nil
}

func (x *ConfirmationSelfConfirmResponse) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_biconom_admin_confirmation_v1_confirmation_self_proto protoreflect.FileDescriptor

var file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDesc = []byte{
	0x0a, 0x35, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6c,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x73, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x78, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x52, 0x04,
	0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0xd6, 0x01, 0x0a,
	0x2c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c,
	0x66, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3b, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x49,
	0x44, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0xb9, 0x01, 0x0a, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49,
	0x64, 0x22, 0xc1, 0x02, 0x0a, 0x1e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x44, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd2, 0x02, 0x0a, 0x1f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x78, 0x0a, 0x0e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0x69, 0x0a, 0x12, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x80, 0x06, 0x0a, 0x10, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x12,
	0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x49, 0x44, 0x1a, 0x1a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x2e, 0x62, 0x69,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x47, 0x65,
	0x74, 0x12, 0x21, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f,
	0x67, 0x2e, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x88, 0x01, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4b, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c,
	0x66, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x99, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x6e,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x4d, 0x2e,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62,
	0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x8a, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x3d, 0x2e, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e,
	0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x12, 0x1d, 0x2e, 0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x5c, 0x5a,
	0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x63, 0x6f,
	0x6e, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x69, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescOnce sync.Once
	file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescData = file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDesc
)

func file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescGZIP() []byte {
	file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescOnce.Do(func() {
		file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescData = protoimpl.X.CompressGZIP(file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescData)
	})
	return file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDescData
}

var file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_biconom_admin_confirmation_v1_confirmation_self_proto_goTypes = []interface{}{
	(*ConfirmationSelfListRequest)(nil),                    // 0: biconom.admin.confirmation.v1.ConfirmationSelfListRequest
	(*ConfirmationSelfLogListByConfirmationRequest)(nil),   // 1: biconom.admin.confirmation.v1.ConfirmationSelfLogListByConfirmationRequest
	(*ConfirmationSelfGenerateOneTimePasswordRequest)(nil), // 2: biconom.admin.confirmation.v1.ConfirmationSelfGenerateOneTimePasswordRequest
	(*ConfirmationSelfConfirmRequest)(nil),                 // 3: biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest
	(*ConfirmationSelfConfirmResponse)(nil),                // 4: biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse
	nil,                                                    // 5: biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest.FieldValuesEntry
	nil,                                                    // 6: biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse.FieldStatusesEntry
	(*confirmation.Confirmation_ID)(nil),                   // 7: biconom.type.Confirmation.ID
	(*sort.Sort)(nil),                                      // 8: biconom.type.Sort
	(*confirmation.Confirmation_Log_ID_Inner)(nil),         // 9: biconom.type.Confirmation.Log.ID.Inner
	(*agent.Agent)(nil),                                    // 10: biconom.type.Agent
	(*anypb.Any)(nil),                                      // 11: google.protobuf.Any
	(*confirmation.Confirmation_Field_Status)(nil),         // 12: biconom.type.Confirmation.Field.Status
	(*confirmation.Confirmation_Log_ID)(nil),               // 13: biconom.type.Confirmation.Log.ID
	(*confirmation.Confirmation)(nil),                      // 14: biconom.type.Confirmation
	(*confirmation.Confirmation_Log)(nil),                  // 15: biconom.type.Confirmation.Log
	(*confirmation.Confirmation_Field_Notification)(nil),   // 16: biconom.type.Confirmation.Field.Notification
	(*emptypb.Empty)(nil),                                  // 17: google.protobuf.Empty
}
var file_biconom_admin_confirmation_v1_confirmation_self_proto_depIdxs = []int32{
	7,  // 0: biconom.admin.confirmation.v1.ConfirmationSelfListRequest.step:type_name -> biconom.type.Confirmation.ID
	8,  // 1: biconom.admin.confirmation.v1.ConfirmationSelfListRequest.sort:type_name -> biconom.type.Sort
	7,  // 2: biconom.admin.confirmation.v1.ConfirmationSelfLogListByConfirmationRequest.confirmation:type_name -> biconom.type.Confirmation.ID
	9,  // 3: biconom.admin.confirmation.v1.ConfirmationSelfLogListByConfirmationRequest.step:type_name -> biconom.type.Confirmation.Log.ID.Inner
	8,  // 4: biconom.admin.confirmation.v1.ConfirmationSelfLogListByConfirmationRequest.sort:type_name -> biconom.type.Sort
	10, // 5: biconom.admin.confirmation.v1.ConfirmationSelfGenerateOneTimePasswordRequest.agent:type_name -> biconom.type.Agent
	7,  // 6: biconom.admin.confirmation.v1.ConfirmationSelfGenerateOneTimePasswordRequest.confirmation:type_name -> biconom.type.Confirmation.ID
	10, // 7: biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest.agent:type_name -> biconom.type.Agent
	7,  // 8: biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest.confirmation:type_name -> biconom.type.Confirmation.ID
	5,  // 9: biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest.field_values:type_name -> biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest.FieldValuesEntry
	6,  // 10: biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse.field_statuses:type_name -> biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse.FieldStatusesEntry
	11, // 11: biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse.metadata:type_name -> google.protobuf.Any
	12, // 12: biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse.FieldStatusesEntry.value:type_name -> biconom.type.Confirmation.Field.Status
	7,  // 13: biconom.admin.confirmation.v1.ConfirmationSelf.Get:input_type -> biconom.type.Confirmation.ID
	0,  // 14: biconom.admin.confirmation.v1.ConfirmationSelf.List:input_type -> biconom.admin.confirmation.v1.ConfirmationSelfListRequest
	13, // 15: biconom.admin.confirmation.v1.ConfirmationSelf.LogGet:input_type -> biconom.type.Confirmation.Log.ID
	1,  // 16: biconom.admin.confirmation.v1.ConfirmationSelf.LogListByConfirmation:input_type -> biconom.admin.confirmation.v1.ConfirmationSelfLogListByConfirmationRequest
	2,  // 17: biconom.admin.confirmation.v1.ConfirmationSelf.GenerateOneTimePassword:input_type -> biconom.admin.confirmation.v1.ConfirmationSelfGenerateOneTimePasswordRequest
	3,  // 18: biconom.admin.confirmation.v1.ConfirmationSelf.Confirm:input_type -> biconom.admin.confirmation.v1.ConfirmationSelfConfirmRequest
	7,  // 19: biconom.admin.confirmation.v1.ConfirmationSelf.Cancel:input_type -> biconom.type.Confirmation.ID
	14, // 20: biconom.admin.confirmation.v1.ConfirmationSelf.Get:output_type -> biconom.type.Confirmation
	14, // 21: biconom.admin.confirmation.v1.ConfirmationSelf.List:output_type -> biconom.type.Confirmation
	15, // 22: biconom.admin.confirmation.v1.ConfirmationSelf.LogGet:output_type -> biconom.type.Confirmation.Log
	15, // 23: biconom.admin.confirmation.v1.ConfirmationSelf.LogListByConfirmation:output_type -> biconom.type.Confirmation.Log
	16, // 24: biconom.admin.confirmation.v1.ConfirmationSelf.GenerateOneTimePassword:output_type -> biconom.type.Confirmation.Field.Notification
	4,  // 25: biconom.admin.confirmation.v1.ConfirmationSelf.Confirm:output_type -> biconom.admin.confirmation.v1.ConfirmationSelfConfirmResponse
	17, // 26: biconom.admin.confirmation.v1.ConfirmationSelf.Cancel:output_type -> google.protobuf.Empty
	20, // [20:27] is the sub-list for method output_type
	13, // [13:20] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_biconom_admin_confirmation_v1_confirmation_self_proto_init() }
func file_biconom_admin_confirmation_v1_confirmation_self_proto_init() {
	if File_biconom_admin_confirmation_v1_confirmation_self_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationSelfListRequest); i {
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
		file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationSelfLogListByConfirmationRequest); i {
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
		file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationSelfGenerateOneTimePasswordRequest); i {
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
		file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationSelfConfirmRequest); i {
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
		file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmationSelfConfirmResponse); i {
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
			RawDescriptor: file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_biconom_admin_confirmation_v1_confirmation_self_proto_goTypes,
		DependencyIndexes: file_biconom_admin_confirmation_v1_confirmation_self_proto_depIdxs,
		MessageInfos:      file_biconom_admin_confirmation_v1_confirmation_self_proto_msgTypes,
	}.Build()
	File_biconom_admin_confirmation_v1_confirmation_self_proto = out.File
	file_biconom_admin_confirmation_v1_confirmation_self_proto_rawDesc = nil
	file_biconom_admin_confirmation_v1_confirmation_self_proto_goTypes = nil
	file_biconom_admin_confirmation_v1_confirmation_self_proto_depIdxs = nil
}
