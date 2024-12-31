package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_MYSTERY_NUMBER = 100

func GetWelcomeMessage() string {
	return `Welcome to the Mystery Number Game! 🎉

The goal is to guess the secret number between 0 and 100.
You will get clues to help you find it.
For each attempt, you will be told if the mystery number is higher or lower than your guess.
Good luck! Let's start! 🚀
	`
}

func GetMysteryNumber() uint8 {
	return generateRandomNumberBetween0And100()
}

func generateRandomNumberBetween0And100() uint8 {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(randomSource)

	return uint8(randomGenerator.Intn(MAX_MYSTERY_NUMBER + 1)) // 0 <= number <= 100
}

func GetUserInput() uint8 {
	var inputNumber uint8
	for {
		fmt.Print("Enter a number between 0 and 100: ")

		_, errors := fmt.Scanln(&inputNumber)
		if errors != nil || inputNumber > 100 {
			fmt.Print("[ERROR] Invalid input, enter a valid number between 0 and 100 !\n")

			// Clear the invalid input from the buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		return inputNumber
	}
}
