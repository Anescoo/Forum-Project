<<<<<<< HEAD
package bdd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "src/bdd/DataBase.db") //ouverture de la bdd
	if err != nil {                                      	//cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	return 0, db //renvoie la base de donnée
}

func MakeUser(pseudo string, email string, hash string) int {
	_, db := OpenDataBase()                                                                //ouverture de la bdd
	result, err := db.Prepare("INSERT INTO User (Pseudo, Email, HashMDP) VALUES(?, ?, ?)") //Preparation de la commande
	if err != nil {                                                                        //cas d'erreur
		return 500
	} else {
		result.Exec(pseudo, email, hash) //creation du l'utilisateur
	}
	return 0
}

func GetUser(username string) (int, [2]string) {

	_, db := OpenDataBase()

	var resultFunc [2]string
	result, err := db.Query("SELECT Pseudo, Email FROM User WHERE Pseudo = ?", username) // recherche de l'utilisateur avec le pseudo voulue
	if err != nil {
		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() { //recuperation des info sur la bdd
		result.Scan(&Pseudo, &Email)
	}

	resultFunc[0] = Pseudo
	resultFunc[1] = Email

	return 0, resultFunc //renvoie l'utilisateur
}

func GetUserHash(username string) (int, string) {

	_, db := OpenDataBase()
	result, err := db.Query("SELECT HashMDP FROM User WHERE Pseudo = ?", username) //renvoie du hash de l'utilisateur voulue
	if err != nil {
		fmt.Println("Error GetUserHash")
		return 500, "error 500"
	} else {

		var HashMDP string

		result.Next() //lecture de la bdd
		result.Scan(&HashMDP)
		return 0, HashMDP //renvoie le hash
	}
}

func MakePoste(user string, contenue string, categorie string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) Values(?,?,?,?)")
	if err != nil {
		return 500, 0
	} else {
		result.Exec(user, contenue, categorie, 0)
	}
	return 0, 0
}

func GetPosteByID(id int) (int, []string) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT PseudoUser, Contenue, Categorie, nbLike FROM Poste WHERE id = ?", id)
	if err != nil {
		return 500, nil
	} else {
		var resultFunc []string

		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int

		result.Scan(&PseudoUser, &Contenue, &Categorie, &nbLike)

		resultFunc[0] = PseudoUser
		resultFunc[1] = Contenue
		resultFunc[3] = Categorie
		resultFunc[4] = strconv.Itoa(nbLike)

		return 0, resultFunc
	}
}

func GetPosteByCategorie(categorie string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT PseudoUser, Contenue, nbLike FROME Poste Where Categorie = ?", categorie)
	if err != nil {
		return 500, resultFunc
	} else {

		var PseudoUser string
		var Contenue string
		var nbLike int
		temp := 0

		for result.Next() {
			result.Scan(&PseudoUser, &Contenue, &nbLike)
			resultFunc[temp][0] = PseudoUser
			resultFunc[temp][1] = Contenue
			resultFunc[temp][2] = strconv.Itoa(nbLike)
			temp++
		}

		return 0, resultFunc
	}
}

func MakeCategorie(name string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		return 500, 0
	} else {
		result.Exec(name)
	}
	return 0, 0
}
=======
package bdd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "../bdd/DataBase.db") //ouverture de la bdd
	if err != nil {                                      //cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	return 0, db //renvoie la base de donner
}

func MakeUser(pseudo string, email string, hash string) int {
	_, db := OpenDataBase()                                                                //ouverture de la bdd
	result, err := db.Prepare("INSERT INTO User (Pseudo, Email, HashMDP) VALUES(?, ?, ?)") //Preparation de la commande
	if err != nil {                                                                        //cas d'erreur
		return 500
	} else {
		result.Exec(pseudo, email, hash) //creation du l'utilisateur
	}
	return 0
}

func GetUser(username string) (int, [2]string) {

	_, db := OpenDataBase()

	var resultFunc [2]string
	result, err := db.Query("SELECT Pseudo, Email FROM User WHERE Pseudo = ?", username) // recherche de l'utilisateur avec le pseudo voulue
	if err != nil {
		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() { //recuperation des info sur la bdd
		result.Scan(&Pseudo, &Email)
	}

	resultFunc[0] = Pseudo
	resultFunc[1] = Email

	return 0, resultFunc //renvoie l'utilisateur
}

func GetUserHash(username string) (int, string) {

	_, db := OpenDataBase()
	result, err := db.Query("SELECT HashMDP FROM User WHERE Pseudo = ?", username) //renvoie du hash de l'utilisateur voulue
	if err != nil {
		fmt.Println("Error GetUserHash")
		return 500, "error 500"
	} else {

		var HashMDP string

		result.Next() //lecture de la bdd
		result.Scan(&HashMDP)
		return 0, HashMDP //renvoie le hash
	}
}

func MakePoste(user string, contenue string, categorie string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) Values(?,?,?,?)")
	if err != nil {
		return 500, 0
	} else {
		result.Exec(user, contenue, categorie, 0)
	}
	return 0, 0
}

func GetPosteByID(id int) (int, []string) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT PseudoUser, Contenue, Categorie, nbLike FROM Poste WHERE id = ?", id)
	if err != nil {
		return 500, nil
	} else {
		var resultFunc []string

		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int

		result.Scan(&PseudoUser, &Contenue, &Categorie, &nbLike)

		resultFunc[0] = PseudoUser
		resultFunc[1] = Contenue
		resultFunc[3] = Categorie
		resultFunc[4] = strconv.Itoa(nbLike)

		return 0, resultFunc
	}
}

func GetPosteByCategorie(categorie string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT PseudoUser, Contenue, nbLike FROME Poste Where Categorie = ?", categorie)
	if err != nil {
		return 500, resultFunc
	} else {

		var PseudoUser string
		var Contenue string
		var nbLike int
		temp := 0

		for result.Next() {
			result.Scan(&PseudoUser, &Contenue, &nbLike)
			resultFunc[temp][0] = PseudoUser
			resultFunc[temp][1] = Contenue
			resultFunc[temp][2] = strconv.Itoa(nbLike)
			temp++
		}

		return 0, resultFunc
	}
}

func MakeCategorie(name string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		return 500, 0
	} else {
		result.Exec(name)
	}
	return 0, 0
}
>>>>>>> 54a15459cdb70964db55023c3d0e1fce5387cfdb
