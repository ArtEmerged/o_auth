package main

import (
	"context"
	"log"

	"github.com/ArtEmerged/o_auth-server/internal/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.New(ctx)
	if err != nil {
		log.Printf("failed to init app: %v", err)
		return
	}

	err = a.Run()
	if err != nil {
		log.Printf("failed to run app: %v", err)
	}

}
