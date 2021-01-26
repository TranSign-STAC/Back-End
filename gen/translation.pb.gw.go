// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: translation.proto

/*
Package transign is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package transign

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_TextToSignLang_Translate_0(ctx context.Context, marshaler runtime.Marshaler, client TextToSignLangClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TextToSignLangRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Translate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TextToSignLang_Translate_0(ctx context.Context, marshaler runtime.Marshaler, server TextToSignLangServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TextToSignLangRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Translate(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_TranslationHistory_GetHistory_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_TranslationHistory_GetHistory_0(ctx context.Context, marshaler runtime.Marshaler, client TranslationHistoryClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UUIDMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_TranslationHistory_GetHistory_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_TranslationHistory_GetHistory_0(ctx context.Context, marshaler runtime.Marshaler, server TranslationHistoryServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UUIDMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_TranslationHistory_GetHistory_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetHistory(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_FavoriteTranslation_GetFavorite_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_FavoriteTranslation_GetFavorite_0(ctx context.Context, marshaler runtime.Marshaler, client FavoriteTranslationClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UUIDMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FavoriteTranslation_GetFavorite_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetFavorite(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FavoriteTranslation_GetFavorite_0(ctx context.Context, marshaler runtime.Marshaler, server FavoriteTranslationServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UUIDMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_FavoriteTranslation_GetFavorite_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetFavorite(ctx, &protoReq)
	return msg, metadata, err

}

func request_FavoriteTranslation_ToggleFavorite_0(ctx context.Context, marshaler runtime.Marshaler, client FavoriteTranslationClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ToggleFavoriteTranslationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ToggleFavorite(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FavoriteTranslation_ToggleFavorite_0(ctx context.Context, marshaler runtime.Marshaler, server FavoriteTranslationServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ToggleFavoriteTranslationRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ToggleFavorite(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterTextToSignLangHandlerServer registers the http handlers for service TextToSignLang to "mux".
// UnaryRPC     :call TextToSignLangServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTextToSignLangHandlerFromEndpoint instead.
func RegisterTextToSignLangHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TextToSignLangServer) error {

	mux.Handle("POST", pattern_TextToSignLang_Translate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/transign.TextToSignLang/Translate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TextToSignLang_Translate_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TextToSignLang_Translate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTranslationHistoryHandlerServer registers the http handlers for service TranslationHistory to "mux".
// UnaryRPC     :call TranslationHistoryServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterTranslationHistoryHandlerFromEndpoint instead.
func RegisterTranslationHistoryHandlerServer(ctx context.Context, mux *runtime.ServeMux, server TranslationHistoryServer) error {

	mux.Handle("GET", pattern_TranslationHistory_GetHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/transign.TranslationHistory/GetHistory")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_TranslationHistory_GetHistory_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TranslationHistory_GetHistory_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFavoriteTranslationHandlerServer registers the http handlers for service FavoriteTranslation to "mux".
// UnaryRPC     :call FavoriteTranslationServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFavoriteTranslationHandlerFromEndpoint instead.
func RegisterFavoriteTranslationHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FavoriteTranslationServer) error {

	mux.Handle("GET", pattern_FavoriteTranslation_GetFavorite_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/transign.FavoriteTranslation/GetFavorite")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FavoriteTranslation_GetFavorite_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FavoriteTranslation_GetFavorite_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_FavoriteTranslation_ToggleFavorite_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/transign.FavoriteTranslation/ToggleFavorite")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FavoriteTranslation_ToggleFavorite_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FavoriteTranslation_ToggleFavorite_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterTextToSignLangHandlerFromEndpoint is same as RegisterTextToSignLangHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTextToSignLangHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTextToSignLangHandler(ctx, mux, conn)
}

// RegisterTextToSignLangHandler registers the http handlers for service TextToSignLang to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTextToSignLangHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTextToSignLangHandlerClient(ctx, mux, NewTextToSignLangClient(conn))
}

// RegisterTextToSignLangHandlerClient registers the http handlers for service TextToSignLang
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TextToSignLangClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TextToSignLangClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TextToSignLangClient" to call the correct interceptors.
func RegisterTextToSignLangHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TextToSignLangClient) error {

	mux.Handle("POST", pattern_TextToSignLang_Translate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/transign.TextToSignLang/Translate")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TextToSignLang_Translate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TextToSignLang_Translate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TextToSignLang_Translate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"translate", "text"}, ""))
)

var (
	forward_TextToSignLang_Translate_0 = runtime.ForwardResponseMessage
)

// RegisterTranslationHistoryHandlerFromEndpoint is same as RegisterTranslationHistoryHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTranslationHistoryHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterTranslationHistoryHandler(ctx, mux, conn)
}

// RegisterTranslationHistoryHandler registers the http handlers for service TranslationHistory to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTranslationHistoryHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTranslationHistoryHandlerClient(ctx, mux, NewTranslationHistoryClient(conn))
}

// RegisterTranslationHistoryHandlerClient registers the http handlers for service TranslationHistory
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "TranslationHistoryClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TranslationHistoryClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TranslationHistoryClient" to call the correct interceptors.
func RegisterTranslationHistoryHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TranslationHistoryClient) error {

	mux.Handle("GET", pattern_TranslationHistory_GetHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/transign.TranslationHistory/GetHistory")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TranslationHistory_GetHistory_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TranslationHistory_GetHistory_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TranslationHistory_GetHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"translate", "history"}, ""))
)

var (
	forward_TranslationHistory_GetHistory_0 = runtime.ForwardResponseMessage
)

// RegisterFavoriteTranslationHandlerFromEndpoint is same as RegisterFavoriteTranslationHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFavoriteTranslationHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterFavoriteTranslationHandler(ctx, mux, conn)
}

// RegisterFavoriteTranslationHandler registers the http handlers for service FavoriteTranslation to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFavoriteTranslationHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFavoriteTranslationHandlerClient(ctx, mux, NewFavoriteTranslationClient(conn))
}

// RegisterFavoriteTranslationHandlerClient registers the http handlers for service FavoriteTranslation
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FavoriteTranslationClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FavoriteTranslationClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FavoriteTranslationClient" to call the correct interceptors.
func RegisterFavoriteTranslationHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FavoriteTranslationClient) error {

	mux.Handle("GET", pattern_FavoriteTranslation_GetFavorite_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/transign.FavoriteTranslation/GetFavorite")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FavoriteTranslation_GetFavorite_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FavoriteTranslation_GetFavorite_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_FavoriteTranslation_ToggleFavorite_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/transign.FavoriteTranslation/ToggleFavorite")
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FavoriteTranslation_ToggleFavorite_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FavoriteTranslation_ToggleFavorite_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FavoriteTranslation_GetFavorite_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"translate", "favorite"}, ""))

	pattern_FavoriteTranslation_ToggleFavorite_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"translate", "favorite"}, ""))
)

var (
	forward_FavoriteTranslation_GetFavorite_0 = runtime.ForwardResponseMessage

	forward_FavoriteTranslation_ToggleFavorite_0 = runtime.ForwardResponseMessage
)
