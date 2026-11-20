package server

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func checkDjangoBackend(ctx context.Context, chError chan<- error) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(40).Second().Do(func() {
		select {
		case <-time.After(time.Second):
			// ping backend here
			fmt.Println("Checked backend")
		case <-ctx.Done():
			scheduler.Stop()
			chError <- ctx.Err()
		}
	})
	scheduler.StartAsync()
}

func checkActiceConnections(ctx context.Context, server *ServerApp, chError chan<- error) {
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(30).Second().Do(func() {
		select {
		case <-time.After(time.Second):
			for _, c := range server.Clients() {
				conn := c.GetConn()
				if conn == nil {
					// Remove the client everywhere
					return
				}

				err := conn.WriteMessage(1, []byte("ping"))
				if err != nil {
					err := conn.Close()
					if err != nil {
						// Remove the client everywhere
						fmt.Println("Error closing connection")
					}
					return
				}
			}
		case <-ctx.Done():
			scheduler.Stop()
			chError <- ctx.Err()
		}
	})
	scheduler.StartAsync()
}
