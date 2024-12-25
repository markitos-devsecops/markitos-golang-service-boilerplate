package domain_test

import (
	"markitos-golang-service-boilerplate/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupInMemoryRepo() *domain.BoilerInMemoryRepository {
	return domain.NewBoilerInMemoryRepository()
}

func TestCreate(t *testing.T) {
	repo := setupInMemoryRepo()

	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	err := repo.Create(boiler)
	assert.NoError(t, err)

	result, err := repo.One(&boiler.Id)
	assert.NoError(t, err)
	assert.Equal(t, boiler.Message, result.Message)
}

func TestDelete(t *testing.T) {
	repo := setupInMemoryRepo()

	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	repo.Create(boiler)

	err := repo.Delete(&boiler.Id)
	assert.NoError(t, err)

	_, err = repo.One(&boiler.Id)
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	repo := setupInMemoryRepo()

	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	err := repo.Create(boiler)
	assert.NoError(t, err)

	updatedMessage := "Updated Boiler"
	boiler.Message = updatedMessage
	err = repo.Update(boiler)
	assert.NoError(t, err)

	result, err := repo.One(&boiler.Id)
	assert.NoError(t, err)
	assert.Equal(t, updatedMessage, result.Message)
}

func TestOne(t *testing.T) {
	repo := setupInMemoryRepo()

	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler"}
	err := repo.Create(boiler)
	assert.NoError(t, err)

	result, err := repo.One(&boiler.Id)
	assert.NoError(t, err)
	assert.Equal(t, boiler.Message, result.Message)
}

func TestList(t *testing.T) {
	repo := setupInMemoryRepo()

	boiler1 := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler 1"}
	boiler2 := &domain.Boiler{Id: domain.UUIDv4(), Message: "Test Boiler 2"}
	err := repo.Create(boiler1)
	assert.NoError(t, err)
	err = repo.Create(boiler2)
	assert.NoError(t, err)

	results, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestSearchAndPaginate(t *testing.T) {
	repo := setupInMemoryRepo()

	for i := 0; i < 25; i++ {
		message := "Test Boiler " + domain.RandomString(5)
		boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: message}
		err := repo.Create(boiler)
		assert.NoError(t, err)
	}

	results, err := repo.SearchAndPaginate("Test", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 10)
}
