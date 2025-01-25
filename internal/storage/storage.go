package storage

import (
	"FetchRewardsAssessment/internal/models"
	"sync"
)

var (
	store = make(map[string]models.Receipt)
	mu    sync.RWMutex
)

// SaveReceipt stores a receipt in memory
func SaveReceipt(receipt models.Receipt) {
	mu.Lock()
	defer mu.Unlock()
	store[receipt.ID] = receipt
}

// GetReceipt retrieves a receipt by ID
func GetReceipt(id string) (models.Receipt, bool) {
	mu.RLock()
	defer mu.RUnlock()
	receipt, exists := store[id]
	return receipt, exists
}
