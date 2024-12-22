package services_test

import (
	"markitos-service-boilerplate/internal/domain"
	"markitos-service-boilerplate/internal/services"
	"os"
	"testing"
)

const VALID_UUIDV4 = "f47ac10b-58cc-4372-a567-0e02b2c3d479"
const VALID_MESSAGE = "any valid message"

var boilerMockSpyRepository domain.BoilerRepository
var boilerCreateService services.BoilerCreateService
var boilerOneService services.BoilerOneService
var boilerListService services.BoilerListService
var boilerUpdateService services.BoilerUpdateService

func TestMain(m *testing.M) {
	boilerMockSpyRepository = NewMockSpyBoilerRepository()
	boilerCreateService = services.NewBoilerCreateService(boilerMockSpyRepository)
	boilerOneService = services.NewBoilerOneService(boilerMockSpyRepository)
	boilerListService = services.NewBoilerListService(boilerMockSpyRepository)
	boilerUpdateService = services.NewBoilerUpdateService(boilerMockSpyRepository)

	os.Exit(m.Run())
}

type MockSpyBoilerRepository struct {
	LastCreatedBoiler       *domain.Boiler
	LastCreatedForOneBoiler *domain.Boiler
	OneCalled               bool
	LastUpdatedBoiler       *domain.Boiler
}

func NewMockSpyBoilerRepository() *MockSpyBoilerRepository {
	return &MockSpyBoilerRepository{
		LastCreatedBoiler:       nil,
		LastCreatedForOneBoiler: nil,
		OneCalled:               false,
		LastUpdatedBoiler:       nil,
	}
}

func (m *MockSpyBoilerRepository) Create(boiler *domain.Boiler) error {
	m.LastCreatedBoiler = boiler
	m.LastCreatedForOneBoiler = boiler

	return nil
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWith(boiler *domain.Boiler) bool {
	var result bool = m.LastCreatedBoiler.Id == boiler.Id && m.LastCreatedBoiler.Message == boiler.Message

	m.LastCreatedBoiler = nil

	return result
}

func (m *MockSpyBoilerRepository) CreateHaveBeenCalledWithMessage(boiler *domain.Boiler) bool {
	var result bool = m.LastCreatedBoiler.Message == boiler.Message

	m.LastCreatedBoiler = nil

	return result
}

func (m *MockSpyBoilerRepository) Delete(id *string) error {
	return nil
}

func (m *MockSpyBoilerRepository) Update(boiler *domain.Boiler) error {
	m.LastUpdatedBoiler = boiler

	return nil
}

func (m *MockSpyBoilerRepository) One(id *string) (*domain.Boiler, error) {
	return &domain.Boiler{
		Id:      *id,
		Message: VALID_MESSAGE,
	}, nil
}

func (m *MockSpyBoilerRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*domain.Boiler, error) {
	return []*domain.Boiler{
		{
			Id:      VALID_UUIDV4,
			Message: VALID_MESSAGE,
		},
	}, nil
}

func (m *MockSpyBoilerRepository) OneHaveBeenCalledWith(boiler *domain.Boiler) bool {
	var result bool = m.LastCreatedForOneBoiler.Id == boiler.Id && m.LastCreatedForOneBoiler.Message == boiler.Message

	m.LastCreatedForOneBoiler = nil

	return result
}

func (m *MockSpyBoilerRepository) OneHaveBeenCalledWithMessage(boiler *domain.Boiler) bool {
	var result bool = m.LastCreatedForOneBoiler.Message == boiler.Message && m.LastCreatedForOneBoiler.Id == boiler.Id

	m.LastCreatedBoiler = nil

	return result
}

func (m *MockSpyBoilerRepository) UpdateHaveBeenCalledWithMessage(boiler *domain.Boiler) bool {
	var result bool = m.LastUpdatedBoiler.Message == boiler.Message && m.LastUpdatedBoiler.Id == boiler.Id

	m.LastUpdatedBoiler = nil

	return result
}

func (m *MockSpyBoilerRepository) List() ([]*domain.Boiler, error) {
	m.OneCalled = true

	return []*domain.Boiler{}, nil
}
func (m *MockSpyBoilerRepository) ListHaveBeenCalled() bool {
	var result bool = m.OneCalled

	m.OneCalled = false

	return result
}
