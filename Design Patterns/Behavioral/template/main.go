package main

import "fmt"

//Template Design Pattern
type OrderProcessor interface {
	ProcessOrder()
}

type BaseOrderProcessor struct{}

func (b *BaseOrderProcessor) ProcessOrder() {
	b.validateOrder()
	b.fetchItems()
	b.calculateTotal()
	b.confirmOrder()
}

func (b *BaseOrderProcessor) validateOrder() {
	fmt.Println("Step 1: Validating order...")
}

func (b *BaseOrderProcessor) fetchItems() {
	fmt.Println("Step 2: Fetching items from inventory...")
}

func (b *BaseOrderProcessor) calculateTotal() {
	fmt.Println("Step 3: Calculating order total...")
}

func (b *BaseOrderProcessor) confirmOrder() {
	fmt.Println("Step 4: Confirming order...")
}

type CustomOrderProcessor struct {
	BaseOrderProcessor
}

func (c *CustomOrderProcessor) validateOrder() {
	fmt.Println("Custom Step 1: Custom validation for the order...")
}

func main() {

	orderProcessor := &CustomOrderProcessor{}

	orderProcessor.ProcessOrder()
}
