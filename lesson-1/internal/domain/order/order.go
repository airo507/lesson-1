package order

import "lesson-1/internal/domain/card"

type Order struct {
	OrderId string    `json:"order_id"`
	UserID  string    `json:"user_id"`
	Card    card.Card `json:"cards"`
}
