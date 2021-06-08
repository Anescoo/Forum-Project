package authent

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"regexp"
	"time"

	bdd "../../bdd"
)

func Register(username string, email string, password string) int {
	var passwordHash string
	verifemail, _ := regexp.Compile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+.[A-Za-z]{2,}")
	majLetter, _ := regexp.Compile("[A-Z]")
	minLetter, _ := regexp.Compile("[a-z]")
	number, _ := regexp.Compile("[0-9]")
	if len(username) < 3 || len(username) > 15 {
		return 1
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
	passwordHashBytes := md5.Sum([]byte(password))
	passwordHash = hex.EncodeToString(passwordHashBytes[:])
	bdd.MakeUser(username, email, passwordHash)
	return 0
}

<<<<<<< HEAD
func Login(w http.ResponseWriter, getPseudo string, getMdp string) {
=======
func Login(w http.ResponseWriter, getPseudo string, getMdp string) int {
>>>>>>> 209af0bdbf7c27bbe1c51105fd3fd7a7387c6aac
	err, _ := bdd.GetUser(getPseudo)
	_, bddMdp := bdd.GetUserHash(getPseudo)
	var key string
	loginPassHashBytes := md5.Sum([]byte(getMdp))
	loginPassHash := hex.EncodeToString(loginPassHashBytes[:])
	if err != 0 {

	}
	if bddMdp == loginPassHash {
		keyBytes := make([]byte, 16)
		_, err := rand.Read(keyBytes)
		if err != nil {
			return 1
		}
		key = hex.EncodeToString(keyBytes)
		expiration := time.Now().Add(6 * time.Hour)
		cookieForKey := http.Cookie{Name: "sessionKey", Value: key, Expires: expiration}
		cookieForName := http.Cookie{Name: "sessionOwner", Value: getPseudo, Expires: expiration}
		http.SetCookie(w, &cookieForKey)
		http.SetCookie(w, &cookieForName)
		return 0
	}
	return 2
}
