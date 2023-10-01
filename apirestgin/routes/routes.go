package routes

import (
	controller "cards/apirestgin/controllers"

	"github.com/gin-gonic/gin"
)

func Config(router *gin.Engine) {
	router.GET("/people", controller.GetAlbums)
	router.POST("/people", controller.PostAlbums)
	router.GET("/people/:id", controller.GetAlbumByID)
}
