// Code generated by protoc-gen-go.
// source: pdpb.proto
// DO NOT EDIT!

/*
Package pdpb is a generated protocol buffer package.

It is generated from these files:
	pdpb.proto

It has these top-level messages:
	Leader
	TsoRequest
	Timestamp
	TsoResponse
	BootstrapRequest
	BootstrapResponse
	IsBootstrappedRequest
	IsBootstrappedResponse
	AllocIdRequest
	AllocIdResponse
	GetStoreRequest
	GetRegionRequest
	GetClusterConfigRequest
	GetStoreResponse
	GetRegionResponse
	GetClusterConfigResponse
	PutStoreRequest
	PutStoreResponse
	PutClusterConfigRequest
	PutClusterConfigResponse
	AskChangePeerRequest
	AskChangePeerResponse
	AskSplitRequest
	AskSplitResponse
	RequestHeader
	ResponseHeader
	Request
	Response
	BootstrappedError
	Error
*/
package pdpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import metapb "github.com/pingcap/kvproto/pkg/metapb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CommandType int32

const (
	CommandType_Invalid          CommandType = 0
	CommandType_Tso              CommandType = 1
	CommandType_Bootstrap        CommandType = 2
	CommandType_IsBootstrapped   CommandType = 3
	CommandType_AllocId          CommandType = 4
	CommandType_GetStore         CommandType = 5
	CommandType_PutStore         CommandType = 6
	CommandType_DeleteMeta       CommandType = 7
	CommandType_AskChangePeer    CommandType = 8
	CommandType_AskSplit         CommandType = 9
	CommandType_GetRegion        CommandType = 10
	CommandType_GetClusterConfig CommandType = 11
	CommandType_PutClusterConfig CommandType = 12
)

var CommandType_name = map[int32]string{
	0:  "Invalid",
	1:  "Tso",
	2:  "Bootstrap",
	3:  "IsBootstrapped",
	4:  "AllocId",
	5:  "GetStore",
	6:  "PutStore",
	7:  "DeleteMeta",
	8:  "AskChangePeer",
	9:  "AskSplit",
	10: "GetRegion",
	11: "GetClusterConfig",
	12: "PutClusterConfig",
}
var CommandType_value = map[string]int32{
	"Invalid":          0,
	"Tso":              1,
	"Bootstrap":        2,
	"IsBootstrapped":   3,
	"AllocId":          4,
	"GetStore":         5,
	"PutStore":         6,
	"DeleteMeta":       7,
	"AskChangePeer":    8,
	"AskSplit":         9,
	"GetRegion":        10,
	"GetClusterConfig": 11,
	"PutClusterConfig": 12,
}

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}
func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (x *CommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandType_value, data, "CommandType")
	if err != nil {
		return err
	}
	*x = CommandType(value)
	return nil
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Leader struct {
	Addr             *string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Pid              *int64  `protobuf:"varint,2,opt,name=pid" json:"pid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Leader) Reset()                    { *m = Leader{} }
func (m *Leader) String() string            { return proto.CompactTextString(m) }
func (*Leader) ProtoMessage()               {}
func (*Leader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Leader) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *Leader) GetPid() int64 {
	if m != nil && m.Pid != nil {
		return *m.Pid
	}
	return 0
}

type TsoRequest struct {
	Number           *uint32 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TsoRequest) Reset()                    { *m = TsoRequest{} }
func (m *TsoRequest) String() string            { return proto.CompactTextString(m) }
func (*TsoRequest) ProtoMessage()               {}
func (*TsoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TsoRequest) GetNumber() uint32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type Timestamp struct {
	Physical         *int64 `protobuf:"varint,1,opt,name=physical" json:"physical,omitempty"`
	Logical          *int64 `protobuf:"varint,2,opt,name=logical" json:"logical,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Timestamp) GetPhysical() int64 {
	if m != nil && m.Physical != nil {
		return *m.Physical
	}
	return 0
}

func (m *Timestamp) GetLogical() int64 {
	if m != nil && m.Logical != nil {
		return *m.Logical
	}
	return 0
}

type TsoResponse struct {
	Timestamps       []*Timestamp `protobuf:"bytes,1,rep,name=timestamps" json:"timestamps,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *TsoResponse) Reset()                    { *m = TsoResponse{} }
func (m *TsoResponse) String() string            { return proto.CompactTextString(m) }
func (*TsoResponse) ProtoMessage()               {}
func (*TsoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TsoResponse) GetTimestamps() []*Timestamp {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type BootstrapRequest struct {
	Store            *metapb.Store  `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	Region           *metapb.Region `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *BootstrapRequest) Reset()                    { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string            { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()               {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BootstrapRequest) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

func (m *BootstrapRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type BootstrapResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrapResponse) Reset()                    { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string            { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()               {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type IsBootstrappedRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedRequest) Reset()                    { *m = IsBootstrappedRequest{} }
func (m *IsBootstrappedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedRequest) ProtoMessage()               {}
func (*IsBootstrappedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type IsBootstrappedResponse struct {
	Bootstrapped     *bool  `protobuf:"varint,1,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsBootstrappedResponse) Reset()                    { *m = IsBootstrappedResponse{} }
func (m *IsBootstrappedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsBootstrappedResponse) ProtoMessage()               {}
func (*IsBootstrappedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsBootstrappedResponse) GetBootstrapped() bool {
	if m != nil && m.Bootstrapped != nil {
		return *m.Bootstrapped
	}
	return false
}

type AllocIdRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AllocIdRequest) Reset()                    { *m = AllocIdRequest{} }
func (m *AllocIdRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocIdRequest) ProtoMessage()               {}
func (*AllocIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AllocIdResponse struct {
	Id               *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AllocIdResponse) Reset()                    { *m = AllocIdResponse{} }
func (m *AllocIdResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocIdResponse) ProtoMessage()               {}
func (*AllocIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AllocIdResponse) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type GetStoreRequest struct {
	StoreId          *uint64 `protobuf:"varint,1,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStoreRequest) Reset()                    { *m = GetStoreRequest{} }
func (m *GetStoreRequest) String() string            { return proto.CompactTextString(m) }
func (*GetStoreRequest) ProtoMessage()               {}
func (*GetStoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetStoreRequest) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type GetRegionRequest struct {
	RegionKey        []byte `protobuf:"bytes,1,opt,name=region_key" json:"region_key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetRegionRequest) Reset()                    { *m = GetRegionRequest{} }
func (m *GetRegionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRegionRequest) ProtoMessage()               {}
func (*GetRegionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetRegionRequest) GetRegionKey() []byte {
	if m != nil {
		return m.RegionKey
	}
	return nil
}

type GetClusterConfigRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetClusterConfigRequest) Reset()                    { *m = GetClusterConfigRequest{} }
func (m *GetClusterConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetClusterConfigRequest) ProtoMessage()               {}
func (*GetClusterConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type GetStoreResponse struct {
	Store            *metapb.Store `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *GetStoreResponse) Reset()                    { *m = GetStoreResponse{} }
func (m *GetStoreResponse) String() string            { return proto.CompactTextString(m) }
func (*GetStoreResponse) ProtoMessage()               {}
func (*GetStoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetStoreResponse) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type GetRegionResponse struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GetRegionResponse) Reset()                    { *m = GetRegionResponse{} }
func (m *GetRegionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRegionResponse) ProtoMessage()               {}
func (*GetRegionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetRegionResponse) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type GetClusterConfigResponse struct {
	Cluster          *metapb.Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetClusterConfigResponse) Reset()                    { *m = GetClusterConfigResponse{} }
func (m *GetClusterConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetClusterConfigResponse) ProtoMessage()               {}
func (*GetClusterConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetClusterConfigResponse) GetCluster() *metapb.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type PutStoreRequest struct {
	Store            *metapb.Store `protobuf:"bytes,1,opt,name=store" json:"store,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *PutStoreRequest) Reset()                    { *m = PutStoreRequest{} }
func (m *PutStoreRequest) String() string            { return proto.CompactTextString(m) }
func (*PutStoreRequest) ProtoMessage()               {}
func (*PutStoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PutStoreRequest) GetStore() *metapb.Store {
	if m != nil {
		return m.Store
	}
	return nil
}

type PutStoreResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutStoreResponse) Reset()                    { *m = PutStoreResponse{} }
func (m *PutStoreResponse) String() string            { return proto.CompactTextString(m) }
func (*PutStoreResponse) ProtoMessage()               {}
func (*PutStoreResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type PutClusterConfigRequest struct {
	Cluster          *metapb.Cluster `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PutClusterConfigRequest) Reset()                    { *m = PutClusterConfigRequest{} }
func (m *PutClusterConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*PutClusterConfigRequest) ProtoMessage()               {}
func (*PutClusterConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *PutClusterConfigRequest) GetCluster() *metapb.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

type PutClusterConfigResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutClusterConfigResponse) Reset()                    { *m = PutClusterConfigResponse{} }
func (m *PutClusterConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*PutClusterConfigResponse) ProtoMessage()               {}
func (*PutClusterConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

type AskChangePeerRequest struct {
	Region *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	// The current leader peer of the region.
	// Pd can first try to send command to this peer,
	// if the peer is not leader now, pd will try to
	// find the new leader of the region and then send
	// command again.
	Leader           *metapb.Peer `protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *AskChangePeerRequest) Reset()                    { *m = AskChangePeerRequest{} }
func (m *AskChangePeerRequest) String() string            { return proto.CompactTextString(m) }
func (*AskChangePeerRequest) ProtoMessage()               {}
func (*AskChangePeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *AskChangePeerRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *AskChangePeerRequest) GetLeader() *metapb.Peer {
	if m != nil {
		return m.Leader
	}
	return nil
}

type AskChangePeerResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AskChangePeerResponse) Reset()                    { *m = AskChangePeerResponse{} }
func (m *AskChangePeerResponse) String() string            { return proto.CompactTextString(m) }
func (*AskChangePeerResponse) ProtoMessage()               {}
func (*AskChangePeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

type AskSplitRequest struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	SplitKey         []byte         `protobuf:"bytes,2,opt,name=split_key" json:"split_key,omitempty"`
	Leader           *metapb.Peer   `protobuf:"bytes,3,opt,name=leader" json:"leader,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *AskSplitRequest) Reset()                    { *m = AskSplitRequest{} }
func (m *AskSplitRequest) String() string            { return proto.CompactTextString(m) }
func (*AskSplitRequest) ProtoMessage()               {}
func (*AskSplitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *AskSplitRequest) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *AskSplitRequest) GetSplitKey() []byte {
	if m != nil {
		return m.SplitKey
	}
	return nil
}

func (m *AskSplitRequest) GetLeader() *metapb.Peer {
	if m != nil {
		return m.Leader
	}
	return nil
}

type AskSplitResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *AskSplitResponse) Reset()                    { *m = AskSplitResponse{} }
func (m *AskSplitResponse) String() string            { return proto.CompactTextString(m) }
func (*AskSplitResponse) ProtoMessage()               {}
func (*AskSplitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

type RequestHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *RequestHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *RequestHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

type ResponseHeader struct {
	// 16 bytes, to distinguish request.
	Uuid             []byte  `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	ClusterId        *uint64 `protobuf:"varint,2,opt,name=cluster_id" json:"cluster_id,omitempty"`
	Error            *Error  `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *ResponseHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

func (m *ResponseHeader) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *ResponseHeader) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Request struct {
	Header           *RequestHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType             `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoRequest              `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapRequest        `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedRequest   `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdRequest          `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetStore         *GetStoreRequest         `protobuf:"bytes,7,opt,name=get_store" json:"get_store,omitempty"`
	PutStore         *PutStoreRequest         `protobuf:"bytes,8,opt,name=put_store" json:"put_store,omitempty"`
	AskChangePeer    *AskChangePeerRequest    `protobuf:"bytes,9,opt,name=ask_change_peer" json:"ask_change_peer,omitempty"`
	AskSplit         *AskSplitRequest         `protobuf:"bytes,10,opt,name=ask_split" json:"ask_split,omitempty"`
	GetRegion        *GetRegionRequest        `protobuf:"bytes,11,opt,name=get_region" json:"get_region,omitempty"`
	GetClusterConfig *GetClusterConfigRequest `protobuf:"bytes,12,opt,name=get_cluster_config" json:"get_cluster_config,omitempty"`
	PutClusterConfig *PutClusterConfigRequest `protobuf:"bytes,13,opt,name=put_cluster_config" json:"put_cluster_config,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

func (m *Request) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Request) GetTso() *TsoRequest {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Request) GetBootstrap() *BootstrapRequest {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Request) GetIsBootstrapped() *IsBootstrappedRequest {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Request) GetAllocId() *AllocIdRequest {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Request) GetGetStore() *GetStoreRequest {
	if m != nil {
		return m.GetStore
	}
	return nil
}

func (m *Request) GetPutStore() *PutStoreRequest {
	if m != nil {
		return m.PutStore
	}
	return nil
}

func (m *Request) GetAskChangePeer() *AskChangePeerRequest {
	if m != nil {
		return m.AskChangePeer
	}
	return nil
}

func (m *Request) GetAskSplit() *AskSplitRequest {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

func (m *Request) GetGetRegion() *GetRegionRequest {
	if m != nil {
		return m.GetRegion
	}
	return nil
}

func (m *Request) GetGetClusterConfig() *GetClusterConfigRequest {
	if m != nil {
		return m.GetClusterConfig
	}
	return nil
}

func (m *Request) GetPutClusterConfig() *PutClusterConfigRequest {
	if m != nil {
		return m.PutClusterConfig
	}
	return nil
}

type Response struct {
	Header           *ResponseHeader           `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	CmdType          *CommandType              `protobuf:"varint,2,opt,name=cmd_type,enum=pdpb.CommandType" json:"cmd_type,omitempty"`
	Tso              *TsoResponse              `protobuf:"bytes,3,opt,name=tso" json:"tso,omitempty"`
	Bootstrap        *BootstrapResponse        `protobuf:"bytes,4,opt,name=bootstrap" json:"bootstrap,omitempty"`
	IsBootstrapped   *IsBootstrappedResponse   `protobuf:"bytes,5,opt,name=is_bootstrapped" json:"is_bootstrapped,omitempty"`
	AllocId          *AllocIdResponse          `protobuf:"bytes,6,opt,name=alloc_id" json:"alloc_id,omitempty"`
	GetStore         *GetStoreResponse         `protobuf:"bytes,7,opt,name=get_store" json:"get_store,omitempty"`
	PutStore         *PutStoreResponse         `protobuf:"bytes,8,opt,name=put_store" json:"put_store,omitempty"`
	AskChangePeer    *AskChangePeerResponse    `protobuf:"bytes,9,opt,name=ask_change_peer" json:"ask_change_peer,omitempty"`
	AskSplit         *AskSplitResponse         `protobuf:"bytes,10,opt,name=ask_split" json:"ask_split,omitempty"`
	GetRegion        *GetRegionResponse        `protobuf:"bytes,11,opt,name=get_region" json:"get_region,omitempty"`
	GetClusterConfig *GetClusterConfigResponse `protobuf:"bytes,12,opt,name=get_cluster_config" json:"get_cluster_config,omitempty"`
	PutClusterConfig *PutClusterConfigResponse `protobuf:"bytes,13,opt,name=put_cluster_config" json:"put_cluster_config,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

func (m *Response) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Response) GetTso() *TsoResponse {
	if m != nil {
		return m.Tso
	}
	return nil
}

func (m *Response) GetBootstrap() *BootstrapResponse {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *Response) GetIsBootstrapped() *IsBootstrappedResponse {
	if m != nil {
		return m.IsBootstrapped
	}
	return nil
}

func (m *Response) GetAllocId() *AllocIdResponse {
	if m != nil {
		return m.AllocId
	}
	return nil
}

func (m *Response) GetGetStore() *GetStoreResponse {
	if m != nil {
		return m.GetStore
	}
	return nil
}

func (m *Response) GetPutStore() *PutStoreResponse {
	if m != nil {
		return m.PutStore
	}
	return nil
}

func (m *Response) GetAskChangePeer() *AskChangePeerResponse {
	if m != nil {
		return m.AskChangePeer
	}
	return nil
}

func (m *Response) GetAskSplit() *AskSplitResponse {
	if m != nil {
		return m.AskSplit
	}
	return nil
}

func (m *Response) GetGetRegion() *GetRegionResponse {
	if m != nil {
		return m.GetRegion
	}
	return nil
}

func (m *Response) GetGetClusterConfig() *GetClusterConfigResponse {
	if m != nil {
		return m.GetClusterConfig
	}
	return nil
}

func (m *Response) GetPutClusterConfig() *PutClusterConfigResponse {
	if m != nil {
		return m.PutClusterConfig
	}
	return nil
}

type BootstrappedError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *BootstrappedError) Reset()                    { *m = BootstrappedError{} }
func (m *BootstrappedError) String() string            { return proto.CompactTextString(m) }
func (*BootstrappedError) ProtoMessage()               {}
func (*BootstrappedError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{28} }

type Error struct {
	Message          *string            `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Bootstrapped     *BootstrappedError `protobuf:"bytes,2,opt,name=bootstrapped" json:"bootstrapped,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{29} }

func (m *Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Error) GetBootstrapped() *BootstrappedError {
	if m != nil {
		return m.Bootstrapped
	}
	return nil
}

func init() {
	proto.RegisterType((*Leader)(nil), "pdpb.Leader")
	proto.RegisterType((*TsoRequest)(nil), "pdpb.TsoRequest")
	proto.RegisterType((*Timestamp)(nil), "pdpb.Timestamp")
	proto.RegisterType((*TsoResponse)(nil), "pdpb.TsoResponse")
	proto.RegisterType((*BootstrapRequest)(nil), "pdpb.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "pdpb.BootstrapResponse")
	proto.RegisterType((*IsBootstrappedRequest)(nil), "pdpb.IsBootstrappedRequest")
	proto.RegisterType((*IsBootstrappedResponse)(nil), "pdpb.IsBootstrappedResponse")
	proto.RegisterType((*AllocIdRequest)(nil), "pdpb.AllocIdRequest")
	proto.RegisterType((*AllocIdResponse)(nil), "pdpb.AllocIdResponse")
	proto.RegisterType((*GetStoreRequest)(nil), "pdpb.GetStoreRequest")
	proto.RegisterType((*GetRegionRequest)(nil), "pdpb.GetRegionRequest")
	proto.RegisterType((*GetClusterConfigRequest)(nil), "pdpb.GetClusterConfigRequest")
	proto.RegisterType((*GetStoreResponse)(nil), "pdpb.GetStoreResponse")
	proto.RegisterType((*GetRegionResponse)(nil), "pdpb.GetRegionResponse")
	proto.RegisterType((*GetClusterConfigResponse)(nil), "pdpb.GetClusterConfigResponse")
	proto.RegisterType((*PutStoreRequest)(nil), "pdpb.PutStoreRequest")
	proto.RegisterType((*PutStoreResponse)(nil), "pdpb.PutStoreResponse")
	proto.RegisterType((*PutClusterConfigRequest)(nil), "pdpb.PutClusterConfigRequest")
	proto.RegisterType((*PutClusterConfigResponse)(nil), "pdpb.PutClusterConfigResponse")
	proto.RegisterType((*AskChangePeerRequest)(nil), "pdpb.AskChangePeerRequest")
	proto.RegisterType((*AskChangePeerResponse)(nil), "pdpb.AskChangePeerResponse")
	proto.RegisterType((*AskSplitRequest)(nil), "pdpb.AskSplitRequest")
	proto.RegisterType((*AskSplitResponse)(nil), "pdpb.AskSplitResponse")
	proto.RegisterType((*RequestHeader)(nil), "pdpb.RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "pdpb.ResponseHeader")
	proto.RegisterType((*Request)(nil), "pdpb.Request")
	proto.RegisterType((*Response)(nil), "pdpb.Response")
	proto.RegisterType((*BootstrappedError)(nil), "pdpb.BootstrappedError")
	proto.RegisterType((*Error)(nil), "pdpb.Error")
	proto.RegisterEnum("pdpb.CommandType", CommandType_name, CommandType_value)
}

var fileDescriptor0 = []byte{
	// 960 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x96, 0x5b, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x49, 0x73, 0x3f, 0x76, 0x92, 0xc9, 0x6c, 0x9b, 0x84, 0x6e, 0x5b, 0xad, 0x5c, 0xb4,
	0x74, 0x8b, 0x28, 0x90, 0x5d, 0x1e, 0xb8, 0xbc, 0x2c, 0x05, 0x2d, 0x95, 0x00, 0x55, 0xbb, 0x7d,
	0x8f, 0xdc, 0x78, 0x48, 0xad, 0x3a, 0xb1, 0xf1, 0x38, 0x48, 0xfd, 0x96, 0xbc, 0x23, 0xbe, 0x0b,
	0x73, 0x39, 0xe3, 0x38, 0x63, 0xa7, 0xdb, 0xb7, 0x76, 0xfc, 0xff, 0x9f, 0x73, 0xe6, 0xcc, 0x6f,
	0x4e, 0x06, 0x20, 0x09, 0x92, 0xdb, 0x8b, 0x24, 0x8d, 0xb3, 0x98, 0x36, 0xe4, 0xdf, 0x87, 0xee,
	0x92, 0x65, 0xbe, 0x59, 0xf3, 0x4e, 0xa1, 0xf5, 0x1b, 0xf3, 0x03, 0x96, 0x52, 0x17, 0x1a, 0x7e,
	0x10, 0xa4, 0x93, 0xda, 0x8b, 0xda, 0x59, 0x97, 0x3a, 0x50, 0x4f, 0xc2, 0x60, 0xb2, 0x27, 0xfe,
	0xa9, 0x7b, 0x47, 0x00, 0x37, 0x3c, 0x7e, 0xcf, 0xfe, 0x5a, 0x33, 0x9e, 0xd1, 0x3e, 0xb4, 0x56,
	0xeb, 0xe5, 0x2d, 0xd3, 0xd2, 0x9e, 0x77, 0x01, 0xdd, 0x9b, 0x70, 0x29, 0xbe, 0xf8, 0xcb, 0x84,
	0x12, 0xe8, 0x24, 0x77, 0x0f, 0x3c, 0x9c, 0xfb, 0x91, 0xfa, 0x5c, 0xa7, 0x03, 0x68, 0x47, 0xf1,
	0x42, 0x2d, 0xe8, 0x68, 0x53, 0x70, 0x54, 0x34, 0x9e, 0xc4, 0x2b, 0xce, 0xe8, 0x29, 0x40, 0x66,
	0xec, 0x5c, 0x78, 0xea, 0x67, 0xce, 0x74, 0x70, 0xa1, 0xca, 0xce, 0xc3, 0x7a, 0xd7, 0x40, 0x7e,
	0x8a, 0xe3, 0x8c, 0x67, 0xa9, 0x9f, 0x98, 0x3a, 0x8e, 0xa0, 0xc9, 0xb3, 0x38, 0x65, 0x2a, 0x8f,
	0x33, 0xed, 0x5d, 0xe0, 0xc6, 0x3e, 0xc8, 0x45, 0x7a, 0x02, 0xad, 0x94, 0x2d, 0xc2, 0x78, 0xa5,
	0xb2, 0x3a, 0xd3, 0xbe, 0xf9, 0xfc, 0x5e, 0xad, 0x7a, 0xcf, 0x60, 0x58, 0x88, 0xa8, 0x6b, 0xf1,
	0xc6, 0x70, 0x70, 0xc5, 0xf3, 0xe5, 0x84, 0x05, 0x98, 0x4b, 0xec, 0x71, 0x64, 0x7f, 0xc0, 0xf2,
	0xf7, 0xc1, 0xbd, 0x2d, 0xac, 0xab, 0x62, 0x3a, 0x1e, 0x81, 0xfe, 0xdb, 0x28, 0x8a, 0xe7, 0x57,
	0x79, 0x84, 0x63, 0x18, 0xe4, 0x2b, 0x68, 0x05, 0xd8, 0x0b, 0xb5, 0xa1, 0x21, 0xce, 0x61, 0xf0,
	0x8e, 0x65, 0xaa, 0x74, 0xb3, 0x3f, 0xd1, 0x4a, 0xb5, 0xbf, 0x59, 0x2e, 0x7a, 0x09, 0x44, 0x88,
	0xf4, 0x06, 0x8c, 0x4a, 0x44, 0xd1, 0xfb, 0x9c, 0xdd, 0xb3, 0x07, 0xa5, 0x73, 0xbd, 0x4f, 0x61,
	0x2c, 0x74, 0x97, 0xd1, 0x9a, 0x67, 0x2c, 0xbd, 0x8c, 0x57, 0x7f, 0x86, 0x0b, 0x53, 0xc6, 0xd7,
	0x2a, 0x04, 0xe6, 0xc1, 0x3a, 0x1e, 0x6d, 0xa4, 0xf7, 0x1a, 0x86, 0x85, 0xa4, 0x68, 0xd9, 0x74,
	0xb7, 0x56, 0xd9, 0xdd, 0x1f, 0x61, 0x52, 0xae, 0x00, 0xbd, 0x2f, 0xa0, 0x3d, 0xd7, 0x1f, 0xd0,
	0x3c, 0x30, 0x66, 0xd4, 0x7b, 0x5f, 0xc1, 0xe0, 0x7a, 0xbd, 0xdd, 0x8c, 0xc7, 0x6b, 0xa4, 0x40,
	0x36, 0x06, 0x3c, 0xcb, 0x1f, 0x60, 0x2c, 0xd6, 0xaa, 0x9a, 0xf0, 0x84, 0x0a, 0x0e, 0x61, 0x52,
	0x36, 0x63, 0xe0, 0x1b, 0xd8, 0x7f, 0xcb, 0xef, 0x2f, 0xef, 0xfc, 0xd5, 0x82, 0x5d, 0x33, 0x96,
	0x9a, 0xa8, 0x1f, 0xe9, 0x89, 0xd8, 0x42, 0x2b, 0x52, 0x57, 0x0d, 0x89, 0x74, 0xcd, 0x77, 0x19,
	0x44, 0xa2, 0x67, 0x45, 0xc5, 0x74, 0xb7, 0x02, 0x1c, 0x7e, 0xff, 0x21, 0x89, 0xc2, 0xec, 0xa9,
	0x99, 0x86, 0xd0, 0xe5, 0x52, 0xaf, 0x90, 0x90, 0xc9, 0xdc, 0x42, 0xf2, 0x7a, 0x45, 0x72, 0xd1,
	0xbf, 0x4d, 0x0e, 0xcc, 0xfb, 0x0d, 0xf4, 0x30, 0xdf, 0xaf, 0xf9, 0x80, 0x58, 0xaf, 0x91, 0x45,
	0x57, 0x72, 0x87, 0x3d, 0x9c, 0xe1, 0x9c, 0x68, 0x78, 0x7f, 0x40, 0xdf, 0xd8, 0x9f, 0xea, 0xa1,
	0x87, 0xd0, 0x64, 0x69, 0x1a, 0x9b, 0xba, 0x1c, 0x7d, 0xf3, 0x7f, 0x91, 0x4b, 0xde, 0x3f, 0x0d,
	0x68, 0x9b, 0x3d, 0x8b, 0x41, 0x75, 0xa7, 0x37, 0xa0, 0xf7, 0xfc, 0x4c, 0x0b, 0xb7, 0x4b, 0x3c,
	0x85, 0xce, 0x7c, 0x19, 0xcc, 0xb2, 0x87, 0x84, 0xa9, 0xf0, 0xfd, 0xe9, 0x50, 0xcb, 0x2e, 0xe3,
	0xe5, 0xd2, 0x5f, 0x05, 0x37, 0xe2, 0x03, 0x3d, 0x86, 0x7a, 0xc6, 0x63, 0xcc, 0x47, 0x70, 0xd2,
	0x6c, 0xc6, 0xdb, 0x2b, 0xe8, 0xe6, 0x17, 0x7a, 0xd2, 0x50, 0xa2, 0x91, 0x16, 0x95, 0x26, 0xd0,
	0x1b, 0x18, 0x84, 0x7c, 0xb6, 0x75, 0xfd, 0x9b, 0xca, 0xf0, 0x5c, 0x1b, 0x2a, 0x67, 0x09, 0x7d,
	0x09, 0x1d, 0x5f, 0x4e, 0x02, 0xd9, 0x83, 0x96, 0x92, 0xef, 0x6b, 0xf9, 0xf6, 0xc4, 0xa0, 0x67,
	0xd0, 0x5d, 0xb0, 0x6c, 0xa6, 0xb1, 0x6f, 0x2b, 0xe1, 0x81, 0x16, 0xda, 0x93, 0x42, 0x28, 0x93,
	0xb5, 0x51, 0x76, 0x8a, 0x4a, 0xfb, 0x1a, 0xbd, 0x86, 0x81, 0xcf, 0xef, 0x67, 0x73, 0x85, 0xd9,
	0x2c, 0x11, 0x67, 0x3f, 0xe9, 0x2a, 0xfd, 0x21, 0x96, 0x50, 0x05, 0xb6, 0x08, 0x2f, 0x4d, 0x0a,
	0xa9, 0x09, 0x14, 0xc3, 0xdb, 0x60, 0x9e, 0x03, 0xc8, 0x92, 0x11, 0x4e, 0xa7, 0xd8, 0xbc, 0xd2,
	0xe0, 0xfa, 0x0e, 0xa8, 0xd4, 0x1a, 0x20, 0xe6, 0xea, 0x92, 0x4d, 0x5c, 0xe5, 0x39, 0xce, 0x3d,
	0x95, 0xf7, 0x57, 0x58, 0xe5, 0x7e, 0x2d, 0x6b, 0xaf, 0x68, 0xdd, 0x71, 0xf5, 0xbd, 0xff, 0x1a,
	0xd0, 0xc9, 0x27, 0xd1, 0x67, 0x16, 0x53, 0xfb, 0x86, 0xa9, 0x2d, 0x86, 0x9f, 0x04, 0xd5, 0x49,
	0x11, 0xaa, 0x61, 0x01, 0x2a, 0x4c, 0x75, 0x5e, 0xa6, 0x6a, 0x5c, 0xa2, 0x0a, 0xb5, 0xdf, 0xee,
	0xc2, 0xea, 0xa8, 0x1a, 0x2b, 0xb4, 0x7d, 0x5e, 0xe2, 0xea, 0xc0, 0xe2, 0x0a, 0x85, 0xaf, 0xca,
	0x60, 0x8d, 0x6c, 0xb0, 0x36, 0x52, 0x9b, 0xac, 0x91, 0x4d, 0x16, 0x4a, 0xdf, 0xec, 0x42, 0xeb,
	0x79, 0x25, 0x5a, 0x9b, 0x04, 0x36, 0x5b, 0x23, 0x9b, 0x2d, 0x94, 0x7e, 0x51, 0x01, 0xd7, 0xb8,
	0x04, 0x17, 0x8a, 0xbf, 0x7f, 0x84, 0xae, 0x93, 0x5d, 0x74, 0x6d, 0xbc, 0x3b, 0xf1, 0x3a, 0xd9,
	0x85, 0x17, 0x4e, 0xcd, 0xe2, 0xb3, 0x42, 0x1c, 0x8e, 0x9e, 0x63, 0xef, 0xa0, 0xa9, 0xfe, 0x90,
	0x6f, 0x21, 0xf1, 0xa2, 0xe1, 0xfe, 0x82, 0xe1, 0x33, 0xeb, 0x4b, 0xeb, 0xf5, 0xb0, 0x57, 0x49,
	0x86, 0x09, 0x74, 0xfe, 0x6f, 0x0d, 0x9c, 0x22, 0x75, 0x0e, 0xb4, 0xaf, 0x56, 0x7f, 0xfb, 0x51,
	0x18, 0x90, 0x4f, 0x68, 0x1b, 0xea, 0x82, 0x38, 0x52, 0xa3, 0x3d, 0xe8, 0xe6, 0x56, 0xb2, 0x27,
	0xa6, 0x6e, 0x7f, 0x9b, 0x18, 0x52, 0x97, 0x46, 0xa4, 0x82, 0x34, 0xc4, 0x90, 0xee, 0x98, 0x83,
	0x27, 0x4d, 0xf9, 0x9f, 0x39, 0x5b, 0xd2, 0x12, 0x8f, 0x3d, 0xf8, 0x99, 0x45, 0x2c, 0x63, 0xbf,
	0x8b, 0x9f, 0x0b, 0xd2, 0x16, 0x3f, 0x2d, 0xbd, 0xad, 0x83, 0x24, 0x1d, 0x69, 0x30, 0x67, 0x45,
	0xba, 0x32, 0x79, 0x7e, 0x1a, 0x04, 0xc4, 0xf3, 0x88, 0xd8, 0x7d, 0x26, 0x8e, 0x5c, 0xb5, 0x3b,
	0x48, 0xdc, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x68, 0xc2, 0xbd, 0xac, 0x0a, 0x00, 0x00,
}
