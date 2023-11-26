package repository

import (
	"gorm_hex/model"

	"gorm.io/gorm"
)

// repository struct (<name>DB)
type countryRepo struct {
	db *gorm.DB
}

// repository contructor (New<name>Repository)
func NewCountryRepository() CountryRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return countryRepo{db: db}
}

// repository interface (<name>Repository)
type CountryRepository interface {
	FindAll() ([]model.Country, error)
}

// implement functions conformed interface
func (l countryRepo) FindAll() ([]model.Country, error) {
	countries := []model.Country{}
	result := l.db.Find(&countries)
	if result.Error != nil {
		return nil, result.Error
	}
	return countries, nil
}
