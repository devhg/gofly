package main

import (
	"fmt"
	"github.com/gofly-dev/gofly/gofly"
	"html/template"
	"log"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	r := gofly.Default()

	// 添加模板渲染函数
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	// 加载模板
	r.LoadHtmlGlob("templates/*")

	r.GET("/", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gofly.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gofly.H{
			"title": "gofly demo",
			"stuArr": []*Student{
				{Name: "dev", Age: 11},
				{Name: "hui", Age: 22},
			},
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gofly.Context) {
			c.HTML(http.StatusOK, "test_func.tmpl", gofly.H{
				"title": "test dateFormat",
				"now":   time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local),
			})
		})

		v1.GET("/hello", func(c *gofly.Context) {
			// expect /hello?name=hui
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	//v2.Use(middlewareForV1())
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

	r.GET("/panic", func(c *gofly.Context) {
		names := []string{"devhui"}
		c.String(http.StatusOK, names[100])
	})

	r.Static("/assets", "./static")

	log.Fatal(r.Run(":8081"))
}
