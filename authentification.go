package main

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"regexp"
)

func main() {
	var username string
	var password string
	var email string
	var retry bool = false
	var retryLog string
	verifemail, _ := regexp.Compile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,}")
	majLetter, _ := regexp.Compile("[A-Z]")
	minLetter, _ := regexp.Compile("[a-z]")
	number, _ := regexp.Compile("[0-9]")
	for i := 0; i < 1; i++ { //retiré les for
		fmt.Println("entrez un username")
		fmt.Scanln(&username)
		if len(username) < 3 || len(username) > 15 {
			i = -1
			//retry = true
			//retryLog = retryLog + "pseudo trop court,  "
		}

	}
	for e := 0; e < 1; e++ { //retiré les for
		fmt.Println("entrez un mot de passe")
		fmt.Scanln(&password)
		if len(password) < 8 {
			e = -1
			//retry = true
			//retryLog = retryLog + "mot de passe trop court,  "
		}
		if len(majLetter.FindAllStringSubmatchIndex(password, -1)) < 1 || len(minLetter.FindAllStringSubmatchIndex(password, -1)) < 2 || len(number.FindAllStringSubmatchIndex(password, -1)) < 2 {
			e = -1
			//retry = true
			//retryLog = retryLog + "mot de passe trop faible,  "
		}
	}
	for u := 0; u < 1; u++ { //retiré les for
		fmt.Println("entrez une email")
		fmt.Scanln(&email)
		if !verifemail.MatchString(email) {
			u = -1
			//retry = true
			//retryLog = retryLog + "email invalide,  "
		}
	}
	if retry {
		fmt.Println(retryLog)
		//retourné une érreur ici
	} else {
		fmt.Println(username, "  ", password, "  ", email)
		passwordHash := md5.Sum([]byte(password))
		key := make([]byte, 16)
		_, err := rand.Read(key)
		if err != nil {
			// return a error
		}
		fmt.Println(username, "  ", email)
		fmt.Print("password : ")
		fmt.Printf("%x", passwordHash)
		fmt.Println()
		fmt.Print("clef id unique : ")
		fmt.Printf("%x", key)
		fmt.Println()
	}

}
