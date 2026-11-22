package utils

import "strings"

const (
	globalStorageKey = "minimalgo"
)

// CreateStorageKey makes a key with the storage prefix and values
func CreateStorageKey(values ...string) string {
	return globalStorageKey + ":" + strings.Join(values, ":")
}
