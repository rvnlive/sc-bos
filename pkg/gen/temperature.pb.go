// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.1
// source: temperature.proto

package gen

import (
	types "github.com/smart-core-os/sc-api/go/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Temperature represents a target and measured temperature.
type Temperature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read/write, the target temperature.
	SetPoint *types.Temperature `protobuf:"bytes,1,opt,name=set_point,json=setPoint,proto3" json:"set_point,omitempty"`
	// Output only, the measured temperature.
	Measured *types.Temperature `protobuf:"bytes,2,opt,name=measured,proto3" json:"measured,omitempty"`
}

func (x *Temperature) Reset() {
	*x = Temperature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Temperature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Temperature) ProtoMessage() {}

func (x *Temperature) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Temperature.ProtoReflect.Descriptor instead.
func (*Temperature) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{0}
}

func (x *Temperature) GetSetPoint() *types.Temperature {
	if x != nil {
		return x.SetPoint
	}
	return nil
}

func (x *Temperature) GetMeasured() *types.Temperature {
	if x != nil {
		return x.Measured
	}
	return nil
}

type GetTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReadMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
}

func (x *GetTemperatureRequest) Reset() {
	*x = GetTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTemperatureRequest) ProtoMessage() {}

func (x *GetTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTemperatureRequest.ProtoReflect.Descriptor instead.
func (*GetTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{1}
}

func (x *GetTemperatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetTemperatureRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type PullTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReadMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	UpdatesOnly bool                   `protobuf:"varint,3,opt,name=updates_only,json=updatesOnly,proto3" json:"updates_only,omitempty"`
}

func (x *PullTemperatureRequest) Reset() {
	*x = PullTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTemperatureRequest) ProtoMessage() {}

func (x *PullTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTemperatureRequest.ProtoReflect.Descriptor instead.
func (*PullTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{2}
}

func (x *PullTemperatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullTemperatureRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

func (x *PullTemperatureRequest) GetUpdatesOnly() bool {
	if x != nil {
		return x.UpdatesOnly
	}
	return false
}

type PullTemperatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Changes []*PullTemperatureResponse_Change `protobuf:"bytes,1,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *PullTemperatureResponse) Reset() {
	*x = PullTemperatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTemperatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTemperatureResponse) ProtoMessage() {}

func (x *PullTemperatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTemperatureResponse.ProtoReflect.Descriptor instead.
func (*PullTemperatureResponse) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{3}
}

func (x *PullTemperatureResponse) GetChanges() []*PullTemperatureResponse_Change {
	if x != nil {
		return x.Changes
	}
	return nil
}

type UpdateTemperatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Temperature *Temperature           `protobuf:"bytes,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	UpdateMask  *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// When true, temperature is a change to the devices current value.
	Delta bool `protobuf:"varint,4,opt,name=delta,proto3" json:"delta,omitempty"`
}

func (x *UpdateTemperatureRequest) Reset() {
	*x = UpdateTemperatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTemperatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTemperatureRequest) ProtoMessage() {}

func (x *UpdateTemperatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTemperatureRequest.ProtoReflect.Descriptor instead.
func (*UpdateTemperatureRequest) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTemperatureRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTemperatureRequest) GetTemperature() *Temperature {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *UpdateTemperatureRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateTemperatureRequest) GetDelta() bool {
	if x != nil {
		return x.Delta
	}
	return false
}

type PullTemperatureResponse_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ChangeTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=change_time,json=changeTime,proto3" json:"change_time,omitempty"`
	Temperature *Temperature           `protobuf:"bytes,3,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *PullTemperatureResponse_Change) Reset() {
	*x = PullTemperatureResponse_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temperature_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PullTemperatureResponse_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullTemperatureResponse_Change) ProtoMessage() {}

func (x *PullTemperatureResponse_Change) ProtoReflect() protoreflect.Message {
	mi := &file_temperature_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullTemperatureResponse_Change.ProtoReflect.Descriptor instead.
func (*PullTemperatureResponse_Change) Descriptor() ([]byte, []int) {
	return file_temperature_proto_rawDescGZIP(), []int{3, 0}
}

func (x *PullTemperatureResponse_Change) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PullTemperatureResponse_Change) GetChangeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ChangeTime
	}
	return nil
}

func (x *PullTemperatureResponse_Change) GetTemperature() *Temperature {
	if x != nil {
		return x.Temperature
	}
	return nil
}

var File_temperature_proto protoreflect.FileDescriptor

var file_temperature_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62,
	0x6f, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x75, 0x6e, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x73, 0x65, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x22, 0x64, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x88, 0x01, 0x0a, 0x16, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xfc, 0x01,
	0x0a, 0x17, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x1a, 0x97, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x62, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xbf, 0x01, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62,
	0x6f, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x32, 0xa8,
	0x02, 0x0a, 0x0e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x12, 0x54, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x62, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x50, 0x75, 0x6c, 0x6c, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f,
	0x73, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5a, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62,
	0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6e, 0x74, 0x69, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x73, 0x63, 0x2d, 0x62, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temperature_proto_rawDescOnce sync.Once
	file_temperature_proto_rawDescData = file_temperature_proto_rawDesc
)

func file_temperature_proto_rawDescGZIP() []byte {
	file_temperature_proto_rawDescOnce.Do(func() {
		file_temperature_proto_rawDescData = protoimpl.X.CompressGZIP(file_temperature_proto_rawDescData)
	})
	return file_temperature_proto_rawDescData
}

var file_temperature_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_temperature_proto_goTypes = []any{
	(*Temperature)(nil),                    // 0: smartcore.bos.Temperature
	(*GetTemperatureRequest)(nil),          // 1: smartcore.bos.GetTemperatureRequest
	(*PullTemperatureRequest)(nil),         // 2: smartcore.bos.PullTemperatureRequest
	(*PullTemperatureResponse)(nil),        // 3: smartcore.bos.PullTemperatureResponse
	(*UpdateTemperatureRequest)(nil),       // 4: smartcore.bos.UpdateTemperatureRequest
	(*PullTemperatureResponse_Change)(nil), // 5: smartcore.bos.PullTemperatureResponse.Change
	(*types.Temperature)(nil),              // 6: smartcore.types.Temperature
	(*fieldmaskpb.FieldMask)(nil),          // 7: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
}
var file_temperature_proto_depIdxs = []int32{
	6,  // 0: smartcore.bos.Temperature.set_point:type_name -> smartcore.types.Temperature
	6,  // 1: smartcore.bos.Temperature.measured:type_name -> smartcore.types.Temperature
	7,  // 2: smartcore.bos.GetTemperatureRequest.read_mask:type_name -> google.protobuf.FieldMask
	7,  // 3: smartcore.bos.PullTemperatureRequest.read_mask:type_name -> google.protobuf.FieldMask
	5,  // 4: smartcore.bos.PullTemperatureResponse.changes:type_name -> smartcore.bos.PullTemperatureResponse.Change
	0,  // 5: smartcore.bos.UpdateTemperatureRequest.temperature:type_name -> smartcore.bos.Temperature
	7,  // 6: smartcore.bos.UpdateTemperatureRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 7: smartcore.bos.PullTemperatureResponse.Change.change_time:type_name -> google.protobuf.Timestamp
	0,  // 8: smartcore.bos.PullTemperatureResponse.Change.temperature:type_name -> smartcore.bos.Temperature
	1,  // 9: smartcore.bos.TemperatureApi.GetTemperature:input_type -> smartcore.bos.GetTemperatureRequest
	2,  // 10: smartcore.bos.TemperatureApi.PullTemperature:input_type -> smartcore.bos.PullTemperatureRequest
	4,  // 11: smartcore.bos.TemperatureApi.UpdateTemperature:input_type -> smartcore.bos.UpdateTemperatureRequest
	0,  // 12: smartcore.bos.TemperatureApi.GetTemperature:output_type -> smartcore.bos.Temperature
	3,  // 13: smartcore.bos.TemperatureApi.PullTemperature:output_type -> smartcore.bos.PullTemperatureResponse
	0,  // 14: smartcore.bos.TemperatureApi.UpdateTemperature:output_type -> smartcore.bos.Temperature
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_temperature_proto_init() }
func file_temperature_proto_init() {
	if File_temperature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temperature_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Temperature); i {
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
		file_temperature_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetTemperatureRequest); i {
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
		file_temperature_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PullTemperatureRequest); i {
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
		file_temperature_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PullTemperatureResponse); i {
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
		file_temperature_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTemperatureRequest); i {
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
		file_temperature_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PullTemperatureResponse_Change); i {
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
			RawDescriptor: file_temperature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_temperature_proto_goTypes,
		DependencyIndexes: file_temperature_proto_depIdxs,
		MessageInfos:      file_temperature_proto_msgTypes,
	}.Build()
	File_temperature_proto = out.File
	file_temperature_proto_rawDesc = nil
	file_temperature_proto_goTypes = nil
	file_temperature_proto_depIdxs = nil
}
