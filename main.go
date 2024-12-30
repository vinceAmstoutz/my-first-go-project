package main

import (
	"fmt"
	"my-first-go-project/utils"
)

func main() {
	var name string

	fmt.Print("Entrez votre nom : ")
	fmt.Scanln(&name)

	message := utils.Greeting(name)
	fmt.Println(message)
}
