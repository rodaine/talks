package grpc_lyft

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const RequestIDHeader = "x-request-id" // OMIT

func RequestID(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if meta, ok := metadata.FromContext(ctx); ok { // HL
		if hdrs, ok := meta[RequestIDHeader]; ok && len(hdrs) > 0 { // HL
			ctx = context.WithValue(ctx, RequestIDHeader, hdrs[0]) // HL
		} // HL
	} // HL

	return handler(ctx, req) // HL
}

// END RequestID OMIT

func Chain(wrappers ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		for i := len(wrappers) - 1; i >= 0; i-- { // HL
			handler = wrapHandler(wrappers[i], info, handler) // HL
		} // HL

		return handler(ctx, req) // HL
	}
}

func wrapHandler(
	wrapper grpc.UnaryServerInterceptor,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) grpc.UnaryHandler {

	return func(ctx context.Context, req interface{}) (interface{}, error) { // HL
		return wrapper(ctx, req, info, handler) // HL
	} // HL
}

// END Chain OMIT

func Metrics(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	return nil, nil
}

func InitServer() { // OMIT
	grpc.NewServer(
		grpc.UnaryInterceptor(
			Chain( // HL
				RequestID, // HL
				Metrics,   // HL
				// etc..., // HL
			), // HL
		),
		// other server options...
	)
} // OMIT
