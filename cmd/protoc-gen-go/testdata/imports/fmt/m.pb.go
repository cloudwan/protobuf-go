// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/fmt/m.proto

package fmt

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type M struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_M struct{ m *M }

func (m *M) ProtoReflect() protoreflect.Message {
	return xxx_M{m}
}
func (m xxx_M) Type() protoreflect.MessageType {
	return xxx_M_ProtoFile_MessageTypes[0].Type
}
func (m xxx_M) KnownFields() protoreflect.KnownFields {
	return xxx_M_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_M) UnknownFields() protoreflect.UnknownFields {
	return xxx_M_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_M) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M) ProtoMutable() {}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c126fcd452e392, []int{0}
}

func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

func init() {
	proto.RegisterFile("imports/fmt/m.proto", fileDescriptor_72c126fcd452e392)
	proto.RegisterType((*M)(nil), "fmt.M")
}

var fileDescriptor_72c126fcd452e392 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x4f, 0xcb, 0x2d, 0xd1, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4e, 0xcb, 0x2d, 0x51, 0x62, 0xe6, 0x62, 0xf4, 0x75, 0xb2, 0x8f, 0xb2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x07,
	0x2b, 0x4a, 0x2a, 0x4d, 0x83, 0x30, 0x92, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd3, 0xf3, 0xf5, 0x4b,
	0x52, 0x8b, 0x4b, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0x91, 0x8c, 0x4c, 0x62, 0x03, 0xab, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xc9, 0xee, 0xbe, 0x68, 0x00, 0x00, 0x00,
}

func init() {
	xxx_M_ProtoFile_FileDesc.Messages = xxx_M_ProtoFile_MessageDescs[0:1]
	var err error
	M_ProtoFile, err = prototype.NewFile(&xxx_M_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var M_ProtoFile protoreflect.FileDescriptor

var xxx_M_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto3,
	Path:    "imports/fmt/m.proto",
	Package: "fmt",
}
var xxx_M_ProtoFile_MessageTypes = [1]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_M_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M)
		},
	)},
}
var xxx_M_ProtoFile_MessageDescs = [1]prototype.Message{
	{
		Name: "M",
	},
}
