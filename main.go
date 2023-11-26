package main

import (
	"gorm_hex/repository"
	"log"
)

func main() {
	_, err := repository.NewDatabaseConnection()
	if err != nil {
		panic(err)
	}
	log.Println("Connection Success")

	// searchServ := service.NewSearchService()
	// landmarks, err := searchServ.GetAllLandmarks()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, landmark := range landmarks {
	// 	log.Printf("%v %v", landmark.Name, landmark.Country.Name)
	// }

	// landmarks, err := searchServ.GetLandmarksByName("แก่ง")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, landmark := range landmarks {
	// 	log.Printf("%v %v", landmark.Name, landmark.Country.Name)
	// }
	StartServer()

}
