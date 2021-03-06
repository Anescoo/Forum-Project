package bdd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Like(idPost int, Pseudo string) int {
	_, db := OpenDataBase()
	temp, err := db.Query("SELECT PseudoUser FROM Like WHERE PseudoUser = ? AND idPoste = ?", Pseudo, idPost)
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return 500
	} else {
		var PseudoUser string
		temp.Next()
		temp.Scan(&PseudoUser)
		temp.Close()
		db.Close()
		if PseudoUser != Pseudo {
			_, db := OpenDataBase()
			temp, err := db.Prepare("INSERT INTO Like (PseudoUser, idPoste, isLike) VALUES(?,?,?)")
			if err != nil {
				return 500
			} else {
				temp.Exec(Pseudo, idPost, 1)
				db.Close()
				return 0
			}
		} else {
			_, db := OpenDataBase()

			result, err := db.Prepare("UPDATE Like SET isLike = 1 WHERE idPoste = ? AND PseudoUser =?")

			if err != nil {
				fmt.Println(err.Error())
				db.Close()
				return 500
			} else {
				result.Exec(Pseudo, idPost)
				db.Close()
				return 0
			}
		}
	}
}

func Dislike(idPost int, PseudoUser string) int {
	_, db := OpenDataBase()
	result, err := db.Prepare("UPDATE Like SET isLike = 0 WHERE idPoste = ? AND PseudoUser =?")
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

	temp, err := db.Query("SELECT idPoste FROM Like WHERE PseudoUser = ? AND isLike = 1 ", UserPseuso)

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
		temp.Close()
		fmt.Println(IdLiker)

		db.Close()
		for _, i := range IdLiker {
			_, temp := GetPosteByID(i)

			ResultFunc = append(ResultFunc, temp)
		}
		return 0, ResultFunc
	}
}

func IsLike(idPoste int, PseudoUser string) (int, int) {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT isLike FROM Like WHERE PseudoUser = ? AND IdPoste = ?", PseudoUser, idPoste)
	if err != nil {
		db.Close()
		return 500, 0
	} else {
		var isLike int

		result.Next()
		result.Scan(&isLike)
		result.Close()
		db.Close()

		return 0, isLike
	}
}

func GetLikeNb(id int) int {
	_, db := OpenDataBase()

	result, err := db.Query("SELECT idPoste FROM Like WHERE idPoste = ? AND isLike = 1", id)
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
		result.Close()

		db.Close()
		// fmt.Println(len(resultFunc))
		return len(resultFunc)
	}
}

func GetDislikeNb(id int) int {
	_, db := OpenDataBase()
	result, err := db.Query("SELECT idPoste FROM Like WHERE idPoste = ? AND isLike = 0", id)

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
		result.Close()

		db.Close()
		// fmt.Println(len(resultFunc))
		return len(resultFunc)
	}
}
