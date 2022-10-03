// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: comms/comms.proto

package comms

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

type CounterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CounterReq) Reset() {
	*x = CounterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_comms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterReq) ProtoMessage() {}

func (x *CounterReq) ProtoReflect() protoreflect.Message {
	mi := &file_comms_comms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterReq.ProtoReflect.Descriptor instead.
func (*CounterReq) Descriptor() ([]byte, []int) {
	return file_comms_comms_proto_rawDescGZIP(), []int{0}
}

type CounterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter int32 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *CounterResp) Reset() {
	*x = CounterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_comms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterResp) ProtoMessage() {}

func (x *CounterResp) ProtoReflect() protoreflect.Message {
	mi := &file_comms_comms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterResp.ProtoReflect.Descriptor instead.
func (*CounterResp) Descriptor() ([]byte, []int) {
	return file_comms_comms_proto_rawDescGZIP(), []int{1}
}

func (x *CounterResp) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_comms_comms_proto protoreflect.FileDescriptor

var file_comms_comms_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x27, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x32, 0x3d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x73,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x61, 0x6c, 0x72, 0x69, 0x31, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67, 0x72, 0x70, 0x73, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_comms_comms_proto_rawDescOnce sync.Once
	file_comms_comms_proto_rawDescData = file_comms_comms_proto_rawDesc
)

func file_comms_comms_proto_rawDescGZIP() []byte {
	file_comms_comms_proto_rawDescOnce.Do(func() {
		file_comms_comms_proto_rawDescData = protoimpl.X.CompressGZIP(file_comms_comms_proto_rawDescData)
	})
	return file_comms_comms_proto_rawDescData
}

var file_comms_comms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comms_comms_proto_goTypes = []interface{}{
	(*CounterReq)(nil),  // 0: comms.CounterReq
	(*CounterResp)(nil), // 1: comms.CounterResp
}
var file_comms_comms_proto_depIdxs = []int32{
	0, // 0: comms.Counter.Count:input_type -> comms.CounterReq
	1, // 1: comms.Counter.Count:output_type -> comms.CounterResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comms_comms_proto_init() }
func file_comms_comms_proto_init() {
	if File_comms_comms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comms_comms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterReq); i {
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
		file_comms_comms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterResp); i {
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
			RawDescriptor: file_comms_comms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comms_comms_proto_goTypes,
		DependencyIndexes: file_comms_comms_proto_depIdxs,
		MessageInfos:      file_comms_comms_proto_msgTypes,
	}.Build()
	File_comms_comms_proto = out.File
	file_comms_comms_proto_rawDesc = nil
	file_comms_comms_proto_goTypes = nil
	file_comms_comms_proto_depIdxs = nil
}