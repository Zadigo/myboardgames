package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Handler for creating a new game table. This is used when the user creates
// a new game and needs a unique table ID that needs to be shared with other players.
func CreateTableHandler(response http.ResponseWriter, request *http.Request, redisClient *redis.Client) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if request.Header.Get("Content-Type") != "application/json" {
		http.Error(response, "Unsupported media type", http.StatusUnsupportedMediaType)
		return
	}

	tableId := uuid.NewString()

	var message struct{ Username string }
	err := json.NewDecoder(request.Body).Decode(&message)

	if message.Username == "" {
		http.Error(response, "Username is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(response, fmt.Sprintf("Failed to parse data: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if message.Username == "" {
		http.Error(response, "Username is required", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	redisClient.HSet(ctx, tableId, []string{"initiator", message.Username, "createdAt", time.Now().Format(time.RFC3339), "openForJoin", "true"})

	err = json.NewEncoder(response).Encode(PostDataMessage{
		TableId: tableId,
	})

	if err != nil {
		http.Error(response, fmt.Sprintf("Failed to encode response: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Created new table with ID: %s", tableId)
}
