package models

import "time"

type Product struct {
	Id          string    `json:"id"`
	ProductName string    `json:"product_name"`
	CreateAt    time.Time `json:"create_at"`
	CreatedBy   string    `json:"created_by"`
}
