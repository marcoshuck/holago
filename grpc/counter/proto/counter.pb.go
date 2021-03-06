// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: proto/counter.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{0}
}

type OperationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OperationInput) Reset() {
	*x = OperationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationInput) ProtoMessage() {}

func (x *OperationInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationInput.ProtoReflect.Descriptor instead.
func (*OperationInput) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{1}
}

func (x *OperationInput) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type OperationOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Result int32 `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OperationOutput) Reset() {
	*x = OperationOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationOutput) ProtoMessage() {}

func (x *OperationOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationOutput.ProtoReflect.Descriptor instead.
func (*OperationOutput) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{2}
}

func (x *OperationOutput) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *OperationOutput) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ResultOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultOutput) Reset() {
	*x = ResultOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_counter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultOutput) ProtoMessage() {}

func (x *ResultOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_counter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultOutput.ProtoReflect.Descriptor instead.
func (*ResultOutput) Descriptor() ([]byte, []int) {
	return file_proto_counter_proto_rawDescGZIP(), []int{3}
}

func (x *ResultOutput) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_counter_proto protoreflect.FileDescriptor

var file_proto_counter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x3f, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xbe, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x12, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_counter_proto_rawDescOnce sync.Once
	file_proto_counter_proto_rawDescData = file_proto_counter_proto_rawDesc
)

func file_proto_counter_proto_rawDescGZIP() []byte {
	file_proto_counter_proto_rawDescOnce.Do(func() {
		file_proto_counter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_counter_proto_rawDescData)
	})
	return file_proto_counter_proto_rawDescData
}

var file_proto_counter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_counter_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: counter.Empty
	(*OperationInput)(nil),  // 1: counter.OperationInput
	(*OperationOutput)(nil), // 2: counter.OperationOutput
	(*ResultOutput)(nil),    // 3: counter.ResultOutput
}
var file_proto_counter_proto_depIdxs = []int32{
	1, // 0: counter.Counter.Increase:input_type -> counter.OperationInput
	1, // 1: counter.Counter.Decrease:input_type -> counter.OperationInput
	0, // 2: counter.Counter.Result:input_type -> counter.Empty
	2, // 3: counter.Counter.Increase:output_type -> counter.OperationOutput
	2, // 4: counter.Counter.Decrease:output_type -> counter.OperationOutput
	3, // 5: counter.Counter.Result:output_type -> counter.ResultOutput
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_counter_proto_init() }
func file_proto_counter_proto_init() {
	if File_proto_counter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_counter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_counter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationInput); i {
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
		file_proto_counter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationOutput); i {
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
		file_proto_counter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultOutput); i {
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
			RawDescriptor: file_proto_counter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_counter_proto_goTypes,
		DependencyIndexes: file_proto_counter_proto_depIdxs,
		MessageInfos:      file_proto_counter_proto_msgTypes,
	}.Build()
	File_proto_counter_proto = out.File
	file_proto_counter_proto_rawDesc = nil
	file_proto_counter_proto_goTypes = nil
	file_proto_counter_proto_depIdxs = nil
}
