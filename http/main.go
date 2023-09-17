package http

import (
	"fmt"
	"net/http"
	"os"
)

func GetHttp(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	// to read all body data in a slace
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println("************* head ************")
	fmt.Println(resp)
	fmt.Println("************* body ************")
	fmt.Println(string(bs))
	return resp.StatusCode
}