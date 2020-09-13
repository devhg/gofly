package gofly

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// addRoute defines the methods to add handler router container
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle the request with the router table
func (r *router) handler(c *Context) {
	key := c.Method + "-" + c.Path

	if handlerFunc, ok := r.handlers[key]; ok {
		handlerFunc(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
