package gofly

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	split := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range split {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute defines the methods to add handler router container
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {

	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern

	// Trie树添加动态路由部分
	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
	// Trie树获取动态路由部分
	pathParts := parsePattern(path) // 请求传入的path /go/cn/doc
	params := make(map[string]string)

	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(pathParts, 0)
	if n != nil {
		searchParts := parsePattern(n.pattern)
		for index, part := range searchParts {
			if part[0] == ':' {
				params[part[1:]] = pathParts[index] // params["lang"] = cn
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(pathParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

// handle the request with the router table
func (r *router) handler(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
