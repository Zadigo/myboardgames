package handlers

import (
	"context"
	"net/http"

	"github.com/Zadigo/flipseven2/internal/backend/broadcasting"
	"github.com/Zadigo/flipseven2/internal/models"
	"github.com/redis/go-redis/v9"
)

type ContextHandlerFunc func(w http.ResponseWriter, r *http.Request, ctx context.Context, redisConn *redis.Client, broadcastRegistry *broadcasting.BroadcasterRegistry, baseRegistry *models.BaseRegistry)
