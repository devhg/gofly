package main

import (
	"github.com/gofly-dev/gofly/gofly"
	"log"
	"net/http"
)

func main() {

	r := gofly.New()
	r.GET("/", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gofly</h1>")
	})

	r.GET("/hello/:name", func(c *gofly.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/hello/:name/age", func(c *gofly.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gofly.Context) {
		c.JSON(http.StatusOK, gofly.H{"filepath": c.Param("filepath")})
	})

	log.Fatal(r.Run(":8081"))
}
