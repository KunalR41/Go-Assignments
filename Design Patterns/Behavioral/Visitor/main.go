package main

import "fmt"

// Visitor Design Pattern
type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(c)
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(c)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visited ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visited ConcreteElementB")
}

func main() {

	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	visitor := &ConcreteVisitor{}

	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
