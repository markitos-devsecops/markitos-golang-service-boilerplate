package domain

import (
	"errors"
	"strings"
)

type BoilerInMemoryRepository struct {
	boilers map[string]*Boiler
}

func NewBoilerInMemoryRepository() *BoilerInMemoryRepository {
	return &BoilerInMemoryRepository{
		boilers: make(map[string]*Boiler),
	}
}

func (s *BoilerInMemoryRepository) Create(boiler *Boiler) error {
	s.boilers[boiler.Id] = boiler
	return nil
}

func (s *BoilerInMemoryRepository) Delete(id *string) error {
	if _, exists := s.boilers[*id]; !exists {
		return errors.New("boiler not found")
	}
	delete(s.boilers, *id)
	return nil
}

func (s *BoilerInMemoryRepository) One(id *string) (*Boiler, error) {
	boiler, exists := s.boilers[*id]
	if !exists {
		return nil, errors.New("boiler not found")
	}
	return boiler, nil
}

func (s *BoilerInMemoryRepository) Update(boiler *Boiler) error {
	existingBoiler, err := s.One(&boiler.Id)
	if err != nil {
		return err
	}
	existingBoiler.Message = boiler.Message
	existingBoiler.UpdatedAt = boiler.UpdatedAt
	return nil
}

func (s *BoilerInMemoryRepository) List() ([]*Boiler, error) {
	var result []*Boiler
	for _, value := range s.boilers {
		result = append(result, value)
	}
	return result, nil
}

func (s *BoilerInMemoryRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*Boiler, error) {
	var filtered []*Boiler
	for _, boiler := range s.boilers {
		if strings.Contains(boiler.Message, searchTerm) {
			filtered = append(filtered, boiler)
		}
	}

	start := (pageNumber - 1) * pageSize
	end := start + pageSize

	if start > len(filtered) {
		return []*Boiler{}, nil
	}

	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[start:end], nil
}
