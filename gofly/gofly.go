package gofly

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc defines the request handler used by gofly
type HandlerFunc func(c *Context)

// Engine implement the interface of Handler
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup

	htmlTemplates *template.Template // 将模板加载进内存
	funcMap       template.FuncMap   // 自定义模板渲染函数
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	context := NewContext(w, r)
	context.handlers = middlewares
	context.engine = engine
	engine.router.handler(context)
}

// New is a constructor of gofly.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}             // 先创建一个engine
	engine.RouterGroup = &RouterGroup{engine: engine}  // 把已经创建的引擎定义为顶级RouterGroup
	engine.groups = []*RouterGroup{engine.RouterGroup} // 把自己的RouterGroup 存进groups
	return engine
}

// Default is a constructor of default engine
func Default() *Engine {
	engine := New()
	engine.Use(Logger, Recovery())
	return engine
}

// RUN defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

//加载模板
func (engine *Engine) LoadHtmlGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// Group create a route group
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

// Use defines the methods to add a middlewares to the RouterGroup
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

// Static defines the methods to serve static files
func (group *RouterGroup) Static(relativePath, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	path := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(path, http.FileServer(fs))
	return func(c *Context) {
		file := c.Params["filepath"]
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

//设置自定义渲染函数map
func (group *RouterGroup) SetFuncMap(funcMap template.FuncMap) {
	group.engine.funcMap = funcMap
}
