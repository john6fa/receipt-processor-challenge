package tests

import (
	"testing"

	model "github.com/john6fa/receipt-processor-challenge/internal/model"
	service "github.com/john6fa/receipt-processor-challenge/internal/service"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  model.Receipt
		expected int
	}{
		{
			name: "Test 1 - retailer name has 6 characters, 5 items, item description multiple of 3, odd purchase day",
			receipt: model.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2024-01-01",
				PurchaseTime: "13:01",
				Items: []model.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Test 2 - total is a round dollar amount, total is muliple of 0.25, retailer name has 14 alphanum chars, time between 2 and 4 PM, 4 items",
			receipt: model.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2024-03-20",
				PurchaseTime: "14:33",
				Items: []model.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
				Total: "9.00",
			},
			expected: 109,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := service.CalculatePoints(tt.receipt)
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
