package gofly

/**
the first beginning version of my web framework
*/
//
//import (
//	"fmt"
//	"net/http"
//)
//
//// HandlerFunc defines the request handler used by gofly
//type HandlerFunc func(http.ResponseWriter, *http.Request)
//
//// Engine implement the interface of Handler
//type Engine struct {
//	router map[string]HandlerFunc
//}
//
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	key := r.Method + "-" + r.URL.Path
//
//	if handler, ok := engine.router[key]; ok {
//		handler(w, r)
//	} else {
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
//	}
//}
//
//// New is a constructor of gofly.Engine
//func New() *Engine {
//	return &Engine{router: make(map[string]HandlerFunc)}
//}
//
//// addRoute defines the methods to add handler router container
//func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
//	key := method + "-" + pattern
//	engine.router[key] = handler
//}
//
//// RUN defines the method to start a http server
//func (engine *Engine) Run(addr string) error {
//	return http.ListenAndServe(addr, engine)
//}
//
//// GET defines the method to add GET request
//func (engine *Engine) GET(pattern string, handler HandlerFunc) {
//	engine.addRoute("GET", pattern, handler)
//}
//
//// POST defines the method to add POST request
//func (engine *Engine) POST(pattern string, handler HandlerFunc) {
//	engine.addRoute("POST", pattern, handler)
//}
