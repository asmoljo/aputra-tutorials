// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: event/event.proto

package event

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version         string          `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	RouteKey        string          `protobuf:"bytes,2,opt,name=route_key,json=routeKey,proto3" json:"route_key,omitempty"`
	RawPath         string          `protobuf:"bytes,3,opt,name=raw_path,json=rawPath,proto3" json:"raw_path,omitempty"`
	RawQueryString  string          `protobuf:"bytes,4,opt,name=raw_query_string,json=rawQueryString,proto3" json:"raw_query_string,omitempty"`
	Headers         *Headers        `protobuf:"bytes,5,opt,name=headers,proto3" json:"headers,omitempty"`
	RequestContext  *RequestContext `protobuf:"bytes,6,opt,name=request_context,json=requestContext,proto3" json:"request_context,omitempty"`
	Body            string          `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	IsBase64Encoded bool            `protobuf:"varint,8,opt,name=is_base64_encoded,json=isBase64Encoded,proto3" json:"is_base64_encoded,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Event) GetRouteKey() string {
	if x != nil {
		return x.RouteKey
	}
	return ""
}

func (x *Event) GetRawPath() string {
	if x != nil {
		return x.RawPath
	}
	return ""
}

func (x *Event) GetRawQueryString() string {
	if x != nil {
		return x.RawQueryString
	}
	return ""
}

func (x *Event) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Event) GetRequestContext() *RequestContext {
	if x != nil {
		return x.RequestContext
	}
	return nil
}

func (x *Event) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Event) GetIsBase64Encoded() bool {
	if x != nil {
		return x.IsBase64Encoded
	}
	return false
}

type Headers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accept                 string `protobuf:"bytes,1,opt,name=accept,proto3" json:"accept,omitempty"`
	AcceptEncoding         string `protobuf:"bytes,2,opt,name=accept_encoding,json=acceptEncoding,proto3" json:"accept_encoding,omitempty"`
	ContentLength          string `protobuf:"bytes,3,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	ContentType            string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Host                   string `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	UserAgent              string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	XAmznTraceId           string `protobuf:"bytes,7,opt,name=x_amzn_trace_id,json=xAmznTraceId,proto3" json:"x_amzn_trace_id,omitempty"`
	XForwardedFor          string `protobuf:"bytes,8,opt,name=x_forwarded_for,json=xForwardedFor,proto3" json:"x_forwarded_for,omitempty"`
	XForwardedPort         string `protobuf:"bytes,9,opt,name=x_forwarded_port,json=xForwardedPort,proto3" json:"x_forwarded_port,omitempty"`
	XForwardedProto        string `protobuf:"bytes,10,opt,name=x_forwarded_proto,json=xForwardedProto,proto3" json:"x_forwarded_proto,omitempty"`
	XSlackRequestTimestamp string `protobuf:"bytes,11,opt,name=x_slack_request_timestamp,json=xSlackRequestTimestamp,proto3" json:"x_slack_request_timestamp,omitempty"`
	XSlackSignature        string `protobuf:"bytes,12,opt,name=x_slack_signature,json=xSlackSignature,proto3" json:"x_slack_signature,omitempty"`
}

func (x *Headers) Reset() {
	*x = Headers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Headers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Headers) ProtoMessage() {}

func (x *Headers) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Headers.ProtoReflect.Descriptor instead.
func (*Headers) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *Headers) GetAccept() string {
	if x != nil {
		return x.Accept
	}
	return ""
}

func (x *Headers) GetAcceptEncoding() string {
	if x != nil {
		return x.AcceptEncoding
	}
	return ""
}

func (x *Headers) GetContentLength() string {
	if x != nil {
		return x.ContentLength
	}
	return ""
}

func (x *Headers) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Headers) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Headers) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Headers) GetXAmznTraceId() string {
	if x != nil {
		return x.XAmznTraceId
	}
	return ""
}

func (x *Headers) GetXForwardedFor() string {
	if x != nil {
		return x.XForwardedFor
	}
	return ""
}

func (x *Headers) GetXForwardedPort() string {
	if x != nil {
		return x.XForwardedPort
	}
	return ""
}

func (x *Headers) GetXForwardedProto() string {
	if x != nil {
		return x.XForwardedProto
	}
	return ""
}

func (x *Headers) GetXSlackRequestTimestamp() string {
	if x != nil {
		return x.XSlackRequestTimestamp
	}
	return ""
}

func (x *Headers) GetXSlackSignature() string {
	if x != nil {
		return x.XSlackSignature
	}
	return ""
}

type RequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId    string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	ApiId        string `protobuf:"bytes,2,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	DomainName   string `protobuf:"bytes,3,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	DomainPrefix string `protobuf:"bytes,4,opt,name=domain_prefix,json=domainPrefix,proto3" json:"domain_prefix,omitempty"`
	Http         *Http  `protobuf:"bytes,5,opt,name=http,proto3" json:"http,omitempty"`
	RequestId    string `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	RouteKey     string `protobuf:"bytes,7,opt,name=route_key,json=routeKey,proto3" json:"route_key,omitempty"`
	Stage        string `protobuf:"bytes,8,opt,name=stage,proto3" json:"stage,omitempty"`
	Time         string `protobuf:"bytes,9,opt,name=time,proto3" json:"time,omitempty"`
	TimeEpoch    int64  `protobuf:"varint,10,opt,name=time_epoch,json=timeEpoch,proto3" json:"time_epoch,omitempty"`
}

func (x *RequestContext) Reset() {
	*x = RequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestContext) ProtoMessage() {}

func (x *RequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestContext.ProtoReflect.Descriptor instead.
func (*RequestContext) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *RequestContext) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *RequestContext) GetApiId() string {
	if x != nil {
		return x.ApiId
	}
	return ""
}

func (x *RequestContext) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *RequestContext) GetDomainPrefix() string {
	if x != nil {
		return x.DomainPrefix
	}
	return ""
}

func (x *RequestContext) GetHttp() *Http {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *RequestContext) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *RequestContext) GetRouteKey() string {
	if x != nil {
		return x.RouteKey
	}
	return ""
}

func (x *RequestContext) GetStage() string {
	if x != nil {
		return x.Stage
	}
	return ""
}

func (x *RequestContext) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *RequestContext) GetTimeEpoch() int64 {
	if x != nil {
		return x.TimeEpoch
	}
	return 0
}

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method    string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Protocol  string `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SourcePp  string `protobuf:"bytes,4,opt,name=source_pp,json=sourcePp,proto3" json:"source_pp,omitempty"`
	UserAgent string `protobuf:"bytes,5,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *Http) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Http) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Http) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Http) GetSourcePp() string {
	if x != nil {
		return x.SourcePp
	}
	return ""
}

func (x *Http) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_event_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_event_event_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_event_event_proto protoreflect.FileDescriptor

var file_event_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xad, 0x02, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x61, 0x77, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x61, 0x77, 0x50, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x77, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x28, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2a,
	0x0a, 0x11, 0x69, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x5f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x42, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x22, 0xd3, 0x03, 0x0a, 0x07, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0f, 0x78, 0x5f, 0x61, 0x6d, 0x7a, 0x6e, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x78,
	0x41, 0x6d, 0x7a, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x78,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x78, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x78, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x78,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x78, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x78, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x19, 0x78, 0x5f, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x78, 0x53,
	0x6c, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x78, 0x5f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x78, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x22, 0xb2, 0x02, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x69, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x1f, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65,
	0x70, 0x6f, 0x63, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x48, 0x74, 0x74, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x22, 0x19, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x35, 0x0a,
	0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x70, 0x75, 0x74, 0x72, 0x61, 0x2f, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x2f,
	0x31, 0x34, 0x39, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_event_proto_rawDescOnce sync.Once
	file_event_event_proto_rawDescData = file_event_event_proto_rawDesc
)

func file_event_event_proto_rawDescGZIP() []byte {
	file_event_event_proto_rawDescOnce.Do(func() {
		file_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_event_proto_rawDescData)
	})
	return file_event_event_proto_rawDescData
}

var file_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_event_event_proto_goTypes = []interface{}{
	(*Event)(nil),          // 0: event.Event
	(*Headers)(nil),        // 1: event.Headers
	(*RequestContext)(nil), // 2: event.RequestContext
	(*Http)(nil),           // 3: event.Http
	(*Request)(nil),        // 4: event.Request
}
var file_event_event_proto_depIdxs = []int32{
	1, // 0: event.Event.headers:type_name -> event.Headers
	2, // 1: event.Event.request_context:type_name -> event.RequestContext
	3, // 2: event.RequestContext.http:type_name -> event.Http
	4, // 3: event.Manager.GetEvent:input_type -> event.Request
	0, // 4: event.Manager.GetEvent:output_type -> event.Event
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_event_proto_init() }
func file_event_event_proto_init() {
	if File_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Headers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestContext); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_event_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_event_proto_goTypes,
		DependencyIndexes: file_event_event_proto_depIdxs,
		MessageInfos:      file_event_event_proto_msgTypes,
	}.Build()
	File_event_event_proto = out.File
	file_event_event_proto_rawDesc = nil
	file_event_event_proto_goTypes = nil
	file_event_event_proto_depIdxs = nil
}
