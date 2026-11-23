package grpcclient

import "context"

type playerKey struct{}

// Returns a copy of ctx that carries the player ID
func WithPlayer(ctx context.Context, playerID string) context.Context {
	return context.WithValue(ctx, playerKey{}, playerID)
}

// Reads the player ID back out
func PlayerFrom(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(playerKey{}).(string)
	return id, ok && id != ""
}
