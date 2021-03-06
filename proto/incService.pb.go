// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: proto/incService.proto

package proto

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

type IncMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncCount int32 `protobuf:"varint,1,opt,name=incCount,proto3" json:"incCount,omitempty"`
}

func (x *IncMessage) Reset() {
	*x = IncMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncMessage) ProtoMessage() {}

func (x *IncMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncMessage.ProtoReflect.Descriptor instead.
func (*IncMessage) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{0}
}

func (x *IncMessage) GetIncCount() int32 {
	if x != nil {
		return x.IncCount
	}
	return 0
}

type IncResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentCount int32 `protobuf:"varint,1,opt,name=currentCount,proto3" json:"currentCount,omitempty"`
}

func (x *IncResponse) Reset() {
	*x = IncResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncResponse) ProtoMessage() {}

func (x *IncResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncResponse.ProtoReflect.Descriptor instead.
func (*IncResponse) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{1}
}

func (x *IncResponse) GetCurrentCount() int32 {
	if x != nil {
		return x.CurrentCount
	}
	return 0
}

type ReplicateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body      string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	TimeStamp uint32 `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	IncSum    int32  `protobuf:"varint,3,opt,name=incSum,proto3" json:"incSum,omitempty"`
}

func (x *ReplicateMessage) Reset() {
	*x = ReplicateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateMessage) ProtoMessage() {}

func (x *ReplicateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateMessage.ProtoReflect.Descriptor instead.
func (*ReplicateMessage) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{2}
}

func (x *ReplicateMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ReplicateMessage) GetTimeStamp() uint32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *ReplicateMessage) GetIncSum() int32 {
	if x != nil {
		return x.IncSum
	}
	return 0
}

type ReplicateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ReplicateResponse) Reset() {
	*x = ReplicateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateResponse) ProtoMessage() {}

func (x *ReplicateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateResponse.ProtoReflect.Descriptor instead.
func (*ReplicateResponse) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{3}
}

func (x *ReplicateResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{4}
}

type CutReplica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *CutReplica) Reset() {
	*x = CutReplica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CutReplica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CutReplica) ProtoMessage() {}

func (x *CutReplica) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CutReplica.ProtoReflect.Descriptor instead.
func (*CutReplica) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{5}
}

func (x *CutReplica) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PortsAndClocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfPorts  []int32  `protobuf:"varint,1,rep,packed,name=listOfPorts,proto3" json:"listOfPorts,omitempty"`
	ListOfClocks []uint32 `protobuf:"varint,2,rep,packed,name=listOfClocks,proto3" json:"listOfClocks,omitempty"`
}

func (x *PortsAndClocks) Reset() {
	*x = PortsAndClocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortsAndClocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortsAndClocks) ProtoMessage() {}

func (x *PortsAndClocks) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortsAndClocks.ProtoReflect.Descriptor instead.
func (*PortsAndClocks) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{6}
}

func (x *PortsAndClocks) GetListOfPorts() []int32 {
	if x != nil {
		return x.ListOfPorts
	}
	return nil
}

func (x *PortsAndClocks) GetListOfClocks() []uint32 {
	if x != nil {
		return x.ListOfClocks
	}
	return nil
}

type ElectionPorts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListOfPorts []int32 `protobuf:"varint,1,rep,packed,name=listOfPorts,proto3" json:"listOfPorts,omitempty"`
}

func (x *ElectionPorts) Reset() {
	*x = ElectionPorts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_incService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionPorts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionPorts) ProtoMessage() {}

func (x *ElectionPorts) ProtoReflect() protoreflect.Message {
	mi := &file_proto_incService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionPorts.ProtoReflect.Descriptor instead.
func (*ElectionPorts) Descriptor() ([]byte, []int) {
	return file_proto_incService_proto_rawDescGZIP(), []int{7}
}

func (x *ElectionPorts) GetListOfPorts() []int32 {
	if x != nil {
		return x.ListOfPorts
	}
	return nil
}

var File_proto_incService_proto protoreflect.FileDescriptor

var file_proto_incService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0a, 0x49, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x69, 0x6e, 0x63, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x0b, 0x49, 0x6e, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x10,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x53, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x53, 0x75, 0x6d, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0a, 0x63,
	0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x56, 0x0a,
	0x0e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x43,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x31, 0x0a, 0x0d, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x4f, 0x66,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x32, 0x9b, 0x03, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x09, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x6c, 0x73,
	0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x43, 0x75, 0x74,
	0x4f, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x1a, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x52,
	0x69, 0x6e, 0x67, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x73, 0x4e, 0x65, 0x77,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x12, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x55, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x63, 0x61, 0x72, 0x6c, 0x39, 0x2f, 0x4f, 0x6e, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x53,
	0x6b, 0x72, 0x69, 0x76, 0x65, 0x62, 0x6f, 0x72, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x79, 0x73, 0x4d,
	0x4f, 0x43, 0x4b, 0x2f, 0x44, 0x49, 0x53, 0x59, 0x53, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_incService_proto_rawDescOnce sync.Once
	file_proto_incService_proto_rawDescData = file_proto_incService_proto_rawDesc
)

func file_proto_incService_proto_rawDescGZIP() []byte {
	file_proto_incService_proto_rawDescOnce.Do(func() {
		file_proto_incService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_incService_proto_rawDescData)
	})
	return file_proto_incService_proto_rawDescData
}

var file_proto_incService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_incService_proto_goTypes = []interface{}{
	(*IncMessage)(nil),        // 0: proto.IncMessage
	(*IncResponse)(nil),       // 1: proto.IncResponse
	(*ReplicateMessage)(nil),  // 2: proto.ReplicateMessage
	(*ReplicateResponse)(nil), // 3: proto.ReplicateResponse
	(*Void)(nil),              // 4: proto.Void
	(*CutReplica)(nil),        // 5: proto.cutReplica
	(*PortsAndClocks)(nil),    // 6: proto.PortsAndClocks
	(*ElectionPorts)(nil),     // 7: proto.ElectionPorts
}
var file_proto_incService_proto_depIdxs = []int32{
	0, // 0: proto.ChatService.Increment:input_type -> proto.IncMessage
	2, // 1: proto.ChatService.Replicate:input_type -> proto.ReplicateMessage
	2, // 2: proto.ChatService.RegisterPulse:input_type -> proto.ReplicateMessage
	5, // 3: proto.ChatService.CutOfReplica:input_type -> proto.cutReplica
	6, // 4: proto.ChatService.RingElection:input_type -> proto.PortsAndClocks
	4, // 5: proto.ChatService.SelectAsNewLeader:input_type -> proto.Void
	7, // 6: proto.ChatService.BroadcastNewLeader:input_type -> proto.ElectionPorts
	1, // 7: proto.ChatService.Increment:output_type -> proto.IncResponse
	3, // 8: proto.ChatService.Replicate:output_type -> proto.ReplicateResponse
	4, // 9: proto.ChatService.RegisterPulse:output_type -> proto.Void
	4, // 10: proto.ChatService.CutOfReplica:output_type -> proto.Void
	4, // 11: proto.ChatService.RingElection:output_type -> proto.Void
	7, // 12: proto.ChatService.SelectAsNewLeader:output_type -> proto.ElectionPorts
	4, // 13: proto.ChatService.BroadcastNewLeader:output_type -> proto.Void
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_incService_proto_init() }
func file_proto_incService_proto_init() {
	if File_proto_incService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_incService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncMessage); i {
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
		file_proto_incService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncResponse); i {
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
		file_proto_incService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicateMessage); i {
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
		file_proto_incService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicateResponse); i {
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
		file_proto_incService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_proto_incService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CutReplica); i {
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
		file_proto_incService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortsAndClocks); i {
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
		file_proto_incService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionPorts); i {
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
			RawDescriptor: file_proto_incService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_incService_proto_goTypes,
		DependencyIndexes: file_proto_incService_proto_depIdxs,
		MessageInfos:      file_proto_incService_proto_msgTypes,
	}.Build()
	File_proto_incService_proto = out.File
	file_proto_incService_proto_rawDesc = nil
	file_proto_incService_proto_goTypes = nil
	file_proto_incService_proto_depIdxs = nil
}
