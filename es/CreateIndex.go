package main

import (
	"context"
	"fmt"
)

func CreateIndex(mapping string) error {
	client := NewEsClient()
	indexName := "users"
	ctx := context.Background()
	createIndex, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged {
		return fmt.Errorf("no ACK")
	}
	return nil
}
