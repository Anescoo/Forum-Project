package main

import (
	"fmt"
	"net/http"

	handlers "./src/go"
)

func main() {

	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Accueil)
	// http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/timeline", handlers.Timeline)
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/inscription", handlers.Inscription)
	http.HandleFunc("/pop", handlers.Pop)
	http.HandleFunc("/poptest", handlers.Poptest)

	http.ListenAndServe(":8000", nil)
}
