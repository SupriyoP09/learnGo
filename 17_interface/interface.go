package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

// we can create an interface for payment gateway and implement it in different payment gateway structs
func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

// we can also create other payment gateways like paypal, stripe etc and implement the same interface
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment using razorpay
	fmt.Println("making payment using razorpay", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fake payment gateway", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("refunding payment using razorpay", amount, account)
}


func main() {
	razorpayPaymentGw := razorpay{}
	// fakeGw := fakePayment{}
	// paypalGw := paypal{}
	// stripeGw := stripe{}
	newPayment := payment{
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
}
