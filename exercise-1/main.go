package main

import (
	"exercise-1/utils"
	"fmt"
)

// Say Hello xxx using a prompt
func main() {
	var name string

	fmt.Print("Entrez votre nom : ")
	fmt.Scanln(&name)

	message := utils.Greeting(name)
	fmt.Println(message)
}
