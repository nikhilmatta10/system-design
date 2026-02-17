package main

import "fmt"

type PaymentService struct{}

func (ps *PaymentService) processPayment(paymentMethod string) {
	switch paymentMethod {
		case "Credit Card":
				fmt.Println("making payment via credit card")
		case "Debit Card":
				fmt.Println("making payment via debit card")
		default:
			fmt.Println("method unsupported")
	}
}

func main() {
	payService := PaymentService{}
	payService.processPayment("Credi1t Card")
}