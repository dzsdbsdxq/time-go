package server

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	h := &HTTPServer{}
	h.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
		fmt.Println("处理第一件事")
		fmt.Println("处理第二件事")
	})
	handle1 := func(ctx *Context) {
		fmt.Println("处理第一件事")
	}
	handle2 := func(ctx *Context) {
		fmt.Println("处理第二件事")
	}

	h.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
		handle1(ctx)
		handle2(ctx)
	})

	h.Post("/uu", func(ctx *Context) {

	})

	h.Get("/uget", func(ctx *Context) {

	})

}
