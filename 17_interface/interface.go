package main

import "fmt"

// payment interface
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	paymentService paymenter
}

func (p payment) pay(amount float32) {
	p.paymentService.pay(amount)
}

// payment service

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Printf("Payment of %.2f made successfully using Razorpay\n", amount)
}

// stripe service
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Printf("Payment of %.2f made successfully using Stripe\n", amount)
}

func main() {
	p := payment{
		paymentService: razorpay{},
	}
	p.pay(100.50)

}
