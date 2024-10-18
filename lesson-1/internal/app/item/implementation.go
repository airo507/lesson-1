package item

import (
	"context"
	domainitem "lesson-1/internal/domain/item"
	"net/http"
)

type Service interface {
	GetItemByID(ctx context.Context, itemID string) (domainitem.Item, error)
	GetItems(ctx context.Context) ([]domainitem.Item, error)
}

type Implementation struct {
	service Service
}

func NewItemServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}

func RegisterRoutes(mux *http.ServeMux, i *Implementation) {
	mux.HandleFunc("GET /items", i.GetItems)
	mux.HandleFunc("GET /items/{item_id}", i.GetItemByID)
}
