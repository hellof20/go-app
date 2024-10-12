package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigtable"
)

func main() {
	projectID := "speedy-victory-336109"
	instanceID := "mybtinstance"
	tableName := "mytable"
	rowKey := "cab4c3687092169dd5208f02ae3a994e"
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, projectID, instanceID)

	if err != nil {
		fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()

	tbl := client.Open(tableName)
	for {
		row, err := tbl.ReadRow(ctx, rowKey)
		if err != nil {
			fmt.Errorf("could not read row with key %s: %w", rowKey, err)
		}
		print(row)
	}
}
