// Code generated by protoc-gen-go.
// source: pfs/route/protolog/protolog.proto
// DO NOT EDIT!

/*
Package routeprotolog is a generated protocol buffer package.

It is generated from these files:
	pfs/route/protolog/protolog.proto

It has these top-level messages:
	StartRegister
	FinishRegister
	Version
	StartAssignRoles
	FinishAssignRoles
	SetServerState
	AddServerRole
	RemoveServerRole
	SetServerRole
	DeleteServerRole
	SetAddresses
	GetMasterAddress
	GetReplicaAddresses
	GetShardToMasterAddress
	ReplicaAddresses
	GetShardToReplicaAddresses
*/
package routeprotolog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import routeproto "github.com/pachyderm/pachyderm/src/pfs/route/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StartRegister struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *StartRegister) Reset()         { *m = StartRegister{} }
func (m *StartRegister) String() string { return proto.CompactTextString(m) }
func (*StartRegister) ProtoMessage()    {}

type FinishRegister struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FinishRegister) Reset()         { *m = FinishRegister{} }
func (m *FinishRegister) String() string { return proto.CompactTextString(m) }
func (*FinishRegister) ProtoMessage()    {}

type Version struct {
	Result int64  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}

type StartAssignRoles struct {
}

func (m *StartAssignRoles) Reset()         { *m = StartAssignRoles{} }
func (m *StartAssignRoles) String() string { return proto.CompactTextString(m) }
func (*StartAssignRoles) ProtoMessage()    {}

type FinishAssignRoles struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *FinishAssignRoles) Reset()         { *m = FinishAssignRoles{} }
func (m *FinishAssignRoles) String() string { return proto.CompactTextString(m) }
func (*FinishAssignRoles) ProtoMessage()    {}

type SetServerState struct {
	ServerState *routeproto.ServerState `protobuf:"bytes,1,opt,name=serverState" json:"serverState,omitempty"`
}

func (m *SetServerState) Reset()         { *m = SetServerState{} }
func (m *SetServerState) String() string { return proto.CompactTextString(m) }
func (*SetServerState) ProtoMessage()    {}

func (m *SetServerState) GetServerState() *routeproto.ServerState {
	if m != nil {
		return m.ServerState
	}
	return nil
}

type AddServerRole struct {
	ServerRole *routeproto.ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string                 `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *AddServerRole) Reset()         { *m = AddServerRole{} }
func (m *AddServerRole) String() string { return proto.CompactTextString(m) }
func (*AddServerRole) ProtoMessage()    {}

func (m *AddServerRole) GetServerRole() *routeproto.ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type RemoveServerRole struct {
	ServerRole *routeproto.ServerRole `protobuf:"bytes,1,opt,name=serverRole" json:"serverRole,omitempty"`
	Error      string                 `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveServerRole) Reset()         { *m = RemoveServerRole{} }
func (m *RemoveServerRole) String() string { return proto.CompactTextString(m) }
func (*RemoveServerRole) ProtoMessage()    {}

func (m *RemoveServerRole) GetServerRole() *routeproto.ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetServerRole struct {
	ServerRole *routeproto.ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *SetServerRole) Reset()         { *m = SetServerRole{} }
func (m *SetServerRole) String() string { return proto.CompactTextString(m) }
func (*SetServerRole) ProtoMessage()    {}

func (m *SetServerRole) GetServerRole() *routeproto.ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type DeleteServerRole struct {
	ServerRole *routeproto.ServerRole `protobuf:"bytes,2,opt,name=serverRole" json:"serverRole,omitempty"`
}

func (m *DeleteServerRole) Reset()         { *m = DeleteServerRole{} }
func (m *DeleteServerRole) String() string { return proto.CompactTextString(m) }
func (*DeleteServerRole) ProtoMessage()    {}

func (m *DeleteServerRole) GetServerRole() *routeproto.ServerRole {
	if m != nil {
		return m.ServerRole
	}
	return nil
}

type SetAddresses struct {
	Addresses *routeproto.Addresses `protobuf:"bytes,1,opt,name=addresses" json:"addresses,omitempty"`
}

func (m *SetAddresses) Reset()         { *m = SetAddresses{} }
func (m *SetAddresses) String() string { return proto.CompactTextString(m) }
func (*SetAddresses) ProtoMessage()    {}

func (m *SetAddresses) GetAddresses() *routeproto.Addresses {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetMasterAddress struct {
	Shard   uint64 `protobuf:"varint,1,opt,name=shard" json:"shard,omitempty"`
	Version int64  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Ok      bool   `protobuf:"varint,4,opt,name=ok" json:"ok,omitempty"`
	Error   string `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
}

func (m *GetMasterAddress) Reset()         { *m = GetMasterAddress{} }
func (m *GetMasterAddress) String() string { return proto.CompactTextString(m) }
func (*GetMasterAddress) ProtoMessage()    {}

type GetReplicaAddresses struct {
	Shard   uint64          `protobuf:"varint,1,opt,name=shard" json:"shard,omitempty"`
	Version int64           `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	Result  map[string]bool `protobuf:"bytes,3,rep,name=result" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Error   string          `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *GetReplicaAddresses) Reset()         { *m = GetReplicaAddresses{} }
func (m *GetReplicaAddresses) String() string { return proto.CompactTextString(m) }
func (*GetReplicaAddresses) ProtoMessage()    {}

func (m *GetReplicaAddresses) GetResult() map[string]bool {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetShardToMasterAddress struct {
	Version int64             `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Result  map[uint64]string `protobuf:"bytes,2,rep,name=result" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error   string            `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *GetShardToMasterAddress) Reset()         { *m = GetShardToMasterAddress{} }
func (m *GetShardToMasterAddress) String() string { return proto.CompactTextString(m) }
func (*GetShardToMasterAddress) ProtoMessage()    {}

func (m *GetShardToMasterAddress) GetResult() map[uint64]string {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReplicaAddresses struct {
	Addresses map[string]bool `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ReplicaAddresses) Reset()         { *m = ReplicaAddresses{} }
func (m *ReplicaAddresses) String() string { return proto.CompactTextString(m) }
func (*ReplicaAddresses) ProtoMessage()    {}

func (m *ReplicaAddresses) GetAddresses() map[string]bool {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type GetShardToReplicaAddresses struct {
	Version int64                        `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Result  map[uint64]*ReplicaAddresses `protobuf:"bytes,2,rep,name=result" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error   string                       `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *GetShardToReplicaAddresses) Reset()         { *m = GetShardToReplicaAddresses{} }
func (m *GetShardToReplicaAddresses) String() string { return proto.CompactTextString(m) }
func (*GetShardToReplicaAddresses) ProtoMessage()    {}

func (m *GetShardToReplicaAddresses) GetResult() map[uint64]*ReplicaAddresses {
	if m != nil {
		return m.Result
	}
	return nil
}
