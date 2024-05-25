package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"os"
)

func Client() *firestore.Client {
	err := os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9000")

	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "alessandro-423819")

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
