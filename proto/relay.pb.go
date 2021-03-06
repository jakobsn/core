// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relay.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// grpccmd imports
import (
	"io"

	"github.com/spf13/cobra"
	"github.com/sshaman1101/grpccmd"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PeerType int32

const (
	PeerType_SERVER   PeerType = 0
	PeerType_CLIENT   PeerType = 1
	PeerType_DISCOVER PeerType = 2
)

var PeerType_name = map[int32]string{
	0: "SERVER",
	1: "CLIENT",
	2: "DISCOVER",
}
var PeerType_value = map[string]int32{
	"SERVER":   0,
	"CLIENT":   1,
	"DISCOVER": 2,
}

func (x PeerType) String() string {
	return proto.EnumName(PeerType_name, int32(x))
}
func (PeerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type HandshakeRequest struct {
	// PeerType describes a peer's source.
	PeerType PeerType `protobuf:"varint,1,opt,name=peerType,enum=sonm.PeerType" json:"peerType,omitempty"`
	// Addr represents a common Ethereum address both peers are connecting
	// around.
	// In case of servers it's their own id. Must be signed. In case of
	// clients - it's the target server id.
	//
	// In case of discovery requests this field has special meaning.
	// Both client and server must discover the same relay server to be able to
	// meet each other. At this stage there is no parameter verification.
	// It is done in the Handshake method.
	Addr []byte `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	// Signature for ETH address.
	// Should be empty for clients.
	Sign []byte `protobuf:"bytes,3,opt,name=sign,proto3" json:"sign,omitempty"`
	// Optional connection id.
	// It is used when a client wants to connect to a specific server avoiding
	// random select.
	// Should be empty for servers.
	UUID string `protobuf:"bytes,4,opt,name=UUID" json:"UUID,omitempty"`
	// Protocol describes the network protocol the peer wants to publish or to
	// resolve.
	Protocol string `protobuf:"bytes,5,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *HandshakeRequest) Reset()                    { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()               {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *HandshakeRequest) GetPeerType() PeerType {
	if m != nil {
		return m.PeerType
	}
	return PeerType_SERVER
}

func (m *HandshakeRequest) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *HandshakeRequest) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *HandshakeRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *HandshakeRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type DiscoverResponse struct {
	// Addr represents network address in form "host:port".
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *DiscoverResponse) Reset()                    { *m = DiscoverResponse{} }
func (m *DiscoverResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoverResponse) ProtoMessage()               {}
func (*DiscoverResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *DiscoverResponse) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type HandshakeResponse struct {
	// Error describes an error number.
	// Zero value means that there is no error.
	Error int32 `protobuf:"varint,1,opt,name=error" json:"error,omitempty"`
	// Description describes an error above.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *HandshakeResponse) Reset()                    { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string            { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()               {}
func (*HandshakeResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *HandshakeResponse) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *HandshakeResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type RelayClusterReply struct {
	Members []string `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
}

func (m *RelayClusterReply) Reset()                    { *m = RelayClusterReply{} }
func (m *RelayClusterReply) String() string            { return proto.CompactTextString(m) }
func (*RelayClusterReply) ProtoMessage()               {}
func (*RelayClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *RelayClusterReply) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type RelayMetrics struct {
	ConnCurrent uint64                 `protobuf:"varint,1,opt,name=connCurrent" json:"connCurrent,omitempty"`
	Net         map[string]*NetMetrics `protobuf:"bytes,2,rep,name=net" json:"net,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Uptime      uint64                 `protobuf:"varint,3,opt,name=uptime" json:"uptime,omitempty"`
}

func (m *RelayMetrics) Reset()                    { *m = RelayMetrics{} }
func (m *RelayMetrics) String() string            { return proto.CompactTextString(m) }
func (*RelayMetrics) ProtoMessage()               {}
func (*RelayMetrics) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *RelayMetrics) GetConnCurrent() uint64 {
	if m != nil {
		return m.ConnCurrent
	}
	return 0
}

func (m *RelayMetrics) GetNet() map[string]*NetMetrics {
	if m != nil {
		return m.Net
	}
	return nil
}

func (m *RelayMetrics) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

type NetMetrics struct {
	TxBytes uint64 `protobuf:"varint,1,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes uint64 `protobuf:"varint,2,opt,name=rxBytes" json:"rxBytes,omitempty"`
}

func (m *NetMetrics) Reset()                    { *m = NetMetrics{} }
func (m *NetMetrics) String() string            { return proto.CompactTextString(m) }
func (*NetMetrics) ProtoMessage()               {}
func (*NetMetrics) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *NetMetrics) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetMetrics) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*HandshakeRequest)(nil), "sonm.HandshakeRequest")
	proto.RegisterType((*DiscoverResponse)(nil), "sonm.DiscoverResponse")
	proto.RegisterType((*HandshakeResponse)(nil), "sonm.HandshakeResponse")
	proto.RegisterType((*RelayClusterReply)(nil), "sonm.RelayClusterReply")
	proto.RegisterType((*RelayMetrics)(nil), "sonm.RelayMetrics")
	proto.RegisterType((*NetMetrics)(nil), "sonm.NetMetrics")
	proto.RegisterEnum("sonm.PeerType", PeerType_name, PeerType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Relay service

type RelayClient interface {
	Cluster(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RelayClusterReply, error)
	Metrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RelayMetrics, error)
}

type relayClient struct {
	cc *grpc.ClientConn
}

func NewRelayClient(cc *grpc.ClientConn) RelayClient {
	return &relayClient{cc}
}

func (c *relayClient) Cluster(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RelayClusterReply, error) {
	out := new(RelayClusterReply)
	err := grpc.Invoke(ctx, "/sonm.Relay/Cluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relayClient) Metrics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RelayMetrics, error) {
	out := new(RelayMetrics)
	err := grpc.Invoke(ctx, "/sonm.Relay/Metrics", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Relay service

type RelayServer interface {
	Cluster(context.Context, *Empty) (*RelayClusterReply, error)
	Metrics(context.Context, *Empty) (*RelayMetrics, error)
}

func RegisterRelayServer(s *grpc.Server, srv RelayServer) {
	s.RegisterService(&_Relay_serviceDesc, srv)
}

func _Relay_Cluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayServer).Cluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Relay/Cluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayServer).Cluster(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relay_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Relay/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayServer).Metrics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Relay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Relay",
	HandlerType: (*RelayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Cluster",
			Handler:    _Relay_Cluster_Handler,
		},
		{
			MethodName: "Metrics",
			Handler:    _Relay_Metrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay.proto",
}

// Begin grpccmd
var _ = grpccmd.RunE

// Relay
var _RelayCmd = &cobra.Command{
	Use:   "relay [method]",
	Short: "Subcommand for the Relay service.",
}

var _Relay_ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Make the Cluster method call, input-type: sonm.Empty output-type: sonm.RelayClusterReply",
	RunE: grpccmd.RunE(
		"Cluster",
		"sonm.Empty",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewRelayClient(cc)
		},
	),
}

var _Relay_ClusterCmd_gen = &cobra.Command{
	Use:   "cluster-gen",
	Short: "Generate JSON for method call of Cluster (input-type: sonm.Empty)",
	RunE:  grpccmd.TypeToJson("sonm.Empty"),
}

var _Relay_MetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Make the Metrics method call, input-type: sonm.Empty output-type: sonm.RelayMetrics",
	RunE: grpccmd.RunE(
		"Metrics",
		"sonm.Empty",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewRelayClient(cc)
		},
	),
}

var _Relay_MetricsCmd_gen = &cobra.Command{
	Use:   "metrics-gen",
	Short: "Generate JSON for method call of Metrics (input-type: sonm.Empty)",
	RunE:  grpccmd.TypeToJson("sonm.Empty"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_RelayCmd)
	_RelayCmd.AddCommand(
		_Relay_ClusterCmd,
		_Relay_ClusterCmd_gen,
		_Relay_MetricsCmd,
		_Relay_MetricsCmd_gen,
	)
}

// End grpccmd

func init() { proto.RegisterFile("relay.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0xad, 0xdb, 0xa6, 0x4d, 0x6f, 0xaa, 0x7d, 0x99, 0xf5, 0x09, 0xa2, 0xf2, 0x12, 0xe5, 0x61,
	0x8a, 0x26, 0x56, 0x41, 0x79, 0x41, 0x3c, 0x21, 0xda, 0x48, 0xab, 0x80, 0x82, 0xbc, 0x8d, 0xf7,
	0x2c, 0xbd, 0x62, 0xd1, 0x12, 0x27, 0xd8, 0xce, 0x44, 0xfe, 0x08, 0xff, 0x87, 0x7f, 0x86, 0x6c,
	0x27, 0x5b, 0x10, 0x4f, 0xb9, 0xf7, 0x9c, 0x6b, 0x9f, 0x73, 0x4f, 0x0c, 0x9e, 0xc0, 0x22, 0x6d,
	0xd7, 0xb5, 0xa8, 0x54, 0x45, 0xa7, 0xb2, 0xe2, 0xe5, 0xea, 0xbf, 0x9c, 0xeb, 0x2f, 0xcf, 0x53,
	0x0b, 0x47, 0xbf, 0x08, 0xf8, 0x97, 0x29, 0x3f, 0xca, 0xbb, 0xf4, 0x1e, 0x19, 0xfe, 0x68, 0x50,
	0x2a, 0x7a, 0x0e, 0x6e, 0x8d, 0x28, 0xae, 0xdb, 0x1a, 0x03, 0x12, 0x92, 0xf8, 0x64, 0x73, 0xb2,
	0xd6, 0xc7, 0xd6, 0x5f, 0x3b, 0x94, 0x3d, 0xf2, 0x94, 0xc2, 0x34, 0x3d, 0x1e, 0x45, 0x30, 0x0e,
	0x49, 0xbc, 0x64, 0xa6, 0xd6, 0x98, 0xcc, 0xbf, 0xf3, 0x60, 0x62, 0x31, 0x5d, 0x6b, 0xec, 0xe6,
	0x66, 0xbf, 0x0b, 0xa6, 0x21, 0x89, 0x17, 0xcc, 0xd4, 0x74, 0x05, 0xae, 0x71, 0x91, 0x55, 0x45,
	0xe0, 0x18, 0xfc, 0xb1, 0x8f, 0xce, 0xc0, 0xdf, 0xe5, 0x32, 0xab, 0x1e, 0x50, 0x30, 0x94, 0x75,
	0xc5, 0xe5, 0x93, 0x16, 0xb1, 0x77, 0xe8, 0x3a, 0xfa, 0x08, 0xa7, 0x03, 0xff, 0xdd, 0xe0, 0xff,
	0xe0, 0xa0, 0x10, 0x95, 0x9d, 0x74, 0x98, 0x6d, 0x68, 0x08, 0xde, 0x11, 0x65, 0x26, 0xf2, 0x5a,
	0xe5, 0x15, 0x37, 0x8e, 0x17, 0x6c, 0x08, 0x45, 0x17, 0x70, 0xca, 0x74, 0x66, 0xdb, 0xa2, 0x91,
	0x4a, 0x0b, 0xd7, 0x45, 0x4b, 0x03, 0x98, 0x97, 0x58, 0xde, 0xa2, 0x90, 0x01, 0x09, 0x27, 0xf1,
	0x82, 0xf5, 0x6d, 0xf4, 0x9b, 0xc0, 0xd2, 0xcc, 0x7f, 0x46, 0x25, 0xf2, 0x4c, 0x6a, 0x85, 0xac,
	0xe2, 0x7c, 0xdb, 0x08, 0x81, 0x5c, 0x19, 0xf5, 0x29, 0x1b, 0x42, 0xf4, 0x02, 0x26, 0x1c, 0x55,
	0x30, 0x0e, 0x27, 0xb1, 0xb7, 0x79, 0x61, 0x53, 0x1d, 0x5e, 0xb1, 0x3e, 0xa0, 0x4a, 0xb8, 0x12,
	0x2d, 0xd3, 0x73, 0xf4, 0x19, 0xcc, 0x9a, 0x5a, 0xe5, 0x25, 0x9a, 0x2c, 0xa7, 0xac, 0xeb, 0x56,
	0x97, 0xe0, 0xf6, 0x83, 0xd4, 0x87, 0xc9, 0x3d, 0xb6, 0x5d, 0x28, 0xba, 0xa4, 0x67, 0xe0, 0x3c,
	0xa4, 0x45, 0x83, 0x66, 0x45, 0x6f, 0xe3, 0x5b, 0x99, 0x03, 0xaa, 0x4e, 0x84, 0x59, 0xfa, 0xdd,
	0xf8, 0x2d, 0x89, 0xde, 0x03, 0x3c, 0x11, 0x7a, 0x57, 0xf5, 0xf3, 0x43, 0xab, 0x50, 0x76, 0xe6,
	0xfb, 0x56, 0x33, 0xa2, 0x63, 0xc6, 0x96, 0xe9, 0xda, 0xf3, 0x57, 0xe0, 0xf6, 0xef, 0x82, 0x02,
	0xcc, 0xae, 0x12, 0xf6, 0x2d, 0x61, 0xfe, 0x48, 0xd7, 0xdb, 0x4f, 0xfb, 0xe4, 0x70, 0xed, 0x13,
	0xba, 0x04, 0x77, 0xb7, 0xbf, 0xda, 0x7e, 0xd1, 0xcc, 0x78, 0x73, 0x07, 0x8e, 0xd9, 0x99, 0xbe,
	0x86, 0x79, 0x17, 0x35, 0xf5, 0xac, 0xc9, 0xa4, 0xac, 0x55, 0xbb, 0x7a, 0x3e, 0x08, 0x66, 0xf8,
	0x2f, 0xa2, 0x11, 0x7d, 0x09, 0xf3, 0xde, 0xec, 0x5f, 0x47, 0xe8, 0xbf, 0x59, 0x46, 0xa3, 0xdb,
	0x99, 0x79, 0x4f, 0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x03, 0x51, 0xd4, 0x0b, 0x03,
	0x00, 0x00,
}
