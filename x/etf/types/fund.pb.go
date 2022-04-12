// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etf/fund.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Holding struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Percent int64  `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"`
	// Pool ID of the Pool for this holding on Broker
	PoolId string `protobuf:"bytes,3,opt,name=poolId,proto3" json:"poolId,omitempty"`
}

func (m *Holding) Reset()         { *m = Holding{} }
func (m *Holding) String() string { return proto.CompactTextString(m) }
func (*Holding) ProtoMessage()    {}
func (*Holding) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f7de7f2b67d1612, []int{0}
}
func (m *Holding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Holding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Holding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Holding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Holding.Merge(m, src)
}
func (m *Holding) XXX_Size() int {
	return m.Size()
}
func (m *Holding) XXX_DiscardUnknown() {
	xxx_messageInfo_Holding.DiscardUnknown(m)
}

var xxx_messageInfo_Holding proto.InternalMessageInfo

func (m *Holding) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Holding) GetPercent() int64 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *Holding) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type Fund struct {
	Id           string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address      string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Symbol       string     `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name         string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description  string     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Shares       types.Coin `protobuf:"bytes,6,opt,name=shares,proto3" json:"shares" yaml:"total_shares"`
	Broker       string     `protobuf:"bytes,7,opt,name=broker,proto3" json:"broker,omitempty"`
	Holdings     []Holding  `protobuf:"bytes,8,rep,name=holdings,proto3" json:"holdings" yaml:"holdings" json:"holdings"`
	Rebalance    int64      `protobuf:"varint,9,opt,name=rebalance,proto3" json:"rebalance,omitempty"`
	BaseDenom    string     `protobuf:"bytes,10,opt,name=baseDenom,proto3" json:"baseDenom,omitempty"`
	ConnectionId string     `protobuf:"bytes,11,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	Creator      string     `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Fund) Reset()         { *m = Fund{} }
func (m *Fund) String() string { return proto.CompactTextString(m) }
func (*Fund) ProtoMessage()    {}
func (*Fund) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f7de7f2b67d1612, []int{1}
}
func (m *Fund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fund.Merge(m, src)
}
func (m *Fund) XXX_Size() int {
	return m.Size()
}
func (m *Fund) XXX_DiscardUnknown() {
	xxx_messageInfo_Fund.DiscardUnknown(m)
}

var xxx_messageInfo_Fund proto.InternalMessageInfo

func (m *Fund) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Fund) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Fund) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Fund) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fund) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Fund) GetShares() types.Coin {
	if m != nil {
		return m.Shares
	}
	return types.Coin{}
}

func (m *Fund) GetBroker() string {
	if m != nil {
		return m.Broker
	}
	return ""
}

func (m *Fund) GetHoldings() []Holding {
	if m != nil {
		return m.Holdings
	}
	return nil
}

func (m *Fund) GetRebalance() int64 {
	if m != nil {
		return m.Rebalance
	}
	return 0
}

func (m *Fund) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *Fund) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Fund) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Holding)(nil), "defundhub.defund.etf.Holding")
	proto.RegisterType((*Fund)(nil), "defundhub.defund.etf.Fund")
}

func init() { proto.RegisterFile("etf/fund.proto", fileDescriptor_2f7de7f2b67d1612) }

var fileDescriptor_2f7de7f2b67d1612 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xce, 0x26, 0x69, 0xd2, 0x38, 0x55, 0x0f, 0x26, 0x42, 0xa6, 0xc0, 0x66, 0xb5, 0x17, 0xc2,
	0x81, 0x5d, 0xb5, 0xdc, 0x38, 0xa6, 0x08, 0xd1, 0x0b, 0x12, 0x39, 0x72, 0x41, 0x5e, 0x7b, 0x92,
	0x2c, 0xec, 0x7a, 0x56, 0xb6, 0x83, 0xc8, 0x5b, 0xf0, 0x50, 0x1c, 0x7a, 0xec, 0x91, 0x53, 0x85,
	0x92, 0x37, 0xe0, 0x09, 0x90, 0xed, 0x0d, 0x2d, 0x52, 0x6f, 0xf3, 0xcd, 0xcf, 0xe7, 0x99, 0xef,
	0x33, 0x39, 0x05, 0xbb, 0xcc, 0x97, 0x1b, 0x25, 0xb3, 0x46, 0xa3, 0x45, 0x3a, 0x91, 0xe0, 0xd0,
	0x7a, 0x53, 0x64, 0x21, 0xca, 0xc0, 0x2e, 0xcf, 0x26, 0x2b, 0x5c, 0xa1, 0x6f, 0xc8, 0x5d, 0x14,
	0x7a, 0xcf, 0x62, 0x81, 0xa6, 0x46, 0x93, 0x17, 0xdc, 0x40, 0xfe, 0xed, 0xbc, 0x00, 0xcb, 0xcf,
	0x73, 0x81, 0xa5, 0x0a, 0xf5, 0xf4, 0x23, 0x19, 0xbe, 0xc7, 0x4a, 0x96, 0x6a, 0x45, 0x27, 0xe4,
	0x48, 0x82, 0xc2, 0x9a, 0x45, 0x49, 0x34, 0x1b, 0x2d, 0x02, 0xa0, 0x8c, 0x0c, 0x1b, 0xd0, 0x02,
	0x94, 0x65, 0xdd, 0x24, 0x9a, 0xf5, 0x16, 0x07, 0x48, 0x1f, 0x93, 0x41, 0x83, 0x58, 0x5d, 0x49,
	0xd6, 0xf3, 0x03, 0x2d, 0x4a, 0x7f, 0xf6, 0x48, 0xff, 0xdd, 0x46, 0x49, 0x7a, 0x4a, 0xba, 0xa5,
	0x6c, 0xd9, 0xba, 0xa5, 0x74, 0x54, 0x5c, 0x4a, 0x0d, 0xc6, 0x78, 0xaa, 0xd1, 0xe2, 0x00, 0x1d,
	0x95, 0xd9, 0xd6, 0x05, 0x56, 0x07, 0xaa, 0x80, 0x28, 0x25, 0x7d, 0xc5, 0x6b, 0x60, 0x7d, 0x9f,
	0xf5, 0x31, 0x4d, 0xc8, 0x58, 0x82, 0x11, 0xba, 0x6c, 0x6c, 0x89, 0x8a, 0x1d, 0xf9, 0xd2, 0xfd,
	0x14, 0xfd, 0x40, 0x06, 0x66, 0xcd, 0x35, 0x18, 0x36, 0x48, 0xa2, 0xd9, 0xf8, 0xe2, 0x49, 0x16,
	0x44, 0xc8, 0x9c, 0x08, 0x59, 0x2b, 0x42, 0x76, 0x89, 0xa5, 0x9a, 0x3f, 0xbd, 0xbe, 0x9d, 0x76,
	0xfe, 0xdc, 0x4e, 0x1f, 0x6d, 0x79, 0x5d, 0xbd, 0x49, 0x2d, 0x5a, 0x5e, 0x7d, 0x0e, 0xc3, 0xe9,
	0xa2, 0x65, 0x71, 0xdb, 0x15, 0x1a, 0xbf, 0x82, 0x66, 0xc3, 0xb0, 0x5d, 0x40, 0xb4, 0x20, 0xc7,
	0xeb, 0xa0, 0x9d, 0x61, 0xc7, 0x49, 0x6f, 0x36, 0xbe, 0x78, 0x9e, 0x3d, 0x64, 0x4d, 0xd6, 0x2a,
	0x3c, 0x7f, 0xd1, 0xbe, 0x36, 0x0d, 0xaf, 0x1d, 0x86, 0xd3, 0xe4, 0x8b, 0x41, 0x75, 0x0f, 0x2f,
	0xfe, 0xf1, 0xd2, 0x67, 0x64, 0xa4, 0xa1, 0xe0, 0x15, 0x57, 0x02, 0xd8, 0xc8, 0x1b, 0x70, 0x97,
	0x70, 0x55, 0x77, 0xd3, 0x5b, 0x6f, 0x1b, 0xf1, 0xcb, 0xdd, 0x25, 0x68, 0x4a, 0x4e, 0x04, 0x2a,
	0x05, 0xc2, 0xa9, 0x72, 0x25, 0xd9, 0xd8, 0x37, 0xfc, 0x97, 0x73, 0x9e, 0x08, 0x0d, 0xdc, 0xa2,
	0x66, 0x27, 0xc1, 0x93, 0x16, 0xce, 0x2f, 0xaf, 0x77, 0x71, 0x74, 0xb3, 0x8b, 0xa3, 0xdf, 0xbb,
	0x38, 0xfa, 0xb1, 0x8f, 0x3b, 0x37, 0xfb, 0xb8, 0xf3, 0x6b, 0x1f, 0x77, 0x3e, 0xbd, 0x5c, 0x95,
	0xd6, 0x5d, 0x28, 0xb0, 0xce, 0xc3, 0x95, 0xaf, 0x2a, 0x5e, 0x98, 0x36, 0xce, 0xbf, 0xe7, 0xee,
	0xbf, 0xda, 0x6d, 0x03, 0xa6, 0x18, 0xf8, 0x5f, 0xf6, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0xc1, 0xf6, 0xcf, 0xc3, 0x02, 0x00, 0x00,
}

func (m *Holding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Holding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Holding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintFund(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Percent != 0 {
		i = encodeVarintFund(dAtA, i, uint64(m.Percent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Fund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintFund(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintFund(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x52
	}
	if m.Rebalance != 0 {
		i = encodeVarintFund(dAtA, i, uint64(m.Rebalance))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Holdings) > 0 {
		for iNdEx := len(m.Holdings) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Holdings[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFund(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Broker) > 0 {
		i -= len(m.Broker)
		copy(dAtA[i:], m.Broker)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Broker)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size, err := m.Shares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFund(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintFund(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFund(dAtA []byte, offset int, v uint64) int {
	offset -= sovFund(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Holding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	if m.Percent != 0 {
		n += 1 + sovFund(uint64(m.Percent))
	}
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	return n
}

func (m *Fund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = m.Shares.Size()
	n += 1 + l + sovFund(uint64(l))
	l = len(m.Broker)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	if len(m.Holdings) > 0 {
		for _, e := range m.Holdings {
			l = e.Size()
			n += 1 + l + sovFund(uint64(l))
		}
	}
	if m.Rebalance != 0 {
		n += 1 + sovFund(uint64(m.Rebalance))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFund(uint64(l))
	}
	return n
}

func sovFund(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFund(x uint64) (n int) {
	return sovFund(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Holding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFund
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
			return fmt.Errorf("proto: Holding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Holding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percent", wireType)
			}
			m.Percent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Percent |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFund
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
func (m *Fund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFund
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
			return fmt.Errorf("proto: Fund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Broker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Broker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Holdings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Holdings = append(m.Holdings, Holding{})
			if err := m.Holdings[len(m.Holdings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rebalance", wireType)
			}
			m.Rebalance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rebalance |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFund
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
				return ErrInvalidLengthFund
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFund
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFund(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFund
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
func skipFund(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFund
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
					return 0, ErrIntOverflowFund
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
					return 0, ErrIntOverflowFund
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
				return 0, ErrInvalidLengthFund
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFund
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFund
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFund        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFund          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFund = fmt.Errorf("proto: unexpected end of group")
)
