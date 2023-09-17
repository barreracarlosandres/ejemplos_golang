package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func GetHttpExample1(url string) int {
	fmt.Println("------------ Http Example 1 ------------")
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

func GetHttpExample2(url string) int {
	fmt.Println("------------ Http Example 2 ------------")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
	return resp.StatusCode
}

func GetHttpExample3(url string) int {
	fmt.Println("------------ Http Example 3 ------------")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	return resp.StatusCode
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}