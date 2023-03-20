// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: hub.proto

package gen

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

type HubNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *HubNode) Reset() {
	*x = HubNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HubNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HubNode) ProtoMessage() {}

func (x *HubNode) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HubNode.ProtoReflect.Descriptor instead.
func (*HubNode) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{0}
}

func (x *HubNode) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HubNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HubNode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetHubNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetHubNodeRequest) Reset() {
	*x = GetHubNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHubNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHubNodeRequest) ProtoMessage() {}

func (x *GetHubNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHubNodeRequest.ProtoReflect.Descriptor instead.
func (*GetHubNodeRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{1}
}

func (x *GetHubNodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type EnrollHubNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node *HubNode `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *EnrollHubNodeRequest) Reset() {
	*x = EnrollHubNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollHubNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollHubNodeRequest) ProtoMessage() {}

func (x *EnrollHubNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollHubNodeRequest.ProtoReflect.Descriptor instead.
func (*EnrollHubNodeRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{2}
}

func (x *EnrollHubNodeRequest) GetNode() *HubNode {
	if x != nil {
		return x.Node
	}
	return nil
}

type ListHubNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHubNodesRequest) Reset() {
	*x = ListHubNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHubNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHubNodesRequest) ProtoMessage() {}

func (x *ListHubNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHubNodesRequest.ProtoReflect.Descriptor instead.
func (*ListHubNodesRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{3}
}

type ListHubNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*HubNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ListHubNodesResponse) Reset() {
	*x = ListHubNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHubNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHubNodesResponse) ProtoMessage() {}

func (x *ListHubNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHubNodesResponse.ProtoReflect.Descriptor instead.
func (*ListHubNodesResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{4}
}

func (x *ListHubNodesResponse) GetNodes() []*HubNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type TestHubNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TestHubNodeRequest) Reset() {
	*x = TestHubNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestHubNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestHubNodeRequest) ProtoMessage() {}

func (x *TestHubNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestHubNodeRequest.ProtoReflect.Descriptor instead.
func (*TestHubNodeRequest) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{5}
}

func (x *TestHubNodeRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type TestHubNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TestHubNodeResponse) Reset() {
	*x = TestHubNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hub_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestHubNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestHubNodeResponse) ProtoMessage() {}

func (x *TestHubNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hub_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestHubNodeResponse.ProtoReflect.Descriptor instead.
func (*TestHubNodeResponse) Descriptor() ([]byte, []int) {
	return file_hub_proto_rawDescGZIP(), []int{6}
}

var File_hub_proto protoreflect.FileDescriptor

var file_hub_proto_rawDesc = []byte{
	0x0a, 0x09, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x22, 0x59, 0x0a, 0x07, 0x48, 0x75,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x14, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x48, 0x75,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x48, 0x75, 0x62, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x12, 0x54, 0x65, 0x73, 0x74, 0x48, 0x75, 0x62,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x54, 0x65, 0x73, 0x74, 0x48, 0x75, 0x62,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xcd, 0x02, 0x0a,
	0x06, 0x48, 0x75, 0x62, 0x41, 0x70, 0x69, 0x12, 0x46, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x75,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x57, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x62, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x48,
	0x75, 0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x48, 0x75,
	0x62, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x75, 0x62, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x48, 0x75, 0x62,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6e, 0x74, 0x69,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x63, 0x2d, 0x62, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hub_proto_rawDescOnce sync.Once
	file_hub_proto_rawDescData = file_hub_proto_rawDesc
)

func file_hub_proto_rawDescGZIP() []byte {
	file_hub_proto_rawDescOnce.Do(func() {
		file_hub_proto_rawDescData = protoimpl.X.CompressGZIP(file_hub_proto_rawDescData)
	})
	return file_hub_proto_rawDescData
}

var file_hub_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hub_proto_goTypes = []interface{}{
	(*HubNode)(nil),              // 0: smartcore.bos.HubNode
	(*GetHubNodeRequest)(nil),    // 1: smartcore.bos.GetHubNodeRequest
	(*EnrollHubNodeRequest)(nil), // 2: smartcore.bos.EnrollHubNodeRequest
	(*ListHubNodesRequest)(nil),  // 3: smartcore.bos.ListHubNodesRequest
	(*ListHubNodesResponse)(nil), // 4: smartcore.bos.ListHubNodesResponse
	(*TestHubNodeRequest)(nil),   // 5: smartcore.bos.TestHubNodeRequest
	(*TestHubNodeResponse)(nil),  // 6: smartcore.bos.TestHubNodeResponse
}
var file_hub_proto_depIdxs = []int32{
	0, // 0: smartcore.bos.EnrollHubNodeRequest.node:type_name -> smartcore.bos.HubNode
	0, // 1: smartcore.bos.ListHubNodesResponse.nodes:type_name -> smartcore.bos.HubNode
	1, // 2: smartcore.bos.HubApi.GetHubNode:input_type -> smartcore.bos.GetHubNodeRequest
	3, // 3: smartcore.bos.HubApi.ListHubNodes:input_type -> smartcore.bos.ListHubNodesRequest
	2, // 4: smartcore.bos.HubApi.EnrollHubNode:input_type -> smartcore.bos.EnrollHubNodeRequest
	5, // 5: smartcore.bos.HubApi.TestHubNode:input_type -> smartcore.bos.TestHubNodeRequest
	0, // 6: smartcore.bos.HubApi.GetHubNode:output_type -> smartcore.bos.HubNode
	4, // 7: smartcore.bos.HubApi.ListHubNodes:output_type -> smartcore.bos.ListHubNodesResponse
	0, // 8: smartcore.bos.HubApi.EnrollHubNode:output_type -> smartcore.bos.HubNode
	6, // 9: smartcore.bos.HubApi.TestHubNode:output_type -> smartcore.bos.TestHubNodeResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hub_proto_init() }
func file_hub_proto_init() {
	if File_hub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HubNode); i {
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
		file_hub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHubNodeRequest); i {
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
		file_hub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollHubNodeRequest); i {
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
		file_hub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHubNodesRequest); i {
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
		file_hub_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHubNodesResponse); i {
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
		file_hub_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestHubNodeRequest); i {
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
		file_hub_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestHubNodeResponse); i {
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
			RawDescriptor: file_hub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hub_proto_goTypes,
		DependencyIndexes: file_hub_proto_depIdxs,
		MessageInfos:      file_hub_proto_msgTypes,
	}.Build()
	File_hub_proto = out.File
	file_hub_proto_rawDesc = nil
	file_hub_proto_goTypes = nil
	file_hub_proto_depIdxs = nil
}