package main

import (
	"github.com/gofly-dev/gofly/gofly"
	"log"
	"net/http"
	"time"
)

func middlewareForV1() gofly.HandlerFunc {
	return func(c *gofly.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gofly.New()
	r.Use(gofly.Logger)

	r.GET("/index", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
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
	v2.Use(middlewareForV1())
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
