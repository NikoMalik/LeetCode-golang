package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		panic("invalid input")
	}
	sort.Ints(nums)

	ans := 0
	n := len(nums)
	i := 0

	for i < n-1 {

		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		for j := i + 1; j < n; j++ {
			diff := nums[j] - nums[i]
			if diff == k {
				ans++
				break
			} else if diff > k {
				break
			}
		}

		i++
	}

	return ans
}

func main() {

	f := findPairs([]int{3, 1, 4, 1, 5}, 2)
	fmt.Println("SUCCESS 31415")
	if f != 2 {

		fmt.Println("expected 2, got", f)
	}

	f1 := (findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println("SUCCESS 12345")
	if f1 != 4 {
		fmt.Println("expected 4, got", f1)
	}
}
