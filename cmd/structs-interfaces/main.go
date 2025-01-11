package main

import "fmt"

// this defines a type called gasEngine which is a struct with two fields
type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
}

type gasEngine2 struct {
	mpg uint8
	gallons uint8
	// the other way is to do owner just type ( avoids doing myEngine.---.---)
	// instead myEngine.---
	owner
	number int
	// or do 
	// int
}

type owner struct {
	name string
}

// interface which both engine1 and engine2 types use to call the func canMakeit which
// now accepts an engine type which all you need to use the canMakeit function is simply
// a type that has the method milesLeft

type engine interface {
	// define the signature needed
	// return type of uint16
	milesLeft() uint16
}

// anynonmous structs types define and initialize in same location (Line 46)

// to make the method to find gas left in tank make a function that is assigned to a type
// in this case the gasEngineType
// same as function has parameter empty func name and return needs to be uint8 type
func (e gasEngine) milesLeft() uint16 {
	// however the (e gasEngine) will assign the function to the gasEngine type
	// has access to the attributes of the gasEngine type
	// like class methods attributes
	return uint16(e.gallons) * uint16(e.mpg)
}

func (e gasEngine2) milesLeft() uint16 {
	// however the (e gasEngine) will assign the function to the gasEngine type
	// has access to the attributes of the gasEngine type
	// like class methods attributes
	return uint16(e.gallons) * uint16(e.mpg)
}

func canMakeIt(e engine, miles uint16) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it")
	}else{
		fmt.Println("You cant make it")
	}
}

func main() {
	var myEngine gasEngine
	// default value of 0 if not assigned
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	// the third parameter is the ownerInfo which is a struct of type owner with a field name
	var myEngine2 gasEngine = gasEngine{mpg:20, gallons:15, ownerInfo: owner{name:"ahmad"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.ownerInfo.name)

	var myEngine3 gasEngine = gasEngine{20, 15, owner{"ahmad"}}
	fmt.Println(myEngine3.ownerInfo.name, myEngine3.gallons, myEngine3.mpg)

	var myEngine4 gasEngine2 = gasEngine2{20, 15, owner{"nayef"}, 10}
	// can also do it with just the type and do myEngine4.int
	fmt.Println(myEngine4.name, myEngine4.mpg, myEngine4.gallons, myEngine4.number)

	// anynomous struct meaning cant be reused initialized and defined smae locat
	var myEngine5 = struct{
		mpg uint8
		gallons uint8
	}{25, 15}

	fmt.Println(myEngine5.mpg, myEngine5.gallons)

	// object myEngine2 of type myEngine has a function milesLeft which expects to be called
	// by a type myEngine since func declaration initialized to the type myEngine
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
	fmt.Printf("Total miles left: %d\n", myEngine2.milesLeft())

	fmt.Println(myEngine4.mpg, myEngine4.gallons)
	fmt.Printf("Total miles left: %d\n", myEngine4.milesLeft())
	canMakeIt(myEngine2, 200)
	// because canMakeIt expects myEngine type not myEngine2 type we need interface to use that func for both types of engines
	canMakeIt(myEngine4, 100)
}