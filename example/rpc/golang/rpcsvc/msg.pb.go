// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: msg.proto

package rpcsvc

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

// The request message containing the user's name.
type CheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *CheckReq) Reset() {
	*x = CheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReq) ProtoMessage() {}

func (x *CheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReq.ProtoReflect.Descriptor instead.
func (*CheckReq) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *CheckReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type WordData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword  string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=Category,proto3" json:"Category,omitempty"`
	Position string `protobuf:"bytes,3,opt,name=Position,proto3" json:"Position,omitempty"`
	Level    string `protobuf:"bytes,4,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *WordData) Reset() {
	*x = WordData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordData) ProtoMessage() {}

func (x *WordData) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordData.ProtoReflect.Descriptor instead.
func (*WordData) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *WordData) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *WordData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *WordData) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *WordData) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

// The response message containing the greetings
type CheckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string      `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg       string      `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ReturnStr string      `protobuf:"bytes,3,opt,name=ReturnStr,proto3" json:"ReturnStr,omitempty"`
	WordList  []*WordData `protobuf:"bytes,4,rep,name=WordList,proto3" json:"WordList,omitempty"`
}

func (x *CheckResp) Reset() {
	*x = CheckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResp) ProtoMessage() {}

func (x *CheckResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResp.ProtoReflect.Descriptor instead.
func (*CheckResp) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *CheckResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CheckResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CheckResp) GetReturnStr() string {
	if x != nil {
		return x.ReturnStr
	}
	return ""
}

func (x *CheckResp) GetWordList() []*WordData {
	if x != nil {
		return x.WordList
	}
	return nil
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x70, 0x63,
	0x73, 0x76, 0x63, 0x22, 0x24, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x08, 0x57, 0x6f, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x7d, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x72, 0x12, 0x2c,
	0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x39, 0x0a, 0x07,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x73, 0x76, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x73, 0x76, 0x63, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x72, 0x70,
	0x63, 0x73, 0x76, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_msg_proto_goTypes = []interface{}{
	(*CheckReq)(nil),  // 0: rpcsvc.CheckReq
	(*WordData)(nil),  // 1: rpcsvc.WordData
	(*CheckResp)(nil), // 2: rpcsvc.CheckResp
}
var file_msg_proto_depIdxs = []int32{
	1, // 0: rpcsvc.CheckResp.WordList:type_name -> rpcsvc.WordData
	0, // 1: rpcsvc.Greeter.Check:input_type -> rpcsvc.CheckReq
	2, // 2: rpcsvc.Greeter.Check:output_type -> rpcsvc.CheckResp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckReq); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordData); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResp); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
