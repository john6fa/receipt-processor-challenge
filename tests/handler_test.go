package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
	api "github.com/john6fa/receipt-processor-challenge/internal/api"
	model "github.com/john6fa/receipt-processor-challenge/internal/model"
)

func TestProcessReceipt(t *testing.T) {
	handler := &api.Handler{
		ReceiptStore: make(map[string]model.Receipt),
	}

	payload := model.Receipt{
		Retailer:     "Test Store",
		PurchaseDate: "2024-11-30",
		PurchaseTime: "16:45",
		Items: []model.Item{
			{ShortDescription: "Test Item 1", Price: "5.00"},
			{ShortDescription: "Test Item 2", Price: "10.00"},
			{ShortDescription: "Test Item 3", Price: "15.00"},
		},
		Total: "30.00",
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/receipts/process", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.ProcessReceipts(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var result map[string]string
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err, "Expected no error while decoding response")

	_, exists := result["id"]
	assert.True(t, exists, "Response should contain an 'id' field")
	assert.NotEmpty(t, result["id"], "The 'id' field should not be empty")
}

func TestGetPoints(t *testing.T) {
	mockStore := map[string]model.Receipt{
		"test-receipt-id": {Points: 100},
	}

	handler := &api.Handler{
		ReceiptStore: mockStore,
	}

	req := httptest.NewRequest(http.MethodGet, "/receipts/test-receipt-id/points", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "test-receipt-id"})

	w := httptest.NewRecorder()

	handler.GetPoints(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

	var result map[string]int
	err := json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err, "Expected no error while decoding response")
	assert.Equal(t, 100, result["points"], "Expected points to be 42")
}
