package bdd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() *sql.DB {
	db, err := sql.Open("sqlite3", "../bdd/DataBase.db")
	if err != nil {
		fmt.Println("Error database")
	}
	return db
}

func MakeUser(pseudo string, email string, hash string) int {
	db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO User (Pseudo, Email, HashMDP) VALUES(?, ?, ?)")
	if err != nil {

		return 500
	} else {
		result.Exec(pseudo, email, hash)
	}
	return 0
}

func GetUser(username string) [2]string {

	db := OpenDataBase()

	var temp [2]string
	result, err := db.Query("SELECT Pseudo, Email FROM User WHERE Pseudo = ?", username)
	if err != nil {
		var temp [2]string
		temp[0] = "500"
		return temp
	}

	var Pseudo string
	var Email string

	for result.Next() {
		result.Scan(&Pseudo, &Email)
	}

	temp[0] = Pseudo
	temp[1] = Email

	return temp
}

func GetUserHash(username string) string {

	db := OpenDataBase()
	result, err := db.Query("SELECT HashMDP FROM User WHERE Pseudo = ?", username)
	if err != nil {
		fmt.Println("Error GetUserHash")
		return "500"
	} else {

		var HashMDP string

		result.Next()
		result.Scan(&HashMDP)

		return HashMDP
	}
}

func MakePoste(user string, contenue string, categorie string, nblike int) int {
	db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) Values(?,?,?,?)")
	if err != nil {
		return 500
	} else {
		result.Exec(user, contenue, categorie, 0)
	}
	return 0
}

func GetPosteByID(id int) []string {
	db := OpenDataBase()
	result, err := db.Query("SELECT PseudoUser, Contenue, Categorie, nbLike FROM Poste WHERE id = ?", id)
	if err != nil {
		var temp []string
		temp[0] = "500"
		return temp
	} else {
		var temp []string

		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int

		result.Scan(&PseudoUser, &Contenue, &Categorie, &nbLike)

		temp[0] = PseudoUser
		temp[1] = Contenue
		temp[3] = Categorie
		temp[4] = strconv.Itoa(nbLike)

		return temp
	}
}

func MakeCategorie(name string) int {
	db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		return 500
	} else {
		result.Exec(name)
	}
	return 0
}

func GetPosteByCategorie(categorie string) [][]string {
	db := OpenDataBase()
	result, err := db.Query("SELECT PseudoUser, Contenue, nbLike FROME Poste Where Categorie = ?", categorie)
	if err != nil {
		var temp [][]string
		return temp
	} else {
		var temp [][]string

		var PseudoUser string
		var Contenue string
		var nbLike int

		for result.Next() {
			result.Scan(&PseudoUser, &Contenue, &nbLike)
		}

		return temp
	}
}
