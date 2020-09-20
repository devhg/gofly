# Web FrameWork

*It like [gin](https://github.com/gin-gonic/gin)*

### What's support

- [x] 支持上下文及多格式返回 Context
- [x] 动态路由 /doc/:lang
- [x] 分组控制路由 Group
- [x] 静态文件 StaticFS
- [x] 支持中间件 middleware
- [x] 支持模板引擎 html/templates
- [x] 支持错误恢复 Panic recover

### How to use

just like
```golang
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
```

<hr>
Just for learning. 奥利给！！！