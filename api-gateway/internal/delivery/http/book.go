package http

import (
	"encoding/json"
	"net/http"

	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/dto"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := h.BookHandler.GetBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func (h *Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book dto.AddBookInput
	err := json.NewDecoder(r.Body).Decode(&book)
	ctx := r.Context()
	if err != nil {
		return
	}
	if book.Name == "" || book.Author == "" {
		http.Error(w, "name and author are required", http.StatusBadRequest)
		return
	}
	output, err := h.BookHandler.AddBook(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		return
	}

}
