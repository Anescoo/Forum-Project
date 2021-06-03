package handlers 

import (
	"fmt"
	"net/http"
)

func Connexion(w http.ResponseWriter, req *http.Request){

	fmt.Print("Page de connexion ✔️ \n")
}