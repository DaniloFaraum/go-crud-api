package handler

import (
	"net/http"
	"github.com/DaniloFaraum/go-crud-api/utils"
	"github.com/DaniloFaraum/go-crud-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreatePersonHandler(ctx *gin.Context) {
	request := CreatePersonRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	person := schemas.Person{
		Name:      request.Name,
		Age:       request.Age,
		Gender:    request.Gender,
		Ocupation: request.Ocupation,
		Salary:    request.Salary,
		Alive:     *request.Alive,
	}

	if err := db.Create(&person).Error; err != nil {
		logger.Errorf("person request could not be created: %v", err)
		sendError(ctx, http.StatusInternalServerError, "could not create entry on db")
		return
	}

	sendSucess(ctx, "create-person", utils.ConvertToPersonResponse(person))
}
