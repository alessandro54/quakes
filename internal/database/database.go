package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func Client() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "alessandro-423819")

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()
}
