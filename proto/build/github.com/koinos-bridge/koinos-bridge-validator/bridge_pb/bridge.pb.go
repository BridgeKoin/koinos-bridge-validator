// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: proto/bridge.proto

package bridge_pb

import (
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

type TransactionType int32

const (
	TransactionType_koinos   TransactionType = 0
	TransactionType_ethereum TransactionType = 1
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "koinos",
		1: "ethereum",
	}
	TransactionType_value = map[string]int32{
		"koinos":   0,
		"ethereum": 1,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bridge_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_proto_bridge_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{0}
}

type TransactionStatus int32

const (
	TransactionStatus_gathering_signatures TransactionStatus = 0
	TransactionStatus_signed               TransactionStatus = 1
	TransactionStatus_completed            TransactionStatus = 2
)

// Enum value maps for TransactionStatus.
var (
	TransactionStatus_name = map[int32]string{
		0: "gathering_signatures",
		1: "signed",
		2: "completed",
	}
	TransactionStatus_value = map[string]int32{
		"gathering_signatures": 0,
		"signed":               1,
		"completed":            2,
	}
)

func (x TransactionStatus) Enum() *TransactionStatus {
	p := new(TransactionStatus)
	*p = x
	return p
}

func (x TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bridge_proto_enumTypes[1].Descriptor()
}

func (TransactionStatus) Type() protoreflect.EnumType {
	return &file_proto_bridge_proto_enumTypes[1]
}

func (x TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionStatus.Descriptor instead.
func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{1}
}

type ActionId int32

const (
	ActionId_reserved_action                ActionId = 0
	ActionId_add_validator                  ActionId = 1
	ActionId_remove_validator               ActionId = 2
	ActionId_add_supported_token            ActionId = 3
	ActionId_remove_supported_token         ActionId = 4
	ActionId_add_supported_wrapped_token    ActionId = 5
	ActionId_remove_supported_wrapped_token ActionId = 6
	ActionId_set_pause                      ActionId = 7
	ActionId_complete_transfer              ActionId = 8
)

// Enum value maps for ActionId.
var (
	ActionId_name = map[int32]string{
		0: "reserved_action",
		1: "add_validator",
		2: "remove_validator",
		3: "add_supported_token",
		4: "remove_supported_token",
		5: "add_supported_wrapped_token",
		6: "remove_supported_wrapped_token",
		7: "set_pause",
		8: "complete_transfer",
	}
	ActionId_value = map[string]int32{
		"reserved_action":                0,
		"add_validator":                  1,
		"remove_validator":               2,
		"add_supported_token":            3,
		"remove_supported_token":         4,
		"add_supported_wrapped_token":    5,
		"remove_supported_wrapped_token": 6,
		"set_pause":                      7,
		"complete_transfer":              8,
	}
)

func (x ActionId) Enum() *ActionId {
	p := new(ActionId)
	*p = x
	return p
}

func (x ActionId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionId) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bridge_proto_enumTypes[2].Descriptor()
}

func (ActionId) Type() protoreflect.EnumType {
	return &file_proto_bridge_proto_enumTypes[2]
}

func (x ActionId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionId.Descriptor instead.
func (ActionId) EnumDescriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{2}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastEthereumBlockParsed uint64 `protobuf:"varint,1,opt,name=last_ethereum_block_parsed,json=lastEthereumBlockParsed,proto3" json:"last_ethereum_block_parsed,omitempty"`
	LastKoinosBlockParsed   uint64 `protobuf:"varint,2,opt,name=last_koinos_block_parsed,json=lastKoinosBlockParsed,proto3" json:"last_koinos_block_parsed,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetLastEthereumBlockParsed() uint64 {
	if x != nil {
		return x.LastEthereumBlockParsed
	}
	return 0
}

func (x *Metadata) GetLastKoinosBlockParsed() uint64 {
	if x != nil {
		return x.LastKoinosBlockParsed
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                    TransactionType   `protobuf:"varint,1,opt,name=type,proto3,enum=bridge.TransactionType" json:"type,omitempty"`
	Id                      string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	OpId                    string            `protobuf:"bytes,3,opt,name=op_id,json=opId,proto3" json:"op_id,omitempty"`
	BlockNumber             uint64            `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	BlockTime               uint64            `protobuf:"varint,5,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	Expiration              uint64            `protobuf:"varint,6,opt,name=expiration,proto3" json:"expiration,omitempty"`
	From                    string            `protobuf:"bytes,7,opt,name=from,proto3" json:"from,omitempty"`
	EthToken                string            `protobuf:"bytes,8,opt,name=eth_token,json=ethToken,proto3" json:"eth_token,omitempty"`
	KoinosToken             string            `protobuf:"bytes,9,opt,name=koinos_token,json=koinosToken,proto3" json:"koinos_token,omitempty"`
	Amount                  string            `protobuf:"bytes,10,opt,name=amount,proto3" json:"amount,omitempty"`
	Recipient               string            `protobuf:"bytes,11,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Hash                    string            `protobuf:"bytes,12,opt,name=hash,proto3" json:"hash,omitempty"`
	Validators              []string          `protobuf:"bytes,13,rep,name=validators,proto3" json:"validators,omitempty"`
	Signatures              []string          `protobuf:"bytes,14,rep,name=signatures,proto3" json:"signatures,omitempty"`
	Status                  TransactionStatus `protobuf:"varint,15,opt,name=status,proto3,enum=bridge.TransactionStatus" json:"status,omitempty"`
	CompletionTransactionId string            `protobuf:"bytes,16,opt,name=completion_transaction_id,json=completionTransactionId,proto3" json:"completion_transaction_id,omitempty"`
	ToChain                 string            `protobuf:"bytes,17,opt,name=to_chain,json=toChain,proto3" json:"to_chain,omitempty"`
	Metadata                string            `protobuf:"bytes,18,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Relayer                 string            `protobuf:"bytes,19,opt,name=relayer,proto3" json:"relayer,omitempty"`
	Payment                 string            `protobuf:"bytes,20,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_koinos
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetOpId() string {
	if x != nil {
		return x.OpId
	}
	return ""
}

func (x *Transaction) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Transaction) GetBlockTime() uint64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *Transaction) GetExpiration() uint64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

func (x *Transaction) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Transaction) GetEthToken() string {
	if x != nil {
		return x.EthToken
	}
	return ""
}

func (x *Transaction) GetKoinosToken() string {
	if x != nil {
		return x.KoinosToken
	}
	return ""
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Transaction) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Transaction) GetValidators() []string {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *Transaction) GetSignatures() []string {
	if x != nil {
		return x.Signatures
	}
	return nil
}

func (x *Transaction) GetStatus() TransactionStatus {
	if x != nil {
		return x.Status
	}
	return TransactionStatus_gathering_signatures
}

func (x *Transaction) GetCompletionTransactionId() string {
	if x != nil {
		return x.CompletionTransactionId
	}
	return ""
}

func (x *Transaction) GetToChain() string {
	if x != nil {
		return x.ToChain
	}
	return ""
}

func (x *Transaction) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Transaction) GetRelayer() string {
	if x != nil {
		return x.Relayer
	}
	return ""
}

func (x *Transaction) GetPayment() string {
	if x != nil {
		return x.Payment
	}
	return ""
}

type CompleteTransferHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action        ActionId `protobuf:"varint,1,opt,name=action,proto3,enum=bridge.ActionId" json:"action,omitempty"`
	TransactionId []byte   `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Token         []byte   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Relayer       []byte   `protobuf:"bytes,4,opt,name=relayer,proto3" json:"relayer,omitempty"`
	Recipient     []byte   `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Amount        uint64   `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Payment       uint64   `protobuf:"varint,7,opt,name=payment,proto3" json:"payment,omitempty"`
	Metadata      string   `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ContractId    []byte   `protobuf:"bytes,9,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Expiration    uint64   `protobuf:"varint,10,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Chain         uint32   `protobuf:"varint,11,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (x *CompleteTransferHash) Reset() {
	*x = CompleteTransferHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTransferHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTransferHash) ProtoMessage() {}

func (x *CompleteTransferHash) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTransferHash.ProtoReflect.Descriptor instead.
func (*CompleteTransferHash) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteTransferHash) GetAction() ActionId {
	if x != nil {
		return x.Action
	}
	return ActionId_reserved_action
}

func (x *CompleteTransferHash) GetTransactionId() []byte {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *CompleteTransferHash) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *CompleteTransferHash) GetRelayer() []byte {
	if x != nil {
		return x.Relayer
	}
	return nil
}

func (x *CompleteTransferHash) GetRecipient() []byte {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *CompleteTransferHash) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CompleteTransferHash) GetPayment() uint64 {
	if x != nil {
		return x.Payment
	}
	return 0
}

func (x *CompleteTransferHash) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *CompleteTransferHash) GetContractId() []byte {
	if x != nil {
		return x.ContractId
	}
	return nil
}

func (x *CompleteTransferHash) GetExpiration() uint64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

func (x *CompleteTransferHash) GetChain() uint32 {
	if x != nil {
		return x.Chain
	}
	return 0
}

type SubmittedSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Signature   string       `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Expiration  int64        `protobuf:"varint,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *SubmittedSignature) Reset() {
	*x = SubmittedSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmittedSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmittedSignature) ProtoMessage() {}

func (x *SubmittedSignature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmittedSignature.ProtoReflect.Descriptor instead.
func (*SubmittedSignature) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{3}
}

func (x *SubmittedSignature) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *SubmittedSignature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SubmittedSignature) GetExpiration() int64 {
	if x != nil {
		return x.Expiration
	}
	return 0
}

type TokensLockedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Token     []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Payment   string `protobuf:"bytes,4,opt,name=payment,proto3" json:"payment,omitempty"`
	Relayer   string `protobuf:"bytes,5,opt,name=relayer,proto3" json:"relayer,omitempty"`
	Recipient string `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Metadata  string `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *TokensLockedEvent) Reset() {
	*x = TokensLockedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensLockedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensLockedEvent) ProtoMessage() {}

func (x *TokensLockedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensLockedEvent.ProtoReflect.Descriptor instead.
func (*TokensLockedEvent) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{4}
}

func (x *TokensLockedEvent) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TokensLockedEvent) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokensLockedEvent) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TokensLockedEvent) GetPayment() string {
	if x != nil {
		return x.Payment
	}
	return ""
}

func (x *TokensLockedEvent) GetRelayer() string {
	if x != nil {
		return x.Relayer
	}
	return ""
}

func (x *TokensLockedEvent) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *TokensLockedEvent) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *TokensLockedEvent) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

type TransferCompletedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxId []byte `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
}

func (x *TransferCompletedEvent) Reset() {
	*x = TransferCompletedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferCompletedEvent) ProtoMessage() {}

func (x *TransferCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferCompletedEvent.ProtoReflect.Descriptor instead.
func (*TransferCompletedEvent) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{5}
}

func (x *TransferCompletedEvent) GetTxId() []byte {
	if x != nil {
		return x.TxId
	}
	return nil
}

type RequestNewSignaturesEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	OperationId   string `protobuf:"bytes,2,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *RequestNewSignaturesEvent) Reset() {
	*x = RequestNewSignaturesEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestNewSignaturesEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestNewSignaturesEvent) ProtoMessage() {}

func (x *RequestNewSignaturesEvent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestNewSignaturesEvent.ProtoReflect.Descriptor instead.
func (*RequestNewSignaturesEvent) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{6}
}

func (x *RequestNewSignaturesEvent) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *RequestNewSignaturesEvent) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

var File_proto_bridge_proto protoreflect.FileDescriptor

var file_proto_bridge_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x1a, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x6c,
	0x61, 0x73, 0x74, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x22,
	0xfb, 0x04, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x13, 0x0a,
	0x05, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x74, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xdd, 0x02,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x29, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x8a, 0x01,
	0x0a, 0x13, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe0, 0x01, 0x0a, 0x13, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69,
	0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x2f, 0x0a,
	0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x68,
	0x0a, 0x1c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x2c, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x10, 0x01, 0x2a, 0x49, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14,
	0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10,
	0x02, 0x2a, 0xe9, 0x01, 0x0a, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x12,
	0x13, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x10, 0x05, 0x12, 0x22, 0x0a, 0x1e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x5f, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x08, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2d, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_bridge_proto_rawDescOnce sync.Once
	file_proto_bridge_proto_rawDescData = file_proto_bridge_proto_rawDesc
)

func file_proto_bridge_proto_rawDescGZIP() []byte {
	file_proto_bridge_proto_rawDescOnce.Do(func() {
		file_proto_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bridge_proto_rawDescData)
	})
	return file_proto_bridge_proto_rawDescData
}

var file_proto_bridge_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_bridge_proto_goTypes = []interface{}{
	(TransactionType)(0),              // 0: bridge.transaction_type
	(TransactionStatus)(0),            // 1: bridge.transaction_status
	(ActionId)(0),                     // 2: bridge.action_id
	(*Metadata)(nil),                  // 3: bridge.metadata
	(*Transaction)(nil),               // 4: bridge.transaction
	(*CompleteTransferHash)(nil),      // 5: bridge.complete_transfer_hash
	(*SubmittedSignature)(nil),        // 6: bridge.submitted_signature
	(*TokensLockedEvent)(nil),         // 7: bridge.tokens_locked_event
	(*TransferCompletedEvent)(nil),    // 8: bridge.transfer_completed_event
	(*RequestNewSignaturesEvent)(nil), // 9: bridge.request_new_signatures_event
}
var file_proto_bridge_proto_depIdxs = []int32{
	0, // 0: bridge.transaction.type:type_name -> bridge.transaction_type
	1, // 1: bridge.transaction.status:type_name -> bridge.transaction_status
	2, // 2: bridge.complete_transfer_hash.action:type_name -> bridge.action_id
	4, // 3: bridge.submitted_signature.transaction:type_name -> bridge.transaction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_bridge_proto_init() }
func file_proto_bridge_proto_init() {
	if File_proto_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_proto_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_proto_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTransferHash); i {
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
		file_proto_bridge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmittedSignature); i {
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
		file_proto_bridge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensLockedEvent); i {
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
		file_proto_bridge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferCompletedEvent); i {
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
		file_proto_bridge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestNewSignaturesEvent); i {
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
			RawDescriptor: file_proto_bridge_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bridge_proto_goTypes,
		DependencyIndexes: file_proto_bridge_proto_depIdxs,
		EnumInfos:         file_proto_bridge_proto_enumTypes,
		MessageInfos:      file_proto_bridge_proto_msgTypes,
	}.Build()
	File_proto_bridge_proto = out.File
	file_proto_bridge_proto_rawDesc = nil
	file_proto_bridge_proto_goTypes = nil
	file_proto_bridge_proto_depIdxs = nil
}
