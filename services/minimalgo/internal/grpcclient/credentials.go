package grpcclient

import (
	"context"
	"crypto/ed25519"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// tokenCreds implements grpc's credentials.PerRPCCredentials interface.
// gRPC calls GetRequestMetadata before EVERY RPC and adds the returned map
// to the request headers. You never call it yourself.
type tokenCreds struct {
	priv ed25519.PrivateKey // private key, only minimalgo has it
	kid  string             // key ID, so minimalgame knows which public key to use
}

func (t tokenCreds) GetRequestMetadata(ctx context.Context, _ ...string) (map[string]string, error) {
	// ctx is the SAME ctx you passed to backend.Table.CreateTable(ctx, ...),
	// so the player ID set with WithPlayer is available here.
	playerID, ok := PlayerFrom(ctx)
	if !ok {
		// Fail closed: refuse to send a call that doesn't say who it is for.
		return nil, status.Error(codes.Unauthenticated, "no player in context")
	}

	now := time.Now()
	tok := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"iss":    "minimalgo",   // must match the issuer minimalgame expects
		"aud":    "minimalgame", // must match the audience minimalgame expects
		"sub":    "minimalgo",
		"iat":    now.Unix(),
		"exp":    now.Add(time.Minute).Unix(), // short-lived
		"jti":    uuid.NewString(),
		"scopes": []string{"table:create", "table:read", "table:join"},
		"obo":    playerID, // "on behalf of": the player for this call
	})
	tok.Header["kid"] = t.kid

	signed, err := tok.SignedString(t.priv)
	if err != nil {
		return nil, err
	}
	return map[string]string{"authorization": "Bearer " + signed}, nil
}

// true = gRPC refuses to send the token over a non-TLS connection.
func (tokenCreds) RequireTransportSecurity() bool { return true }
