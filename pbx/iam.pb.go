// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pbx/iam.proto

package pbx

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RequestHeader struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Domain               string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	ClientKey            string   `protobuf:"bytes,3,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{0}
}

func (m *RequestHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHeader.Unmarshal(m, b)
}
func (m *RequestHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHeader.Marshal(b, m, deterministic)
}
func (m *RequestHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHeader.Merge(m, src)
}
func (m *RequestHeader) XXX_Size() int {
	return xxx_messageInfo_RequestHeader.Size(m)
}
func (m *RequestHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHeader proto.InternalMessageInfo

func (m *RequestHeader) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *RequestHeader) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RequestHeader) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

// ----------------------------------------
// AdminUser Model
// ----------------------------------------
type AdminUser struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Fullname             string               `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	AccountAccess        bool                 `protobuf:"varint,6,opt,name=account_access,json=accountAccess,proto3" json:"account_access,omitempty"`
	LoginCounter         int32                `protobuf:"varint,7,opt,name=login_counter,json=loginCounter,proto3" json:"login_counter,omitempty"`
	LastLogin            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	AccountExpiry        *timestamp.Timestamp `protobuf:"bytes,9,opt,name=account_expiry,json=accountExpiry,proto3" json:"account_expiry,omitempty"`
	Photo                string               `protobuf:"bytes,10,opt,name=photo,proto3" json:"photo,omitempty"`
	PwdExpiry            bool                 `protobuf:"varint,11,opt,name=pwd_expiry,json=pwdExpiry,proto3" json:"pwd_expiry,omitempty"`
	PwdExpiryTime        *timestamp.Timestamp `protobuf:"bytes,12,opt,name=pwd_expiry_time,json=pwdExpiryTime,proto3" json:"pwd_expiry_time,omitempty"`
	PwdLifeInDays        int32                `protobuf:"varint,13,opt,name=pwd_life_in_days,json=pwdLifeInDays,proto3" json:"pwd_life_in_days,omitempty"`
	ForcePwdChange       bool                 `protobuf:"varint,14,opt,name=force_pwd_change,json=forcePwdChange,proto3" json:"force_pwd_change,omitempty"`
	Institution          string               `protobuf:"bytes,15,opt,name=institution,proto3" json:"institution,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AdminUser) Reset()         { *m = AdminUser{} }
func (m *AdminUser) String() string { return proto.CompactTextString(m) }
func (*AdminUser) ProtoMessage()    {}
func (*AdminUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{1}
}

func (m *AdminUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminUser.Unmarshal(m, b)
}
func (m *AdminUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminUser.Marshal(b, m, deterministic)
}
func (m *AdminUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminUser.Merge(m, src)
}
func (m *AdminUser) XXX_Size() int {
	return xxx_messageInfo_AdminUser.Size(m)
}
func (m *AdminUser) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminUser.DiscardUnknown(m)
}

var xxx_messageInfo_AdminUser proto.InternalMessageInfo

func (m *AdminUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AdminUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AdminUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AdminUser) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *AdminUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AdminUser) GetAccountAccess() bool {
	if m != nil {
		return m.AccountAccess
	}
	return false
}

func (m *AdminUser) GetLoginCounter() int32 {
	if m != nil {
		return m.LoginCounter
	}
	return 0
}

func (m *AdminUser) GetLastLogin() *timestamp.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func (m *AdminUser) GetAccountExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.AccountExpiry
	}
	return nil
}

func (m *AdminUser) GetPhoto() string {
	if m != nil {
		return m.Photo
	}
	return ""
}

func (m *AdminUser) GetPwdExpiry() bool {
	if m != nil {
		return m.PwdExpiry
	}
	return false
}

func (m *AdminUser) GetPwdExpiryTime() *timestamp.Timestamp {
	if m != nil {
		return m.PwdExpiryTime
	}
	return nil
}

func (m *AdminUser) GetPwdLifeInDays() int32 {
	if m != nil {
		return m.PwdLifeInDays
	}
	return 0
}

func (m *AdminUser) GetForcePwdChange() bool {
	if m != nil {
		return m.ForcePwdChange
	}
	return false
}

func (m *AdminUser) GetInstitution() string {
	if m != nil {
		return m.Institution
	}
	return ""
}

type AdminUserFindRequest struct {
	Filter               string   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminUserFindRequest) Reset()         { *m = AdminUserFindRequest{} }
func (m *AdminUserFindRequest) String() string { return proto.CompactTextString(m) }
func (*AdminUserFindRequest) ProtoMessage()    {}
func (*AdminUserFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{2}
}

func (m *AdminUserFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminUserFindRequest.Unmarshal(m, b)
}
func (m *AdminUserFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminUserFindRequest.Marshal(b, m, deterministic)
}
func (m *AdminUserFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminUserFindRequest.Merge(m, src)
}
func (m *AdminUserFindRequest) XXX_Size() int {
	return xxx_messageInfo_AdminUserFindRequest.Size(m)
}
func (m *AdminUserFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminUserFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdminUserFindRequest proto.InternalMessageInfo

func (m *AdminUserFindRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *AdminUserFindRequest) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type AdminUserFindResponse struct {
	AdminUsers           []*AdminUser `protobuf:"bytes,1,rep,name=adminUsers,proto3" json:"adminUsers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AdminUserFindResponse) Reset()         { *m = AdminUserFindResponse{} }
func (m *AdminUserFindResponse) String() string { return proto.CompactTextString(m) }
func (*AdminUserFindResponse) ProtoMessage()    {}
func (*AdminUserFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{3}
}

func (m *AdminUserFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminUserFindResponse.Unmarshal(m, b)
}
func (m *AdminUserFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminUserFindResponse.Marshal(b, m, deterministic)
}
func (m *AdminUserFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminUserFindResponse.Merge(m, src)
}
func (m *AdminUserFindResponse) XXX_Size() int {
	return xxx_messageInfo_AdminUserFindResponse.Size(m)
}
func (m *AdminUserFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminUserFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdminUserFindResponse proto.InternalMessageInfo

func (m *AdminUserFindResponse) GetAdminUsers() []*AdminUser {
	if m != nil {
		return m.AdminUsers
	}
	return nil
}

// ----------------------------------------
// AuthRole Model
// ----------------------------------------
type AuthRole struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Default              bool     `protobuf:"varint,4,opt,name=default,proto3" json:"default,omitempty"`
	Status               bool     `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRole) Reset()         { *m = AuthRole{} }
func (m *AuthRole) String() string { return proto.CompactTextString(m) }
func (*AuthRole) ProtoMessage()    {}
func (*AuthRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{4}
}

func (m *AuthRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRole.Unmarshal(m, b)
}
func (m *AuthRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRole.Marshal(b, m, deterministic)
}
func (m *AuthRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRole.Merge(m, src)
}
func (m *AuthRole) XXX_Size() int {
	return xxx_messageInfo_AuthRole.Size(m)
}
func (m *AuthRole) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRole.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRole proto.InternalMessageInfo

func (m *AuthRole) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRole) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AuthRole) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func (m *AuthRole) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// AuthLogin
type AuthLoginRequest struct {
	Username             string         `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string         `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Header               *RequestHeader `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AuthLoginRequest) Reset()         { *m = AuthLoginRequest{} }
func (m *AuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*AuthLoginRequest) ProtoMessage()    {}
func (*AuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{5}
}

func (m *AuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthLoginRequest.Unmarshal(m, b)
}
func (m *AuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthLoginRequest.Marshal(b, m, deterministic)
}
func (m *AuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthLoginRequest.Merge(m, src)
}
func (m *AuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_AuthLoginRequest.Size(m)
}
func (m *AuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthLoginRequest proto.InternalMessageInfo

func (m *AuthLoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthLoginRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthLoginResponse struct {
	Token                string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Roles                []*AuthRole `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	ForcePwdChange       bool        `protobuf:"varint,4,opt,name=force_pwd_change,json=forcePwdChange,proto3" json:"force_pwd_change,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AuthLoginResponse) Reset()         { *m = AuthLoginResponse{} }
func (m *AuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*AuthLoginResponse) ProtoMessage()    {}
func (*AuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeac6748badcf6d, []int{6}
}

func (m *AuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthLoginResponse.Unmarshal(m, b)
}
func (m *AuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthLoginResponse.Marshal(b, m, deterministic)
}
func (m *AuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthLoginResponse.Merge(m, src)
}
func (m *AuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_AuthLoginResponse.Size(m)
}
func (m *AuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthLoginResponse proto.InternalMessageInfo

func (m *AuthLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthLoginResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AuthLoginResponse) GetRoles() []*AuthRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *AuthLoginResponse) GetForcePwdChange() bool {
	if m != nil {
		return m.ForcePwdChange
	}
	return false
}

func init() {
	proto.RegisterType((*RequestHeader)(nil), "pbx.RequestHeader")
	proto.RegisterType((*AdminUser)(nil), "pbx.AdminUser")
	proto.RegisterType((*AdminUserFindRequest)(nil), "pbx.AdminUserFindRequest")
	proto.RegisterType((*AdminUserFindResponse)(nil), "pbx.AdminUserFindResponse")
	proto.RegisterType((*AuthRole)(nil), "pbx.AuthRole")
	proto.RegisterType((*AuthLoginRequest)(nil), "pbx.AuthLoginRequest")
	proto.RegisterType((*AuthLoginResponse)(nil), "pbx.AuthLoginResponse")
}

func init() { proto.RegisterFile("pbx/iam.proto", fileDescriptor_dfeac6748badcf6d) }

var fileDescriptor_dfeac6748badcf6d = []byte{
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xfa, 0xb5, 0xf6, 0x76, 0xed, 0x8a, 0xb5, 0x0d, 0x53, 0x09, 0x51, 0x75, 0x42, 0x54,
	0x3c, 0x64, 0x52, 0x79, 0x82, 0xb7, 0x6e, 0x30, 0x98, 0xd8, 0x03, 0x0a, 0xf0, 0x1c, 0xb9, 0x89,
	0xd3, 0x5a, 0x4b, 0xec, 0x10, 0x3b, 0x6b, 0xfb, 0xca, 0x3f, 0xe0, 0x85, 0x7f, 0xc3, 0x7f, 0x43,
	0xb6, 0x93, 0x2c, 0x1b, 0x13, 0x7b, 0xeb, 0x39, 0xf7, 0xfa, 0x7e, 0x9c, 0x9c, 0x5b, 0x18, 0xa4,
	0xcb, 0xed, 0x29, 0x23, 0x89, 0x9b, 0x66, 0x42, 0x09, 0xd4, 0x4c, 0x97, 0xdb, 0xf1, 0x8b, 0x95,
	0x10, 0xab, 0x98, 0x9e, 0x1a, 0x6a, 0x99, 0x47, 0xa7, 0x8a, 0x25, 0x54, 0x2a, 0x92, 0xa4, 0x36,
	0x6b, 0xba, 0x84, 0x81, 0x47, 0x7f, 0xe4, 0x54, 0xaa, 0x4f, 0x94, 0x84, 0x34, 0x43, 0x63, 0xe8,
	0xa6, 0x31, 0x51, 0x91, 0xc8, 0x12, 0xec, 0x4c, 0x9c, 0x59, 0xcf, 0xab, 0x30, 0x3a, 0x86, 0x4e,
	0x28, 0x12, 0xc2, 0x38, 0x6e, 0x98, 0x48, 0x81, 0xd0, 0x73, 0x80, 0x20, 0x66, 0x94, 0x2b, 0xff,
	0x9a, 0xee, 0x70, 0xd3, 0xc4, 0x7a, 0x96, 0xf9, 0x4c, 0x77, 0xd3, 0x3f, 0x2d, 0xe8, 0x2d, 0xc2,
	0x84, 0xf1, 0xef, 0x92, 0x66, 0x68, 0x08, 0x0d, 0x16, 0x16, 0xa5, 0x1b, 0x2c, 0xd4, 0x0d, 0x73,
	0x49, 0x33, 0x4e, 0x12, 0x5a, 0x94, 0xad, 0xb0, 0x19, 0x86, 0x48, 0xb9, 0x11, 0x59, 0x58, 0x94,
	0xad, 0xb0, 0x8e, 0x45, 0x79, 0x1c, 0x9b, 0x77, 0x2d, 0x1b, 0x2b, 0x31, 0x3a, 0x84, 0x36, 0x4d,
	0x08, 0x8b, 0x71, 0xdb, 0x04, 0x2c, 0x40, 0x2f, 0x61, 0x48, 0x82, 0x40, 0xe4, 0x5c, 0xf9, 0x24,
	0x08, 0xa8, 0x94, 0xb8, 0x33, 0x71, 0x66, 0x5d, 0x6f, 0x50, 0xb0, 0x0b, 0x43, 0xa2, 0x13, 0x18,
	0xc4, 0x62, 0xc5, 0xb8, 0x6f, 0x48, 0x9a, 0xe1, 0xbd, 0x89, 0x33, 0x6b, 0x7b, 0xfb, 0x86, 0x3c,
	0xb7, 0x1c, 0x7a, 0x0b, 0x10, 0x13, 0xa9, 0x7c, 0x43, 0xe2, 0xee, 0xc4, 0x99, 0xf5, 0xe7, 0x63,
	0xd7, 0xaa, 0xed, 0x96, 0x6a, 0xbb, 0xdf, 0x4a, 0xb5, 0xbd, 0x9e, 0xce, 0xbe, 0xd2, 0xc9, 0x68,
	0x71, 0x3b, 0x06, 0xdd, 0xa6, 0x2c, 0xdb, 0xe1, 0xde, 0xa3, 0xcf, 0xcb, 0x11, 0x3f, 0x98, 0x07,
	0x7a, 0xbf, 0x74, 0x2d, 0x94, 0xc0, 0x60, 0xf7, 0x33, 0x40, 0x7f, 0x86, 0x74, 0x13, 0x96, 0x45,
	0xfb, 0x66, 0xb7, 0x5e, 0xba, 0x09, 0x8b, 0x47, 0x67, 0x70, 0x70, 0x1b, 0xf6, 0xb5, 0x11, 0xf0,
	0xfe, 0xe3, 0x8d, 0xab, 0xf7, 0x9a, 0x43, 0xaf, 0x60, 0xa4, 0x6b, 0xc4, 0x2c, 0xa2, 0x3e, 0xe3,
	0x7e, 0x48, 0x76, 0x12, 0x0f, 0x8c, 0x3c, 0x3a, 0xf1, 0x8a, 0x45, 0xf4, 0x92, 0xbf, 0x27, 0x3b,
	0x89, 0x66, 0x30, 0x8a, 0x44, 0x16, 0x50, 0x5f, 0xa7, 0x07, 0x6b, 0xc2, 0x57, 0x14, 0x0f, 0xcd,
	0x44, 0x43, 0xc3, 0x7f, 0xd9, 0x84, 0xe7, 0x86, 0x45, 0x13, 0xe8, 0x33, 0x2e, 0x15, 0x53, 0xb9,
	0x62, 0x82, 0xe3, 0x03, 0xb3, 0x51, 0x9d, 0x9a, 0x5e, 0xc0, 0x61, 0x65, 0x9f, 0x0b, 0xc6, 0xc3,
	0xc2, 0xb0, 0xda, 0x8e, 0x11, 0x8b, 0xf5, 0x17, 0xb2, 0x6e, 0x2a, 0x90, 0xe6, 0xad, 0xf9, 0x4a,
	0x9b, 0x5a, 0x34, 0xfd, 0x08, 0x47, 0xf7, 0xea, 0xc8, 0x54, 0x70, 0x49, 0x91, 0x0b, 0x40, 0xca,
	0x80, 0xc4, 0xce, 0xa4, 0x39, 0xeb, 0xcf, 0x87, 0x6e, 0xba, 0xdc, 0xba, 0x55, 0xbe, 0x57, 0xcb,
	0x98, 0xfe, 0x74, 0xa0, 0xbb, 0xc8, 0xd5, 0xda, 0x13, 0x31, 0xfd, 0xc7, 0xcf, 0x08, 0x5a, 0x35,
	0x2f, 0x9b, 0xdf, 0x7a, 0xc7, 0x90, 0xca, 0x20, 0x63, 0xa9, 0xd9, 0xd1, 0x5a, 0xb9, 0x4e, 0x21,
	0x0c, 0x7b, 0x21, 0x8d, 0x48, 0x1e, 0x2b, 0x63, 0xe6, 0xae, 0x57, 0x42, 0xbd, 0x8d, 0x54, 0x44,
	0xe5, 0xd2, 0x98, 0xb9, 0xeb, 0x15, 0x68, 0x7a, 0x03, 0x23, 0x3d, 0x83, 0xf1, 0x54, 0xa9, 0x48,
	0xfd, 0x96, 0x9c, 0xff, 0xdc, 0x52, 0xe3, 0xde, 0x2d, 0xbd, 0x86, 0xce, 0xda, 0x9c, 0xbf, 0x19,
	0xad, 0x3f, 0x47, 0x66, 0xf9, 0x3b, 0x7f, 0x0c, 0x5e, 0x91, 0x31, 0xfd, 0xe5, 0xc0, 0x93, 0x5a,
	0xe3, 0x42, 0xc2, 0x43, 0x68, 0x2b, 0x71, 0x4d, 0x79, 0xd1, 0xd6, 0x02, 0xf4, 0x14, 0xf6, 0x74,
	0x7f, 0x9f, 0x95, 0x2d, 0x3b, 0x1a, 0x5e, 0x86, 0xe8, 0x04, 0xda, 0x99, 0x88, 0xa9, 0xc4, 0x4d,
	0x23, 0xf6, 0xc0, 0x8a, 0x5d, 0x48, 0xea, 0xd9, 0xd8, 0x83, 0x1e, 0x6a, 0x3d, 0xe4, 0xa1, 0xf9,
	0x6f, 0x07, 0x46, 0x67, 0x94, 0x04, 0x82, 0x5f, 0x92, 0xe4, 0x2b, 0xcd, 0x6e, 0x58, 0x40, 0xd1,
	0x3b, 0xe8, 0x55, 0x73, 0xa2, 0xa3, 0xaa, 0x43, 0x5d, 0xb0, 0xf1, 0xf1, 0x7d, 0xba, 0x58, 0xe7,
	0x02, 0x06, 0x77, 0xac, 0x82, 0x9e, 0xdd, 0xb5, 0x43, 0xcd, 0x86, 0xe3, 0xf1, 0x43, 0x21, 0x5b,
	0x67, 0xd9, 0x31, 0x27, 0xf5, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x5c, 0xb1, 0xcf,
	0x9c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconIamServiceClient is the client API for BeaconIamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconIamServiceClient interface {
	// AuthLogin
	AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error)
	// AdminUser
	AdminUserFind(ctx context.Context, in *AdminUserFindRequest, opts ...grpc.CallOption) (*AdminUserFindResponse, error)
}

type beaconIamServiceClient struct {
	cc *grpc.ClientConn
}

func NewBeaconIamServiceClient(cc *grpc.ClientConn) BeaconIamServiceClient {
	return &beaconIamServiceClient{cc}
}

func (c *beaconIamServiceClient) AuthLogin(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error) {
	out := new(AuthLoginResponse)
	err := c.cc.Invoke(ctx, "/pbx.BeaconIamService/AuthLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconIamServiceClient) AdminUserFind(ctx context.Context, in *AdminUserFindRequest, opts ...grpc.CallOption) (*AdminUserFindResponse, error) {
	out := new(AdminUserFindResponse)
	err := c.cc.Invoke(ctx, "/pbx.BeaconIamService/AdminUserFind", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconIamServiceServer is the server API for BeaconIamService service.
type BeaconIamServiceServer interface {
	// AuthLogin
	AuthLogin(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error)
	// AdminUser
	AdminUserFind(context.Context, *AdminUserFindRequest) (*AdminUserFindResponse, error)
}

func RegisterBeaconIamServiceServer(s *grpc.Server, srv BeaconIamServiceServer) {
	s.RegisterService(&_BeaconIamService_serviceDesc, srv)
}

func _BeaconIamService_AuthLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconIamServiceServer).AuthLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbx.BeaconIamService/AuthLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconIamServiceServer).AuthLogin(ctx, req.(*AuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconIamService_AdminUserFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUserFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconIamServiceServer).AdminUserFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbx.BeaconIamService/AdminUserFind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconIamServiceServer).AdminUserFind(ctx, req.(*AdminUserFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconIamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbx.BeaconIamService",
	HandlerType: (*BeaconIamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthLogin",
			Handler:    _BeaconIamService_AuthLogin_Handler,
		},
		{
			MethodName: "AdminUserFind",
			Handler:    _BeaconIamService_AdminUserFind_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbx/iam.proto",
}
