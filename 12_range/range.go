package main

import (
	"fmt"
	
)

func main(){
	nums := []int {6,7,8}
	for i := 0; i < len(nums);i++{
		fmt.Println(nums[i])
	}


	// using range

	sum := 0

	for idx,num := range nums {
		fmt.Println(num,idx)
		sum+=num
	}

	fmt.Println(sum)



	mp := map[string]int{"price":40,"phones":3,"fruits":20}

	for key,it := range mp {
		fmt.Println(key,it)
	}


	for startingByte,charCode := range "golang" {
		fmt.Println(startingByte,string(charCode))
	}

	
}