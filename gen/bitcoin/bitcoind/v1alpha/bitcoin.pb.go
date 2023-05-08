// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: bitcoin/bitcoind/v1alpha/bitcoin.proto

package bitcoindv1alpha

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

type GetBlockchainInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlockchainInfoRequest) Reset() {
	*x = GetBlockchainInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoRequest) ProtoMessage() {}

func (x *GetBlockchainInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockchainInfoRequest.ProtoReflect.Descriptor instead.
func (*GetBlockchainInfoRequest) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{0}
}

type GetBlockchainInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestBlockHash        string `protobuf:"bytes,1,opt,name=best_block_hash,json=bestBlockHash,proto3" json:"best_block_hash,omitempty"`
	Chain                string `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	ChainWork            string `protobuf:"bytes,3,opt,name=chain_work,json=chainWork,proto3" json:"chain_work,omitempty"`
	InitialBlockDownload bool   `protobuf:"varint,4,opt,name=initial_block_download,json=initialBlockDownload,proto3" json:"initial_block_download,omitempty"`
}

func (x *GetBlockchainInfoResponse) Reset() {
	*x = GetBlockchainInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoResponse) ProtoMessage() {}

func (x *GetBlockchainInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockchainInfoResponse.ProtoReflect.Descriptor instead.
func (*GetBlockchainInfoResponse) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockchainInfoResponse) GetBestBlockHash() string {
	if x != nil {
		return x.BestBlockHash
	}
	return ""
}

func (x *GetBlockchainInfoResponse) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *GetBlockchainInfoResponse) GetChainWork() string {
	if x != nil {
		return x.ChainWork
	}
	return ""
}

func (x *GetBlockchainInfoResponse) GetInitialBlockDownload() bool {
	if x != nil {
		return x.InitialBlockDownload
	}
	return false
}

type GetNewAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label       string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	AddressType string `protobuf:"bytes,2,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
}

func (x *GetNewAddressRequest) Reset() {
	*x = GetNewAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewAddressRequest) ProtoMessage() {}

func (x *GetNewAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewAddressRequest.ProtoReflect.Descriptor instead.
func (*GetNewAddressRequest) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{2}
}

func (x *GetNewAddressRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *GetNewAddressRequest) GetAddressType() string {
	if x != nil {
		return x.AddressType
	}
	return ""
}

type GetNewAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetNewAddressResponse) Reset() {
	*x = GetNewAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNewAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNewAddressResponse) ProtoMessage() {}

func (x *GetNewAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNewAddressResponse.ProtoReflect.Descriptor instead.
func (*GetNewAddressResponse) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{3}
}

func (x *GetNewAddressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetWalletInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWalletInfoRequest) Reset() {
	*x = GetWalletInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletInfoRequest) ProtoMessage() {}

func (x *GetWalletInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletInfoRequest.ProtoReflect.Descriptor instead.
func (*GetWalletInfoRequest) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{4}
}

type GetWalletInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletName            string  `protobuf:"bytes,1,opt,name=wallet_name,json=walletName,proto3" json:"wallet_name,omitempty"`
	WalletVersion         int64   `protobuf:"varint,2,opt,name=wallet_version,json=walletVersion,proto3" json:"wallet_version,omitempty"`
	Format                string  `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	TxCount               int64   `protobuf:"varint,7,opt,name=tx_count,json=txCount,proto3" json:"tx_count,omitempty"`
	KeyPoolSize           int64   `protobuf:"varint,8,opt,name=key_pool_size,json=keyPoolSize,proto3" json:"key_pool_size,omitempty"`
	KeyPoolSizeHdInternal int64   `protobuf:"varint,9,opt,name=key_pool_size_hd_internal,json=keyPoolSizeHdInternal,proto3" json:"key_pool_size_hd_internal,omitempty"`
	PayTxFee              float64 `protobuf:"fixed64,10,opt,name=pay_tx_fee,json=payTxFee,proto3" json:"pay_tx_fee,omitempty"`
	PrivateKeysEnabled    bool    `protobuf:"varint,11,opt,name=private_keys_enabled,json=privateKeysEnabled,proto3" json:"private_keys_enabled,omitempty"`
	AvoidReuse            bool    `protobuf:"varint,12,opt,name=avoid_reuse,json=avoidReuse,proto3" json:"avoid_reuse,omitempty"`
	// Not set if no scan is in progress.
	Scanning       *WalletScan `protobuf:"bytes,13,opt,name=scanning,proto3" json:"scanning,omitempty"`
	Descriptors    bool        `protobuf:"varint,14,opt,name=descriptors,proto3" json:"descriptors,omitempty"`
	ExternalSigner bool        `protobuf:"varint,15,opt,name=external_signer,json=externalSigner,proto3" json:"external_signer,omitempty"`
}

func (x *GetWalletInfoResponse) Reset() {
	*x = GetWalletInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletInfoResponse) ProtoMessage() {}

func (x *GetWalletInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletInfoResponse.ProtoReflect.Descriptor instead.
func (*GetWalletInfoResponse) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{5}
}

func (x *GetWalletInfoResponse) GetWalletName() string {
	if x != nil {
		return x.WalletName
	}
	return ""
}

func (x *GetWalletInfoResponse) GetWalletVersion() int64 {
	if x != nil {
		return x.WalletVersion
	}
	return 0
}

func (x *GetWalletInfoResponse) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *GetWalletInfoResponse) GetTxCount() int64 {
	if x != nil {
		return x.TxCount
	}
	return 0
}

func (x *GetWalletInfoResponse) GetKeyPoolSize() int64 {
	if x != nil {
		return x.KeyPoolSize
	}
	return 0
}

func (x *GetWalletInfoResponse) GetKeyPoolSizeHdInternal() int64 {
	if x != nil {
		return x.KeyPoolSizeHdInternal
	}
	return 0
}

func (x *GetWalletInfoResponse) GetPayTxFee() float64 {
	if x != nil {
		return x.PayTxFee
	}
	return 0
}

func (x *GetWalletInfoResponse) GetPrivateKeysEnabled() bool {
	if x != nil {
		return x.PrivateKeysEnabled
	}
	return false
}

func (x *GetWalletInfoResponse) GetAvoidReuse() bool {
	if x != nil {
		return x.AvoidReuse
	}
	return false
}

func (x *GetWalletInfoResponse) GetScanning() *WalletScan {
	if x != nil {
		return x.Scanning
	}
	return nil
}

func (x *GetWalletInfoResponse) GetDescriptors() bool {
	if x != nil {
		return x.Descriptors
	}
	return false
}

func (x *GetWalletInfoResponse) GetExternalSigner() bool {
	if x != nil {
		return x.ExternalSigner
	}
	return false
}

type WalletScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration int64   `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"`
	Progress float64 `protobuf:"fixed64,2,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *WalletScan) Reset() {
	*x = WalletScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletScan) ProtoMessage() {}

func (x *WalletScan) ProtoReflect() protoreflect.Message {
	mi := &file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletScan.ProtoReflect.Descriptor instead.
func (*WalletScan) Descriptor() ([]byte, []int) {
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP(), []int{6}
}

func (x *WalletScan) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *WalletScan) GetProgress() float64 {
	if x != nil {
		return x.Progress
	}
	return 0
}

var File_bitcoin_bitcoind_v1alpha_bitcoin_proto protoreflect.FileDescriptor

var file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xae,
	0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x62, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x4f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x31, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xee, 0x03, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x22, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x50, 0x6f, 0x6f, 0x6c,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x38, 0x0a, 0x19, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x6f, 0x6f, 0x6c,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x68, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6b, 0x65, 0x79, 0x50, 0x6f, 0x6f, 0x6c,
	0x53, 0x69, 0x7a, 0x65, 0x48, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x0a, 0x70, 0x61, 0x79, 0x5f, 0x74, 0x78, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x79, 0x54, 0x78, 0x46, 0x65, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x75, 0x73, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x52, 0x65, 0x75, 0x73, 0x65, 0x12,
	0x40, 0x0a, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63,
	0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0a,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x32, 0xf2, 0x02, 0x0a, 0x0e, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xf7, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x62, 0x74, 0x63, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xa2, 0x02, 0x03, 0x42, 0x42, 0x58, 0xaa, 0x02, 0x18, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xca, 0x02, 0x18, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x24,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64,
	0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x3a, 0x3a,
	0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescOnce sync.Once
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescData = file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDesc
)

func file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescGZIP() []byte {
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescOnce.Do(func() {
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescData = protoimpl.X.CompressGZIP(file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescData)
	})
	return file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDescData
}

var file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bitcoin_bitcoind_v1alpha_bitcoin_proto_goTypes = []interface{}{
	(*GetBlockchainInfoRequest)(nil),  // 0: bitcoin.bitcoind.v1alpha.GetBlockchainInfoRequest
	(*GetBlockchainInfoResponse)(nil), // 1: bitcoin.bitcoind.v1alpha.GetBlockchainInfoResponse
	(*GetNewAddressRequest)(nil),      // 2: bitcoin.bitcoind.v1alpha.GetNewAddressRequest
	(*GetNewAddressResponse)(nil),     // 3: bitcoin.bitcoind.v1alpha.GetNewAddressResponse
	(*GetWalletInfoRequest)(nil),      // 4: bitcoin.bitcoind.v1alpha.GetWalletInfoRequest
	(*GetWalletInfoResponse)(nil),     // 5: bitcoin.bitcoind.v1alpha.GetWalletInfoResponse
	(*WalletScan)(nil),                // 6: bitcoin.bitcoind.v1alpha.WalletScan
}
var file_bitcoin_bitcoind_v1alpha_bitcoin_proto_depIdxs = []int32{
	6, // 0: bitcoin.bitcoind.v1alpha.GetWalletInfoResponse.scanning:type_name -> bitcoin.bitcoind.v1alpha.WalletScan
	0, // 1: bitcoin.bitcoind.v1alpha.BitcoinService.GetBlockchainInfo:input_type -> bitcoin.bitcoind.v1alpha.GetBlockchainInfoRequest
	2, // 2: bitcoin.bitcoind.v1alpha.BitcoinService.GetNewAddress:input_type -> bitcoin.bitcoind.v1alpha.GetNewAddressRequest
	4, // 3: bitcoin.bitcoind.v1alpha.BitcoinService.GetWalletInfo:input_type -> bitcoin.bitcoind.v1alpha.GetWalletInfoRequest
	1, // 4: bitcoin.bitcoind.v1alpha.BitcoinService.GetBlockchainInfo:output_type -> bitcoin.bitcoind.v1alpha.GetBlockchainInfoResponse
	3, // 5: bitcoin.bitcoind.v1alpha.BitcoinService.GetNewAddress:output_type -> bitcoin.bitcoind.v1alpha.GetNewAddressResponse
	5, // 6: bitcoin.bitcoind.v1alpha.BitcoinService.GetWalletInfo:output_type -> bitcoin.bitcoind.v1alpha.GetWalletInfoResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bitcoin_bitcoind_v1alpha_bitcoin_proto_init() }
func file_bitcoin_bitcoind_v1alpha_bitcoin_proto_init() {
	if File_bitcoin_bitcoind_v1alpha_bitcoin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockchainInfoRequest); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockchainInfoResponse); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewAddressRequest); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNewAddressResponse); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletInfoRequest); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletInfoResponse); i {
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
		file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletScan); i {
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
			RawDescriptor: file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bitcoin_bitcoind_v1alpha_bitcoin_proto_goTypes,
		DependencyIndexes: file_bitcoin_bitcoind_v1alpha_bitcoin_proto_depIdxs,
		MessageInfos:      file_bitcoin_bitcoind_v1alpha_bitcoin_proto_msgTypes,
	}.Build()
	File_bitcoin_bitcoind_v1alpha_bitcoin_proto = out.File
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_rawDesc = nil
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_goTypes = nil
	file_bitcoin_bitcoind_v1alpha_bitcoin_proto_depIdxs = nil
}
