package main

import (
	"fmt"
	"sort"
)

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
	// ------------------------------------------------------------
	// Aggregate variable types

	//array
	var myStrings [3]string
	myStrings[0] = "Cat"
	myStrings[1] = "Dog"
	myStrings[2] = "Fish"

	fmt.Println("The first element of the array is", myStrings[0])
	// ------------------------------------------------------------
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

	// ----------------------------------------
	// reference variable types

	//pointers

	x := 10
	// & references the address of a variable
	myFirstPointer := &x //storing the memory location of x to myFirstPointer

	fmt.Println("X is", x)
	fmt.Println("The value of my first pointer is", myFirstPointer)

	//* references the content of the variable that is addressed to that pointer
	*myFirstPointer = 15 //referencing the content of what is stored in that memory address and assigning a new value
	fmt.Println("Now, x is", x)
	changeValueOfPointerTo25(&x) //passing a reference to the variable we want to change, and not the variable itself
	fmt.Println("Now the x value has been changed by changing the content of a pointer, and X is", x)

	// ------------------------------

	// Slices
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals)

	//printing items using a for loop
	for _, x := range animals {
		fmt.Println(x)
	}
	fmt.Println("The second element is", animals[1])
	fmt.Println("But i will edit it right now")
	animals[1] = "sagu"
	fmt.Println("The second element is", animals[1])
	fmt.Println("The slice is", len(animals), "items long")
	fmt.Println("Are they sorted?", sort.StringsAreSorted(animals))
	sort.Strings(animals)
	fmt.Println("Are they sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
	animals = deleteFromSliceOfStrings(animals, 1, 2)
	fmt.Println(animals)

	// ------------------------------

	// Maps

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}
	delete(intMap, "four")
	fmt.Println("-----")

	for key, value := range intMap {
		fmt.Println(key, value)
	}
	checkKeyInMap(intMap, "four")

}

func changeValueOfPointerTo25(num *int) { //accepting a pointer
	*num = 25 //changing the value of the content of the address the pointer is pointing at
}

// deletes an element from i to j
func deleteFromSliceOfStrings(slice []string, i, j int) []string {
	return append(slice[:i], slice[j:]...)
}

func checkKeyInMap(m map[string]int, key string) {
	el, ok := m[key]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}
}
