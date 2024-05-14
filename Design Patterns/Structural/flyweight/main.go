package main

import "fmt"

//Flyweight Design Pattern
const (
	NikeAirForce = "Nike Air Force 1"

	AdidasSuperstar = "Adidas Superstar"
)

func main() {
	sneaker1 := Sneaker{
		shoeType: NikeAirForce,
		size:     9,
		color:    "white",
	}

	sneaker2 := Sneaker{
		shoeType: AdidasSuperstar,
		size:     8,
		color:    "black",
	}

	fmt.Println(sneaker1.describe())
	fmt.Println(sneaker2.describe())
}
