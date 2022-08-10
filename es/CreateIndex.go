package es

import (
	"context"
	"fmt"
)

func CreateIndex(indexName, mapping string) error {

	ctx := context.Background()
	createIndex, err := esClient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !createIndex.Acknowledged {
		return fmt.Errorf("no ACK")
	}
	return nil
}
