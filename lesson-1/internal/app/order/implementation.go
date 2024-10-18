package order

import (
	"context"
	"lesson-1/internal/app/card"
	domaincard "lesson-1/internal/domain/card"
	"lesson-1/internal/domain/order"
	"net/http"
)

type Service interface {
	CreateOrder(ctx context.Context, orderId string, curCard domaincard.Card) error
	GetOrderById(ctx context.Context, orderId string) (order.Order, error)
	GetOrdersByUserId(ctx context.Context, userId string) (map[string]order.Order, error)
}

type CardService interface {
	GetByUserID(ctx context.Context, userID string) (domaincard.Card, error)
}

type Implementation struct {
	orderService Service
	cardService  CardService
}

func NewOrderImplementation(orderService Service, cardService card.Service) *Implementation {
	return &Implementation{
		orderService: orderService,
		cardService:  cardService,
	}
}

func RegisterRoutes(mux *http.ServeMux, i *Implementation) {
	mux.HandleFunc("POST  /orders/{order_id}/card/{card_id}", i.CreateOrder)
	mux.HandleFunc("GET /order/{order_id}", i.GetOrderById)
	mux.HandleFunc("GET /orders/{user_id}", i.GetOrdersByUser)
}
