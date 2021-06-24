package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func MakeComment(Username string, Contenue string, IDPoste int) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Commentaire (User, Contenue, IDPoste) VALUES(?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(Username, Contenue, IDPoste)
		db.Close()
		return 0
	}
}

func GetCommentByPoste(idPoste int) (int, [][]string) {
	_, db := OpenDataBase()
	var ResultFunc [][]string
	result, err := db.Query("SELECT User, Contenue FROM Commentaire WHERE IDPoste = ?", idPoste)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var User string
		var Contenue string

		for result.Next() {
			result.Scan(&User, &Contenue)
			temp := []string{User, Contenue}
			ResultFunc = append(ResultFunc, temp)
		}
		db.Close()
		return 0, ResultFunc
	}
}
