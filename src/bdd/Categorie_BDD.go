package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func MakeCategorie(name string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Categorie (Name) Values(?)")
	if err != nil {
		db.Close()
		return 500
	} else {
		result.Exec(name)
	}

	db.Close()
	return 0
}

func GetAllCategorie() (int, []string) {
	_, db := OpenDataBase()
	var ResultFunc []string
	result, err := db.Query("SELECT Name FROM Categorie")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var NameCategorie string

		for result.Next() {
			result.Scan(&NameCategorie)
			ResultFunc = append(ResultFunc, NameCategorie)
		}
		result.Close()

		db.Close()
		return 0, ResultFunc
	}
}

func NbPosteByCategorie(name string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT ID FROM Poste WHERE Categorie = ?", name)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 0, 500
	} else {
		var temp []int
		var ID int

		for result.Next() {
			result.Scan(&ID)
			temp = append(temp, ID)
		}
		result.Close()
		return 0, len(temp)
	}
}

func DeleteCategorire(name string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELETE FROM Categorie WHERE Name = ? ")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(name)
		db.Close()
		return 0
	}
}
