package main

import (
	"fmt"
	"net/http"
)

// Decorator Design Pattern
type IPizza interface {
	getPrice() int
}
type VeggieMania struct{}

func (p *VeggieMania) getPrice() int {
	return 50
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 25
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 25
}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {

	pizza := &VeggieMania{}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	price := pizzaWithCheeseAndTomato.getPrice()

	fmt.Fprintln(w, "~~~Welcome to IPizza Store~~~")
	fmt.Fprintf(w, "Price of VeggieMania with tomato and cheese topping is %d\n", price)
}

func main() {
	http.HandleFunc("/price", getPriceHandler)
	fmt.Println("Server is running on port 8888...")
	http.ListenAndServe(":8888", nil)
}
