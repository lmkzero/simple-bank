// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: bank.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner    string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *CreateAccountReq) Reset() {
	*x = CreateAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReq) ProtoMessage() {}

func (x *CreateAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountReq.ProtoReflect.Descriptor instead.
func (*CreateAccountReq) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateAccountReq) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type CreateAccountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAccount *Account `protobuf:"bytes,1,opt,name=created_account,json=createdAccount,proto3" json:"created_account,omitempty"`
}

func (x *CreateAccountRsp) Reset() {
	*x = CreateAccountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRsp) ProtoMessage() {}

func (x *CreateAccountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRsp.ProtoReflect.Descriptor instead.
func (*CreateAccountRsp) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountRsp) GetCreatedAccount() *Account {
	if x != nil {
		return x.CreatedAccount
	}
	return nil
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string                 `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Balance  int64                  `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	CreateAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
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
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Account) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Account) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Account) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

type GetAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" uri:"id"` // @gotags: uri:"id"
}

func (x *GetAccountReq) Reset() {
	*x = GetAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReq) ProtoMessage() {}

func (x *GetAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReq.ProtoReflect.Descriptor instead.
func (*GetAccountReq) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAccountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetAccountRsp) Reset() {
	*x = GetAccountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRsp) ProtoMessage() {}

func (x *GetAccountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRsp.ProtoReflect.Descriptor instead.
func (*GetAccountRsp) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccountRsp) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type ListAccountsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty" form:"limit"`   // @gotags: form:"limit"
	Offset uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty" form:"offset"` // @gotags: form:"offset"
}

func (x *ListAccountsReq) Reset() {
	*x = ListAccountsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsReq) ProtoMessage() {}

func (x *ListAccountsReq) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsReq.ProtoReflect.Descriptor instead.
func (*ListAccountsReq) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{5}
}

func (x *ListAccountsReq) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAccountsReq) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListAccountsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ListAccountsRsp) Reset() {
	*x = ListAccountsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccountsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccountsRsp) ProtoMessage() {}

func (x *ListAccountsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccountsRsp.ProtoReflect.Descriptor instead.
func (*ListAccountsRsp) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{6}
}

func (x *ListAccountsRsp) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type TransferReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountId int64  `protobuf:"varint,1,opt,name=from_account_id,json=fromAccountId,proto3" json:"from_account_id,omitempty"`
	ToAccountId   int64  `protobuf:"varint,2,opt,name=to_account_id,json=toAccountId,proto3" json:"to_account_id,omitempty"`
	Amount        int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *TransferReq) Reset() {
	*x = TransferReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferReq) ProtoMessage() {}

func (x *TransferReq) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferReq.ProtoReflect.Descriptor instead.
func (*TransferReq) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{7}
}

func (x *TransferReq) GetFromAccountId() int64 {
	if x != nil {
		return x.FromAccountId
	}
	return 0
}

func (x *TransferReq) GetToAccountId() int64 {
	if x != nil {
		return x.ToAccountId
	}
	return 0
}

func (x *TransferReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferReq) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type TransferRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccount *Account `protobuf:"bytes,1,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	ToAccount   *Account `protobuf:"bytes,2,opt,name=to_account,json=toAccount,proto3" json:"to_account,omitempty"`
}

func (x *TransferRsp) Reset() {
	*x = TransferRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRsp) ProtoMessage() {}

func (x *TransferRsp) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRsp.ProtoReflect.Descriptor instead.
func (*TransferRsp) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{8}
}

func (x *TransferRsp) GetFromAccount() *Account {
	if x != nil {
		return x.FromAccount
	}
	return nil
}

func (x *TransferRsp) GetToAccount() *Account {
	if x != nil {
		return x.ToAccount
	}
	return nil
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x1d, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x52, 0x03, 0x55, 0x53, 0x44, 0x52, 0x03,
	0x45, 0x55, 0x52, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x4d, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73,
	0x70, 0x12, 0x39, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x28, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x0a, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x2f, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0d, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x0b, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2b, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x52, 0x03, 0x55, 0x53, 0x44, 0x52, 0x03,
	0x45, 0x55, 0x52, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x73, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0c,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xd8, 0x02, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x58, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x73, 0x70, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x3a, 0x69, 0x64, 0x3a, 0x12, 0x55, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x73, 0x70, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x49, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x14,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0b, 0x22, 0x09, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x2f, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6d, 0x6b, 0x7a,
	0x65, 0x72, 0x6f, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x62, 0x61, 0x6e, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bank_proto_goTypes = []any{
	(*CreateAccountReq)(nil),      // 0: bank.v1.CreateAccountReq
	(*CreateAccountRsp)(nil),      // 1: bank.v1.CreateAccountRsp
	(*Account)(nil),               // 2: bank.v1.Account
	(*GetAccountReq)(nil),         // 3: bank.v1.GetAccountReq
	(*GetAccountRsp)(nil),         // 4: bank.v1.GetAccountRsp
	(*ListAccountsReq)(nil),       // 5: bank.v1.ListAccountsReq
	(*ListAccountsRsp)(nil),       // 6: bank.v1.ListAccountsRsp
	(*TransferReq)(nil),           // 7: bank.v1.TransferReq
	(*TransferRsp)(nil),           // 8: bank.v1.TransferRsp
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_bank_proto_depIdxs = []int32{
	2,  // 0: bank.v1.CreateAccountRsp.created_account:type_name -> bank.v1.Account
	9,  // 1: bank.v1.Account.create_at:type_name -> google.protobuf.Timestamp
	2,  // 2: bank.v1.GetAccountRsp.account:type_name -> bank.v1.Account
	2,  // 3: bank.v1.ListAccountsRsp.accounts:type_name -> bank.v1.Account
	2,  // 4: bank.v1.TransferRsp.from_account:type_name -> bank.v1.Account
	2,  // 5: bank.v1.TransferRsp.to_account:type_name -> bank.v1.Account
	0,  // 6: bank.v1.Bank.CreateAccount:input_type -> bank.v1.CreateAccountReq
	3,  // 7: bank.v1.Bank.GetAccount:input_type -> bank.v1.GetAccountReq
	5,  // 8: bank.v1.Bank.ListAccounts:input_type -> bank.v1.ListAccountsReq
	7,  // 9: bank.v1.Bank.Transfer:input_type -> bank.v1.TransferReq
	1,  // 10: bank.v1.Bank.CreateAccount:output_type -> bank.v1.CreateAccountRsp
	4,  // 11: bank.v1.Bank.GetAccount:output_type -> bank.v1.GetAccountRsp
	6,  // 12: bank.v1.Bank.ListAccounts:output_type -> bank.v1.ListAccountsRsp
	8,  // 13: bank.v1.Bank.Transfer:output_type -> bank.v1.TransferRsp
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAccountReq); i {
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
		file_bank_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAccountRsp); i {
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
		file_bank_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_bank_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountReq); i {
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
		file_bank_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountRsp); i {
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
		file_bank_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListAccountsReq); i {
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
		file_bank_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListAccountsRsp); i {
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
		file_bank_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TransferReq); i {
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
		file_bank_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*TransferRsp); i {
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
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}
