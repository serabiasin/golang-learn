package main

import (
	"jnet/json_loader"
	"jnet/tarif_util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func getQuotes(c echo.Context) error {
	buffer := json_loader.Load_file("./json_loader/carrier-data.json")
	pickup_postcode := c.QueryParam("pickup_postcode")
	delivery_postcode := c.QueryParam("delivery_postcode")
	vehicle := c.QueryParam("vehicle")
	tarif := tarif_util.AdjustPrice(vehicle, float32(tarif_util.GetPrice(pickup_postcode, delivery_postcode)))

	resultSearch := json_loader.Search_json(vehicle, tarif, buffer)
	QueryResult := json_loader.Final_result(vehicle, pickup_postcode, delivery_postcode, resultSearch)

	// j, err := json.Marshal(QueryResult)

	// if err != nil {
	// 	log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
	// }
	// fmt.Printf("\nemployee2 JSON: %s\n", string(j))

	// tidak perlu di convert ke byte, library echo sudah menghandle
	return c.JSON(http.StatusCreated, QueryResult)

}
func main() {
	e := echo.New()

	e.GET("/quotes", getQuotes)
	e.Logger.Fatal(e.Start(":1323"))
}
