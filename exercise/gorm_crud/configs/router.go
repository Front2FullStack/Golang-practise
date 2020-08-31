package configs

import (
	"master_go_programming/Golang-practise/exercise/gorm_crud/helpers"
	"master_go_programming/Golang-practise/exercise/gorm_crud/models"
	"master_go_programming/Golang-practise/exercise/gorm_crud/repositories"
	"master_go_programming/Golang-practise/exercise/gorm_crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(contactRepository *repositories.ContactRepository) *gin.Engine {
	route := gin.Default()
	//create route
	route.POST("/create", func(context *gin.Context) {

		//initialization contact model
		var contact models.Contact

		//validate json
		err := context.ShouldBindJSON(&contact)

		//validation errors
		if err != nil {
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)
			return
		}

		//default status code = 200
		code := http.StatusOK

		//save contact & get it's response
		response := services.CreateContact(&contact, *contactRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)

	})

	return route
}
