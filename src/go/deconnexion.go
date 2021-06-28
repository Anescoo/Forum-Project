package handlers

import (
	"net/http"
)

func Deconnexion(w http.ResponseWriter, req *http.Request){

	deleteCookie(w, req)

}