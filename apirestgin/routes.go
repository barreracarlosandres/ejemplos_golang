package apirestgin

import (
	"github.com/gin-gonic/gin"
)

func configRoutes(router *gin.Engine) {

	router.GET("/person", handleGetPeople)
	router.POST("/person", handleAddPerson)
	router.DELETE("/person", handleDeletePerson)
	router.GET("/person/:uuid", handleGetPersonByID)

}
