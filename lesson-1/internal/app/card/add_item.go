package card

import (
	"encoding/json"
	stderr "errors"
	"lesson-1/internal/api"
	"lesson-1/internal/errors"
	"net/http"
)

func (i *Implementation) AddItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}

	itemId, ok := api.PathValueOrError(w, r, "item_id")
	if !ok {
		return
	}
	req := AddItemRequest{
		ItemID: itemId,
	}
	//err := json.NewDecoder(r.Body).Decode(&req)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	_ = json.NewEncoder(w).Encode(api.DefaultResponse{
	//		Code:    api.InvalidRequest,
	//		Message: "failed to decode request",
	//		Request: err.Error(),
	//	})
	//	return
	//}

	item, err := i.itemService.GetItemByID(r.Context(), req.ItemID)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "item not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to get item by id",
		})
		return
	}

	card, err := i.cardService.AddItem(r.Context(), userID, item)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "card not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to add item",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(DefaultResponse{
		Card: card,
	})
}
