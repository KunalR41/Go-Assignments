package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Strategy Design pattern
type PaymentStrategy interface {
	ProcessPayment(amount float64) string
}

type CreditCardPayment struct{}

func (cc *CreditCardPayment) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Paid %.2f using credit card", amount)
}

type NetBankingPayment struct{}

func (nb *NetBankingPayment) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Paid %.2f using NetBanking", amount)
}

type UPIPayment struct{}

func (upi *UPIPayment) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Paid %.2f using UPI", amount)
}

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{
		strategy: strategy,
	}
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) string {
	return pp.strategy.ProcessPayment(amount)
}
func CreditCardHandler(w http.ResponseWriter, r *http.Request) {
	amount := 1000.0
	processor := NewPaymentProcessor(&CreditCardPayment{})
	message := processor.ProcessPayment(amount)

	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func NetBankingHandler(w http.ResponseWriter, r *http.Request) {
	amount := 1000.0
	processor := NewPaymentProcessor(&NetBankingPayment{})
	message := processor.ProcessPayment(amount)

	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UPIHandler(w http.ResponseWriter, r *http.Request) {
	amount := 1000.0
	processor := NewPaymentProcessor(&UPIPayment{})
	message := processor.ProcessPayment(amount)

	response := map[string]string{"message": message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/pay/creditcard", CreditCardHandler)
	http.HandleFunc("/pay/netbanking", NetBankingHandler)
	http.HandleFunc("/pay/upi", UPIHandler)

	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
