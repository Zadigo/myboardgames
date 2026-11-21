package internal

import (
	"context"
	"log/slog"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoveryUnary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("Panic recovered", "method", info.FullMethod, "panic", r, "stack", string(debug.Stack()))
			err = status.Errorf(codes.Internal, "Internal error")
		}
	}()

	return handler(ctx, req)
}

func recoveryStream(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("Panic recovered", "panic", r, "stack", string(debug.Stack()))
			err = status.Errorf(codes.Internal, "Internal error")
		}
	}()

	return handler(srv, ss)
}
