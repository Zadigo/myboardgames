package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func reader(ch <-chan error) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			log.Printf("reading %v", v)
		default:
			log.Print("Waiting...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan error)
	wg.Add(1)
	go reader(ch)
	for range [10]int{} {
		ch <- fmt.Errorf("Some error")
	}
	close(ch)
	wg.Done()
}
