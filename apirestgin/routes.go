package apirestgin

import (
	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine) {

	router.GET("/person", handleGetAll)
	router.POST("/person", handleAddPerson)
	router.DELETE("/person/:uuid", handleDeletePerson)
	router.GET("/person/:uuid", handleGetPersonByID)
}
