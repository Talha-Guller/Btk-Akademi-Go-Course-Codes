package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	adress             string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carİnfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 100 / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credits1 := Mortgage{rate: 10, creditPaymentTotal: 100000, adress: "ankara"}
	credits2 := Mortgage{rate: 12, creditPaymentTotal: 50000, adress: "ankara"}
	credits3 := Car{rate: 15, creditPaymentTotal: 60000, carİnfo: "polo"}

	credits := []CreditCalculator{credits1, credits2, credits3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("Toplam ödenecek tutar :", total)
}
