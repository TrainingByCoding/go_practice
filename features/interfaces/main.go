package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct {
	cardno string
}
type Paypal struct {
	email string
}

func ProcessPayment(pp PaymentMethod, amount float64) {
	result := pp.Pay(amount)
	fmt.Println(result)
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("paid %.2f using creditcard, no ending with %s", amount, cc.cardno[len(cc.cardno)-4:])
}
func (pp Paypal) Pay(amount float64) string {
	return fmt.Sprintf("paid %.2f using paypal with email %s", amount, pp.email)
}

func main() {
	// Creating instances of CreditCard and PayPal
	mycard := CreditCard{cardno: "asdfdsfdsf"}
	mypaypal := Paypal{email: "test@gmail.com"}

	ProcessPayment(mycard, 100)
	ProcessPayment(mypaypal, 200)

	// Create an interface variable of type PaymentMethod
	var pm PaymentMethod
	pm = CreditCard{cardno: "asdfsdfdsf"}
	ProcessPayment(pm, 1000)
	pm = Paypal{email: "test@gmail.com"}
	ProcessPayment(pm, 2000)
}
