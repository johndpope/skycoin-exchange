// Code generated by protoc-gen-go.
// source: pp.deposit.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GetDepositAddrReq struct {
	Pubkey           *string `protobuf:"bytes,10,opt,name=pubkey" json:"pubkey,omitempty"`
	CoinType         *string `protobuf:"bytes,11,opt,name=coin_type" json:"coin_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDepositAddrReq) Reset()                    { *m = GetDepositAddrReq{} }
func (m *GetDepositAddrReq) String() string            { return proto.CompactTextString(m) }
func (*GetDepositAddrReq) ProtoMessage()               {}
func (*GetDepositAddrReq) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *GetDepositAddrReq) GetPubkey() string {
	if m != nil && m.Pubkey != nil {
		return *m.Pubkey
	}
	return ""
}

func (m *GetDepositAddrReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

type GetDepositAddrRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	CoinType         *string `protobuf:"bytes,11,opt,name=coin_type" json:"coin_type,omitempty"`
	Address          *string `protobuf:"bytes,12,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDepositAddrRes) Reset()                    { *m = GetDepositAddrRes{} }
func (m *GetDepositAddrRes) String() string            { return proto.CompactTextString(m) }
func (*GetDepositAddrRes) ProtoMessage()               {}
func (*GetDepositAddrRes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *GetDepositAddrRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetDepositAddrRes) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *GetDepositAddrRes) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDepositAddrReq)(nil), "pp.GetDepositAddrReq")
	proto.RegisterType((*GetDepositAddrRes)(nil), "pp.GetDepositAddrRes")
}

func init() { proto.RegisterFile("pp.deposit.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0xcc, 0xb1, 0x0e, 0x82, 0x30,
	0x14, 0x85, 0xe1, 0xc0, 0x80, 0xe1, 0x62, 0x44, 0x3a, 0x35, 0x4c, 0x84, 0x89, 0xa9, 0x83, 0x83,
	0xbb, 0x89, 0x89, 0x3b, 0x3e, 0x80, 0x51, 0x7a, 0x07, 0xa2, 0x70, 0x8f, 0x6d, 0x19, 0x78, 0x7b,
	0x83, 0x8c, 0xba, 0x7e, 0x39, 0xe7, 0xa7, 0x3d, 0x60, 0x2c, 0x43, 0x7c, 0x1f, 0x0c, 0x9c, 0x04,
	0x51, 0x31, 0x50, 0xe6, 0x80, 0xe9, 0x64, 0x18, 0x64, 0x5c, 0xb1, 0x3e, 0x52, 0x71, 0xe1, 0x70,
	0x5e, 0x87, 0x27, 0x6b, 0x5d, 0xcb, 0x6f, 0xb5, 0xa3, 0x04, 0xd3, 0xe3, 0xc9, 0xb3, 0xa6, 0x2a,
	0x6a, 0x52, 0x55, 0x50, 0xda, 0x49, 0x3f, 0xde, 0xc2, 0x0c, 0xd6, 0xd9, 0x42, 0xf5, 0xf5, 0xf7,
	0xe7, 0x55, 0x49, 0x89, 0x63, 0x3f, 0xbd, 0x82, 0x8e, 0xaa, 0xb8, 0xc9, 0x0e, 0x64, 0x00, 0xd3,
	0x7e, 0xe5, 0x4f, 0x43, 0xe5, 0xb4, 0xb9, 0x5b, 0xeb, 0xd8, 0x7b, 0xbd, 0x5d, 0xe0, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x86, 0x8a, 0x61, 0xb4, 0x00, 0x00, 0x00,
}
