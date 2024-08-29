package handler

import "github.com/gin-gonic/gin"

func CreatePersonHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "placeholder",
	})
}