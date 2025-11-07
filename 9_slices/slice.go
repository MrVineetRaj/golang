package main

import "fmt"

// slices
// Mostly used construct
// dynamic size
// it comes with many useful methods
func main(){
	// uninitialized slice is nil
	var nums []int;


	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var nums1 = make([]int,2,5) // not nill | 5 is initial capaicut
	nums1 = append(nums1, 1)
	nums1 = append(nums1, 3)
	nums1 = append(nums1, 4)
	nums1 = append(nums1, 50)
	fmt.Println(len(nums1))
	// capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nums1))


	nums1 = make([]int,0)

	for i:=0; i < 10;i++{
		nums1 = append(nums1, i+1)
	}

	nums = make([]int, len(nums1)) // Allocate space for nums
	copy(nums, nums1) // before calling make sure that both slices is of same length

	nums1[2] = 200;
	fmt.Println(nums1,nums)
	fmt.Println(nums[0:2])  // print element from index 0 to 2(index 2  is excluded)
	fmt.Println(nums[:2])  // 0 index is default start
	fmt.Println(nums[5:])  // last index is default last

	
}