package utils

import (
	"FetchRewardsAssessment/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name        string
		receipt     models.Receipt
		expectedPts int
	}{
		{
			name: "Target Receipt",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Total:        "35.35",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
			},
			expectedPts: 28,
		},
		{
			name: "M&M Corner Market Receipt",
			receipt: models.Receipt{
				Retailer:     "M&M Corner Market",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "14:33",
				Total:        "9.00",
				Items: []models.Item{
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
					{ShortDescription: "Gatorade", Price: "2.25"},
				},
			},
			expectedPts: 109, // Breakdown:
			/*
				- Retailer name: 14 alphanumeric characters = 14 points
				- Total is a round dollar amount: 50 points
				- Total is a multiple of 0.25: 25 points
				- 4 items (2 pairs): 10 points
				- Purchase time is between 2:00 PM and 4:00 PM: 10 points
			*/
		},
		{
			name: "Receipt with Edge Cases",
			receipt: models.Receipt{
				Retailer:     "EdgeCaseMart",
				PurchaseDate: "2023-03-01",
				PurchaseTime: "10:15",
				Total:        "0.75",
				Items: []models.Item{
					{ShortDescription: "Small Item", Price: "0.25"},
					{ShortDescription: "Discounted", Price: "0.50"},
				},
			},
			expectedPts: 48, // Breakdown:
			/*
				- Retailer name: 12 alphanumeric characters = 12 points
				- Total is a multiple of 0.25: 25 points
				- 2 items (1 pair): 5 points
				- Item description length is a multiple of 3: 0.2*0.25 = 0.05 ~ 0 points
				- Purchase date day is odd: 6 points

			*/
		},
	}

	// Iterate over all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := CalculatePoints(tt.receipt)
			println(tt.name, points)
			assert.Equal(t, tt.expectedPts, points, "Points should match the expected value")
		})
	}
}
