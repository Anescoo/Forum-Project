package handlers

import (
	"net/http"

	bdd "../bdd"
)

func Deconnexion(w http.ResponseWriter, req *http.Request) {

	deleteCookie(w, req)
	uuidValue := readCookie(w, req)
	errBdd, userValue := bdd.GetUserByUUID(uuidValue)
	if ReturnError500(w, errBdd) {
		return
	}
	errBdd = bdd.DeleteSession(userValue)
	if ReturnError500(w, errBdd) {
		return
	}
	http.Redirect(w, req, "/home", http.StatusSeeOther)

}
