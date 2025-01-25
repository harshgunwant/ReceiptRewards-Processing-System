package handlers

import (
	"FetchRewardsAssessment/internal/models"
	"FetchRewardsAssessment/internal/storage"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Set up the router
	r := gin.Default()
	r.POST("/receipts/process", ProcessReceipt)

	// Create a valid receipt payload
	payload := models.Receipt{
		Retailer:     "TestStore123",
		PurchaseDate: "2025-01-01",
		PurchaseTime: "15:30",
		Total:        "50.00",
		Items: []models.Item{
			{ShortDescription: "Product1", Price: "10.00"},
			{ShortDescription: "Product2", Price: "15.00"},
		},
	}

	// Convert payload to JSON
	body, _ := json.Marshal(payload)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)

	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Contains(t, resp, "id")

	// Verify the receipt was saved in memory
	id := resp["id"]
	savedReceipt, found := storage.GetReceipt(id)
	assert.True(t, found)
	assert.Equal(t, payload.Retailer, savedReceipt.Retailer)
}

func TestProcessReceipt_InvalidReceiptFields(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Set up the router
	r := gin.Default()
	r.POST("/receipts/process", ProcessReceipt)

	// Create an invalid receipt payload (missing or incorrect required fields)
	payload := models.Receipt{
		Retailer:     "", // Missing retailer
		PurchaseDate: "", // Missing purchase date
		PurchaseTime: "15:30",
		Total:        "5.00", // Valid total, but the payload still has missing fields
		Items:        nil,    // No items provided
	}

	// Convert payload to JSON
	body, _ := json.Marshal(payload)

	// Create a new HTTP request
	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Parse the response body
	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Check error message
	assert.Equal(t, "The receipt is invalid.", resp["message"])
	assert.Contains(t, resp["details"], "one or more required fields are missing")
}
