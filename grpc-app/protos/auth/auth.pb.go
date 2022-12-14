// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/auth.proto

package auth

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

// global param
type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Info) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Info) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// request
type RequestRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *RequestRegister_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestRegister) Reset() {
	*x = RequestRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister) ProtoMessage() {}

func (x *RequestRegister) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister.ProtoReflect.Descriptor instead.
func (*RequestRegister) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RequestRegister) GetData() *RequestRegister_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type RequestLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *RequestLogin_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestLogin) Reset() {
	*x = RequestLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogin) ProtoMessage() {}

func (x *RequestLogin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogin.ProtoReflect.Descriptor instead.
func (*RequestLogin) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RequestLogin) GetData() *RequestLogin_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// response
type ResponseRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info                  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Data *ResponseRegister_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseRegister) Reset() {
	*x = ResponseRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister) ProtoMessage() {}

func (x *ResponseRegister) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister.ProtoReflect.Descriptor instead.
func (*ResponseRegister) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseRegister) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ResponseRegister) GetData() *ResponseRegister_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Info               `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Data *ResponseLogin_Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResponseLogin) Reset() {
	*x = ResponseLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogin) ProtoMessage() {}

func (x *ResponseLogin) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogin.ProtoReflect.Descriptor instead.
func (*ResponseLogin) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ResponseLogin) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ResponseLogin) GetData() *ResponseLogin_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type RequestRegister_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail    string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserName     string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPassword string `protobuf:"bytes,3,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
}

func (x *RequestRegister_Data) Reset() {
	*x = RequestRegister_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRegister_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRegister_Data) ProtoMessage() {}

func (x *RequestRegister_Data) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRegister_Data.ProtoReflect.Descriptor instead.
func (*RequestRegister_Data) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RequestRegister_Data) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *RequestRegister_Data) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RequestRegister_Data) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

type RequestLogin_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail      string `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserPassword   string `protobuf:"bytes,2,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	UserRememberMe bool   `protobuf:"varint,3,opt,name=user_remember_me,json=userRememberMe,proto3" json:"user_remember_me,omitempty"`
}

func (x *RequestLogin_Data) Reset() {
	*x = RequestLogin_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogin_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogin_Data) ProtoMessage() {}

func (x *RequestLogin_Data) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogin_Data.ProtoReflect.Descriptor instead.
func (*RequestLogin_Data) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RequestLogin_Data) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *RequestLogin_Data) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *RequestLogin_Data) GetUserRememberMe() bool {
	if x != nil {
		return x.UserRememberMe
	}
	return false
}

type ResponseRegister_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *ResponseRegister_Data) Reset() {
	*x = ResponseRegister_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRegister_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRegister_Data) ProtoMessage() {}

func (x *ResponseRegister_Data) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRegister_Data.ProtoReflect.Descriptor instead.
func (*ResponseRegister_Data) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ResponseRegister_Data) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ResponseRegister_Data) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

type ResponseLogin_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Session string `protobuf:"bytes,4,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *ResponseLogin_Data) Reset() {
	*x = ResponseLogin_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseLogin_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseLogin_Data) ProtoMessage() {}

func (x *ResponseLogin_Data) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseLogin_Data.ProtoReflect.Descriptor instead.
func (*ResponseLogin_Data) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{4, 0}
}

func (x *ResponseLogin_Data) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ResponseLogin_Data) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

var File_protos_auth_proto protoreflect.FileDescriptor

var file_protos_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x32, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xac, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x67, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xb3,
	0x01, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x74,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4d, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x42,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x32, 0x8a, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x52, 0x50, 0x43, 0x12, 0x43, 0x0a,
	0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0x00, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_auth_proto_rawDescOnce sync.Once
	file_protos_auth_proto_rawDescData = file_protos_auth_proto_rawDesc
)

func file_protos_auth_proto_rawDescGZIP() []byte {
	file_protos_auth_proto_rawDescOnce.Do(func() {
		file_protos_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_auth_proto_rawDescData)
	})
	return file_protos_auth_proto_rawDescData
}

var file_protos_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_auth_proto_goTypes = []interface{}{
	(*Info)(nil),                  // 0: protos.Info
	(*RequestRegister)(nil),       // 1: protos.RequestRegister
	(*RequestLogin)(nil),          // 2: protos.RequestLogin
	(*ResponseRegister)(nil),      // 3: protos.ResponseRegister
	(*ResponseLogin)(nil),         // 4: protos.ResponseLogin
	(*RequestRegister_Data)(nil),  // 5: protos.RequestRegister.Data
	(*RequestLogin_Data)(nil),     // 6: protos.RequestLogin.Data
	(*ResponseRegister_Data)(nil), // 7: protos.ResponseRegister.Data
	(*ResponseLogin_Data)(nil),    // 8: protos.ResponseLogin.Data
}
var file_protos_auth_proto_depIdxs = []int32{
	5, // 0: protos.RequestRegister.data:type_name -> protos.RequestRegister.Data
	6, // 1: protos.RequestLogin.data:type_name -> protos.RequestLogin.Data
	0, // 2: protos.ResponseRegister.info:type_name -> protos.Info
	7, // 3: protos.ResponseRegister.data:type_name -> protos.ResponseRegister.Data
	0, // 4: protos.ResponseLogin.info:type_name -> protos.Info
	8, // 5: protos.ResponseLogin.data:type_name -> protos.ResponseLogin.Data
	1, // 6: protos.AuthRPC.RegisterUser:input_type -> protos.RequestRegister
	2, // 7: protos.AuthRPC.LoginUser:input_type -> protos.RequestLogin
	3, // 8: protos.AuthRPC.RegisterUser:output_type -> protos.ResponseRegister
	4, // 9: protos.AuthRPC.LoginUser:output_type -> protos.ResponseLogin
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_auth_proto_init() }
func file_protos_auth_proto_init() {
	if File_protos_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_protos_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRegister); i {
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
		file_protos_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogin); i {
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
		file_protos_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRegister); i {
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
		file_protos_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLogin); i {
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
		file_protos_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRegister_Data); i {
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
		file_protos_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogin_Data); i {
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
		file_protos_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRegister_Data); i {
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
		file_protos_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseLogin_Data); i {
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
			RawDescriptor: file_protos_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_auth_proto_goTypes,
		DependencyIndexes: file_protos_auth_proto_depIdxs,
		MessageInfos:      file_protos_auth_proto_msgTypes,
	}.Build()
	File_protos_auth_proto = out.File
	file_protos_auth_proto_rawDesc = nil
	file_protos_auth_proto_goTypes = nil
	file_protos_auth_proto_depIdxs = nil
}
