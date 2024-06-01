// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: model/model.proto

//option go_package = "db_proj/model/proto;model_pb"; // 更新完改为这个, 用于本文件被import

package model_pb

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin         *string `protobuf:"bytes,1,opt,name=uin" json:"uin,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password    *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	PhoneNumber *string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	Perm        *int32  `protobuf:"varint,5,opt,name=perm" json:"perm,omitempty"`
	Email       *string `protobuf:"bytes,6,opt,name=email" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUin() string {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *User) GetPerm() int32 {
	if x != nil && x.Perm != nil {
		return *x.Perm
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type Dish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Price    *float64 `protobuf:"fixed64,3,opt,name=price" json:"price,omitempty"`
	Discount *float64 `protobuf:"fixed64,4,opt,name=discount" json:"discount,omitempty"`
	Detail   *string  `protobuf:"bytes,5,opt,name=detail" json:"detail,omitempty"`
}

func (x *Dish) Reset() {
	*x = Dish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dish) ProtoMessage() {}

func (x *Dish) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dish.ProtoReflect.Descriptor instead.
func (*Dish) Descriptor() ([]byte, []int) {
	return file_model_model_proto_rawDescGZIP(), []int{1}
}

func (x *Dish) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Dish) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Dish) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *Dish) GetDiscount() float64 {
	if x != nil && x.Discount != nil {
		return *x.Discount
	}
	return 0
}

func (x *Dish) GetDetail() string {
	if x != nil && x.Detail != nil {
		return *x.Detail
	}
	return ""
}

var File_model_model_proto protoreflect.FileDescriptor

var file_model_model_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x95, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x72, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x65, 0x72, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x74, 0x0a, 0x04, 0x44, 0x69, 0x73, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x16, 0x5a, 0x14, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x62,
}

var (
	file_model_model_proto_rawDescOnce sync.Once
	file_model_model_proto_rawDescData = file_model_model_proto_rawDesc
)

func file_model_model_proto_rawDescGZIP() []byte {
	file_model_model_proto_rawDescOnce.Do(func() {
		file_model_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_proto_rawDescData)
	})
	return file_model_model_proto_rawDescData
}

var file_model_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_model_proto_goTypes = []interface{}{
	(*User)(nil), // 0: model.User
	(*Dish)(nil), // 1: model.Dish
}
var file_model_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_model_proto_init() }
func file_model_model_proto_init() {
	if File_model_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_model_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dish); i {
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
			RawDescriptor: file_model_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_proto_goTypes,
		DependencyIndexes: file_model_model_proto_depIdxs,
		MessageInfos:      file_model_model_proto_msgTypes,
	}.Build()
	File_model_model_proto = out.File
	file_model_model_proto_rawDesc = nil
	file_model_model_proto_goTypes = nil
	file_model_model_proto_depIdxs = nil
}
