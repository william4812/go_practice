package main

import (
	"fmt"
	"reflect"
)

/*
Given an array of integers nums sorted in non-decreasing order,
find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func searchRange(nums []int, target int) []int {
	// O(log n) using binary or divide and conquer paradign
	nSize := len(nums)
	left := 0          // index Left to the target
	right := nSize - 1 // index Right to the target

	for left < right {
		mid := left + (right-left)/2 // middle of Left and Right
		if target > nums[mid] {
			// target at the Right of mid, update left to middle
			left = mid + 1
		} else if target < nums[mid] {
			// target at the Left of mid, update right to middle
			right = mid - 1
		} else if target == nums[mid] {
			// nums[mid] hit a target, then search from mid to see
			// the range of indices covering the other targets

			// search Left to the mid
			left = mid
			for left > 0 && nums[left-1] == target {
				left--
			}

			// search Right to the mid
			right = mid
			for right < nSize && nums[right+1] == target {
				right++
			}

			//fmt.Println("answer", left, right)
			return []int{left, right} // found the range
		}
	}

	return []int{-1, -1} // target not found
}

func test() {
	fmt.Println("testing...")

	// test 1
	n1 := []int{5, 7, 7, 8, 8, 10}
	t1 := 8
	fmt.Println("Passing? %t",
		reflect.DeepEqual(searchRange(n1, t1), []int{3, 4}))

	// test 2
	n2 := []int{5, 7, 7, 8, 8, 10}
	t2 := 6
	fmt.Println("Passing? %t",
		reflect.DeepEqual(searchRange(n2, t2), []int{-1, -1}))

	// test 3
	n3 := []int{}
	t3 := 0
	fmt.Println("Passing? %t",
		reflect.DeepEqual(searchRange(n3, t3), []int{-1, -1}))

	fmt.Println("test ends!")
}

func main() {
	test()
}
