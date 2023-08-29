package server_v2

import (
	"fmt"
	"strings"
)

// 用来支持对路由树的操作
// 代表路由树
type router struct {
	//http method => 路由树根节点
	trees map[string]*node
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}

func (r *router) addRoute(method, path string, handleFunc HandleFunc) {
	if path == "" {
		panic("路由是空字符串")
	}
	if path[0] != '/' {
		panic("路由必须以/开头")
	}
	if path[0] != '/' && path[len(path)-1] == '/' {
		panic("路由不能以 / 结尾")
	}

	//首先找到树来
	root, ok := r.trees[method]
	if !ok {
		//说明还没有根节点
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}

	//重复问题
	if path == "/" {
		if root.handler != nil {
			panic("路由冲突")
		}
		root.handler = handleFunc
		return
	}

	segs := strings.Split(path[1:], "/")
	for _, seg := range segs {
		if seg == "" {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		//递归下去，找准位置
		//如果中途节点不存在，你就要创建出来
		root = root.childOrCreate(seg)
	}
	if root.handler != nil {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}
	root.handler = handleFunc
}

func (r *router) findRoute(method, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}
	seqs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seq := range seqs {
		root, ok = root.childOf(seq)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

type node struct {
	path string
	//子 path到子节点的映射
	children map[string]*node
	//缺一个代表用户注册的业务逻辑
	handler HandleFunc
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	res, ok := n.children[path]
	return res, ok
}
func (n *node) childOrCreate(seq string) *node {
	if n.children == nil {
		n.children = make(map[string]*node, 1)
	}
	res, ok := n.children[seq]
	if !ok {
		res = &node{
			path: seq,
		}
		n.children[seq] = res
	}
	return res
}
