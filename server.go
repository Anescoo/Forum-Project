package main 


import (
	"net/http"
	"fmt"
	handlers "./src"
)

func main(){

	fmt.Print("Démarrage du serveur... 💬\n")

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/timeline", handlers.Timeline)


	http.ListenAndServe(":8000", nil)
}