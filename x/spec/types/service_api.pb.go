// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spec/service_api.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PARSER_FUNC int32

const (
	PARSER_FUNC_EMPTY                 PARSER_FUNC = 0
	PARSER_FUNC_PARSE_PARAM_BY_ARG    PARSER_FUNC = 1
	PARSER_FUNC_PARSE_PARAM_CANONICAL PARSER_FUNC = 2
)

var PARSER_FUNC_name = map[int32]string{
	0: "EMPTY",
	1: "PARSE_PARAM_BY_ARG",
	2: "PARSE_PARAM_CANONICAL",
}

var PARSER_FUNC_value = map[string]int32{
	"EMPTY":                 0,
	"PARSE_PARAM_BY_ARG":    1,
	"PARSE_PARAM_CANONICAL": 2,
}

func (x PARSER_FUNC) String() string {
	return proto.EnumName(PARSER_FUNC_name, int32(x))
}

func (PARSER_FUNC) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{0}
}

type ServiceApi struct {
	Name          string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BlockParsing  BlockParser    `protobuf:"bytes,2,opt,name=blockParsing,proto3" json:"blockParsing"`
	ComputeUnits  uint64         `protobuf:"varint,3,opt,name=computeUnits,proto3" json:"computeUnits,omitempty"`
	Enabled       bool           `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ApiInterfaces []ApiInterface `protobuf:"bytes,5,rep,name=apiInterfaces,proto3" json:"apiInterfaces"`
}

func (m *ServiceApi) Reset()         { *m = ServiceApi{} }
func (m *ServiceApi) String() string { return proto.CompactTextString(m) }
func (*ServiceApi) ProtoMessage()    {}
func (*ServiceApi) Descriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{0}
}
func (m *ServiceApi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServiceApi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServiceApi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServiceApi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceApi.Merge(m, src)
}
func (m *ServiceApi) XXX_Size() int {
	return m.Size()
}
func (m *ServiceApi) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceApi.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceApi proto.InternalMessageInfo

func (m *ServiceApi) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceApi) GetBlockParsing() BlockParser {
	if m != nil {
		return m.BlockParsing
	}
	return BlockParser{}
}

func (m *ServiceApi) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *ServiceApi) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ServiceApi) GetApiInterfaces() []ApiInterface {
	if m != nil {
		return m.ApiInterfaces
	}
	return nil
}

type ApiInterface struct {
	Interface         string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Type              string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ExtraComputeUnits uint64 `protobuf:"varint,3,opt,name=extraComputeUnits,proto3" json:"extraComputeUnits,omitempty"`
}

func (m *ApiInterface) Reset()         { *m = ApiInterface{} }
func (m *ApiInterface) String() string { return proto.CompactTextString(m) }
func (*ApiInterface) ProtoMessage()    {}
func (*ApiInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{1}
}
func (m *ApiInterface) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApiInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApiInterface.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApiInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiInterface.Merge(m, src)
}
func (m *ApiInterface) XXX_Size() int {
	return m.Size()
}
func (m *ApiInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiInterface.DiscardUnknown(m)
}

var xxx_messageInfo_ApiInterface proto.InternalMessageInfo

func (m *ApiInterface) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ApiInterface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ApiInterface) GetExtraComputeUnits() uint64 {
	if m != nil {
		return m.ExtraComputeUnits
	}
	return 0
}

type BlockParser struct {
	ParserArg  []string    `protobuf:"bytes,1,rep,name=parserArg,proto3" json:"parserArg,omitempty"`
	ParserFunc PARSER_FUNC `protobuf:"varint,2,opt,name=parserFunc,proto3,enum=lavanet.lava.spec.PARSER_FUNC" json:"parserFunc,omitempty"`
}

func (m *BlockParser) Reset()         { *m = BlockParser{} }
func (m *BlockParser) String() string { return proto.CompactTextString(m) }
func (*BlockParser) ProtoMessage()    {}
func (*BlockParser) Descriptor() ([]byte, []int) {
	return fileDescriptor_3323a3ad252c5ed4, []int{2}
}
func (m *BlockParser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockParser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockParser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockParser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockParser.Merge(m, src)
}
func (m *BlockParser) XXX_Size() int {
	return m.Size()
}
func (m *BlockParser) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockParser.DiscardUnknown(m)
}

var xxx_messageInfo_BlockParser proto.InternalMessageInfo

func (m *BlockParser) GetParserArg() []string {
	if m != nil {
		return m.ParserArg
	}
	return nil
}

func (m *BlockParser) GetParserFunc() PARSER_FUNC {
	if m != nil {
		return m.ParserFunc
	}
	return PARSER_FUNC_EMPTY
}

func init() {
	proto.RegisterEnum("lavanet.lava.spec.PARSER_FUNC", PARSER_FUNC_name, PARSER_FUNC_value)
	proto.RegisterType((*ServiceApi)(nil), "lavanet.lava.spec.ServiceApi")
	proto.RegisterType((*ApiInterface)(nil), "lavanet.lava.spec.ApiInterface")
	proto.RegisterType((*BlockParser)(nil), "lavanet.lava.spec.BlockParser")
}

func init() { proto.RegisterFile("spec/service_api.proto", fileDescriptor_3323a3ad252c5ed4) }

var fileDescriptor_3323a3ad252c5ed4 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xf5, 0x36, 0x29, 0x90, 0x49, 0x40, 0xe9, 0x0a, 0x2a, 0x83, 0xd0, 0xd6, 0x8a, 0x38, 0x58,
	0x08, 0xd9, 0x52, 0xb9, 0x23, 0xad, 0xa3, 0x16, 0xaa, 0xd2, 0x10, 0x6d, 0xe9, 0xa1, 0x5c, 0xac,
	0xb5, 0x59, 0xcc, 0xaa, 0xc9, 0x7a, 0x65, 0x6f, 0xaa, 0xf2, 0x17, 0x7c, 0x06, 0x9f, 0xd2, 0x63,
	0x8f, 0x9c, 0x10, 0x72, 0xbe, 0x81, 0x3b, 0xf2, 0x26, 0x01, 0x47, 0xc9, 0x69, 0x67, 0xe6, 0xbd,
	0xa7, 0x37, 0xf3, 0xb4, 0xb0, 0x5f, 0x6a, 0x91, 0x86, 0xa5, 0x28, 0xae, 0x65, 0x2a, 0x62, 0xae,
	0x65, 0xa0, 0x8b, 0xdc, 0xe4, 0x78, 0x6f, 0xc2, 0xaf, 0xb9, 0x12, 0x26, 0xa8, 0xdf, 0xa0, 0x26,
	0x3d, 0x7b, 0x9c, 0xe5, 0x59, 0x6e, 0xd1, 0xb0, 0xae, 0x16, 0xc4, 0xc1, 0x1f, 0x04, 0x70, 0xbe,
	0x90, 0x53, 0x2d, 0x31, 0x86, 0xb6, 0xe2, 0x53, 0xe1, 0x22, 0x0f, 0xf9, 0x1d, 0x66, 0x6b, 0xfc,
	0x0e, 0x7a, 0xc9, 0x24, 0x4f, 0xaf, 0xc6, 0xbc, 0x28, 0xa5, 0xca, 0xdc, 0x1d, 0x0f, 0xf9, 0xdd,
	0x43, 0x12, 0x6c, 0x58, 0x04, 0xd1, 0x8a, 0x26, 0x8a, 0xa8, 0x7d, 0xfb, 0xeb, 0xc0, 0x61, 0x6b,
	0x4a, 0x3c, 0x80, 0x5e, 0x9a, 0x4f, 0xf5, 0xcc, 0x88, 0x0b, 0x25, 0x4d, 0xe9, 0xb6, 0x3c, 0xe4,
	0xb7, 0xd9, 0xda, 0x0c, 0xbb, 0x70, 0x5f, 0x28, 0x9e, 0x4c, 0xc4, 0x67, 0xb7, 0xed, 0x21, 0xff,
	0x01, 0x5b, 0xb5, 0xf8, 0x14, 0x1e, 0x72, 0x2d, 0x4f, 0x94, 0x11, 0xc5, 0x17, 0x9e, 0x8a, 0xd2,
	0xdd, 0xf5, 0x5a, 0x7e, 0xf7, 0xf0, 0x60, 0xcb, 0x22, 0xb4, 0xc1, 0x5b, 0x6e, 0xb2, 0xae, 0x1d,
	0x28, 0xe8, 0x35, 0x49, 0xf8, 0x39, 0x74, 0xe4, 0xaa, 0x59, 0x5e, 0xff, 0x7f, 0x50, 0xc7, 0x62,
	0xbe, 0x69, 0x61, 0x4f, 0xef, 0x30, 0x5b, 0xe3, 0x57, 0xb0, 0x27, 0x6e, 0x4c, 0xc1, 0x87, 0x9b,
	0x17, 0x6d, 0x02, 0x83, 0x2b, 0xe8, 0x36, 0xd2, 0xa9, 0xed, 0xb4, 0xad, 0x68, 0x91, 0xb9, 0xc8,
	0x6b, 0xd5, 0x76, 0xff, 0x06, 0xf8, 0x0d, 0xc0, 0xa2, 0x39, 0x9e, 0xa9, 0xd4, 0x9a, 0x3e, 0xda,
	0x9a, 0xf7, 0x98, 0xb2, 0xf3, 0x23, 0x16, 0x1f, 0x5f, 0x8c, 0x86, 0xac, 0xa1, 0x78, 0x79, 0x0a,
	0xdd, 0x06, 0x84, 0x3b, 0xb0, 0x7b, 0x74, 0x36, 0xfe, 0x78, 0xd9, 0x77, 0xf0, 0x3e, 0x60, 0x8b,
	0xc4, 0x63, 0xca, 0xe8, 0x59, 0x1c, 0x5d, 0xc6, 0x94, 0xbd, 0xed, 0x23, 0xfc, 0x14, 0x9e, 0x34,
	0xe7, 0x43, 0x3a, 0xfa, 0x30, 0x3a, 0x19, 0xd2, 0xf7, 0xfd, 0x9d, 0x28, 0xfa, 0x51, 0x11, 0x74,
	0x5b, 0x11, 0x74, 0x57, 0x11, 0xf4, 0xbb, 0x22, 0xe8, 0xfb, 0x9c, 0x38, 0x77, 0x73, 0xe2, 0xfc,
	0x9c, 0x13, 0xe7, 0xd3, 0x8b, 0x4c, 0x9a, 0xaf, 0xb3, 0x24, 0x48, 0xf3, 0x69, 0xb8, 0x5c, 0xd0,
	0xbe, 0xe1, 0x4d, 0x68, 0xbf, 0x66, 0x1d, 0x55, 0x99, 0xdc, 0xb3, 0x9f, 0xed, 0xf5, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x21, 0x8e, 0x5c, 0x86, 0xaf, 0x02, 0x00, 0x00,
}

func (this *ServiceApi) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceApi)
	if !ok {
		that2, ok := that.(ServiceApi)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.BlockParsing.Equal(&that1.BlockParsing) {
		return false
	}
	if this.ComputeUnits != that1.ComputeUnits {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if len(this.ApiInterfaces) != len(that1.ApiInterfaces) {
		return false
	}
	for i := range this.ApiInterfaces {
		if !this.ApiInterfaces[i].Equal(&that1.ApiInterfaces[i]) {
			return false
		}
	}
	return true
}
func (this *ApiInterface) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ApiInterface)
	if !ok {
		that2, ok := that.(ApiInterface)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Interface != that1.Interface {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.ExtraComputeUnits != that1.ExtraComputeUnits {
		return false
	}
	return true
}
func (this *BlockParser) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BlockParser)
	if !ok {
		that2, ok := that.(BlockParser)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ParserArg) != len(that1.ParserArg) {
		return false
	}
	for i := range this.ParserArg {
		if this.ParserArg[i] != that1.ParserArg[i] {
			return false
		}
	}
	if this.ParserFunc != that1.ParserFunc {
		return false
	}
	return true
}
func (m *ServiceApi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServiceApi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServiceApi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApiInterfaces) > 0 {
		for iNdEx := len(m.ApiInterfaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApiInterfaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintServiceApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintServiceApi(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.BlockParsing.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintServiceApi(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApiInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApiInterface) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApiInterface) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExtraComputeUnits != 0 {
		i = encodeVarintServiceApi(dAtA, i, uint64(m.ExtraComputeUnits))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Interface) > 0 {
		i -= len(m.Interface)
		copy(dAtA[i:], m.Interface)
		i = encodeVarintServiceApi(dAtA, i, uint64(len(m.Interface)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockParser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockParser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockParser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ParserFunc != 0 {
		i = encodeVarintServiceApi(dAtA, i, uint64(m.ParserFunc))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ParserArg) > 0 {
		for iNdEx := len(m.ParserArg) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ParserArg[iNdEx])
			copy(dAtA[i:], m.ParserArg[iNdEx])
			i = encodeVarintServiceApi(dAtA, i, uint64(len(m.ParserArg[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintServiceApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovServiceApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ServiceApi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	l = m.BlockParsing.Size()
	n += 1 + l + sovServiceApi(uint64(l))
	if m.ComputeUnits != 0 {
		n += 1 + sovServiceApi(uint64(m.ComputeUnits))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.ApiInterfaces) > 0 {
		for _, e := range m.ApiInterfaces {
			l = e.Size()
			n += 1 + l + sovServiceApi(uint64(l))
		}
	}
	return n
}

func (m *ApiInterface) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Interface)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovServiceApi(uint64(l))
	}
	if m.ExtraComputeUnits != 0 {
		n += 1 + sovServiceApi(uint64(m.ExtraComputeUnits))
	}
	return n
}

func (m *BlockParser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ParserArg) > 0 {
		for _, s := range m.ParserArg {
			l = len(s)
			n += 1 + l + sovServiceApi(uint64(l))
		}
	}
	if m.ParserFunc != 0 {
		n += 1 + sovServiceApi(uint64(m.ParserFunc))
	}
	return n
}

func sovServiceApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozServiceApi(x uint64) (n int) {
	return sovServiceApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ServiceApi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServiceApi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServiceApi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockParsing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockParsing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiInterfaces", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiInterfaces = append(m.ApiInterfaces, ApiInterface{})
			if err := m.ApiInterfaces[len(m.ApiInterfaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ApiInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ApiInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApiInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interface", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Interface = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraComputeUnits", wireType)
			}
			m.ExtraComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExtraComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipServiceApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockParser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockParser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockParser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParserArg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthServiceApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParserArg = append(m.ParserArg, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParserFunc", wireType)
			}
			m.ParserFunc = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ParserFunc |= PARSER_FUNC(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipServiceApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthServiceApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipServiceApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowServiceApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthServiceApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupServiceApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthServiceApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthServiceApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupServiceApi = fmt.Errorf("proto: unexpected end of group")
)
