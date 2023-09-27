package controllers

import (

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/people", getAlbums)
	router.POST("/people", postAlbums)
	router.GET("/people/:id", getAlbumByID)
}
