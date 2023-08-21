package server

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var h Server
	http.ListenAndServe(":8081", h)
}
