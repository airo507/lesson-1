package item

import domainitem "lesson-1/internal/domain/item"

type GetItemByIDResponse struct {
	domainitem.Item
}

type GetItemsResponse struct {
	Items []domainitem.Item `json:"items"`
}
