package main

import "fmt"

// numbered sequence of specific length
// fixed size
func main(){
	var strs [4]string

	strs[0] = "Hello"
	strs[1] = "Hello"
	strs[2] = "Hello"
	strs[3] = "Hello"

	var nums  = [3]int{1,2,3}
	fmt.Println(nums)
	// array length
	// fmt.Println(len(nums))

	var grid = [2][2]int{{3,4},{5,6}}
	fmt.Println(grid)
}