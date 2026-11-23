package http

import (
	"log"
	"myboardgames/minimalgo/internal/constants"
	"net/http"
)

func Cors(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		origin := r.Header.Get("Origin")

		if _, ok := constants.AllowedOrigins[origin]; !ok {
			log.Println("❌ Origin not allowed:", origin)
			http.Error(w, "Origin not allowed", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
