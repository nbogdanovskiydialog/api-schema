// Code generated by protoc-gen-go.
// source: typing_and_online.proto
// DO NOT EDIT!

package dialog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TypingType int32

const (
	TypingType_TYPINGTYPE_UNKNOWN TypingType = 0
	TypingType_TYPINGTYPE_TEXT    TypingType = 1
)

var TypingType_name = map[int32]string{
	0: "TYPINGTYPE_UNKNOWN",
	1: "TYPINGTYPE_TEXT",
}
var TypingType_value = map[string]int32{
	"TYPINGTYPE_UNKNOWN": 0,
	"TYPINGTYPE_TEXT":    1,
}

func (x TypingType) String() string {
	return proto.EnumName(TypingType_name, int32(x))
}
func (TypingType) EnumDescriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

type DeviceType int32

const (
	DeviceType_DEVICETYPE_UNKNOWN DeviceType = 0
	DeviceType_DEVICETYPE_GENERIC DeviceType = 1
	DeviceType_DEVICETYPE_PC      DeviceType = 2
	DeviceType_DEVICETYPE_MOBILE  DeviceType = 3
	DeviceType_DEVICETYPE_TABLET  DeviceType = 4
	DeviceType_DEVICETYPE_WATCH   DeviceType = 5
	DeviceType_DEVICETYPE_MIRROR  DeviceType = 6
	DeviceType_DEVICETYPE_CAR     DeviceType = 7
	DeviceType_DEVICETYPE_TABLE   DeviceType = 8
)

var DeviceType_name = map[int32]string{
	0: "DEVICETYPE_UNKNOWN",
	1: "DEVICETYPE_GENERIC",
	2: "DEVICETYPE_PC",
	3: "DEVICETYPE_MOBILE",
	4: "DEVICETYPE_TABLET",
	5: "DEVICETYPE_WATCH",
	6: "DEVICETYPE_MIRROR",
	7: "DEVICETYPE_CAR",
	8: "DEVICETYPE_TABLE",
}
var DeviceType_value = map[string]int32{
	"DEVICETYPE_UNKNOWN": 0,
	"DEVICETYPE_GENERIC": 1,
	"DEVICETYPE_PC":      2,
	"DEVICETYPE_MOBILE":  3,
	"DEVICETYPE_TABLET":  4,
	"DEVICETYPE_WATCH":   5,
	"DEVICETYPE_MIRROR":  6,
	"DEVICETYPE_CAR":     7,
	"DEVICETYPE_TABLE":   8,
}

func (x DeviceType) String() string {
	return proto.EnumName(DeviceType_name, int32(x))
}
func (DeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

// Sending typing notification
type RequestTyping struct {
	Peer       *OutPeer   `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	TypingType TypingType `protobuf:"varint,3,opt,name=typing_type,json=typingType,enum=dialog.TypingType" json:"typing_type,omitempty"`
}

func (m *RequestTyping) Reset()                    { *m = RequestTyping{} }
func (m *RequestTyping) String() string            { return proto.CompactTextString(m) }
func (*RequestTyping) ProtoMessage()               {}
func (*RequestTyping) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{0} }

func (m *RequestTyping) GetPeer() *OutPeer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *RequestTyping) GetTypingType() TypingType {
	if m != nil {
		return m.TypingType
	}
	return TypingType_TYPINGTYPE_UNKNOWN
}

// Stop typing
type RequestStopTyping struct {
	Peer       *OutPeer   `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	TypingType TypingType `protobuf:"varint,2,opt,name=typing_type,json=typingType,enum=dialog.TypingType" json:"typing_type,omitempty"`
}

func (m *RequestStopTyping) Reset()                    { *m = RequestStopTyping{} }
func (m *RequestStopTyping) String() string            { return proto.CompactTextString(m) }
func (*RequestStopTyping) ProtoMessage()               {}
func (*RequestStopTyping) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{1} }

func (m *RequestStopTyping) GetPeer() *OutPeer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *RequestStopTyping) GetTypingType() TypingType {
	if m != nil {
		return m.TypingType
	}
	return TypingType_TYPINGTYPE_UNKNOWN
}

// Sending online state
type RequestSetOnline struct {
	IsOnline       bool                         `protobuf:"varint,1,opt,name=is_online,json=isOnline" json:"is_online,omitempty"`
	Timeout        int64                        `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	DeviceType     DeviceType                   `protobuf:"varint,3,opt,name=device_type,json=deviceType,enum=dialog.DeviceType" json:"device_type,omitempty"`
	DeviceCategory *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=device_category,json=deviceCategory" json:"device_category,omitempty"`
}

func (m *RequestSetOnline) Reset()                    { *m = RequestSetOnline{} }
func (m *RequestSetOnline) String() string            { return proto.CompactTextString(m) }
func (*RequestSetOnline) ProtoMessage()               {}
func (*RequestSetOnline) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{2} }

func (m *RequestSetOnline) GetIsOnline() bool {
	if m != nil {
		return m.IsOnline
	}
	return false
}

func (m *RequestSetOnline) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RequestSetOnline) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICETYPE_UNKNOWN
}

func (m *RequestSetOnline) GetDeviceCategory() *google_protobuf.StringValue {
	if m != nil {
		return m.DeviceCategory
	}
	return nil
}

// Update about pausing notifications
type UpdatePauseNotifications struct {
	Timeout int32 `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *UpdatePauseNotifications) Reset()                    { *m = UpdatePauseNotifications{} }
func (m *UpdatePauseNotifications) String() string            { return proto.CompactTextString(m) }
func (*UpdatePauseNotifications) ProtoMessage()               {}
func (*UpdatePauseNotifications) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{3} }

func (m *UpdatePauseNotifications) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Update about restoring notifications
type UpdateRestoreNotifications struct {
}

func (m *UpdateRestoreNotifications) Reset()                    { *m = UpdateRestoreNotifications{} }
func (m *UpdateRestoreNotifications) String() string            { return proto.CompactTextString(m) }
func (*UpdateRestoreNotifications) ProtoMessage()               {}
func (*UpdateRestoreNotifications) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{4} }

// Pause notifications
type RequestPauseNotifications struct {
	Timeout int32 `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *RequestPauseNotifications) Reset()                    { *m = RequestPauseNotifications{} }
func (m *RequestPauseNotifications) String() string            { return proto.CompactTextString(m) }
func (*RequestPauseNotifications) ProtoMessage()               {}
func (*RequestPauseNotifications) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{5} }

func (m *RequestPauseNotifications) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Restoring notifications
type RequestRestoreNotifications struct {
}

func (m *RequestRestoreNotifications) Reset()                    { *m = RequestRestoreNotifications{} }
func (m *RequestRestoreNotifications) String() string            { return proto.CompactTextString(m) }
func (*RequestRestoreNotifications) ProtoMessage()               {}
func (*RequestRestoreNotifications) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{6} }

// Update about user's typing
type UpdateTyping struct {
	Peer       *Peer      `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	Uid        int32      `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	TypingType TypingType `protobuf:"varint,3,opt,name=typing_type,json=typingType,enum=dialog.TypingType" json:"typing_type,omitempty"`
}

func (m *UpdateTyping) Reset()                    { *m = UpdateTyping{} }
func (m *UpdateTyping) String() string            { return proto.CompactTextString(m) }
func (*UpdateTyping) ProtoMessage()               {}
func (*UpdateTyping) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{7} }

func (m *UpdateTyping) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *UpdateTyping) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateTyping) GetTypingType() TypingType {
	if m != nil {
		return m.TypingType
	}
	return TypingType_TYPINGTYPE_UNKNOWN
}

// Update about user's typing stop
type UpdateTypingStop struct {
	Peer       *Peer      `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	Uid        int32      `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	TypingType TypingType `protobuf:"varint,3,opt,name=typing_type,json=typingType,enum=dialog.TypingType" json:"typing_type,omitempty"`
}

func (m *UpdateTypingStop) Reset()                    { *m = UpdateTypingStop{} }
func (m *UpdateTypingStop) String() string            { return proto.CompactTextString(m) }
func (*UpdateTypingStop) ProtoMessage()               {}
func (*UpdateTypingStop) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{8} }

func (m *UpdateTypingStop) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *UpdateTypingStop) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateTypingStop) GetTypingType() TypingType {
	if m != nil {
		return m.TypingType
	}
	return TypingType_TYPINGTYPE_UNKNOWN
}

// Update about user became online
type UpdateUserOnline struct {
	Uid            int32                        `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	DeviceType     DeviceType                   `protobuf:"varint,2,opt,name=device_type,json=deviceType,enum=dialog.DeviceType" json:"device_type,omitempty"`
	DeviceCategory *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=device_category,json=deviceCategory" json:"device_category,omitempty"`
}

func (m *UpdateUserOnline) Reset()                    { *m = UpdateUserOnline{} }
func (m *UpdateUserOnline) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserOnline) ProtoMessage()               {}
func (*UpdateUserOnline) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{9} }

func (m *UpdateUserOnline) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateUserOnline) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICETYPE_UNKNOWN
}

func (m *UpdateUserOnline) GetDeviceCategory() *google_protobuf.StringValue {
	if m != nil {
		return m.DeviceCategory
	}
	return nil
}

// Update about user became offline
type UpdateUserOffline struct {
	Uid            int32                        `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	DeviceType     DeviceType                   `protobuf:"varint,2,opt,name=device_type,json=deviceType,enum=dialog.DeviceType" json:"device_type,omitempty"`
	DeviceCategory *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=device_category,json=deviceCategory" json:"device_category,omitempty"`
}

func (m *UpdateUserOffline) Reset()                    { *m = UpdateUserOffline{} }
func (m *UpdateUserOffline) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserOffline) ProtoMessage()               {}
func (*UpdateUserOffline) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{10} }

func (m *UpdateUserOffline) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateUserOffline) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICETYPE_UNKNOWN
}

func (m *UpdateUserOffline) GetDeviceCategory() *google_protobuf.StringValue {
	if m != nil {
		return m.DeviceCategory
	}
	return nil
}

// Update about user's last seen state
type UpdateUserLastSeen struct {
	Uid            int32                        `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Date           int64                        `protobuf:"varint,2,opt,name=date" json:"date,omitempty"`
	DeviceType     DeviceType                   `protobuf:"varint,3,opt,name=device_type,json=deviceType,enum=dialog.DeviceType" json:"device_type,omitempty"`
	DeviceCategory *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=device_category,json=deviceCategory" json:"device_category,omitempty"`
}

func (m *UpdateUserLastSeen) Reset()                    { *m = UpdateUserLastSeen{} }
func (m *UpdateUserLastSeen) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserLastSeen) ProtoMessage()               {}
func (*UpdateUserLastSeen) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{11} }

func (m *UpdateUserLastSeen) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateUserLastSeen) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *UpdateUserLastSeen) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICETYPE_UNKNOWN
}

func (m *UpdateUserLastSeen) GetDeviceCategory() *google_protobuf.StringValue {
	if m != nil {
		return m.DeviceCategory
	}
	return nil
}

// Update about group online change
type UpdateGroupOnline struct {
	GroupId int32 `protobuf:"varint,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	// / amount of online users
	Count int32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *UpdateGroupOnline) Reset()                    { *m = UpdateGroupOnline{} }
func (m *UpdateGroupOnline) String() string            { return proto.CompactTextString(m) }
func (*UpdateGroupOnline) ProtoMessage()               {}
func (*UpdateGroupOnline) Descriptor() ([]byte, []int) { return fileDescriptor25, []int{12} }

func (m *UpdateGroupOnline) GetGroupId() int32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *UpdateGroupOnline) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestTyping)(nil), "dialog.RequestTyping")
	proto.RegisterType((*RequestStopTyping)(nil), "dialog.RequestStopTyping")
	proto.RegisterType((*RequestSetOnline)(nil), "dialog.RequestSetOnline")
	proto.RegisterType((*UpdatePauseNotifications)(nil), "dialog.UpdatePauseNotifications")
	proto.RegisterType((*UpdateRestoreNotifications)(nil), "dialog.UpdateRestoreNotifications")
	proto.RegisterType((*RequestPauseNotifications)(nil), "dialog.RequestPauseNotifications")
	proto.RegisterType((*RequestRestoreNotifications)(nil), "dialog.RequestRestoreNotifications")
	proto.RegisterType((*UpdateTyping)(nil), "dialog.UpdateTyping")
	proto.RegisterType((*UpdateTypingStop)(nil), "dialog.UpdateTypingStop")
	proto.RegisterType((*UpdateUserOnline)(nil), "dialog.UpdateUserOnline")
	proto.RegisterType((*UpdateUserOffline)(nil), "dialog.UpdateUserOffline")
	proto.RegisterType((*UpdateUserLastSeen)(nil), "dialog.UpdateUserLastSeen")
	proto.RegisterType((*UpdateGroupOnline)(nil), "dialog.UpdateGroupOnline")
	proto.RegisterEnum("dialog.TypingType", TypingType_name, TypingType_value)
	proto.RegisterEnum("dialog.DeviceType", DeviceType_name, DeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TypingAndOnline service

type TypingAndOnlineClient interface {
	Typing(ctx context.Context, in *RequestTyping, opts ...grpc.CallOption) (*ResponseVoid, error)
	StopTyping(ctx context.Context, in *RequestStopTyping, opts ...grpc.CallOption) (*ResponseVoid, error)
	SetOnline(ctx context.Context, in *RequestSetOnline, opts ...grpc.CallOption) (*ResponseVoid, error)
	PauseNotifications(ctx context.Context, in *RequestPauseNotifications, opts ...grpc.CallOption) (*ResponseVoid, error)
	RestoreNotifications(ctx context.Context, in *RequestRestoreNotifications, opts ...grpc.CallOption) (*ResponseVoid, error)
}

type typingAndOnlineClient struct {
	cc *grpc.ClientConn
}

func NewTypingAndOnlineClient(cc *grpc.ClientConn) TypingAndOnlineClient {
	return &typingAndOnlineClient{cc}
}

func (c *typingAndOnlineClient) Typing(ctx context.Context, in *RequestTyping, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := grpc.Invoke(ctx, "/dialog.TypingAndOnline/Typing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typingAndOnlineClient) StopTyping(ctx context.Context, in *RequestStopTyping, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := grpc.Invoke(ctx, "/dialog.TypingAndOnline/StopTyping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typingAndOnlineClient) SetOnline(ctx context.Context, in *RequestSetOnline, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := grpc.Invoke(ctx, "/dialog.TypingAndOnline/SetOnline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typingAndOnlineClient) PauseNotifications(ctx context.Context, in *RequestPauseNotifications, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := grpc.Invoke(ctx, "/dialog.TypingAndOnline/PauseNotifications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typingAndOnlineClient) RestoreNotifications(ctx context.Context, in *RequestRestoreNotifications, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := grpc.Invoke(ctx, "/dialog.TypingAndOnline/RestoreNotifications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TypingAndOnline service

type TypingAndOnlineServer interface {
	Typing(context.Context, *RequestTyping) (*ResponseVoid, error)
	StopTyping(context.Context, *RequestStopTyping) (*ResponseVoid, error)
	SetOnline(context.Context, *RequestSetOnline) (*ResponseVoid, error)
	PauseNotifications(context.Context, *RequestPauseNotifications) (*ResponseVoid, error)
	RestoreNotifications(context.Context, *RequestRestoreNotifications) (*ResponseVoid, error)
}

func RegisterTypingAndOnlineServer(s *grpc.Server, srv TypingAndOnlineServer) {
	s.RegisterService(&_TypingAndOnline_serviceDesc, srv)
}

func _TypingAndOnline_Typing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTyping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypingAndOnlineServer).Typing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.TypingAndOnline/Typing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypingAndOnlineServer).Typing(ctx, req.(*RequestTyping))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypingAndOnline_StopTyping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestStopTyping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypingAndOnlineServer).StopTyping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.TypingAndOnline/StopTyping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypingAndOnlineServer).StopTyping(ctx, req.(*RequestStopTyping))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypingAndOnline_SetOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSetOnline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypingAndOnlineServer).SetOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.TypingAndOnline/SetOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypingAndOnlineServer).SetOnline(ctx, req.(*RequestSetOnline))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypingAndOnline_PauseNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPauseNotifications)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypingAndOnlineServer).PauseNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.TypingAndOnline/PauseNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypingAndOnlineServer).PauseNotifications(ctx, req.(*RequestPauseNotifications))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypingAndOnline_RestoreNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRestoreNotifications)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypingAndOnlineServer).RestoreNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.TypingAndOnline/RestoreNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypingAndOnlineServer).RestoreNotifications(ctx, req.(*RequestRestoreNotifications))
	}
	return interceptor(ctx, in, info, handler)
}

var _TypingAndOnline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.TypingAndOnline",
	HandlerType: (*TypingAndOnlineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Typing",
			Handler:    _TypingAndOnline_Typing_Handler,
		},
		{
			MethodName: "StopTyping",
			Handler:    _TypingAndOnline_StopTyping_Handler,
		},
		{
			MethodName: "SetOnline",
			Handler:    _TypingAndOnline_SetOnline_Handler,
		},
		{
			MethodName: "PauseNotifications",
			Handler:    _TypingAndOnline_PauseNotifications_Handler,
		},
		{
			MethodName: "RestoreNotifications",
			Handler:    _TypingAndOnline_RestoreNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "typing_and_online.proto",
}

func init() { proto.RegisterFile("typing_and_online.proto", fileDescriptor25) }

var fileDescriptor25 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x49, 0x9a, 0xa6, 0xaf, 0xff, 0x9c, 0xd9, 0x76, 0x37, 0x2d, 0x15, 0xbb, 0xeb, 0x22,
	0xb6, 0x0a, 0x5a, 0xa7, 0x64, 0x25, 0x04, 0xcb, 0x61, 0x95, 0x64, 0xad, 0x12, 0x51, 0x92, 0x28,
	0xeb, 0x76, 0xd9, 0x53, 0xe5, 0xc4, 0x13, 0x33, 0xc8, 0xf5, 0x18, 0x7b, 0x5c, 0xd4, 0x2b, 0x12,
	0x97, 0xe5, 0xc8, 0x07, 0xe0, 0x02, 0x42, 0x9c, 0xe0, 0xbb, 0xb0, 0x9f, 0x60, 0x05, 0x97, 0xfd,
	0x14, 0xc8, 0xe3, 0x71, 0xe3, 0x38, 0xb6, 0x82, 0xaa, 0x22, 0xed, 0x29, 0xf1, 0x7b, 0xf3, 0x7e,
	0xbf, 0xdf, 0x7b, 0xf3, 0xde, 0xcc, 0xc0, 0x1d, 0x76, 0xe9, 0x12, 0xc7, 0x3a, 0x33, 0x1c, 0xf3,
	0x8c, 0x3a, 0x36, 0x71, 0xb0, 0xea, 0x7a, 0x94, 0x51, 0x54, 0x36, 0x89, 0x61, 0x53, 0x6b, 0xf7,
	0x3d, 0x8b, 0x52, 0xcb, 0xc6, 0x0d, 0x6e, 0x1d, 0x05, 0x93, 0xc6, 0x77, 0x9e, 0xe1, 0xba, 0xd8,
	0xf3, 0xa3, 0x75, 0xbb, 0x7b, 0xc2, 0x6f, 0xb8, 0xa4, 0x61, 0x38, 0x0e, 0x65, 0x06, 0x23, 0xd4,
	0x89, 0xbd, 0x55, 0x13, 0x4f, 0x88, 0x43, 0x92, 0xa6, 0x5b, 0xe7, 0xc4, 0x1f, 0x63, 0xdb, 0x36,
	0x1c, 0x4c, 0x83, 0xd8, 0xb8, 0xea, 0xe2, 0x29, 0xe4, 0xb6, 0x3f, 0x36, 0x6c, 0xc3, 0x1d, 0x35,
	0xc4, 0x6f, 0x64, 0x56, 0x7e, 0x95, 0x60, 0x7d, 0x88, 0xbf, 0x0d, 0xb0, 0xcf, 0x74, 0x2e, 0x1a,
	0x35, 0xa1, 0x14, 0xc6, 0xd5, 0xa4, 0x7b, 0xd2, 0xc1, 0x6a, 0x73, 0x53, 0x8d, 0x24, 0xab, 0xfd,
	0x80, 0x0d, 0x30, 0xf6, 0xda, 0xeb, 0x2f, 0xdf, 0x1c, 0xae, 0xc0, 0xf2, 0x05, 0xf1, 0xc9, 0xc8,
	0xc6, 0x43, 0xbe, 0x16, 0xb5, 0x61, 0x55, 0xa4, 0xcc, 0x2e, 0x5d, 0x5c, 0x2b, 0xde, 0x93, 0x0e,
	0x36, 0x9a, 0x28, 0x0e, 0x8d, 0x80, 0xf5, 0x4b, 0x17, 0xa7, 0xa3, 0x81, 0x5d, 0xb9, 0x1e, 0xef,
	0xbd, 0x7e, 0xb2, 0x03, 0x77, 0xc8, 0xb9, 0x6a, 0xda, 0x96, 0x6a, 0x79, 0xee, 0x58, 0x3d, 0xf2,
	0xdc, 0xb1, 0xd0, 0xa6, 0xfc, 0x2e, 0x41, 0x55, 0xfc, 0x7f, 0xc6, 0xa8, 0x7b, 0x73, 0x5a, 0x0b,
	0x37, 0xaf, 0xf5, 0xe7, 0x02, 0xc8, 0xb1, 0x56, 0xcc, 0xfa, 0xbc, 0x01, 0x50, 0x1d, 0x56, 0x88,
	0x2f, 0xba, 0x81, 0xeb, 0xad, 0xa4, 0x09, 0x2a, 0xc4, 0x17, 0x6b, 0x1f, 0xc0, 0x32, 0x23, 0xe7,
	0x98, 0x06, 0x8c, 0xcb, 0x2b, 0xa6, 0x57, 0xc6, 0xde, 0x30, 0x17, 0x13, 0x5f, 0x90, 0x31, 0xce,
	0xac, 0xfb, 0x53, 0xee, 0xca, 0xcc, 0xc5, 0xbc, 0x72, 0x21, 0x1d, 0x36, 0x05, 0xc6, 0xd8, 0x60,
	0xd8, 0xa2, 0xde, 0x65, 0xad, 0xc4, 0xcb, 0xb9, 0xa7, 0x46, 0x5d, 0xa8, 0xc6, 0x5d, 0xaa, 0x3e,
	0x63, 0x1e, 0x71, 0xac, 0x53, 0xc3, 0x0e, 0xe6, 0x10, 0x37, 0x22, 0x8c, 0x8e, 0x80, 0x58, 0x50,
	0xa1, 0x0e, 0xd4, 0x4e, 0x5c, 0xd3, 0x60, 0x78, 0x60, 0x04, 0x3e, 0xee, 0x51, 0x46, 0x26, 0x64,
	0x1c, 0xf5, 0x78, 0x32, 0xf9, 0xb0, 0x4c, 0x4b, 0x79, 0xc9, 0x2b, 0x7b, 0xb0, 0x1b, 0x81, 0x0c,
	0xb1, 0xcf, 0xa8, 0x37, 0x0b, 0xa3, 0x18, 0xb0, 0x23, 0xd8, 0x32, 0x38, 0x3e, 0x48, 0x73, 0xac,
	0xbd, 0x7c, 0x73, 0x58, 0x81, 0xf2, 0xd7, 0xc4, 0x34, 0xb1, 0x73, 0x45, 0xb1, 0x20, 0x8b, 0xcf,
	0xe0, 0x5d, 0xf1, 0x37, 0x4b, 0xc1, 0x82, 0xe0, 0x5f, 0x24, 0x58, 0x8b, 0xe4, 0x8b, 0x5e, 0x6e,
	0xcc, 0xf4, 0xf2, 0x5a, 0xbc, 0x89, 0xf9, 0x8d, 0x7c, 0x17, 0x8a, 0x01, 0x31, 0x79, 0x87, 0xcc,
	0x15, 0x29, 0xf4, 0xdc, 0xc4, 0x54, 0x2a, 0xbf, 0x49, 0x20, 0x27, 0x65, 0x86, 0xc3, 0xf7, 0x3f,
	0x48, 0x6d, 0xfd, 0x57, 0xa9, 0xb3, 0x1b, 0x95, 0x54, 0xfa, 0xc7, 0x95, 0xd2, 0x13, 0x1f, 0x7b,
	0x62, 0x92, 0x04, 0xb1, 0x94, 0x4b, 0xfc, 0x68, 0x76, 0x82, 0x0a, 0x79, 0x13, 0x34, 0x33, 0x32,
	0xda, 0xfc, 0xc8, 0x14, 0x17, 0x8f, 0x4c, 0x7a, 0x46, 0x94, 0x3f, 0x25, 0xa8, 0x26, 0x14, 0x4f,
	0x26, 0x6f, 0xbf, 0xe4, 0x57, 0x12, 0xa0, 0xa9, 0xe4, 0x63, 0x23, 0x3c, 0xe1, 0xb0, 0xb3, 0x58,
	0xf3, 0x7d, 0x28, 0x85, 0x41, 0xd9, 0xc7, 0x19, 0x77, 0xa5, 0xd3, 0x2a, 0x5e, 0x37, 0xad, 0xd2,
	0x35, 0xd2, 0x1a, 0xc5, 0x1b, 0x71, 0xe4, 0xd1, 0xc0, 0x15, 0xbd, 0x73, 0x00, 0x15, 0x2b, 0xfc,
	0x3c, 0xcb, 0xcb, 0x6c, 0x99, 0xbb, 0xbb, 0x26, 0xda, 0x87, 0xa5, 0x31, 0x0d, 0x1c, 0x96, 0xdd,
	0xe0, 0x91, 0xaf, 0xfe, 0x29, 0xc0, 0xb4, 0x8f, 0xd1, 0x6d, 0x40, 0xfa, 0x8b, 0x41, 0xb7, 0x77,
	0xa4, 0xbf, 0x18, 0x68, 0x67, 0x27, 0xbd, 0x2f, 0x7a, 0xfd, 0xe7, 0x3d, 0xf9, 0x1d, 0x74, 0x0b,
	0x36, 0x13, 0x76, 0x5d, 0xfb, 0x4a, 0x97, 0xa5, 0xfa, 0x2b, 0x09, 0x60, 0x5a, 0x80, 0x30, 0xf6,
	0xa9, 0x76, 0xda, 0xed, 0x68, 0xa9, 0xd8, 0x59, 0xfb, 0x91, 0xd6, 0xd3, 0x86, 0xdd, 0x8e, 0x2c,
	0xa1, 0x2a, 0xac, 0x27, 0xec, 0x83, 0x8e, 0x5c, 0x40, 0xdb, 0x50, 0x4d, 0x98, 0xbe, 0xec, 0xb7,
	0xbb, 0xc7, 0x9a, 0x5c, 0x4c, 0x99, 0xf5, 0x56, 0xfb, 0x58, 0xd3, 0xe5, 0x12, 0xda, 0x02, 0x39,
	0x61, 0x7e, 0xde, 0xd2, 0x3b, 0x9f, 0xcb, 0x4b, 0x69, 0x8c, 0xee, 0x70, 0xd8, 0x1f, 0xca, 0x65,
	0x84, 0x60, 0x23, 0x61, 0xee, 0xb4, 0x86, 0xf2, 0x72, 0x0a, 0x80, 0xe3, 0xca, 0x95, 0xe6, 0x3f,
	0x25, 0xd8, 0x8c, 0x4a, 0xd2, 0x72, 0x4c, 0x51, 0x74, 0x03, 0xca, 0xe2, 0x3c, 0xdc, 0x8e, 0xb7,
	0x7e, 0xe6, 0x79, 0xb2, 0xbb, 0x35, 0x35, 0xfb, 0x2e, 0x75, 0x7c, 0x7c, 0x4a, 0x89, 0xa9, 0xd4,
	0xbf, 0xff, 0xeb, 0xef, 0x9f, 0x0a, 0xef, 0x2b, 0x77, 0x1b, 0x17, 0x1f, 0x35, 0xc2, 0xb3, 0xb6,
	0x91, 0x02, 0x16, 0xdf, 0x8f, 0xa5, 0x3a, 0x3a, 0x07, 0x48, 0x3c, 0x21, 0x76, 0x52, 0x34, 0x53,
	0x57, 0x0e, 0x95, 0xca, 0xa9, 0x0e, 0x94, 0xfd, 0x5c, 0xaa, 0x29, 0x44, 0x48, 0xf7, 0x0d, 0xac,
	0x4c, 0x5f, 0x01, 0xb5, 0x34, 0x5b, 0xec, 0xc9, 0x21, 0x7b, 0xc8, 0xc9, 0x1e, 0x28, 0x4a, 0x3e,
	0x59, 0x8c, 0x10, 0x72, 0xfd, 0x20, 0x01, 0xca, 0xb8, 0xee, 0xee, 0xa7, 0x58, 0xe7, 0x97, 0xe4,
	0xd0, 0x7f, 0xcc, 0xe9, 0x0f, 0x95, 0x0f, 0x73, 0xe9, 0xe7, 0xa1, 0x42, 0x1d, 0x3f, 0x4a, 0xb0,
	0x95, 0x75, 0x27, 0xa2, 0xfd, 0x94, 0x92, 0xac, 0x45, 0x39, 0x5a, 0x3e, 0xe1, 0x5a, 0x9a, 0xca,
	0xc3, 0x5c, 0x2d, 0x99, 0xb7, 0xb0, 0x54, 0x6f, 0xef, 0xbc, 0x7e, 0x72, 0x1b, 0xb6, 0x92, 0x17,
	0xb1, 0x8f, 0xbd, 0x70, 0x9a, 0xfc, 0x51, 0x99, 0x1f, 0x0f, 0x8f, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x28, 0x41, 0x36, 0x39, 0xb2, 0x0b, 0x00, 0x00,
}