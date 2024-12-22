package services

import (
	"markitos-service-boilerplate/internal/domain"
)

type BoilerListService struct {
	Repository domain.BoilerRepository
}

func NewBoilerListService(repository domain.BoilerRepository) BoilerListService {
	return BoilerListService{Repository: repository}
}

func (s *BoilerListService) Execute() ([]*domain.Boiler, error) {
	response, err := s.Repository.List()
	if err != nil {
		return nil, err
	}

	return response, nil
}
