package main

import "fmt"

// we can declare variables outside the functions as well with var and const

func main(){
	const name = "Golang" // can not be re assigned

	const (
		port = 5000
		host = "host"
	)

	fmt.Println(port,host)
}