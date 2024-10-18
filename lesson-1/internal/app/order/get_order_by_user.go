package order

import (
	"encoding/json"
	stderr "errors"
	"lesson-1/internal/api"
	"lesson-1/internal/errors"
	"net/http"
)

func (i *Implementation) GetOrdersByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID, ok := api.PathValueOrError(w, r, "user_id")
	if !ok {
		return
	}

	ordersMap, err := i.orderService.GetOrdersByUserId(r.Context(), userID)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "orders not found",
		})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.InternalError,
			Message: "failed to get orders by user id",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(OrdersMapResponse{
		Orders: ordersMap,
	})
}
