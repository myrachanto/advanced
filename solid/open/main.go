package main

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type CreditCardProcessor struct{}

func (ccp *CreditCardProcessor) ProcessPayment(amount float64) error {
	// process payment with credit card
	return nil
}

type PaypalProcessor struct{}

func (pp *PaypalProcessor) ProcessPayment(amount float64) error {
	// process payment with PayPal
	return nil
}
func main() {
	var processor PaymentProcessor

	processor = &CreditCardProcessor{}
	processor.ProcessPayment(100.0)

	processor = &PaypalProcessor{}
	processor.ProcessPayment(200.0)
}
