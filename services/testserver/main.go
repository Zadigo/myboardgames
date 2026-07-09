package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/redis/go-redis/v9"
)

// // reader registers a periodic job and reports only the registration
// // outcome on ch. Runtime errors from the ticking job itself are logged
// // directly, not sent over ch, since nothing may still be listening on it
// // by the time they occur.
// func reader(ctx context.Context, ch chan<- error) {
// 	tickCtx, cancel := context.WithTimeout(ctx, 10*time.Second)

// 	log.Print("Starting reader...")

// 	s := gocron.NewScheduler(time.UTC)
// 	_, err := s.Every(1).Second().Do(func() {
// 		ch <- fmt.Errorf("Error from the job")
// 		select {
// 		case <-tickCtx.Done():
// 			log.Printf("Error %v", tickCtx.Err())
// 			s.Stop()
// 			cancel()
// 			return
// 		default:
// 			log.Print("Continuing...")
// 		}
// 	})

// 	if err != nil {
// 		cancel()
// 		ch <- err
// 		return
// 	}

// 	s.StartBlocking()
// 	ch <- nil
// }

// func main() {
// 	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
// 	defer done()

// 	// Buffered so reader's send can never block, even if main isn't at
// 	// the select yet (or has already moved on and stopped listening).
// 	ch := make(chan error, 1)

// 	go reader(ctx, ch)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Printf("Error %v", ctx.Err())
// 			return
// 		case err := <-ch:
// 			if err != nil {
// 				log.Printf("Job: %v", err)
// 			}
// 		default:
// 			log.Print("Waiting for job...")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
	defer done()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	ch := make(chan string)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					log.Print("Channel closed")
					return
				} else {
					log.Printf("Received message: %s", msg)
					client.Publish(ctx, "kendall", msg)
				}
			case <-ctx.Done():
				log.Printf("Reader error %v", ctx.Err())
				client.Close()
				done()
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			// client.Publish(ctx, "kendall", fmt.Sprintf("Hello! %s", faker.Name()))
			ch <- fmt.Sprintf("Hello! %s", faker.Name())
			time.Sleep(5 * time.Second)

			select {
			case <-ctx.Done():
				log.Printf("Sender error %v", ctx.Err())
				close(ch)
				client.Close()
				done()
				return
			default:
				log.Print("Generating name...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wg.Add(3)
	go func() {
		sub := client.Subscribe(ctx, "kendall")
		defer sub.Close()

		ctx4, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
		defer cancel()

		for {
			msg, err := sub.ReceiveMessage(ctx4)
			if err != nil {
				log.Printf("Subscriber error: %v", err)
				return
			}
			log.Printf("Received message from channel %s: %s", msg.Channel, msg.Payload)
		}
	}()
	wg.Wait()

	<-ctx.Done()
}
