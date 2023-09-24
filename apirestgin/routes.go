package apirestgin

import (
	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	router := gin.Default()
	router.GET("/people", getAlbums)
	router.POST("/people", postAlbums)
	router.GET("/people/:id", getAlbumByID)

	return router
}
