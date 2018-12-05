// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	math "math"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type E int32

const (
	E_ZERO E = 0
)

type xxx_E E

func (e E) ProtoReflect() protoreflect.Enum {
	return (xxx_E)(e)
}
func (e xxx_E) Type() protoreflect.EnumType {
	return xxx_A_ProtoFile_EnumTypes[0]
}
func (e xxx_E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var E_name = map[int32]string{
	0: "ZERO",
}

var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	p := new(E)
	*p = x
	return p
}

func (x E) String() string {
	return proto.EnumName(E_name, int32(x))
}

func (x *E) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(E_value, data, "E")
	if err != nil {
		return err
	}
	*x = E(value)
	return nil
}

func (E) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

type xxx_M_Subenum M_Subenum

func (e M_Subenum) ProtoReflect() protoreflect.Enum {
	return (xxx_M_Subenum)(e)
}
func (e xxx_M_Subenum) Type() protoreflect.EnumType {
	return xxx_A_ProtoFile_EnumTypes[1]
}
func (e xxx_M_Subenum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	p := new(M_Subenum)
	*p = x
	return p
}

func (x M_Subenum) String() string {
	return proto.EnumName(M_Subenum_name, int32(x))
}

func (x *M_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Subenum_value, data, "M_Subenum")
	if err != nil {
		return err
	}
	*x = M_Subenum(value)
	return nil
}

func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

type xxx_M_Submessage_Submessage_Subenum M_Submessage_Submessage_Subenum

func (e M_Submessage_Submessage_Subenum) ProtoReflect() protoreflect.Enum {
	return (xxx_M_Submessage_Submessage_Subenum)(e)
}
func (e xxx_M_Submessage_Submessage_Subenum) Type() protoreflect.EnumType {
	return xxx_A_ProtoFile_EnumTypes[2]
}
func (e xxx_M_Submessage_Submessage_Subenum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	p := new(M_Submessage_Submessage_Subenum)
	*p = x
	return p
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return proto.EnumName(M_Submessage_Submessage_Subenum_name, int32(x))
}

func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(M_Submessage_Submessage_Subenum_value, data, "M_Submessage_Submessage_Subenum")
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(value)
	return nil
}

func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2      `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	S  *string  `protobuf:"bytes,4,opt,name=s,def=default" json:"s,omitempty"`
	B  []byte   `protobuf:"bytes,5,opt,name=b,def=default" json:"b,omitempty"`
	F  *float64 `protobuf:"fixed64,6,opt,name=f,def=nan" json:"f,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField                   isM_OneofField `protobuf_oneof:"oneof_field"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

type xxx_M struct{ m *M }

func (m *M) ProtoReflect() protoreflect.Message {
	return xxx_M{m}
}
func (m xxx_M) Type() protoreflect.MessageType {
	return xxx_A_ProtoFile_MessageTypes[0].Type
}
func (m xxx_M) KnownFields() protoreflect.KnownFields {
	return xxx_A_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_M) UnknownFields() protoreflect.UnknownFields {
	return xxx_A_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_M) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M) ProtoMutable() {}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0}
}

var extRange_M = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*M) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_M
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

const Default_M_S string = "default"

var Default_M_B []byte = []byte("default")
var Default_M_F float64 = math.NaN()

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M) GetS() string {
	if m != nil && m.S != nil {
		return *m.S
	}
	return Default_M_S
}

func (m *M) GetB() []byte {
	if m != nil && m.B != nil {
		return m.B
	}
	return append([]byte(nil), Default_M_B...)
}

func (m *M) GetF() float64 {
	if m != nil && m.F != nil {
		return *m.F
	}
	return Default_M_F
}

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof"`
}

type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof"`
}

func (*M_OneofInt32) isM_OneofField() {}

func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

type xxx_M_Submessage struct{ m *M_Submessage }

func (m *M_Submessage) ProtoReflect() protoreflect.Message {
	return xxx_M_Submessage{m}
}
func (m xxx_M_Submessage) Type() protoreflect.MessageType {
	return xxx_A_ProtoFile_MessageTypes[1].Type
}
func (m xxx_M_Submessage) KnownFields() protoreflect.KnownFields {
	return xxx_A_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_M_Submessage) UnknownFields() protoreflect.UnknownFields {
	return xxx_A_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_M_Submessage) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M_Submessage) ProtoMutable() {}

func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_382f7805394b5c4e, []int{0, 0}
}

func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof"`
}

type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}

func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

var E_ExtensionField = &proto.ExtensionDesc{
	ExtendedType:  (*M)(nil),
	ExtensionType: (*string)(nil),
	Field:         100,
	Name:          "goproto.protoc.import_public.sub.extension_field",
	Tag:           "bytes,100,opt,name=extension_field",
	Filename:      "import_public/sub/a.proto",
}

func init() {
	proto.RegisterFile("import_public/sub/a.proto", fileDescriptor_382f7805394b5c4e)
	proto.RegisterEnum("goproto.protoc.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.protoc.import_public.sub.M")
	proto.RegisterType((*M_Submessage)(nil), "goproto.protoc.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
}

var fileDescriptor_382f7805394b5c4e = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x7b, 0xe2, 0x5e, 0xc2, 0x29, 0x85, 0x76, 0x44, 0xd0, 0xd0, 0x95, 0x09, 0x2c, 0xac,
	0xa2, 0x7a, 0x24, 0x63, 0x79, 0x91, 0x1d, 0x91, 0xc2, 0x4d, 0xb5, 0x2a, 0xd9, 0x62, 0xd3, 0x8d,
	0xe5, 0xb1, 0xc7, 0xc6, 0x52, 0x3c, 0x13, 0x75, 0x66, 0x10, 0xcb, 0x3c, 0x15, 0x2f, 0xc0, 0x8b,
	0xa1, 0xd8, 0x6d, 0x93, 0x88, 0x20, 0xba, 0xf3, 0xf9, 0xfd, 0x7d, 0xff, 0xb1, 0x3c, 0x83, 0xaf,
	0x9a, 0x76, 0xa1, 0x6e, 0x4d, 0xb6, 0xb0, 0x7c, 0xde, 0x14, 0x4c, 0x5b, 0xce, 0x72, 0x7f, 0x71,
	0xab, 0x8c, 0x22, 0x6e, 0xad, 0xba, 0x87, 0x7e, 0x2c, 0xfc, 0x2d, 0xd2, 0xd7, 0x96, 0x9f, 0xef,
	0x90, 0x79, 0x4f, 0x8f, 0x7f, 0x39, 0x08, 0x31, 0x09, 0x71, 0xd0, 0x06, 0x14, 0x5c, 0xf0, 0x8e,
	0x83, 0xb7, 0xfe, 0xff, 0xfa, 0xfc, 0x38, 0x48, 0x06, 0x6d, 0x40, 0x46, 0x08, 0x9a, 0xee, 0xbb,
	0xe0, 0x3d, 0x99, 0x1c, 0x95, 0xa2, 0xca, 0xed, 0xdc, 0x24, 0xa0, 0x57, 0x31, 0xa7, 0x07, 0x2e,
	0x78, 0x4f, 0x37, 0x62, 0x4e, 0xce, 0x10, 0x2a, 0x7a, 0xe8, 0x82, 0x07, 0x13, 0x47, 0xe6, 0x32,
	0x81, 0x8a, 0xbc, 0xc6, 0x63, 0x25, 0x85, 0xaa, 0xb2, 0x46, 0x9a, 0xf7, 0x01, 0x1d, 0xb8, 0xe0,
	0x1d, 0x7c, 0xde, 0x4b, 0xb0, 0x0b, 0xbf, 0xac, 0xb2, 0x2d, 0x24, 0x0a, 0xa9, 0xe3, 0x82, 0xe7,
	0x6c, 0x22, 0x51, 0x78, 0xfe, 0x1b, 0x10, 0x53, 0xcb, 0x5b, 0xa1, 0x75, 0x5e, 0x0b, 0x12, 0xe1,
	0x4b, 0xfd, 0x30, 0x65, 0x9b, 0xfd, 0x70, 0xd7, 0xff, 0x62, 0xfd, 0xfe, 0x7a, 0xbd, 0xe9, 0x1f,
	0x5e, 0x14, 0x76, 0xdf, 0xe5, 0xec, 0xf6, 0xa2, 0x70, 0xfc, 0x0e, 0xc9, 0x7a, 0x7b, 0x96, 0x5a,
	0x2e, 0xa4, 0x6d, 0xc9, 0x08, 0xcf, 0xe2, 0x2c, 0xfd, 0x36, 0x8d, 0x67, 0x69, 0xfa, 0xe1, 0xd3,
	0x2c, 0xbb, 0x99, 0x25, 0xd7, 0xa7, 0x7b, 0x53, 0xba, 0x63, 0x49, 0xd5, 0x88, 0x79, 0x39, 0x1e,
	0xe1, 0xd1, 0xbd, 0x8b, 0x78, 0x18, 0xdf, 0x09, 0x17, 0xc3, 0x61, 0x79, 0xba, 0x5c, 0x2e, 0x97,
	0x83, 0xe9, 0xc9, 0xfd, 0x9f, 0xe8, 0xf8, 0x8b, 0x13, 0x84, 0x19, 0x19, 0xe2, 0x7e, 0xcf, 0x4d,
	0xae, 0xf0, 0xb9, 0xf8, 0x69, 0x84, 0xd4, 0x8d, 0x92, 0x3d, 0x41, 0xde, 0x3c, 0xe2, 0x20, 0x69,
	0xb9, 0x3a, 0xbe, 0xe4, 0xd9, 0x83, 0xfb, 0x71, 0xa5, 0x4e, 0xaf, 0x6e, 0xbe, 0xd6, 0x8d, 0xf9,
	0x6e, 0xb9, 0x5f, 0xa8, 0x96, 0xd5, 0x6a, 0x9e, 0xcb, 0x9a, 0x75, 0x2d, 0xdc, 0x56, 0xec, 0x47,
	0xc0, 0x8a, 0xb6, 0xec, 0xe7, 0xe2, 0xb2, 0x16, 0xf2, 0xb2, 0x56, 0xcc, 0x08, 0x6d, 0xca, 0xdc,
	0xe4, 0xec, 0xaf, 0xbb, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xa9, 0x17, 0xb5, 0xbc, 0x02,
	0x00, 0x00,
}

func init() {
	xxx_A_ProtoFile_FileDesc.Enums = xxx_A_ProtoFile_EnumDescs[0:1]
	xxx_A_ProtoFile_FileDesc.Messages = xxx_A_ProtoFile_MessageDescs[0:1]
	xxx_A_ProtoFile_MessageDescs[0].Enums = xxx_A_ProtoFile_EnumDescs[1:2]
	xxx_A_ProtoFile_MessageDescs[0].Messages = xxx_A_ProtoFile_MessageDescs[1:2]
	xxx_A_ProtoFile_MessageDescs[1].Enums = xxx_A_ProtoFile_EnumDescs[2:3]
	xxx_A_ProtoFile_MessageDescs[0].Fields[0].MessageType = protoimpl.X.MessageTypeOf((*M2)(nil))
	var err error
	A_ProtoFile, err = prototype.NewFile(&xxx_A_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var A_ProtoFile protoreflect.FileDescriptor

var xxx_A_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "import_public/sub/a.proto",
	Package: "goproto.protoc.import_public.sub",
	Imports: []protoreflect.FileImport{
		{FileDescriptor: prototype.PlaceholderFile("import_public/sub/b.proto", "goproto.protoc.import_public.sub")},
	},
}
var xxx_A_ProtoFile_EnumTypes = [3]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_A_ProtoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return E(n)
		},
	),
	prototype.GoEnum(
		xxx_A_ProtoFile_EnumDescs[1].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return M_Subenum(n)
		},
	),
	prototype.GoEnum(
		xxx_A_ProtoFile_EnumDescs[2].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return M_Submessage_Submessage_Subenum(n)
		},
	),
}
var xxx_A_ProtoFile_EnumDescs = [3]prototype.Enum{
	{
		Name: "E",
		Values: []prototype.EnumValue{
			{Name: "ZERO", Number: 0},
		},
	},
	{
		Name: "Subenum",
		Values: []prototype.EnumValue{
			{Name: "M_ZERO", Number: 0},
		},
	},
	{
		Name: "Submessage_Subenum",
		Values: []prototype.EnumValue{
			{Name: "M_SUBMESSAGE_ZERO", Number: 0},
		},
	},
}
var xxx_A_ProtoFile_MessageTypes = [2]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_A_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_A_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M_Submessage)
		},
	)},
}
var xxx_A_ProtoFile_MessageDescs = [2]prototype.Message{
	{
		Name: "M",
		Fields: []prototype.Field{
			{
				Name:        "m2",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "m2",
			},
			{
				Name:        "s",
				Number:      4,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "s",
				Default:     protoreflect.ValueOf(string("default")),
			},
			{
				Name:        "b",
				Number:      5,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.BytesKind,
				JSONName:    "b",
				Default:     protoreflect.ValueOf(("default")),
			},
			{
				Name:        "f",
				Number:      6,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.DoubleKind,
				JSONName:    "f",
				Default:     protoreflect.ValueOf(float64(math.NaN())),
			},
			{
				Name:        "oneof_int32",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int32Kind,
				JSONName:    "oneofInt32",
				OneofName:   "oneof_field",
			},
			{
				Name:        "oneof_int64",
				Number:      3,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int64Kind,
				JSONName:    "oneofInt64",
				OneofName:   "oneof_field",
			},
		},
		Oneofs: []prototype.Oneof{
			{Name: "oneof_field"},
		},
		ExtensionRanges: [][2]protoreflect.FieldNumber{{100, 536870912}},
	},
	{
		Name: "Submessage",
		Fields: []prototype.Field{
			{
				Name:        "submessage_oneof_int32",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int32Kind,
				JSONName:    "submessageOneofInt32",
				OneofName:   "submessage_oneof_field",
			},
			{
				Name:        "submessage_oneof_int64",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.Int64Kind,
				JSONName:    "submessageOneofInt64",
				OneofName:   "submessage_oneof_field",
			},
		},
		Oneofs: []prototype.Oneof{
			{Name: "submessage_oneof_field"},
		},
	},
}
