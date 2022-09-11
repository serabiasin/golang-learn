package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinlindhe/base36"
)

type resultSearch struct {
	CarrierName  string `json:"carrier_name"`
	TotalHarga   uint32 `json:"base_price"`
	DeliveryTime int    `json:"DeliveryTime"`
}

type requestResult struct {
	pickup_postcode   string         `json:"pickup_postcode"`
	delivery_postcode string         `json:"delivery_postcode"`
	vehicle           string         `json:"vehicle"`
	price_list        []resultSearch `json:"price_list"`
}

// nama entry dalam bentuk json (untuk tag `json:"nama"`)
type CarrierData []struct {
	CarrierName string `json:"carrier_name"`
	BasePrice   int    `json:"base_price"`
	Services    []struct {
		DeliveryTime int      `json:"delivery_time"`
		Markup       int      `json:"markup"`
		Vehicles     []string `json:"vehicles"`
	} `json:"services"`
}

func getPrice(pickup_postcode string, delivery_postcode string) uint64 {

	return uint64(base36.Decode(pickup_postcode)-base36.Decode(delivery_postcode)) / 100000000
}

func adjustPrice(jenisKendaran string, harga float32) uint32 {
	if jenisKendaran == "bicycle" {
		return uint32(float32(harga) + float32(harga*0.1))

	} else if jenisKendaran == "motorbike" {
		return uint32(float32(harga) + float32(harga*0.15))
	} else if jenisKendaran == "parcel_car" {
		return uint32(float32(harga) + float32(harga*0.2))
	} else if jenisKendaran == "small_van" {
		return uint32(float32(harga) + float32(harga*0.3))
	} else if jenisKendaran == "large_van" {
		return uint32(float32(harga) + float32(harga*0.4))
	}
	return 0
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("./carrier-data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var users CarrierData

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	var key_find string = "bicycle"
	var sortedSearch = make([]resultSearch, 0)
	//  sortedSearch := []resultSearch
	// fmt.Println(users)
	for indeks_object, object := range users {
		// fmt.Println(object)
		for indeks_service, service := range object.Services {
			// fmt.Println(service)
			for indeks_vehic, vehicle := range service.Vehicles {
				if key_find == vehicle {
					var buffer resultSearch
					buffer.CarrierName = users[indeks_object].CarrierName
					buffer.DeliveryTime = users[indeks_object].Services[indeks_service].DeliveryTime
					buffer.TotalHarga = adjustPrice(key_find, float32(getPrice("SW1A1AA", "EC2A3LT")))

					fmt.Println(users[indeks_object].Services[indeks_service].Vehicles[indeks_vehic])

					// append tapi harus assign, dia bukan pass by reference
					sortedSearch = append(sortedSearch, buffer)
				}
			}
		}
	}
	fmt.Println("Hasil", float32(getPrice("SW1A1AA", "EC2A3LT")))
	var queryResult requestResult
	queryResult.pickup_postcode = "SW1A1AA"
	queryResult.delivery_postcode = "EC2A3LT"
	queryResult.vehicle = key_find
	queryResult.price_list = sortedSearch
	fmt.Println(queryResult)
}
