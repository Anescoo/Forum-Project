<<<<<<< HEAD
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
=======
package bdd

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "src/bdd/DataBase.db") //ouverture de la bdd
	if err != nil {                                       //cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	return 0, db //renvoie la base de donner
}

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
		return 500, resultFunc
	}

	var Pseudo string
	var Email string

	for result.Next() {
		result.Scan(&Pseudo, &Email)
		temp := []string{Pseudo, Email}
		resultFunc = append(resultFunc, temp)
	}

	return 0, resultFunc
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

		db.Close()
		return 0, HashMDP //renvoie le hash
	}
}

func MakePoste(user string, contenue string, categorie string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println("Error MakePoste")
		return 500
	} else {
		fmt.Println(user, contenue, categorie)
		result.Exec(user, contenue, categorie, 0)

		fmt.Println("MakePoste done")
	}

	db.Close()
	return 0
}

func GetPosteByID(id int) (int, []string) {
	_, db := OpenDataBase()
	var resultFunc []string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike FROM Poste WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error GetPosteById")
		return 500, resultFunc
	} else {
		fmt.Println("Starting GetPosteById")

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int

		result.Next()
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike)

		resultFunc = append(resultFunc, strconv.Itoa(ID))
		resultFunc = append(resultFunc, PseudoUser)
		resultFunc = append(resultFunc, Contenue)
		resultFunc = append(resultFunc, Categorie)
		resultFunc = append(resultFunc, strconv.Itoa(nbLike))

		db.Close()
		return 0, resultFunc
	}
}

func GetAllPoste() (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike FROM Poste")
	if err != nil {
		fmt.Println("Error GetAllPoste")
		return 500, resultFunc
	}

	var ID int
	var PseudoUser string
	var Contenue string
	var Categorie string
	var nbLike int

	for result.Next() {
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike)
		temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike)}
		resultFunc = append(resultFunc, temp)
	}

	return 0, resultFunc
}

func GetPosteByCategorie(categorie string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, nbLike FROM Poste WHERE Categorie = ?", categorie)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var nbLike int

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &nbLike)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, strconv.Itoa(nbLike)}
			resultFunc = append(resultFunc, temp)
		}

		db.Close()
		return 0, resultFunc
	}
}

func GetPosteByUser(UserPseudo string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike FROM Poste WHERE PseudoUser = ?", UserPseudo)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike)}
			resultFunc = append(resultFunc, temp)
		}

		db.Close()
		return 0, resultFunc
	}
}

func DeletePoste(id int) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELETE FROM Poste WHERE ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		return 500
	} else {
		result.Exec(id)
		return 0
	}
}

func MakeCategorie(name string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		return 500
	} else {
		result.Exec(name)
	}

	db.Close()
	return 0
}
>>>>>>> 910468e1ddfd93cc3295e08774f8ed21a05d8af2
