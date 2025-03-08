package main

import "fmt"

type Payment_Metthod interface {
	pay(amount float64)
}

// for Bitcoin
type Bitcoin struct {
}

func NewBitcoin() *Bitcoin {
	return &Bitcoin{}
}
func (b *Bitcoin) pay(amount float64) {
	fmt.Printf("Paid $%.2f using Bitcoin\n", amount)
}

// for Credit Card
type CreditCard struct {
}

func NewCreditCard() *CreditCard {
	return &CreditCard{}
}
func (c *CreditCard) pay(amount float64) {
	fmt.Printf("Paying with Bitcoin: %f\n", amount)
}

// for Paypal
type Paypal struct {
}

func NewPaypal() *Paypal {
	return &Paypal{}
}

func (p *Paypal) pay(amount float64) {
	fmt.Printf("Paying with Bitcoin: %f\n", amount)
}

// PaymentProcessor Struct

type PaymentProcessor struct {
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{}
}

func (p *PaymentProcessor) ProcessPayment(method Payment_Metthod, amount float64) {
	method.pay(amount)
}

func main() {
	processor := NewPaymentProcessor()
	bitcoin := NewBitcoin()
	creditCard := NewCreditCard()
	paypal := NewPaypal()
	processor.ProcessPayment(bitcoin, 12.3)
	processor.ProcessPayment(paypal, 545.3)
	processor.ProcessPayment(creditCard, 342.3)

}
