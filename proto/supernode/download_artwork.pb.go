// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: download_artwork.proto

package supernode

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

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid      string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Ttxid     string `protobuf:"bytes,4,opt,name=ttxid,proto3" json:"ttxid,omitempty"`
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_artwork_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_download_artwork_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_download_artwork_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadRequest) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *DownloadRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *DownloadRequest) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *DownloadRequest) GetTtxid() string {
	if x != nil {
		return x.Ttxid
	}
	return ""
}

type DownloadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File []byte `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *DownloadReply) Reset() {
	*x = DownloadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_download_artwork_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadReply) ProtoMessage() {}

func (x *DownloadReply) ProtoReflect() protoreflect.Message {
	mi := &file_download_artwork_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadReply.ProtoReflect.Descriptor instead.
func (*DownloadReply) Descriptor() ([]byte, []int) {
	return file_download_artwork_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadReply) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_download_artwork_proto protoreflect.FileDescriptor

var file_download_artwork_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x61, 0x72, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e,
	0x6f, 0x64, 0x65, 0x22, 0x77, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x74, 0x78, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x74, 0x78, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x0d,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x32, 0x55, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x12, 0x42, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x1a, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x28, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x73, 0x74, 0x65, 0x6c, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x6f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_download_artwork_proto_rawDescOnce sync.Once
	file_download_artwork_proto_rawDescData = file_download_artwork_proto_rawDesc
)

func file_download_artwork_proto_rawDescGZIP() []byte {
	file_download_artwork_proto_rawDescOnce.Do(func() {
		file_download_artwork_proto_rawDescData = protoimpl.X.CompressGZIP(file_download_artwork_proto_rawDescData)
	})
	return file_download_artwork_proto_rawDescData
}

var file_download_artwork_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_download_artwork_proto_goTypes = []interface{}{
	(*DownloadRequest)(nil), // 0: supernode.DownloadRequest
	(*DownloadReply)(nil),   // 1: supernode.DownloadReply
}
var file_download_artwork_proto_depIdxs = []int32{
	0, // 0: supernode.DownloadArtwork.Download:input_type -> supernode.DownloadRequest
	1, // 1: supernode.DownloadArtwork.Download:output_type -> supernode.DownloadReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_download_artwork_proto_init() }
func file_download_artwork_proto_init() {
	if File_download_artwork_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_download_artwork_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_download_artwork_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadReply); i {
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
			RawDescriptor: file_download_artwork_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_download_artwork_proto_goTypes,
		DependencyIndexes: file_download_artwork_proto_depIdxs,
		MessageInfos:      file_download_artwork_proto_msgTypes,
	}.Build()
	File_download_artwork_proto = out.File
	file_download_artwork_proto_rawDesc = nil
	file_download_artwork_proto_goTypes = nil
	file_download_artwork_proto_depIdxs = nil
}
