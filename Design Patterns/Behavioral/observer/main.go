package main

import (
	"fmt"
)

// Observer Design pattern
type Observer interface {
	Update(data interface{})
}
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data interface{})
}

type Publisher struct {
	observers []Observer
}

func NewPublisher() *Publisher {
	return &Publisher{
		observers: make([]Observer, 0),
	}
}
func (p *Publisher) RegisterObserver(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *Publisher) RemoveObserver(observer Observer) {
	for i, obs := range p.observers {
		if obs == observer {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *Publisher) NotifyObservers(data interface{}) {
	for _, observer := range p.observers {
		observer.Update(data)
	}
}

type Subscriber struct {
	ID int
}

func NewSubscriber(id int) *Subscriber {
	return &Subscriber{
		ID: id,
	}
}
func (s *Subscriber) Update(data interface{}) {
	fmt.Printf("Subscriber %d received update: %v\n", s.ID, data)
}

func main() {
	publisher := NewPublisher()

	subscriber1 := NewSubscriber(1)
	subscriber2 := NewSubscriber(2)

	publisher.RegisterObserver(subscriber1)
	publisher.RegisterObserver(subscriber2)

	eventData := "Event data"
	publisher.NotifyObservers(eventData)
}
