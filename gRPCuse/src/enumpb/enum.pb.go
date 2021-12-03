// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: second/enum.proto

package enumpb

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

//枚举
type Gender int32

const (
	Gender_NOT_SPECIFIED Gender = 0
	Gender_FEMALE        Gender = 1
	Gender_MALE          Gender = 2
	Gender_WOMAN         Gender = 1
	Gender_MAN           Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "NOT_SPECIFIED",
		1: "FEMALE",
		2: "MALE",
		// Duplicate value: 1: "WOMAN",
		// Duplicate value: 2: "MAN",
	}
	Gender_value = map[string]int32{
		"NOT_SPECIFIED": 0,
		"FEMALE":        1,
		"MALE":          2,
		"WOMAN":         1,
		"MAN":           2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_second_enum_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_second_enum_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_second_enum_proto_rawDescGZIP(), []int{0}
}

type EnumMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Gender Gender `protobuf:"varint,2,opt,name=gender,proto3,enum=example.second.Gender" json:"gender,omitempty"`
}

func (x *EnumMessage) Reset() {
	*x = EnumMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_second_enum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumMessage) ProtoMessage() {}

func (x *EnumMessage) ProtoReflect() protoreflect.Message {
	mi := &file_second_enum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumMessage.ProtoReflect.Descriptor instead.
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return file_second_enum_proto_rawDescGZIP(), []int{0}
}

func (x *EnumMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EnumMessage) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_NOT_SPECIFIED
}

var File_second_enum_proto protoreflect.FileDescriptor

var file_second_enum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x22, 0x4d, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2a, 0x71, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x11, 0x0a, 0x0d,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x04, 0x08,
	0x09, 0x10, 0x09, 0x22, 0x04, 0x08, 0x0a, 0x10, 0x0a, 0x22, 0x04, 0x08, 0x14, 0x10, 0x64, 0x22,
	0x09, 0x08, 0xc8, 0x01, 0x10, 0xff, 0xff, 0xff, 0xff, 0x07, 0x2a, 0x03, 0x42, 0x4f, 0x59, 0x2a,
	0x04, 0x47, 0x49, 0x52, 0x4c, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_second_enum_proto_rawDescOnce sync.Once
	file_second_enum_proto_rawDescData = file_second_enum_proto_rawDesc
)

func file_second_enum_proto_rawDescGZIP() []byte {
	file_second_enum_proto_rawDescOnce.Do(func() {
		file_second_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_second_enum_proto_rawDescData)
	})
	return file_second_enum_proto_rawDescData
}

var file_second_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_second_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_second_enum_proto_goTypes = []interface{}{
	(Gender)(0),         // 0: example.second.Gender
	(*EnumMessage)(nil), // 1: example.second.EnumMessage
}
var file_second_enum_proto_depIdxs = []int32{
	0, // 0: example.second.EnumMessage.gender:type_name -> example.second.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_second_enum_proto_init() }
func file_second_enum_proto_init() {
	if File_second_enum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_second_enum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumMessage); i {
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
			RawDescriptor: file_second_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_second_enum_proto_goTypes,
		DependencyIndexes: file_second_enum_proto_depIdxs,
		EnumInfos:         file_second_enum_proto_enumTypes,
		MessageInfos:      file_second_enum_proto_msgTypes,
	}.Build()
	File_second_enum_proto = out.File
	file_second_enum_proto_rawDesc = nil
	file_second_enum_proto_goTypes = nil
	file_second_enum_proto_depIdxs = nil
}
