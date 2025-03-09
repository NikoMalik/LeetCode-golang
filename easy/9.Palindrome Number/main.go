package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

var (
	is_64 = bits.UintSize == 64
	is_32 = bits.UintSize == 32
)

//go:nosplit
//go:nocheckptr
func Noescape(up unsafe.Pointer) unsafe.Pointer {
	x := uintptr(up)
	return unsafe.Pointer(x ^ 0)
}

func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	var copy_number = n
	reversed := 0
	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return copy_number == reversed

}

func intToDigitArray(x int, out []int) []int {
	if x < 0 {
		return nil
	}
	var digits [20]int
	count := 0

	if x == 0 {
		digits[count] = 0
		count++
	} else {
		for x > 0 {
			digits[count] = x % 10
			count++
			x /= 10
		}
		for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}
	}
	copy(out, digits[:count])
	_ = Noescape(unsafe.Pointer(&out))
	return out[:count]
}

func isBytePalindrome(x int) bool {
	// runtime.GC()
	if x < 0 {
		return false
	}
	var buf [20]int
	digits := intToDigitArray(x, buf[:])
	n := len(digits)
	for i := 0; i < n/2; i++ {
		if digits[i] != digits[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	const example_number = 121
	const example_number_2 = -121

	is_ := isPalindrome(example_number)
	fmt.Println(is_)
	is_2 := isPalindrome(example_number_2)
	fmt.Println(is_2)

	var x int = 121
	bytesArr := *(*[bits.UintSize / 8]byte)(unsafe.Pointer(&x))
	fmt.Printf("number: %d, array: % x\n", x, bytesArr)

	ten := isBytePalindrome(10)
	if ten {
		fmt.Println("YOOOOO")
	} else {
		fmt.Println("noooooooooo")
	}

	if isBytePalindrome(x) {
		fmt.Println("yes ye syes 121 isPalindrome")
	} else {
		fmt.Println("no 121 fake")
	}
}
