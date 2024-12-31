package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const MaxMysteryNumber = 100

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
	n, err := rand.Int(rand.Reader, big.NewInt(MaxMysteryNumber+1))
	if err != nil {
		fmt.Println("[ERROR] Failed to generate random number:", err)
		return 0
	}

	return uint8(n.Int64())
}

func GetUserInput() uint8 {
	var inputNumber uint8

	for {
		fmt.Print("Enter a number between 0 and 100: ")

		_, error := fmt.Scanln(&inputNumber)
		if error != nil || inputNumber > 100 {
			fmt.Print("[ERROR] Invalid input, enter a valid number between 0 and 100!\n")

			// Clear the invalid input from the buffer
			var discard string
			if _, discardErr := fmt.Scanln(&discard); discardErr != nil {
				fmt.Println("[ERROR] Unable to clear the buffer.")
			}

			continue
		}

		return inputNumber
	}
}

func ShouldRestartGame() bool {
	var input string

	for {
		fmt.Print("Do you want to start a new game? (yes/no, default: yes): ")

		_, error := fmt.Scanln(&input)
		if error != nil {
			fmt.Println("[ERROR] Failed to read input.")
			continue
		}

		switch input {
		case "yes", "":
			return true
		case "no":
			return false
		default:
			fmt.Println("[ERROR] Please enter 'yes' or 'no'.")
			continue
		}
	}
}
