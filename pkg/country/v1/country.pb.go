// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: country.proto

package country

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Countries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country []*Country `protobuf:"bytes,1,rep,name=country,proto3" json:"country,omitempty"`
}

func (x *Countries) Reset() {
	*x = Countries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Countries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Countries) ProtoMessage() {}

func (x *Countries) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Countries.ProtoReflect.Descriptor instead.
func (*Countries) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{0}
}

func (x *Countries) GetCountry() []*Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type AddCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryName string `protobuf:"bytes,1,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
}

func (x *AddCountryRequest) Reset() {
	*x = AddCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCountryRequest) ProtoMessage() {}

func (x *AddCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCountryRequest.ProtoReflect.Descriptor instead.
func (*AddCountryRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{1}
}

func (x *AddCountryRequest) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CountryName string               `protobuf:"bytes,2,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{2}
}

func (x *Country) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Country) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Country) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Country) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UpdateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CountryName string `protobuf:"bytes,2,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
}

func (x *UpdateCountryRequest) Reset() {
	*x = UpdateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryRequest) ProtoMessage() {}

func (x *UpdateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCountryRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCountryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCountryRequest) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

type DeleteCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCountryRequest) Reset() {
	*x = DeleteCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCountryRequest) ProtoMessage() {}

func (x *DeleteCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCountryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCountryRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCountryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCountryRequest) Reset() {
	*x = GetCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryRequest) ProtoMessage() {}

func (x *GetCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryRequest.ProtoReflect.Descriptor instead.
func (*GetCountryRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{5}
}

func (x *GetCountryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListCountriesRequest) Reset() {
	*x = ListCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesRequest) ProtoMessage() {}

func (x *ListCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesRequest.ProtoReflect.Descriptor instead.
func (*ListCountriesRequest) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{6}
}

func (x *ListCountriesRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListCountriesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Countries  *Countries `protobuf:"bytes,1,opt,name=countries,proto3" json:"countries,omitempty"`
	TotalCount int64      `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	PageNumber int32      `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32      `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListCountriesResponse) Reset() {
	*x = ListCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_country_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesResponse) ProtoMessage() {}

func (x *ListCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesResponse.ProtoReflect.Descriptor instead.
func (*ListCountriesResponse) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{7}
}

func (x *ListCountriesResponse) GetCountries() *Countries {
	if x != nil {
		return x.Countries
	}
	return nil
}

func (x *ListCountriesResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListCountriesResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListCountriesResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_country_proto protoreflect.FileDescriptor

var file_country_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x09,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x36, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xb2, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x32, 0xfe, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x2c, 0x5a, 0x2a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x2d, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_country_proto_rawDescOnce sync.Once
	file_country_proto_rawDescData = file_country_proto_rawDesc
)

func file_country_proto_rawDescGZIP() []byte {
	file_country_proto_rawDescOnce.Do(func() {
		file_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_country_proto_rawDescData)
	})
	return file_country_proto_rawDescData
}

var file_country_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_country_proto_goTypes = []any{
	(*Countries)(nil),             // 0: country.v1.Countries
	(*AddCountryRequest)(nil),     // 1: country.v1.AddCountryRequest
	(*Country)(nil),               // 2: country.v1.Country
	(*UpdateCountryRequest)(nil),  // 3: country.v1.UpdateCountryRequest
	(*DeleteCountryRequest)(nil),  // 4: country.v1.DeleteCountryRequest
	(*GetCountryRequest)(nil),     // 5: country.v1.GetCountryRequest
	(*ListCountriesRequest)(nil),  // 6: country.v1.ListCountriesRequest
	(*ListCountriesResponse)(nil), // 7: country.v1.ListCountriesResponse
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_country_proto_depIdxs = []int32{
	2, // 0: country.v1.Countries.country:type_name -> country.v1.Country
	8, // 1: country.v1.Country.created_at:type_name -> google.protobuf.Timestamp
	8, // 2: country.v1.Country.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: country.v1.ListCountriesResponse.countries:type_name -> country.v1.Countries
	1, // 4: country.v1.CountryService.AddCountry:input_type -> country.v1.AddCountryRequest
	6, // 5: country.v1.CountryService.ListCountries:input_type -> country.v1.ListCountriesRequest
	5, // 6: country.v1.CountryService.GetCountryById:input_type -> country.v1.GetCountryRequest
	3, // 7: country.v1.CountryService.UpdateCountry:input_type -> country.v1.UpdateCountryRequest
	4, // 8: country.v1.CountryService.DeleteCountry:input_type -> country.v1.DeleteCountryRequest
	2, // 9: country.v1.CountryService.AddCountry:output_type -> country.v1.Country
	7, // 10: country.v1.CountryService.ListCountries:output_type -> country.v1.ListCountriesResponse
	2, // 11: country.v1.CountryService.GetCountryById:output_type -> country.v1.Country
	2, // 12: country.v1.CountryService.UpdateCountry:output_type -> country.v1.Country
	2, // 13: country.v1.CountryService.DeleteCountry:output_type -> country.v1.Country
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_country_proto_init() }
func file_country_proto_init() {
	if File_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_country_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Countries); i {
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
		file_country_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddCountryRequest); i {
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
		file_country_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Country); i {
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
		file_country_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCountryRequest); i {
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
		file_country_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCountryRequest); i {
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
		file_country_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetCountryRequest); i {
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
		file_country_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListCountriesRequest); i {
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
		file_country_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListCountriesResponse); i {
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
			RawDescriptor: file_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_country_proto_goTypes,
		DependencyIndexes: file_country_proto_depIdxs,
		MessageInfos:      file_country_proto_msgTypes,
	}.Build()
	File_country_proto = out.File
	file_country_proto_rawDesc = nil
	file_country_proto_goTypes = nil
	file_country_proto_depIdxs = nil
}
