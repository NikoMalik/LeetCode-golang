package main

import (
	"fmt"
	"slices"
	"unsafe"
)

func findPairs(nums []int, k int) int {
	if k < 0 || len(nums) == 0 {
		panic("invalid input")
	}
	slices.Sort(nums)

	n := len(nums)
	ans := 0
	i := 0

	base := unsafe.Pointer(unsafe.SliceData(nums))
	size := unsafe.Sizeof(nums[0])

	for i < n-1 {
		if i > 0 && *(*int)(unsafe.Add(base, uintptr(i)*size)) == *(*int)(unsafe.Add(base, uintptr(i-1)*size)) {
			i++
			continue
		}

		a := *(*int)(unsafe.Add(base, uintptr(i)*size))
		j := i + 1

		for j < n {
			for ; j+3 < n; j += 4 {
				b0 := *(*int)(unsafe.Add(base, uintptr(j+0)*size))
				b1 := *(*int)(unsafe.Add(base, uintptr(j+1)*size))
				b2 := *(*int)(unsafe.Add(base, uintptr(j+2)*size))
				b3 := *(*int)(unsafe.Add(base, uintptr(j+3)*size))

				if b0-a == k || b1-a == k || b2-a == k || b3-a == k {
					ans++
					goto nextI
				}
				if b3-a > k {
					goto nextI
				}
			}

			for ; j < n; j++ {
				b := *(*int)(unsafe.Add(base, uintptr(j)*size))
				diff := b - a
				if diff == k {
					ans++
					goto nextI
				} else if diff > k {
					goto nextI
				}
			}
		}
	nextI:
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
