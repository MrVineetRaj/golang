package main

import (
	"fmt"
)

func printSlice[T int | string](items []T){
	fmt.Println(items)
}


// LIFO
type Stack[t comparable] struct {
	elements []int
}
func main(){
	// nums := []int{1,2,3}

	// strs := []string{"Golang","System","Go"}

	st := Stack[int]{
		elements: []int{1,2,3,4},
	}


	fmt.Println(st)
	

}