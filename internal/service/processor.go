package utils

import (
	"math"
	"strconv"
	"strings"
	"time"

	model "github.com/john6fa/receipt-processor-challenge/internal/model"
)

func CalculatePoints(receipt model.Receipt) int {
	points := 0

	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	if int(total*100)%25 == 0 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		descriptionLength := len(strings.TrimSpace(item.ShortDescription))
		if descriptionLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	time, _ := time.Parse("15:04", receipt.PurchaseTime)
	if time.Hour() >= 14 && time.Hour() <= 16 {
		points += 10
	}

	return points
}
