// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pkg/grpc/replica/replica.proto

package replica

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

type BidResponseStatus int32

const (
	BidResponseStatus_Success BidResponseStatus = 0
	BidResponseStatus_Fail    BidResponseStatus = 1
)

// Enum value maps for BidResponseStatus.
var (
	BidResponseStatus_name = map[int32]string{
		0: "Success",
		1: "Fail",
	}
	BidResponseStatus_value = map[string]int32{
		"Success": 0,
		"Fail":    1,
	}
)

func (x BidResponseStatus) Enum() *BidResponseStatus {
	p := new(BidResponseStatus)
	*p = x
	return p
}

func (x BidResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BidResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_grpc_replica_replica_proto_enumTypes[0].Descriptor()
}

func (BidResponseStatus) Type() protoreflect.EnumType {
	return &file_pkg_grpc_replica_replica_proto_enumTypes[0]
}

func (x BidResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BidResponseStatus.Descriptor instead.
func (BidResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_pkg_grpc_replica_replica_proto_rawDescGZIP(), []int{0}
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount uint64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_replica_replica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_replica_replica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_replica_replica_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status BidResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=replica.BidResponseStatus" json:"status,omitempty"`
}

func (x *BidResponse) Reset() {
	*x = BidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_replica_replica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidResponse) ProtoMessage() {}

func (x *BidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_replica_replica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidResponse.ProtoReflect.Descriptor instead.
func (*BidResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_replica_replica_proto_rawDescGZIP(), []int{1}
}

func (x *BidResponse) GetStatus() BidResponseStatus {
	if x != nil {
		return x.Status
	}
	return BidResponseStatus_Success
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_replica_replica_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_replica_replica_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_replica_replica_proto_rawDescGZIP(), []int{2}
}

var File_pkg_grpc_replica_replica_proto protoreflect.FileDescriptor

var file_pkg_grpc_replica_replica_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x0b, 0x42,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x06,
	0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x2a, 0x2a, 0x0a, 0x11, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c,
	0x10, 0x01, 0x32, 0xb1, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x2c,
	0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x2e, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x0f, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0f,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x05, 0x43, 0x72, 0x61, 0x73, 0x68, 0x12, 0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0d, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_replica_replica_proto_rawDescOnce sync.Once
	file_pkg_grpc_replica_replica_proto_rawDescData = file_pkg_grpc_replica_replica_proto_rawDesc
)

func file_pkg_grpc_replica_replica_proto_rawDescGZIP() []byte {
	file_pkg_grpc_replica_replica_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_replica_replica_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_replica_replica_proto_rawDescData)
	})
	return file_pkg_grpc_replica_replica_proto_rawDescData
}

var file_pkg_grpc_replica_replica_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_grpc_replica_replica_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_grpc_replica_replica_proto_goTypes = []interface{}{
	(BidResponseStatus)(0), // 0: replica.BidResponseStatus
	(*Amount)(nil),         // 1: replica.Amount
	(*BidResponse)(nil),    // 2: replica.BidResponse
	(*Void)(nil),           // 3: replica.Void
}
var file_pkg_grpc_replica_replica_proto_depIdxs = []int32{
	0, // 0: replica.BidResponse.status:type_name -> replica.BidResponseStatus
	1, // 1: replica.Replica.Bid:input_type -> replica.Amount
	3, // 2: replica.Replica.Result:input_type -> replica.Void
	1, // 3: replica.Replica.End:input_type -> replica.Amount
	3, // 4: replica.Replica.Crash:input_type -> replica.Void
	2, // 5: replica.Replica.Bid:output_type -> replica.BidResponse
	1, // 6: replica.Replica.Result:output_type -> replica.Amount
	1, // 7: replica.Replica.End:output_type -> replica.Amount
	3, // 8: replica.Replica.Crash:output_type -> replica.Void
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpc_replica_replica_proto_init() }
func file_pkg_grpc_replica_replica_proto_init() {
	if File_pkg_grpc_replica_replica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_replica_replica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_pkg_grpc_replica_replica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidResponse); i {
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
		file_pkg_grpc_replica_replica_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_pkg_grpc_replica_replica_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_replica_replica_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_replica_replica_proto_depIdxs,
		EnumInfos:         file_pkg_grpc_replica_replica_proto_enumTypes,
		MessageInfos:      file_pkg_grpc_replica_replica_proto_msgTypes,
	}.Build()
	File_pkg_grpc_replica_replica_proto = out.File
	file_pkg_grpc_replica_replica_proto_rawDesc = nil
	file_pkg_grpc_replica_replica_proto_goTypes = nil
	file_pkg_grpc_replica_replica_proto_depIdxs = nil
}
