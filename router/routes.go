package router

import (
	"github.com/DaniloFaraum/go-crud-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1") //Creates an endpoint group "api/version", with all the CRUD operations
	{
		v1.GET("/person", handler.ShowPersonHandler)

		v1.POST("/person", handler.CreatePersonHandler)

		v1.PUT("/person", handler.UpdatePersonHandler)

		v1.DELETE("/person", handler.DeletePersonHandler)

		v1.GET("/persons", handler.ListPersonHandler)
	}
}
