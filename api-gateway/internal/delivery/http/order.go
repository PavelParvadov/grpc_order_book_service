package http

import (
	"encoding/json"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/dto"
	"net/http"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orders, err := h.OrderBookHandler.GetOrders(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		return
	}
}

func (h *Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var order dto.AddOrderInput
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	output, err := h.OrderBookHandler.AddOrder(ctx, order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}
}
