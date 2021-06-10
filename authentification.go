package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

func register() {
	var username string
	var password string
	var email string
	var retry bool = false
	var retryLog string
	var passwordHash string
	verifemail, _ := regexp.Compile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,}")
	majLetter, _ := regexp.Compile("[A-Z]")
	minLetter, _ := regexp.Compile("[a-z]")
	number, _ := regexp.Compile("[0-9]")
	for i := 0; i < 1; i++ { //retirer les for
		fmt.Println("entrez un username")
		fmt.Scanln(&username)
		if len(username) < 3 || len(username) > 15 {
			i = -1
			//retry = true
			//retryLog = retryLog + "pseudo trop court,  "
		}

	}
	for e := 0; e < 1; e++ { //retirer les for
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
	for u := 0; u < 1; u++ { //retirer les for
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
		//retourner une érreur ici
	} else {

		passwordHashBytes := md5.Sum([]byte(password))
		passwordHash = hex.EncodeToString(passwordHashBytes[:])
		fmt.Println(username, "  ", password, "  ", email)
		fmt.Println(username, "  ", email)
		fmt.Print("password : ")
		fmt.Printf("%x", passwordHash)
		fmt.Println()
	}

}

func login(w http.ResponseWriter, userLogin string, loginPass string) {
	//recevoir les info
	var userList = []string{}
	var key string

	for i := 0; i < len(userList); i++ {
		fmt.Print(i)
		if userLogin == userList[i].username {
			fmt.Print(" User is valid")
			loginPassHash := md5.Sum([]byte(loginPass))
			if loginPassHash == userList[i].passwordHash {
				keyBytes := make([]byte, 16)
				_, err := rand.Read(keyBytes)
				if err != nil {
					// return a error
				}
				fmt.Print(",  password is right")
				fmt.Println()

				key = hex.EncodeToString(keyBytes)
				fmt.Print("clef id unique : ", key)
				expiration := time.Now().Add(6 * time.Hour)
				cookie := http.Cookie{Name: "sessionKey", Value: key, Expires: expiration}
				http.SetCookie(w, &cookie)

			} else {
				fmt.Print(",  password is wrong")
			}
		}
		fmt.Println()
	}

}
