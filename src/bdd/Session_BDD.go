package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddSession(PseudoUser string, UUID string) int {
	_, db := OpenDataBase()
	resultFunc, err := db.Prepare("INSERT INTO Session (PseudoUser, UUID) VALUES(?,?)")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		_, err2 := resultFunc.Exec(PseudoUser, UUID)
		if err2 != nil {
			fmt.Println(err2.Error())

		}
	}
	
	db.Close()

	return 0
}

func GetUserByUUID(UUID string) (int, string) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT PseudoUser FROM Session WHERE UUID = ?", UUID)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ""
	} else {
		var PseudoUser string
		result.Next()
		result.Scan(&PseudoUser)
		result.Close()
		db.Close()
		return 0, PseudoUser
	}
}

func DeleteSession(PseudoUser string) int {
	_, db := OpenDataBase()
	resultFunc, err := db.Prepare("DELETE FROM Session WHERE PseudoUser = ?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		resultFunc.Exec(PseudoUser)
	}
	db.Close()
	return 0
}

func SessionAlreadyExiste(PseudoUser string) (int, bool) {
	_, db := OpenDataBase()
	resultFunc, err := db.Prepare("SELECT PseudoUser From Session WHERE PseudoUser = ?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, false
	} else {
		var Pseudo string
		temp, err := resultFunc.Query(PseudoUser)
		if err != nil {
			fmt.Println(err.Error())
			db.Close()
			return 500, false
		} else {
			temp.Next()
			temp.Scan(&Pseudo)
			db.Close()
			if Pseudo != PseudoUser {
				return 0, false
			} else {
				return 0, true
			}
		}
	}
}
