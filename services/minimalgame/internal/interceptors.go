package internal

import (
	"context"
	"log/slog"
	"myboardgames/minimalgame/internal/backend"
	"runtime/debug"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	pb "myboardgames/minimalgame/minimalgameproto"
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

// Default deny: a method that is not listed here is rejected.
var methodScopes = map[string]string{
	pb.TableService_CreateTable_FullMethodName: "table:create",
	pb.TableService_GetTable_FullMethodName:    "table:read",
	pb.TableService_JoinTable_FullMethodName:   "table:join",
	pb.TableService_LeaveTable_FullMethodName:  "table:join",
	pb.TableService_EndGame_FullMethodName:     "table:admin",
	pb.TableService_BanPlayer_FullMethodName:   "table:admin",
}

var publicMethods = map[string]struct{}{
	// "/grpc.health.v1.Health/Check": {},
}

type callerKey struct{}

func authUnary(v *backend.Verifier) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, next grpc.UnaryHandler) (any, error) {
		if _, ok := publicMethods[info.FullMethod]; ok {
			return next(ctx, req)
		}

		// Checks
		metadata, ok := metadata.FromIncomingContext(ctx)
		values := metadata.Get("authorization")
		if !ok || len(values) != 1 {
			return nil, status.Error(codes.Unauthenticated, "unauthenticated")
		}

		token, found := strings.CutPrefix(values[0], "Bearer ")
		if !found || token == "" {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}

		// Parse the JWT
		claims, err := v.Parse(token)
		if err != nil {
			slog.Warn("Authentication failed", "method", info.FullMethod, "error", err)
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}

		required, ok := methodScopes[info.FullMethod]
		if !ok {
			return nil, status.Error(codes.PermissionDenied, "forbidden")
		}

		caller := &backend.Caller{
			Service:    claims.Subject,
			OnBehalfOf: claims.OnBehalfOf,
			Scopes:     make(map[string]struct{}, len(claims.Scopes)),
		}

		if !caller.Has(required) {
			slog.Warn("scope denied", "method", info.FullMethod, "caller", caller.Service)
			return nil, status.Error(codes.PermissionDenied, "forbidden")
		}

		return next(context.WithValue(ctx, callerKey{}, caller), req)
	}
}

// Client mistakes are warnings, server faults are errors
func isClientError(c codes.Code) bool {
	switch c {
	case codes.InvalidArgument, codes.NotFound, codes.AlreadyExists,
		codes.PermissionDenied, codes.Unauthenticated, codes.FailedPrecondition,
		codes.ResourceExhausted, codes.Canceled:
		return true
	}
	return false
}

// loggingUnary wraps every RPC and records what happens: method, duration, status code, and caller.
func loggingUnary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	response, err := handler(ctx, req)

	code := status.Code(err) // codes.OK when err == nil

	attrs := []any{
		"method", info.FullMethod,
		"code", code.String(),
		"duration", time.Since(start),
	}

	if p, ok := peer.FromContext(ctx); ok {
		attrs = append(attrs, "peer", p.Addr.String())
	}

	switch {
	case code == codes.OK:
		slog.Info("rpc", attrs...)
	case isClientError(code):
		slog.Warn("rpc", append(attrs, "err", err)...)
	default:
		slog.Error("rpc", append(attrs, "err", err)...)
	}

	return response, err
}
