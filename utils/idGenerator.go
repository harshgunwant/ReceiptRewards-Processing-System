package utils

import (
	"math/rand"
	"time"
)

const idPrefix = "REC"

// Random ID prefixed with "REC" and 10 random characters
func GenerateID() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 10)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return idPrefix + string(result)
}
