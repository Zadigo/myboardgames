package tests

import (
	"fmt"
	"slices"
	"testing"

	"github.com/google/uuid"
)

func TestOther(t *testing.T) {
	type c struct {
		uuid string
	}

	v := []*c{}

	v = append(v, &c{uuid: uuid.NewString()})
	v = append(v, &c{uuid: uuid.NewString()})
	v = append(v, &c{uuid: uuid.NewString()})

	fmt.Print(v)

	w := slices.Delete(v, 1, 2)

	fmt.Print(w[0])

	// c := redisclient.NewRedisClient("redis://@localhost:6379/0")

	// var wg sync.WaitGroup
	// wg.Add(1)

	// wg.Go(func() {
	// 	p := c.Subscribe(context.Background(), "simplechannel")
	// 	chann := p.Channel()

	// 	for msg := range chann {
	// 		t.Logf("Received message: %s", msg.Payload)
	// 	}
	// 	wg.Done()
	// })

	// wg.Go(func() {
	// 	for range 10 {
	// 		err := c.Publish(context.Background(), "simplechannel", "Hello, World!")
	// 		if err != nil {
	// 			t.Errorf("Failed to publish message: %v", err)
	// 		}
	// 		time.Sleep(5 * time.Millisecond)
	// 	}
	// 	wg.Done()
	// })

	// wg.Wait()
}
