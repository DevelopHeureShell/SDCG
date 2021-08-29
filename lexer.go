package main

import "fmt"

type Token struct {
	Nature string
	Value string
}

func lexer(toAnalyse string) []Token {
	toReturn := make([]Token, 0) // Le slice qui va être renvoyée, contenant les tokens obtenus grâce au lexer.
	toMakeANumber := make([]string, 0) // Le slice qui va permettre à la fonction makeNumber de créer un nombre.

	for i, charAnalysedShadowed := range toAnalyse {

		charAnalysed := string(charAnalysedShadowed) // Le caractère actuellement analysé par le lexer.
		isCharacterANumber := false // Le booléen qui stocke si le caractère est un nombre ou non.
		isNextCharacterANumber := false // Le booléen qui stocke si le prochain caractère est un nombre ou non.
		var currentToken Token // L'objet de type Token qui permet de créer le token avant de l'ajouter au tableau.

		for _, aNumber := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			// Détecte si le caractère actuel est un nombre.
			if charAnalysed == aNumber {
				isCharacterANumber = true
				break
			}
		}
		for _, aNumber := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
			// Détecte si le prochain caractère est un nombre.
			if i < len(toAnalyse)-1 {
				if string(toAnalyse[i+1]) == aNumber {
					isNextCharacterANumber = true
					break
				}
			}
		}
		switch {
		case isCharacterANumber == true:
			// Si le caractère est un nombre, on l'ajoute à toMakeANumber pour en faire un nombre ensuite.
			toMakeANumber = append(toMakeANumber, charAnalysed)
			if isNextCharacterANumber == false {
				numberToAdd := ""
				for _, numberUnit := range toMakeANumber {
					numberToAdd = numberToAdd + numberUnit
				}
				currentToken.Nature = "NUMBER"
				currentToken.Value = numberToAdd
				toReturn = append(toReturn, currentToken)
				toMakeANumber = nil
			}
			break
		case isCharacterANumber == false:
			// Si le caractère n'est pas un nombre, on construit le nombre précédemment obtenu, avec la fonction makeNumber.
			switch {
			// On étudie ensuite en quoi consiste ce caractère puis on l'ajoute ou on affiche un message d'erreur.
			case charAnalysed == "+":
				currentToken.Nature = "OPERATOR"
				currentToken.Value = "+"
				toReturn = append(toReturn, currentToken)
				break
			case charAnalysed == "-":
				currentToken.Nature = "OPERATOR"
				currentToken.Value = "-"
				toReturn = append(toReturn, currentToken)
				break
			case charAnalysed == "X":
				currentToken.Nature = "OPERATOR"
				currentToken.Value = "*"
				toReturn = append(toReturn, currentToken)
				fmt.Printf("Warning : The character %d is an invalid symbol to multiply. Please consider to use \"*\""+
					" in the future.", i+1)
				break
			case charAnalysed == "*":
				currentToken.Nature = "OPERATOR"
				currentToken.Value = "*"
				toReturn = append(toReturn, currentToken)
				break
			case charAnalysed == "/":
				currentToken.Nature = "OPERATOR"
				currentToken.Value = "/"
				toReturn = append(toReturn, currentToken)
				break
			case charAnalysed == "(":
				currentToken.Nature = "BRACKET"
				currentToken.Value = "("
				toReturn = append(toReturn, currentToken)
				break
			case charAnalysed == ")":
				currentToken.Nature = "BRACKET"
				currentToken.Value = ")"
				toReturn = append(toReturn, currentToken)
				break
			default:
				fmt.Printf("Warning : The character %d is invalid, omitting.", i+1)
				break
			}
		}
		i++
	}

	for i, currentToken := range toReturn {
		if currentToken.Value == "" {
			toReturn = append(toReturn[:i], toReturn[i+1])
		}
	}
	return toReturn
}