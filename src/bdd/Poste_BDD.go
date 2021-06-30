package bdd

import (
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func MakePoste(user string, contenue string, categorie string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Poste (PseudoUser, Contenue, Categorie, nbLike) VALUES(?,?,?,?)")
	if err != nil {
		fmt.Println("Error MakePoste")
		fmt.Println(err.Error())
		return 500
	} else {
		_, err := result.Exec(user, contenue, categorie, 0)
		if err != nil {
			fmt.Println(err.Error())
			db.Close()
			return 500
		} else {
			db.Close()
			return 0
		}
	}
}

func GetPosteByID(id int) (int, []string) {
	_, db := OpenDataBase()
	var resultFunc []string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error GetPosteById")
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {
		

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var DatePoste string
		var nbLike int

		result.Next()
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
		result.Close()

		resultFunc = append(resultFunc, strconv.Itoa(ID))
		resultFunc = append(resultFunc, PseudoUser)
		resultFunc = append(resultFunc, Contenue)
		resultFunc = append(resultFunc, Categorie)
		resultFunc = append(resultFunc, strconv.Itoa(nbLike))
		resultFunc = append(resultFunc, DatePoste)

		db.Close()
		return 0, resultFunc
	}
}

func GetAllPoste() (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste ORDER BY DatePoste DESC LIMIT 5")
	if err != nil {
		fmt.Println("Error GetAllPoste")
		fmt.Println(err.Error())
		return 500, resultFunc
	}

	var ID int
	var PseudoUser string
	var Contenue string
	var Categorie string
	var DatePoste string
	var nbLike int

	for result.Next() {
		result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
		temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike), DatePoste}
		resultFunc = append(resultFunc, temp)
	}
	result.Close()
	db.Close()
	return 0, resultFunc
}

func GetPosteByCategorie(categorie string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, nbLike, DatePoste FROM Poste WHERE Categorie = ?", categorie)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var nbLike int
		var DatePoste string

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &nbLike, &DatePoste)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, strconv.Itoa(nbLike), DatePoste}
			resultFunc = append(resultFunc, temp)
		}
		result.Close()

		db.Close()
		return 0, resultFunc
	}
}

func GetPosteByUser(UserPseudo string) (int, [][]string) {
	_, db := OpenDataBase()
	var resultFunc [][]string
	result, err := db.Query("SELECT ID, PseudoUser, Contenue, Categorie, nbLike, DatePoste FROM Poste WHERE PseudoUser = ? ORDER BY DatePoste DESC", UserPseudo)
	if err != nil {
		fmt.Println(err.Error())
		return 500, resultFunc
	} else {

		var ID int
		var PseudoUser string
		var Contenue string
		var Categorie string
		var nbLike int
		var DatePoste string

		for result.Next() {
			result.Scan(&ID, &PseudoUser, &Contenue, &Categorie, &nbLike, &DatePoste)
			temp := []string{strconv.Itoa(ID), PseudoUser, Contenue, Categorie, strconv.Itoa(nbLike), DatePoste}
			resultFunc = append(resultFunc, temp)
		}
		result.Close()

		db.Close()
		return 0, resultFunc
	}
}

func UpdatePoste(id int, update string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("UPDATE Poste SET Contenue = ? WHERE ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(update, id)
		db.Close()
		return 0
	}
}

func DeletePoste(id int) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELETE FROM Poste WHERE ID = ?")
	if err != nil {
		fmt.Println(err.Error())
		return 500
	} else {
		_, err := result.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
			return 500
		} else {
			db.Close()
			return 0
		}
	}
}
