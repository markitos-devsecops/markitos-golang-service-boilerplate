package services

import (
	"markitos-golang-service-boilerplate/internal/domain"
)

type BoilerOneRequest struct {
	Id string `uri:"id" binding:"required" minLength:"36" maxLength:"36"`
}

func NewBoilerOneRequest(id string) BoilerOneRequest {
	return BoilerOneRequest{Id: id}
}

type BoilerOneService struct {
	Repository domain.BoilerRepository
}

func NewBoilerOneService(repository domain.BoilerRepository) BoilerOneService {
	return BoilerOneService{Repository: repository}
}

func (s *BoilerOneService) Execute(request BoilerOneRequest) (*domain.Boiler, error) {
	requestedId := &request.Id
	boilerId, err := domain.NewBoilerId(*requestedId)
	if err != nil {
		return nil, err
	}

	secureId := boilerId.Value()
	response, err := s.Repository.One(&secureId)
	if err != nil {
		return nil, err
	}

	return response, nil
}
