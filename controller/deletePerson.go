package handler

import (
	"fmt"
	"net/http"

	"github.com/DaniloFaraum/go-crud-api/schemas"
	"github.com/DaniloFaraum/go-crud-api/utils"
	"github.com/gin-gonic/gin"
)

func DeletePersonHandler(ctx *gin.Context) {
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

	if err := db.Delete(&person).Error; err!=nil{
		logger.Errorf("could not delete person: %v", err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("could not delete person with id: %s", id))
		return
	}

	sendSucess(ctx, "delete-person", utils.ConvertToPersonResponse(person))
}