package mvc

import (
	"cards/mvc/routes"
	"fmt"
	"net/http"
)

func RunApiRest() {
	routes.Routes()
	fmt.Println("Running mvc")
	http.ListenAndServe(":3000", nil)
}
