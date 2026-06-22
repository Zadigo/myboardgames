package utils

var AllowedOrigins = map[string]bool{
	"http://localhost":      true,
	"http://localhost:3000": true,
	"http://127.0.0.1:8000": true,
	"PostmanRuntime/7.*":    true,
}
