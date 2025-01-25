package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()

	// Assert ID format
	assert.Contains(t, id1, "REC")
	assert.Contains(t, id2, "REC")

	// Assert IDs are unique
	assert.NotEqual(t, id1, id2)
}
