// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: echo/echo.proto

package echo

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_echo_echo_proto protoreflect.FileDescriptor

var file_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x1f, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x20, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x25, 0x0a, 0x04, 0x45, 0x63,
	0x68, 0x6f, 0x12, 0x1d, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x09, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3b, 0x65, 0x63, 0x68, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_echo_proto_rawDescOnce sync.Once
	file_echo_echo_proto_rawDescData = file_echo_echo_proto_rawDesc
)

func file_echo_echo_proto_rawDescGZIP() []byte {
	file_echo_echo_proto_rawDescOnce.Do(func() {
		file_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_echo_proto_rawDescData)
	})
	return file_echo_echo_proto_rawDescData
}

var file_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_echo_proto_goTypes = []any{
	(*Req)(nil),  // 0: echo.Req
	(*Resp)(nil), // 1: echo.Resp
}
var file_echo_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.Echo:input_type -> echo.Req
	1, // 1: echo.Echo.Echo:output_type -> echo.Resp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_echo_proto_init() }
func file_echo_echo_proto_init() {
	if File_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_echo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Req); i {
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
		file_echo_echo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Resp); i {
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
			RawDescriptor: file_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_echo_proto_goTypes,
		DependencyIndexes: file_echo_echo_proto_depIdxs,
		MessageInfos:      file_echo_echo_proto_msgTypes,
	}.Build()
	File_echo_echo_proto = out.File
	file_echo_echo_proto_rawDesc = nil
	file_echo_echo_proto_goTypes = nil
	file_echo_echo_proto_depIdxs = nil
}
