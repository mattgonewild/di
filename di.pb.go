// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: dat.proto

package di

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

type CreateImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Private bool   `protobuf:"varint,3,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *CreateImgRequest) Reset() {
	*x = CreateImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImgRequest) ProtoMessage() {}

func (x *CreateImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImgRequest.ProtoReflect.Descriptor instead.
func (*CreateImgRequest) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{0}
}

func (x *CreateImgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateImgRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateImgRequest) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type CreateImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	BaseUrl string `protobuf:"bytes,2,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
}

func (x *CreateImgResponse) Reset() {
	*x = CreateImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImgResponse) ProtoMessage() {}

func (x *CreateImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImgResponse.ProtoReflect.Descriptor instead.
func (*CreateImgResponse) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateImgResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *CreateImgResponse) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

type ReadImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadImgRequest) Reset() {
	*x = ReadImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadImgRequest) ProtoMessage() {}

func (x *ReadImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadImgRequest.ProtoReflect.Descriptor instead.
func (*ReadImgRequest) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{2}
}

func (x *ReadImgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadImgResponse) Reset() {
	*x = ReadImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadImgResponse) ProtoMessage() {}

func (x *ReadImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadImgResponse.ProtoReflect.Descriptor instead.
func (*ReadImgResponse) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{3}
}

func (x *ReadImgResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ReadImgResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Private bool   `protobuf:"varint,2,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *UpdateImgRequest) Reset() {
	*x = UpdateImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImgRequest) ProtoMessage() {}

func (x *UpdateImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImgRequest.ProtoReflect.Descriptor instead.
func (*UpdateImgRequest) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateImgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateImgRequest) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

type UpdateImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateImgResponse) Reset() {
	*x = UpdateImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImgResponse) ProtoMessage() {}

func (x *UpdateImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImgResponse.ProtoReflect.Descriptor instead.
func (*UpdateImgResponse) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateImgResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type DeleteImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteImgRequest) Reset() {
	*x = DeleteImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImgRequest) ProtoMessage() {}

func (x *DeleteImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImgRequest.ProtoReflect.Descriptor instead.
func (*DeleteImgRequest) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteImgRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteImgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteImgResponse) Reset() {
	*x = DeleteImgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImgResponse) ProtoMessage() {}

func (x *DeleteImgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImgResponse.ProtoReflect.Descriptor instead.
func (*DeleteImgResponse) Descriptor() ([]byte, []int) {
	return file_dat_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteImgResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_dat_proto protoreflect.FileDescriptor

var file_dat_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22, 0x43, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x3c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x22,
	0x29, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd5, 0x01, 0x0a, 0x09, 0x44, 0x61,
	0x74, 0x61, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6d, 0x67, 0x12, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x52,
	0x65, 0x61, 0x64, 0x49, 0x6d, 0x67, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x49, 0x6d,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x12, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x61, 0x74, 0x74, 0x67, 0x6f, 0x6e, 0x65, 0x77, 0x69, 0x6c, 0x64, 0x2f, 0x64, 0x61, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dat_proto_rawDescOnce sync.Once
	file_dat_proto_rawDescData = file_dat_proto_rawDesc
)

func file_dat_proto_rawDescGZIP() []byte {
	file_dat_proto_rawDescOnce.Do(func() {
		file_dat_proto_rawDescData = protoimpl.X.CompressGZIP(file_dat_proto_rawDescData)
	})
	return file_dat_proto_rawDescData
}

var file_dat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dat_proto_goTypes = []interface{}{
	(*CreateImgRequest)(nil),  // 0: CreateImgRequest
	(*CreateImgResponse)(nil), // 1: CreateImgResponse
	(*ReadImgRequest)(nil),    // 2: ReadImgRequest
	(*ReadImgResponse)(nil),   // 3: ReadImgResponse
	(*UpdateImgRequest)(nil),  // 4: UpdateImgRequest
	(*UpdateImgResponse)(nil), // 5: UpdateImgResponse
	(*DeleteImgRequest)(nil),  // 6: DeleteImgRequest
	(*DeleteImgResponse)(nil), // 7: DeleteImgResponse
}
var file_dat_proto_depIdxs = []int32{
	0, // 0: DataLayer.CreateImg:input_type -> CreateImgRequest
	2, // 1: DataLayer.ReadImg:input_type -> ReadImgRequest
	4, // 2: DataLayer.UpdateImg:input_type -> UpdateImgRequest
	6, // 3: DataLayer.DeleteImg:input_type -> DeleteImgRequest
	1, // 4: DataLayer.CreateImg:output_type -> CreateImgResponse
	3, // 5: DataLayer.ReadImg:output_type -> ReadImgResponse
	5, // 6: DataLayer.UpdateImg:output_type -> UpdateImgResponse
	7, // 7: DataLayer.DeleteImg:output_type -> DeleteImgResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dat_proto_init() }
func file_dat_proto_init() {
	if File_dat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImgRequest); i {
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
		file_dat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImgResponse); i {
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
		file_dat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadImgRequest); i {
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
		file_dat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadImgResponse); i {
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
		file_dat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImgRequest); i {
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
		file_dat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImgResponse); i {
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
		file_dat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImgRequest); i {
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
		file_dat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImgResponse); i {
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
			RawDescriptor: file_dat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dat_proto_goTypes,
		DependencyIndexes: file_dat_proto_depIdxs,
		MessageInfos:      file_dat_proto_msgTypes,
	}.Build()
	File_dat_proto = out.File
	file_dat_proto_rawDesc = nil
	file_dat_proto_goTypes = nil
	file_dat_proto_depIdxs = nil
}
