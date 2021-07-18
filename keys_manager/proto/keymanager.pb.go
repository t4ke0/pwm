// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: keymanager.proto

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

type Mode int32

const (
	Mode_Server     Mode = 0
	Mode_ServerAuth Mode = 1
	Mode_User       Mode = 2
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "Server",
		1: "ServerAuth",
		2: "User",
	}
	Mode_value = map[string]int32{
		"Server":     0,
		"ServerAuth": 1,
		"User":       2,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_keymanager_proto_enumTypes[0].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_keymanager_proto_enumTypes[0]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_keymanager_proto_rawDescGZIP(), []int{0}
}

type KeyGenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Mode Mode  `protobuf:"varint,2,opt,name=mode,proto3,enum=proto.Mode" json:"mode,omitempty"`
}

func (x *KeyGenRequest) Reset() {
	*x = KeyGenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyGenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyGenRequest) ProtoMessage() {}

func (x *KeyGenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keymanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyGenRequest.ProtoReflect.Descriptor instead.
func (*KeyGenRequest) Descriptor() ([]byte, []int) {
	return file_keymanager_proto_rawDescGZIP(), []int{0}
}

func (x *KeyGenRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *KeyGenRequest) GetMode() Mode {
	if x != nil {
		return x.Mode
	}
	return Mode_Server
}

type KeyFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *KeyFetchRequest) Reset() {
	*x = KeyFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyFetchRequest) ProtoMessage() {}

func (x *KeyFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keymanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyFetchRequest.ProtoReflect.Descriptor instead.
func (*KeyFetchRequest) Descriptor() ([]byte, []int) {
	return file_keymanager_proto_rawDescGZIP(), []int{1}
}

func (x *KeyFetchRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type KeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *KeyResponse) Reset() {
	*x = KeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keymanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyResponse) ProtoMessage() {}

func (x *KeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keymanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyResponse.ProtoReflect.Descriptor instead.
func (*KeyResponse) Descriptor() ([]byte, []int) {
	return file_keymanager_proto_rawDescGZIP(), []int{2}
}

func (x *KeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_keymanager_proto protoreflect.FileDescriptor

var file_keymanager_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0d, 0x4b, 0x65, 0x79,
	0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x2d, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f,
	0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x2a,
	0x2c, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x10, 0x02, 0x32, 0x7e, 0x0a,
	0x0a, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a,
	0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x34, 0x6b, 0x65,
	0x30, 0x2f, 0x70, 0x77, 0x6d, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keymanager_proto_rawDescOnce sync.Once
	file_keymanager_proto_rawDescData = file_keymanager_proto_rawDesc
)

func file_keymanager_proto_rawDescGZIP() []byte {
	file_keymanager_proto_rawDescOnce.Do(func() {
		file_keymanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_keymanager_proto_rawDescData)
	})
	return file_keymanager_proto_rawDescData
}

var file_keymanager_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_keymanager_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_keymanager_proto_goTypes = []interface{}{
	(Mode)(0),               // 0: proto.Mode
	(*KeyGenRequest)(nil),   // 1: proto.KeyGenRequest
	(*KeyFetchRequest)(nil), // 2: proto.KeyFetchRequest
	(*KeyResponse)(nil),     // 3: proto.KeyResponse
}
var file_keymanager_proto_depIdxs = []int32{
	0, // 0: proto.KeyGenRequest.mode:type_name -> proto.Mode
	1, // 1: proto.KeyManager.GenKey:input_type -> proto.KeyGenRequest
	2, // 2: proto.KeyManager.GetUserKey:input_type -> proto.KeyFetchRequest
	3, // 3: proto.KeyManager.GenKey:output_type -> proto.KeyResponse
	3, // 4: proto.KeyManager.GetUserKey:output_type -> proto.KeyResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keymanager_proto_init() }
func file_keymanager_proto_init() {
	if File_keymanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keymanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyGenRequest); i {
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
		file_keymanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyFetchRequest); i {
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
		file_keymanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyResponse); i {
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
			RawDescriptor: file_keymanager_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keymanager_proto_goTypes,
		DependencyIndexes: file_keymanager_proto_depIdxs,
		EnumInfos:         file_keymanager_proto_enumTypes,
		MessageInfos:      file_keymanager_proto_msgTypes,
	}.Build()
	File_keymanager_proto = out.File
	file_keymanager_proto_rawDesc = nil
	file_keymanager_proto_goTypes = nil
	file_keymanager_proto_depIdxs = nil
}
