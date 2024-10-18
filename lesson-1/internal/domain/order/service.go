package order

import (
	"context"
	"lesson-1/internal/domain/card"
)

type Repository interface {
	Save(ctx context.Context, order Order) error
	GetOrdersByUserId(ctx context.Context, userId string) (map[string]Order, error)
	GetOrderById(ctx context.Context, orderId string) (Order, error)
}

type Service struct {
	repo Repository
}

func NewOrderService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(ctx context.Context, orderId string, currentCard card.Card) error {
	err := s.repo.Save(ctx, Order{
		OrderId: orderId,
		UserID:  currentCard.UserID,
		Card:    currentCard,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetOrdersByUserId(ctx context.Context, orderId string) (map[string]Order, error) {
	order, err := s.repo.GetOrdersByUserId(ctx, orderId)
	if err != nil {
		return map[string]Order{}, err
	}

	return order, nil
}

func (s *Service) GetOrderById(ctx context.Context, userId string) (Order, error) {
	order, err := s.repo.GetOrderById(ctx, userId)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}
