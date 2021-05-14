// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: walletnode/register_artowrk.proto

package proto

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

type HandshakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPrimary bool `protobuf:"varint,1,opt,name=is_primary,json=isPrimary,proto3" json:"is_primary,omitempty"`
}

func (x *HandshakeRequest) Reset() {
	*x = HandshakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeRequest) ProtoMessage() {}

func (x *HandshakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeRequest.ProtoReflect.Descriptor instead.
func (*HandshakeRequest) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeRequest) GetIsPrimary() bool {
	if x != nil {
		return x.IsPrimary
	}
	return false
}

type HandshakeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
}

func (x *HandshakeReply) Reset() {
	*x = HandshakeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeReply) ProtoMessage() {}

func (x *HandshakeReply) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeReply.ProtoReflect.Descriptor instead.
func (*HandshakeReply) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{1}
}

func (x *HandshakeReply) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

type AcceptedNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AcceptedNodesRequest) Reset() {
	*x = AcceptedNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptedNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptedNodesRequest) ProtoMessage() {}

func (x *AcceptedNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptedNodesRequest.ProtoReflect.Descriptor instead.
func (*AcceptedNodesRequest) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{2}
}

type AcceptedNodesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*AcceptedNodesReply_Peer `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *AcceptedNodesReply) Reset() {
	*x = AcceptedNodesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptedNodesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptedNodesReply) ProtoMessage() {}

func (x *AcceptedNodesReply) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptedNodesReply.ProtoReflect.Descriptor instead.
func (*AcceptedNodesReply) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{3}
}

func (x *AcceptedNodesReply) GetPeers() []*AcceptedNodesReply_Peer {
	if x != nil {
		return x.Peers
	}
	return nil
}

type ConnectToRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskID  string `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	NodeKey string `protobuf:"bytes,2,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
}

func (x *ConnectToRequest) Reset() {
	*x = ConnectToRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectToRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectToRequest) ProtoMessage() {}

func (x *ConnectToRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectToRequest.ProtoReflect.Descriptor instead.
func (*ConnectToRequest) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectToRequest) GetTaskID() string {
	if x != nil {
		return x.TaskID
	}
	return ""
}

func (x *ConnectToRequest) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

type ConnectToReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectToReply) Reset() {
	*x = ConnectToReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectToReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectToReply) ProtoMessage() {}

func (x *ConnectToReply) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectToReply.ProtoReflect.Descriptor instead.
func (*ConnectToReply) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{5}
}

type SendImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SendImageRequest) Reset() {
	*x = SendImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendImageRequest) ProtoMessage() {}

func (x *SendImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendImageRequest.ProtoReflect.Descriptor instead.
func (*SendImageRequest) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{6}
}

func (x *SendImageRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type SendImageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fgpt string `protobuf:"bytes,1,opt,name=fgpt,proto3" json:"fgpt,omitempty"`
}

func (x *SendImageReply) Reset() {
	*x = SendImageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendImageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendImageReply) ProtoMessage() {}

func (x *SendImageReply) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendImageReply.ProtoReflect.Descriptor instead.
func (*SendImageReply) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{7}
}

func (x *SendImageReply) GetFgpt() string {
	if x != nil {
		return x.Fgpt
	}
	return ""
}

type SendTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket          []byte `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
	TicketSignature string `protobuf:"bytes,2,opt,name=ticket_signature,json=ticketSignature,proto3" json:"ticket_signature,omitempty"`
	Fgpt            string `protobuf:"bytes,3,opt,name=fgpt,proto3" json:"fgpt,omitempty"`
	FgptSignature   string `protobuf:"bytes,4,opt,name=fgpt_signature,json=fgptSignature,proto3" json:"fgpt_signature,omitempty"`
	FeeTxid         string `protobuf:"bytes,5,opt,name=fee_txid,json=feeTxid,proto3" json:"fee_txid,omitempty"`
	Thumbnail       []byte `protobuf:"bytes,6,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *SendTicketRequest) Reset() {
	*x = SendTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTicketRequest) ProtoMessage() {}

func (x *SendTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTicketRequest.ProtoReflect.Descriptor instead.
func (*SendTicketRequest) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{8}
}

func (x *SendTicketRequest) GetTicket() []byte {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *SendTicketRequest) GetTicketSignature() string {
	if x != nil {
		return x.TicketSignature
	}
	return ""
}

func (x *SendTicketRequest) GetFgpt() string {
	if x != nil {
		return x.Fgpt
	}
	return ""
}

func (x *SendTicketRequest) GetFgptSignature() string {
	if x != nil {
		return x.FgptSignature
	}
	return ""
}

func (x *SendTicketRequest) GetFeeTxid() string {
	if x != nil {
		return x.FeeTxid
	}
	return ""
}

func (x *SendTicketRequest) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

type SendTicketReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketTxid string `protobuf:"bytes,1,opt,name=ticket_txid,json=ticketTxid,proto3" json:"ticket_txid,omitempty"`
}

func (x *SendTicketReply) Reset() {
	*x = SendTicketReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTicketReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTicketReply) ProtoMessage() {}

func (x *SendTicketReply) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTicketReply.ProtoReflect.Descriptor instead.
func (*SendTicketReply) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{9}
}

func (x *SendTicketReply) GetTicketTxid() string {
	if x != nil {
		return x.TicketTxid
	}
	return ""
}

type AcceptedNodesReply_Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeKey string `protobuf:"bytes,1,opt,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
}

func (x *AcceptedNodesReply_Peer) Reset() {
	*x = AcceptedNodesReply_Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_walletnode_register_artowrk_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptedNodesReply_Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptedNodesReply_Peer) ProtoMessage() {}

func (x *AcceptedNodesReply_Peer) ProtoReflect() protoreflect.Message {
	mi := &file_walletnode_register_artowrk_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptedNodesReply_Peer.ProtoReflect.Descriptor instead.
func (*AcceptedNodesReply_Peer) Descriptor() ([]byte, []int) {
	return file_walletnode_register_artowrk_proto_rawDescGZIP(), []int{3, 0}
}

func (x *AcceptedNodesReply_Peer) GetNodeKey() string {
	if x != nil {
		return x.NodeKey
	}
	return ""
}

var File_walletnode_register_artowrk_proto protoreflect.FileDescriptor

var file_walletnode_register_artowrk_proto_rawDesc = []byte{
	0x0a, 0x21, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x74, 0x6f, 0x77, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x22,
	0x31, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x22, 0x28, 0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x22, 0x16, 0x0a, 0x14,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x70, 0x65,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x1a, 0x21, 0x0a, 0x04, 0x50, 0x65, 0x65, 0x72, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x45, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x22,
	0x10, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x24, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x67, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x67, 0x70, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x67, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x67,
	0x70, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x67, 0x70, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x67, 0x70, 0x74,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65,
	0x5f, 0x74, 0x78, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x65,
	0x54, 0x78, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x22, 0x32, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x54, 0x78, 0x69, 0x64, 0x32, 0x89, 0x03, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x72, 0x74, 0x6f, 0x77, 0x72, 0x6b, 0x12, 0x49, 0x0a, 0x09, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x28, 0x01, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x54, 0x6f, 0x12, 0x1c, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x47, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x73, 0x74, 0x65, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67,
	0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_walletnode_register_artowrk_proto_rawDescOnce sync.Once
	file_walletnode_register_artowrk_proto_rawDescData = file_walletnode_register_artowrk_proto_rawDesc
)

func file_walletnode_register_artowrk_proto_rawDescGZIP() []byte {
	file_walletnode_register_artowrk_proto_rawDescOnce.Do(func() {
		file_walletnode_register_artowrk_proto_rawDescData = protoimpl.X.CompressGZIP(file_walletnode_register_artowrk_proto_rawDescData)
	})
	return file_walletnode_register_artowrk_proto_rawDescData
}

var file_walletnode_register_artowrk_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_walletnode_register_artowrk_proto_goTypes = []interface{}{
	(*HandshakeRequest)(nil),        // 0: walletnode.HandshakeRequest
	(*HandshakeReply)(nil),          // 1: walletnode.HandshakeReply
	(*AcceptedNodesRequest)(nil),    // 2: walletnode.AcceptedNodesRequest
	(*AcceptedNodesReply)(nil),      // 3: walletnode.AcceptedNodesReply
	(*ConnectToRequest)(nil),        // 4: walletnode.ConnectToRequest
	(*ConnectToReply)(nil),          // 5: walletnode.ConnectToReply
	(*SendImageRequest)(nil),        // 6: walletnode.SendImageRequest
	(*SendImageReply)(nil),          // 7: walletnode.SendImageReply
	(*SendTicketRequest)(nil),       // 8: walletnode.SendTicketRequest
	(*SendTicketReply)(nil),         // 9: walletnode.SendTicketReply
	(*AcceptedNodesReply_Peer)(nil), // 10: walletnode.AcceptedNodesReply.Peer
}
var file_walletnode_register_artowrk_proto_depIdxs = []int32{
	10, // 0: walletnode.AcceptedNodesReply.peers:type_name -> walletnode.AcceptedNodesReply.Peer
	0,  // 1: walletnode.RegisterArtowrk.Handshake:input_type -> walletnode.HandshakeRequest
	2,  // 2: walletnode.RegisterArtowrk.AcceptedNodes:input_type -> walletnode.AcceptedNodesRequest
	4,  // 3: walletnode.RegisterArtowrk.ConnectTo:input_type -> walletnode.ConnectToRequest
	6,  // 4: walletnode.RegisterArtowrk.SendImage:input_type -> walletnode.SendImageRequest
	8,  // 5: walletnode.RegisterArtowrk.SendTicket:input_type -> walletnode.SendTicketRequest
	1,  // 6: walletnode.RegisterArtowrk.Handshake:output_type -> walletnode.HandshakeReply
	3,  // 7: walletnode.RegisterArtowrk.AcceptedNodes:output_type -> walletnode.AcceptedNodesReply
	5,  // 8: walletnode.RegisterArtowrk.ConnectTo:output_type -> walletnode.ConnectToReply
	7,  // 9: walletnode.RegisterArtowrk.SendImage:output_type -> walletnode.SendImageReply
	9,  // 10: walletnode.RegisterArtowrk.SendTicket:output_type -> walletnode.SendTicketReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_walletnode_register_artowrk_proto_init() }
func file_walletnode_register_artowrk_proto_init() {
	if File_walletnode_register_artowrk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_walletnode_register_artowrk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeRequest); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeReply); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptedNodesRequest); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptedNodesReply); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectToRequest); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectToReply); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendImageRequest); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendImageReply); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTicketRequest); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTicketReply); i {
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
		file_walletnode_register_artowrk_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptedNodesReply_Peer); i {
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
			RawDescriptor: file_walletnode_register_artowrk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_walletnode_register_artowrk_proto_goTypes,
		DependencyIndexes: file_walletnode_register_artowrk_proto_depIdxs,
		MessageInfos:      file_walletnode_register_artowrk_proto_msgTypes,
	}.Build()
	File_walletnode_register_artowrk_proto = out.File
	file_walletnode_register_artowrk_proto_rawDesc = nil
	file_walletnode_register_artowrk_proto_goTypes = nil
	file_walletnode_register_artowrk_proto_depIdxs = nil
}
