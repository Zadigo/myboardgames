package server

type ServerErrors string

func (e ServerErrors) Error() string {
	return string(e)
}

const (
	FailedToLoadRedisError     ServerErrors = "Failed to load redis"
	RpcConnectionNotFoundError ServerErrors = "RPC connection not found"
)
