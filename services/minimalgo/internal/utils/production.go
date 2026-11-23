package utils

import "os"

// IsProduction checks if GRPC_ENV is prod or set
func IsProduction() bool {
	return os.Getenv("GRPC_ENV") == "prod" || os.Getenv("GRPC_ENV") == "1" || os.Getenv("GRPC_ENV") != ""
}
