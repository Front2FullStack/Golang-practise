package main

import (
	"log"
	"master_go_programming/Golang-practise/exercise/gorm_crud/configs"
	"master_go_programming/Golang-practise/exercise/gorm_crud/database"
	"master_go_programming/Golang-practise/exercise/gorm_crud/models"
	"master_go_programming/Golang-practise/exercise/gorm_crud/repositories"
)

func main() {
	// database configs
	dbUser, dbPassword, dbName := "root", "", "gorm_crud"
	db, err := database.ConnectToDB(dbUser, dbPassword, dbName)

	// unable to connect db
	if err != nil {
		log.Fatalln(err)
	}

	//ping to database
	err = db.DB().Ping()

	//error ping to database
	if err != nil {
		log.Fatalln(err)
	}

	//migrations
	db.AutoMigrate(&models.Contact{})

	defer db.Close()

	contactRepository := repositories.NewContactRepository(db)

	route := configs.SetupRoutes(contactRepository)

	route.Run(":8282")
}
