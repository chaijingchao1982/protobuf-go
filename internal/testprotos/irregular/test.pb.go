// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains a message which references a message that implements the
// proto.Message interface but does not have the structure of a normal generated
// message.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/irregular/test.proto

package irregular

import (
	protoreflect "github.com/chaijingchao1982/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/chaijingchao1982/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionalMessage *IrregularMessage            `protobuf:"bytes,1,opt,name=optional_message,json=optionalMessage" json:"optional_message,omitempty"`
	RepeatedMessage []*IrregularMessage          `protobuf:"bytes,2,rep,name=repeated_message,json=repeatedMessage" json:"repeated_message,omitempty"`
	RequiredMessage *IrregularMessage            `protobuf:"bytes,3,req,name=required_message,json=requiredMessage" json:"required_message,omitempty"`
	MapMessage      map[string]*IrregularMessage `protobuf:"bytes,4,rep,name=map_message,json=mapMessage" json:"map_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are assignable to Union:
	//
	//	*Message_OneofMessage
	//	*Message_OneofAberrantMessage
	Union                   isMessage_Union             `protobuf_oneof:"union"`
	OptionalAberrantMessage *AberrantMessage            `protobuf:"bytes,7,opt,name=optional_aberrant_message,json=optionalAberrantMessage" json:"optional_aberrant_message,omitempty"`
	RepeatedAberrantMessage []*AberrantMessage          `protobuf:"bytes,8,rep,name=repeated_aberrant_message,json=repeatedAberrantMessage" json:"repeated_aberrant_message,omitempty"`
	RequiredAberrantMessage *AberrantMessage            `protobuf:"bytes,9,req,name=required_aberrant_message,json=requiredAberrantMessage" json:"required_aberrant_message,omitempty"`
	MapAberrantMessage      map[string]*AberrantMessage `protobuf:"bytes,10,rep,name=map_aberrant_message,json=mapAberrantMessage" json:"map_aberrant_message,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_irregular_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_irregular_test_proto_msgTypes[0]
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
	return file_internal_testprotos_irregular_test_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetOptionalMessage() *IrregularMessage {
	if x != nil {
		return x.OptionalMessage
	}
	return nil
}

func (x *Message) GetRepeatedMessage() []*IrregularMessage {
	if x != nil {
		return x.RepeatedMessage
	}
	return nil
}

func (x *Message) GetRequiredMessage() *IrregularMessage {
	if x != nil {
		return x.RequiredMessage
	}
	return nil
}

func (x *Message) GetMapMessage() map[string]*IrregularMessage {
	if x != nil {
		return x.MapMessage
	}
	return nil
}

func (m *Message) GetUnion() isMessage_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *Message) GetOneofMessage() *IrregularMessage {
	if x, ok := x.GetUnion().(*Message_OneofMessage); ok {
		return x.OneofMessage
	}
	return nil
}

func (x *Message) GetOneofAberrantMessage() *AberrantMessage {
	if x, ok := x.GetUnion().(*Message_OneofAberrantMessage); ok {
		return x.OneofAberrantMessage
	}
	return nil
}

func (x *Message) GetOptionalAberrantMessage() *AberrantMessage {
	if x != nil {
		return x.OptionalAberrantMessage
	}
	return nil
}

func (x *Message) GetRepeatedAberrantMessage() []*AberrantMessage {
	if x != nil {
		return x.RepeatedAberrantMessage
	}
	return nil
}

func (x *Message) GetRequiredAberrantMessage() *AberrantMessage {
	if x != nil {
		return x.RequiredAberrantMessage
	}
	return nil
}

func (x *Message) GetMapAberrantMessage() map[string]*AberrantMessage {
	if x != nil {
		return x.MapAberrantMessage
	}
	return nil
}

type isMessage_Union interface {
	isMessage_Union()
}

type Message_OneofMessage struct {
	OneofMessage *IrregularMessage `protobuf:"bytes,5,opt,name=oneof_message,json=oneofMessage,oneof"`
}

type Message_OneofAberrantMessage struct {
	OneofAberrantMessage *AberrantMessage `protobuf:"bytes,6,opt,name=oneof_aberrant_message,json=oneofAberrantMessage,oneof"`
}

func (*Message_OneofMessage) isMessage_Union() {}

func (*Message_OneofAberrantMessage) isMessage_Union() {}

var File_internal_testprotos_irregular_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_irregular_test_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x1a, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x2f, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x54,
	0x0a, 0x10, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69,
	0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49,
	0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x51, 0x0a, 0x0b, 0x6d, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x16, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x61,
	0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e,
	0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x14, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x64, 0x0a, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x2e, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x17, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x62,
	0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x64, 0x0a,
	0x19, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x62, 0x65, 0x72, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x41, 0x62, 0x65, 0x72, 0x72,
	0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x17, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x64, 0x0a, 0x19, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72,
	0x2e, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6a, 0x0a, 0x14, 0x6d, 0x61, 0x70,
	0x5f, 0x61, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x41, 0x62, 0x65,
	0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x12, 0x6d, 0x61, 0x70, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x68, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x2e, 0x49, 0x72, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x6f, 0x0a, 0x17, 0x4d, 0x61, 0x70, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x69, 0x72, 0x72, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x41, 0x62, 0x65, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x72, 0x72, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72,
}

var (
	file_internal_testprotos_irregular_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_irregular_test_proto_rawDescData = file_internal_testprotos_irregular_test_proto_rawDesc
)

func file_internal_testprotos_irregular_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_irregular_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_irregular_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_irregular_test_proto_rawDescData)
	})
	return file_internal_testprotos_irregular_test_proto_rawDescData
}

var file_internal_testprotos_irregular_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_testprotos_irregular_test_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: goproto.proto.irregular.Message
	nil,                      // 1: goproto.proto.irregular.Message.MapMessageEntry
	nil,                      // 2: goproto.proto.irregular.Message.MapAberrantMessageEntry
	(*IrregularMessage)(nil), // 3: goproto.proto.irregular.IrregularMessage
	(*AberrantMessage)(nil),  // 4: goproto.proto.irregular.AberrantMessage
}
var file_internal_testprotos_irregular_test_proto_depIdxs = []int32{
	3,  // 0: goproto.proto.irregular.Message.optional_message:type_name -> goproto.proto.irregular.IrregularMessage
	3,  // 1: goproto.proto.irregular.Message.repeated_message:type_name -> goproto.proto.irregular.IrregularMessage
	3,  // 2: goproto.proto.irregular.Message.required_message:type_name -> goproto.proto.irregular.IrregularMessage
	1,  // 3: goproto.proto.irregular.Message.map_message:type_name -> goproto.proto.irregular.Message.MapMessageEntry
	3,  // 4: goproto.proto.irregular.Message.oneof_message:type_name -> goproto.proto.irregular.IrregularMessage
	4,  // 5: goproto.proto.irregular.Message.oneof_aberrant_message:type_name -> goproto.proto.irregular.AberrantMessage
	4,  // 6: goproto.proto.irregular.Message.optional_aberrant_message:type_name -> goproto.proto.irregular.AberrantMessage
	4,  // 7: goproto.proto.irregular.Message.repeated_aberrant_message:type_name -> goproto.proto.irregular.AberrantMessage
	4,  // 8: goproto.proto.irregular.Message.required_aberrant_message:type_name -> goproto.proto.irregular.AberrantMessage
	2,  // 9: goproto.proto.irregular.Message.map_aberrant_message:type_name -> goproto.proto.irregular.Message.MapAberrantMessageEntry
	3,  // 10: goproto.proto.irregular.Message.MapMessageEntry.value:type_name -> goproto.proto.irregular.IrregularMessage
	4,  // 11: goproto.proto.irregular.Message.MapAberrantMessageEntry.value:type_name -> goproto.proto.irregular.AberrantMessage
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_internal_testprotos_irregular_test_proto_init() }
func file_internal_testprotos_irregular_test_proto_init() {
	if File_internal_testprotos_irregular_test_proto != nil {
		return
	}
	file_internal_testprotos_irregular_irregular_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_irregular_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_internal_testprotos_irregular_test_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_OneofMessage)(nil),
		(*Message_OneofAberrantMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_irregular_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_irregular_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_irregular_test_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_irregular_test_proto_msgTypes,
	}.Build()
	File_internal_testprotos_irregular_test_proto = out.File
	file_internal_testprotos_irregular_test_proto_rawDesc = nil
	file_internal_testprotos_irregular_test_proto_goTypes = nil
	file_internal_testprotos_irregular_test_proto_depIdxs = nil
}
