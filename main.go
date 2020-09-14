package main

import (
	"github.com/gofly-dev/gofly/gofly"
	"log"
	"net/http"
)

func main() {
	engine := gofly.New()

	r := gofly.New()
	r.GET("/", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gofly</h1>")
	})

	r.GET("/hello", func(c *gofly.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gofly.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gofly.Context) {
		c.JSON(http.StatusOK, gofly.H{"filepath": c.Param("filepath")})
	})

	/*
		engine.GET("/", func(c *gofly.Context) {
			c.HTML(http.StatusOK, "<h1>你好，世界</h1>")
		})

		engine.GET("/hello", func(c *gofly.Context) {
			for k, v := range c.Req.Header {
				c.String(http.StatusOK, "Header[%q] = %q\n", k, v)
			}
		})

		engine.GET("/json", func(c *gofly.Context) {
			m := make(map[string]interface{})
			m["data"] = "sasdas"
			m["code"] = 200
			m["msg"] = "成功"
			c.JSON(http.StatusOK, m)
		})

		engine.GET("/html", func(c *gofly.Context) {
			for k, v := range c.Req.Header {
				c.String(http.StatusOK, "Header[%q] = %q\n", k, v)
			}
			c.HTML(http.StatusOK, "<p>hello world</p>")
		})

		engine.GET("/query", func(c *gofly.Context) {
			query := c.Query("key")
			c.Data(http.StatusOK, []byte(query))
		})

		engine.POST("/postForm", func(c *gofly.Context) {
			name := c.PostForm("name")
			c.String(http.StatusOK, "test postForm name=%s", name)
		})
	*/
	log.Fatal(engine.Run(":8081"))

}
