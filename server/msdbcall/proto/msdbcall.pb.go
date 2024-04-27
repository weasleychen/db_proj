// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: msdbcall.proto

package msdbcall

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

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin         *string `protobuf:"bytes,1,opt,name=uin" json:"uin,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password    *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	PhoneNumber *string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	Perm        *int32  `protobuf:"varint,5,opt,name=perm" json:"perm,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserReq) GetUin() string {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return ""
}

func (x *CreateUserReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *CreateUserReq) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *CreateUserReq) GetPerm() int32 {
	if x != nil && x.Perm != nil {
		return *x.Perm
	}
	return 0
}

type CreateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"` // 返回新插入的user的json
}

func (x *CreateUserResp) Reset() {
	*x = CreateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResp) ProtoMessage() {}

func (x *CreateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResp.ProtoReflect.Descriptor instead.
func (*CreateUserResp) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResp) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type AddDishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Price    *float64 `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	Discount *float64 `protobuf:"fixed64,3,opt,name=discount" json:"discount,omitempty"`
	Detail   *string  `protobuf:"bytes,4,opt,name=detail" json:"detail,omitempty"`
}

func (x *AddDishReq) Reset() {
	*x = AddDishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDishReq) ProtoMessage() {}

func (x *AddDishReq) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDishReq.ProtoReflect.Descriptor instead.
func (*AddDishReq) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{2}
}

func (x *AddDishReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AddDishReq) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *AddDishReq) GetDiscount() float64 {
	if x != nil && x.Discount != nil {
		return *x.Discount
	}
	return 0
}

func (x *AddDishReq) GetDetail() string {
	if x != nil && x.Detail != nil {
		return *x.Detail
	}
	return ""
}

type AddDishResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (x *AddDishResp) Reset() {
	*x = AddDishResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDishResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDishResp) ProtoMessage() {}

func (x *AddDishResp) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDishResp.ProtoReflect.Descriptor instead.
func (*AddDishResp) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{3}
}

func (x *AddDishResp) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type CheckUserPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password *string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (x *CheckUserPasswordReq) Reset() {
	*x = CheckUserPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserPasswordReq) ProtoMessage() {}

func (x *CheckUserPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserPasswordReq.ProtoReflect.Descriptor instead.
func (*CheckUserPasswordReq) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{4}
}

func (x *CheckUserPasswordReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CheckUserPasswordReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

type CheckUserPasswordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (x *CheckUserPasswordResp) Reset() {
	*x = CheckUserPasswordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserPasswordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserPasswordResp) ProtoMessage() {}

func (x *CheckUserPasswordResp) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserPasswordResp.ProtoReflect.Descriptor instead.
func (*CheckUserPasswordResp) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{5}
}

func (x *CheckUserPasswordResp) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

type CheckUserNameUniqueReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *CheckUserNameUniqueReq) Reset() {
	*x = CheckUserNameUniqueReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserNameUniqueReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserNameUniqueReq) ProtoMessage() {}

func (x *CheckUserNameUniqueReq) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserNameUniqueReq.ProtoReflect.Descriptor instead.
func (*CheckUserNameUniqueReq) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{6}
}

func (x *CheckUserNameUniqueReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type CheckUserNameUniqueResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unique *bool `protobuf:"varint,1,opt,name=unique" json:"unique,omitempty"`
}

func (x *CheckUserNameUniqueResp) Reset() {
	*x = CheckUserNameUniqueResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msdbcall_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserNameUniqueResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserNameUniqueResp) ProtoMessage() {}

func (x *CheckUserNameUniqueResp) ProtoReflect() protoreflect.Message {
	mi := &file_msdbcall_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserNameUniqueResp.ProtoReflect.Descriptor instead.
func (*CheckUserNameUniqueResp) Descriptor() ([]byte, []int) {
	return file_msdbcall_proto_rawDescGZIP(), []int{7}
}

func (x *CheckUserNameUniqueResp) GetUnique() bool {
	if x != nil && x.Unique != nil {
		return *x.Unique
	}
	return false
}

var File_msdbcall_proto protoreflect.FileDescriptor

var file_msdbcall_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x65, 0x72, 0x6d, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6a, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x44, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x14, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x32, 0xb5, 0x02, 0x0a, 0x08, 0x4d, 0x53, 0x44, 0x42, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x2e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6d, 0x73, 0x64,
	0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x68, 0x12,
	0x14, 0x2e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x11,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x2e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x5a, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x20, 0x2e, 0x6d, 0x73, 0x64, 0x62,
	0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d, 0x73,
	0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x6d, 0x73, 0x64, 0x62, 0x63, 0x61, 0x6c, 0x6c,
}

var (
	file_msdbcall_proto_rawDescOnce sync.Once
	file_msdbcall_proto_rawDescData = file_msdbcall_proto_rawDesc
)

func file_msdbcall_proto_rawDescGZIP() []byte {
	file_msdbcall_proto_rawDescOnce.Do(func() {
		file_msdbcall_proto_rawDescData = protoimpl.X.CompressGZIP(file_msdbcall_proto_rawDescData)
	})
	return file_msdbcall_proto_rawDescData
}

var file_msdbcall_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_msdbcall_proto_goTypes = []interface{}{
	(*CreateUserReq)(nil),           // 0: msdbcall.CreateUserReq
	(*CreateUserResp)(nil),          // 1: msdbcall.CreateUserResp
	(*AddDishReq)(nil),              // 2: msdbcall.AddDishReq
	(*AddDishResp)(nil),             // 3: msdbcall.AddDishResp
	(*CheckUserPasswordReq)(nil),    // 4: msdbcall.CheckUserPasswordReq
	(*CheckUserPasswordResp)(nil),   // 5: msdbcall.CheckUserPasswordResp
	(*CheckUserNameUniqueReq)(nil),  // 6: msdbcall.CheckUserNameUniqueReq
	(*CheckUserNameUniqueResp)(nil), // 7: msdbcall.CheckUserNameUniqueResp
}
var file_msdbcall_proto_depIdxs = []int32{
	0, // 0: msdbcall.MSDBCall.CreateUser:input_type -> msdbcall.CreateUserReq
	2, // 1: msdbcall.MSDBCall.AddDish:input_type -> msdbcall.AddDishReq
	4, // 2: msdbcall.MSDBCall.CheckUserPassword:input_type -> msdbcall.CheckUserPasswordReq
	6, // 3: msdbcall.MSDBCall.CheckUserNameUnique:input_type -> msdbcall.CheckUserNameUniqueReq
	1, // 4: msdbcall.MSDBCall.CreateUser:output_type -> msdbcall.CreateUserResp
	3, // 5: msdbcall.MSDBCall.AddDish:output_type -> msdbcall.AddDishResp
	5, // 6: msdbcall.MSDBCall.CheckUserPassword:output_type -> msdbcall.CheckUserPasswordResp
	7, // 7: msdbcall.MSDBCall.CheckUserNameUnique:output_type -> msdbcall.CheckUserNameUniqueResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msdbcall_proto_init() }
func file_msdbcall_proto_init() {
	if File_msdbcall_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msdbcall_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReq); i {
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
		file_msdbcall_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResp); i {
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
		file_msdbcall_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDishReq); i {
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
		file_msdbcall_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDishResp); i {
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
		file_msdbcall_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserPasswordReq); i {
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
		file_msdbcall_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserPasswordResp); i {
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
		file_msdbcall_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserNameUniqueReq); i {
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
		file_msdbcall_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserNameUniqueResp); i {
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
			RawDescriptor: file_msdbcall_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msdbcall_proto_goTypes,
		DependencyIndexes: file_msdbcall_proto_depIdxs,
		MessageInfos:      file_msdbcall_proto_msgTypes,
	}.Build()
	File_msdbcall_proto = out.File
	file_msdbcall_proto_rawDesc = nil
	file_msdbcall_proto_goTypes = nil
	file_msdbcall_proto_depIdxs = nil
}
