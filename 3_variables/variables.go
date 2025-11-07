package main

import "fmt"


func main(){
	var name string = "Golang" // type defined
	var age = 5 // type is inferred from the value
	var isAdult = age > 18


	// shorthand index
	city := "Delhi"
	var country string
	country = "India"

	fmt.Println(name,age,isAdult,city,country)
}

