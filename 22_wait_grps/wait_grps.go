package main

import (
	"fmt"
	"sync"
)


func task(id int,w *sync.WaitGroup){
	defer w.Done() // defer will make it run at end of function call
	fmt.Println("Doing task",id)
}

func main(){
	var waitGroup sync.WaitGroup
	for i:= 0;  i <= 10; i++{
		waitGroup.Add(1)
		go task(i,&waitGroup) // coroutines running in parallel with light weight thread
	}

	waitGroup.Wait()
	
}