// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo_indexer.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IndexRepoRequest struct {
	RepoId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
}

func (m *IndexRepoRequest) Reset()                    { *m = IndexRepoRequest{} }
func (m *IndexRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*IndexRepoRequest) ProtoMessage()               {}
func (*IndexRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *IndexRepoRequest) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type IndexRepoResponse struct {
	RepoTask *RepoTask `protobuf:"bytes,1,opt,name=repo_task,json=repoTask" json:"repo_task,omitempty"`
}

func (m *IndexRepoResponse) Reset()                    { *m = IndexRepoResponse{} }
func (m *IndexRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*IndexRepoResponse) ProtoMessage()               {}
func (*IndexRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *IndexRepoResponse) GetRepoTask() *RepoTask {
	if m != nil {
		return m.RepoTask
	}
	return nil
}

type RepoTask struct {
	RepoTaskId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=repo_task_id,json=repoTaskId" json:"repo_task_id,omitempty"`
	RepoId     *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Status     *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Result     *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=result" json:"result,omitempty"`
	CreateTime *google_protobuf1.Timestamp  `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime *google_protobuf1.Timestamp  `protobuf:"bytes,6,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *RepoTask) Reset()                    { *m = RepoTask{} }
func (m *RepoTask) String() string            { return proto.CompactTextString(m) }
func (*RepoTask) ProtoMessage()               {}
func (*RepoTask) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *RepoTask) GetRepoTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoTaskId
	}
	return nil
}

func (m *RepoTask) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *RepoTask) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RepoTask) GetResult() *google_protobuf.StringValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *RepoTask) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *RepoTask) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeRepoTasksRequest struct {
	RepoTaskId []string `protobuf:"bytes,1,rep,name=repo_task_id,json=repoTaskId" json:"repo_task_id,omitempty"`
	RepoId     []string `protobuf:"bytes,2,rep,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Status     []string `protobuf:"bytes,3,rep,name=status" json:"status,omitempty"`
	Limit      uint32   `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Offset     uint32   `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeRepoTasksRequest) Reset()                    { *m = DescribeRepoTasksRequest{} }
func (m *DescribeRepoTasksRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRepoTasksRequest) ProtoMessage()               {}
func (*DescribeRepoTasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DescribeRepoTasksRequest) GetRepoTaskId() []string {
	if m != nil {
		return m.RepoTaskId
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetRepoId() []string {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeRepoTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeRepoTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeRepoTasksResponse struct {
	TotalCount  uint32      `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	RepoTaskSet []*RepoTask `protobuf:"bytes,2,rep,name=repo_task_set,json=repoTaskSet" json:"repo_task_set,omitempty"`
}

func (m *DescribeRepoTasksResponse) Reset()                    { *m = DescribeRepoTasksResponse{} }
func (m *DescribeRepoTasksResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeRepoTasksResponse) ProtoMessage()               {}
func (*DescribeRepoTasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *DescribeRepoTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeRepoTasksResponse) GetRepoTaskSet() []*RepoTask {
	if m != nil {
		return m.RepoTaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexRepoRequest)(nil), "openpitrix.IndexRepoRequest")
	proto.RegisterType((*IndexRepoResponse)(nil), "openpitrix.IndexRepoResponse")
	proto.RegisterType((*RepoTask)(nil), "openpitrix.RepoTask")
	proto.RegisterType((*DescribeRepoTasksRequest)(nil), "openpitrix.DescribeRepoTasksRequest")
	proto.RegisterType((*DescribeRepoTasksResponse)(nil), "openpitrix.DescribeRepoTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepoIndexer service

type RepoIndexerClient interface {
	IndexRepo(ctx context.Context, in *IndexRepoRequest, opts ...grpc.CallOption) (*IndexRepoResponse, error)
	DescribeRepoTasks(ctx context.Context, in *DescribeRepoTasksRequest, opts ...grpc.CallOption) (*DescribeRepoTasksResponse, error)
}

type repoIndexerClient struct {
	cc *grpc.ClientConn
}

func NewRepoIndexerClient(cc *grpc.ClientConn) RepoIndexerClient {
	return &repoIndexerClient{cc}
}

func (c *repoIndexerClient) IndexRepo(ctx context.Context, in *IndexRepoRequest, opts ...grpc.CallOption) (*IndexRepoResponse, error) {
	out := new(IndexRepoResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoIndexer/IndexRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoIndexerClient) DescribeRepoTasks(ctx context.Context, in *DescribeRepoTasksRequest, opts ...grpc.CallOption) (*DescribeRepoTasksResponse, error) {
	out := new(DescribeRepoTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoIndexer/DescribeRepoTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepoIndexer service

type RepoIndexerServer interface {
	IndexRepo(context.Context, *IndexRepoRequest) (*IndexRepoResponse, error)
	DescribeRepoTasks(context.Context, *DescribeRepoTasksRequest) (*DescribeRepoTasksResponse, error)
}

func RegisterRepoIndexerServer(s *grpc.Server, srv RepoIndexerServer) {
	s.RegisterService(&_RepoIndexer_serviceDesc, srv)
}

func _RepoIndexer_IndexRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoIndexerServer).IndexRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoIndexer/IndexRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoIndexerServer).IndexRepo(ctx, req.(*IndexRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoIndexer_DescribeRepoTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRepoTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoIndexerServer).DescribeRepoTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoIndexer/DescribeRepoTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoIndexerServer).DescribeRepoTasks(ctx, req.(*DescribeRepoTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoIndexer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.RepoIndexer",
	HandlerType: (*RepoIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexRepo",
			Handler:    _RepoIndexer_IndexRepo_Handler,
		},
		{
			MethodName: "DescribeRepoTasks",
			Handler:    _RepoIndexer_DescribeRepoTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo_indexer.proto",
}

func init() { proto.RegisterFile("repo_indexer.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xd3, 0x36, 0x34, 0x13, 0x02, 0xed, 0x12, 0x68, 0x1a, 0x05, 0x6a, 0x59, 0x20, 0x55,
	0x88, 0xda, 0x6a, 0xf9, 0x11, 0x02, 0x09, 0xa9, 0x80, 0x90, 0x72, 0x75, 0x2b, 0x0e, 0x5c, 0xa2,
	0x4d, 0x32, 0xb1, 0xac, 0x3a, 0xde, 0x65, 0x77, 0xdd, 0xf6, 0x88, 0x78, 0x02, 0x54, 0x8e, 0x1c,
	0x79, 0x24, 0x5e, 0x01, 0x89, 0xd7, 0x40, 0xbb, 0xeb, 0x6d, 0x4c, 0x43, 0xd4, 0x9c, 0x92, 0x9d,
	0xf9, 0xbe, 0x99, 0x6f, 0xbe, 0x9d, 0x35, 0x10, 0x81, 0x9c, 0x0d, 0xd2, 0x7c, 0x8c, 0xe7, 0x28,
	0x42, 0x2e, 0x98, 0x62, 0x04, 0x18, 0xc7, 0x9c, 0xa7, 0x4a, 0xa4, 0xe7, 0xdd, 0x07, 0x09, 0x63,
	0x49, 0x86, 0x91, 0xc9, 0x0c, 0x8b, 0x49, 0x74, 0x26, 0x28, 0xe7, 0x28, 0xa4, 0xc5, 0x76, 0x77,
	0xae, 0xe6, 0x55, 0x3a, 0x45, 0xa9, 0xe8, 0x94, 0x97, 0x80, 0x5e, 0x09, 0xa0, 0x3c, 0x8d, 0x68,
	0x9e, 0x33, 0x45, 0x55, 0xca, 0x72, 0x47, 0x7f, 0x62, 0x7e, 0x46, 0x7b, 0x09, 0xe6, 0x7b, 0xf2,
	0x8c, 0x26, 0x09, 0x8a, 0x88, 0x71, 0x83, 0x98, 0x47, 0x07, 0x7d, 0xd8, 0xe8, 0x6b, 0xa5, 0x31,
	0x72, 0x16, 0xe3, 0xe7, 0x02, 0xa5, 0x22, 0xcf, 0xe1, 0x86, 0x1d, 0x61, 0xdc, 0xf1, 0x7c, 0x6f,
	0xb7, 0x79, 0xd0, 0x0b, 0x6d, 0xc7, 0xd0, 0x49, 0x0a, 0x8f, 0x94, 0x48, 0xf3, 0xe4, 0x23, 0xcd,
	0x0a, 0x8c, 0xeb, 0x1a, 0xdc, 0x1f, 0x07, 0x1f, 0x60, 0xb3, 0x52, 0x4a, 0x72, 0x96, 0x4b, 0x24,
	0xfb, 0xd0, 0x30, 0xb5, 0x14, 0x95, 0x27, 0x65, 0xb5, 0x76, 0x38, 0x33, 0x23, 0xd4, 0xe0, 0x63,
	0x2a, 0x4f, 0xe2, 0x75, 0x51, 0xfe, 0x0b, 0xfe, 0xd4, 0x60, 0xdd, 0x85, 0xc9, 0x1b, 0xb8, 0x79,
	0xc9, 0x5f, 0x56, 0x10, 0xb8, 0x52, 0xfd, 0x71, 0x75, 0x96, 0xda, 0xf2, 0xb3, 0x90, 0x67, 0x50,
	0x97, 0x8a, 0xaa, 0x42, 0x76, 0x56, 0x96, 0x61, 0x59, 0xac, 0x66, 0x09, 0x94, 0x45, 0xa6, 0x3a,
	0xab, 0xcb, 0xf5, 0xd2, 0x58, 0xf2, 0x1a, 0x9a, 0x23, 0x81, 0x54, 0xe1, 0x40, 0x5f, 0x74, 0x67,
	0xcd, 0x50, 0xbb, 0x73, 0xd4, 0x63, 0xb7, 0x05, 0x31, 0x58, 0xb8, 0x0e, 0x68, 0xb2, 0x6d, 0x6e,
	0xc9, 0xf5, 0xeb, 0xc9, 0x16, 0xae, 0x03, 0xc1, 0x0f, 0x0f, 0x3a, 0xef, 0x51, 0x8e, 0x44, 0x3a,
	0x44, 0xe7, 0xb8, 0x74, 0x5b, 0xe0, 0xcf, 0x39, 0xbf, 0xb2, 0xdb, 0xf8, 0xc7, 0xdb, 0xad, 0xaa,
	0xb7, 0x3a, 0xe9, 0xdc, 0xbb, 0x57, 0x71, 0xcf, 0xc4, 0x4b, 0x7f, 0xda, 0xb0, 0x96, 0xa5, 0xd3,
	0xd4, 0xda, 0xd3, 0x8a, 0xed, 0x41, 0xa3, 0xd9, 0x64, 0x22, 0x51, 0x99, 0xd1, 0x5b, 0x71, 0x79,
	0x0a, 0x4e, 0x61, 0xfb, 0x3f, 0xe2, 0xca, 0xbd, 0xda, 0x81, 0xa6, 0x62, 0x8a, 0x66, 0x83, 0x11,
	0x2b, 0x72, 0x65, 0xd6, 0xa2, 0x15, 0x83, 0x09, 0xbd, 0xd3, 0x11, 0xf2, 0x12, 0x5a, 0x33, 0xf9,
	0xba, 0xb8, 0x96, 0xb8, 0x68, 0xf9, 0x9a, 0x6e, 0xaa, 0x23, 0x54, 0x07, 0x3f, 0x6b, 0xd0, 0xd4,
	0x99, 0xbe, 0x7d, 0xc1, 0xe4, 0x8b, 0x07, 0x8d, 0xcb, 0xc5, 0x26, 0xbd, 0x6a, 0x81, 0xab, 0x4f,
	0xa7, 0x7b, 0x7f, 0x41, 0xd6, 0xaa, 0x0e, 0x5e, 0x5c, 0x1c, 0x6e, 0x93, 0x2d, 0xa9, 0xa8, 0x50,
	0x3e, 0xf5, 0xcd, 0x27, 0xc2, 0xd7, 0xad, 0x7d, 0xad, 0xf2, 0xeb, 0xaf, 0xdf, 0xdf, 0x6b, 0xed,
	0xe0, 0x76, 0x74, 0xba, 0x1f, 0xe9, 0xa0, 0x8c, 0x0c, 0xe0, 0x95, 0xf7, 0x98, 0x7c, 0xf3, 0x60,
	0x73, 0xce, 0x0b, 0xf2, 0xb0, 0xda, 0x6c, 0xd1, 0x3d, 0x76, 0x1f, 0x5d, 0x83, 0x2a, 0xa5, 0x85,
	0x17, 0x87, 0x77, 0xc9, 0x9d, 0x71, 0x99, 0x9f, 0xa9, 0x92, 0x46, 0xd6, 0x06, 0xb9, 0xe5, 0x64,
	0x19, 0x47, 0xe5, 0xdb, 0xd5, 0x4f, 0x35, 0x3e, 0x1c, 0xd6, 0xcd, 0x86, 0x3d, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0x73, 0x1d, 0xae, 0xf4, 0x04, 0x00, 0x00,
}
