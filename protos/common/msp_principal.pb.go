// Code generated by protoc-gen-go.
// source: common/msp_principal.proto
// DO NOT EDIT!

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MSPPrincipal_Classification int32

const (
	MSPPrincipal_ROLE MSPPrincipal_Classification = 0
	// one of a member of MSP network, and the one of an
	// administrator of an MSP network
	MSPPrincipal_ORGANIZATION_UNIT MSPPrincipal_Classification = 1
	// groupping of entities, per MSP affiliation
	// E.g., this can well be represented by an MSP's
	// Organization unit
	MSPPrincipal_IDENTITY MSPPrincipal_Classification = 2
)

var MSPPrincipal_Classification_name = map[int32]string{
	0: "ROLE",
	1: "ORGANIZATION_UNIT",
	2: "IDENTITY",
}
var MSPPrincipal_Classification_value = map[string]int32{
	"ROLE":              0,
	"ORGANIZATION_UNIT": 1,
	"IDENTITY":          2,
}

func (x MSPPrincipal_Classification) String() string {
	return proto.EnumName(MSPPrincipal_Classification_name, int32(x))
}
func (MSPPrincipal_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

type MSPRole_MSPRoleType int32

const (
	MSPRole_MEMBER MSPRole_MSPRoleType = 0
	MSPRole_ADMIN  MSPRole_MSPRoleType = 1
)

var MSPRole_MSPRoleType_name = map[int32]string{
	0: "MEMBER",
	1: "ADMIN",
}
var MSPRole_MSPRoleType_value = map[string]int32{
	"MEMBER": 0,
	"ADMIN":  1,
}

func (x MSPRole_MSPRoleType) String() string {
	return proto.EnumName(MSPRole_MSPRoleType_name, int32(x))
}
func (MSPRole_MSPRoleType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

// MSPPrincipal aims to represent an MSP-centric set of identities.
// In particular, this structure allows for definition of
//  - a group of identities that are member of the same MSP
//  - a group of identities that are member of the same organization unit
//    in the same MSP
//  - a group of identities that are administering a specific MSP
//  - a specific identity
// Expressing these groups is done given two fields of the fields below
//  - Classification, that defines the type of classification of identities
//    in an MSP this principal would be defined on; Classification can take
//    three values:
//     (i)  ByMSPRole: that represents a classification of identities within
//          MSP based on one of the two pre-defined MSP rules, "member" and "admin"
//     (ii) ByOrganizationUnit: that represents a classification of identities
//          within MSP based on the organization unit an identity belongs to
//     (iii)ByIdentity that denotes that MSPPrincipal is mapped to a single
//          identity/certificate; this would mean that the Principal bytes
//          message
type MSPPrincipal struct {
	// Classification describes the way that one should process
	// Principal. An Classification value of "ByOrganizationUnit" reflects
	// that "Principal" contains the name of an organization this MSP
	// handles. A Classification value "ByIdentity" means that
	// "Principal" contains a specific identity. Default value
	// denotes that Principal contains one of the groups by
	// default supported by all MSPs ("admin" or "member").
	PrincipalClassification MSPPrincipal_Classification `protobuf:"varint,1,opt,name=principal_classification,json=principalClassification,enum=common.MSPPrincipal_Classification" json:"principal_classification,omitempty"`
	// Principal completes the policy principal definition. For the default
	// principal types, Principal can be either "Admin" or "Member".
	// For the ByOrganizationUnit/ByIdentity values of Classification,
	// PolicyPrincipal acquires its value from an organization unit or
	// identity, respectively.
	Principal []byte `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
}

func (m *MSPPrincipal) Reset()                    { *m = MSPPrincipal{} }
func (m *MSPPrincipal) String() string            { return proto.CompactTextString(m) }
func (*MSPPrincipal) ProtoMessage()               {}
func (*MSPPrincipal) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// OrganizationUnit governs the organization of the Principal
// field of a policy principal when a specific organization unity members
// are to be defined within a policy principal.
type OrganizationUnit struct {
	// MSPIdentifier represents the identifier of the MSP this organization unit
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// OrganizationUnitIdentifier defines the organization unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
}

func (m *OrganizationUnit) Reset()                    { *m = OrganizationUnit{} }
func (m *OrganizationUnit) String() string            { return proto.CompactTextString(m) }
func (*OrganizationUnit) ProtoMessage()               {}
func (*OrganizationUnit) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

// MSPRole governs the organization of the Principal
// field of an MSPPrincipal when it aims to define one of the
// two dedicated roles within an MSP: Admin and Members.
type MSPRole struct {
	// MSPIdentifier represents the identifier of the MSP this principal
	// refers to
	MSPIdentifier string `protobuf:"bytes,1,opt,name=MSPIdentifier" json:"MSPIdentifier,omitempty"`
	// MSPRoleType defines which of the available, pre-defined MSP-roles
	// an identiy should posess inside the MSP with identifier MSPidentifier
	Role MSPRole_MSPRoleType `protobuf:"varint,2,opt,name=Role,enum=common.MSPRole_MSPRoleType" json:"Role,omitempty"`
}

func (m *MSPRole) Reset()                    { *m = MSPRole{} }
func (m *MSPRole) String() string            { return proto.CompactTextString(m) }
func (*MSPRole) ProtoMessage()               {}
func (*MSPRole) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func init() {
	proto.RegisterType((*MSPPrincipal)(nil), "common.MSPPrincipal")
	proto.RegisterType((*OrganizationUnit)(nil), "common.OrganizationUnit")
	proto.RegisterType((*MSPRole)(nil), "common.MSPRole")
	proto.RegisterEnum("common.MSPPrincipal_Classification", MSPPrincipal_Classification_name, MSPPrincipal_Classification_value)
	proto.RegisterEnum("common.MSPRole_MSPRoleType", MSPRole_MSPRoleType_name, MSPRole_MSPRoleType_value)
}

func init() { proto.RegisterFile("common/msp_principal.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6a, 0xea, 0x40,
	0x1c, 0xc5, 0x8d, 0x78, 0xbd, 0xfa, 0xbf, 0x5e, 0xc9, 0x1d, 0xb8, 0x54, 0x5a, 0x29, 0x92, 0xba,
	0x10, 0x4a, 0x13, 0xb0, 0x0f, 0x50, 0xb4, 0x86, 0x12, 0x68, 0x3e, 0x18, 0xe3, 0xa2, 0x2e, 0x1a,
	0x62, 0x1c, 0x75, 0x20, 0xce, 0x84, 0x49, 0x5c, 0xd8, 0x45, 0x97, 0x7d, 0xbb, 0xbe, 0x53, 0xc9,
	0x88, 0x9a, 0x74, 0xd5, 0x55, 0xc8, 0x39, 0xe7, 0x77, 0x0e, 0x33, 0x0c, 0x5c, 0x46, 0x7c, 0xbb,
	0xe5, 0xcc, 0xd8, 0xa6, 0x49, 0x90, 0x08, 0xca, 0x22, 0x9a, 0x84, 0xb1, 0x9e, 0x08, 0x9e, 0x71,
	0x54, 0x3f, 0x78, 0xda, 0xa7, 0x02, 0x2d, 0x7b, 0xea, 0x79, 0x47, 0x1b, 0xbd, 0x42, 0xe7, 0x94,
	0x0d, 0xa2, 0x38, 0x4c, 0x53, 0xba, 0xa2, 0x51, 0x98, 0x51, 0xce, 0x3a, 0x4a, 0x4f, 0x19, 0xb4,
	0x87, 0x37, 0xfa, 0x81, 0xd5, 0x8b, 0x9c, 0xfe, 0x58, 0x8a, 0xe2, 0x8b, 0x53, 0x49, 0xd9, 0x40,
	0x5d, 0x68, 0x9e, 0xac, 0x4e, 0xb5, 0xa7, 0x0c, 0x5a, 0xf8, 0x2c, 0x68, 0x0f, 0xd0, 0xfe, 0x96,
	0x6f, 0x40, 0x0d, 0xbb, 0xcf, 0xa6, 0x5a, 0x41, 0xff, 0xe1, 0x9f, 0x8b, 0x9f, 0x46, 0x8e, 0x35,
	0x1f, 0xf9, 0x96, 0xeb, 0x04, 0x33, 0xc7, 0xf2, 0x55, 0x05, 0xb5, 0xa0, 0x61, 0x4d, 0x4c, 0xc7,
	0xb7, 0xfc, 0x17, 0xb5, 0xaa, 0xbd, 0x83, 0xea, 0x8a, 0x75, 0xc8, 0xe8, 0x9b, 0xc4, 0x67, 0x8c,
	0x66, 0xa8, 0x0f, 0x7f, 0xed, 0xa9, 0x67, 0x2d, 0x09, 0xcb, 0xe8, 0x8a, 0x12, 0x21, 0xcf, 0xd1,
	0xc4, 0x65, 0x11, 0x4d, 0xe0, 0x9a, 0x17, 0xc8, 0x30, 0x0e, 0x76, 0x8c, 0x66, 0x01, 0x3d, 0x63,
	0x55, 0x89, 0x75, 0xcb, 0xa9, 0x7c, 0xe1, 0xdc, 0xa2, 0x7d, 0x28, 0xf0, 0xdb, 0x9e, 0x7a, 0x98,
	0xc7, 0xe4, 0x87, 0xbb, 0x06, 0xd4, 0xf2, 0xb4, 0x6c, 0x6f, 0x0f, 0xaf, 0x0a, 0x97, 0x9b, 0xcb,
	0xc7, 0xaf, 0xbf, 0x4f, 0x08, 0x96, 0x41, 0xad, 0x0f, 0x7f, 0x0a, 0x22, 0x02, 0xa8, 0xdb, 0xa6,
	0x3d, 0x36, 0xb1, 0x5a, 0x41, 0x4d, 0xf8, 0x35, 0x9a, 0xd8, 0x96, 0xa3, 0x2a, 0xe3, 0xbb, 0xf9,
	0xed, 0x9a, 0x66, 0x9b, 0xdd, 0x22, 0x2f, 0x34, 0x36, 0xfb, 0x84, 0x88, 0x98, 0x2c, 0xd7, 0x44,
	0x18, 0xab, 0x70, 0x21, 0x68, 0x64, 0xc8, 0x77, 0x90, 0x1a, 0x87, 0xb9, 0x45, 0x5d, 0xfe, 0xde,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xa4, 0x4f, 0x15, 0x34, 0x02, 0x00, 0x00,
}
