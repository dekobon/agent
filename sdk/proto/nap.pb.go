// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nap.proto

package proto

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

type AppProtectWAFHealth_AppProtectWAFStatus int32

const (
	AppProtectWAFHealth_UNKNOWN  AppProtectWAFHealth_AppProtectWAFStatus = 0
	AppProtectWAFHealth_ACTIVE   AppProtectWAFHealth_AppProtectWAFStatus = 1
	AppProtectWAFHealth_DEGRADED AppProtectWAFHealth_AppProtectWAFStatus = 2
)

var AppProtectWAFHealth_AppProtectWAFStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACTIVE",
	2: "DEGRADED",
}

var AppProtectWAFHealth_AppProtectWAFStatus_value = map[string]int32{
	"UNKNOWN":  0,
	"ACTIVE":   1,
	"DEGRADED": 2,
}

func (x AppProtectWAFHealth_AppProtectWAFStatus) String() string {
	return proto.EnumName(AppProtectWAFHealth_AppProtectWAFStatus_name, int32(x))
}

func (AppProtectWAFHealth_AppProtectWAFStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f34234efeae954d9, []int{1, 0}
}

// AppProtectWAFDetails reports the details of Nginx App Protect
type AppProtectWAFDetails struct {
	WafVersion              string               `protobuf:"bytes,1,opt,name=waf_version,json=wafVersion,proto3" json:"waf_version"`
	AttackSignaturesVersion string               `protobuf:"bytes,2,opt,name=attack_signatures_version,json=attackSignaturesVersion,proto3" json:"attack_signatures_version"`
	ThreatCampaignsVersion  string               `protobuf:"bytes,3,opt,name=threat_campaigns_version,json=threatCampaignsVersion,proto3" json:"threat_campaigns_version"`
	Health                  *AppProtectWAFHealth `protobuf:"bytes,4,opt,name=health,proto3" json:"health"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *AppProtectWAFDetails) Reset()         { *m = AppProtectWAFDetails{} }
func (m *AppProtectWAFDetails) String() string { return proto.CompactTextString(m) }
func (*AppProtectWAFDetails) ProtoMessage()    {}
func (*AppProtectWAFDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34234efeae954d9, []int{0}
}
func (m *AppProtectWAFDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppProtectWAFDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppProtectWAFDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppProtectWAFDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProtectWAFDetails.Merge(m, src)
}
func (m *AppProtectWAFDetails) XXX_Size() int {
	return m.Size()
}
func (m *AppProtectWAFDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProtectWAFDetails.DiscardUnknown(m)
}

var xxx_messageInfo_AppProtectWAFDetails proto.InternalMessageInfo

func (m *AppProtectWAFDetails) GetWafVersion() string {
	if m != nil {
		return m.WafVersion
	}
	return ""
}

func (m *AppProtectWAFDetails) GetAttackSignaturesVersion() string {
	if m != nil {
		return m.AttackSignaturesVersion
	}
	return ""
}

func (m *AppProtectWAFDetails) GetThreatCampaignsVersion() string {
	if m != nil {
		return m.ThreatCampaignsVersion
	}
	return ""
}

func (m *AppProtectWAFDetails) GetHealth() *AppProtectWAFHealth {
	if m != nil {
		return m.Health
	}
	return nil
}

// AppProtectWAFHealth reports the health details of Nginx App Protect
type AppProtectWAFHealth struct {
	SystemId             string                                  `protobuf:"bytes,1,opt,name=system_id,json=systemId,proto3" json:"system_id"`
	AppProtectWafStatus  AppProtectWAFHealth_AppProtectWAFStatus `protobuf:"varint,2,opt,name=app_protect_waf_status,json=appProtectWafStatus,proto3,enum=f5.nginx.agent.sdk.AppProtectWAFHealth_AppProtectWAFStatus" json:"app_protect_waf_status"`
	DegradedReason       string                                  `protobuf:"bytes,3,opt,name=degraded_reason,json=degradedReason,proto3" json:"degraded_reason"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AppProtectWAFHealth) Reset()         { *m = AppProtectWAFHealth{} }
func (m *AppProtectWAFHealth) String() string { return proto.CompactTextString(m) }
func (*AppProtectWAFHealth) ProtoMessage()    {}
func (*AppProtectWAFHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34234efeae954d9, []int{1}
}
func (m *AppProtectWAFHealth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AppProtectWAFHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AppProtectWAFHealth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AppProtectWAFHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProtectWAFHealth.Merge(m, src)
}
func (m *AppProtectWAFHealth) XXX_Size() int {
	return m.Size()
}
func (m *AppProtectWAFHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProtectWAFHealth.DiscardUnknown(m)
}

var xxx_messageInfo_AppProtectWAFHealth proto.InternalMessageInfo

func (m *AppProtectWAFHealth) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *AppProtectWAFHealth) GetAppProtectWafStatus() AppProtectWAFHealth_AppProtectWAFStatus {
	if m != nil {
		return m.AppProtectWafStatus
	}
	return AppProtectWAFHealth_UNKNOWN
}

func (m *AppProtectWAFHealth) GetDegradedReason() string {
	if m != nil {
		return m.DegradedReason
	}
	return ""
}

func init() {
	proto.RegisterEnum("f5.nginx.agent.sdk.AppProtectWAFHealth_AppProtectWAFStatus", AppProtectWAFHealth_AppProtectWAFStatus_name, AppProtectWAFHealth_AppProtectWAFStatus_value)
	proto.RegisterType((*AppProtectWAFDetails)(nil), "f5.nginx.agent.sdk.AppProtectWAFDetails")
	proto.RegisterType((*AppProtectWAFHealth)(nil), "f5.nginx.agent.sdk.AppProtectWAFHealth")
}

func init() { proto.RegisterFile("nap.proto", fileDescriptor_f34234efeae954d9) }

var fileDescriptor_f34234efeae954d9 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xd3, 0x30,
	0x1c, 0xc7, 0x49, 0x41, 0x65, 0xfd, 0x15, 0xb6, 0xca, 0x45, 0xa3, 0x4c, 0xd0, 0x4c, 0xbd, 0x30,
	0x21, 0x91, 0xa0, 0xa2, 0x5d, 0xd8, 0x0e, 0xb4, 0x6b, 0x81, 0x69, 0x52, 0x41, 0x1e, 0x74, 0x82,
	0x4b, 0xf4, 0x5b, 0xe2, 0xa4, 0x51, 0xdb, 0xc4, 0xb2, 0xdd, 0x0d, 0xde, 0x81, 0x07, 0xdb, 0x81,
	0x03, 0x4f, 0x10, 0xa1, 0x1e, 0xf3, 0x14, 0x08, 0xbb, 0x7f, 0x58, 0xa1, 0xd2, 0x2e, 0xb1, 0xf3,
	0xfd, 0xa7, 0xe8, 0xa3, 0x40, 0x29, 0x41, 0xee, 0x70, 0x91, 0xaa, 0x94, 0x90, 0x70, 0xdf, 0x49,
	0xa2, 0x38, 0xf9, 0xea, 0x60, 0xc4, 0x12, 0xe5, 0xc8, 0x60, 0xb8, 0x03, 0x51, 0x1a, 0xa5, 0xc6,
	0x6f, 0x5c, 0x15, 0xe0, 0x41, 0x8b, 0xf3, 0x0f, 0x22, 0x55, 0xcc, 0x57, 0x67, 0xad, 0x37, 0x1d,
	0xa6, 0x30, 0x1e, 0x49, 0xf2, 0x02, 0xca, 0x97, 0x18, 0x7a, 0x17, 0x4c, 0xc8, 0x38, 0x4d, 0x6a,
	0xd6, 0xae, 0xb5, 0x57, 0x6a, 0x6f, 0xe5, 0x99, 0xfd, 0xb7, 0x4c, 0xe1, 0x12, 0xc3, 0xbe, 0xb9,
	0x93, 0xcf, 0xf0, 0x08, 0x95, 0x42, 0x7f, 0xe8, 0xc9, 0x38, 0x4a, 0x50, 0x4d, 0x04, 0x93, 0x8b,
	0x7e, 0x41, 0xf7, 0x9f, 0xe4, 0x99, 0xbd, 0x3e, 0x44, 0x1f, 0x1a, 0xeb, 0x74, 0xe1, 0xcc, 0xa7,
	0xfb, 0x50, 0x53, 0x03, 0xc1, 0x50, 0x79, 0x3e, 0x8e, 0x39, 0xc6, 0x51, 0xb2, 0x5c, 0xbe, 0xad,
	0x97, 0x1f, 0xe7, 0x99, 0xbd, 0x36, 0x43, 0xb7, 0x8d, 0x73, 0x34, 0x37, 0xe6, 0xbb, 0x27, 0x50,
	0x1c, 0x30, 0x1c, 0xa9, 0x41, 0xed, 0xce, 0xae, 0xb5, 0x57, 0x6e, 0x3e, 0x75, 0xfe, 0xc5, 0xe5,
	0x5c, 0xc3, 0xf3, 0x4e, 0xc7, 0xdb, 0x90, 0x67, 0xf6, 0xac, 0x4a, 0x67, 0x67, 0xe3, 0x47, 0x01,
	0xaa, 0xff, 0xc9, 0x92, 0x67, 0x50, 0x92, 0xdf, 0xa4, 0x62, 0x63, 0x2f, 0x0e, 0x66, 0x1c, 0xef,
	0xe7, 0x99, 0xbd, 0x14, 0xe9, 0x86, 0xb9, 0x1e, 0x07, 0xe4, 0xbb, 0x05, 0xdb, 0xc8, 0xb9, 0xc7,
	0xcd, 0x88, 0xf7, 0x87, 0xb5, 0x54, 0xa8, 0x26, 0x52, 0x13, 0xdc, 0x6c, 0x1e, 0xdc, 0xf0, 0x0b,
	0xaf, 0x6b, 0xa7, 0x7a, 0xa2, 0xbd, 0x93, 0x67, 0xf6, 0x9a, 0x79, 0x5a, 0xc5, 0x65, 0x01, 0x43,
	0x53, 0x20, 0x87, 0xb0, 0x15, 0xb0, 0x48, 0x60, 0xc0, 0x02, 0x4f, 0x30, 0x94, 0x0b, 0xdc, 0xd5,
	0x3c, 0xb3, 0x57, 0x2d, 0xba, 0x39, 0x17, 0xa8, 0x7e, 0x6f, 0x1c, 0xae, 0xf0, 0x98, 0x8d, 0x96,
	0xe1, 0xee, 0xa7, 0xde, 0x49, 0xef, 0xfd, 0x59, 0xaf, 0x72, 0x8b, 0x00, 0x14, 0x5b, 0x47, 0x1f,
	0x8f, 0xfb, 0xdd, 0x8a, 0x45, 0xee, 0xc1, 0x46, 0xa7, 0xfb, 0x96, 0xb6, 0x3a, 0xdd, 0x4e, 0xa5,
	0xd0, 0xee, 0x5d, 0x4d, 0xeb, 0xd6, 0xcf, 0x69, 0xdd, 0xfa, 0x35, 0xad, 0x5b, 0x5f, 0x5e, 0x47,
	0xb1, 0x1a, 0xe1, 0xb9, 0xe3, 0xa7, 0xe3, 0x57, 0xe1, 0xbe, 0xab, 0x09, 0xb8, 0x9a, 0x80, 0xcb,
	0x45, 0x1a, 0x4c, 0x7c, 0x65, 0xb4, 0xe7, 0x46, 0x93, 0xc1, 0xd0, 0xbd, 0x68, 0xba, 0xfa, 0x1f,
	0x3f, 0xd0, 0xcf, 0xf3, 0xa2, 0x3e, 0x5e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x98, 0x6c, 0x67,
	0x98, 0x1d, 0x03, 0x00, 0x00,
}

func (m *AppProtectWAFDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppProtectWAFDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppProtectWAFDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Health != nil {
		{
			size, err := m.Health.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.ThreatCampaignsVersion) > 0 {
		i -= len(m.ThreatCampaignsVersion)
		copy(dAtA[i:], m.ThreatCampaignsVersion)
		i = encodeVarintNap(dAtA, i, uint64(len(m.ThreatCampaignsVersion)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AttackSignaturesVersion) > 0 {
		i -= len(m.AttackSignaturesVersion)
		copy(dAtA[i:], m.AttackSignaturesVersion)
		i = encodeVarintNap(dAtA, i, uint64(len(m.AttackSignaturesVersion)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.WafVersion) > 0 {
		i -= len(m.WafVersion)
		copy(dAtA[i:], m.WafVersion)
		i = encodeVarintNap(dAtA, i, uint64(len(m.WafVersion)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AppProtectWAFHealth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AppProtectWAFHealth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AppProtectWAFHealth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.DegradedReason) > 0 {
		i -= len(m.DegradedReason)
		copy(dAtA[i:], m.DegradedReason)
		i = encodeVarintNap(dAtA, i, uint64(len(m.DegradedReason)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AppProtectWafStatus != 0 {
		i = encodeVarintNap(dAtA, i, uint64(m.AppProtectWafStatus))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SystemId) > 0 {
		i -= len(m.SystemId)
		copy(dAtA[i:], m.SystemId)
		i = encodeVarintNap(dAtA, i, uint64(len(m.SystemId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNap(dAtA []byte, offset int, v uint64) int {
	offset -= sovNap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AppProtectWAFDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WafVersion)
	if l > 0 {
		n += 1 + l + sovNap(uint64(l))
	}
	l = len(m.AttackSignaturesVersion)
	if l > 0 {
		n += 1 + l + sovNap(uint64(l))
	}
	l = len(m.ThreatCampaignsVersion)
	if l > 0 {
		n += 1 + l + sovNap(uint64(l))
	}
	if m.Health != nil {
		l = m.Health.Size()
		n += 1 + l + sovNap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AppProtectWAFHealth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SystemId)
	if l > 0 {
		n += 1 + l + sovNap(uint64(l))
	}
	if m.AppProtectWafStatus != 0 {
		n += 1 + sovNap(uint64(m.AppProtectWafStatus))
	}
	l = len(m.DegradedReason)
	if l > 0 {
		n += 1 + l + sovNap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNap(x uint64) (n int) {
	return sovNap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AppProtectWAFDetails) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNap
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
			return fmt.Errorf("proto: AppProtectWAFDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppProtectWAFDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WafVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WafVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttackSignaturesVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttackSignaturesVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreatCampaignsVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThreatCampaignsVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Health", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Health == nil {
				m.Health = &AppProtectWAFHealth{}
			}
			if err := m.Health.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AppProtectWAFHealth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNap
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
			return fmt.Errorf("proto: AppProtectWAFHealth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AppProtectWAFHealth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SystemId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppProtectWafStatus", wireType)
			}
			m.AppProtectWafStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppProtectWafStatus |= AppProtectWAFHealth_AppProtectWAFStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DegradedReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNap
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
				return ErrInvalidLengthNap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DegradedReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNap
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNap
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
					return 0, ErrIntOverflowNap
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
					return 0, ErrIntOverflowNap
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
				return 0, ErrInvalidLengthNap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNap = fmt.Errorf("proto: unexpected end of group")
)
