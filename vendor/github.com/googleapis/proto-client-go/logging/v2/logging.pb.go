// Code generated by protoc-gen-go.
// source: google/logging/v2/logging.proto
// DO NOT EDIT!

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/googleapis/proto-client-go/api"
import google_api3 "github.com/googleapis/proto-client-go/api"
import google_protobuf4 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/googleapis/proto-client-go/rpc"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The parameters to DeleteLog.
type DeleteLogRequest struct {
	// Required. The resource name of the log to delete.  Example:
	// `"projects/my-project/logs/syslog"`.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
}

func (m *DeleteLogRequest) Reset()                    { *m = DeleteLogRequest{} }
func (m *DeleteLogRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteLogRequest) ProtoMessage()               {}
func (*DeleteLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

// The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Optional. A default log resource name for those log entries in `entries`
	// that do not specify their own `logName`.  Example:
	// `"projects/my-project/logs/syslog"`.  See
	// [LogEntry][google.logging.v2.LogEntry].
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName" json:"log_name,omitempty"`
	// Optional. A default monitored resource for those log entries in `entries`
	// that do not specify their own `resource`.
	Resource *google_api3.MonitoredResource `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
	// Optional. User-defined `key:value` items that are added to
	// the `labels` field of each log entry in `entries`, except when a log
	// entry specifies its own `key:value` item with the same key.
	// Example: `{ "size": "large", "color":"red" }`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Required. The log entries to write. The log entries must have values for
	// all required fields.
	Entries []*LogEntry `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
}

func (m *WriteLogEntriesRequest) Reset()                    { *m = WriteLogEntriesRequest{} }
func (m *WriteLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesRequest) ProtoMessage()               {}
func (*WriteLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *WriteLogEntriesRequest) GetResource() *google_api3.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WriteLogEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Result returned from WriteLogEntries.
type WriteLogEntriesResponse struct {
}

func (m *WriteLogEntriesResponse) Reset()                    { *m = WriteLogEntriesResponse{} }
func (m *WriteLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteLogEntriesResponse) ProtoMessage()               {}
func (*WriteLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

// The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Required. One or more project IDs or project numbers from which to retrieve
	// log entries.  Examples of a project ID: `"my-project-1A"`, `"1234567890"`.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds" json:"project_ids,omitempty"`
	// Optional. An [advanced logs filter](/logging/docs/view/advanced_filters).
	// The filter is compared against all log entries in the projects specified by
	// `projectIds`.  Only entries that match the filter are retrieved.  An empty
	// filter matches all log entries.
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// Optional. How the results should be sorted.  Presently, the only permitted
	// values are `"timestamp"` (default) and `"timestamp desc"`.  The first
	// option returns entries in order of increasing values of
	// `LogEntry.timestamp` (oldest first), and the second option returns entries
	// in order of decreasing timestamps (newest first).  Entries with equal
	// timestamps are returned in order of `LogEntry.insertId`.
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy" json:"order_by,omitempty"`
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.  The values of `projectIds`, `filter`, and `orderBy` must
	// be the same as in the previous request.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListLogEntriesRequest) Reset()                    { *m = ListLogEntriesRequest{} }
func (m *ListLogEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesRequest) ProtoMessage()               {}
func (*ListLogEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

// Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// A list of log entries.
	Entries []*LogEntry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// given a value in the response.  To get the next batch of results, call
	// this method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListLogEntriesResponse) Reset()                    { *m = ListLogEntriesResponse{} }
func (m *ListLogEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListLogEntriesResponse) ProtoMessage()               {}
func (*ListLogEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ListLogEntriesResponse) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// The parameters to ListMonitoredResourceDescriptors
type ListMonitoredResourceDescriptorsRequest struct {
	// Optional. The maximum number of results to return from this request.  Fewer
	// results might be returned. You must check for the `nextPageToken` result to
	// determine if additional results are available, which you can retrieve by
	// passing the `nextPageToken` value in the `pageToken` parameter to the next
	// request.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. If the `pageToken` request parameter is supplied, then the next
	// page of results in the set are retrieved.  The `pageToken` parameter must
	// be set with the value of the `nextPageToken` result parameter from the
	// previous request.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsRequest) Reset() {
	*m = ListMonitoredResourceDescriptorsRequest{}
}
func (m *ListMonitoredResourceDescriptorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsRequest) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{5}
}

// Result returned from ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// A list of resource descriptors.
	ResourceDescriptors []*google_api3.MonitoredResourceDescriptor `protobuf:"bytes,1,rep,name=resource_descriptors,json=resourceDescriptors" json:"resource_descriptors,omitempty"`
	// If there are more results than were returned, then `nextPageToken` is
	// returned in the response.  To get the next batch of results, call this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListMonitoredResourceDescriptorsResponse) Reset() {
	*m = ListMonitoredResourceDescriptorsResponse{}
}
func (m *ListMonitoredResourceDescriptorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMonitoredResourceDescriptorsResponse) ProtoMessage()    {}
func (*ListMonitoredResourceDescriptorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{6}
}

func (m *ListMonitoredResourceDescriptorsResponse) GetResourceDescriptors() []*google_api3.MonitoredResourceDescriptor {
	if m != nil {
		return m.ResourceDescriptors
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteLogRequest)(nil), "google.logging.v2.DeleteLogRequest")
	proto.RegisterType((*WriteLogEntriesRequest)(nil), "google.logging.v2.WriteLogEntriesRequest")
	proto.RegisterType((*WriteLogEntriesResponse)(nil), "google.logging.v2.WriteLogEntriesResponse")
	proto.RegisterType((*ListLogEntriesRequest)(nil), "google.logging.v2.ListLogEntriesRequest")
	proto.RegisterType((*ListLogEntriesResponse)(nil), "google.logging.v2.ListLogEntriesResponse")
	proto.RegisterType((*ListMonitoredResourceDescriptorsRequest)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsRequest")
	proto.RegisterType((*ListMonitoredResourceDescriptorsResponse)(nil), "google.logging.v2.ListMonitoredResourceDescriptorsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LoggingServiceV2 service

type LoggingServiceV2Client interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error)
	// Writes log entries to Cloud Logging.
	// All log entries in Cloud Logging are written by this method.
	//
	WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error)
	// Lists monitored resource descriptors that are used by Cloud Logging.
	ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error)
}

type loggingServiceV2Client struct {
	cc *grpc.ClientConn
}

func NewLoggingServiceV2Client(cc *grpc.ClientConn) LoggingServiceV2Client {
	return &loggingServiceV2Client{cc}
}

func (c *loggingServiceV2Client) DeleteLog(ctx context.Context, in *DeleteLogRequest, opts ...grpc.CallOption) (*google_protobuf4.Empty, error) {
	out := new(google_protobuf4.Empty)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/DeleteLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) WriteLogEntries(ctx context.Context, in *WriteLogEntriesRequest, opts ...grpc.CallOption) (*WriteLogEntriesResponse, error) {
	out := new(WriteLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/WriteLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListLogEntries(ctx context.Context, in *ListLogEntriesRequest, opts ...grpc.CallOption) (*ListLogEntriesResponse, error) {
	out := new(ListLogEntriesResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListLogEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loggingServiceV2Client) ListMonitoredResourceDescriptors(ctx context.Context, in *ListMonitoredResourceDescriptorsRequest, opts ...grpc.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
	out := new(ListMonitoredResourceDescriptorsResponse)
	err := grpc.Invoke(ctx, "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoggingServiceV2 service

type LoggingServiceV2Server interface {
	// Deletes a log and all its log entries.
	// The log will reappear if it receives new entries.
	//
	DeleteLog(context.Context, *DeleteLogRequest) (*google_protobuf4.Empty, error)
	// Writes log entries to Cloud Logging.
	// All log entries in Cloud Logging are written by this method.
	//
	WriteLogEntries(context.Context, *WriteLogEntriesRequest) (*WriteLogEntriesResponse, error)
	// Lists log entries.  Use this method to retrieve log entries from Cloud
	// Logging.  For ways to export log entries, see
	// [Exporting Logs](/logging/docs/export).
	//
	ListLogEntries(context.Context, *ListLogEntriesRequest) (*ListLogEntriesResponse, error)
	// Lists monitored resource descriptors that are used by Cloud Logging.
	ListMonitoredResourceDescriptors(context.Context, *ListMonitoredResourceDescriptorsRequest) (*ListMonitoredResourceDescriptorsResponse, error)
}

func RegisterLoggingServiceV2Server(s *grpc.Server, srv LoggingServiceV2Server) {
	s.RegisterService(&_LoggingServiceV2_serviceDesc, srv)
}

func _LoggingServiceV2_DeleteLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/DeleteLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).DeleteLog(ctx, req.(*DeleteLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_WriteLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/WriteLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).WriteLogEntries(ctx, req.(*WriteLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListLogEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListLogEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListLogEntries(ctx, req.(*ListLogEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMonitoredResourceDescriptorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.logging.v2.LoggingServiceV2/ListMonitoredResourceDescriptors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceV2Server).ListMonitoredResourceDescriptors(ctx, req.(*ListMonitoredResourceDescriptorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingServiceV2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.logging.v2.LoggingServiceV2",
	HandlerType: (*LoggingServiceV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteLog",
			Handler:    _LoggingServiceV2_DeleteLog_Handler,
		},
		{
			MethodName: "WriteLogEntries",
			Handler:    _LoggingServiceV2_WriteLogEntries_Handler,
		},
		{
			MethodName: "ListLogEntries",
			Handler:    _LoggingServiceV2_ListLogEntries_Handler,
		},
		{
			MethodName: "ListMonitoredResourceDescriptors",
			Handler:    _LoggingServiceV2_ListMonitoredResourceDescriptors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("google/logging/v2/logging.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xd6, 0x24, 0xfd, 0xcb, 0xe9, 0xbd, 0xb7, 0xbd, 0x43, 0x9b, 0xa6, 0x0e, 0x55, 0x53, 0x57,
	0xd0, 0x34, 0x52, 0x6c, 0x08, 0xaa, 0x44, 0x83, 0xd8, 0x54, 0xed, 0x02, 0x29, 0x45, 0x95, 0x8b,
	0x40, 0xaa, 0x90, 0x22, 0x27, 0x39, 0x35, 0xa6, 0x8e, 0x27, 0xd8, 0x93, 0x94, 0x80, 0xd8, 0xb0,
	0x61, 0xc1, 0x92, 0x87, 0x60, 0x07, 0xef, 0xc1, 0x96, 0x57, 0xe0, 0x01, 0x58, 0xb2, 0x64, 0x6c,
	0x8f, 0xd3, 0x34, 0x0e, 0x6d, 0xc4, 0x2a, 0x39, 0xe7, 0x7c, 0xe7, 0xe7, 0x3b, 0x3f, 0x63, 0x58,
	0xb7, 0x18, 0xb3, 0x1c, 0xd4, 0x1d, 0x66, 0x59, 0xb6, 0x6b, 0xe9, 0xbd, 0x4a, 0xfc, 0x57, 0xeb,
	0x78, 0x8c, 0x33, 0xfa, 0x7f, 0x04, 0xd0, 0x62, 0x6d, 0xaf, 0xa2, 0xdc, 0x94, 0x3e, 0x66, 0xc7,
	0xd6, 0x4d, 0xd7, 0x65, 0xdc, 0xe4, 0x36, 0x73, 0xfd, 0xc8, 0x41, 0xd9, 0x1c, 0xb2, 0xb6, 0x99,
	0x6b, 0x73, 0xe6, 0x61, 0xab, 0xee, 0xa1, 0xcf, 0xba, 0x5e, 0x13, 0x25, 0x68, 0x63, 0x6c, 0xda,
	0x3a, 0xba, 0xdc, 0xeb, 0x4b, 0x48, 0x5e, 0x42, 0x42, 0xa9, 0xd1, 0x3d, 0xd5, 0xb1, 0xdd, 0xe1,
	0xb1, 0x71, 0x45, 0x1a, 0xbd, 0x4e, 0x53, 0xf7, 0x45, 0xfe, 0xae, 0xcc, 0xae, 0x96, 0x61, 0x71,
	0x1f, 0x1d, 0xe4, 0x58, 0x63, 0x96, 0x81, 0xaf, 0xba, 0xe8, 0x73, 0xba, 0x0a, 0x73, 0x41, 0x70,
	0xd7, 0x6c, 0x63, 0x8e, 0x14, 0x48, 0x31, 0x63, 0xcc, 0x0a, 0xf9, 0xb1, 0x10, 0xd5, 0xaf, 0x29,
	0xc8, 0x3e, 0xf3, 0xec, 0x10, 0x7e, 0x20, 0x92, 0xdb, 0xe8, 0x5f, 0xef, 0x45, 0x77, 0x61, 0x2e,
	0xe6, 0x93, 0x4b, 0x09, 0xd3, 0x7c, 0x65, 0x4d, 0x93, 0x6d, 0x12, 0xac, 0xb5, 0xc3, 0x98, 0xb5,
	0x21, 0x41, 0xc6, 0x00, 0x4e, 0x0f, 0x61, 0xc6, 0x31, 0x1b, 0xe8, 0xf8, 0xb9, 0x74, 0x21, 0x2d,
	0x1c, 0x77, 0xb4, 0x44, 0x7f, 0xb5, 0xf1, 0x05, 0x69, 0xb5, 0xd0, 0x2f, 0x50, 0xf6, 0x0d, 0x19,
	0x84, 0xee, 0xc0, 0x2c, 0x46, 0xa8, 0xdc, 0x54, 0x18, 0x2f, 0x3f, 0x26, 0x9e, 0x0c, 0xd5, 0x37,
	0x62, 0xac, 0xb2, 0x0b, 0xf3, 0x43, 0xd1, 0xe8, 0x22, 0xa4, 0xcf, 0xb0, 0x2f, 0x59, 0x06, 0x7f,
	0xe9, 0x12, 0x4c, 0xf7, 0x4c, 0xa7, 0x1b, 0xd1, 0xcb, 0x18, 0x91, 0x50, 0x4d, 0xdd, 0x27, 0xea,
	0x2a, 0xac, 0x24, 0xea, 0xf3, 0x3b, 0x62, 0xfc, 0xa8, 0x7e, 0x26, 0xb0, 0x5c, 0xb3, 0x7d, 0x9e,
	0xec, 0xe5, 0x3a, 0xcc, 0x8b, 0xf1, 0xbc, 0xc4, 0x26, 0xaf, 0xdb, 0x2d, 0x5f, 0x24, 0x4a, 0x8b,
	0xa0, 0x20, 0x55, 0x8f, 0x5a, 0x3e, 0xcd, 0xc2, 0xcc, 0xa9, 0xed, 0x70, 0xf4, 0x64, 0x42, 0x29,
	0x05, 0x43, 0x60, 0x5e, 0x0b, 0xbd, 0x7a, 0xa3, 0x2f, 0x1a, 0x16, 0x0e, 0x21, 0x94, 0xf7, 0xfa,
	0x34, 0x0f, 0x99, 0x8e, 0x69, 0x61, 0xdd, 0xb7, 0xdf, 0xa0, 0x20, 0x4f, 0x8a, 0xd3, 0xc6, 0x5c,
	0xa0, 0x38, 0x16, 0x32, 0x5d, 0x03, 0x08, 0x8d, 0x9c, 0x9d, 0xa1, 0x9b, 0x9b, 0x0e, 0x3d, 0x43,
	0xf8, 0x93, 0x40, 0xa1, 0x9e, 0x43, 0x76, 0xb4, 0xd0, 0x88, 0xc3, 0x70, 0x43, 0xc9, 0xe4, 0x0d,
	0xa5, 0xb7, 0x61, 0xc1, 0xc5, 0xd7, 0xbc, 0x3e, 0x94, 0x34, 0x22, 0xf2, 0x6f, 0xa0, 0x3e, 0x1a,
	0x24, 0x46, 0xd8, 0x0a, 0x12, 0x27, 0x36, 0x64, 0x1f, 0xfd, 0xa6, 0x67, 0x77, 0x84, 0x6e, 0xd0,
	0xb3, 0x4b, 0xfc, 0xc8, 0x95, 0xfc, 0x52, 0xa3, 0xfc, 0xbe, 0x10, 0x28, 0x5e, 0x9f, 0x47, 0x52,
	0x3e, 0x81, 0xa5, 0x78, 0x3d, 0xeb, 0xad, 0x0b, 0xbb, 0xe4, 0xbf, 0x75, 0xe5, 0x66, 0x5f, 0xc4,
	0x33, 0x6e, 0x78, 0xc9, 0x1c, 0x93, 0xf6, 0xa5, 0xf2, 0x73, 0x0a, 0x16, 0x6b, 0x51, 0x83, 0x8f,
	0xd1, 0xeb, 0xd9, 0x4d, 0x7c, 0x5a, 0xa1, 0xe7, 0x90, 0x19, 0xdc, 0x32, 0xdd, 0x1c, 0x33, 0x87,
	0xd1, 0x4b, 0x57, 0xb2, 0x31, 0x28, 0x7e, 0x34, 0xb4, 0x83, 0xe0, 0xd1, 0x50, 0xcb, 0xef, 0xbf,
	0xff, 0xf8, 0x94, 0xda, 0x2a, 0xdd, 0x12, 0x0f, 0x4d, 0x03, 0xb9, 0x79, 0x57, 0x7f, 0x1b, 0xdf,
	0xf6, 0x43, 0xb9, 0x85, 0xbe, 0x5e, 0x0a, 0x9e, 0x20, 0xf1, 0xf3, 0x8e, 0x7e, 0x24, 0xb0, 0x30,
	0xb2, 0xe4, 0x74, 0x7b, 0xe2, 0x43, 0x55, 0x4a, 0x93, 0x40, 0xe5, 0xcd, 0x6c, 0x84, 0x95, 0xe5,
	0xd5, 0xec, 0xa0, 0x32, 0xb9, 0x52, 0xd5, 0xf3, 0xc0, 0xa3, 0x4a, 0x4a, 0xf4, 0x03, 0x81, 0xff,
	0x2e, 0x6f, 0x2b, 0x2d, 0x8e, 0x5b, 0xca, 0x71, 0x97, 0xa7, 0x6c, 0x4f, 0x80, 0x94, 0xa5, 0x14,
	0xc2, 0x52, 0x14, 0x75, 0x39, 0x51, 0x8a, 0x23, 0x1c, 0x82, 0x4a, 0xbe, 0x11, 0x28, 0x5c, 0xb7,
	0x56, 0xb4, 0xfa, 0x87, 0x8c, 0x13, 0xec, 0xbc, 0xf2, 0xe0, 0xaf, 0x7c, 0x65, 0xfd, 0x72, 0xc8,
	0xf4, 0x62, 0xc8, 0xed, 0x2b, 0xdc, 0xf6, 0x9e, 0xc3, 0x72, 0x93, 0xb5, 0x93, 0x09, 0xf7, 0xfe,
	0x91, 0x8b, 0x78, 0x14, 0xec, 0xd0, 0x11, 0x39, 0xb9, 0x63, 0xd9, 0xfc, 0x45, 0xb7, 0xa1, 0x09,
	0xb4, 0x1e, 0xa1, 0xc5, 0x29, 0xf8, 0xd1, 0x77, 0xa9, 0xdc, 0x74, 0x6c, 0xd1, 0xa5, 0xb2, 0xc5,
	0x86, 0x3e, 0x65, 0xbf, 0x08, 0x69, 0xcc, 0x84, 0xe6, 0x7b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x5e, 0x23, 0xb8, 0x60, 0x07, 0x00, 0x00,
}
