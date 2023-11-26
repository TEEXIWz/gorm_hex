package repository

import (
	"gorm_hex/model"

	"gorm.io/gorm"
)

// repository struct (<name>DB)
type landmarkRepo struct {
	db *gorm.DB
}

// repository contructor (New<name>Repository)
func NewLandmarkRepository() LandmarkRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}
	return landmarkRepo{db: db}
}

// repository interface (<name>Repository)
type LandmarkRepository interface {
	FindAll() ([]model.Landmark, error)
	FindByName(key string) ([]model.Landmark, error)
}

// implement functions conformed interface
func (l landmarkRepo) FindAll() ([]model.Landmark, error) {
	landmarks := []model.Landmark{}
	result := l.db.Joins("Country").Find(&landmarks)
	if result.Error != nil {
		return nil, result.Error
	}
	return landmarks, nil
}

func (l landmarkRepo) FindByName(key string) ([]model.Landmark, error) {
	landmarks := []model.Landmark{}
	result := l.db.Joins("Country").Where("landmark.name like ?", "%"+key+"%").Find(&landmarks)
	if result.Error != nil {
		return nil, result.Error
	}
	return landmarks, nil
}
