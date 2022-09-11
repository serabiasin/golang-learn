package json_loader

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

type ResultSearch struct {
	CarrierName  string `json:"carrier_name"`
	TotalHarga   uint32 `json:"base_price"`
	DeliveryTime int    `json:"DeliveryTime"`
}

// Must! setiap awal huruf variabel wajib kapital
type RequestResult struct {
	Pickup_postcode   string         `json:"pickup_postcode"`
	Delivery_postcode string         `json:"delivery_postcode"`
	Vehicle           string         `json:"vehicle"`
	Price_list        []ResultSearch `json:"price_list"`
}
