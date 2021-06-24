package main

import (
	"fmt"
	"net/http"

	handlers "./src/go"
)

// var ListeUserConnected []string
// var ListeUserKey []string

// func ListImport(UserConnected string, userKey string) {

//     for i := 0; i < len(ListeUserConnected); i++ {
// 		if ListeUserConnected[i] == UserConnected {
// 			ListeUserKey[i] = UserConnected
// 			return
// 		}
// 	}
// 	ListeUserConnected = append(ListeUserConnected, UserConnected)
// 	ListeUserKey = append(ListeUserKey, userKey)
// }

func main() {

	fmt.Print("Démarrage du serveur... 💬\n")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Création des pages de notre site
	http.HandleFunc("/home", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/inscription", handlers.Inscription)
	http.HandleFunc("/userpost", handlers.UserPost)
	http.HandleFunc("/userlikes", handlers.UserLikes)
	http.HandleFunc("/comms", handlers.Comments)
	http.HandleFunc("/pop", handlers.Pop)

	// Choix du port que l'on sélectionne
	http.ListenAndServe(":8000", nil)
}
