// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: pokemon.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Pokemon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height         int32  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight         int32  `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	BaseExperience int32  `protobuf:"varint,5,opt,name=base_experience,json=baseExperience,proto3" json:"base_experience,omitempty"`
	Attack         int32  `protobuf:"varint,6,opt,name=attack,proto3" json:"attack,omitempty"`
	Defense        int32  `protobuf:"varint,7,opt,name=defense,proto3" json:"defense,omitempty"`
}

func (x *Pokemon) Reset() {
	*x = Pokemon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pokemon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pokemon) ProtoMessage() {}

func (x *Pokemon) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pokemon.ProtoReflect.Descriptor instead.
func (*Pokemon) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{0}
}

func (x *Pokemon) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pokemon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pokemon) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Pokemon) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Pokemon) GetBaseExperience() int32 {
	if x != nil {
		return x.BaseExperience
	}
	return 0
}

func (x *Pokemon) GetAttack() int32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *Pokemon) GetDefense() int32 {
	if x != nil {
		return x.Defense
	}
	return 0
}

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{1}
}

func (x *Params) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Reason         string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	ErrorUserTitle string `protobuf:"bytes,3,opt,name=error_user_title,json=errorUserTitle,proto3" json:"error_user_title,omitempty"`
	ErrorUserMsg   string `protobuf:"bytes,4,opt,name=error_user_msg,json=errorUserMsg,proto3" json:"error_user_msg,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorMessage) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ErrorMessage) GetErrorUserTitle() string {
	if x != nil {
		return x.ErrorUserTitle
	}
	return ""
}

func (x *ErrorMessage) GetErrorUserMsg() string {
	if x != nil {
		return x.ErrorUserMsg
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *Pokemon      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Error   *ErrorMessage `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pokemon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_pokemon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_pokemon_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Data) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Data) GetData() *Pokemon {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Data) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_pokemon_proto protoreflect.FileDescriptor

var file_pokemon_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x6b,
	0x65, 0x6d, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x66,
	0x65, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x66, 0x65,
	0x6e, 0x73, 0x65, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x24,
	0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x73, 0x67, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x3e, 0x0a, 0x0e, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6b, 0x65, 0x6d,
	0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6b, 0x65, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pokemon_proto_rawDescOnce sync.Once
	file_pokemon_proto_rawDescData = file_pokemon_proto_rawDesc
)

func file_pokemon_proto_rawDescGZIP() []byte {
	file_pokemon_proto_rawDescOnce.Do(func() {
		file_pokemon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pokemon_proto_rawDescData)
	})
	return file_pokemon_proto_rawDescData
}

var file_pokemon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pokemon_proto_goTypes = []interface{}{
	(*Pokemon)(nil),      // 0: pokemon.Pokemon
	(*Params)(nil),       // 1: pokemon.Params
	(*ErrorMessage)(nil), // 2: pokemon.ErrorMessage
	(*Data)(nil),         // 3: pokemon.Data
}
var file_pokemon_proto_depIdxs = []int32{
	0, // 0: pokemon.Data.data:type_name -> pokemon.Pokemon
	2, // 1: pokemon.Data.error:type_name -> pokemon.ErrorMessage
	1, // 2: pokemon.PokemonHandler.GetPokemon:input_type -> pokemon.Params
	3, // 3: pokemon.PokemonHandler.GetPokemon:output_type -> pokemon.Data
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pokemon_proto_init() }
func file_pokemon_proto_init() {
	if File_pokemon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pokemon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pokemon); i {
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
		file_pokemon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_pokemon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
		file_pokemon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_pokemon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pokemon_proto_goTypes,
		DependencyIndexes: file_pokemon_proto_depIdxs,
		MessageInfos:      file_pokemon_proto_msgTypes,
	}.Build()
	File_pokemon_proto = out.File
	file_pokemon_proto_rawDesc = nil
	file_pokemon_proto_goTypes = nil
	file_pokemon_proto_depIdxs = nil
}
