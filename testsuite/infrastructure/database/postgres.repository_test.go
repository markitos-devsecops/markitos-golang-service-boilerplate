package database_test

import (
	"log"
	"markitos-golang-service-boilerplate/infrastructure/configuration"
	"markitos-golang-service-boilerplate/infrastructure/database"
	"markitos-golang-service-boilerplate/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestBoilerCreate(t *testing.T) {
	db := setupTestDB()
	repository := database.NewBoilerPostgresRepository(db)

	boiler, _ := domain.NewBoiler(domain.UUIDv4(), "Hello, World!")
	err := repository.Create(boiler)
	require.NoError(t, err)

	var result domain.Boiler
	err = db.First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Message, result.Message)
	require.WithinDuration(t, boiler.CreatedAt, result.CreatedAt, time.Second)
	require.WithinDuration(t, boiler.UpdatedAt, result.UpdatedAt, time.Second)

	db.Delete(&result)
}

func TestSearch(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewBoilerPostgresRepository(db)

	randomMessage := "Test " + domain.RandomString(10)
	boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: randomMessage}
	db.Create(boiler)

	results, err := repo.SearchAndPaginate(randomMessage, 1, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, randomMessage, results[0].Message)

	cleanDB(db)
}
func TestBoilerDelete(t *testing.T) {
	db := setupTestDB()
	repository := database.NewBoilerPostgresRepository(db)

	boiler, _ := domain.NewBoiler(domain.UUIDv4(), "Hello, World!")
	db.Create(boiler)

	err := repository.Delete(&boiler.Id)
	require.NoError(t, err)

	var result domain.Boiler
	err = db.First(&result, "id = ?", boiler.Id).Error
	require.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestBoilerUpdate(t *testing.T) {
	db := setupTestDB()
	repository := database.NewBoilerPostgresRepository(db)

	boiler, _ := domain.NewBoiler(domain.UUIDv4(), "Hello, World!")
	db.Create(boiler)

	boiler.Message = "Updated Message"
	err := repository.Update(boiler)
	require.NoError(t, err)

	var result domain.Boiler
	err = db.First(&result, "id = ?", boiler.Id).Error
	require.NoError(t, err)
	require.Equal(t, "Updated Message", result.Message)
}

func TestBoilerOne(t *testing.T) {
	db := setupTestDB()
	repository := database.NewBoilerPostgresRepository(db)

	boiler, _ := domain.NewBoiler(domain.UUIDv4(), "Hello, World!")
	db.Create(boiler)

	result, err := repository.One(&boiler.Id)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, boiler.Id, result.Id)
	require.Equal(t, boiler.Message, result.Message)
}

func TestBoilerList(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repository := database.NewBoilerPostgresRepository(db)

	for i := 0; i < 5; i++ {
		boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: domain.RandomString(10)}
		db.Create(boiler)
	}

	results, err := repository.List()
	require.NoError(t, err)
	require.Len(t, results, 5)
}
func TestPagination(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewBoilerPostgresRepository(db)

	for i := 0; i < 15; i++ {
		boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: domain.RandomString(10)}
		db.Create(boiler)
	}

	results, err := repo.SearchAndPaginate("", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 5)

	cleanDB(db)
}

func TestSearchAndPagination(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)
	repo := database.NewBoilerPostgresRepository(db)

	for i := 0; i < 25; i++ {
		message := "Test Boiler " + domain.RandomString(5)
		boiler := &domain.Boiler{Id: domain.UUIDv4(), Message: message}
		db.Create(boiler)
	}

	results, err := repo.SearchAndPaginate("Test", 2, 10)
	assert.NoError(t, err)
	assert.Len(t, results, 10)

	cleanDB(db)
}

func setupTestDB() *gorm.DB {
	config, err := configuration.LoadConfiguration("../../..")
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(config.DsnDatabase), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Boiler{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func cleanDB(db *gorm.DB) {
	db.Exec("DELETE FROM boilers")
}
