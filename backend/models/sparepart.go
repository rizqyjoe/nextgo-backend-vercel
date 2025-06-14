package models

type Sparepart struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Unit     string `json:"unit"`
	Category string `json:"category"`
}
