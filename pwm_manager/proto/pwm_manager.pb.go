// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pwm_manager.proto

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

type PasswordMode int32

const (
	PasswordMode_Easy    PasswordMode = 0
	PasswordMode_Medium  PasswordMode = 1
	PasswordMode_Complex PasswordMode = 2
)

// Enum value maps for PasswordMode.
var (
	PasswordMode_name = map[int32]string{
		0: "Easy",
		1: "Medium",
		2: "Complex",
	}
	PasswordMode_value = map[string]int32{
		"Easy":    0,
		"Medium":  1,
		"Complex": 2,
	}
)

func (x PasswordMode) Enum() *PasswordMode {
	p := new(PasswordMode)
	*p = x
	return p
}

func (x PasswordMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PasswordMode) Descriptor() protoreflect.EnumDescriptor {
	return file_pwm_manager_proto_enumTypes[0].Descriptor()
}

func (PasswordMode) Type() protoreflect.EnumType {
	return &file_pwm_manager_proto_enumTypes[0]
}

func (x PasswordMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PasswordMode.Descriptor instead.
func (PasswordMode) EnumDescriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{0}
}

type ItemToUpdate int32

const (
	ItemToUpdate_Password ItemToUpdate = 0
	ItemToUpdate_Category ItemToUpdate = 1
	ItemToUpdate_Site     ItemToUpdate = 2
)

// Enum value maps for ItemToUpdate.
var (
	ItemToUpdate_name = map[int32]string{
		0: "Password",
		1: "Category",
		2: "Site",
	}
	ItemToUpdate_value = map[string]int32{
		"Password": 0,
		"Category": 1,
		"Site":     2,
	}
)

func (x ItemToUpdate) Enum() *ItemToUpdate {
	p := new(ItemToUpdate)
	*p = x
	return p
}

func (x ItemToUpdate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemToUpdate) Descriptor() protoreflect.EnumDescriptor {
	return file_pwm_manager_proto_enumTypes[1].Descriptor()
}

func (ItemToUpdate) Type() protoreflect.EnumType {
	return &file_pwm_manager_proto_enumTypes[1]
}

func (x ItemToUpdate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemToUpdate.Descriptor instead.
func (ItemToUpdate) EnumDescriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{1}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[0]
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
	return file_pwm_manager_proto_rawDescGZIP(), []int{0}
}

type GeneratedPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *GeneratedPassword) Reset() {
	*x = GeneratedPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratedPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratedPassword) ProtoMessage() {}

func (x *GeneratedPassword) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratedPassword.ProtoReflect.Descriptor instead.
func (*GeneratedPassword) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratedPassword) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GeneratePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int64        `protobuf:"varint,1,opt,name=Length,proto3" json:"Length,omitempty"`
	Mode   PasswordMode `protobuf:"varint,2,opt,name=Mode,proto3,enum=proto.PasswordMode" json:"Mode,omitempty"`
}

func (x *GeneratePasswordRequest) Reset() {
	*x = GeneratePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePasswordRequest) ProtoMessage() {}

func (x *GeneratePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePasswordRequest.ProtoReflect.Descriptor instead.
func (*GeneratePasswordRequest) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{2}
}

func (x *GeneratePasswordRequest) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *GeneratePasswordRequest) GetMode() PasswordMode {
	if x != nil {
		return x.Mode
	}
	return PasswordMode_Easy
}

type ManagerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken string        `protobuf:"bytes,1,opt,name=JwtToken,proto3" json:"JwtToken,omitempty"`
	Password *PasswordData `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *ManagerRequest) Reset() {
	*x = ManagerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerRequest) ProtoMessage() {}

func (x *ManagerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerRequest.ProtoReflect.Descriptor instead.
func (*ManagerRequest) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{3}
}

func (x *ManagerRequest) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *ManagerRequest) GetPassword() *PasswordData {
	if x != nil {
		return x.Password
	}
	return nil
}

type PasswordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearTextPassword string `protobuf:"bytes,1,opt,name=ClearTextPassword,proto3" json:"ClearTextPassword,omitempty"`
	Category          string `protobuf:"bytes,2,opt,name=Category,proto3" json:"Category,omitempty"`
	Site              string `protobuf:"bytes,3,opt,name=Site,proto3" json:"Site,omitempty"`
}

func (x *PasswordData) Reset() {
	*x = PasswordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordData) ProtoMessage() {}

func (x *PasswordData) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordData.ProtoReflect.Descriptor instead.
func (*PasswordData) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{4}
}

func (x *PasswordData) GetClearTextPassword() string {
	if x != nil {
		return x.ClearTextPassword
	}
	return ""
}

func (x *PasswordData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *PasswordData) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

type ManagerUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken   string       `protobuf:"bytes,1,opt,name=JwtToken,proto3" json:"JwtToken,omitempty"`
	PasswordID int64        `protobuf:"varint,2,opt,name=PasswordID,proto3" json:"PasswordID,omitempty"`
	Mode       ItemToUpdate `protobuf:"varint,3,opt,name=Mode,proto3,enum=proto.ItemToUpdate" json:"Mode,omitempty"`
}

func (x *ManagerUpdateRequest) Reset() {
	*x = ManagerUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerUpdateRequest) ProtoMessage() {}

func (x *ManagerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerUpdateRequest.ProtoReflect.Descriptor instead.
func (*ManagerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{5}
}

func (x *ManagerUpdateRequest) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *ManagerUpdateRequest) GetPasswordID() int64 {
	if x != nil {
		return x.PasswordID
	}
	return 0
}

func (x *ManagerUpdateRequest) GetMode() ItemToUpdate {
	if x != nil {
		return x.Mode
	}
	return ItemToUpdate_Password
}

type PasswordItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PasswordID int64         `protobuf:"varint,1,opt,name=PasswordID,proto3" json:"PasswordID,omitempty"`
	Data       *PasswordData `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PasswordItem) Reset() {
	*x = PasswordItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasswordItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordItem) ProtoMessage() {}

func (x *PasswordItem) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordItem.ProtoReflect.Descriptor instead.
func (*PasswordItem) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{6}
}

func (x *PasswordItem) GetPasswordID() int64 {
	if x != nil {
		return x.PasswordID
	}
	return 0
}

func (x *PasswordItem) GetData() *PasswordData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UserPasswords struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passwords []*PasswordItem `protobuf:"bytes,1,rep,name=Passwords,proto3" json:"Passwords,omitempty"`
}

func (x *UserPasswords) Reset() {
	*x = UserPasswords{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPasswords) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPasswords) ProtoMessage() {}

func (x *UserPasswords) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPasswords.ProtoReflect.Descriptor instead.
func (*UserPasswords) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{7}
}

func (x *UserPasswords) GetPasswords() []*PasswordItem {
	if x != nil {
		return x.Passwords
	}
	return nil
}

type GetPasswordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JwtToken string `protobuf:"bytes,1,opt,name=JwtToken,proto3" json:"JwtToken,omitempty"`
}

func (x *GetPasswordsRequest) Reset() {
	*x = GetPasswordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pwm_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPasswordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPasswordsRequest) ProtoMessage() {}

func (x *GetPasswordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pwm_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPasswordsRequest.ProtoReflect.Descriptor instead.
func (*GetPasswordsRequest) Descriptor() ([]byte, []int) {
	return file_pwm_manager_proto_rawDescGZIP(), []int{8}
}

func (x *GetPasswordsRequest) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

var File_pwm_manager_proto protoreflect.FileDescriptor

var file_pwm_manager_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x77, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x5a, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x4d, 0x6f, 0x64, 0x65,
	0x22, 0x5d, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x6c, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2c, 0x0a, 0x11, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x65, 0x78, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x54, 0x65, 0x78, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x69, 0x74, 0x65, 0x22, 0x7b, 0x0a,
	0x14, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49,
	0x44, 0x12, 0x27, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x0c, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x42, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x31, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x31, 0x0a, 0x0c, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x61,
	0x73, 0x79, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x10, 0x02, 0x2a, 0x34, 0x0a,
	0x0c, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x69, 0x74,
	0x65, 0x10, 0x02, 0x32, 0x93, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x4c, 0x0a, 0x10, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x34, 0x6b, 0x65, 0x30, 0x2f, 0x70, 0x77,
	0x6d, 0x2f, 0x70, 0x77, 0x6d, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pwm_manager_proto_rawDescOnce sync.Once
	file_pwm_manager_proto_rawDescData = file_pwm_manager_proto_rawDesc
)

func file_pwm_manager_proto_rawDescGZIP() []byte {
	file_pwm_manager_proto_rawDescOnce.Do(func() {
		file_pwm_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_pwm_manager_proto_rawDescData)
	})
	return file_pwm_manager_proto_rawDescData
}

var file_pwm_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pwm_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pwm_manager_proto_goTypes = []interface{}{
	(PasswordMode)(0),               // 0: proto.PasswordMode
	(ItemToUpdate)(0),               // 1: proto.ItemToUpdate
	(*Empty)(nil),                   // 2: proto.Empty
	(*GeneratedPassword)(nil),       // 3: proto.GeneratedPassword
	(*GeneratePasswordRequest)(nil), // 4: proto.GeneratePasswordRequest
	(*ManagerRequest)(nil),          // 5: proto.ManagerRequest
	(*PasswordData)(nil),            // 6: proto.PasswordData
	(*ManagerUpdateRequest)(nil),    // 7: proto.ManagerUpdateRequest
	(*PasswordItem)(nil),            // 8: proto.PasswordItem
	(*UserPasswords)(nil),           // 9: proto.UserPasswords
	(*GetPasswordsRequest)(nil),     // 10: proto.GetPasswordsRequest
}
var file_pwm_manager_proto_depIdxs = []int32{
	0,  // 0: proto.GeneratePasswordRequest.Mode:type_name -> proto.PasswordMode
	6,  // 1: proto.ManagerRequest.Password:type_name -> proto.PasswordData
	1,  // 2: proto.ManagerUpdateRequest.Mode:type_name -> proto.ItemToUpdate
	6,  // 3: proto.PasswordItem.Data:type_name -> proto.PasswordData
	8,  // 4: proto.UserPasswords.Passwords:type_name -> proto.PasswordItem
	5,  // 5: proto.Manager.StorePassword:input_type -> proto.ManagerRequest
	7,  // 6: proto.Manager.UpdatePassword:input_type -> proto.ManagerUpdateRequest
	10, // 7: proto.Manager.GetPasswords:input_type -> proto.GetPasswordsRequest
	4,  // 8: proto.Manager.GeneratePassword:input_type -> proto.GeneratePasswordRequest
	2,  // 9: proto.Manager.StorePassword:output_type -> proto.Empty
	8,  // 10: proto.Manager.UpdatePassword:output_type -> proto.PasswordItem
	9,  // 11: proto.Manager.GetPasswords:output_type -> proto.UserPasswords
	3,  // 12: proto.Manager.GeneratePassword:output_type -> proto.GeneratedPassword
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pwm_manager_proto_init() }
func file_pwm_manager_proto_init() {
	if File_pwm_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pwm_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pwm_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratedPassword); i {
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
		file_pwm_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneratePasswordRequest); i {
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
		file_pwm_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerRequest); i {
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
		file_pwm_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordData); i {
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
		file_pwm_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerUpdateRequest); i {
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
		file_pwm_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasswordItem); i {
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
		file_pwm_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPasswords); i {
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
		file_pwm_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPasswordsRequest); i {
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
			RawDescriptor: file_pwm_manager_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pwm_manager_proto_goTypes,
		DependencyIndexes: file_pwm_manager_proto_depIdxs,
		EnumInfos:         file_pwm_manager_proto_enumTypes,
		MessageInfos:      file_pwm_manager_proto_msgTypes,
	}.Build()
	File_pwm_manager_proto = out.File
	file_pwm_manager_proto_rawDesc = nil
	file_pwm_manager_proto_goTypes = nil
	file_pwm_manager_proto_depIdxs = nil
}