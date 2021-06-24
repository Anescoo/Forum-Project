package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Like(idPost int, PseudoUser string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("INSERT INTO Like (PseudoUser, idPoste) VALUES(?,?)")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(PseudoUser, idPost)
		db.Close()
		return 0
	}
}

func Unlike(idPost int, PseudoUser string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("DELET FROM Like WHERE idPoste = ? AND PseudoUser =?")
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		result.Exec(idPost, PseudoUser)
		db.Close()
		return 0
	}
}

func GetPosteLikeByUser(UserPseuso string) (int, [][]string) {
	_, db := OpenDataBase()
	var ResultFunc [][]string
	var IdLiker []int
	temp, err := db.Query("SELECT idPoste FROM Like WHERE PseudoUser = ?", UserPseuso)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500, ResultFunc
	} else {
		var IDPoste int

		for temp.Next() {
			temp.Scan(&IDPoste)
			IdLiker = append(IdLiker, IDPoste)
		}

		db.Close()
		for _, i := range IdLiker {
			_, temp := GetPosteByID(IdLiker[i])
			ResultFunc = append(ResultFunc, temp)
		}
		return 0, ResultFunc
	}
}

func IsLike(idPoste int, PseudoUser string) (int, bool) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT isLike FROM Like WHERE PseudoUser = ? AND IdPoste = ?", PseudoUser, idPoste)
	if err != nil {
		db.Close()
		return 500, false
	} else {
		var isLike bool

		result.Next()
		result.Scan(&isLike)
		db.Close()

		if isLike == true {
			return 0, true
		} else {
			return 0, false
		}
	}
}

func GetLikeNb(id int) int {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT idPoste FROM Like WHERE idPoste = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		var resultFunc []int
		var idPoste int

		for result.Next() {
			result.Scan(&idPoste)
			resultFunc = append(resultFunc, idPoste)
		}

		db.Close()
		fmt.Println(len(resultFunc))
		return len(resultFunc)
	}
}
