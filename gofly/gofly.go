package gofly

import (
	"net/http"
)

// HandlerFunc defines the request handler used by gofly
type HandlerFunc func(c *Context)

// Engine implement the interface of Handler
type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	engine.router.handler(context)
}

// New is a constructor of gofly.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute defines the methods to add handler router container
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// RUN defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
