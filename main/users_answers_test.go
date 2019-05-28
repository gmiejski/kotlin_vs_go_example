package main

import (
	"kotlin-presentation/main/users"
	"testing"
)

func BenchmarkUsersStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		users.Test(10000, 50)
	}
}
