package main

import (
	"fmt"
	lorem "github.com/drhodes/golorem"
	"math/rand"
)

type Entry struct {
	PrimaryKey int
	Text string
	Value int
}

func main() {
	var entries []Entry

	for i := 0; i < 100; i++ {
		value := Entry{
			PrimaryKey: i,
			Text:       lorem.Word(3, 10),
			Value: rand.Intn(100),
		}
		entries = append(entries, value)
	}

	basicShardingHash(entries, 3)
}

func basicShardingHash(entries []Entry, shardSize int) {
	for _, entry := range entries {
		shard := entry.PrimaryKey % shardSize
		fmt.Printf("[BASIC] ID: %d | Value: %d | Shard: %d\n", entry.PrimaryKey, entry.Value, shard)
	}
}

func rangeShardingHash(entries []Entry, low int, high int) {
	for _, entry := range entries {
		fmt.Printf("[RANGE] ID: %d | Value: %d | Shard: %d\n", entry.PrimaryKey, entry.Value, shard)
	}
}