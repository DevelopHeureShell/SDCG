package main

import (
	"errors"
	"fmt"
)

func main() {
	// Display "Calculator > ", then wait that collectAndTreatTheUserInput() send the entry to calculate.

	fmt.Println("Bienvenue dans SDCG - Simple and Dirty Calculator in Golang !" +
		"Pour consulter l'aide, entrez \"aide\" ou \"help\" puis appuyez sur entrer.")
	fmt.Println("Welcome to SDCG - Simple and Dirty Calculator in Golang ! " +
		"To consult the help guide, please enter \"aide\" or \"help\" then press enter.")

	for true {
		fmt.Printf("Calculator > ")
		toDisplay, err := collectUserInput()
		IfError(err)
		if toDisplay == "exit" {
			break
		}
		toParse, err2 := lexer(toDisplay)
		IfError(err2)
		fmt.Println(toParse)
	}
}

func collectUserInput() (string, error) {
	// Collect the user input, verify that the input is not empty then return the user input.
	userInput := ""
	var err error = nil
	fmt.Scanln(&userInput, &err)
	IfError(err)
	if userInput == "" {
		err = errors.New("empty entry, nothing were entred")
	}
	return userInput, err
}

func IfError(err error) () {
	if err != nil {
		fmt.Println("An error occured. Error :", err)
	}
}