package order

import "lesson-1/internal/domain/order"

type EmptyResponse struct{}

type GetOrderResponse struct {
	order.Order
}

type OrdersMapResponse struct {
	Orders map[string]order.Order
}
