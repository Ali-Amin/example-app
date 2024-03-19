package models

type Product struct {
	Name string `json:"name" db:"name"`
	ID   string `json:"id" db:"id"`
}
