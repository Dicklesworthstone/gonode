// Copyright (c) 2021-2021 The Pastel Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: raptorq.proto

package raptorq

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

type UploadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadDataRequest) Reset() {
	*x = UploadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDataRequest) ProtoMessage() {}

func (x *UploadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDataRequest.ProtoReflect.Descriptor instead.
func (*UploadDataRequest) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{0}
}

func (x *UploadDataRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EncoderInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          []string           `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	EncoderParams *EncoderParameters `protobuf:"bytes,2,opt,name=encoder_params,json=encoderParams,proto3" json:"encoder_params,omitempty"`
}

func (x *EncoderInfoReply) Reset() {
	*x = EncoderInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncoderInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncoderInfoReply) ProtoMessage() {}

func (x *EncoderInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncoderInfoReply.ProtoReflect.Descriptor instead.
func (*EncoderInfoReply) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{1}
}

func (x *EncoderInfoReply) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *EncoderInfoReply) GetEncoderParams() *EncoderParameters {
	if x != nil {
		return x.EncoderParams
	}
	return nil
}

type EncoderParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransferLength  uint64 `protobuf:"varint,1,opt,name=transfer_length,json=transferLength,proto3" json:"transfer_length,omitempty"`      //u64
	SymbolSize      uint32 `protobuf:"varint,2,opt,name=symbol_size,json=symbolSize,proto3" json:"symbol_size,omitempty"`                  //u16
	NumSourceBlocks uint32 `protobuf:"varint,3,opt,name=num_source_blocks,json=numSourceBlocks,proto3" json:"num_source_blocks,omitempty"` //u8
	NumSubBlocks    uint32 `protobuf:"varint,4,opt,name=num_sub_blocks,json=numSubBlocks,proto3" json:"num_sub_blocks,omitempty"`          //u16
	SymbolAlignment uint32 `protobuf:"varint,5,opt,name=symbol_alignment,json=symbolAlignment,proto3" json:"symbol_alignment,omitempty"`   //u8
}

func (x *EncoderParameters) Reset() {
	*x = EncoderParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncoderParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncoderParameters) ProtoMessage() {}

func (x *EncoderParameters) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncoderParameters.ProtoReflect.Descriptor instead.
func (*EncoderParameters) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{2}
}

func (x *EncoderParameters) GetTransferLength() uint64 {
	if x != nil {
		return x.TransferLength
	}
	return 0
}

func (x *EncoderParameters) GetSymbolSize() uint32 {
	if x != nil {
		return x.SymbolSize
	}
	return 0
}

func (x *EncoderParameters) GetNumSourceBlocks() uint32 {
	if x != nil {
		return x.NumSourceBlocks
	}
	return 0
}

func (x *EncoderParameters) GetNumSubBlocks() uint32 {
	if x != nil {
		return x.NumSubBlocks
	}
	return 0
}

func (x *EncoderParameters) GetSymbolAlignment() uint32 {
	if x != nil {
		return x.SymbolAlignment
	}
	return 0
}

type SymbolReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol []byte `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *SymbolReply) Reset() {
	*x = SymbolReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolReply) ProtoMessage() {}

func (x *SymbolReply) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolReply.ProtoReflect.Descriptor instead.
func (*SymbolReply) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{3}
}

func (x *SymbolReply) GetSymbol() []byte {
	if x != nil {
		return x.Symbol
	}
	return nil
}

type UploadSymbolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ParamsOrSymbolsOneof:
	//	*UploadSymbolsRequest_EncoderParams
	//	*UploadSymbolsRequest_Symbol
	ParamsOrSymbolsOneof isUploadSymbolsRequest_ParamsOrSymbolsOneof `protobuf_oneof:"params_or_symbols_oneof"`
}

func (x *UploadSymbolsRequest) Reset() {
	*x = UploadSymbolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadSymbolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadSymbolsRequest) ProtoMessage() {}

func (x *UploadSymbolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadSymbolsRequest.ProtoReflect.Descriptor instead.
func (*UploadSymbolsRequest) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{4}
}

func (m *UploadSymbolsRequest) GetParamsOrSymbolsOneof() isUploadSymbolsRequest_ParamsOrSymbolsOneof {
	if m != nil {
		return m.ParamsOrSymbolsOneof
	}
	return nil
}

func (x *UploadSymbolsRequest) GetEncoderParams() *EncoderParameters {
	if x, ok := x.GetParamsOrSymbolsOneof().(*UploadSymbolsRequest_EncoderParams); ok {
		return x.EncoderParams
	}
	return nil
}

func (x *UploadSymbolsRequest) GetSymbol() []byte {
	if x, ok := x.GetParamsOrSymbolsOneof().(*UploadSymbolsRequest_Symbol); ok {
		return x.Symbol
	}
	return nil
}

type isUploadSymbolsRequest_ParamsOrSymbolsOneof interface {
	isUploadSymbolsRequest_ParamsOrSymbolsOneof()
}

type UploadSymbolsRequest_EncoderParams struct {
	EncoderParams *EncoderParameters `protobuf:"bytes,1,opt,name=encoder_params,json=encoderParams,proto3,oneof"`
}

type UploadSymbolsRequest_Symbol struct {
	Symbol []byte `protobuf:"bytes,2,opt,name=symbol,proto3,oneof"`
}

func (*UploadSymbolsRequest_EncoderParams) isUploadSymbolsRequest_ParamsOrSymbolsOneof() {}

func (*UploadSymbolsRequest_Symbol) isUploadSymbolsRequest_ParamsOrSymbolsOneof() {}

type DownloadDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadDataReply) Reset() {
	*x = DownloadDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raptorq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadDataReply) ProtoMessage() {}

func (x *DownloadDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_raptorq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadDataReply.ProtoReflect.Descriptor instead.
func (*DownloadDataReply) Descriptor() ([]byte, []int) {
	return file_raptorq_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadDataReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_raptorq_proto protoreflect.FileDescriptor

var file_raptorq_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x22, 0x27, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x69, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0xda, 0x01, 0x0a,
	0x11, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x5f,
	0x73, 0x75, 0x62, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x53, 0x75, 0x62, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x41, 0x6c, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x22, 0x90, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52,
	0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x42, 0x19, 0x0a, 0x17, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x5f, 0x6f, 0x72, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x5f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd4, 0x01, 0x0a,
	0x07, 0x52, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x51, 0x12, 0x44, 0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72,
	0x71, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x45, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3c,
	0x0a, 0x06, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f,
	0x72, 0x71, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x06,
	0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x28, 0x01, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x61, 0x73, 0x74, 0x65, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f,
	0x67, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x71, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raptorq_proto_rawDescOnce sync.Once
	file_raptorq_proto_rawDescData = file_raptorq_proto_rawDesc
)

func file_raptorq_proto_rawDescGZIP() []byte {
	file_raptorq_proto_rawDescOnce.Do(func() {
		file_raptorq_proto_rawDescData = protoimpl.X.CompressGZIP(file_raptorq_proto_rawDescData)
	})
	return file_raptorq_proto_rawDescData
}

var file_raptorq_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raptorq_proto_goTypes = []interface{}{
	(*UploadDataRequest)(nil),    // 0: raptorq.UploadDataRequest
	(*EncoderInfoReply)(nil),     // 1: raptorq.EncoderInfoReply
	(*EncoderParameters)(nil),    // 2: raptorq.EncoderParameters
	(*SymbolReply)(nil),          // 3: raptorq.SymbolReply
	(*UploadSymbolsRequest)(nil), // 4: raptorq.UploadSymbolsRequest
	(*DownloadDataReply)(nil),    // 5: raptorq.DownloadDataReply
}
var file_raptorq_proto_depIdxs = []int32{
	2, // 0: raptorq.EncoderInfoReply.encoder_params:type_name -> raptorq.EncoderParameters
	2, // 1: raptorq.UploadSymbolsRequest.encoder_params:type_name -> raptorq.EncoderParameters
	0, // 2: raptorq.RaptorQ.EncoderInfo:input_type -> raptorq.UploadDataRequest
	0, // 3: raptorq.RaptorQ.Encode:input_type -> raptorq.UploadDataRequest
	4, // 4: raptorq.RaptorQ.Decode:input_type -> raptorq.UploadSymbolsRequest
	1, // 5: raptorq.RaptorQ.EncoderInfo:output_type -> raptorq.EncoderInfoReply
	3, // 6: raptorq.RaptorQ.Encode:output_type -> raptorq.SymbolReply
	5, // 7: raptorq.RaptorQ.Decode:output_type -> raptorq.DownloadDataReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_raptorq_proto_init() }
func file_raptorq_proto_init() {
	if File_raptorq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_raptorq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDataRequest); i {
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
		file_raptorq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncoderInfoReply); i {
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
		file_raptorq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncoderParameters); i {
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
		file_raptorq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolReply); i {
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
		file_raptorq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadSymbolsRequest); i {
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
		file_raptorq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadDataReply); i {
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
	file_raptorq_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*UploadSymbolsRequest_EncoderParams)(nil),
		(*UploadSymbolsRequest_Symbol)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_raptorq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_raptorq_proto_goTypes,
		DependencyIndexes: file_raptorq_proto_depIdxs,
		MessageInfos:      file_raptorq_proto_msgTypes,
	}.Build()
	File_raptorq_proto = out.File
	file_raptorq_proto_rawDesc = nil
	file_raptorq_proto_goTypes = nil
	file_raptorq_proto_depIdxs = nil
}