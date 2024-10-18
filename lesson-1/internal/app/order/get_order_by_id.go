package order

import (
	"encoding/json"
	stderr "errors"
	"lesson-1/internal/api"
	"lesson-1/internal/errors"
	"net/http"
)

func (i *Implementation) GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orderID, ok := api.PathValueOrError(w, r, "order_id")
	if !ok {
		return
	}

	orderElem, err := i.orderService.GetOrderById(r.Context(), orderID)
	if stderr.Is(err, errors.NotFound) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.DefaultResponse{
			Code:    api.NotFound,
			Message: "order not found",
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

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(GetOrderResponse{
		orderElem,
	})
}
