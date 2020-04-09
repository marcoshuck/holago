package main

import (
	"fmt"
	"math/rand"
	"time"
)

type QueueItem struct {
	Value int
	Zone string
	Subnet string
}

func main() {
	queue := make(chan QueueItem, 1000)
	packages := generateRandomSlice(500000)
	var received []int
	zones := []string{"us-east-1a", "us-east-1c", "us-east-1d", "us-east-1e"}
	go func() {
		var count int
		for {
			for _, v := range packages {
				q := QueueItem{
					Value:  v,
					Zone:   zones[count % len(zones)],
					Subnet: "",
				}
				fmt.Printf("Package [%d] sent. Value: [%d]. AZ: [%s]\n", count, q.Value, q.Zone)
				queue <- q
				generateRandomDelay()
				count++
			}
		}
	}()

	go func() {
		var sum int
		for {
			select {
			case q := <- queue:
				fmt.Printf("Package received. Value [%d]. AZ: [%s]\n", q.Value, q.Zone)
				sum += q.Value
				fmt.Printf("SUM: [%d]\n", sum)
				received = append(received, q.Value)
			}
		}
	}()

	for len(received) != 500000 {

	}
}

func generateRandomSlice(size int) []int {
	var i int
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i = 0; i < size; i++ {
		slice[i] = rand.Intn(999)
	}
	return slice
}

func generateRandomDelay() {
	n := rand.Intn(250)
	time.Sleep(time.Duration(n)*time.Millisecond)
}