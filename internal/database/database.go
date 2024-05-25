package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"os"
)

type ClientInterface interface {
	Collection(name string) CollectionInterface
	Close() error
}

// CollectionInterface represents the collection interface
type CollectionInterface interface {
	Add(ctx context.Context, data interface{}) (string, string, error)
}

func Client() *firestore.Client {
	ctx := context.Background()

	projectId := os.Getenv("FIRESTORE_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client
}
