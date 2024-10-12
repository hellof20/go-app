package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"cloud.google.com/go/bigtable"
)

func generateRandomRowKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// func writeSimple(projectID, instanceID, tableName, columnFamilyName string) error {
// 	ctx := context.Background()
// 	client, err := bigtable.NewClient(ctx, projectID, instanceID)
// 	if err != nil {
// 		return fmt.Errorf("bigtable.NewClient: %w", err)
// 	}
// 	defer client.Close()

// 	tbl := client.Open(tableName)
// 	timestamp := bigtable.Now()

// 	mut := bigtable.NewMutation()
// 	mut.Set(columnFamilyName, "connected_cell", timestamp, []byte("1"))
// 	mut.Set(columnFamilyName, "connected_wifi", timestamp, []byte("1"))
// 	mut.Set(columnFamilyName, "os_build", timestamp, []byte("PQ2A.190405.003"))

// 	rowKey, err := generateRandomRowKey(16) // 生成16字节（32个十六进制字符）的随机rowKey
// 	if err != nil {
// 		return fmt.Errorf("generateRandomRowKey: %w", err)
// 	}

// 	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
// 		return fmt.Errorf("apply: %w", err)
// 	}

// 	fmt.Printf("Successfully wrote row with key '%s' to table %s\n", rowKey, tableName) // 打印rowKey以便确认
// 	return nil
// }

func main() {
	projectID := "speedy-victory-336109"
	instanceID := "mybtinstance"
	tableName := "mytable"
	columnFamilyName := "cf1"

	ctx := context.Background()
	client, err := bigtable.NewClientWithConfig(ctx, projectID, instanceID, bigtable.ClientConfig{MetricsProvider: bigtable.NoopMetricsProvider{}})
	if err != nil {
		fmt.Errorf("bigtable.NewClient: %w", err)
	}
	defer client.Close()

	tbl := client.Open(tableName)

	for {
		timestamp := bigtable.Now()

		mut := bigtable.NewMutation()
		mut.Set(columnFamilyName, "connected_cell", timestamp, []byte("1"))
		mut.Set(columnFamilyName, "connected_wifi", timestamp, []byte("1"))
		mut.Set(columnFamilyName, "os_build", timestamp, []byte("PQ2A.190405.003"))

		rowKey, err := generateRandomRowKey(16) // 生成16字节（32个十六进制字符）的随机rowKey
		if err != nil {
			fmt.Errorf("generateRandomRowKey: %w", err)
		}

		if err := tbl.Apply(ctx, rowKey, mut); err != nil {
			fmt.Errorf("apply: %w", err)
		}

		fmt.Printf("Successfully wrote row with key '%s' to table %s\n", rowKey, tableName) // 打印rowKey以便确认
	}

}
