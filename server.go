package main

import (
	"fmt"
	"net/http"

	handlers "./src/go"
)

// var ListeUserConnected []string
// var ListeUserKey []string

// func ListImport(UserConnected string, userKey string) {
// 	for i := 0; i < len(ListeUserConnected); i++ {
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

	http.HandleFunc("/", handlers.Accueil)
	http.HandleFunc("/connexion", handlers.Connexion)
	http.HandleFunc("/timeline", handlers.Timeline)
	http.HandleFunc("/profil", handlers.Profil)
	http.HandleFunc("/inscription", handlers.Inscription)

	http.ListenAndServe(":8000", nil)
}
