package main

import (
	"fmt"
	"time"
)

// struct

type Customer struct {
	name string
	age int
}
type Order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
	customer Customer  // composition
}

// constructors
func newOrder(id string,amount float32,status string) *Order{
	myOrder := Order{
		id: id,
		amount: amount,	
		status: status,
		createdAt: time.Now(),
	}

	return &myOrder
}

func (o *Order) changeStatus(status string){ // this function is now attached with struct
	 o.status = status
}

func (o *Order) getAmount() float32{ // this function is now attached with struct
	 return  o.amount
}

func main(){
	customer := Customer{
		name:"John",
		age:22,
	}
	order := Order {
		id: "1",
		amount: 50.00,
		status: "received", 
		customer:customer,
	}

	// if you skip any value then it will default to falsy value of that type like for int it will be 0 , for string it will be empty string for object it will be nil
	newOrder := Order {
		id:"2",
		amount:54.00,
		createdAt: time.Now(),
	}

	order.createdAt = time.Now()

	fmt.Println(order.id)
	newOrder.changeStatus("Paid")

	fmt.Println("Order Struct",order,newOrder)
}