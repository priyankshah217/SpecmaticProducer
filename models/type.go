package models

type Product struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Inventory int    `json:"inventory"`
	Cost      int    `json:"cost"`
}
