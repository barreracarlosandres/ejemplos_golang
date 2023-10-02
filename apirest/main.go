package apirest

import (
	"cards/apirest/routes"
	"fmt"
	"net/http"
)

func Example() {
	routes.Routes()
	fmt.Println("Running mvc")
	http.ListenAndServe(":8080", nil)
}
