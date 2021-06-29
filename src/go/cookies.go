package handlers

import (
	"fmt"
<<<<<<< HEAD
	"net/http"
=======
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
	"time"

	guuid "github.com/google/uuid"
)

func sessionCookie(w http.ResponseWriter, req *http.Request) string {
<<<<<<< HEAD

	sessionToken := guuid.New().String()

	// Création du cookie
	cookie := http.Cookie{
		Name:    "session_Token",
		Value:   sessionToken,
		Expires: time.Now().Add(2 * time.Hour),
		Path:    "/",
=======
	
	sessionToken := guuid.New().String()
	
	// Création du cookie 
	cookie := http.Cookie{ 
		Name:    "session_Token",
		Value:   sessionToken,
		Expires: time.Now().Add(2 * time.Hour), 
		Path : "/",
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
	}
	http.SetCookie(w, &cookie)

	uuidValue := cookie.Value
	fmt.Println("Value Cookies : ", uuidValue)
	return uuidValue
}

func readCookie(w http.ResponseWriter, req *http.Request) string {

	cookieContent, err := req.Cookie("session_Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return ""
		}
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	sessionToken := cookieContent.Value
	return sessionToken
}

func VerifyCookie(req *http.Request) bool {
<<<<<<< HEAD

=======
    
>>>>>>> 14fa70e34b9264b2e8c8028503e58f8f575ad285
	_, err := req.Cookie("session_Token")

	if err != nil {
		return false
	} else {
		return true
	}
}

func deleteCookie(w http.ResponseWriter, req *http.Request) {

	c := http.Cookie{
		Name:   "session_Token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
}

// fonction de verification connexion user

func VerifyUser(w http.ResponseWriter, req *http.Request) bool {

	uuidValue := sessionCookie(w, req)

	// if uuidValue == {
	//     return true
	// }else {
	//     return false
	// }
	fmt.Println(uuidValue)
	return true
}

func deleteCookie(w http.ResponseWriter, req *http.Request){

	c := http.Cookie {
		Name : "session_Token",
		MaxAge : -1,
	}
	http.SetCookie(w, &c)
}

// fonction de verification connexion user 

func VerifyUser(w http.ResponseWriter, req *http.Request) bool {

	
	uuidValue := sessionCookie(w, req)

	// if uuidValue == {
	// 	return true
	// }else {
	// 	return false
	// }
	fmt.Println(uuidValue)
	return true
}
