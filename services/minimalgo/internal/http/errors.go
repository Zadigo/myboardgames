package http

type HttpErrors string

func (e HttpErrors) Error() string {
	return string(e)
}

const (
	ConnectionUpgradeFailError HttpErrors = "Failed to upgrade connection"
	InvalidMessageError        HttpErrors = "Invalid message"
)
