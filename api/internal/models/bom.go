package models

type BOM struct {
	Materials []Material
}

type Material struct {
	Name            string `json:"name"`
	Unit            string `json:"unit"`
	PricePerUnit    string `json:"price_per_unit"`
	Quantity        string `json:"quantity"`
	Currency        string `json:"currency"`
	ManufactureDate string `json:"manufacture_date"`
}
