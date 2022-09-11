package json_loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func Load_file(filepath string) CarrierData {
	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Database.json")
	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Database array
	var Database CarrierData

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'Database' which we defined above
	json.Unmarshal(byteValue, &Database)

	return Database
}

func Search_json(key_find string, final_price uint32, Database CarrierData) []ResultSearch {
	var sortedSearch = make([]ResultSearch, 0)
	//  sortedSearch := []resultSearch
	// fmt.Println(Database)
	for indeks_object, object := range Database {
		// fmt.Println(object)
		for indeks_service, service := range object.Services {
			// fmt.Println(service)
			for _, vehicle := range service.Vehicles {
				if key_find == vehicle {
					var buffer ResultSearch
					buffer.CarrierName = Database[indeks_object].CarrierName
					buffer.DeliveryTime = Database[indeks_object].Services[indeks_service].DeliveryTime
					buffer.TotalHarga = final_price + uint32(Database[indeks_object].BasePrice)
					//  adjustPrice(key_find, float32(getPrice("SW1A1AA", "EC2A3LT")))

					// fmt.Println(Database[indeks_object].Services[indeks_service].Vehicles[indeks_vehic])

					// append tapi harus assign, dia bukan pass by reference
					sortedSearch = append(sortedSearch, buffer)
				}
			}
		}

	}
	sort.Slice(sortedSearch, func(i, j int) bool {
		return sortedSearch[i].TotalHarga < sortedSearch[j].TotalHarga
	})
	return sortedSearch
}

func Final_result(vehicle string, pickup_postcode string, delivery_postcode string,
	buffer []ResultSearch) RequestResult {
	// fmt.Println("Hasil", float32(getPrice("SW1A1AA", "EC2A3LT")))
	var queryResult RequestResult
	queryResult.Pickup_postcode = pickup_postcode
	queryResult.Delivery_postcode = delivery_postcode
	queryResult.Vehicle = vehicle
	queryResult.Price_list = buffer

	return queryResult
}
