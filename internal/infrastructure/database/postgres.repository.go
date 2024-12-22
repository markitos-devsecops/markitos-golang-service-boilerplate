package database

import (
	"fmt"
	"markitos-service-boilerplate/internal/domain"

	"gorm.io/gorm"
)

type BoilerPostgresRepository struct {
	db *gorm.DB
}

func NewBoilerPostgresRepository(db *gorm.DB) *BoilerPostgresRepository {
	return &BoilerPostgresRepository{db: db}
}

func (r *BoilerPostgresRepository) Create(boiler *domain.Boiler) error {
	return r.db.Create(boiler).Error
}

func (r *BoilerPostgresRepository) Delete(id *string) error {
	return r.db.Delete(&domain.Boiler{}, "id = ?", *id).Error
}

func (r *BoilerPostgresRepository) Update(boiler *domain.Boiler) error {
	return r.db.Save(boiler).Error
}

func (r *BoilerPostgresRepository) One(id *string) (*domain.Boiler, error) {
	var boiler domain.Boiler
	if err := r.db.First(&boiler, "id = ?", *id).Error; err != nil {
		return nil, err
	}
	return &boiler, nil
}

func (r *BoilerPostgresRepository) List() ([]*domain.Boiler, error) {
	var boilers []*domain.Boiler
	if err := r.db.Find(&boilers).Error; err != nil {
		return nil, err
	}
	return boilers, nil
}

func (r *BoilerPostgresRepository) SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*domain.Boiler, error) {
	offset := (pageNumber - 1) * pageSize
	var boilers []*domain.Boiler
	if err := r.db.Where("message ILIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).
		Order("message").
		Limit(pageSize).
		Offset(offset).
		Find(&boilers).Error; err != nil {
		return nil, err
	}
	return boilers, nil
}
