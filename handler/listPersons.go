package handler

import "github.com/gin-gonic/gin"

func ListPersonHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "placeholder",
	})
}