package win

import (
	"log"
	"net/http"
)

// HandlerFunc 用来定义路由映射的处理方法， 将http.ResponseWriter, *http.Request封装为context
type HandlerFunc func(ctx *Context)

// Engine 提供了一个router表来进行 key到方法的映射
type Engine struct {
	router *router
}

// New Win的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 本地函数，根据请求的方式(method)和模式(pattern)来得到对应的函数
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

// GET 添加GET请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动http服务，监听相应的端口
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
	/*
		第二个参数的类型是什么呢？通过查看net/http的源码可以发现，Handler是一个接口，需要实现方法 ServeHTTP ，
		也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了。
	*/
}

// ServeHTTP 实现http方法，则http.ListenAndServe可以接收engine参数。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
