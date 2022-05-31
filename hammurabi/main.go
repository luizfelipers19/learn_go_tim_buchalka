package main

import (
	"fmt"
	"myapp/game"
)

func main() {
	playagain := true

	for playagain {
		game.Play()
		playagain = game.GetYesOrNo("Would you like to play again? (y/n)")
	}

	fmt.Println("")
	fmt.Println("Good bye! See you later!")
}
