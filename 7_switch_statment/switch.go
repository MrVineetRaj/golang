package main

import (
	"fmt"
	"time"
)

func main(){

	i := 8

	switch i {
	case 1 :
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Invalid number")
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default :
		fmt.Println("Workday")
	}
}
