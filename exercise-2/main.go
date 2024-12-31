package main

import (
	"exercise-2/utils"
	"fmt"
)

// The aim of this program is to create a guessing game where the user tries to find a mystery number chosen in advance by the program.
// The aim is to master the basics of Go, such as loops, conditions and user interaction.

// What the program should contain :
// - A number to guess: A fixed number (e.g. 42) that the user must guess.
// - Welcome messages: Instructions to tell the user how to play.
// - A trial counter: A variable that counts the number of attempts.
// - Game loop: A loop that continues until the mystery number is found.
// - Comparisons: Clues (larger or smaller) to guide the user.
// - Victory message: A message congratulating the user with the number of trials used.
// - Bonus : allow replay by adding a loop to restart a game after winning.
func main() {
	displayWelcomeMessage()
	startGame()
}

func displayWelcomeMessage() {
	welcomeMessage := utils.GetWelcomeMessage()
	fmt.Println(welcomeMessage)
}

func startGame() {

forPlay:
	for {
		var attempts uint8
		mysteryNumber := utils.GetMysteryNumber()

	forGame:
		for {
			attempts++
			inputNumber := utils.GetUserInput()

			if inputNumber == mysteryNumber {
				break forGame
			}

			provideFeedback(inputNumber, mysteryNumber)
		}

		fmt.Printf("Congratulations! 🎉 You've uncovered the mystery number %d in just %d attempts!\n", mysteryNumber, attempts)

		shouldContinue := utils.ShouldRestartGame()
		if !shouldContinue {
			break forPlay
		}

		continue
	}
}

func provideFeedback(inputNumber, mysteryNumber uint8) {
	if inputNumber < mysteryNumber {
		fmt.Println("The mystery number is higher.")
	} else {
		fmt.Println("The mystery number is lower.")
	}
}
