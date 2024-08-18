package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

// first solution easy to figure out
// we use two loops for enumerating all possible pairs of elements in nums.
//If a pair is found that satisfies the condition, the function returns their indices. Otherwise it returns an empty slice
func twoSum(nums []int, target int) []int { // O(n²)

	for i := 0; i < len(nums); i++ {
		//This loop iterates over all elements of slice nums with index i.
		//i starts at 0 and increases until it is equal to the length of slice nums

		for j := i + 1; j < len(nums); j++ {
			// This loop iterates over all elements of the slice nums, starting at index i + 1 and ending at the end of the slice
			// It does this so that it doesn't compare one element to itself or compare pairs in reverse order

			if nums[i]+nums[j] == target {
				//It is checked whether the sum of elements at indices i and j is equal to the target value of target
				return []int{i, j}
				//If yes, these are the elements to be returned.
			}
		}
		// very simple and easy to start solution
		// try to look at everything and then write a solution from scratch without peeking

	}

	/*
		This algorithm has a time complexity of O(n²) due to two nested loops, where n is the slice length of nums.
		It is inefficient for large arrays, but works fine for small arrays. To improve performance, a hash table can be used to find pairs with sum equal to target in O(n) time



	*/

	return []int{}

}

// second solution using hash table more efficient but more hard to understand O(n)
func twoSum1(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	//make(map[int]int, len(nums)) creates a hash table with initial capacity equal to the slice length of nums.
	//This means that the hash table initially allocates enough memory to store the len(nums) elements, without having to be expanded as new elements are added.

	for i := 0; i < len(nums); i++ { // for loop faster than range loop
		// For loop: Loops through all elements of the nums slice using index i
		complement := target - nums[i] //calculate what number we need to find in the hash table so that the sum with the current number equals target.
		//Calculate complement: Calculates what value (complement) is needed to reach the target if the current value of nums[i] is already added to the hash table
		// example  6 - 3 = 3 | nums[0] = 3
		//check complement is in numMap: numMap[3] = 0
		// second iteration 6 - 2 = 4 | nums[1] = 2
		// check complement is in numMap: numMap[2] = 1
		// third iteration 6 - 4 = 2 | nums[2] = 4
		// check complement is in numMap: numMap[4] = 2
		//Check if there is 2 in numMap. There is (with index 1).
		//Return indices [1, 2].

		if j, ok := numMap[complement]; ok { // if
			return []int{j, i}
		}
		/*
			Check if complement is in numMap: Checks if the complement value is in the numMap hash table.
			If yes (ok returns true), it means that we have found two numbers whose sum is equal to target:
			j: This is the index of the element that, together with nums[i], gives the right sum.
			Return result: If complement is found, return indices [j, i], where j is the index of the found element and i is the index of the current element.


		*/
		numMap[nums[i]] = i
		// Add the current element and its index to the numMap hash table
	}

	/*

		At each step, we add the current number (nums[i]) to numMap with its index.
		This allows us to keep track of the numbers we have already processed and their indexes. When we check a complement,
		we look to see if the complement value exists in numMap. If the complement exists, it means that we already have a number that,
		when combined with the current number, gives the target sum. Thus, we find the desired indices and return them.

		Bottom line: We add nums[i] to numMap rather than complement,
		because nums[i] represents a value we have already processed that can be used in combination with subsequent values to achieve the target sum.

	*/

	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)

	result1 := twoSum([]int{3, 2, 4}, 6)

	fmt.Println(result)

	fmt.Println(result1)

	result2 := twoSum1([]int{1, 5, 11, 15, 26}, 6)
	fmt.Println(result2)

	result3 := twoSum1([]int{3, 2, 4, 3}, 6)

	fmt.Println(result3)

}
