package main

import "fmt"

func main() {

	// FLOW CONTROL

	//LOOPS

	//FOR
	// tree part loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for j := 10; j > 0; j-- {
		fmt.Println("Going to 0: ", j)
	}

	//While loop
	t := 1
	for t < 50 { // while loop is done with a for loop kk
		fmt.Println("Raising my value to 50: ", t)
		t++
	}
	// infinite loop
	for {
		fmt.Println("Infinitely printing stuff")
	}
}
