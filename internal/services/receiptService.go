package services

import (
	"FetchRewardsAssessment/internal/models"
	"FetchRewardsAssessment/internal/storage"
	"FetchRewardsAssessment/utils"
	"errors"
)

// ValidateReceipt validates the receipt fields
func ValidateReceipt(receipt models.Receipt) error {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) == 0 {
		return errors.New("one or more required fields are missing")
	}
	return nil
}

// ProcessReceipt processes the receipt by validating, generating an ID, and saving it
func ProcessReceipt(receipt *models.Receipt) (string, error) {
	if err := ValidateReceipt(*receipt); err != nil {
		return "", err
	}

	receipt.ID = utils.GenerateID()

	storage.SaveReceipt(*receipt)

	return receipt.ID, nil
}

// GetPoints retrieves a receipt by ID and calculates points
func GetPoints(id string) (int, error) {
	receipt, found := storage.GetReceipt(id)
	if !found {
		return 0, errors.New("receipt not found")
	}

	points := utils.CalculatePoints(receipt)

	return points, nil
}
