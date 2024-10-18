package repositories

import (
	"context"
	"fmt"
	domainitem "lesson-1/internal/domain/item"
	"lesson-1/internal/errors"
)

type ItemRepository struct {
	items []domainitem.Item
}

func NewItemRepository(items []domainitem.Item) *ItemRepository {
	return &ItemRepository{
		items: items,
	}
}

func (r *ItemRepository) GetItemByID(ctx context.Context, itemID string) (domainitem.Item, error) {
	select {
	case <-ctx.Done():
		return domainitem.Item{}, ctx.Err()
	default:
	}

	itemsMap := r.getItemsMap()
	item, exist := itemsMap[itemID]
	if exist {
		return item, nil
	}

	return domainitem.Item{}, fmt.Errorf("item not found: %w", errors.NotFound)
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]domainitem.Item, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return r.items, nil
}

func (r *ItemRepository) getItemsMap() map[string]domainitem.Item {
	itemsMap := make(map[string]domainitem.Item)
	for _, item := range r.items {
		itemsMap[item.ID] = item
	}
	return itemsMap
}
