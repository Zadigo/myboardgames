package tests

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Zadigo/oriflamme/internal/backend/redisclient"
)

func TestOther(t *testing.T) {
	c := redisclient.NewRedisClient("redis://@localhost:6379/0")

	var wg sync.WaitGroup
	wg.Add(1)

	wg.Go(func() {
		p := c.Subscribe(context.Background(), "simplechannel")
		chann := p.Channel()

		for msg := range chann {
			t.Logf("Received message: %s", msg.Payload)
		}
		wg.Done()
	})

	wg.Go(func() {
		for range 10 {
			err := c.Publish(context.Background(), "simplechannel", "Hello, World!")
			if err != nil {
				t.Errorf("Failed to publish message: %v", err)
			}
			time.Sleep(5 * time.Millisecond)
		}
		wg.Done()
	})

	wg.Wait()
}
