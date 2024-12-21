package services

import (
	"markitos-svc-boilerplate/internal/domain"
)

type BoilerCreateRequest struct {
	Message string `json:"message" binding:"required"`
}

func NewBolilerCreateRequest(message string) BoilerCreateRequest {
	return BoilerCreateRequest{Message: message}
}

type BoilerCreateService struct {
	Repository domain.BoilerRepository
}

func NewBoilerCreateService(repository domain.BoilerRepository) BoilerCreateService {
	return BoilerCreateService{Repository: repository}
}

func (s *BoilerCreateService) Execute(request BoilerCreateRequest) (*domain.Boiler, error) {
	boiler, err := domain.NewBoiler(domain.UUIDv4(), request.Message)
	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(boiler)
	if err != nil {
		return nil, err
	}

	return boiler, nil
}
