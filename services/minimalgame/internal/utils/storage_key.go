package utils

import "strings"

const (
	storageKey = "minimalgame"
)

// CreateStorageKey makes a key with the storage prefix and values
func CreateStorageKey(values ...string) string {
	return storageKey + ":" + strings.Join(values, ":")
}
