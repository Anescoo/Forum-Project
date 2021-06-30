package main

import (
	"fmt"
	"net/http"

	handlers "./src/go"
)

func main() {

	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Création des actions de notre site
	http.HandleFunc("/home", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/inscription", handlers.Inscription)
	http.HandleFunc("/userpost", handlers.UserPost)
	http.HandleFunc("/userlikes", handlers.UserLikes)
	http.HandleFunc("/deconnexion", handlers.Deconnexion)

	// Choix du port que l'on sélectionne
	http.ListenAndServe(":8000", nil)
}
