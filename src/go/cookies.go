package handlers


import (
	"fmt"
	"time"
	guuid "github.com/google/uuid"
	"net/http"
)

func sessionCookie(w http.ResponseWriter, req *http.Request) {
	
	sessionToken := guuid.New().String()
	
	// Création du cookie 
	cookie := http.Cookie{ 
		Name:    "session_Token",
		Value:   sessionToken,
		Expires: time.Now().Add(2 * time.Hour), 
		Path : "/",
	}
	http.SetCookie(w, &cookie)

	// return sessionToken
	fmt.Println("Value Cookies : ", sessionToken)
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

func verifyCookie(w http.ResponseWriter, req *http.Request) bool {
    
	_, err := req.Cookie("session_Token")
    
	if err != nil {
        return false
    } else {
        return true
    }
}


func deleteCookie(w http.ResponseWriter, req *http.Request){

	c := http.Cookie {
		Name : "session_Token",
		MaxAge : -1,
	}
	http.SetCookie(w, &c)

}