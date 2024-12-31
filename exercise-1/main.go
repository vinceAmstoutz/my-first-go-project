package main

import (
	"exercise-1/utils"
	"fmt"
)

// Say Hello xxx using a prompt
func main() {
	var name string

	fmt.Print("Entrez votre nom : ")

	_, error := fmt.Scanln(&name)
	if error != nil {
		fmt.Println("[ERROR] Failed to read input:", error)
		return
	}

	message := utils.Greeting(name)
	fmt.Println(message)
}
