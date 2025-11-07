package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main(){
	// creating map
	
	temp := make(map[string]string)
	fmt.Println(temp["Hello"] == "") // no such key hence will get empty string

	m:= make(map[string]int) // ordered

	m["area"] = 1
	m["pare"] = 1
	m["case"] = 1

	delete(m,"area")
	clear(m)
	fmt.Println(m,m["phone"]) // no such key hence will get 0

	mp := map[string]int{"price":40,"phones":3}
	fmt.Println(mp)
	
	value,ok := m["price"]

	fmt.Println("value",value)
	if ok  {
		fmt.Println("All OK")
	} else {
		fmt.Println("Not OK")
	}




	mp1 := map[string]int{"price":40,"phones":3}
	mp2 := map[string]int{"price":40,"phones":3}


	fmt.Println(maps.Equal(mp1,mp2))

	
}