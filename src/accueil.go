package handlers 

import (
	"fmt"
	"net/http"
)

func Accueil(w http.ResponseWriter, req *http.Request){

	fmt.Print("Page d'accueil ✔️ \n")


	if req.URL.Path == "/" {	//verification de l'URL
	} else if req.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}