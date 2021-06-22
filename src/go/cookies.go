package handlers


import (
	"time"
	guuid "github.com/google/uuid"
	"net/http"
)

func sessionCookie(w http.ResponseWriter, req *http.Request) {
	
	sessionToken := guuid.New().String()
	
	// Création du cookie 
	http.SetCookie(w, &http.Cookie{ 
		Name:    "session_Token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second), 
		Path : "/",
	})
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

func VerifyCookie(w http.ResponseWriter, req *http.Request) bool {
    
	_, err := req.Cookie("session_Token")
    
	if err != nil {
        return false
    } else {
        return true
    }
}
