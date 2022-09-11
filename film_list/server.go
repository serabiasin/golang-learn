package main

// same as from filmlist.db import driver
import (
	driver "filmlist/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func getDatabase(c echo.Context) error {
	db := driver.LoadDB("./db/sqlite.db")
	listDB := driver.ShowTable(db)

	driver.CloseDB(db)

	return c.JSON(http.StatusCreated, listDB)

}

func updateDatabase(c echo.Context) error {
	db := driver.LoadDB("./db/sqlite.db")

	film := c.QueryParam("film")
	genre := c.QueryParam("genre")
	rating, _ := strconv.ParseFloat(c.QueryParam("rating"), 32)

	code := driver.UpdateDB(db, film, genre, float32(rating))
	driver.CloseDB(db)

	if code != 500 {
		return c.String(http.StatusOK, "OK")
	} else {
		return c.String(http.StatusBadRequest, "Bad Query")
	}
}

func SearchRating(c echo.Context) error {
	db := driver.LoadDB("./db/sqlite.db")
	rating, _ := strconv.ParseFloat(c.QueryParam("rating"), 32)

	listDB := driver.RatingTable(db, rating)
	driver.CloseDB(db)

	return c.JSON(http.StatusCreated, listDB)

}

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/list", getDatabase)
	e.GET("/list_rating", SearchRating)
	e.POST("/update", updateDatabase)

	e.Logger.Fatal(e.Start(":1323"))
}
