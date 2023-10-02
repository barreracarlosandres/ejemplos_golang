package controller

import (
	"fmt"
	"net/http"
)

type product struct {
	ID   int
	name string
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo desde controller en home")
}

func Products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, product{ID: 1, name: "prueba"})
	fmt.Fprintln(w, "Hola mundo desde controller con productos")
}
