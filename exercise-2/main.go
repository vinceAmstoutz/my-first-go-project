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

func main() {
	var attempts uint8
	mysteryNumber := utils.GetMysteryNumber()

	displayWelcomeMessage()
	startGame(mysteryNumber, &attempts)

	fmt.Printf("Congratulations! 🎉 You've uncovered the mystery number %d in just %d attempts!\n", mysteryNumber, attempts)

}

func displayWelcomeMessage() {
	welcomeMessage := utils.GetWelcomeMessage()
	fmt.Println(welcomeMessage)
}

func startGame(mysteryNumber uint8, attempts *uint8) {
	for {
		(*attempts)++
		inputNumber := utils.GetUserInput()

		if inputNumber == mysteryNumber {
			break
		}

		provideFeedback(inputNumber, mysteryNumber)
	}
}

func provideFeedback(inputNumber, mysteryNumber uint8) {
	if inputNumber < mysteryNumber {
		fmt.Println("The mystery number is higher.")
	} else {
		fmt.Println("The mystery number is lower.")
	}
}
