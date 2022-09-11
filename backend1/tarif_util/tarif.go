package tarif_util

import "github.com/martinlindhe/base36"

func GetPrice(pickup_postcode string, delivery_postcode string) uint64 {

	return uint64(base36.Decode(pickup_postcode)-base36.Decode(delivery_postcode)) / 100000000
}

func AdjustPrice(jenisKendaran string, harga float32) uint32 {
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
