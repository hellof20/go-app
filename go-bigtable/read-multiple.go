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
	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, projectID, instanceID)

	if err != nil {
		fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()

	tbl := client.Open(tableName)
	for {
		err = tbl.ReadRows(ctx, bigtable.RowList{
			"12bd032492b948eaec9b092e908d300c",
			"993f9af1b0ddec77c36cd8f3d4f93957",
			"bd3b4d44b1c2eb12f891c49fe66d9ad8",
			"2f7565e2194e1eff82055b0f64c30a56",
			"6b08422953cb1a24cea93d14cabedf56",
			"87408885a92b074d6c46037fe7816ba5",
			"6a6117f7a2b86467d543b37308da669f",
			"34f95c8e9c4b944794aff3c34460d491",
			"6cbb0be1bce7b4afbcdde65de0e75975",
			"ec7987e599f51cc63e14fb8c19186e89",
			"cab4c3687092169dd5208f02ae3a994e",
			"805a86d59de241bcf5089035a10b4a47",
		},
			func(row bigtable.Row) bool {
				print(row)
				return true
			},
		)
		if err != nil {
			fmt.Printf("could not read row with key")
		}
	}
}
