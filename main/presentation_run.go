package main

import (
	"fmt"
	"kotlin-presentation/main/users"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	results := make([]time.Duration, 0)
	for i := 0; i < 50; i++ {
		start := time.Now()
		users.Test(10000, 50)
		elapsed := time.Since(start)
		results = append(results, elapsed)
	}

	fmt.Println("**********************************\nResults:")
	for _, result := range results {
		fmt.Printf("Total took: %s\n", result)
	}
}
