package handler

import (
	"fmt"
	"net/http"

	"github.com/DaniloFaraum/go-crud-api/schemas"
	"github.com/DaniloFaraum/go-crud-api/utils"
	"github.com/gin-gonic/gin"
)

func UpdatePersonHandler(ctx *gin.Context) {
	request := UpdatePersonRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == ""{
		sendError(ctx, http.StatusBadRequest, errParamIsrequired("id","queryParam").Error())
		return
	}

	person := schemas.Person{}

	if err := db.First(&person, id).Error; err!=nil{
		logger.Errorf("could not find person: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("person with id: %s not found", id))
		return
	}

	if request.Name != "" {
		person.Name = request.Name
	}

	if request.Gender != "" {
		person.Gender = request.Gender
	}

	if request.Ocupation != "" {
		person.Ocupation = request.Ocupation
	}

	if request.Alive != nil {
		person.Alive = *request.Alive
	}

	if request.Age > 0 {
		person.Age = request.Age
	}

	if request.Salary > 0 {
		person.Salary = request.Salary
	}

	if err := db.Save(&person).Error; err!=nil{
		logger.Errorf("could not update person: %v", err)
		sendError(ctx, http.StatusInternalServerError, "could not save update in server")
	}
	sendSucess(ctx, "update-person", utils.ConvertToPersonResponse(person))
}