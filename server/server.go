package server

import "net/http"

type Server struct {
	http.Handler
}
