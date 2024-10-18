package repositories

import (
	"context"
	"fmt"
	"lesson-1/internal/domain/order"
	"lesson-1/internal/errors"
)

type OrderRepository struct {
	orders map[string]order.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]order.Order),
	}
}

func (o *OrderRepository) GetOrderById(ctx context.Context, orderId string) (order.Order, error) {
	select {
	case <-ctx.Done():
		return order.Order{}, ctx.Err()
	default:
	}

	orderItem, exist := o.orders[orderId]
	if exist {
		return orderItem, nil
	}
	return order.Order{}, fmt.Errorf("order not found: %w", errors.NotFound)
}

func (o *OrderRepository) GetOrdersByUserId(ctx context.Context, userId string) (map[string]order.Order, error) {
	select {
	case <-ctx.Done():
		return map[string]order.Order{}, ctx.Err()
	default:
	}

	userOrders := make(map[string]order.Order)
	for _, orderElem := range o.orders {
		if orderElem.UserID == userId {
			userOrders[orderElem.OrderId] = orderElem
		}
	}
	if len(userOrders) == 0 {
		return userOrders, fmt.Errorf("orderList not found: %w", errors.NotFound)
	}
	return userOrders, nil
}

func (o *OrderRepository) Save(ctx context.Context, order order.Order) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	o.orders[order.UserID] = order

	return nil
}
