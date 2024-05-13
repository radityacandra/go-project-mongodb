package handler

import (
	"testing"

	"github.com/radityacandra/go-project-mongodb/internal/server/types"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	handler := NewHandler(&types.Dependency{})

	assert.NotNil(t, handler.service)
}
