package handlers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"regexp"

	bdd "../bdd"
)

func Register(username string, email string, password string) int {

	var passwordHash string
	_, pseudo := bdd.GetUser(username)
	verifemail, _ := regexp.Compile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,}")
	majLetter, _ := regexp.Compile("[A-Z]")
	minLetter, _ := regexp.Compile("[a-z]")
	number, _ := regexp.Compile("[0-9]")
	if len(username) < 3 || len(username) > 15 {
		return 1
	}
	if pseudo[0] == username {
		return 5
	}
	if len(password) < 8 {
		return 2
	}
	if len(majLetter.FindAllStringSubmatchIndex(password, -1)) < 1 || len(minLetter.FindAllStringSubmatchIndex(password, -1)) < 2 || len(number.FindAllStringSubmatchIndex(password, -1)) < 2 {
		return 3
	}
	if !verifemail.MatchString(email) {
		return 4
	}
	if email == pseudo[1] {
		return 6
	}
	passwordHashBytes := md5.Sum([]byte(password))
	passwordHash = hex.EncodeToString(passwordHashBytes[:])
	bdd.MakeUser(username, email, passwordHash)
	return 0
}

func Login(w http.ResponseWriter, getPseudo string, getMdp string) int {

	err, _ := bdd.GetUser(getPseudo)
	_, bddMdp := bdd.GetUserHash(getPseudo)
	loginPassHashBytes := md5.Sum([]byte(getMdp))
	loginPassHash := hex.EncodeToString(loginPassHashBytes[:])
	if err != 0 {
		if err == 500 {
			return 3
		}
	}
	if bddMdp == loginPassHash {
		keyBytes := make([]byte, 16)
		_, err := rand.Read(keyBytes)
		if err != nil {
			return 1
		}
		return 0
	}
	return 2
}
