package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

const ENTERMSG string = "don't inform the chosen number but press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	fmt.Println("Funciona")
	fmt.Println("GUESS THE NUMBER GAME! \n")

	playGame(firstNumber, secondNumber, subtraction, answer)

}

func playGame(firstNumber int, secondNumber int, subtraction int, answer int) {
	fmt.Println("Choose a number between 1 and 10 and", ENTERMSG)
	reader.ReadString('\n')
	fmt.Println("Multiply your number by ", firstNumber, " and ", ENTERMSG)
	reader.ReadString('\n')
	fmt.Println("Now multiply the result number by", secondNumber, "and", ENTERMSG)
	reader.ReadString('\n')

	fmt.Println("Now divide the result by the number you originally thought of", ENTERMSG)

	fmt.Println("Subtract the total value by", subtraction, "and", ENTERMSG)
	reader.ReadString('\n')

	fmt.Println("The final number is", answer)
}
