package handler

import (
	"net/http"

	"github.com/DaniloFaraum/go-crud-api/schemas"
	"github.com/DaniloFaraum/go-crud-api/utils"
	"github.com/gin-gonic/gin"
)

func ListPersonHandler(ctx *gin.Context) {
	var persons []schemas.Person

	if err := db.Find(&persons).Error; err!=nil{
		logger.Errorf("could not find persons: %v", err)
		sendError(ctx, http.StatusInternalServerError, "could not get persons")
		return
	}
	sendSucess(ctx, "list-persons", utils.ConvertToPersonResponses(persons))
}