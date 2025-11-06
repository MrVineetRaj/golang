package main

import "fmt"

func add(a int,b int) int{
	return a + b
}


func getLanguages() (string,string,string){
	return"golang","js","c++"
}

func processIt(fn func(a int) int) {
	fmt.Println(fn(1))
}

func main(){
	fmt.Println(add(2,5))


	lang1, lang2, _ := getLanguages()

	fmt.Println(lang1,lang2)


	fn := func(a int) int {
		return 2
	}

	processIt(fn)
}