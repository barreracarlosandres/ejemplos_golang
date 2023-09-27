package apirestgin

import (
	"cards/apirestgin/controllers"

	"github.com/gin-gonic/gin"
)

func Example() {
	router := gin.Default()
	controllers.Routes(router)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
