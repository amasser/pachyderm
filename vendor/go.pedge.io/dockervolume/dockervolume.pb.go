// Code generated by protoc-gen-go.
// source: dockervolume.proto
// DO NOT EDIT!

package dockervolume

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import google_api1 "google/api"
import google_protobuf1 "go.pedge.io/google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RemoveVolumeAttempt struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *RemoveVolumeAttempt) Reset()         { *m = RemoveVolumeAttempt{} }
func (m *RemoveVolumeAttempt) String() string { return proto.CompactTextString(m) }
func (*RemoveVolumeAttempt) ProtoMessage()    {}

type ActivateResponse struct {
	Implements []string `protobuf:"bytes,1,rep,name=implements" json:"implements,omitempty"`
}

func (m *ActivateResponse) Reset()         { *m = ActivateResponse{} }
func (m *ActivateResponse) String() string { return proto.CompactTextString(m) }
func (*ActivateResponse) ProtoMessage()    {}

type CreateRequest struct {
	Name string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Opts map[string]string `protobuf:"bytes,2,rep,name=opts" json:"opts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}

func (m *CreateRequest) GetOpts() map[string]string {
	if m != nil {
		return m.Opts
	}
	return nil
}

type CreateResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}

type RemoveRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}

type RemoveResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}

type PathRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PathRequest) Reset()         { *m = PathRequest{} }
func (m *PathRequest) String() string { return proto.CompactTextString(m) }
func (*PathRequest) ProtoMessage()    {}

type PathResponse struct {
	Mountpoint string `protobuf:"bytes,1,opt,name=mountpoint" json:"mountpoint,omitempty"`
	Err        string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *PathResponse) Reset()         { *m = PathResponse{} }
func (m *PathResponse) String() string { return proto.CompactTextString(m) }
func (*PathResponse) ProtoMessage()    {}

type MountRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *MountRequest) Reset()         { *m = MountRequest{} }
func (m *MountRequest) String() string { return proto.CompactTextString(m) }
func (*MountRequest) ProtoMessage()    {}

type MountResponse struct {
	Mountpoint string `protobuf:"bytes,1,opt,name=mountpoint" json:"mountpoint,omitempty"`
	Err        string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *MountResponse) Reset()         { *m = MountResponse{} }
func (m *MountResponse) String() string { return proto.CompactTextString(m) }
func (*MountResponse) ProtoMessage()    {}

type UnmountRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UnmountRequest) Reset()         { *m = UnmountRequest{} }
func (m *UnmountRequest) String() string { return proto.CompactTextString(m) }
func (*UnmountRequest) ProtoMessage()    {}

type UnmountResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *UnmountResponse) Reset()         { *m = UnmountResponse{} }
func (m *UnmountResponse) String() string { return proto.CompactTextString(m) }
func (*UnmountResponse) ProtoMessage()    {}

type CleanupResponse struct {
	RemoveVolumeAttempt []*RemoveVolumeAttempt `protobuf:"bytes,1,rep,name=remove_volume_attempt" json:"remove_volume_attempt,omitempty"`
	Err                 string                 `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CleanupResponse) Reset()         { *m = CleanupResponse{} }
func (m *CleanupResponse) String() string { return proto.CompactTextString(m) }
func (*CleanupResponse) ProtoMessage()    {}

func (m *CleanupResponse) GetRemoveVolumeAttempt() []*RemoveVolumeAttempt {
	if m != nil {
		return m.RemoveVolumeAttempt
	}
	return nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	Activate(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ActivateResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	Path(ctx context.Context, in *PathRequest, opts ...grpc.CallOption) (*PathResponse, error)
	Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*MountResponse, error)
	Unmount(ctx context.Context, in *UnmountRequest, opts ...grpc.CallOption) (*UnmountResponse, error)
	Cleanup(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CleanupResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Activate(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Activate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Path(ctx context.Context, in *PathRequest, opts ...grpc.CallOption) (*PathResponse, error) {
	out := new(PathResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Path", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*MountResponse, error) {
	out := new(MountResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Mount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Unmount(ctx context.Context, in *UnmountRequest, opts ...grpc.CallOption) (*UnmountResponse, error) {
	out := new(UnmountResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Unmount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Cleanup(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*CleanupResponse, error) {
	out := new(CleanupResponse)
	err := grpc.Invoke(ctx, "/dockervolume.API/Cleanup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Activate(context.Context, *google_protobuf1.Empty) (*ActivateResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	Path(context.Context, *PathRequest) (*PathResponse, error)
	Mount(context.Context, *MountRequest) (*MountResponse, error)
	Unmount(context.Context, *UnmountRequest) (*UnmountResponse, error)
	Cleanup(context.Context, *google_protobuf1.Empty) (*CleanupResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Activate_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Activate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Create_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(CreateRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Remove_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(RemoveRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Remove(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Path_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(PathRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Path(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Mount_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(MountRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Mount(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Unmount_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(UnmountRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Unmount(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _API_Cleanup_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Cleanup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dockervolume.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _API_Activate_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _API_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _API_Remove_Handler,
		},
		{
			MethodName: "Path",
			Handler:    _API_Path_Handler,
		},
		{
			MethodName: "Mount",
			Handler:    _API_Mount_Handler,
		},
		{
			MethodName: "Unmount",
			Handler:    _API_Unmount_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _API_Cleanup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
