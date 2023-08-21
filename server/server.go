package server

import (
	"net"
	"net/http"
)

var _ Server = &HTTPServer{}

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
	Start(addr string) error
	// AddRoute 路由注册功能
	// method 是 HTTP 方法
	// path 是路由
	// handleFunc 是你的业务逻辑
	AddRoute(method string, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
	// addr string 创建的时候传递，而不是 Start 接收。这个都是可以的
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {

}

func (h *HTTPServer) Start(addr string) error {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 在这里，可以让用户注册所谓的 after start 回调
	// 比如说往你的 admin 注册一下自己这个实例
	// 在这里执行一些你业务所需的前置条件

	return http.Serve(listen, h)
}

func (h *HTTPServer) Post(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodPost, path, handleFunc)
}
func (h *HTTPServer) Get(path string, handleFunc HandleFunc) {
	h.AddRoute(http.MethodGet, path, handleFunc)
}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	//TODO implement me
	//panic("implement me")
}
