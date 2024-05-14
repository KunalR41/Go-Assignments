package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Command_Desing_Pattern
type Order interface {
	Execute()
}

type Stock struct {
	name  string
	count int
}

func (s *Stock) Buy() {
	fmt.Println("Bought new item.")
	s.count++
}

func (s *Stock) Sell() {
	fmt.Println("Sold one item.")
	s.count--
}

type BuyStock struct {
	Stock *Stock
}

func (b *BuyStock) Execute() {
	b.Stock.Buy()
}

type SellStock struct {
	Stock *Stock
}

func (s *SellStock) Execute() {
	s.Stock.Sell()
}

type Broker struct {
	Orders []Order
}

func (b *Broker) TakeOrder(o Order) {
	b.Orders = append(b.Orders, o)
}

func (b *Broker) PlaceOrders() {
	for _, v := range b.Orders {
		v.Execute()
	}

	b.Orders = make([]Order, 0)
}

func buyStockHandler(w http.ResponseWriter, r *http.Request) {
	stock := &Stock{name: "phone", count: 2}
	buyStock := &BuyStock{Stock: stock}
	broker := &Broker{}

	broker.TakeOrder(buyStock)
	broker.PlaceOrders()

	response := map[string]string{"message": "Buy order executed successfully"}
	json.NewEncoder(w).Encode(response)
}
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WELCOME To Stock Api....!!!")
}
func sellStockHandler(w http.ResponseWriter, r *http.Request) {
	stock := &Stock{name: "phone", count: 2}
	sellStock := &SellStock{Stock: stock}
	broker := &Broker{}

	broker.TakeOrder(sellStock)
	broker.PlaceOrders()

	response := map[string]string{"message": "Sell order executed successfully"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/buy", buyStockHandler)
	http.HandleFunc("/api/sell", sellStockHandler)
	//http://localhost:8000
	fmt.Println("API server started on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
