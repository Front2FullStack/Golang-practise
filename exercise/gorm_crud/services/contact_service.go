package services

import (
	"log"
	"master_go_programming/Golang-practise/exercise/gorm_crud/dtos"
	"master_go_programming/Golang-practise/exercise/gorm_crud/models"
	"master_go_programming/Golang-practise/exercise/gorm_crud/repositories"

	"github.com/google/uuid"
)

func CreateContact(contact *models.Contact, repository repositories.ContactRepository) dtos.Response {
	uuidResult, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln(err)
	}
	contact.ID = uuidResult.String()

	operationResult := repository.Save(contact)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Contact)
	return dtos.Response{Success: true, Data: data}
}
