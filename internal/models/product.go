package models

import "time"

// Product представляет продукт в системе OZON
type Product struct {
	ProductID  int64     `json:"product_id"`
	OfferID    string    `json:"offer_id"`
	Name       string    `json:"name"`
	Price      string    `json:"price"`
	Stock      int       `json:"stock"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
}

// Order представляет заказ в системе OZON
type Order struct {
	PostingNumber string    `json:"posting_number"`
	OrderID       int64     `json:"order_id"`
	OrderNumber   string    `json:"order_number"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	Products      []Product `json:"products"`
}

// StockUpdate представляет обновление запаса продуктов
type StockUpdate struct {
	OfferID string `json:"offer_id"`
	Stock   int    `json:"stock"`
}
