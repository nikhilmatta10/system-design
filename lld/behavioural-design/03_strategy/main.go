package main

import "fmt"

type PaymentStrategy interface {
	processPayment()
}

type PaymentService struct {
	strategy PaymentStrategy
}

func (ps *PaymentService) setPaymentStrategy(strategy PaymentStrategy) {
	ps.strategy = strategy
}

func (ps *PaymentService) pay() {
	ps.strategy.processPayment()
}

type CreditCardPayment struct{}

func (ccp *CreditCardPayment) processPayment() {
	fmt.Println("Making paymetn via credit card")
}

type DebitCardPayment struct{}

func (dcp *DebitCardPayment) processPayment() {
	fmt.Println("Making payment via debit card")
}

func main() {
	paymentService := PaymentService{}
	paymentService.setPaymentStrategy(&CreditCardPayment{})
	paymentService.pay()
}