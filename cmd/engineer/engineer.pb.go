// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc3
// source: engineer.proto

package engineer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEngineerByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEngineerByIdRequest) Reset() {
	*x = GetEngineerByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEngineerByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEngineerByIdRequest) ProtoMessage() {}

func (x *GetEngineerByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEngineerByIdRequest.ProtoReflect.Descriptor instead.
func (*GetEngineerByIdRequest) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{0}
}

func (x *GetEngineerByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AddEngineerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender    int32  `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	CountryId int64  `protobuf:"varint,4,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *AddEngineerRequest) Reset() {
	*x = AddEngineerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEngineerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEngineerRequest) ProtoMessage() {}

func (x *AddEngineerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEngineerRequest.ProtoReflect.Descriptor instead.
func (*AddEngineerRequest) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{1}
}

func (x *AddEngineerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AddEngineerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AddEngineerRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *AddEngineerRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *AddEngineerRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Engineer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender    int32                  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	CountryId int64                  `protobuf:"varint,5,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	Title     string                 `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Engineer) Reset() {
	*x = Engineer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engineer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engineer) ProtoMessage() {}

func (x *Engineer) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engineer.ProtoReflect.Descriptor instead.
func (*Engineer) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{2}
}

func (x *Engineer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Engineer) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Engineer) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Engineer) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Engineer) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *Engineer) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Engineer) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Engineer) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateEngineerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Gender    int32  `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Title     string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateEngineerRequest) Reset() {
	*x = UpdateEngineerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEngineerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEngineerRequest) ProtoMessage() {}

func (x *UpdateEngineerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEngineerRequest.ProtoReflect.Descriptor instead.
func (*UpdateEngineerRequest) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEngineerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateEngineerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateEngineerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateEngineerRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UpdateEngineerRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DeleteEngineerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEngineerRequest) Reset() {
	*x = DeleteEngineerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEngineerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEngineerRequest) ProtoMessage() {}

func (x *DeleteEngineerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEngineerRequest.ProtoReflect.Descriptor instead.
func (*DeleteEngineerRequest) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteEngineerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Engineers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engineers []*Engineer `protobuf:"bytes,1,rep,name=engineers,proto3" json:"engineers,omitempty"`
}

func (x *Engineers) Reset() {
	*x = Engineers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engineer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Engineers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engineers) ProtoMessage() {}

func (x *Engineers) ProtoReflect() protoreflect.Message {
	mi := &file_engineer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Engineers.ProtoReflect.Descriptor instead.
func (*Engineers) Descriptor() ([]byte, []int) {
	return file_engineer_proto_rawDescGZIP(), []int{5}
}

func (x *Engineers) GetEngineers() []*Engineer {
	if x != nil {
		return x.Engineers
	}
	return nil
}

var File_engineer_proto protoreflect.FileDescriptor

var file_engineer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x99, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91,
	0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x09, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x52,
	0x09, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x73, 0x32, 0xe7, 0x02, 0x0a, 0x0f, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f,
	0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12,
	0x3c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x73, 0x12, 0x47, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x20, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12, 0x45, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x12,
	0x1f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x65, 0x72, 0x42, 0x2a, 0x5a, 0x28, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72,
	0x2d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engineer_proto_rawDescOnce sync.Once
	file_engineer_proto_rawDescData = file_engineer_proto_rawDesc
)

func file_engineer_proto_rawDescGZIP() []byte {
	file_engineer_proto_rawDescOnce.Do(func() {
		file_engineer_proto_rawDescData = protoimpl.X.CompressGZIP(file_engineer_proto_rawDescData)
	})
	return file_engineer_proto_rawDescData
}

var file_engineer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_engineer_proto_goTypes = []any{
	(*GetEngineerByIdRequest)(nil), // 0: engineer.GetEngineerByIdRequest
	(*AddEngineerRequest)(nil),     // 1: engineer.AddEngineerRequest
	(*Engineer)(nil),               // 2: engineer.Engineer
	(*UpdateEngineerRequest)(nil),  // 3: engineer.UpdateEngineerRequest
	(*DeleteEngineerRequest)(nil),  // 4: engineer.DeleteEngineerRequest
	(*Engineers)(nil),              // 5: engineer.Engineers
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 7: google.protobuf.Empty
}
var file_engineer_proto_depIdxs = []int32{
	6, // 0: engineer.Engineer.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: engineer.Engineer.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: engineer.Engineers.engineers:type_name -> engineer.Engineer
	1, // 3: engineer.EngineerService.AddEngineer:input_type -> engineer.AddEngineerRequest
	7, // 4: engineer.EngineerService.ListEngineers:input_type -> google.protobuf.Empty
	0, // 5: engineer.EngineerService.GetEngineerById:input_type -> engineer.GetEngineerByIdRequest
	3, // 6: engineer.EngineerService.UpdateEngineer:input_type -> engineer.UpdateEngineerRequest
	4, // 7: engineer.EngineerService.DeleteEngineer:input_type -> engineer.DeleteEngineerRequest
	2, // 8: engineer.EngineerService.AddEngineer:output_type -> engineer.Engineer
	5, // 9: engineer.EngineerService.ListEngineers:output_type -> engineer.Engineers
	2, // 10: engineer.EngineerService.GetEngineerById:output_type -> engineer.Engineer
	2, // 11: engineer.EngineerService.UpdateEngineer:output_type -> engineer.Engineer
	2, // 12: engineer.EngineerService.DeleteEngineer:output_type -> engineer.Engineer
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_engineer_proto_init() }
func file_engineer_proto_init() {
	if File_engineer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engineer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetEngineerByIdRequest); i {
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
		file_engineer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddEngineerRequest); i {
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
		file_engineer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Engineer); i {
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
		file_engineer_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateEngineerRequest); i {
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
		file_engineer_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteEngineerRequest); i {
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
		file_engineer_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Engineers); i {
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
			RawDescriptor: file_engineer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_engineer_proto_goTypes,
		DependencyIndexes: file_engineer_proto_depIdxs,
		MessageInfos:      file_engineer_proto_msgTypes,
	}.Build()
	File_engineer_proto = out.File
	file_engineer_proto_rawDesc = nil
	file_engineer_proto_goTypes = nil
	file_engineer_proto_depIdxs = nil
}
