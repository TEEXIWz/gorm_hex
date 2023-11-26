package service

import (
	"gorm_hex/model"
	"gorm_hex/repository"
)

// service struct (<name>Serv)
type searchServ struct{}

// service contructor (New<name>Service)
func NewSearchService() SearchService {
	return searchServ{}
}

// create repositories
var landmarkRepo = repository.NewLandmarkRepository()

// service interface (<name>Service)
type SearchService interface {
	GetAllLandmarks() ([]model.Landmark, error)
	GetLandmarksByName(key string) ([]model.Landmark, error)
}

func (d searchServ) GetAllLandmarks() ([]model.Landmark, error) {

	landmarks, err := landmarkRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}

func (d searchServ) GetLandmarksByName(key string) ([]model.Landmark, error) {

	landmarks, err := landmarkRepo.FindByName(key)
	if err != nil {
		return nil, err
	}
	return landmarks, nil
}
