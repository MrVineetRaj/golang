package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func processNum(numChan chan int){
	for num := range numChan {
		fmt.Println("Processing number", num)
	}
}

func main(){
// 	messageChannel := make(chan string)

// 	messageChannel <- "ping" // channels are blocking

// 	// as blocking channel code could not move to next line hence messageChannel data is never sent out of pip hence causing deadlock

// 	msg := <- messageChannel

// 	fmt.Println(msg)


	numChan := make(chan int)
	go processNum(numChan)

	for{
		numChan <- rand.IntN(100)
	time.Sleep(time.Second * 1)
	}


}