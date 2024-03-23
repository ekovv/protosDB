// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: creator/creator.proto

package ekovv_create_v1_createv1

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

type CreateDBRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Login    string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	DbName   string `protobuf:"bytes,4,opt,name=dbName,proto3" json:"dbName,omitempty"`
	DbType   string `protobuf:"bytes,5,opt,name=dbType,proto3" json:"dbType,omitempty"`
}

func (x *CreateDBRequest) Reset() {
	*x = CreateDBRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creator_creator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDBRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDBRequest) ProtoMessage() {}

func (x *CreateDBRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creator_creator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDBRequest.ProtoReflect.Descriptor instead.
func (*CreateDBRequest) Descriptor() ([]byte, []int) {
	return file_creator_creator_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDBRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateDBRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateDBRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateDBRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *CreateDBRequest) GetDbType() string {
	if x != nil {
		return x.DbType
	}
	return ""
}

type CreateDBResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionString string `protobuf:"bytes,1,opt,name=connectionString,proto3" json:"connectionString,omitempty"`
}

func (x *CreateDBResponse) Reset() {
	*x = CreateDBResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_creator_creator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDBResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDBResponse) ProtoMessage() {}

func (x *CreateDBResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creator_creator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDBResponse.ProtoReflect.Descriptor instead.
func (*CreateDBResponse) Descriptor() ([]byte, []int) {
	return file_creator_creator_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDBResponse) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

var File_creator_creator_proto protoreflect.FileDescriptor

var file_creator_creator_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x87, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x4a, 0x0a, 0x07, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44,
	0x42, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x42, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x42, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x65, 0x6b, 0x6f, 0x76, 0x76, 0x2e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creator_creator_proto_rawDescOnce sync.Once
	file_creator_creator_proto_rawDescData = file_creator_creator_proto_rawDesc
)

func file_creator_creator_proto_rawDescGZIP() []byte {
	file_creator_creator_proto_rawDescOnce.Do(func() {
		file_creator_creator_proto_rawDescData = protoimpl.X.CompressGZIP(file_creator_creator_proto_rawDescData)
	})
	return file_creator_creator_proto_rawDescData
}

var file_creator_creator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_creator_creator_proto_goTypes = []interface{}{
	(*CreateDBRequest)(nil),  // 0: creator.CreateDBRequest
	(*CreateDBResponse)(nil), // 1: creator.CreateDBResponse
}
var file_creator_creator_proto_depIdxs = []int32{
	0, // 0: creator.Creator.CreateDB:input_type -> creator.CreateDBRequest
	1, // 1: creator.Creator.CreateDB:output_type -> creator.CreateDBResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_creator_creator_proto_init() }
func file_creator_creator_proto_init() {
	if File_creator_creator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_creator_creator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDBRequest); i {
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
		file_creator_creator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDBResponse); i {
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
			RawDescriptor: file_creator_creator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_creator_creator_proto_goTypes,
		DependencyIndexes: file_creator_creator_proto_depIdxs,
		MessageInfos:      file_creator_creator_proto_msgTypes,
	}.Build()
	File_creator_creator_proto = out.File
	file_creator_creator_proto_rawDesc = nil
	file_creator_creator_proto_goTypes = nil
	file_creator_creator_proto_depIdxs = nil
}
