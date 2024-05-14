package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitBuilding(t *testing.T) {
	building := NewBuilding("userId", "some building", 1, 2)

	assert.NotEqual(t, "", building.Id)
	assert.Equal(t, "userId", building.UserId)
	assert.Equal(t, "some building", building.Name)
	assert.Equal(t, float64(1), building.Latitude)
	assert.Equal(t, float64(2), building.Longitude)
}
