// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/fuzz/fuzz.proto

package fuzz

import (
	test "github.com/chaijingchao1982/protobuf/internal/testprotos/test"
	test3 "github.com/chaijingchao1982/protobuf/internal/testprotos/test3"
	protoreflect "github.com/chaijingchao1982/protobuf/reflect/protoreflect"
	protoimpl "github.com/chaijingchao1982/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

// Fuzz is a container for every message we want to make available to the fuzzer.
type Fuzz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestAllTypes            *test.TestAllTypes            `protobuf:"bytes,1,opt,name=test_all_types,json=testAllTypes" json:"test_all_types,omitempty"`
	TestAllExtensions       *test.TestAllExtensions       `protobuf:"bytes,2,opt,name=test_all_extensions,json=testAllExtensions" json:"test_all_extensions,omitempty"`
	TestRequired            *test.TestRequired            `protobuf:"bytes,3,opt,name=test_required,json=testRequired" json:"test_required,omitempty"`
	TestRequiredForeign     *test.TestRequiredForeign     `protobuf:"bytes,4,opt,name=test_required_foreign,json=testRequiredForeign" json:"test_required_foreign,omitempty"`
	TestRequiredGroupFields *test.TestRequiredGroupFields `protobuf:"bytes,5,opt,name=test_required_group_fields,json=testRequiredGroupFields" json:"test_required_group_fields,omitempty"`
	TestPackedTypes         *test.TestPackedTypes         `protobuf:"bytes,6,opt,name=test_packed_types,json=testPackedTypes" json:"test_packed_types,omitempty"`
	TestPackedExtensions    *test.TestPackedExtensions    `protobuf:"bytes,7,opt,name=test_packed_extensions,json=testPackedExtensions" json:"test_packed_extensions,omitempty"`
	TestAllTypes3           *test3.TestAllTypes           `protobuf:"bytes,8,opt,name=test_all_types3,json=testAllTypes3" json:"test_all_types3,omitempty"`
}

func (x *Fuzz) Reset() {
	*x = Fuzz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_fuzz_fuzz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fuzz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fuzz) ProtoMessage() {}

func (x *Fuzz) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_fuzz_fuzz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fuzz.ProtoReflect.Descriptor instead.
func (*Fuzz) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_fuzz_fuzz_proto_rawDescGZIP(), []int{0}
}

func (x *Fuzz) GetTestAllTypes() *test.TestAllTypes {
	if x != nil {
		return x.TestAllTypes
	}
	return nil
}

func (x *Fuzz) GetTestAllExtensions() *test.TestAllExtensions {
	if x != nil {
		return x.TestAllExtensions
	}
	return nil
}

func (x *Fuzz) GetTestRequired() *test.TestRequired {
	if x != nil {
		return x.TestRequired
	}
	return nil
}

func (x *Fuzz) GetTestRequiredForeign() *test.TestRequiredForeign {
	if x != nil {
		return x.TestRequiredForeign
	}
	return nil
}

func (x *Fuzz) GetTestRequiredGroupFields() *test.TestRequiredGroupFields {
	if x != nil {
		return x.TestRequiredGroupFields
	}
	return nil
}

func (x *Fuzz) GetTestPackedTypes() *test.TestPackedTypes {
	if x != nil {
		return x.TestPackedTypes
	}
	return nil
}

func (x *Fuzz) GetTestPackedExtensions() *test.TestPackedExtensions {
	if x != nil {
		return x.TestPackedExtensions
	}
	return nil
}

func (x *Fuzz) GetTestAllTypes3() *test3.TestAllTypes {
	if x != nil {
		return x.TestAllTypes3
	}
	return nil
}

var File_internal_testprotos_fuzz_fuzz_proto protoreflect.FileDescriptor

var file_internal_testprotos_fuzz_fuzz_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x75, 0x7a, 0x7a, 0x2f, 0x66, 0x75, 0x7a, 0x7a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x66, 0x75, 0x7a, 0x7a, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x05, 0x0a, 0x04, 0x46, 0x75, 0x7a, 0x7a, 0x12, 0x46, 0x0a,
	0x0e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x74, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x0d,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x5b, 0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x52, 0x13, 0x74, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e,
	0x12, 0x68, 0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x52, 0x17, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0f, 0x74, 0x65, 0x73, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x16, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x14, 0x74, 0x65, 0x73, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49, 0x0a, 0x0f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x33, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x33, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x75, 0x7a, 0x7a,
}

var (
	file_internal_testprotos_fuzz_fuzz_proto_rawDescOnce sync.Once
	file_internal_testprotos_fuzz_fuzz_proto_rawDescData = file_internal_testprotos_fuzz_fuzz_proto_rawDesc
)

func file_internal_testprotos_fuzz_fuzz_proto_rawDescGZIP() []byte {
	file_internal_testprotos_fuzz_fuzz_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_fuzz_fuzz_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_fuzz_fuzz_proto_rawDescData)
	})
	return file_internal_testprotos_fuzz_fuzz_proto_rawDescData
}

var file_internal_testprotos_fuzz_fuzz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_fuzz_fuzz_proto_goTypes = []interface{}{
	(*Fuzz)(nil),                         // 0: goproto.proto.fuzz.Fuzz
	(*test.TestAllTypes)(nil),            // 1: goproto.proto.test.TestAllTypes
	(*test.TestAllExtensions)(nil),       // 2: goproto.proto.test.TestAllExtensions
	(*test.TestRequired)(nil),            // 3: goproto.proto.test.TestRequired
	(*test.TestRequiredForeign)(nil),     // 4: goproto.proto.test.TestRequiredForeign
	(*test.TestRequiredGroupFields)(nil), // 5: goproto.proto.test.TestRequiredGroupFields
	(*test.TestPackedTypes)(nil),         // 6: goproto.proto.test.TestPackedTypes
	(*test.TestPackedExtensions)(nil),    // 7: goproto.proto.test.TestPackedExtensions
	(*test3.TestAllTypes)(nil),           // 8: goproto.proto.test3.TestAllTypes
}
var file_internal_testprotos_fuzz_fuzz_proto_depIdxs = []int32{
	1, // 0: goproto.proto.fuzz.Fuzz.test_all_types:type_name -> goproto.proto.test.TestAllTypes
	2, // 1: goproto.proto.fuzz.Fuzz.test_all_extensions:type_name -> goproto.proto.test.TestAllExtensions
	3, // 2: goproto.proto.fuzz.Fuzz.test_required:type_name -> goproto.proto.test.TestRequired
	4, // 3: goproto.proto.fuzz.Fuzz.test_required_foreign:type_name -> goproto.proto.test.TestRequiredForeign
	5, // 4: goproto.proto.fuzz.Fuzz.test_required_group_fields:type_name -> goproto.proto.test.TestRequiredGroupFields
	6, // 5: goproto.proto.fuzz.Fuzz.test_packed_types:type_name -> goproto.proto.test.TestPackedTypes
	7, // 6: goproto.proto.fuzz.Fuzz.test_packed_extensions:type_name -> goproto.proto.test.TestPackedExtensions
	8, // 7: goproto.proto.fuzz.Fuzz.test_all_types3:type_name -> goproto.proto.test3.TestAllTypes
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_internal_testprotos_fuzz_fuzz_proto_init() }
func file_internal_testprotos_fuzz_fuzz_proto_init() {
	if File_internal_testprotos_fuzz_fuzz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_fuzz_fuzz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fuzz); i {
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
			RawDescriptor: file_internal_testprotos_fuzz_fuzz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_fuzz_fuzz_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_fuzz_fuzz_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_fuzz_fuzz_proto_msgTypes,
	}.Build()
	File_internal_testprotos_fuzz_fuzz_proto = out.File
	file_internal_testprotos_fuzz_fuzz_proto_rawDesc = nil
	file_internal_testprotos_fuzz_fuzz_proto_goTypes = nil
	file_internal_testprotos_fuzz_fuzz_proto_depIdxs = nil
}
