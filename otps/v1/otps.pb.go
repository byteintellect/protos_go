// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: otps/v1/otps.proto

package otpsv1

import (
	v1 "github.com/byteintellect/protos_go/commons/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AuthOtpDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalId  string    `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Status      v1.Status `protobuf:"varint,2,opt,name=status,proto3,enum=commons.v1.Status" json:"status,omitempty"`
	PhoneNumber string    `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Otp         string    `protobuf:"bytes,4,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *AuthOtpDto) Reset() {
	*x = AuthOtpDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otps_v1_otps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthOtpDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthOtpDto) ProtoMessage() {}

func (x *AuthOtpDto) ProtoReflect() protoreflect.Message {
	mi := &file_otps_v1_otps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthOtpDto.ProtoReflect.Descriptor instead.
func (*AuthOtpDto) Descriptor() ([]byte, []int) {
	return file_otps_v1_otps_proto_rawDescGZIP(), []int{0}
}

func (x *AuthOtpDto) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *AuthOtpDto) GetStatus() v1.Status {
	if x != nil {
		return x.Status
	}
	return v1.Status(0)
}

func (x *AuthOtpDto) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AuthOtpDto) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type GetOtpForPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	ExternalId  string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *GetOtpForPhoneRequest) Reset() {
	*x = GetOtpForPhoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otps_v1_otps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOtpForPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOtpForPhoneRequest) ProtoMessage() {}

func (x *GetOtpForPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otps_v1_otps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOtpForPhoneRequest.ProtoReflect.Descriptor instead.
func (*GetOtpForPhoneRequest) Descriptor() ([]byte, []int) {
	return file_otps_v1_otps_proto_rawDescGZIP(), []int{1}
}

func (x *GetOtpForPhoneRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetOtpForPhoneRequest) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

type GetOtpForPhoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *AuthOtpDto `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetOtpForPhoneResponse) Reset() {
	*x = GetOtpForPhoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otps_v1_otps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOtpForPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOtpForPhoneResponse) ProtoMessage() {}

func (x *GetOtpForPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otps_v1_otps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOtpForPhoneResponse.ProtoReflect.Descriptor instead.
func (*GetOtpForPhoneResponse) Descriptor() ([]byte, []int) {
	return file_otps_v1_otps_proto_rawDescGZIP(), []int{2}
}

func (x *GetOtpForPhoneResponse) GetResponse() *AuthOtpDto {
	if x != nil {
		return x.Response
	}
	return nil
}

type AuthValidateOtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	OtpId       string `protobuf:"bytes,2,opt,name=otp_id,json=otpId,proto3" json:"otp_id,omitempty"`
	Otp         string `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *AuthValidateOtpRequest) Reset() {
	*x = AuthValidateOtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otps_v1_otps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthValidateOtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidateOtpRequest) ProtoMessage() {}

func (x *AuthValidateOtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_otps_v1_otps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidateOtpRequest.ProtoReflect.Descriptor instead.
func (*AuthValidateOtpRequest) Descriptor() ([]byte, []int) {
	return file_otps_v1_otps_proto_rawDescGZIP(), []int{3}
}

func (x *AuthValidateOtpRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AuthValidateOtpRequest) GetOtpId() string {
	if x != nil {
		return x.OtpId
	}
	return ""
}

func (x *AuthValidateOtpRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type AuthValidateOtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *AuthValidateOtpResponse) Reset() {
	*x = AuthValidateOtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_otps_v1_otps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthValidateOtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthValidateOtpResponse) ProtoMessage() {}

func (x *AuthValidateOtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_otps_v1_otps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthValidateOtpResponse.ProtoReflect.Descriptor instead.
func (*AuthValidateOtpResponse) Descriptor() ([]byte, []int) {
	return file_otps_v1_otps_proto_rawDescGZIP(), []int{4}
}

func (x *AuthValidateOtpResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_otps_v1_otps_proto protoreflect.FileDescriptor

var file_otps_v1_otps_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x74, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x70, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x74,
	0x70, 0x44, 0x74, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x74, 0x70, 0x22, 0x5a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70, 0x46,
	0x6f, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x22, 0x49, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70, 0x46, 0x6f, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6f, 0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x74, 0x70, 0x44,
	0x74, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x64, 0x0a, 0x16,
	0x41, 0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x74, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x74, 0x70, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x74, 0x70, 0x22, 0x2f, 0x0a, 0x17, 0x41, 0x75, 0x74, 0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x32, 0xdc, 0x01, 0x0a, 0x0a, 0x4f, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x6f,
	0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70, 0x46, 0x6f, 0x72,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f,
	0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x70, 0x46, 0x6f, 0x72,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x74, 0x70, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74,
	0x70, 0x12, 0x1f, 0x2e, 0x6f, 0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x74, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x1a, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x74, 0x70, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x42, 0x88, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x74, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x4f, 0x74, 0x70, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x79, 0x74, 0x65,
	0x69, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x5f, 0x67, 0x6f, 0x2f, 0x6f, 0x74, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x74, 0x70, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x74, 0x70, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x07, 0x4f, 0x74, 0x70, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4f,
	0x74, 0x70, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x08, 0x4f, 0x74, 0x70, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_otps_v1_otps_proto_rawDescOnce sync.Once
	file_otps_v1_otps_proto_rawDescData = file_otps_v1_otps_proto_rawDesc
)

func file_otps_v1_otps_proto_rawDescGZIP() []byte {
	file_otps_v1_otps_proto_rawDescOnce.Do(func() {
		file_otps_v1_otps_proto_rawDescData = protoimpl.X.CompressGZIP(file_otps_v1_otps_proto_rawDescData)
	})
	return file_otps_v1_otps_proto_rawDescData
}

var file_otps_v1_otps_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_otps_v1_otps_proto_goTypes = []interface{}{
	(*AuthOtpDto)(nil),              // 0: otps.v1.AuthOtpDto
	(*GetOtpForPhoneRequest)(nil),   // 1: otps.v1.GetOtpForPhoneRequest
	(*GetOtpForPhoneResponse)(nil),  // 2: otps.v1.GetOtpForPhoneResponse
	(*AuthValidateOtpRequest)(nil),  // 3: otps.v1.AuthValidateOtpRequest
	(*AuthValidateOtpResponse)(nil), // 4: otps.v1.AuthValidateOtpResponse
	(v1.Status)(0),                  // 5: commons.v1.Status
}
var file_otps_v1_otps_proto_depIdxs = []int32{
	5, // 0: otps.v1.AuthOtpDto.status:type_name -> commons.v1.Status
	0, // 1: otps.v1.GetOtpForPhoneResponse.response:type_name -> otps.v1.AuthOtpDto
	1, // 2: otps.v1.OtpService.GetOtp:input_type -> otps.v1.GetOtpForPhoneRequest
	3, // 3: otps.v1.OtpService.ValidateOtp:input_type -> otps.v1.AuthValidateOtpRequest
	2, // 4: otps.v1.OtpService.GetOtp:output_type -> otps.v1.GetOtpForPhoneResponse
	4, // 5: otps.v1.OtpService.ValidateOtp:output_type -> otps.v1.AuthValidateOtpResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_otps_v1_otps_proto_init() }
func file_otps_v1_otps_proto_init() {
	if File_otps_v1_otps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_otps_v1_otps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthOtpDto); i {
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
		file_otps_v1_otps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOtpForPhoneRequest); i {
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
		file_otps_v1_otps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOtpForPhoneResponse); i {
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
		file_otps_v1_otps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthValidateOtpRequest); i {
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
		file_otps_v1_otps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthValidateOtpResponse); i {
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
			RawDescriptor: file_otps_v1_otps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_otps_v1_otps_proto_goTypes,
		DependencyIndexes: file_otps_v1_otps_proto_depIdxs,
		MessageInfos:      file_otps_v1_otps_proto_msgTypes,
	}.Build()
	File_otps_v1_otps_proto = out.File
	file_otps_v1_otps_proto_rawDesc = nil
	file_otps_v1_otps_proto_goTypes = nil
	file_otps_v1_otps_proto_depIdxs = nil
}
