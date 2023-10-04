package apirestgin

import (
	"github.com/gin-gonic/gin"
)

func Example() {
	router := gin.Default()
	configRoutes(router)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
