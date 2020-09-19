package main

import (
	"fmt"
	"github.com/gofly-dev/gofly/gofly"
	"log"
	"net/http"
)

func main() {

	r := gofly.New()
	r.GET("/index", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	fmt.Println("v1", v1)
	{
		v1.GET("/", func(c *gofly.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gofly</h1>")
		})

		v1.GET("/hello", func(c *gofly.Context) {
			// expect /hello?name=hui
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gofly.Context) {
			// expect /hello/hui
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gofly.Context) {
			c.JSON(http.StatusOK, gofly.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	log.Fatal(r.Run(":8081"))
}
