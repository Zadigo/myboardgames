package utils

func HasDuplicates[T comparable](slice []T) bool {
	seen := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		if _, exists := seen[item]; exists {
			return true
		}
		seen[item] = struct{}{}
	}
	return false
}
