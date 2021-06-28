package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func MakeUser(pseudo string, email string, hash string) int {
	_, db := OpenDataBase()                                                                //ouverture de la bdd
	result, err := db.Prepare("INSERT INTO User (Pseudo, Email, HashMDP) VALUES(?, ?, ?)") //Preparation de la commande
	if err != nil {
		fmt.Println("Error MakeUser") //cas d'erreur
		return 500
	} else {
		result.Exec(pseudo, email, hash) //creation du l'utilisateur
	}
	db.Close()
	return 0
}

func GetUser(username string) (int, [2]string) {

	_, db := OpenDataBase()

	var resultFunc [2]string
	result, err := db.Query("SELECT Pseudo, Email FROM User WHERE Pseudo = ?", username) // recherche de l'utilisateur avec le pseudo voulue
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() { //recuperation des info sur la bdd
		result.Scan(&Pseudo, &Email)
	}
	
	result.Close()
	resultFunc[0] = Pseudo
	resultFunc[1] = Email

	db.Close()
	return 0, resultFunc //renvoie l'utilisateur
}

func GetAllUser() (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT Pseudo, Email FROM User")
	if err != nil {
		fmt.Println("Error GetAllUser")
		fmt.Println(err.Error())

		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() {
		result.Scan(&Pseudo, &Email)
		temp := []string{Pseudo, Email}
		resultFunc = append(resultFunc, temp)
	}

	db.Close()
	return 0, resultFunc
}

func GetUserHash(username string) (int, string) {

	_, db := OpenDataBase()
	result, err := db.Query("SELECT HashMDP FROM User WHERE Pseudo = ?", username) //renvoie du hash de l'utilisateur voulue
	if err != nil {
		fmt.Println("Error GetUserHash")
		db.Close()
		return 500, "error 500"
	} else {

		var HashMDP string

		result.Next() //lecture de la bdd
		result.Scan(&HashMDP)
		result.Close()

		db.Close()
		return 0, HashMDP //renvoie le hash
	}
}
