package handlers

import (
	"net/http"

	bdd "../bdd"
)

func Deconnexion(w http.ResponseWriter, req *http.Request) {

	deleteCookie(w, req)
	uuidValue := readCookie(w, req)
	errBdd, userValue := bdd.GetUserByUUID(uuidValue)
	ReturnError500(w, errBdd)
	errBdd = bdd.DeleteSession(userValue)
	ReturnError500(w, errBdd)
	http.Redirect(w, req, "/home", http.StatusSeeOther)

}
