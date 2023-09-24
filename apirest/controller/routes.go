package controller

import (
	"fmt"
	"net/http"
)

func Home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo desde controller en home")
}

func Productos (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo desde controller con productos")
}