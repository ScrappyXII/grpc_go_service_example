// Code generated by protoc-gen-go. DO NOT EDIT.
// source: organization.proto

package mbm_service_organization

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CreateOrganizationRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrganizationRequest) Reset()         { *m = CreateOrganizationRequest{} }
func (m *CreateOrganizationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrganizationRequest) ProtoMessage()    {}
func (*CreateOrganizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{1}
}

func (m *CreateOrganizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrganizationRequest.Unmarshal(m, b)
}
func (m *CreateOrganizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrganizationRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrganizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrganizationRequest.Merge(m, src)
}
func (m *CreateOrganizationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrganizationRequest.Size(m)
}
func (m *CreateOrganizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrganizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrganizationRequest proto.InternalMessageInfo

func (m *CreateOrganizationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateOrganizationRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type OrganizationResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizationResponse) Reset()         { *m = OrganizationResponse{} }
func (m *OrganizationResponse) String() string { return proto.CompactTextString(m) }
func (*OrganizationResponse) ProtoMessage()    {}
func (*OrganizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{2}
}

func (m *OrganizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationResponse.Unmarshal(m, b)
}
func (m *OrganizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationResponse.Marshal(b, m, deterministic)
}
func (m *OrganizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationResponse.Merge(m, src)
}
func (m *OrganizationResponse) XXX_Size() int {
	return xxx_messageInfo_OrganizationResponse.Size(m)
}
func (m *OrganizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationResponse proto.InternalMessageInfo

func (m *OrganizationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OrganizationResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type OrganizationListResponse struct {
	Organizations        []*OrganizationResponse `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OrganizationListResponse) Reset()         { *m = OrganizationListResponse{} }
func (m *OrganizationListResponse) String() string { return proto.CompactTextString(m) }
func (*OrganizationListResponse) ProtoMessage()    {}
func (*OrganizationListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{3}
}

func (m *OrganizationListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationListResponse.Unmarshal(m, b)
}
func (m *OrganizationListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationListResponse.Marshal(b, m, deterministic)
}
func (m *OrganizationListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationListResponse.Merge(m, src)
}
func (m *OrganizationListResponse) XXX_Size() int {
	return xxx_messageInfo_OrganizationListResponse.Size(m)
}
func (m *OrganizationListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationListResponse proto.InternalMessageInfo

func (m *OrganizationListResponse) GetOrganizations() []*OrganizationResponse {
	if m != nil {
		return m.Organizations
	}
	return nil
}

type CreateUserRequest struct {
	OrganizationId       string   `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *CreateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId       string   `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{5}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserListResponse struct {
	Users                []*UserResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{6}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUsers() []*UserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

type ByOrganizationRequest struct {
	OrganizationId       string   `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByOrganizationRequest) Reset()         { *m = ByOrganizationRequest{} }
func (m *ByOrganizationRequest) String() string { return proto.CompactTextString(m) }
func (*ByOrganizationRequest) ProtoMessage()    {}
func (*ByOrganizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d10c68ef159b9ed, []int{7}
}

func (m *ByOrganizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByOrganizationRequest.Unmarshal(m, b)
}
func (m *ByOrganizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByOrganizationRequest.Marshal(b, m, deterministic)
}
func (m *ByOrganizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByOrganizationRequest.Merge(m, src)
}
func (m *ByOrganizationRequest) XXX_Size() int {
	return xxx_messageInfo_ByOrganizationRequest.Size(m)
}
func (m *ByOrganizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ByOrganizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ByOrganizationRequest proto.InternalMessageInfo

func (m *ByOrganizationRequest) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "mbm.service.organization.Empty")
	proto.RegisterType((*CreateOrganizationRequest)(nil), "mbm.service.organization.CreateOrganizationRequest")
	proto.RegisterType((*OrganizationResponse)(nil), "mbm.service.organization.OrganizationResponse")
	proto.RegisterType((*OrganizationListResponse)(nil), "mbm.service.organization.OrganizationListResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "mbm.service.organization.CreateUserRequest")
	proto.RegisterType((*UserResponse)(nil), "mbm.service.organization.UserResponse")
	proto.RegisterType((*UserListResponse)(nil), "mbm.service.organization.UserListResponse")
	proto.RegisterType((*ByOrganizationRequest)(nil), "mbm.service.organization.ByOrganizationRequest")
}

func init() { proto.RegisterFile("organization.proto", fileDescriptor_8d10c68ef159b9ed) }

var fileDescriptor_8d10c68ef159b9ed = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xa1, 0x80, 0xc6, 0xe1, 0x8f, 0x3a, 0x91, 0xa4, 0xe2, 0x41, 0xb2, 0x07, 0x25, 0x9a,
	0xac, 0x09, 0x5c, 0x3d, 0x18, 0x8d, 0x26, 0x26, 0x26, 0x62, 0xa3, 0x27, 0x8d, 0xa6, 0x94, 0x8d,
	0xee, 0xa1, 0x7f, 0xdc, 0x5d, 0x4c, 0x50, 0x3f, 0x9b, 0x9f, 0xcd, 0x50, 0x04, 0xb7, 0xd0, 0x85,
	0x72, 0x6b, 0x76, 0xb6, 0xbf, 0x99, 0xf7, 0xe6, 0xb5, 0x80, 0xa1, 0x78, 0x75, 0x03, 0xfe, 0xe9,
	0x2a, 0x1e, 0x06, 0x34, 0x12, 0xa1, 0x0a, 0xd1, 0xf6, 0x7b, 0x3e, 0x95, 0x4c, 0x7c, 0x70, 0x8f,
	0x51, 0xbd, 0x4e, 0xd6, 0xa1, 0x74, 0xe9, 0x47, 0x6a, 0x48, 0xee, 0x60, 0xf7, 0x42, 0x30, 0x57,
	0xb1, 0x5b, 0xad, 0xec, 0xb0, 0xf7, 0x01, 0x93, 0x0a, 0x11, 0x8a, 0x81, 0xeb, 0x33, 0x3b, 0xdf,
	0xcc, 0xb7, 0x36, 0x9c, 0xf8, 0x19, 0x9b, 0x50, 0xee, 0x33, 0xe9, 0x09, 0x1e, 0x8d, 0x6e, 0xda,
	0x56, 0x5c, 0xd2, 0x8f, 0xc8, 0x13, 0xec, 0x24, 0x61, 0x32, 0x0a, 0x03, 0xc9, 0xb0, 0x06, 0x16,
	0xef, 0xff, 0xb1, 0x2c, 0xde, 0x9f, 0xd2, 0x2d, 0x33, 0xbd, 0x30, 0x4f, 0x8f, 0xc0, 0xd6, 0xe9,
	0x37, 0x5c, 0xaa, 0x69, 0x87, 0x7b, 0xa8, 0xea, 0x2a, 0xa5, 0x9d, 0x6f, 0x16, 0x5a, 0xe5, 0x36,
	0xa5, 0x26, 0x1f, 0x68, 0xda, 0xa0, 0x4e, 0x12, 0x42, 0xba, 0xb0, 0x3d, 0xb6, 0xe8, 0x41, 0x32,
	0x31, 0xb1, 0xe6, 0x10, 0x36, 0xf5, 0x5b, 0x2f, 0x53, 0x65, 0x35, 0xfd, 0xf8, 0x3a, 0x55, 0x25,
	0x79, 0x84, 0xca, 0x98, 0x65, 0x70, 0x26, 0x05, 0x6e, 0x2d, 0x84, 0x17, 0x34, 0x78, 0x17, 0xb6,
	0x46, 0xf0, 0x84, 0x31, 0xa7, 0x50, 0x1a, 0x48, 0x26, 0x26, 0x86, 0x1c, 0x98, 0x0d, 0xd1, 0xe7,
	0x72, 0xc6, 0x2f, 0x91, 0x33, 0xa8, 0x9f, 0x0f, 0xd3, 0xf2, 0x91, 0xd5, 0x84, 0xf6, 0x4f, 0x11,
	0x2a, 0x3a, 0x00, 0xbf, 0x00, 0xe7, 0x63, 0x87, 0x1d, 0xf3, 0x5c, 0xc6, 0x90, 0x36, 0x56, 0xdc,
	0x2e, 0xc9, 0x61, 0x00, 0xf5, 0x2b, 0xa6, 0xbc, 0xb7, 0xd9, 0x1c, 0xe1, 0xbe, 0x19, 0x15, 0x7f,
	0x2d, 0x8d, 0x76, 0xb6, 0x5e, 0xba, 0xf7, 0x24, 0x87, 0x1e, 0xc0, 0x7f, 0x80, 0xf0, 0x78, 0x99,
	0x48, 0x2d, 0x66, 0x8d, 0x8c, 0x9b, 0x22, 0x39, 0x7c, 0x86, 0x6a, 0x2c, 0x6a, 0xb2, 0xfb, 0xe5,
	0x62, 0x8e, 0x16, 0xb3, 0x67, 0x44, 0x7c, 0xc3, 0x5e, 0x82, 0x9f, 0x4c, 0x04, 0x9e, 0x98, 0x61,
	0xa9, 0xd9, 0x59, 0xad, 0x7b, 0x6f, 0x2d, 0xfe, 0xa1, 0x75, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x74, 0xe1, 0x7a, 0x1f, 0xe6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrganizationClient is the client API for Organization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrganizationClient interface {
	CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationResponse, error)
	FetchOrganizationList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrganizationListResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	FetchUserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error)
	FetchUserListByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...grpc.CallOption) (*UserListResponse, error)
}

type organizationClient struct {
	cc *grpc.ClientConn
}

func NewOrganizationClient(cc *grpc.ClientConn) OrganizationClient {
	return &organizationClient{cc}
}

func (c *organizationClient) CreateOrganization(ctx context.Context, in *CreateOrganizationRequest, opts ...grpc.CallOption) (*OrganizationResponse, error) {
	out := new(OrganizationResponse)
	err := c.cc.Invoke(ctx, "/mbm.service.organization.Organization/CreateOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) FetchOrganizationList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrganizationListResponse, error) {
	out := new(OrganizationListResponse)
	err := c.cc.Invoke(ctx, "/mbm.service.organization.Organization/FetchOrganizationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/mbm.service.organization.Organization/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) FetchUserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/mbm.service.organization.Organization/FetchUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *organizationClient) FetchUserListByOrganization(ctx context.Context, in *ByOrganizationRequest, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/mbm.service.organization.Organization/FetchUserListByOrganization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrganizationServer is the server API for Organization service.
type OrganizationServer interface {
	CreateOrganization(context.Context, *CreateOrganizationRequest) (*OrganizationResponse, error)
	FetchOrganizationList(context.Context, *Empty) (*OrganizationListResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	FetchUserList(context.Context, *Empty) (*UserListResponse, error)
	FetchUserListByOrganization(context.Context, *ByOrganizationRequest) (*UserListResponse, error)
}

// UnimplementedOrganizationServer can be embedded to have forward compatible implementations.
type UnimplementedOrganizationServer struct {
}

func (*UnimplementedOrganizationServer) CreateOrganization(ctx context.Context, req *CreateOrganizationRequest) (*OrganizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrganization not implemented")
}
func (*UnimplementedOrganizationServer) FetchOrganizationList(ctx context.Context, req *Empty) (*OrganizationListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrganizationList not implemented")
}
func (*UnimplementedOrganizationServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedOrganizationServer) FetchUserList(ctx context.Context, req *Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserList not implemented")
}
func (*UnimplementedOrganizationServer) FetchUserListByOrganization(ctx context.Context, req *ByOrganizationRequest) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUserListByOrganization not implemented")
}

func RegisterOrganizationServer(s *grpc.Server, srv OrganizationServer) {
	s.RegisterService(&_Organization_serviceDesc, srv)
}

func _Organization_CreateOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).CreateOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbm.service.organization.Organization/CreateOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).CreateOrganization(ctx, req.(*CreateOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_FetchOrganizationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).FetchOrganizationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbm.service.organization.Organization/FetchOrganizationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).FetchOrganizationList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbm.service.organization.Organization/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_FetchUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).FetchUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbm.service.organization.Organization/FetchUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).FetchUserList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Organization_FetchUserListByOrganization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByOrganizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrganizationServer).FetchUserListByOrganization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbm.service.organization.Organization/FetchUserListByOrganization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrganizationServer).FetchUserListByOrganization(ctx, req.(*ByOrganizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Organization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mbm.service.organization.Organization",
	HandlerType: (*OrganizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrganization",
			Handler:    _Organization_CreateOrganization_Handler,
		},
		{
			MethodName: "FetchOrganizationList",
			Handler:    _Organization_FetchOrganizationList_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Organization_CreateUser_Handler,
		},
		{
			MethodName: "FetchUserList",
			Handler:    _Organization_FetchUserList_Handler,
		},
		{
			MethodName: "FetchUserListByOrganization",
			Handler:    _Organization_FetchUserListByOrganization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "organization.proto",
}
