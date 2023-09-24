package routes

import (
	"cards/apirest/controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/productos", controller.Productos)
}
