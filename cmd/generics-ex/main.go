package main

import (
	"fmt"
)

// generics can be used with struct types not only functions

type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

// T is a type parameter
// T can be either gasEngine or electricEngine
type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Toyota",
		carModel: "Corolla",
		engine: gasEngine{
			gallons: 10,
			mpg:     30,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   60,
			mpkwh: 4,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)

}
