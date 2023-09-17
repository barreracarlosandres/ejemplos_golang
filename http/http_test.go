package http_test

import (
	"cards/http"
	"testing"
)


func TestGetHttp(t *testing.T) {

	gh := http.GetHttp("http://google.com")

	if gh != 200 {
		t.Errorf("Expected Status Code 200, but got %v:", gh)
	}
}