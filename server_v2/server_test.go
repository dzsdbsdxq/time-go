package server_v2

import (
	"testing"
)

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.Get("/", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello, world"))
	})
	s.Get("/user", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello, user"))
	})
	s.Get("/order/detail", func(ctx *Context) {
		ctx.Resp.Write([]byte("hello, order detail"))
	})

	s.Start(":8081")
	//h.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
	//	fmt.Println("处理第一件事")
	//	fmt.Println("处理第二件事")
	//})
	//handle1 := func(ctx *Context) {
	//	fmt.Println("处理第一件事")
	//}
	//handle2 := func(ctx *Context) {
	//	fmt.Println("处理第二件事")
	//}
	//
	//h.AddRoute(http.MethodGet, "/user", func(ctx *Context) {
	//	handle1(ctx)
	//	handle2(ctx)
	//})
	//
	//h.Post("/uu", func(ctx *Context) {
	//
	//})
	//
	//h.Get("/uget", func(ctx *Context) {
	//
	//})
	//
	//h.Start(":8080")

}
