package utils

import "os"

// IsProduction checks if ENV is prod or set
func IsProduction() bool {
	return os.Getenv("ENV") == "prod" || os.Getenv("ENV") == "1" || os.Getenv("ENV") != ""
}
