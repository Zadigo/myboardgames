package models

type ServiceError string

func (e ServiceError) Error() string {
	return string(e)
}

const (
	TableNotFound ServiceError = "table not found"
)
