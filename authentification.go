package main

import (
	"fmt"
	"regexp"
)

func main() {
	var username string
	var password string
	var email string
	verifemail, _ := regexp.Compile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,}")
	majLetter, _ := regexp.Compile("[A-Z]")
	minLetter, _ := regexp.Compile("[a-z]")
	number, _ := regexp.Compile("[0-9]")
	for i := 0; i < 1; i++ { //retiré les for
		fmt.Println("entrez un username")
		fmt.Scanln(&username)
		if len(username) < 3 || len(username) > 15 {
			i = -1 //retourné une érreur ici
		}

	}
	for e := 0; e < 1; e++ { //retiré les for
		fmt.Println("entrez un mot de passe")
		fmt.Scanln(&password)
		if len(password) < 8 {
			e = -1 //retourné une érreur ici
		}
		if len(majLetter.FindAllStringSubmatchIndex(password, -1)) < 1 || len(minLetter.FindAllStringSubmatchIndex(password, -1)) < 2 || len(number.FindAllStringSubmatchIndex(password, -1)) < 2 {
			e = -1 //retourné une érreur ici
		}
	}
	for u := 0; u < 1; u++ { //retiré les for
		fmt.Println("entrez une email")
		fmt.Scanln(&email)
		if !verifemail.MatchString(email) {
			u = -1 //retourné une érreur ici
		}
	}
	fmt.Println(username, "  ", password, "  ", email)
}
