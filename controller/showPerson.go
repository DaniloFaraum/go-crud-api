package handler

import (
	"fmt"
	"net/http"

	"github.com/DaniloFaraum/go-crud-api/schemas"
	"github.com/DaniloFaraum/go-crud-api/utils"
	"github.com/gin-gonic/gin"
)

func ShowPersonHandler(ctx *gin.Context) {
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

	sendSucess(ctx, "show-person", utils.ConvertToPersonResponse(person))
}