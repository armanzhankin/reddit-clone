package main

import (
	"context"
	"github.com/armanzhankin/reddit-clone/services/post-service/internal/config"
	"os"

	"github.com/armanzhankin/reddit-clone/services/post-service/pkg/postgres"
)

func main() {
	ctx := context.Background()

	config, err := config.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	pool, err := postgres.NewPgxPool(ctx, dbUrl)

}
