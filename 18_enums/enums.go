package main

import "fmt"

// Enums in Go are typically implemented using a custom type and a set of constants.
type OrderStatus string

// Define constants for the OrderStatus type.
const (
	Received  OrderStatus = "received"
	Confirmed OrderStatus = "confirmed"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
)

// Function to change the order status.
func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

// The main function demonstrates how to use the OrderStatus enum.
func main() {
	changeOrderStatus(Shipped)
}
