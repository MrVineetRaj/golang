package main

import "fmt"

func counter() func() int{
	count := 0;

	return func() int {
		count+=1
		return count
	}
}
func main(){
	ctr1 := counter()
	ctr2 := counter()

	fmt.Println(ctr1())
	fmt.Println(ctr1())
	fmt.Println(ctr2())
	fmt.Println(ctr1())
}