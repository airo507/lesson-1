package order

import (
	"encoding/json"
	"lesson-1/internal/api"
	"net/http"
)

func (i *Implementation) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderId, ok := api.PathValueOrError(w, r, "order_id")
	if !ok {
		return
	}

	userId, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}

	card, err := i.cardService.GetByUserID(r.Context(), userId)

	err = i.orderService.CreateOrder(r.Context(), orderId, card)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to create order",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(EmptyResponse{})
}
