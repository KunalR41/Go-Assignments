package main

import "encoding/json"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64 `json:"radius"`
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func DecodeShape(data []byte) (Shape, error) {
	var shape Shape
	err := json.Unmarshal(data, &shape)
	return shape, err
}
