package main

type sneakerDetails struct {
	name      string
	brand     string
	price     float64
	materials []string
}

var NikeAirForceDetails = sneakerDetails{
	name:      "Nike Air Force 1",
	brand:     "Nike",
	price:     144.99,
	materials: []string{"leather", "rubber"},
}

var AdidasSuperstarDetails = sneakerDetails{
	name:      "Adidas Superstar",
	brand:     "Adidas",
	price:     85.99,
	materials: []string{"leather", "rubber"},
}

var DetailsMap = map[string]sneakerDetails{
	NikeAirForce:    NikeAirForceDetails,
	AdidasSuperstar: AdidasSuperstarDetails,
}

func getShoeDetails(shoeType string) sneakerDetails {
	return DetailsMap[shoeType]
}
