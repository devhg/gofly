package gofly

import (
	"fmt"
	"log"
	"time"
)

// print color of terminal
func Color() {
	for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func Logger(c *Context) {
	now := time.Now()
	// process request
	c.Next()
	gofly := fmt.Sprintf("%c[%d;%dm%s%c[0m", 0x1B, 0, 35, "gofly", 0x1B)
	log.Printf("%s -- [%d] %s in %v\n", gofly, c.StatusCode, c.Req.RequestURI, time.Since(now))
}
