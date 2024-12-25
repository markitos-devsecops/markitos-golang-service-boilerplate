package services

import (
	"markitos-golang-service-boilerplate/internal/domain"
)

type BoilerSearchService struct {
	Repository domain.BoilerRepository
}

func NewBoilerSearchService(repository domain.BoilerRepository) BoilerSearchService {
	return BoilerSearchService{Repository: repository}
}

type BoilerSearchRequest struct {
	SearchTerm string `json:"searchTerm"`
	PageNumber int    `json:"pageNumber" bindings:"min=1"`
	PageSize   int    `json:"pageSize" bindings:"min=10,max=100"`
}

func (s *BoilerSearchService) Execute(request BoilerSearchRequest) ([]*domain.Boiler, error) {
	response, err := s.Repository.SearchAndPaginate(request.SearchTerm, request.PageNumber, request.PageSize)
	if err != nil {
		return nil, err
	}

	return response, nil
}
