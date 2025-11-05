package main

import "fmt"

// in go there is only one way to loop => for loop
// we can also use continue and break here


func main(){

	// for i:=0;i<=5;i++ {
	// 	fmt.Println(i)
	// }


	// i:=1
	// // simulated while loop
	// for i <= 3 {
	// 	if i == 2 {
	// 		i++
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }


	// range

	for i:= range 5{  // 3 is excluded
		fmt.Println(i)
	}
}
