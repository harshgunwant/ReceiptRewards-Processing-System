package utils

import (
	"FetchRewardsAssessment/internal/models"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// CalculatePoints calculates points for a given receipt based on rules
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: 1 point for each alphanumeric character in the retailer name
	alphanumericRegex := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(alphanumericRegex.FindAllString(receipt.Retailer, -1))

	// Rule 2: 50 points if the total is a round dollar amount
	if isWholeDollar(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25
	if isMultipleOf(receipt.Total, 0.25) {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Item description length is a multiple of 3
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	if len(dateParts) == 3 {
		day, _ := strconv.Atoi(dateParts[2])
		if day%2 != 0 {
			points += 6
		}
	}

	// Rule 7: 10 points if purchase time is between 2:00 PM and 4:00 PM
	if isBetweenTime(receipt.PurchaseTime, "14:00", "16:00") {
		points += 10
	}

	return points
}

// Helper functions
func isWholeDollar(total string) bool {
	return strings.HasSuffix(total, ".00")
}

func isMultipleOf(total string, factor float64) bool {
	totalFloat, _ := strconv.ParseFloat(total, 64)
	return math.Mod(totalFloat, factor) == 0
}

func isBetweenTime(purchaseTime, start, end string) bool {
	return purchaseTime >= start && purchaseTime < end
}
