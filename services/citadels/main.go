package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"uuid"

	"github.com/Zadigo/citadels/internal/game"
	"github.com/Zadigo/citadels/internal/observer"
	"github.com/joho/godotenv"
)

func initApplication(ctx context.Context) error {
	// Load the absolute path of the root dir
	if rootDir, err := filepath.Abs("."); err != nil {
		return errors.Join(fmt.Errorf("Error when initializing application"), err)
	} else {
		ctx = context.WithValue(ctx, "rootDir", rootDir)
	}

	// Check the validity of the port
	port := os.Getenv("CITADELS_PORT")
	if port == "" {
		os.Setenv("CITADELS_PORT", "8080")
	} else {
		port, err := strconv.Atoi(port)
		if err != nil {
			return err
		}

		if port < 1 || port > 65535 {
			return fmt.Errorf("Invalid port")
		}
	}

	_url := os.Getenv("CITADELS_URL")
	if _url == "" {
		os.Setenv("CITADELS_URL", "http://localhost")
	} else {
		_url, err := url.Parse(_url)
		if err != nil {
			return err
		}

		os.Setenv("CITADELS_HOST", _url.Host)
	}

	os.Setenv("CITADELS_CONFIGURED", "true")
	fmt.Printf("⚡️ Game initizalized on port: %s", os.Getenv("CITADELS_PORT"))
	return nil
}

func main() {
	fmt.Println("⚡️ Starting Citadels...")

	context := context.Background()
	context, cancel := signal.NotifyContext(context, os.Interrupt, os.Kill)
	defer cancel()

	godotenv.Load(".env", ".env-test")

	instance := observer.NewBaseObserver()
	gameObserver := &observer.GameSubject{
		ID:   uuid.New().String(),
		Name: "Citadels",
	}
	instance.Register(gameObserver)

	err := initApplication(context)
	if err != nil {
		panic(err)
	}

	instance.NotifyAll(observer.ObserverMessage{StrMessage: "Citadels initialized"})

	app := game.NewHttpApplication()
	app.Start(context)
}
