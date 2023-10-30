// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/required/required.proto

package required

import (
	protoreflect "github.com/chaijingchao1982/protobuf/reflect/protoreflect"
	protoimpl "github.com/chaijingchao1982/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Int32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *int32 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
}

func (x *Int32) Reset() {
	*x = Int32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32) ProtoMessage() {}

func (x *Int32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32.ProtoReflect.Descriptor instead.
func (*Int32) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{0}
}

func (x *Int32) GetV() int32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Int64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *int64 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
}

func (x *Int64) Reset() {
	*x = Int64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64) ProtoMessage() {}

func (x *Int64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64.ProtoReflect.Descriptor instead.
func (*Int64) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{1}
}

func (x *Int64) GetV() int64 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Uint32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *uint32 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
}

func (x *Uint32) Reset() {
	*x = Uint32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32) ProtoMessage() {}

func (x *Uint32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint32.ProtoReflect.Descriptor instead.
func (*Uint32) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{2}
}

func (x *Uint32) GetV() uint32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Uint64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *uint64 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
}

func (x *Uint64) Reset() {
	*x = Uint64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uint64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint64) ProtoMessage() {}

func (x *Uint64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uint64.ProtoReflect.Descriptor instead.
func (*Uint64) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{3}
}

func (x *Uint64) GetV() uint64 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Sint32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *int32 `protobuf:"zigzag32,1,req,name=v" json:"v,omitempty"`
}

func (x *Sint32) Reset() {
	*x = Sint32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sint32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sint32) ProtoMessage() {}

func (x *Sint32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sint32.ProtoReflect.Descriptor instead.
func (*Sint32) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{4}
}

func (x *Sint32) GetV() int32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Sint64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *int64 `protobuf:"zigzag64,1,req,name=v" json:"v,omitempty"`
}

func (x *Sint64) Reset() {
	*x = Sint64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sint64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sint64) ProtoMessage() {}

func (x *Sint64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sint64.ProtoReflect.Descriptor instead.
func (*Sint64) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{5}
}

func (x *Sint64) GetV() int64 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Fixed32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *uint32 `protobuf:"fixed32,1,req,name=v" json:"v,omitempty"`
}

func (x *Fixed32) Reset() {
	*x = Fixed32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fixed32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed32) ProtoMessage() {}

func (x *Fixed32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fixed32.ProtoReflect.Descriptor instead.
func (*Fixed32) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{6}
}

func (x *Fixed32) GetV() uint32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Fixed64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *uint64 `protobuf:"fixed64,1,req,name=v" json:"v,omitempty"`
}

func (x *Fixed64) Reset() {
	*x = Fixed64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fixed64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed64) ProtoMessage() {}

func (x *Fixed64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fixed64.ProtoReflect.Descriptor instead.
func (*Fixed64) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{7}
}

func (x *Fixed64) GetV() uint64 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Float struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *float32 `protobuf:"fixed32,1,req,name=v" json:"v,omitempty"`
}

func (x *Float) Reset() {
	*x = Float{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float) ProtoMessage() {}

func (x *Float) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float.ProtoReflect.Descriptor instead.
func (*Float) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{8}
}

func (x *Float) GetV() float32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Double struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *float64 `protobuf:"fixed64,1,req,name=v" json:"v,omitempty"`
}

func (x *Double) Reset() {
	*x = Double{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Double) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Double) ProtoMessage() {}

func (x *Double) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Double.ProtoReflect.Descriptor instead.
func (*Double) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{9}
}

func (x *Double) GetV() float64 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

type Bool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *bool `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
}

func (x *Bool) Reset() {
	*x = Bool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bool.ProtoReflect.Descriptor instead.
func (*Bool) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{10}
}

func (x *Bool) GetV() bool {
	if x != nil && x.V != nil {
		return *x.V
	}
	return false
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *string `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{11}
}

func (x *String) GetV() string {
	if x != nil && x.V != nil {
		return *x.V
	}
	return ""
}

type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V []byte `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{12}
}

func (x *Bytes) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *Message_M `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{13}
}

func (x *Message) GetV() *Message_M {
	if x != nil {
		return x.V
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group_Group `protobuf:"group,1,req,name=Group,json=group" json:"group,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{14}
}

func (x *Group) GetGroup() *Group_Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type Message_M struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message_M) Reset() {
	*x = Message_M{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_M) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_M) ProtoMessage() {}

func (x *Message_M) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_M.ProtoReflect.Descriptor instead.
func (*Message_M) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{13, 0}
}

type Group_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *int32 `protobuf:"varint,1,opt,name=v" json:"v,omitempty"`
}

func (x *Group_Group) Reset() {
	*x = Group_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_required_required_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group_Group) ProtoMessage() {}

func (x *Group_Group) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group_Group.ProtoReflect.Descriptor instead.
func (*Group_Group) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_required_required_proto_rawDescGZIP(), []int{14, 0}
}

func (x *Group_Group) GetV() int32 {
	if x != nil && x.V != nil {
		return *x.V
	}
	return 0
}

var File_internal_testprotos_required_required_proto protoreflect.FileDescriptor

var file_internal_testprotos_required_required_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x05, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x01, 0x76,
	0x22, 0x15, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x01, 0x76, 0x22, 0x16, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x01, 0x76, 0x22,
	0x16, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x04, 0x52, 0x01, 0x76, 0x22, 0x16, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x11, 0x52, 0x01, 0x76, 0x22,
	0x16, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x12, 0x52, 0x01, 0x76, 0x22, 0x17, 0x0a, 0x07, 0x46, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x07, 0x52, 0x01, 0x76,
	0x22, 0x17, 0x0a, 0x07, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x76,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x06, 0x52, 0x01, 0x76, 0x22, 0x15, 0x0a, 0x05, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x02, 0x52, 0x01, 0x76,
	0x22, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x01, 0x52, 0x01, 0x76, 0x22, 0x14, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c,
	0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x01, 0x76, 0x22, 0x16,
	0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x01, 0x76, 0x22, 0x15, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x22, 0x43, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x52, 0x01, 0x76, 0x1a, 0x03, 0x0a,
	0x01, 0x4d, 0x22, 0x5d, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3d, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0a, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x15, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01,
	0x76, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
}

var (
	file_internal_testprotos_required_required_proto_rawDescOnce sync.Once
	file_internal_testprotos_required_required_proto_rawDescData = file_internal_testprotos_required_required_proto_rawDesc
)

func file_internal_testprotos_required_required_proto_rawDescGZIP() []byte {
	file_internal_testprotos_required_required_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_required_required_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_required_required_proto_rawDescData)
	})
	return file_internal_testprotos_required_required_proto_rawDescData
}

var file_internal_testprotos_required_required_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_internal_testprotos_required_required_proto_goTypes = []interface{}{
	(*Int32)(nil),       // 0: goproto.proto.testrequired.Int32
	(*Int64)(nil),       // 1: goproto.proto.testrequired.Int64
	(*Uint32)(nil),      // 2: goproto.proto.testrequired.Uint32
	(*Uint64)(nil),      // 3: goproto.proto.testrequired.Uint64
	(*Sint32)(nil),      // 4: goproto.proto.testrequired.Sint32
	(*Sint64)(nil),      // 5: goproto.proto.testrequired.Sint64
	(*Fixed32)(nil),     // 6: goproto.proto.testrequired.Fixed32
	(*Fixed64)(nil),     // 7: goproto.proto.testrequired.Fixed64
	(*Float)(nil),       // 8: goproto.proto.testrequired.Float
	(*Double)(nil),      // 9: goproto.proto.testrequired.Double
	(*Bool)(nil),        // 10: goproto.proto.testrequired.Bool
	(*String)(nil),      // 11: goproto.proto.testrequired.String
	(*Bytes)(nil),       // 12: goproto.proto.testrequired.Bytes
	(*Message)(nil),     // 13: goproto.proto.testrequired.Message
	(*Group)(nil),       // 14: goproto.proto.testrequired.Group
	(*Message_M)(nil),   // 15: goproto.proto.testrequired.Message.M
	(*Group_Group)(nil), // 16: goproto.proto.testrequired.Group.Group
}
var file_internal_testprotos_required_required_proto_depIdxs = []int32{
	15, // 0: goproto.proto.testrequired.Message.v:type_name -> goproto.proto.testrequired.Message.M
	16, // 1: goproto.proto.testrequired.Group.group:type_name -> goproto.proto.testrequired.Group.Group
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_internal_testprotos_required_required_proto_init() }
func file_internal_testprotos_required_required_proto_init() {
	if File_internal_testprotos_required_required_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_required_required_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint32); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uint64); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sint32); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sint64); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fixed32); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fixed64); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Float); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Double); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bool); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_M); i {
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
		file_internal_testprotos_required_required_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group_Group); i {
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
			RawDescriptor: file_internal_testprotos_required_required_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_required_required_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_required_required_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_required_required_proto_msgTypes,
	}.Build()
	File_internal_testprotos_required_required_proto = out.File
	file_internal_testprotos_required_required_proto_rawDesc = nil
	file_internal_testprotos_required_required_proto_goTypes = nil
	file_internal_testprotos_required_required_proto_depIdxs = nil
}
