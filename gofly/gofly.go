package gofly

import (
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gofly
type HandlerFunc func(c *Context)

// Engine implement the interface of Handler
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	engine.router.handler(context)
}

// New is a constructor of gofly.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}             // 先创建一个engine
	engine.RouterGroup = &RouterGroup{engine: engine}  // 把已经创建的引擎定义为顶级RouterGroup
	engine.groups = []*RouterGroup{engine.RouterGroup} // 把自己的RouterGroup 存进groups
	return engine
}

// RUN defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// addRoute defines the methods to add handler router container
func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
