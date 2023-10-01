package apirestgin

import (
	"cards/apirestgin/routes"

	"github.com/gin-gonic/gin"
)

func Example() {
	router := gin.Default()
	routes.Config(router)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
