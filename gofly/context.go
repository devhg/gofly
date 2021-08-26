package gofly

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request into
	Path string
	// 动态route
	Params map[string]string
	Method string

	// response info
	StatusCode int

	// 中间件
	handlers []HandlerFunc
	index    int

	engine *Engine
}

// create a new context obj
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

// get form content from the request
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// get query params from the request
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Param(key string) string {
	return c.Params[key]
}

// set StatusCode of Context obj
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// set Header of the responseWriter
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// response the string type
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// response the json type
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json;charset=utf-8")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// response the other type data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.Writer.Write(data)
}

// response the Html type
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html;charset=utf-8")
	c.Status(code)
	err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data)

	if err != nil {
		c.Fail(http.StatusInternalServerError, err.Error())
	}
}

// fail to response
func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

// pass to next handler 巧妙的设计
func (c *Context) Next() {
	c.index++
	len := len(c.handlers)
	for ; c.index < len; c.index++ {
		c.handlers[c.index](c)
	}
}
