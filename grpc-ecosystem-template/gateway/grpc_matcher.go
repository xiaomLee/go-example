package gateway

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const (
	_minBusinessErrorCode = 999
)

// IncomingMatcher 转换 http header 为 grpc metadata
func IncomingMatcher(key string) (string, bool) {
	switch key {
	case HttpHeaderRequestId:
		return GrpcMetadataRequestId, true
	case HttpHeaderRequestTimeout:
		return GrpcMetadataRequestTimeout, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

// OutgoingMatcher ...
func OutgoingMatcher(key string) (string, bool) {
	switch key {
	case GrpcMetadataRequestId:
		return HttpHeaderRequestId, true
	case GrpcMetadataRequestTimeout:
		return HttpHeaderRequestTimeout, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

type errorBody struct {
	Error string `protobuf:"bytes,100,name=error" json:"error"`
	// This is to make the error more compatible with users that expect errors to be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	Code      int32      `protobuf:"varint,1,name=code" json:"code"`
	Message   string     `protobuf:"bytes,2,name=message" json:"message"`
	Details   []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
	RequestID string     `protobuf:"bytes,4,name=request_id" json:"request_id,omitempty"`
}

// Reset Make this also conform to proto.Message for builtin JSONPb Marshaler
func (e *errorBody) Reset()         { *e = errorBody{} }
func (e *errorBody) String() string { return proto.CompactTextString(e) }
func (*errorBody) ProtoMessage()    {}

// CustomHTTPError 自定义 grpc-runtime 的错误处理方法
func CustomHTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	pb := s.Proto()
	contentType := marshaler.ContentType(s.Proto())
	// Check marshaler on run time in order to keep backwards compatability
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		contentType = httpBodyMarshaler.ContentType(pb)
	}
	w.Header().Set("Content-Type", contentType)

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(w, mux, md)

	body := &errorBody{
		Error:     s.Message(),
		Message:   s.Message(),
		Code:      int32(s.Code()),
		Details:   pb.GetDetails(),
		RequestID: w.Header().Get(HttpHeaderRequestId),
	}

	buf, merr := marshaler.Marshal(body)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", body, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	handleForwardResponseTrailerHeader(w, md)
	st := runtime.HTTPStatusFromCode(s.Code())
	if s.Code() > _minBusinessErrorCode {
		st = http.StatusOK
	}
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	handleForwardResponseTrailer(w, md)
}

func handleForwardResponseTrailer(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k, vs := range md.TrailerMD {
		tKey := fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k)
		for _, v := range vs {
			w.Header().Add(tKey, v)
		}
	}
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, md runtime.ServerMetadata) {
	for k := range md.TrailerMD {
		tKey := textproto.CanonicalMIMEHeaderKey(fmt.Sprintf("%s%s", runtime.MetadataTrailerPrefix, k))
		w.Header().Add("Trailer", tKey)
	}
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, mux *runtime.ServeMux, md runtime.ServerMetadata) {
	for k, vs := range md.HeaderMD {
		if h, ok := OutgoingMatcher(k); ok {
			for _, v := range vs {
				w.Header().Add(h, v)
			}
		}
	}
}
