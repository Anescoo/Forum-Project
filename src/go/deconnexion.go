package handlers

import (
	"net/http"
	bdd "../bdd"
)

func Deconnexion(w http.ResponseWriter, req *http.Request){

	deleteCookie(w, req)
	uuidValue := sessionCookie(w, req)
	_, userValue := bdd.GetUserByUUID(uuidValue)
	bdd.DeleteSession(userValue)
	http.Redirect(w, req, "/home", http.StatusSeeOther)
	
}