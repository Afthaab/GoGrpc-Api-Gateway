// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/user/pb/user.proto

package pb

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

type ViewProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewProfileRequest) Reset() {
	*x = ViewProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewProfileRequest) ProtoMessage() {}

func (x *ViewProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewProfileRequest.ProtoReflect.Descriptor instead.
func (*ViewProfileRequest) Descriptor() ([]byte, []int) {
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *ViewProfileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ViewProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Profile  string `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile,omitempty"`
	Error    string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Gender   string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Dob      string `protobuf:"bytes,8,opt,name=dob,proto3" json:"dob,omitempty"`
}

func (x *ViewProfileResponse) Reset() {
	*x = ViewProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_user_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewProfileResponse) ProtoMessage() {}

func (x *ViewProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_user_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewProfileResponse.ProtoReflect.Descriptor instead.
func (*ViewProfileResponse) Descriptor() ([]byte, []int) {
	return file_pkg_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *ViewProfileResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ViewProfileResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ViewProfileResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ViewProfileResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ViewProfileResponse) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *ViewProfileResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ViewProfileResponse) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *ViewProfileResponse) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

var File_pkg_user_pb_user_proto protoreflect.FileDescriptor

var file_pkg_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x24, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x13, 0x56, 0x69, 0x65, 0x77,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x32, 0x5f, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4a,
	0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pkg_user_pb_user_proto_rawDescOnce sync.Once
	file_pkg_user_pb_user_proto_rawDescData = file_pkg_user_pb_user_proto_rawDesc
)

func file_pkg_user_pb_user_proto_rawDescGZIP() []byte {
	file_pkg_user_pb_user_proto_rawDescOnce.Do(func() {
		file_pkg_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_user_pb_user_proto_rawDescData)
	})
	return file_pkg_user_pb_user_proto_rawDescData
}

var file_pkg_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_user_pb_user_proto_goTypes = []interface{}{
	(*ViewProfileRequest)(nil),  // 0: profile.ViewProfileRequest
	(*ViewProfileResponse)(nil), // 1: profile.ViewProfileResponse
}
var file_pkg_user_pb_user_proto_depIdxs = []int32{
	0, // 0: profile.ProfileManagement.ViewProfile:input_type -> profile.ViewProfileRequest
	1, // 1: profile.ProfileManagement.ViewProfile:output_type -> profile.ViewProfileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_user_pb_user_proto_init() }
func file_pkg_user_pb_user_proto_init() {
	if File_pkg_user_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_user_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewProfileRequest); i {
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
		file_pkg_user_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewProfileResponse); i {
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
			RawDescriptor: file_pkg_user_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_user_pb_user_proto_goTypes,
		DependencyIndexes: file_pkg_user_pb_user_proto_depIdxs,
		MessageInfos:      file_pkg_user_pb_user_proto_msgTypes,
	}.Build()
	File_pkg_user_pb_user_proto = out.File
	file_pkg_user_pb_user_proto_rawDesc = nil
	file_pkg_user_pb_user_proto_goTypes = nil
	file_pkg_user_pb_user_proto_depIdxs = nil
}
