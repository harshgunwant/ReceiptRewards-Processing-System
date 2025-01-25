package handlers

import (
	"FetchRewardsAssessment/internal/models"
	"FetchRewardsAssessment/internal/storage"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPoints(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Set up the router
	r := gin.Default()
	r.GET("/receipts/:id/points", GetPoints)

	// Add a valid receipt to the in-memory storage
	receipt := models.Receipt{
		ID:           "REC12345",
		Retailer:     "Shop123",
		PurchaseDate: "2025-01-01",
		PurchaseTime: "15:30",
		Total:        "20.00",
		Items: []models.Item{
			{ShortDescription: "Item1", Price: "10.00"},
		},
	}
	storage.SaveReceipt(receipt)

	// Test case for valid ID
	t.Run("ValidReceiptID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/receipts/REC12345/points", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp map[string]int
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp, "points")
		assert.Greater(t, resp["points"], 0)
	})

	// Test case for invalid ID
	t.Run("InvalidReceiptID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/receipts/INVALID_ID/points", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)

		assert.Equal(t, "No receipt found for that ID.", resp["message"])
		assert.Contains(t, resp["details"], "receipt not found")
	})
}
