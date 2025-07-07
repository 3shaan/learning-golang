package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
}

func newOrder(id int, amount float32, status string) Order {
	return Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// get amount
func (o Order) getAmount() float32 {
	return o.amount
}

func (o *Order) changeStatus(newStatus string) {
	o.status = newStatus
	fmt.Println("Order status changed to:", o.status)
	
}

func main() {
	myOrder := newOrder(1, 100.50, "Pending")

	myOrder.createdAt = time.Now()

	myOrder.changeStatus("Shipped")
	// Using the method to get the amount
	amount := myOrder.getAmount()
	fmt.Println("Order Amount:", amount)
	// Accessing fields of the struct
	fmt.Println("Order ID:", myOrder)

	// direct struct initialization
	language := struct {
		name   string
		isGood bool
	}{
		name:   "Go",
		isGood: true,
	}
	fmt.Println("Language:", language.name, "Is Good:", language.isGood)

}
