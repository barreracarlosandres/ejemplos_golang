package http

import (
	"fmt"
	"net/http"
	"os"
)

func GetHttp() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}