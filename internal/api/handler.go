package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/john6fa/receipt-processor-challenge/internal/model"
	service "github.com/john6fa/receipt-processor-challenge/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	ReceiptStore map[string]model.Receipt
}

func NewHandler() *Handler {
	return &Handler{
		ReceiptStore: make(map[string]model.Receipt),
	}
}

func (h *Handler) ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var receipt model.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := service.CalculatePoints(receipt)
	receipt.ID = id
	receipt.Points = points
	h.ReceiptStore[id] = receipt

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	receipt, exists := h.ReceiptStore[id]
	if !exists {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": receipt.Points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
