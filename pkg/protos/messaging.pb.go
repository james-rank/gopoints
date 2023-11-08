// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: messaging.proto

package protos

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

type Request_Action int32

const (
	Request_ACTION_ADD    Request_Action = 0
	Request_ACTION_REMOVE Request_Action = 1
	Request_ACTION_RESUME Request_Action = 2
)

// Enum value maps for Request_Action.
var (
	Request_Action_name = map[int32]string{
		0: "ACTION_ADD",
		1: "ACTION_REMOVE",
		2: "ACTION_RESUME",
	}
	Request_Action_value = map[string]int32{
		"ACTION_ADD":    0,
		"ACTION_REMOVE": 1,
		"ACTION_RESUME": 2,
	}
)

func (x Request_Action) Enum() *Request_Action {
	p := new(Request_Action)
	*p = x
	return p
}

func (x Request_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_messaging_proto_enumTypes[0].Descriptor()
}

func (Request_Action) Type() protoreflect.EnumType {
	return &file_messaging_proto_enumTypes[0]
}

func (x Request_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Action.Descriptor instead.
func (Request_Action) EnumDescriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0, 0}
}

type Request_PointType int32

const (
	Request_POINT_PAUSE Request_PointType = 0
	Request_POINT_CRASH Request_PointType = 1
)

// Enum value maps for Request_PointType.
var (
	Request_PointType_name = map[int32]string{
		0: "POINT_PAUSE",
		1: "POINT_CRASH",
	}
	Request_PointType_value = map[string]int32{
		"POINT_PAUSE": 0,
		"POINT_CRASH": 1,
	}
)

func (x Request_PointType) Enum() *Request_PointType {
	p := new(Request_PointType)
	*p = x
	return p
}

func (x Request_PointType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_PointType) Descriptor() protoreflect.EnumDescriptor {
	return file_messaging_proto_enumTypes[1].Descriptor()
}

func (Request_PointType) Type() protoreflect.EnumType {
	return &file_messaging_proto_enumTypes[1]
}

func (x Request_PointType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_PointType.Descriptor instead.
func (Request_PointType) EnumDescriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0, 1}
}

type Reply_Response int32

const (
	Reply_RESPONSE_ACK  Reply_Response = 0
	Reply_RESPONSE_NACK Reply_Response = 1
)

// Enum value maps for Reply_Response.
var (
	Reply_Response_name = map[int32]string{
		0: "RESPONSE_ACK",
		1: "RESPONSE_NACK",
	}
	Reply_Response_value = map[string]int32{
		"RESPONSE_ACK":  0,
		"RESPONSE_NACK": 1,
	}
)

func (x Reply_Response) Enum() *Reply_Response {
	p := new(Reply_Response)
	*p = x
	return p
}

func (x Reply_Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reply_Response) Descriptor() protoreflect.EnumDescriptor {
	return file_messaging_proto_enumTypes[2].Descriptor()
}

func (Reply_Response) Type() protoreflect.EnumType {
	return &file_messaging_proto_enumTypes[2]
}

func (x Reply_Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reply_Response.Descriptor instead.
func (Reply_Response) EnumDescriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{1, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action Request_Action    `protobuf:"varint,1,opt,name=action,proto3,enum=Messaging.Request_Action" json:"action,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type   Request_PointType `protobuf:"varint,3,opt,name=type,proto3,enum=Messaging.Request_PointType" json:"type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAction() Request_Action {
	if x != nil {
		return x.Action
	}
	return Request_ACTION_ADD
}

func (x *Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Request) GetType() Request_PointType {
	if x != nil {
		return x.Type
	}
	return Request_POINT_PAUSE
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response Reply_Response `protobuf:"varint,1,opt,name=response,proto3,enum=Messaging.Reply_Response" json:"response,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetResponse() Reply_Response {
	if x != nil {
		return x.Response
	}
	return Reply_RESPONSE_ACK
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messaging_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_messaging_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_messaging_proto_rawDescGZIP(), []int{2}
}

func (x *Notification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_messaging_proto protoreflect.FileDescriptor

var file_messaging_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0xf1, 0x01, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x3e, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4d, 0x45, 0x10,
	0x02, 0x22, 0x2d, 0x0a, 0x09, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x41, 0x53, 0x48, 0x10, 0x01,
	0x22, 0x6f, 0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x4e, 0x41, 0x43, 0x4b, 0x10,
	0x01, 0x22, 0x22, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x65, 0x73, 0x2d, 0x72, 0x61, 0x6e, 0x6b, 0x2f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_messaging_proto_rawDescOnce sync.Once
	file_messaging_proto_rawDescData = file_messaging_proto_rawDesc
)

func file_messaging_proto_rawDescGZIP() []byte {
	file_messaging_proto_rawDescOnce.Do(func() {
		file_messaging_proto_rawDescData = protoimpl.X.CompressGZIP(file_messaging_proto_rawDescData)
	})
	return file_messaging_proto_rawDescData
}

var file_messaging_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_messaging_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messaging_proto_goTypes = []interface{}{
	(Request_Action)(0),    // 0: Messaging.Request.Action
	(Request_PointType)(0), // 1: Messaging.Request.PointType
	(Reply_Response)(0),    // 2: Messaging.Reply.Response
	(*Request)(nil),        // 3: Messaging.Request
	(*Reply)(nil),          // 4: Messaging.Reply
	(*Notification)(nil),   // 5: Messaging.Notification
}
var file_messaging_proto_depIdxs = []int32{
	0, // 0: Messaging.Request.action:type_name -> Messaging.Request.Action
	1, // 1: Messaging.Request.type:type_name -> Messaging.Request.PointType
	2, // 2: Messaging.Reply.response:type_name -> Messaging.Reply.Response
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_messaging_proto_init() }
func file_messaging_proto_init() {
	if File_messaging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messaging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_messaging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_messaging_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_messaging_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messaging_proto_goTypes,
		DependencyIndexes: file_messaging_proto_depIdxs,
		EnumInfos:         file_messaging_proto_enumTypes,
		MessageInfos:      file_messaging_proto_msgTypes,
	}.Build()
	File_messaging_proto = out.File
	file_messaging_proto_rawDesc = nil
	file_messaging_proto_goTypes = nil
	file_messaging_proto_depIdxs = nil
}
