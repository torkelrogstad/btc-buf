// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: barebitcoin/bitcoind/v22/bitcoin.proto

package bitcoindv22

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
		mi := &file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoRequest) ProtoMessage() {}

func (x *GetBlockchainInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[0]
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
	return file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescGZIP(), []int{0}
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
		mi := &file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockchainInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockchainInfoResponse) ProtoMessage() {}

func (x *GetBlockchainInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[1]
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
	return file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescGZIP(), []int{1}
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

var File_barebitcoin_bitcoind_v22_bitcoin_proto protoreflect.FileDescriptor

var file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDesc = []byte{
	0x0a, 0x26, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x32, 0x32, 0x2f, 0x62, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x32, 0x32, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
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
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x32,
	0x8e, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x32, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x61,
	0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69,
	0x6e, 0x64, 0x2e, 0x76, 0x32, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xf3, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69, 0x74,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x32,
	0x32, 0x42, 0x0c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61,
	0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x62, 0x74, 0x63, 0x2d, 0x62, 0x75,
	0x66, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x32, 0x32, 0x3b, 0x62, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x64, 0x76, 0x32, 0x32, 0xa2, 0x02, 0x03, 0x42, 0x42, 0x58, 0xaa, 0x02, 0x18, 0x42,
	0x61, 0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x74, 0x63, 0x6f,
	0x69, 0x6e, 0x64, 0x2e, 0x56, 0x32, 0x32, 0xca, 0x02, 0x18, 0x42, 0x61, 0x72, 0x65, 0x62, 0x69,
	0x74, 0x63, 0x6f, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x5c, 0x56,
	0x32, 0x32, 0xe2, 0x02, 0x24, 0x42, 0x61, 0x72, 0x65, 0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x5c, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x64, 0x5c, 0x56, 0x32, 0x32, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x42, 0x61, 0x72, 0x65,
	0x62, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x3a, 0x3a, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e,
	0x64, 0x3a, 0x3a, 0x56, 0x32, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescOnce sync.Once
	file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescData = file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDesc
)

func file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescGZIP() []byte {
	file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescOnce.Do(func() {
		file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescData = protoimpl.X.CompressGZIP(file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescData)
	})
	return file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDescData
}

var file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_barebitcoin_bitcoind_v22_bitcoin_proto_goTypes = []interface{}{
	(*GetBlockchainInfoRequest)(nil),  // 0: barebitcoin.bitcoind.v22.GetBlockchainInfoRequest
	(*GetBlockchainInfoResponse)(nil), // 1: barebitcoin.bitcoind.v22.GetBlockchainInfoResponse
}
var file_barebitcoin_bitcoind_v22_bitcoin_proto_depIdxs = []int32{
	0, // 0: barebitcoin.bitcoind.v22.BitcoinService.GetBlockchainInfo:input_type -> barebitcoin.bitcoind.v22.GetBlockchainInfoRequest
	1, // 1: barebitcoin.bitcoind.v22.BitcoinService.GetBlockchainInfo:output_type -> barebitcoin.bitcoind.v22.GetBlockchainInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_barebitcoin_bitcoind_v22_bitcoin_proto_init() }
func file_barebitcoin_bitcoind_v22_bitcoin_proto_init() {
	if File_barebitcoin_bitcoind_v22_bitcoin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_barebitcoin_bitcoind_v22_bitcoin_proto_goTypes,
		DependencyIndexes: file_barebitcoin_bitcoind_v22_bitcoin_proto_depIdxs,
		MessageInfos:      file_barebitcoin_bitcoind_v22_bitcoin_proto_msgTypes,
	}.Build()
	File_barebitcoin_bitcoind_v22_bitcoin_proto = out.File
	file_barebitcoin_bitcoind_v22_bitcoin_proto_rawDesc = nil
	file_barebitcoin_bitcoind_v22_bitcoin_proto_goTypes = nil
	file_barebitcoin_bitcoind_v22_bitcoin_proto_depIdxs = nil
}
