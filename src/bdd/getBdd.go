package bdd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDataBase() (int, *sql.DB) {
	db, err := sql.Open("sqlite3", "src/bdd/DataBase.db") //ouverture de la bdd
	if err != nil {                                       //cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	if err != nil { //cas d'erreur
		fmt.Println("Error database")
		return 500, nil //renvoie du code d'erreur
	}
	return 0, db //renvoie la base de donner
}
