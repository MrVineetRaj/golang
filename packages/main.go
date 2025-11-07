package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mrvineetraj/golang/auth"
	"github.com/mrvineetraj/golang/user"
)

func main(){
	auth.LoginWithCred("mrvineetraj","abcd..00")
	color.Red(auth.GetSession())

	user := user.User{
		Email: "user@email.com",
		Name: "John",
	}

	fmt.Println(user)
}