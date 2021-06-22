package authent


import (
	"time"
	guuid "github.com/google/uuid"
	"net/http"
)

// type Cookie struct {
// 	Password string 
// 	UserName string
// }

func sessionCookie(sessionToken string, w http.ResponseWriter) {
	
	sessionToken = guuid.New().String()
	// startSession := time.Now()

	http.SetCookie(w, &http.Cookie{ // Création du cookie 
		Name:    "session_Token",
		Value:   sessionToken,
		Expires: time.Now().Add(120 * time.Second), // Temps d'expiration de la session 
	})

	// return sessionToken
}


func readCookie(w http.ResponseWriter, r *http.Request) string {
	
	cookieContent, err := r.Cookie("session_Token")
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

