package gofly

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder

	str.WriteString(message + "\nTraceback:")

	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// 优先注册
func Recovery() HandlerFunc {
	return func(c *Context) {
		//c.Next()
		// 如果panic在 defer的上方的话，下面defer recover的将不会被开启，必须将defer recover()放到panic的上方
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

func RecoveryT() {
	fmt.Println(123)
	defer func() {
		fmt.Println(2333)
		//if err := recover(); err != nil {
		//	message := fmt.Sprintf("%s", err)
		//	log.Printf("%s\n\n", trace(message))
		//
		//}
	}()
}
