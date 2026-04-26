package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used contruct in go
// + usefull method
func main() {
	// uninitalised slices is nil
	// var nums []int
	// fmt.Println(nums)

	// fmt.Println(len(nums))



	// var nums = make([]int, 0, 5)

	// capacity -> maximum number  of element can fit
	// fmt.Println(cap(nums))
	// fmt.Println(nums) 
	// fmt.Println(nums == nil)

	// nums := []int{}
	// nums = append(nums, 1)
	// nums = append(nums, 2)

	// var nums = make([]int, 2, 5)
	// nums[0] = 3
	// nums[1] = 4
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)

	// var nums2 = make([]int, len(nums))
	// // copy function

	// copy(nums2, nums)

	// fmt.Println(nums, nums2)


	// slice operator
	// var nums = []int{1,2,3,4}

	// fmt.Println(nums[0:2])
	// fmt.Println(nums[:2])
	// fmt.Println(nums[1:])

	// slice package
	var nums1 = []int{1,2,3}
	var nums2 = []int{1,2,4}

	fmt.Println(slices.Equal(nums1, nums2))

	//2d slices
	// var nums = [][]int{{1,2,3}, {4,5,6}}
	// fmt.Println(nums)

}