package routes

import (
	"net/http"
	"cards/mvc/controller"
)

func Routes () {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/productos", controller.Productos)
}