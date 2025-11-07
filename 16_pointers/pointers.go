package main

import "fmt"


// call by value
func changeNum(num int){
	num = 5
	fmt.Println("In changeNum",num)
}

// call by reference

func changeNum2(num *int){
	*num = 50

	fmt.Println("In changeNum2",*num)
}

func main(){
	num := 1
	
	changeNum(num)
	fmt.Println("After changeNum in main",num)
	changeNum2(&num)
	fmt.Println("After changeNum2 in main",num)

}