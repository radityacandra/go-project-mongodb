package services

import (
	"testing"

	"github.com/radityacandra/go-project-mongodb/mocks/internal_/application/building/repository"
	"github.com/stretchr/testify/assert"
)

func TestInitService(t *testing.T) {
	repository := repository.NewMockIRepository(t)
	service := NewService(repository)

	assert.NotNil(t, service)
	assert.NotNil(t, service.repository)
}
