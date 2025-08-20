package main

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkFindPairs(b *testing.B) {

	n := 1000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.IntN(10000)
	}
	k := 5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPairs(nums, k)
	}
}
