package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Film struct {
	Id     int32   `json:"id"`
	Movie  string  `json:"movie"`
	Genre  string  `json:"genre"`
	Rating float32 `json:"rating"`
}

func CloseDB(db *sql.DB) {
	db.Close()

}
func LoadDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	checkErr(err)

	return db
}

func ShowTable(db *sql.DB) []Film {
	var queryResult = make([]Film, 0)

	rows, err := db.Query("SELECT * FROM FILM")
	checkErr(err)

	for rows.Next() {
		var buffer Film
		err = rows.Scan(&buffer.Id, &buffer.Movie, &buffer.Genre, &buffer.Rating)
		checkErr(err)
		queryResult = append(queryResult, buffer)
	}

	rows.Close()
	return queryResult
}

func RatingTable(db *sql.DB, rating float64) []Film {
	var queryResult = make([]Film, 0)

	rows, err := db.Query("SELECT * FROM FILM WHERE rating > ? ORDER BY rating", rating)
	checkErr(err)

	for rows.Next() {
		var buffer Film
		err = rows.Scan(&buffer.Id, &buffer.Movie, &buffer.Genre, &buffer.Rating)
		checkErr(err)
		queryResult = append(queryResult, buffer)
	}

	rows.Close()
	return queryResult
}

func UpdateDB(db *sql.DB, movie string, genre string, rating float32) int16 {
	//INSERT INTO demo (movie, genre,rating) VALUES ('year', 3021);
	stmt, err := db.Prepare("INSERT INTO FILM (movie, genre,rating) VALUES (?, ?,?)")
	checkErr(err)

	res, err := stmt.Exec(movie, genre, rating)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("Last Id :", id)

	if err == nil {
		return 200
	} else {
		return 500
	}
}
