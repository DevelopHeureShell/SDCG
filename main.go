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
		if err != nil {
			fmt.Println("An error occured. Error :", err)
			continue
		}
		if toDisplay == "exit" {
			break
		}
		var toParse []Token = lexer(toDisplay)
		fmt.Println(toParse)
	}
}

func collectUserInput() (string, error) {
	// Collect the user input, verify that the input is not empty then return the user input.
	userInput := ""
	var err error = nil
	fmt.Scanln(&userInput, &err)
	if err != nil {
		return "ERROR", err
	}
	if err == nil && userInput == "" {
		err = errors.New("empty entry, nothing were entred")
	}
	return userInput, err
}