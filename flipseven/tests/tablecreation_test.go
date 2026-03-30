package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zadigo/flipseven/internal"
	"github.com/redis/go-redis/v9"
)

func redisConnection() *redis.Client {
	options, _ := redis.ParseURL("redis://:@localhost:6379/0")
	client := redis.NewClient(options)
	return client
}

func TestCreateTableHandler(t *testing.T) {
	body := map[string]any{
		"username": "Test Player",
	}
	jsonBody, err := json.Marshal(body)

	if err != nil {
		t.Fatalf("Failed to marshal body: %v", err)
	}

	req := httptest.NewRequest("POST", "/v1/flip-seven/create", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	internal.CreateTableHandler(w, req, redisConnection())
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	var result map[string]any
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status: %d, body: %s", resp.StatusCode, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode response: %v", err.Error())
	}

	if _, ok := result["tableId"]; !ok {
		t.Error("Expected 'tableId' in response")
	}

	fmt.Printf("Response is %v", result)
}
