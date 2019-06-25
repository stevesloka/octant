/*
Copyright (c) 2019 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dashboard.proto

package dashboard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RegisterRequest struct {
	DashboardAPIAddress  string   `protobuf:"bytes,1,opt,name=dashboardAPIAddress,proto3" json:"dashboardAPIAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{1}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetDashboardAPIAddress() string {
	if m != nil {
		return m.DashboardAPIAddress
	}
	return ""
}

type RegisterResponse struct {
	PluginName           string                         `protobuf:"bytes,1,opt,name=pluginName,proto3" json:"pluginName,omitempty"`
	Description          string                         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Capabilities         *RegisterResponse_Capabilities `protobuf:"bytes,3,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{2}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetPluginName() string {
	if m != nil {
		return m.PluginName
	}
	return ""
}

func (m *RegisterResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegisterResponse) GetCapabilities() *RegisterResponse_Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

type RegisterResponse_GroupVersionKind struct {
	Group                string   `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Kind                 string   `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse_GroupVersionKind) Reset()         { *m = RegisterResponse_GroupVersionKind{} }
func (m *RegisterResponse_GroupVersionKind) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse_GroupVersionKind) ProtoMessage()    {}
func (*RegisterResponse_GroupVersionKind) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{2, 0}
}
func (m *RegisterResponse_GroupVersionKind) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse_GroupVersionKind.Unmarshal(m, b)
}
func (m *RegisterResponse_GroupVersionKind) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse_GroupVersionKind.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse_GroupVersionKind) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse_GroupVersionKind.Merge(dst, src)
}
func (m *RegisterResponse_GroupVersionKind) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse_GroupVersionKind.Size(m)
}
func (m *RegisterResponse_GroupVersionKind) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse_GroupVersionKind.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse_GroupVersionKind proto.InternalMessageInfo

func (m *RegisterResponse_GroupVersionKind) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RegisterResponse_GroupVersionKind) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegisterResponse_GroupVersionKind) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type RegisterResponse_Capabilities struct {
	SupportsPrinterConfig []*RegisterResponse_GroupVersionKind `protobuf:"bytes,1,rep,name=supportsPrinterConfig,proto3" json:"supportsPrinterConfig,omitempty"`
	SupportsPrinterStatus []*RegisterResponse_GroupVersionKind `protobuf:"bytes,2,rep,name=supportsPrinterStatus,proto3" json:"supportsPrinterStatus,omitempty"`
	SupportsPrinterItems  []*RegisterResponse_GroupVersionKind `protobuf:"bytes,3,rep,name=supportsPrinterItems,proto3" json:"supportsPrinterItems,omitempty"`
	SupportsObjectStatus  []*RegisterResponse_GroupVersionKind `protobuf:"bytes,4,rep,name=supportsObjectStatus,proto3" json:"supportsObjectStatus,omitempty"`
	SupportsTab           []*RegisterResponse_GroupVersionKind `protobuf:"bytes,5,rep,name=supportsTab,proto3" json:"supportsTab,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *RegisterResponse_Capabilities) Reset()         { *m = RegisterResponse_Capabilities{} }
func (m *RegisterResponse_Capabilities) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse_Capabilities) ProtoMessage()    {}
func (*RegisterResponse_Capabilities) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{2, 1}
}
func (m *RegisterResponse_Capabilities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse_Capabilities.Unmarshal(m, b)
}
func (m *RegisterResponse_Capabilities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse_Capabilities.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse_Capabilities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse_Capabilities.Merge(dst, src)
}
func (m *RegisterResponse_Capabilities) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse_Capabilities.Size(m)
}
func (m *RegisterResponse_Capabilities) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse_Capabilities.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse_Capabilities proto.InternalMessageInfo

func (m *RegisterResponse_Capabilities) GetSupportsPrinterConfig() []*RegisterResponse_GroupVersionKind {
	if m != nil {
		return m.SupportsPrinterConfig
	}
	return nil
}

func (m *RegisterResponse_Capabilities) GetSupportsPrinterStatus() []*RegisterResponse_GroupVersionKind {
	if m != nil {
		return m.SupportsPrinterStatus
	}
	return nil
}

func (m *RegisterResponse_Capabilities) GetSupportsPrinterItems() []*RegisterResponse_GroupVersionKind {
	if m != nil {
		return m.SupportsPrinterItems
	}
	return nil
}

func (m *RegisterResponse_Capabilities) GetSupportsObjectStatus() []*RegisterResponse_GroupVersionKind {
	if m != nil {
		return m.SupportsObjectStatus
	}
	return nil
}

func (m *RegisterResponse_Capabilities) GetSupportsTab() []*RegisterResponse_GroupVersionKind {
	if m != nil {
		return m.SupportsTab
	}
	return nil
}

type ObjectRequest struct {
	Object               []byte   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectRequest) Reset()         { *m = ObjectRequest{} }
func (m *ObjectRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectRequest) ProtoMessage()    {}
func (*ObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{3}
}
func (m *ObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectRequest.Unmarshal(m, b)
}
func (m *ObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectRequest.Marshal(b, m, deterministic)
}
func (dst *ObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectRequest.Merge(dst, src)
}
func (m *ObjectRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectRequest.Size(m)
}
func (m *ObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectRequest proto.InternalMessageInfo

func (m *ObjectRequest) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

type PrintResponse struct {
	Config               []*PrintResponse_SummaryItem `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	Status               []*PrintResponse_SummaryItem `protobuf:"bytes,2,rep,name=status,proto3" json:"status,omitempty"`
	Items                []byte                       `protobuf:"bytes,3,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PrintResponse) Reset()         { *m = PrintResponse{} }
func (m *PrintResponse) String() string { return proto.CompactTextString(m) }
func (*PrintResponse) ProtoMessage()    {}
func (*PrintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{4}
}
func (m *PrintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintResponse.Unmarshal(m, b)
}
func (m *PrintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintResponse.Marshal(b, m, deterministic)
}
func (dst *PrintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintResponse.Merge(dst, src)
}
func (m *PrintResponse) XXX_Size() int {
	return xxx_messageInfo_PrintResponse.Size(m)
}
func (m *PrintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintResponse proto.InternalMessageInfo

func (m *PrintResponse) GetConfig() []*PrintResponse_SummaryItem {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PrintResponse) GetStatus() []*PrintResponse_SummaryItem {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PrintResponse) GetItems() []byte {
	if m != nil {
		return m.Items
	}
	return nil
}

type PrintResponse_SummaryItem struct {
	Header               string   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Component            []byte   `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintResponse_SummaryItem) Reset()         { *m = PrintResponse_SummaryItem{} }
func (m *PrintResponse_SummaryItem) String() string { return proto.CompactTextString(m) }
func (*PrintResponse_SummaryItem) ProtoMessage()    {}
func (*PrintResponse_SummaryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{4, 0}
}
func (m *PrintResponse_SummaryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintResponse_SummaryItem.Unmarshal(m, b)
}
func (m *PrintResponse_SummaryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintResponse_SummaryItem.Marshal(b, m, deterministic)
}
func (dst *PrintResponse_SummaryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintResponse_SummaryItem.Merge(dst, src)
}
func (m *PrintResponse_SummaryItem) XXX_Size() int {
	return xxx_messageInfo_PrintResponse_SummaryItem.Size(m)
}
func (m *PrintResponse_SummaryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintResponse_SummaryItem.DiscardUnknown(m)
}

var xxx_messageInfo_PrintResponse_SummaryItem proto.InternalMessageInfo

func (m *PrintResponse_SummaryItem) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *PrintResponse_SummaryItem) GetComponent() []byte {
	if m != nil {
		return m.Component
	}
	return nil
}

type PrintTabResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layout               []byte   `protobuf:"bytes,2,opt,name=layout,proto3" json:"layout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrintTabResponse) Reset()         { *m = PrintTabResponse{} }
func (m *PrintTabResponse) String() string { return proto.CompactTextString(m) }
func (*PrintTabResponse) ProtoMessage()    {}
func (*PrintTabResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{5}
}
func (m *PrintTabResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrintTabResponse.Unmarshal(m, b)
}
func (m *PrintTabResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrintTabResponse.Marshal(b, m, deterministic)
}
func (dst *PrintTabResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrintTabResponse.Merge(dst, src)
}
func (m *PrintTabResponse) XXX_Size() int {
	return xxx_messageInfo_PrintTabResponse.Size(m)
}
func (m *PrintTabResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrintTabResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrintTabResponse proto.InternalMessageInfo

func (m *PrintTabResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrintTabResponse) GetLayout() []byte {
	if m != nil {
		return m.Layout
	}
	return nil
}

type ObjectStatusResponse struct {
	ObjectStatus         []byte   `protobuf:"bytes,1,opt,name=objectStatus,proto3" json:"objectStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectStatusResponse) Reset()         { *m = ObjectStatusResponse{} }
func (m *ObjectStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ObjectStatusResponse) ProtoMessage()    {}
func (*ObjectStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{6}
}
func (m *ObjectStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectStatusResponse.Unmarshal(m, b)
}
func (m *ObjectStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectStatusResponse.Marshal(b, m, deterministic)
}
func (dst *ObjectStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectStatusResponse.Merge(dst, src)
}
func (m *ObjectStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ObjectStatusResponse.Size(m)
}
func (m *ObjectStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectStatusResponse proto.InternalMessageInfo

func (m *ObjectStatusResponse) GetObjectStatus() []byte {
	if m != nil {
		return m.ObjectStatus
	}
	return nil
}

type WatchRequest struct {
	WatchID              string   `protobuf:"bytes,1,opt,name=watchID,proto3" json:"watchID,omitempty"`
	Object               []byte   `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dashboard_bf0c1b655ae30a4d, []int{7}
}
func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (dst *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(dst, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetWatchID() string {
	if m != nil {
		return m.WatchID
	}
	return ""
}

func (m *WatchRequest) GetObject() []byte {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "dashboard.Empty")
	proto.RegisterType((*RegisterRequest)(nil), "dashboard.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "dashboard.RegisterResponse")
	proto.RegisterType((*RegisterResponse_GroupVersionKind)(nil), "dashboard.RegisterResponse.GroupVersionKind")
	proto.RegisterType((*RegisterResponse_Capabilities)(nil), "dashboard.RegisterResponse.Capabilities")
	proto.RegisterType((*ObjectRequest)(nil), "dashboard.ObjectRequest")
	proto.RegisterType((*PrintResponse)(nil), "dashboard.PrintResponse")
	proto.RegisterType((*PrintResponse_SummaryItem)(nil), "dashboard.PrintResponse.SummaryItem")
	proto.RegisterType((*PrintTabResponse)(nil), "dashboard.PrintTabResponse")
	proto.RegisterType((*ObjectStatusResponse)(nil), "dashboard.ObjectStatusResponse")
	proto.RegisterType((*WatchRequest)(nil), "dashboard.WatchRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Print(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*PrintResponse, error)
	ObjectStatus(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectStatusResponse, error)
	PrintTab(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*PrintTabResponse, error)
	WatchAdd(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error)
	WatchUpdate(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error)
	WatchDelete(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Print(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*PrintResponse, error) {
	out := new(PrintResponse)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/Print", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ObjectStatus(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*ObjectStatusResponse, error) {
	out := new(ObjectStatusResponse)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/ObjectStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) PrintTab(ctx context.Context, in *ObjectRequest, opts ...grpc.CallOption) (*PrintTabResponse, error) {
	out := new(PrintTabResponse)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/PrintTab", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) WatchAdd(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/WatchAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) WatchUpdate(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/WatchUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) WatchDelete(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/dashboard.Plugin/WatchDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Print(context.Context, *ObjectRequest) (*PrintResponse, error)
	ObjectStatus(context.Context, *ObjectRequest) (*ObjectStatusResponse, error)
	PrintTab(context.Context, *ObjectRequest) (*PrintTabResponse, error)
	WatchAdd(context.Context, *WatchRequest) (*Empty, error)
	WatchUpdate(context.Context, *WatchRequest) (*Empty, error)
	WatchDelete(context.Context, *WatchRequest) (*Empty, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Print_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Print(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/Print",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Print(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ObjectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ObjectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/ObjectStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ObjectStatus(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_PrintTab_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).PrintTab(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/PrintTab",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).PrintTab(ctx, req.(*ObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_WatchAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).WatchAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/WatchAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).WatchAdd(ctx, req.(*WatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_WatchUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).WatchUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/WatchUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).WatchUpdate(ctx, req.(*WatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_WatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).WatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dashboard.Plugin/WatchDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).WatchDelete(ctx, req.(*WatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Plugin_Register_Handler,
		},
		{
			MethodName: "Print",
			Handler:    _Plugin_Print_Handler,
		},
		{
			MethodName: "ObjectStatus",
			Handler:    _Plugin_ObjectStatus_Handler,
		},
		{
			MethodName: "PrintTab",
			Handler:    _Plugin_PrintTab_Handler,
		},
		{
			MethodName: "WatchAdd",
			Handler:    _Plugin_WatchAdd_Handler,
		},
		{
			MethodName: "WatchUpdate",
			Handler:    _Plugin_WatchUpdate_Handler,
		},
		{
			MethodName: "WatchDelete",
			Handler:    _Plugin_WatchDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}

func init() { proto.RegisterFile("dashboard.proto", fileDescriptor_dashboard_bf0c1b655ae30a4d) }

var fileDescriptor_dashboard_bf0c1b655ae30a4d = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x6d, 0x8b, 0xd3, 0x40,
	0x10, 0xa6, 0xd7, 0xeb, 0xdb, 0x24, 0xc7, 0x95, 0xb5, 0x6a, 0xc8, 0x89, 0x96, 0x20, 0xd8, 0x0f,
	0x52, 0xa4, 0x22, 0x88, 0x8a, 0x58, 0x7a, 0x22, 0x45, 0x39, 0x4b, 0xee, 0x3c, 0xbf, 0xba, 0x49,
	0xd6, 0x76, 0xb5, 0xc9, 0xc6, 0xdd, 0x8d, 0xd2, 0xdf, 0xe7, 0x7f, 0x51, 0xf0, 0x57, 0x48, 0x36,
	0x9b, 0x76, 0x5b, 0x7b, 0xc5, 0x3b, 0xfd, 0xd6, 0x99, 0x9d, 0x79, 0x9e, 0x99, 0x79, 0x66, 0x1a,
	0x38, 0x8c, 0xb0, 0x98, 0x05, 0x0c, 0xf3, 0xa8, 0x9f, 0x72, 0x26, 0x19, 0x6a, 0x2d, 0x1d, 0x5e,
	0x03, 0x6a, 0x2f, 0xe3, 0x54, 0x2e, 0xbc, 0x11, 0x1c, 0xfa, 0x64, 0x4a, 0x85, 0x24, 0xdc, 0x27,
	0x5f, 0x32, 0x22, 0x24, 0x7a, 0x00, 0xd7, 0x96, 0x81, 0xc3, 0xc9, 0x78, 0x18, 0x45, 0x9c, 0x08,
	0xe1, 0x54, 0xba, 0x95, 0x5e, 0xcb, 0xdf, 0xf6, 0xe4, 0xfd, 0xac, 0x41, 0x7b, 0x85, 0x22, 0x52,
	0x96, 0x08, 0x82, 0x6e, 0x03, 0xa4, 0xf3, 0x6c, 0x4a, 0x93, 0x13, 0x1c, 0x13, 0x9d, 0x6d, 0x78,
	0x50, 0x17, 0xac, 0x88, 0x88, 0x90, 0xd3, 0x54, 0x52, 0x96, 0x38, 0x7b, 0x2a, 0xc0, 0x74, 0xa1,
	0x37, 0x60, 0x87, 0x38, 0xc5, 0x01, 0x9d, 0x53, 0x49, 0x89, 0x70, 0xaa, 0xdd, 0x4a, 0xcf, 0x1a,
	0xf4, 0xfa, 0xab, 0xbe, 0x36, 0x49, 0xfb, 0x23, 0x23, 0xde, 0x5f, 0xcb, 0x76, 0xcf, 0xa1, 0xfd,
	0x8a, 0xb3, 0x2c, 0x3d, 0x27, 0x5c, 0x50, 0x96, 0xbc, 0xa6, 0x49, 0x84, 0x3a, 0x50, 0x9b, 0xe6,
	0x3e, 0x5d, 0x5e, 0x61, 0x20, 0x07, 0x1a, 0x5f, 0x8b, 0x20, 0x5d, 0x55, 0x69, 0x22, 0x04, 0xfb,
	0x9f, 0x69, 0x12, 0xa9, 0x4a, 0x5a, 0xbe, 0xfa, 0xed, 0xfe, 0xaa, 0x82, 0x6d, 0xd2, 0xa2, 0x00,
	0xae, 0x8b, 0x2c, 0x4d, 0x19, 0x97, 0x62, 0xc2, 0x69, 0x22, 0x09, 0x1f, 0xb1, 0xe4, 0x23, 0x9d,
	0x3a, 0x95, 0x6e, 0xb5, 0x67, 0x0d, 0xee, 0xef, 0xaa, 0x7f, 0xb3, 0x42, 0x7f, 0x3b, 0xd4, 0x16,
	0x8e, 0x53, 0x89, 0x65, 0x26, 0x9c, 0xbd, 0xff, 0xc0, 0x51, 0x40, 0xa1, 0x0f, 0xd0, 0xd9, 0x78,
	0x18, 0x4b, 0x12, 0xe7, 0x32, 0x5c, 0x9e, 0x62, 0x2b, 0x92, 0xc9, 0xf0, 0x36, 0xf8, 0x44, 0x42,
	0xa9, 0x9b, 0xd8, 0xff, 0x17, 0x06, 0x13, 0x09, 0x9d, 0x80, 0x55, 0xfa, 0xcf, 0x70, 0xe0, 0xd4,
	0xae, 0x00, 0x6c, 0x02, 0x78, 0xf7, 0xe0, 0xa0, 0xc0, 0x2f, 0x8f, 0xe5, 0x06, 0xd4, 0x99, 0x72,
	0xa8, 0x15, 0xb2, 0x7d, 0x6d, 0x79, 0x3f, 0x2a, 0x70, 0xa0, 0x7a, 0x5d, 0xde, 0xc3, 0x33, 0xa8,
	0x87, 0xe6, 0x1e, 0xdc, 0x35, 0xaa, 0x58, 0x8b, 0xec, 0x9f, 0x66, 0x71, 0x8c, 0xf9, 0x22, 0x9f,
	0x91, 0xaf, 0x73, 0xf2, 0x6c, 0x61, 0x2a, 0xfc, 0x97, 0xd9, 0x45, 0x4e, 0xbe, 0xe7, 0x54, 0x6b,
	0x97, 0x17, 0x59, 0x18, 0xee, 0x08, 0x2c, 0x23, 0x38, 0x6f, 0x65, 0x46, 0x70, 0x44, 0xb8, 0xbe,
	0x06, 0x6d, 0xa1, 0x5b, 0xd0, 0x0a, 0x59, 0x9c, 0xb2, 0x84, 0x24, 0x52, 0x1d, 0x84, 0xed, 0xaf,
	0x1c, 0xde, 0x73, 0x68, 0x2b, 0xfe, 0x33, 0x1c, 0x2c, 0x5b, 0x45, 0xb0, 0x9f, 0xac, 0x8e, 0x5e,
	0xfd, 0xce, 0xd1, 0xe7, 0x78, 0xc1, 0xb2, 0x12, 0x42, 0x5b, 0xde, 0x13, 0xe8, 0x98, 0x8a, 0x2d,
	0x31, 0x3c, 0xb0, 0x99, 0xb9, 0x13, 0xc5, 0x78, 0xd7, 0x7c, 0xde, 0x0b, 0xb0, 0xdf, 0x63, 0x19,
	0xce, 0x4a, 0x31, 0x1c, 0x68, 0x7c, 0xcb, 0xed, 0xf1, 0xb1, 0xa6, 0x2e, 0x4d, 0x43, 0xa6, 0x3d,
	0x53, 0xa6, 0xc1, 0xf7, 0x2a, 0xd4, 0x27, 0xea, 0x3f, 0x09, 0x8d, 0xa0, 0x59, 0x2e, 0x03, 0x72,
	0xb7, 0x6e, 0x88, 0x22, 0x71, 0x8f, 0x76, 0x6c, 0x0f, 0x7a, 0x0a, 0x35, 0x35, 0x0d, 0xe4, 0x18,
	0x51, 0x6b, 0x1b, 0xe3, 0x3a, 0x17, 0x29, 0x87, 0xc6, 0x60, 0xaf, 0x2d, 0xef, 0xc5, 0x18, 0x77,
	0xfe, 0x78, 0xd9, 0x98, 0xde, 0x10, 0x9a, 0xa5, 0x2a, 0x3b, 0x60, 0x8e, 0x36, 0x4b, 0x31, 0x45,
	0x7c, 0x04, 0x4d, 0x35, 0xdc, 0x61, 0x14, 0xa1, 0x9b, 0x46, 0xa0, 0x39, 0x71, 0xb7, 0x6d, 0x3c,
	0xa8, 0x0f, 0x0a, 0x7a, 0x0c, 0x96, 0x8a, 0x78, 0x97, 0x46, 0x58, 0x92, 0xab, 0x64, 0x1e, 0x93,
	0x39, 0xb9, 0x54, 0x66, 0x50, 0x57, 0xdf, 0xb7, 0x87, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfe,
	0xd6, 0x2a, 0xad, 0xf2, 0x06, 0x00, 0x00,
}