package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Mocha"
	coffees[4] = "Espresso"
	coffees[5] = "Machiatto"
	coffees[6] = "Americano"
	coffees[7] = "Short"
	coffees[8] = "Long"

	fmt.Println("COFFEE MENU")
	fmt.Println("-----------")
	fmt.Println("1- Cappuccino")
	fmt.Println("2- Latte")
	fmt.Println("3- Mocha")
	fmt.Println("4- Espresso")
	fmt.Println("5- Machiatto")
	fmt.Println("6- Americano")
	fmt.Println("7- Short")
	fmt.Println("8- Long")
	fmt.Println("Q- Quit the program")

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")
	for {

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {

			break
		}

		i, _ := strconv.Atoi(string(char))

		t := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t)

	}
	fmt.Println("Program ended")
}
