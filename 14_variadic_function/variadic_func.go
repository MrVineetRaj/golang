package main

import "fmt"


// variadic functions: function spread operator
func sum(nums...int)int{
	total := 0

	for _,num:= range nums{
		total+=num
	}

	return  total
}



func main(){
	result := sum(1,2,5,2)
	fmt.Println(result)

	nums:=[]int{3,4,5,6,5}

	fmt.Println(sum(nums...))
}