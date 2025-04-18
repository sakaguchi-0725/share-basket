// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: personal/auth.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_personal_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_personal_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_personal_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personal_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_personal_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type SignUpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	mi := &file_personal_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_personal_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignUpConfirmRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Email            string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmationCode string                 `protobuf:"bytes,2,opt,name=confirmation_code,json=confirmationCode,proto3" json:"confirmation_code,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *SignUpConfirmRequest) Reset() {
	*x = SignUpConfirmRequest{}
	mi := &file_personal_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpConfirmRequest) ProtoMessage() {}

func (x *SignUpConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personal_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpConfirmRequest.ProtoReflect.Descriptor instead.
func (*SignUpConfirmRequest) Descriptor() ([]byte, []int) {
	return file_personal_auth_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpConfirmRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpConfirmRequest) GetConfirmationCode() string {
	if x != nil {
		return x.ConfirmationCode
	}
	return ""
}

var File_personal_auth_proto protoreflect.FileDescriptor

const file_personal_auth_proto_rawDesc = "" +
	"\n" +
	"\x13personal/auth.proto\x12\bpersonal\x1a\x1bgoogle/protobuf/empty.proto\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"2\n" +
	"\rLoginResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"U\n" +
	"\rSignUpRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"Y\n" +
	"\x14SignUpConfirmRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12+\n" +
	"\x11confirmation_code\x18\x02 \x01(\tR\x10confirmationCode2\xcb\x01\n" +
	"\vAuthService\x128\n" +
	"\x05Login\x12\x16.personal.LoginRequest\x1a\x17.personal.LoginResponse\x129\n" +
	"\x06SignUp\x12\x17.personal.SignUpRequest\x1a\x16.google.protobuf.Empty\x12G\n" +
	"\rSignUpConfirm\x12\x1e.personal.SignUpConfirmRequest\x1a\x16.google.protobuf.EmptyBd\n" +
	"\fcom.personalB\tAuthProtoP\x01Z\tproto/gen\xa2\x02\x03PXX\xaa\x02\bPersonal\xca\x02\bPersonal\xe2\x02\x14Personal\\GPBMetadata\xea\x02\bPersonalb\x06proto3"

var (
	file_personal_auth_proto_rawDescOnce sync.Once
	file_personal_auth_proto_rawDescData []byte
)

func file_personal_auth_proto_rawDescGZIP() []byte {
	file_personal_auth_proto_rawDescOnce.Do(func() {
		file_personal_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_personal_auth_proto_rawDesc), len(file_personal_auth_proto_rawDesc)))
	})
	return file_personal_auth_proto_rawDescData
}

var file_personal_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_personal_auth_proto_goTypes = []any{
	(*LoginRequest)(nil),         // 0: personal.LoginRequest
	(*LoginResponse)(nil),        // 1: personal.LoginResponse
	(*SignUpRequest)(nil),        // 2: personal.SignUpRequest
	(*SignUpConfirmRequest)(nil), // 3: personal.SignUpConfirmRequest
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_personal_auth_proto_depIdxs = []int32{
	0, // 0: personal.AuthService.Login:input_type -> personal.LoginRequest
	2, // 1: personal.AuthService.SignUp:input_type -> personal.SignUpRequest
	3, // 2: personal.AuthService.SignUpConfirm:input_type -> personal.SignUpConfirmRequest
	1, // 3: personal.AuthService.Login:output_type -> personal.LoginResponse
	4, // 4: personal.AuthService.SignUp:output_type -> google.protobuf.Empty
	4, // 5: personal.AuthService.SignUpConfirm:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_personal_auth_proto_init() }
func file_personal_auth_proto_init() {
	if File_personal_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_personal_auth_proto_rawDesc), len(file_personal_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personal_auth_proto_goTypes,
		DependencyIndexes: file_personal_auth_proto_depIdxs,
		MessageInfos:      file_personal_auth_proto_msgTypes,
	}.Build()
	File_personal_auth_proto = out.File
	file_personal_auth_proto_goTypes = nil
	file_personal_auth_proto_depIdxs = nil
}
