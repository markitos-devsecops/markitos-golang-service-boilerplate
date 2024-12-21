package services

import (
	"markitos-svc-boilerplate/internal/domain"
	"time"
)

type BoilerUpdateRequestUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BoilerUpdateRequestBody struct {
	Message string `json:"message" binding:"required"`
}

type BoilerUpdateRequest struct {
	Id      string `uri:"id" binding:"required,uuid"`
	Message string `json:"message" binding:"required"`
}

func NewBoilerUpdateRequest(id string, message string) *BoilerUpdateRequest {
	return &BoilerUpdateRequest{
		Id:      id,
		Message: message,
	}
}

type BoilerUpdateService struct {
	Repository domain.BoilerRepository
}

func NewBoilerUpdateService(repository domain.BoilerRepository) BoilerUpdateService {
	return BoilerUpdateService{Repository: repository}
}

func (s *BoilerUpdateService) Execute(request BoilerUpdateRequest) (*domain.Boiler, error) {
	securedBoiler, err := domain.NewBoiler(request.Id, request.Message)
	if err != nil {
		return nil, err
	}

	boilerToUpdate, errExistingBoiler := s.Repository.One(&securedBoiler.Id)
	if errExistingBoiler != nil {
		return nil, errExistingBoiler
	}

	boilerToUpdate.Message = securedBoiler.Message
	boilerToUpdate.UpdatedAt = time.Now()
	err = s.Repository.Update(boilerToUpdate)
	if err != nil {
		return nil, err
	}

	return boilerToUpdate, nil
}
