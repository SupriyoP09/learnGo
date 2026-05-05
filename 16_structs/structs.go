package main

import (
	"fmt"
	"time"
)

// order structs

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // time.Time is a struct from the time package // nano second precision
	customer
}

// func newOrder(id string, amount float64, status string) *order {
// 	// initializing a new order struct with the provided values
// 	myorder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myorder
// }

// // receiver function to change the status of an order
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float64 {
// 	return o.amount
// }

func main() {
	// we can initialize the embedded struct fields separately after initializing the parent struct
	// newCustomer := customer{
	// 	name: "Virat Kohli",
	// 	phone: "1818181818",
	// }

	// we can also initialize the embedded struct fields directly while initializing the parent struct
	newOrder := order{
		id: "1",
		amount: 100,
		status: "received",
		customer: customer{
			name: "Virat Kohli",
			phone: "1818181818",
		},
	}

	// we can also update the fields of the struct after it has been initialized
	newOrder.customer.name = "King Kohli"
	fmt.Println(newOrder)



	// myOrder := newOrder("18", 973, "received") // create a new order using the newOrder function

	// fmt.Println(myOrder.amount)

	// create an anonymous struct and initialize it with values
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"Go", true} // create an anonymous struct and initialize it with values

	// fmt.Println(language)

	// if u dont set any firld, default value will be assigned to the field
	// int => 0, float => 0, string => "", bool => false
	// myorder := order{
	// 	id:     "18",
	// 	amount: 973,
	// 	status: "received",
	// }

	// myorder.changeStatus("conformed") // call the changeStatus method to update the status of myorder
	// fmt.Println(myorder.getAmount()) // call the getAmount method to retrieve the amount of myorder

	// myorder.createdAt = time.Now() // assign the current time to the createdAt field

	// fmt.Println(myorder.status)

	// myorder2 := order{
	// 	id:        "7",
	// 	amount:    5,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myorder.status = "paid" // update the status field of myorder

	// fmt.Println("Order struct", myorder2)

	// fmt.Println("Order struct", myorder)
}