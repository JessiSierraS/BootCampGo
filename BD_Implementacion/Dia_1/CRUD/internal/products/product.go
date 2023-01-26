package products

import "time"

type Product struct {
    ID           int        `json:"id"`
    Name         string     `json:"nombre"`
    Quantity     int        `json:"cantidad"`
    Code_value   string     `json:"codigo valor"`
    Is_published string     `json:"publicado"`
    Expiration   time.Time  `json:"vencimiento"`
    Price        float64    `json:"precio"`
}