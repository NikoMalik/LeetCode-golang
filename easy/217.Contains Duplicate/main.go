package main

import (
	"fmt"
	"sort"
)

type Flags uint32

// Set sets the specified flag to true.
//
// For example, if flag = 0, it corresponds to 1 << 0, which equals 1 (00000001 in binary).
// If flag = 1, this corresponds to 1 << 1, which equals 2 (00000010 in binary).
// If flag = 2, this corresponds to 1 << 2, which equals 4 (00000100 in binary).
//
// The bitwise OR operator (|=) sets the specified flag to f.  This means that if this flag was not set, it will be set.
// For example, if f = 00000000 and we set flag 0 (1 << 0), f will become 00000001.
// If f = 00000001 and we set flag 0 again, the value of f will not change (it will remain 00000001).
func (f *Flags) Set(flag byte) {

	*f |= Flags(1 << flag)
}

// Clear sets the specified flag to false.  This does not change the value of any other flags.
// The bitwise AND operator (&=) is used to clear the specified flag.  The expression
//
//	f &= ^Flags(1 << flag)
//
// will clear the specified flag, but not change the value of any other flag.
// For example, if f = 00000001 and flag = 0, then f will become 00000000.
// If f = 00000101 and flag = 0, then f will become 00000100.

func (f *Flags) Clear(flag byte) {
	*f &^= Flags(1 << flag)

}

// IsSet returns true if the specified flag is set, false otherwise.
// The expression f&(1<<flag) is a binary operation that compares the value of the
// f byte with the mask (1<<flag).  If the flag is set, the expression will be

// non-zero; if the flag is not set, the expression will be zero.

func (f Flags) IsSet(flag byte) bool {
	return f&(1<<flag) != 0
}

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	// sort.Ints(nums) // if not sort i get rundom numbers and cannot iterrate
	var flag Flags

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			flag.Set(0x1f)
			return flag.IsSet(0x1f)

		}

	}

	return flag.IsSet(0x1f)

}

func containsDuplicate_2(nums []int) bool {
	if len(nums) < 2 {
		return false

	}
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true

		}

	}

	return false

}

func main() {
	nums := []int{4, 3, 2, 1, 1}
	fmt.Println(containsDuplicate(nums))
	nums_1 := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums_1))

	nums_3 := []int{94886, 28821, 20675, 98974, 22621, 91823, 44730, 6160, 99710, 46764, 46571, 55716, 11540, 28266, 73209, 13807, 8993, 18841, 89902, 59175, 90398, 77670, 64665, 40145, 45521, 57445, 49570, 11896, 18068, 67432, 83592, 26821, 58985, 26944, 84818, 54572, 20638, 14656, 206, 70860, 93071, 44055, 96137, 3536, 72607, 78294, 97322, 77146, 57043, 54165, 90341, 61515, 8318, 1711, 80599, 37455, 83277, 24794, 77134, 25009, 148, 21904, 71534, 95227, 96105, 27114, 83113, 5950, 11348, 15844, 75966, 47447, 27051, 8887, 35354, 40231, 70795, 11906, 33565, 86815, 71649, 35906, 75200, 70001, 97367, 42827, 83577, 16687, 27315, 16493, 93667, 91973, 48885, 15583, 17731, 54936, 74874, 82474, 1399, 73080, 13759, 25206, 81472, 58730, 5147, 74013, 21594, 53260, 33459, 91020, 52332, 45901, 84370, 44027, 41557, 4514, 87642, 55924, 75849, 89189, 26048, 20291, 90824, 92839, 97513, 16950, 43040, 26594, 20167, 23935, 14649, 8783, 58112, 80561, 40255, 78519, 92347, 63255, 75073, 65568, 16186, 61602, 19834, 80650, 9385, 78664, 32019, 60731, 13899, 20208, 70056, 83777, 96104, 63008, 28428, 71002, 90604, 5387, 79698, 72818, 31054, 82023, 44135, 71636, 57306, 61287, 80836, 78352, 26708}
	fmt.Println(containsDuplicate(nums_3))

	nums_4 := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums_4))
}
