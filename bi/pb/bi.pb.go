// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: bi.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_bi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_bi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_bi_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Res           string                 `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_bi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_bi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_bi_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

var File_bi_proto protoreflect.FileDescriptor

const file_bi_proto_rawDesc = "" +
	"\n" +
	"\bbi.proto\x12\vbidirection\"\x1a\n" +
	"\x04Data\x12\x12\n" +
	"\x04data\x18\x01 \x01(\tR\x04data\"\x1c\n" +
	"\bResponse\x12\x10\n" +
	"\x03res\x18\x01 \x01(\tR\x03res2L\n" +
	"\vBidirection\x12=\n" +
	"\vBidirection\x12\x11.bidirection.Data\x1a\x15.bidirection.Response\"\x00(\x010\x01B\x06Z\x04.;pbb\x06proto3"

var (
	file_bi_proto_rawDescOnce sync.Once
	file_bi_proto_rawDescData []byte
)

func file_bi_proto_rawDescGZIP() []byte {
	file_bi_proto_rawDescOnce.Do(func() {
		file_bi_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bi_proto_rawDesc), len(file_bi_proto_rawDesc)))
	})
	return file_bi_proto_rawDescData
}

var file_bi_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bi_proto_goTypes = []any{
	(*Data)(nil),     // 0: bidirection.Data
	(*Response)(nil), // 1: bidirection.Response
}
var file_bi_proto_depIdxs = []int32{
	0, // 0: bidirection.Bidirection.Bidirection:input_type -> bidirection.Data
	1, // 1: bidirection.Bidirection.Bidirection:output_type -> bidirection.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bi_proto_init() }
func file_bi_proto_init() {
	if File_bi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bi_proto_rawDesc), len(file_bi_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bi_proto_goTypes,
		DependencyIndexes: file_bi_proto_depIdxs,
		MessageInfos:      file_bi_proto_msgTypes,
	}.Build()
	File_bi_proto = out.File
	file_bi_proto_goTypes = nil
	file_bi_proto_depIdxs = nil
}
