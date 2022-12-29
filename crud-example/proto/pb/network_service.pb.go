// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: network_service.proto

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

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Rpc                  string `protobuf:"bytes,4,opt,name=rpc,proto3" json:"rpc,omitempty"`
	NativeTokenSymbol    string `protobuf:"bytes,5,opt,name=native_token_symbol,json=nativeTokenSymbol,proto3" json:"native_token_symbol,omitempty"`
	ImageUrl             string `protobuf:"bytes,6,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	EntryPointAddress    string `protobuf:"bytes,7,opt,name=entry_point_address,json=entryPointAddress,proto3" json:"entry_point_address,omitempty"`
	PreVerificationGas   uint64 `protobuf:"varint,8,opt,name=pre_verification_gas,json=preVerificationGas,proto3" json:"pre_verification_gas,omitempty"`
	VerificationGasLimit uint64 `protobuf:"varint,9,opt,name=verification_gas_limit,json=verificationGasLimit,proto3" json:"verification_gas_limit,omitempty"`
	DeployFee            uint64 `protobuf:"varint,10,opt,name=deploy_fee,json=deployFee,proto3" json:"deploy_fee,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Network) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Network) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Network) GetRpc() string {
	if x != nil {
		return x.Rpc
	}
	return ""
}

func (x *Network) GetNativeTokenSymbol() string {
	if x != nil {
		return x.NativeTokenSymbol
	}
	return ""
}

func (x *Network) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Network) GetEntryPointAddress() string {
	if x != nil {
		return x.EntryPointAddress
	}
	return ""
}

func (x *Network) GetPreVerificationGas() uint64 {
	if x != nil {
		return x.PreVerificationGas
	}
	return 0
}

func (x *Network) GetVerificationGasLimit() uint64 {
	if x != nil {
		return x.VerificationGasLimit
	}
	return 0
}

func (x *Network) GetDeployFee() uint64 {
	if x != nil {
		return x.DeployFee
	}
	return 0
}

type UpdateNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network *Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *UpdateNetworkRequest) Reset() {
	*x = UpdateNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetworkRequest) ProtoMessage() {}

func (x *UpdateNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetworkRequest.ProtoReflect.Descriptor instead.
func (*UpdateNetworkRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNetworkRequest) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

type UpdateNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network *Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *UpdateNetworkResponse) Reset() {
	*x = UpdateNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetworkResponse) ProtoMessage() {}

func (x *UpdateNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetworkResponse.ProtoReflect.Descriptor instead.
func (*UpdateNetworkResponse) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNetworkResponse) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

type CreateNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network *Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *CreateNetworkRequest) Reset() {
	*x = CreateNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkRequest) ProtoMessage() {}

func (x *CreateNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkRequest.ProtoReflect.Descriptor instead.
func (*CreateNetworkRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNetworkRequest) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

type CreateNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network *Network `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *CreateNetworkResponse) Reset() {
	*x = CreateNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNetworkResponse) ProtoMessage() {}

func (x *CreateNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNetworkResponse.ProtoReflect.Descriptor instead.
func (*CreateNetworkResponse) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNetworkResponse) GetNetwork() *Network {
	if x != nil {
		return x.Network
	}
	return nil
}

type GetListNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListNetworkRequest) Reset() {
	*x = GetListNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListNetworkRequest) ProtoMessage() {}

func (x *GetListNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListNetworkRequest.ProtoReflect.Descriptor instead.
func (*GetListNetworkRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{5}
}

type GetListNetworkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Networks []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
}

func (x *GetListNetworkResponse) Reset() {
	*x = GetListNetworkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListNetworkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListNetworkResponse) ProtoMessage() {}

func (x *GetListNetworkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListNetworkResponse.ProtoReflect.Descriptor instead.
func (*GetListNetworkResponse) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetListNetworkResponse) GetNetworks() []*Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

type GetNetworkDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNetworkDetailRequest) Reset() {
	*x = GetNetworkDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkDetailRequest) ProtoMessage() {}

func (x *GetNetworkDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkDetailRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkDetailRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetNetworkDetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteNetworkByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNetworkByIdRequest) Reset() {
	*x = DeleteNetworkByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkByIdRequest) ProtoMessage() {}

func (x *DeleteNetworkByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkByIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteNetworkByIdRequest) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteNetworkByIdRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteNetworkByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successful bool `protobuf:"varint,1,opt,name=successful,proto3" json:"successful,omitempty"`
}

func (x *DeleteNetworkByIdResponse) Reset() {
	*x = DeleteNetworkByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNetworkByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNetworkByIdResponse) ProtoMessage() {}

func (x *DeleteNetworkByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNetworkByIdResponse.ProtoReflect.Descriptor instead.
func (*DeleteNetworkByIdResponse) Descriptor() ([]byte, []int) {
	return file_network_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteNetworkByIdResponse) GetSuccessful() bool {
	if x != nil {
		return x.Successful
	}
	return false
}

var File_network_service_proto protoreflect.FileDescriptor

var file_network_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x22, 0xdb, 0x02, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x70, 0x63, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x70, 0x72, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x46, 0x65, 0x65, 0x22, 0x42,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x22, 0x43, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x43, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x66, 0x75, 0x6c, 0x32, 0xb1, 0x03, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x48, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x20, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x21, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_network_service_proto_rawDescOnce sync.Once
	file_network_service_proto_rawDescData = file_network_service_proto_rawDesc
)

func file_network_service_proto_rawDescGZIP() []byte {
	file_network_service_proto_rawDescOnce.Do(func() {
		file_network_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_service_proto_rawDescData)
	})
	return file_network_service_proto_rawDescData
}

var file_network_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_network_service_proto_goTypes = []interface{}{
	(*Network)(nil),                   // 0: network.Network
	(*UpdateNetworkRequest)(nil),      // 1: network.UpdateNetworkRequest
	(*UpdateNetworkResponse)(nil),     // 2: network.UpdateNetworkResponse
	(*CreateNetworkRequest)(nil),      // 3: network.CreateNetworkRequest
	(*CreateNetworkResponse)(nil),     // 4: network.CreateNetworkResponse
	(*GetListNetworkRequest)(nil),     // 5: network.GetListNetworkRequest
	(*GetListNetworkResponse)(nil),    // 6: network.GetListNetworkResponse
	(*GetNetworkDetailRequest)(nil),   // 7: network.GetNetworkDetailRequest
	(*DeleteNetworkByIdRequest)(nil),  // 8: network.DeleteNetworkByIdRequest
	(*DeleteNetworkByIdResponse)(nil), // 9: network.DeleteNetworkByIdResponse
}
var file_network_service_proto_depIdxs = []int32{
	0,  // 0: network.UpdateNetworkRequest.network:type_name -> network.Network
	0,  // 1: network.UpdateNetworkResponse.network:type_name -> network.Network
	0,  // 2: network.CreateNetworkRequest.network:type_name -> network.Network
	0,  // 3: network.CreateNetworkResponse.network:type_name -> network.Network
	0,  // 4: network.GetListNetworkResponse.networks:type_name -> network.Network
	3,  // 5: network.NetworkService.CreateNetwork:input_type -> network.CreateNetworkRequest
	5,  // 6: network.NetworkService.GetListNetwork:input_type -> network.GetListNetworkRequest
	7,  // 7: network.NetworkService.GetNetworkDetail:input_type -> network.GetNetworkDetailRequest
	1,  // 8: network.NetworkService.UpdateNetwork:input_type -> network.UpdateNetworkRequest
	8,  // 9: network.NetworkService.DeleteNetworkById:input_type -> network.DeleteNetworkByIdRequest
	4,  // 10: network.NetworkService.CreateNetwork:output_type -> network.CreateNetworkResponse
	6,  // 11: network.NetworkService.GetListNetwork:output_type -> network.GetListNetworkResponse
	0,  // 12: network.NetworkService.GetNetworkDetail:output_type -> network.Network
	2,  // 13: network.NetworkService.UpdateNetwork:output_type -> network.UpdateNetworkResponse
	9,  // 14: network.NetworkService.DeleteNetworkById:output_type -> network.DeleteNetworkByIdResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_network_service_proto_init() }
func file_network_service_proto_init() {
	if File_network_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_network_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_network_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNetworkRequest); i {
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
		file_network_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNetworkResponse); i {
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
		file_network_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkRequest); i {
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
		file_network_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNetworkResponse); i {
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
		file_network_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListNetworkRequest); i {
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
		file_network_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListNetworkResponse); i {
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
		file_network_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkDetailRequest); i {
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
		file_network_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkByIdRequest); i {
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
		file_network_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNetworkByIdResponse); i {
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
			RawDescriptor: file_network_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_network_service_proto_goTypes,
		DependencyIndexes: file_network_service_proto_depIdxs,
		MessageInfos:      file_network_service_proto_msgTypes,
	}.Build()
	File_network_service_proto = out.File
	file_network_service_proto_rawDesc = nil
	file_network_service_proto_goTypes = nil
	file_network_service_proto_depIdxs = nil
}
