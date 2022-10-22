package models

type Product struct {
	Id          int64   `json:"id" db: "id"`
	Name        string  `json:"name" db: "name"`
	Description string  `json: "description" db: "description"`
	Value       float64 `json: "value" db: "value"`
	Status      string  `json: "status" db: "status"`
}
