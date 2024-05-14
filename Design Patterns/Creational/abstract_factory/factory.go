package main

//Abstract Factory
type Factory interface {
	CreateShape(params map[string]float64) Shape
}

type RectangleFactory struct{}

func (rf RectangleFactory) CreateShape(params map[string]float64) Shape {
	return Rectangle{
		Width:  params["width"],
		Height: params["height"],
	}
}

type CircleFactory struct{}

func (cf CircleFactory) CreateShape(params map[string]float64) Shape {
	return Circle{
		Radius: params["radius"],
	}
}
