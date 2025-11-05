package main

import "fmt"


func main(){
	age := 87
	if age > 60 {
		fmt.Println("Senior")
	} else if age > 18 {
		fmt.Println("Adult")
	} else{
		fmt.Println("Teenager")
	}

	// golang does not have ternary
}	