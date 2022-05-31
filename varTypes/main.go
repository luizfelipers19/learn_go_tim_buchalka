package main

import "fmt"

func main() {

	//basic variable types
	var myInt int
	var myInt64 int64
	var myFloat float64

	var myBool bool
	var myString string

	myInt = 150
	myInt64 = 32022
	myFloat = 19.19

	myBool = true
	myString = "Isso é uma string"

	fmt.Println(myInt, myInt64, myFloat, myBool, myString)

	// Aggregate variable types

	//array
	var myStrings [3]string
	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("The first element of the array is", myStrings[0])

	//structs

	type Car struct {
		numberOfTires int
		luxury        bool
		numberOfSeats int
		brand         string
		model         string
		year          int
	}

	var teslaS Car
	teslaS.model = "S"
	teslaS.year = 2020
	teslaS.brand = "Tesla"
	teslaS.numberOfSeats = 4
	teslaS.numberOfTires = 4
	teslaS.luxury = true

	fmt.Printf("My car is a %s %s from %d, has %d seats and %d tires, and it's %t that it's a luxury car", teslaS.brand, teslaS.model, teslaS.year, teslaS.numberOfSeats, teslaS.numberOfTires, teslaS.luxury)

}
